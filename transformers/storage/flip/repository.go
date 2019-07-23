package flip

import (
	"fmt"

	"github.com/vulcanize/vulcanizedb/libraries/shared/storage/utils"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
)

const (
	insertFlipVatQuery   = `INSERT INTO maker.flip_vat (block_number, block_hash, contract_address, vat) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	insertFlipIlkQuery   = `INSERT INTO maker.flip_ilk (block_number, block_hash, contract_address, ilk) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	insertFlipBegQuery   = `INSERT INTO maker.flip_beg (block_number, block_hash, contract_address, beg) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	insertFlipTtlQuery   = `INSERT INTO maker.flip_ttl (block_number, block_hash, contract_address, ttl) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	insertFlipTauQuery   = `INSERT INTO maker.flip_tau (block_number, block_hash, contract_address, tau) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	InsertFlipKicksQuery = `INSERT INTO maker.flip_kicks (block_number, block_hash, contract_address, kicks) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`

	InsertFlipBidBidQuery = `INSERT INTO maker.flip_bid_bid (block_number, block_hash, contract_address, bid_id, bid) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidLotQuery = `INSERT INTO maker.flip_bid_lot (block_number, block_hash, contract_address, bid_id, lot) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidGuyQuery = `INSERT INTO maker.flip_bid_guy (block_number, block_hash, contract_address, bid_id, guy) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidTicQuery = `INSERT INTO maker.flip_bid_tic (block_number, block_hash, contract_address, bid_id, tic) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidEndQuery = `INSERT INTO maker.flip_bid_end (block_number, block_hash, contract_address, bid_id, "end") VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidUsrQuery = `INSERT INTO maker.flip_bid_usr (block_number, block_hash, contract_address, bid_id, usr) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidGalQuery = `INSERT INTO maker.flip_bid_gal (block_number, block_hash, contract_address, bid_id, gal) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	InsertFlipBidTabQuery = `INSERT INTO maker.flip_bid_tab (block_number, block_hash, contract_address, bid_id, tab) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
)

type FlipStorageRepository struct {
	ContractAddress string
	db              *postgres.DB
}

func (repository *FlipStorageRepository) Create(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, value interface{}) error {
	switch metadata.Name {
	case Vat:
		return repository.insertVat(blockNumber, blockHash, value.(string))
	case Ilk:
		return repository.insertIlk(blockNumber, blockHash, value.(string))
	case Beg:
		return repository.insertBeg(blockNumber, blockHash, value.(string))
	case Ttl:
		return repository.insertTtl(blockNumber, blockHash, value.(string))
	case Tau:
		return repository.insertTau(blockNumber, blockHash, value.(string))
	case Kicks:
		return repository.insertKicks(blockNumber, blockHash, value.(string))
	case BidBid:
		return repository.insertBidBid(blockNumber, blockHash, metadata, value.(string))
	case BidLot:
		return repository.insertBidLot(blockNumber, blockHash, metadata, value.(string))
	case BidGuy:
		return repository.insertBidGuy(blockNumber, blockHash, metadata, value.(string))
	case BidTic:
		return repository.insertBidTic(blockNumber, blockHash, metadata, value.(string))
	case BidEnd:
		return repository.insertBidEnd(blockNumber, blockHash, metadata, value.(string))
	case BidUsr:
		return repository.insertBidUsr(blockNumber, blockHash, metadata, value.(string))
	case BidGal:
		return repository.insertBidGal(blockNumber, blockHash, metadata, value.(string))
	case BidTab:
		return repository.insertBidTab(blockNumber, blockHash, metadata, value.(string))
	case Packed:
		return repository.insertPackedValueRecord(blockNumber, blockHash, metadata, value.(map[int]string))
	default:
		panic(fmt.Sprintf("unrecognized flip contract storage name: %s", metadata.Name))
	}
}

func (repository *FlipStorageRepository) SetDB(db *postgres.DB) {
	repository.db = db
}

func (repository *FlipStorageRepository) insertVat(blockNumber int, blockHash string, vat string) error {
	_, writeErr := repository.db.Exec(insertFlipVatQuery, blockNumber, blockHash, repository.ContractAddress, vat)
	return writeErr
}

func (repository *FlipStorageRepository) insertIlk(blockNumber int, blockHash string, ilk string) error {
	_, writeErr := repository.db.Exec(insertFlipIlkQuery, blockNumber, blockHash, repository.ContractAddress, ilk)
	return writeErr
}

func (repository *FlipStorageRepository) insertBeg(blockNumber int, blockHash string, beg string) error {
	_, writeErr := repository.db.Exec(insertFlipBegQuery, blockNumber, blockHash, repository.ContractAddress, beg)
	return writeErr
}

func (repository *FlipStorageRepository) insertTtl(blockNumber int, blockHash string, ttl string) error {
	_, writeErr := repository.db.Exec(insertFlipTtlQuery, blockNumber, blockHash, repository.ContractAddress, ttl)
	return writeErr
}

func (repository *FlipStorageRepository) insertTau(blockNumber int, blockHash string, tau string) error {
	_, writeErr := repository.db.Exec(insertFlipTauQuery, blockNumber, blockHash, repository.ContractAddress, tau)
	return writeErr
}

func (repository *FlipStorageRepository) insertKicks(blockNumber int, blockHash, kicks string) error {
	_, writeErr := repository.db.Exec(InsertFlipKicksQuery, blockNumber, blockHash, repository.ContractAddress, kicks)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidBid(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, bid string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidBidQuery, blockNumber, blockHash, repository.ContractAddress, bidId, bid)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidLot(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, lot string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidLotQuery, blockNumber, blockHash, repository.ContractAddress, bidId, lot)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidGuy(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, guy string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidGuyQuery, blockNumber, blockHash, repository.ContractAddress, bidId, guy)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidTic(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, tic string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidTicQuery, blockNumber, blockHash, repository.ContractAddress, bidId, tic)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidEnd(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, end string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidEndQuery, blockNumber, blockHash, repository.ContractAddress, bidId, end)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidUsr(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, usr string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidUsrQuery, blockNumber, blockHash, repository.ContractAddress, bidId, usr)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidGal(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, gal string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidGalQuery, blockNumber, blockHash, repository.ContractAddress, bidId, gal)
	return writeErr
}

func (repository *FlipStorageRepository) insertBidTab(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, tab string) error {
	bidId, err := getBidId(metadata.Keys)
	if err != nil {
		return err
	}
	_, writeErr := repository.db.Exec(InsertFlipBidTabQuery, blockNumber, blockHash, repository.ContractAddress, bidId, tab)
	return writeErr
}

func (repository *FlipStorageRepository) insertPackedValueRecord(blockNumber int, blockHash string, metadata utils.StorageValueMetadata, packedValues map[int]string) error {
	for order, value := range packedValues {
		switch metadata.PackedNames[order] {
		case Ttl:
			ttlErr := repository.insertTtl(blockNumber, blockHash, value)
			if ttlErr != nil {
				return ttlErr
			}
		case Tau:
			tauErr := repository.insertTau(blockNumber, blockHash, value)
			if tauErr != nil {
				return tauErr
			}
		default:
			panic(fmt.Sprintf("unrecognized flip contract storage name in packed values: %s", metadata.Name))
		}
	}
	return nil
}

func getBidId(keys map[utils.Key]string) (string, error) {
	bidId, ok := keys[constants.BidId]
	if !ok {
		return "", utils.ErrMetadataMalformed{MissingData: constants.BidId}
	}
	return bidId, nil
}