package queries

import (
	"math/rand"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/mcd_transformers/test_config"
	"github.com/vulcanize/mcd_transformers/transformers/component_tests/queries/test_helpers"
	"github.com/vulcanize/mcd_transformers/transformers/events/flop_kick"
	"github.com/vulcanize/mcd_transformers/transformers/test_data"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
)

var _ = Describe("Flop computed columns", func() {
	var (
		db              *postgres.DB
		flopKickRepo    flop_kick.FlopKickRepository
		headerRepo      repositories.HeaderRepository
		contractAddress = "0x763ztv6x68exwqrgtl325e7hrcvavid4e3fcb4g"

		fakeBidId      = rand.Int()
		blockOne       = rand.Int()
		timestampOne   = int(rand.Int31())
		blockOneHeader = fakes.GetFakeHeaderWithTimestamp(int64(timestampOne), int64(blockOne))
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		headerRepo = repositories.NewHeaderRepository(db)
		flopKickRepo = flop_kick.FlopKickRepository{}
		flopKickRepo.SetDB(db)
	})

	AfterEach(func() {
		closeErr := db.Close()
		Expect(closeErr).NotTo(HaveOccurred())
	})

	Describe("flop_bid_events", func() {
		It("returns the bid events for flop", func() {
			headerId, headerErr := headerRepo.CreateOrUpdateHeader(blockOneHeader)
			Expect(headerErr).NotTo(HaveOccurred())

			flopStorageValues := test_helpers.GetFlopStorageValues(1, fakeBidId)
			test_helpers.CreateFlop(db, blockOneHeader, flopStorageValues, test_helpers.GetFlopMetadatas(strconv.Itoa(fakeBidId)), contractAddress)

			flopKickEvent := test_data.FlopKickModel
			flopKickEvent.ContractAddress = contractAddress
			flopKickEvent.BidId = strconv.Itoa(fakeBidId)
			flopKickErr := flopKickRepo.Create(headerId, []interface{}{flopKickEvent})
			Expect(flopKickErr).NotTo(HaveOccurred())

			expectedBidEvents := test_helpers.BidEvent{
				BidId:     strconv.Itoa(fakeBidId),
				Lot:       flopKickEvent.Lot,
				BidAmount: flopKickEvent.Bid,
				Act:       "kick",
			}
			var actualBidEvents test_helpers.BidEvent
			queryErr := db.Get(&actualBidEvents,
				`SELECT bid_id, bid_amount, lot, act FROM api.flop_bid_events(
    					(SELECT (bid_id, guy, tic, "end", lot, bid, dealt, created, updated)::api.flop 
    					FROM api.all_flops()))`)
			Expect(queryErr).NotTo(HaveOccurred())
			Expect(actualBidEvents).To(Equal(expectedBidEvents))
		})

		It("does not include bid events for a different flop", func() {
			headerId, headerErr := headerRepo.CreateOrUpdateHeader(blockOneHeader)
			Expect(headerErr).NotTo(HaveOccurred())

			flopStorageValues := test_helpers.GetFlopStorageValues(1, fakeBidId)
			test_helpers.CreateFlop(db, blockOneHeader, flopStorageValues, test_helpers.GetFlopMetadatas(strconv.Itoa(fakeBidId)), contractAddress)

			flopKickEvent := test_data.FlopKickModel
			flopKickEvent.ContractAddress = contractAddress
			flopKickEvent.BidId = strconv.Itoa(fakeBidId)
			flopKickErr := flopKickRepo.Create(headerId, []interface{}{flopKickEvent})
			Expect(flopKickErr).NotTo(HaveOccurred())

			blockTwo := blockOne + 1
			timestampTwo := timestampOne + 111111
			blockTwoHeader := fakes.GetFakeHeaderWithTimestamp(int64(timestampTwo), int64(blockTwo))
			headerTwoId, headerTwoErr := headerRepo.CreateOrUpdateHeader(blockTwoHeader)
			Expect(headerTwoErr).NotTo(HaveOccurred())

			irrelevantBidId := fakeBidId + 9999999999999
			irrelevantFlopStorageValues := test_helpers.GetFlopStorageValues(2, irrelevantBidId)
			test_helpers.CreateFlop(db, blockTwoHeader, irrelevantFlopStorageValues, test_helpers.GetFlopMetadatas(strconv.Itoa(irrelevantBidId)), contractAddress)

			irrelevantFlopKickEvent := test_data.FlopKickModel
			irrelevantFlopKickEvent.ContractAddress = contractAddress
			irrelevantFlopKickEvent.BidId = strconv.Itoa(irrelevantBidId)

			flopKickErr = flopKickRepo.Create(headerTwoId, []interface{}{irrelevantFlopKickEvent})
			Expect(flopKickErr).NotTo(HaveOccurred())

			expectedBidEvents := test_helpers.BidEvent{
				BidId:     strconv.Itoa(fakeBidId),
				Lot:       flopKickEvent.Lot,
				BidAmount: flopKickEvent.Bid,
				Act:       "kick",
			}

			var actualBidEvents []test_helpers.BidEvent
			queryErr := db.Select(&actualBidEvents,
				`SELECT bid_id, bid_amount, lot, act FROM api.flop_bid_events(
    					(SELECT (bid_id, guy, tic, "end", lot, bid, dealt, created, updated)::api.flop
    					FROM api.all_flops() WHERE bid_id = $1))`, fakeBidId)

			Expect(queryErr).NotTo(HaveOccurred())
			Expect(actualBidEvents).To(ConsistOf(expectedBidEvents))
		})
	})
})