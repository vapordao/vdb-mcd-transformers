// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bite

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
)

type BiteRepository struct {
	db *postgres.DB
}

func (repository *BiteRepository) SetDB(db *postgres.DB) {
	repository.db = db
}

func (repository BiteRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Begin()
	if dBaseErr != nil {
		return dBaseErr
	}
	for _, model := range models {
		biteModel, ok := model.(BiteModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, BiteModel{})
		}

		_, execErr := tx.Exec(
			`INSERT into maker.bite (header_id, ilk, urn, ink, art, iart, tab, nflip, log_idx, tx_idx, raw_log)
        VALUES($1, $2, $3, $4::NUMERIC, $5::NUMERIC, $6::NUMERIC, $7::NUMERIC, $8::NUMERIC, $9, $10, $11)`,
			headerID, biteModel.Ilk, biteModel.Urn, biteModel.Ink, biteModel.Art, biteModel.IArt, biteModel.Tab, biteModel.NFlip, biteModel.LogIndex, biteModel.TransactionIndex, biteModel.Raw,
		)
		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := shared.MarkHeaderCheckedInTransaction(headerID, tx, constants.BiteChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}

	return tx.Commit()
}

func (repository BiteRepository) MarkHeaderChecked(headerID int64) error {
	return shared.MarkHeaderChecked(headerID, repository.db, constants.BiteChecked)
}
