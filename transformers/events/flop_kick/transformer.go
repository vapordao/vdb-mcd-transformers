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

package flop_kick

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/repository"
	"github.com/makerdao/vulcanizedb/pkg/core"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
	"github.com/makerdao/vulcanizedb/pkg/eth"
)

type Transformer struct{}

func (Transformer) toEntities(contractAbi string, logs []core.EventLog) ([]FlopKickEntity, error) {
	var results []FlopKickEntity
	for _, log := range logs {
		var entity FlopKickEntity
		address := log.Log.Address
		abi, parseErr := eth.ParseAbi(contractAbi)
		if parseErr != nil {
			return nil, parseErr
		}

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		unpackErr := contract.UnpackLog(&entity, "Kick", log.Log)
		if unpackErr != nil {
			return nil, unpackErr
		}
		entity.ContractAddress = log.Log.Address
		entity.HeaderID = log.HeaderID
		entity.LogID = log.ID
		results = append(results, entity)
	}
	return results, nil
}

func (t Transformer) ToModels(abi string, logs []core.EventLog, db *postgres.DB) ([]event.InsertionModel, error) {
	var results []event.InsertionModel
	entities, entityErr := t.toEntities(abi, logs)
	if entityErr != nil {
		return nil, fmt.Errorf("FlopKick transformer couldn't convert logs to entities: %v", entityErr)
	}
	for _, flopKickEntity := range entities {
		addressId, addressErr := repository.GetOrCreateAddress(db, flopKickEntity.ContractAddress.Hex())
		if addressErr != nil {
			return nil, shared.ErrCouldNotCreateFK(addressErr)
		}

		model := event.InsertionModel{
			SchemaName: constants.MakerSchema,
			TableName:  constants.FlopKickTable,
			OrderedColumns: []event.ColumnName{
				event.HeaderFK,
				event.LogFK,
				event.AddressFK,
				constants.BidIDColumn,
				constants.LotColumn,
				constants.BidColumn,
				constants.GalColumn,
			},
			ColumnValues: event.ColumnValues{
				event.HeaderFK:        flopKickEntity.HeaderID,
				event.LogFK:           flopKickEntity.LogID,
				event.AddressFK:       addressId,
				constants.BidIDColumn: shared.BigIntToString(flopKickEntity.Id),
				constants.LotColumn:   shared.BigIntToString(flopKickEntity.Lot),
				constants.BidColumn:   shared.BigIntToString(flopKickEntity.Bid),
				constants.GalColumn:   flopKickEntity.Gal.String(),
			},
		}
		results = append(results, model)
	}

	return results, nil
}
