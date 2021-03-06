// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package trigger_test

import (
	"database/sql"
	"math/rand"

	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vdb-mcd-transformers/transformers/component_tests/queries/test_helpers"
	. "github.com/makerdao/vdb-mcd-transformers/transformers/storage/test_helpers"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/pkg/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Updating ilk_snapshot table", func() {
	var (
		blockOne,
		blockTwo int
		headerOne,
		headerTwo core.Header
		rawTimestampOne,
		rawTimestampTwo int64
		vatInitModel        event.InsertionModel
		database            = test_config.NewTestDB(test_config.NewTestNode())
		getTimeCreatedQuery = `SELECT created FROM api.ilk_snapshot ORDER BY block_number`
		insertRecordQuery   = `INSERT INTO api.ilk_snapshot (ilk_identifier, block_number, created) VALUES ($1, $2, $3::TIMESTAMP)`
		insertEmptyRowQuery = `INSERT INTO api.ilk_snapshot (ilk_identifier, block_number) VALUES ($1, $2)`
	)

	BeforeEach(func() {
		test_config.CleanTestDB(database)
		blockOne = rand.Int()
		blockTwo = blockOne + 1
		rawTimestampOne = int64(rand.Int31())
		rawTimestampTwo = rawTimestampOne + 1
		headerOne = CreateHeader(rawTimestampOne, blockOne, database)
		headerTwo = CreateHeader(rawTimestampTwo, blockTwo, database)
		vatInitModel = test_helpers.CreateVatInit(database, headerOne.Id, test_helpers.FakeIlk.Hex)
	})

	It("updates time created of all records for an ilk", func() {
		_, setupErr := database.Exec(insertEmptyRowQuery, test_helpers.FakeIlk.Identifier, headerTwo.BlockNumber)
		Expect(setupErr).NotTo(HaveOccurred())
		expectedTimeCreated := sql.NullString{Valid: true, String: FormatTimestamp(rawTimestampOne)}

		vatInitErr := event.PersistModels([]event.InsertionModel{vatInitModel}, database)
		Expect(vatInitErr).NotTo(HaveOccurred())

		var ilkSnapshots []test_helpers.IlkSnapshot
		queryErr := database.Select(&ilkSnapshots, getTimeCreatedQuery)
		Expect(queryErr).NotTo(HaveOccurred())
		Expect(len(ilkSnapshots)).To(Equal(1))
		Expect(ilkSnapshots[0].Created).To(Equal(expectedTimeCreated))
	})

	It("does not update time created if old time created is not null", func() {
		_, setupErr := database.Exec(insertRecordQuery, test_helpers.FakeIlk.Identifier, headerTwo.BlockNumber,
			FormatTimestamp(rawTimestampTwo))
		Expect(setupErr).NotTo(HaveOccurred())
		expectedTimeCreated := sql.NullString{Valid: true, String: FormatTimestamp(rawTimestampTwo)}

		vatInitErr := event.PersistModels([]event.InsertionModel{vatInitModel}, database)
		Expect(vatInitErr).NotTo(HaveOccurred())

		var ilkSnapshots []test_helpers.IlkSnapshot
		queryErr := database.Select(&ilkSnapshots, getTimeCreatedQuery)
		Expect(queryErr).NotTo(HaveOccurred())
		Expect(len(ilkSnapshots)).To(Equal(1))
		Expect(ilkSnapshots[0].Created).To(Equal(expectedTimeCreated))
	})

	It("does not update records with a different ilk", func() {
		_, setupErr := database.Exec(insertEmptyRowQuery, test_helpers.AnotherFakeIlk.Identifier, headerTwo.BlockNumber)
		Expect(setupErr).NotTo(HaveOccurred())
		expectedTimeCreated := sql.NullString{Valid: false, String: ""}

		vatInitErr := event.PersistModels([]event.InsertionModel{vatInitModel}, database)
		Expect(vatInitErr).NotTo(HaveOccurred())

		var ilkSnapshots []test_helpers.IlkSnapshot
		queryErr := database.Select(&ilkSnapshots, getTimeCreatedQuery)
		Expect(queryErr).NotTo(HaveOccurred())
		Expect(len(ilkSnapshots)).To(Equal(1))
		Expect(ilkSnapshots[0].Created).To(Equal(expectedTimeCreated))
	})

	It("sets created to null when record is deleted", func() {
		_, ilkSetupErr := database.Exec(insertEmptyRowQuery, test_helpers.FakeIlk.Identifier, headerOne.BlockNumber)
		Expect(ilkSetupErr).NotTo(HaveOccurred())
		vatInitErr := event.PersistModels([]event.InsertionModel{vatInitModel}, database)
		Expect(vatInitErr).NotTo(HaveOccurred())

		_, err := database.Exec(`DELETE FROM maker.vat_init WHERE header_id = $1`, headerOne.Id)
		Expect(err).NotTo(HaveOccurred())

		var ilkSnapshots []test_helpers.IlkSnapshot
		queryErr := database.Select(&ilkSnapshots, getTimeCreatedQuery)
		Expect(queryErr).NotTo(HaveOccurred())
		Expect(ilkSnapshots[0].Created).To(Equal(sql.NullString{Valid: false, String: ""}))
	})
})
