// VulcanizeDB
// Copyright © 2018 Vulcanize

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

package ilk_test

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/mcd_transformers/test_config"
	"github.com/vulcanize/mcd_transformers/transformers/events/vat_file/ilk"
	"github.com/vulcanize/mcd_transformers/transformers/shared"
	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
	"github.com/vulcanize/mcd_transformers/transformers/test_data"
	"github.com/vulcanize/mcd_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("Vat file ilk repository", func() {
	var (
		db                   *postgres.DB
		vatFileIlkRepository ilk.VatFileIlkRepository
		headerRepository     repositories.HeaderRepository
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		vatFileIlkRepository = ilk.VatFileIlkRepository{}
		vatFileIlkRepository.SetDB(db)
		headerRepository = repositories.NewHeaderRepository(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.VatFileIlkSpotModel
		modelWithDifferentLogIdx.LogIndex = modelWithDifferentLogIdx.LogIndex + 1
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.VatFileIlkChecked,
			LogEventTableName:        "maker.vat_file_ilk",
			TestModel:                test_data.VatFileIlkSpotModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &vatFileIlkRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("adds a vat file ilk event", func() {
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())
			err = vatFileIlkRepository.Create(headerID, []interface{}{test_data.VatFileIlkSpotModel})

			Expect(err).NotTo(HaveOccurred())
			var dbVatFile ilk.VatFileIlkModel
			err = db.Get(&dbVatFile, `SELECT ilk_id, what, data, log_idx, tx_idx, raw_log FROM maker.vat_file_ilk WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			ilkID, err := shared.GetOrCreateIlk(test_data.VatFileIlkSpotModel.Ilk, db)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbVatFile.Ilk).To(Equal(strconv.Itoa(ilkID)))
			Expect(dbVatFile.What).To(Equal(test_data.VatFileIlkSpotModel.What))
			Expect(dbVatFile.Data).To(Equal(test_data.VatFileIlkSpotModel.Data))
			Expect(dbVatFile.LogIndex).To(Equal(test_data.VatFileIlkSpotModel.LogIndex))
			Expect(dbVatFile.TransactionIndex).To(Equal(test_data.VatFileIlkSpotModel.TransactionIndex))
			Expect(dbVatFile.Raw).To(MatchJSON(test_data.VatFileIlkSpotModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.VatFileIlkChecked,
			Repository:              &vatFileIlkRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
