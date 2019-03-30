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

package debt_ceiling

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	repo "github.com/vulcanize/vulcanizedb/libraries/shared/repository"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
)

type VatFileDebtCeilingRepository struct {
	db *postgres.DB
}

func (repository VatFileDebtCeilingRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Beginx()
	if dBaseErr != nil {
		return dBaseErr
	}

	for _, model := range models {
		vatFileDC, ok := model.(VatFileDebtCeilingModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, VatFileDebtCeilingModel{})
		}

		_, execErr := tx.Exec(
			`INSERT into maker.vat_file_debt_ceiling (header_id, what, data, log_idx, tx_idx, raw_log)
        VALUES($1, $2, $3::NUMERIC, $4, $5, $6)
		ON CONFLICT (header_id, tx_idx, log_idx) DO UPDATE SET what = $2, data = $3, raw_log = $6;`,
			headerID, vatFileDC.What, vatFileDC.Data, vatFileDC.LogIndex, vatFileDC.TransactionIndex, vatFileDC.Raw,
		)

		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := repo.MarkHeaderCheckedInTransaction(headerID, tx, constants.VatFileDebtCeilingChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}

	return tx.Commit()
}

func (repository VatFileDebtCeilingRepository) MarkHeaderChecked(headerID int64) error {
	return repo.MarkHeaderChecked(headerID, repository.db, constants.VatFileDebtCeilingChecked)
}

func (repository VatFileDebtCeilingRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	return repo.MissingHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.VatFileDebtCeilingChecked)
}

func (repository VatFileDebtCeilingRepository) RecheckHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	return repo.RecheckHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.VatFileDebtCeilingChecked)
}

func (repository *VatFileDebtCeilingRepository) SetDB(db *postgres.DB) {
	repository.db = db
}