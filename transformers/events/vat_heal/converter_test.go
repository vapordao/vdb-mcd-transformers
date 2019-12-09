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

package vat_heal_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/makerdao/vdb-mcd-transformers/transformers/events/vat_heal"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
	"github.com/makerdao/vulcanizedb/pkg/core"
)

var _ = Describe("VatHeal converter", func() {

	var (
		converter vat_heal.Converter
		db        *postgres.DB
	)

	BeforeEach(func() {
		converter = vat_heal.Converter{}
		db = test_config.NewTestDB(test_config.NewTestNode())
	})
	It("Convert log with positive rad to a model", func() {
		models, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{test_data.VatHealHeaderSyncLog}, db)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(models)).To(Equal(1))
		Expect(models[0]).To(Equal(test_data.VatHealModel))
	})

	It("Returns an error there are missing topics", func() {
		converter := vat_heal.Converter{}
		badLog := core.HeaderSyncLog{
			Log: types.Log{
				Topics: []common.Hash{
					common.HexToHash("0x"),
				}},
		}

		_, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{badLog}, db)

		Expect(err).To(HaveOccurred())
	})
})
