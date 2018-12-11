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

package flip

import (
	"fmt"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
)

type CatFileFlipRepository struct {
	db *postgres.DB
}

func (repository CatFileFlipRepository) Create(headerID int64, models []interface{}) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}
	for _, model := range models {
		flip, ok := model.(CatFileFlipModel)
		if !ok {
			tx.Rollback()
			return fmt.Errorf("model of type %T, not %T", model, CatFileFlipModel{})
		}

		_, err = repository.db.Exec(
			`INSERT into maker.cat_file_flip (header_id, ilk, what, flip, tx_idx, log_idx, raw_log)
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
			headerID, flip.Ilk, flip.What, flip.Flip, flip.TransactionIndex, flip.LogIndex, flip.Raw,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = shared.Repository{}.MarkHeaderCheckedInTransaction(headerID, tx, constants.CatFileFlipChecked)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repository CatFileFlipRepository) MarkHeaderChecked(headerID int64) error {
	return shared.Repository{}.MarkHeaderChecked(headerID, repository.db, constants.CatFileFlipChecked)
}

func (repository *CatFileFlipRepository) SetDB(db *postgres.DB) {
	repository.db = db
}
