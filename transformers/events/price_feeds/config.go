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

package price_feeds

import (
	shared_t "github.com/vulcanize/vulcanizedb/libraries/shared/transformer"

	"github.com/vulcanize/mcd_transformers/transformers/shared"
	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
)

func GetPriceFeedConfig() shared_t.EventTransformerConfig {
	return shared_t.EventTransformerConfig{
		TransformerName: constants.PriceFeedLabel,
		ContractAddresses: []string{
			constants.PipEthContractAddress(), constants.PipRepContractAddress(),
		},
		ContractAbi: constants.MedianizerABI(),
		Topic:       constants.GetLogMedianPriceSignature(),
		StartingBlockNumber: shared.MinInt64([]int64{
			constants.PipEthDeploymentBlock(), constants.PipRepDeploymentBlock(),
		}),
		EndingBlockNumber: -1,
	}
}