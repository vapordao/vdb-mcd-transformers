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

package vat_fork_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"

	"github.com/makerdao/vdb-mcd-transformers/transformers/events/vat_fork"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
)

var _ = Describe("VatFork converter", func() {
	converter := vat_fork.VatForkConverter{}

	It("Converts a log with a negative dink and dart to a model", func() {
		models, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{test_data.VatForkHeaderSyncLogWithNegativeDinkDart})

		Expect(err).NotTo(HaveOccurred())
		Expect(models).To(Equal([]shared.InsertionModel{test_data.VatForkModelWithNegativeDinkDart}))
	})

	It("Converts a log with a positive dink and dart to a model", func() {
		models, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{test_data.VatForkHeaderSyncLogWithPositiveDinkDart})

		Expect(err).NotTo(HaveOccurred())
		Expect(models).To(Equal([]shared.InsertionModel{test_data.VatForkModelWithPositiveDinkDart}))
	})

	It("Returns an error there are missing topics", func() {
		badLog := core.HeaderSyncLog{
			Log: types.Log{
				Topics: []common.Hash{
					common.HexToHash("0x"),
					common.HexToHash("0x"),
					common.HexToHash("0x"),
				}},
		}

		_, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{badLog})
		Expect(err).To(HaveOccurred())
	})

	It("returns err if log is missing data", func() {
		badLog := core.HeaderSyncLog{
			Log: types.Log{
				Topics: []common.Hash{{}, {}, {}, {}},
			}}

		_, err := converter.ToModels(constants.VatABI(), []core.HeaderSyncLog{badLog})
		Expect(err).To(HaveOccurred())
	})
})
