package dsr

import (
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/repository"
	"github.com/makerdao/vulcanizedb/pkg/core"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
)

// Transformer implements the VDB event Transformer interface
type Transformer struct{}

// ToModels transforms log data into general InsertionModels the Repository can persist
func (Transformer) ToModels(_ string, logs []core.EventLog, db *postgres.DB) ([]event.InsertionModel, error) {
	var results []event.InsertionModel
	for _, log := range logs {
		verifyErr := shared.VerifyLog(log.Log, shared.FourTopicsRequired, shared.LogDataNotRequired)
		if verifyErr != nil {
			return nil, verifyErr
		}
		msgSender := shared.GetChecksumAddressString(log.Log.Topics[1].Hex())
		msgSenderID, msgSenderErr := repository.GetOrCreateAddress(db, msgSender)
		if msgSenderErr != nil {
			return nil, shared.ErrCouldNotCreateFK(msgSenderErr)
		}

		what := shared.DecodeHexToText(log.Log.Topics[2].Hex())
		data := shared.ConvertUint256HexToBigInt(log.Log.Topics[3].Hex())
		result := event.InsertionModel{
			SchemaName: constants.MakerSchema,
			TableName:  constants.PotFileDSRTable,
			OrderedColumns: []event.ColumnName{
				event.HeaderFK,
				event.LogFK,
				constants.MsgSenderColumn,
				constants.WhatColumn,
				constants.DataColumn,
			},
			ColumnValues: event.ColumnValues{
				event.HeaderFK:            log.HeaderID,
				event.LogFK:               log.ID,
				constants.MsgSenderColumn: msgSenderID,
				constants.WhatColumn:      what,
				constants.DataColumn:      data.String(),
			},
		}
		results = append(results, result)
	}
	return results, nil
}
