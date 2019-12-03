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

package queries

import (
	"math/rand"
	"strconv"

	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vdb-mcd-transformers/transformers/component_tests/queries/test_helpers"
	"github.com/makerdao/vdb-mcd-transformers/transformers/events/flip_kick"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/storage/vat"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
	"github.com/makerdao/vulcanizedb/pkg/core"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/makerdao/vulcanizedb/pkg/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flip state computed columns", func() {
	var (
		db                     *postgres.DB
		headerOne              core.Header
		headerRepository       repositories.HeaderRepository
		logId                  int64
		flipKickRepo           flip_kick.FlipKickRepository
		contractAddress        = fakes.FakeAddress.Hex()
		fakeBidId              int
		blockOne, timestampOne int
	)

	BeforeEach(func() {
		fakeBidId = rand.Int()
		blockOne = rand.Int()
		timestampOne = int(rand.Int31())

		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)

		flipKickRepo = flip_kick.FlipKickRepository{}
		flipKickRepo.SetDB(db)

		headerRepository = repositories.NewHeaderRepository(db)
		headerOne = createHeader(blockOne, timestampOne, headerRepository)
		logId = test_data.CreateTestLog(headerOne.Id, db).ID

		flipStorageValues := test_helpers.GetFlipStorageValues(1, test_helpers.FakeIlk.Hex, fakeBidId)
		test_helpers.CreateFlip(db, headerOne, flipStorageValues, test_helpers.GetFlipMetadatas(strconv.Itoa(fakeBidId)), contractAddress)

		_, _, err := test_helpers.SetUpFlipBidContext(test_helpers.FlipBidContextInput{
			DealCreationInput: test_helpers.DealCreationInput{
				DB:              db,
				BidId:           fakeBidId,
				ContractAddress: contractAddress,
			},
			Dealt:            false,
			IlkHex:           test_helpers.FakeIlk.Hex,
			UrnGuy:           test_data.FlipKickModel().ForeignKeyValues[constants.UrnFK],
			FlipKickRepo:     flipKickRepo,
			FlipKickHeaderId: headerOne.Id,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		closeErr := db.Close()
		Expect(closeErr).NotTo(HaveOccurred())
	})

	Describe("flip_state_ilk", func() {
		It("returns ilk_state for a flip_state", func() {
			ilkValues := test_helpers.GetIlkValues(0)
			test_helpers.CreateIlk(db, headerOne, ilkValues, test_helpers.FakeIlkVatMetadatas,
				test_helpers.FakeIlkCatMetadatas, test_helpers.FakeIlkJugMetadatas, test_helpers.FakeIlkSpotMetadatas)

			expectedIlk := test_helpers.IlkStateFromValues(test_helpers.FakeIlk.Hex, headerOne.Timestamp, headerOne.Timestamp, ilkValues)

			var result test_helpers.IlkState
			getIlkErr := db.Get(&result, `
				SELECT ilk_identifier, rate, art, spot, line, dust, chop, lump, flip, rho, duty, pip, mat, created, updated
				FROM api.flip_state_ilk(
					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state
					 FROM api.get_flip($1, $2, $3))
			)`, fakeBidId, test_helpers.FakeIlk.Identifier, blockOne)

			Expect(getIlkErr).NotTo(HaveOccurred())
			Expect(result).To(Equal(expectedIlk))
		})
	})

	Describe("flip_state_urn", func() {
		It("returns urn_state for a flip_state", func() {
			urnSetupData := test_helpers.GetUrnSetupData()
			urnMetadata := test_helpers.GetUrnMetadata(test_helpers.FakeIlk.Hex, test_data.FlipKickModel().ForeignKeyValues[constants.UrnFK])
			vatRepository := vat.VatStorageRepository{}
			vatRepository.SetDB(db)
			test_helpers.CreateUrn(urnSetupData, headerOne.Id, urnMetadata, vatRepository)

			var actualUrn test_helpers.UrnState
			getUrnErr := db.Get(&actualUrn, `
				SELECT urn_identifier, ilk_identifier
				FROM api.flip_state_urn(
					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state
					FROM api.get_flip($1, $2, $3))
			)`, fakeBidId, test_helpers.FakeIlk.Identifier, blockOne)

			Expect(getUrnErr).NotTo(HaveOccurred())

			expectedUrn := test_helpers.UrnState{
				UrnIdentifier: test_data.FlipKickModel().ForeignKeyValues[constants.UrnFK],
				IlkIdentifier: test_helpers.FakeIlk.Identifier,
			}

			test_helpers.AssertUrn(actualUrn, expectedUrn)
		})
	})

	Describe("flip_state_bid_events", func() {
		It("returns the bid events for a flip", func() {
			// flip kick created in BeforeEach
			expectedFlipKickEvent := test_helpers.BidEvent{
				BidId:           strconv.Itoa(fakeBidId),
				Lot:             test_data.FlipKickModel().ColumnValues["lot"].(string),
				BidAmount:       test_data.FlipKickModel().ColumnValues["bid"].(string),
				Act:             "kick",
				ContractAddress: contractAddress,
			}

			tendLot := rand.Intn(100)
			tendBidAmount := rand.Intn(100)
			tendLog := test_data.CreateTestLog(headerOne.Id, db)
			flipTendErr := test_helpers.CreateTend(test_helpers.TendCreationInput{
				DB:              db,
				ContractAddress: contractAddress,
				BidId:           fakeBidId,
				Lot:             tendLot,
				BidAmount:       tendBidAmount,
				TendHeaderId:    headerOne.Id,
				TendLogId:       tendLog.ID,
			})
			Expect(flipTendErr).NotTo(HaveOccurred())

			expectedTendEvent := test_helpers.BidEvent{
				BidId:           strconv.Itoa(fakeBidId),
				Lot:             strconv.Itoa(tendLot),
				BidAmount:       strconv.Itoa(tendBidAmount),
				Act:             "tend",
				ContractAddress: contractAddress,
			}

			var actualBidEvents []test_helpers.BidEvent
			queryErr := db.Select(&actualBidEvents,
				`SELECT bid_id, bid_amount, lot, act, contract_address FROM api.flip_state_bid_events(
    					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state 
    					FROM api.get_flip($1, $2)))`, fakeBidId, test_helpers.FakeIlk.Identifier)
			Expect(queryErr).NotTo(HaveOccurred())
			Expect(actualBidEvents).To(ConsistOf(expectedFlipKickEvent, expectedTendEvent))
		})

		Describe("result pagination", func() {
			var (
				tendLot, tendBidAmount int
				flipKickEvent          shared.InsertionModel
			)

			BeforeEach(func() {
				flipKickEvent = test_data.FlipKickModel()
				flipKickEvent.ForeignKeyValues[constants.AddressFK] = contractAddress
				flipKickEvent.ColumnValues["bid_id"] = strconv.Itoa(fakeBidId)
				flipKickEvent.ColumnValues[constants.HeaderFK] = headerOne.Id
				flipKickEvent.ColumnValues[constants.LogFK] = logId
				flipKickErr := flipKickRepo.Create([]shared.InsertionModel{flipKickEvent})
				Expect(flipKickErr).NotTo(HaveOccurred())

				headerTwo := createHeader(blockOne+1, timestampOne+1, headerRepository)
				tendLogId := test_data.CreateTestLog(headerTwo.Id, db).ID

				tendLot = rand.Intn(100)
				tendBidAmount = rand.Intn(100)
				flipTendErr := test_helpers.CreateTend(test_helpers.TendCreationInput{
					DB:              db,
					ContractAddress: contractAddress,
					BidId:           fakeBidId,
					Lot:             tendLot,
					BidAmount:       tendBidAmount,
					TendHeaderId:    headerTwo.Id,
					TendLogId:       tendLogId,
				})
				Expect(flipTendErr).NotTo(HaveOccurred())
			})

			It("limits results to most recent block if max_results argument is provided", func() {
				expectedTendEvent := test_helpers.BidEvent{
					BidId:           strconv.Itoa(fakeBidId),
					Lot:             strconv.Itoa(tendLot),
					BidAmount:       strconv.Itoa(tendBidAmount),
					Act:             "tend",
					ContractAddress: contractAddress,
				}

				maxResults := 1
				var actualBidEvents []test_helpers.BidEvent
				queryErr := db.Select(&actualBidEvents,
					`SELECT bid_id, bid_amount, lot, act, contract_address FROM api.flip_state_bid_events(
    					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state 
    					FROM api.get_flip($1, $2)), $3)`, fakeBidId, test_helpers.FakeIlk.Identifier, maxResults)
				Expect(queryErr).NotTo(HaveOccurred())

				Expect(actualBidEvents).To(ConsistOf(expectedTendEvent))
			})

			It("offsets result if offset is provided", func() {
				expectedTendEvent := test_helpers.BidEvent{
					BidId:           strconv.Itoa(fakeBidId),
					Lot:             flipKickEvent.ColumnValues["lot"].(string),
					BidAmount:       flipKickEvent.ColumnValues["bid"].(string),
					Act:             "kick",
					ContractAddress: contractAddress,
				}

				maxResults := 1
				resultOffset := 1
				var actualBidEvents []test_helpers.BidEvent
				queryErr := db.Select(&actualBidEvents,
					`SELECT bid_id, bid_amount, lot, act, contract_address FROM api.flip_state_bid_events(
    					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state 
    					FROM api.get_flip($1, $2)), $3, $4)`,
					fakeBidId, test_helpers.FakeIlk.Identifier, maxResults, resultOffset)
				Expect(queryErr).NotTo(HaveOccurred())

				Expect(actualBidEvents).To(ConsistOf(expectedTendEvent))
			})
		})

		It("ignores bid events for a flip with a different ilk", func() {
			expectedBidEvent := test_helpers.BidEvent{
				BidId:           strconv.Itoa(fakeBidId),
				Lot:             test_data.FlipKickModel().ColumnValues["lot"].(string),
				BidAmount:       test_data.FlipKickModel().ColumnValues["bid"].(string),
				Act:             "kick",
				ContractAddress: contractAddress,
			}

			irrelevantContractAddress := "different flipper"
			irrelevantFlipStorageValues := test_helpers.GetFlipStorageValues(1, test_helpers.AnotherFakeIlk.Hex, fakeBidId)
			irrelevantFlipMetadatas := test_helpers.GetFlipMetadatas(strconv.Itoa(fakeBidId))
			test_helpers.CreateFlip(db, headerOne, irrelevantFlipStorageValues, irrelevantFlipMetadatas, irrelevantContractAddress)

			_, _, irrelevantFlipContextErr := test_helpers.SetUpFlipBidContext(test_helpers.FlipBidContextInput{
				DealCreationInput: test_helpers.DealCreationInput{
					DB:              db,
					BidId:           fakeBidId,
					ContractAddress: irrelevantContractAddress,
				},
				Dealt:            false,
				IlkHex:           test_helpers.AnotherFakeIlk.Hex,
				UrnGuy:           test_data.FlipKickModel().ForeignKeyValues[constants.UrnFK],
				FlipKickRepo:     flipKickRepo,
				FlipKickHeaderId: headerOne.Id,
			})
			Expect(irrelevantFlipContextErr).NotTo(HaveOccurred())

			var actualBidEvents []test_helpers.BidEvent
			queryErr := db.Select(&actualBidEvents,
				`SELECT bid_id, bid_amount, lot, act, contract_address FROM api.flip_state_bid_events(
    					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state 
    					FROM api.get_flip($1, $2)))`, fakeBidId, test_helpers.FakeIlk.Identifier)
			Expect(queryErr).NotTo(HaveOccurred())
			Expect(actualBidEvents).To(ConsistOf(expectedBidEvent))
		})

		It("returns nothing when no bid events match", func() {
			irrelevantBidId := fakeBidId + 1
			irrelevantContractAddress := "DifferentFlipper"
			irrelevantFlipStorageValues := test_helpers.GetFlipStorageValues(2, test_helpers.FakeIlk.Hex, fakeBidId)
			irrelevantFlipMetadatas := test_helpers.GetFlipMetadatas(strconv.Itoa(fakeBidId))
			test_helpers.CreateFlip(db, headerOne, irrelevantFlipStorageValues, irrelevantFlipMetadatas, irrelevantContractAddress)

			// this function creates a flip kick but we are going to use a different bid id in the select query
			// so the test should return nothing
			_, _, err := test_helpers.SetUpFlipBidContext(test_helpers.FlipBidContextInput{
				DealCreationInput: test_helpers.DealCreationInput{
					DB:              db,
					BidId:           fakeBidId,
					ContractAddress: irrelevantContractAddress,
				},
				Dealt:            false,
				IlkHex:           test_helpers.FakeIlk.Hex,
				UrnGuy:           test_data.FlipKickModel().ForeignKeyValues[constants.UrnFK],
				FlipKickRepo:     flipKickRepo,
				FlipKickHeaderId: headerOne.Id,
			})
			Expect(err).NotTo(HaveOccurred())

			var actualBidEvents []test_helpers.BidEvent
			queryErr := db.Select(&actualBidEvents,
				`SELECT bid_id, bid_amount, lot, act, contract_address FROM api.flip_state_bid_events(
    					(SELECT (block_height, bid_id, ilk_id, urn_id, guy, tic, "end", lot, bid, gal, dealt, tab, created, updated)::api.flip_state 
    					FROM api.get_flip($1, $2)))`, irrelevantBidId, test_helpers.FakeIlk.Identifier)
			Expect(queryErr).NotTo(HaveOccurred())
			Expect(actualBidEvents).To(BeZero())
		})
	})
})
