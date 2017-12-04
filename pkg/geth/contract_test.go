package geth_test

import (
	"math/big"

	cfg "github.com/8thlight/vulcanizedb/pkg/config"
	"github.com/8thlight/vulcanizedb/pkg/geth"
	"github.com/8thlight/vulcanizedb/pkg/geth/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading contracts", func() {

	Describe("Reading the list of attributes", func() {
		It("returns a string attribute for a real contract", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, err := blockchain.GetContract(contractHash)
			//contractAttributes, _ := blockchain.GetContractAttributes(contractHash)

			Expect(err).To(BeNil())
			Expect(len(contract.Attributes)).NotTo(Equal(0))
			symbolAttribute := *testing.FindAttribute(contract.Attributes, "symbol")
			Expect(symbolAttribute.Name).To(Equal("symbol"))
			Expect(symbolAttribute.Type).To(Equal("string"))
		})

		It("does not return an attribute that takes an input", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, err := blockchain.GetContract(contractHash)

			Expect(err).To(BeNil())
			attribute := testing.FindAttribute(contract.Attributes, "balanceOf")
			Expect(attribute).To(BeNil())
		})

		It("does not return an attribute that is not constant", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, err := blockchain.GetContract(contractHash)

			Expect(err).To(BeNil())
			attribute := testing.FindAttribute(contract.Attributes, "unpause")
			Expect(attribute).To(BeNil())
		})
	})

	Describe("Getting a contract attribute", func() {
		It("returns the correct attribute for a real contract", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, _ := blockchain.GetContract(contractHash)
			name, err := blockchain.GetAttribute(contract, "name", nil)

			Expect(err).To(BeNil())
			Expect(name).To(Equal("OMGToken"))
		})

		It("returns the correct attribute for a real contract", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, _ := blockchain.GetContract(contractHash)
			name, err := blockchain.GetAttribute(contract, "name", nil)

			Expect(err).To(BeNil())
			Expect(name).To(Equal("OMGToken"))
		})

		It("returns the correct attribute for a real contract at a specific block height", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, _ := blockchain.GetContract(contractHash)
			name, err := blockchain.GetAttribute(contract, "name", big.NewInt(4652791))

			Expect(name).To(Equal("OMGToken"))
			Expect(err).To(BeNil())
		})

		It("returns an error when there is no ABI for the given contract", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "MISSINGHASH"

			contract, _ := blockchain.GetContract(contractHash)
			name, err := blockchain.GetAttribute(contract, "name", nil)

			Expect(err).To(Equal(geth.ErrMissingAbiFile))
			Expect(name).To(BeNil())
		})

		It("returns an error when asking for an attribute that does not exist", func() {
			config, _ := cfg.NewConfig("public")
			blockchain := geth.NewGethBlockchain(config.Client.IPCPath)
			contractHash := "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"

			contract, _ := blockchain.GetContract(contractHash)
			name, err := blockchain.GetAttribute(contract, "missing_attribute", nil)

			Expect(err).To(Equal(geth.ErrInvalidStateAttribute))
			Expect(name).To(BeNil())
		})
	})

})
