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

package integration_tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/makerdao/vdb-mcd-transformers/test_config"
	"github.com/makerdao/vdb-mcd-transformers/transformers/events/tend"
	"github.com/makerdao/vdb-mcd-transformers/transformers/shared/constants"
	"github.com/makerdao/vdb-mcd-transformers/transformers/test_data"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/fetcher"
	"github.com/makerdao/vulcanizedb/libraries/shared/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tend EventTransformer", func() {
	var (
		tendConfig  event.TransformerConfig
		initializer event.ConfiguredTransformer
		logFetcher  fetcher.ILogFetcher
		addresses   []common.Address
		topics      []common.Hash
	)

	BeforeEach(func() {
		test_config.CleanTestDB(db)

		tendConfig = event.TransformerConfig{
			TransformerName:   constants.TendTable,
			ContractAddresses: []string{test_data.FlipEthAV100Address(), test_data.FlapV100Address()},
			ContractAbi:       constants.FlipV100ABI(),
			Topic:             constants.TendSignature(),
		}

		logFetcher = fetcher.NewLogFetcher(blockChain)
		addresses = event.HexStringsToAddresses(tendConfig.ContractAddresses)
		topics = []common.Hash{common.HexToHash(tendConfig.Topic)}

		initializer = event.ConfiguredTransformer{
			Config:      tendConfig,
			Transformer: tend.Transformer{},
		}
	})

	It("fetches and transforms a Flip Tend event", func() {
		blockNumber := int64(9004844)
		initializer.Config.StartingBlockNumber = blockNumber
		initializer.Config.EndingBlockNumber = blockNumber

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		logs, err := logFetcher.FetchLogs(addresses, topics, header)
		Expect(err).NotTo(HaveOccurred())

		eventLogs := test_data.CreateLogs(header.Id, logs, db)

		transformer := initializer.NewTransformer(db)
		err = transformer.Execute(eventLogs)
		Expect(err).NotTo(HaveOccurred())

		var flipTend tendModel
		err = db.Get(&flipTend, `SELECT address_id, msg_sender, bid, bid_id, lot FROM maker.tend`)
		Expect(err).NotTo(HaveOccurred())

		flipAddressID, addrErr := repository.GetOrCreateAddress(db, test_data.FlipEthAV100Address())
		Expect(addrErr).NotTo(HaveOccurred())

		msgSender := shared.GetChecksumAddressString("0x00000000000000000000000000abe7471ec9b6953a3bd0ed3c06c46f29aa4280")
		msgSenderID, msgSenderErr := repository.GetOrCreateAddress(db, msgSender)
		Expect(msgSenderErr).NotTo(HaveOccurred())

		expectedFlipTend := tendModel{
			AddressID: flipAddressID,
			MsgSender: msgSenderID,
			Bid:       "76840636079422693500873675445736719538580144543",
			BidId:     "121",
			Lot:       "700000000000000000",
		}
		Expect(flipTend).To(Equal(expectedFlipTend))
	})

	It("fetches and transforms a Flap Tend event", func() {
		blockNumber := int64(9656046)
		initializer.Config.StartingBlockNumber = blockNumber
		initializer.Config.EndingBlockNumber = blockNumber

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		logs, err := logFetcher.FetchLogs(addresses, topics, header)
		Expect(err).NotTo(HaveOccurred())

		eventLogs := test_data.CreateLogs(header.Id, logs, db)

		transformer := initializer.NewTransformer(db)
		err = transformer.Execute(eventLogs)
		Expect(err).NotTo(HaveOccurred())

		var flapTend tendModel
		err = db.Get(&flapTend, `SELECT address_id, msg_sender, bid, bid_id, lot FROM maker.tend`)
		Expect(err).NotTo(HaveOccurred())

		flapAddressID, addrErr := repository.GetOrCreateAddress(db, test_data.FlapV100Address())
		Expect(addrErr).NotTo(HaveOccurred())

		msgSender := shared.GetChecksumAddressString("0x000000000000000000000000d9d1e81bb35db066986fa441113a27708663d70b")
		msgSenderID, msgSenderErr := repository.GetOrCreateAddress(db, msgSender)
		Expect(msgSenderErr).NotTo(HaveOccurred())

		expectedFlapTend := tendModel{
			AddressID: flapAddressID,
			MsgSender: msgSenderID,
			Bid:       "22836140232828485845",
			BidId:     "55",
			Lot:       "10000000000000000000000000000000000000000000000000",
		}
		Expect(flapTend).To(Equal(expectedFlapTend))
	})
})

type tendModel struct {
	AddressID int64 `db:"address_id"`
	MsgSender int64 `db:"msg_sender"`
	Bid       string
	BidId     string `db:"bid_id"`
	Lot       string
}
