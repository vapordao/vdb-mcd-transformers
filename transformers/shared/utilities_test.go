package shared_test

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/mcd_transformers/transformers/shared"
	"math/big"
)

var _ = Describe("Shared utilities", func() {
	Describe("getting data at index", func() {
		It("gets bytes for the last index in log data", func() {
			logData := hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c45dd6471a66616b6520696c6b0000000000000000000000000000000000000000000000007d7bee5fcfd8028cf7b00876c5b1421c800561a6000000000000000000000000a3e37186e017747dba34042e83e3f76ad3cce9b00000000000000000000000000f243e26db94b5426032e6dfa6007802dea2a61400000000000000000000000000000000000000000000000000000000000000000000000000000000075bcd15000000000000000000000000000000000000000000000000000000003ade68b1")
			bigIntBytes := big.NewInt(987654321).Bytes()
			// big.Int.Bytes() does not include zero padding, but bytes in data index are of fixed length and include zero padding
			expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			expected = append(expected, bigIntBytes...)

			actual := shared.GetDataBytesAtIndex(-1, logData)

			Expect(expected[:]).To(Equal(actual))
		})

		It("gets bytes for the second-to-last index in log data", func() {
			logData := hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c45dd6471a66616b6520696c6b0000000000000000000000000000000000000000000000007d7bee5fcfd8028cf7b00876c5b1421c800561a6000000000000000000000000a3e37186e017747dba34042e83e3f76ad3cce9b00000000000000000000000000f243e26db94b5426032e6dfa6007802dea2a61400000000000000000000000000000000000000000000000000000000000000000000000000000000075bcd15000000000000000000000000000000000000000000000000000000003ade68b1")
			bigIntBytes := big.NewInt(123456789).Bytes()
			expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			expected = append(expected, bigIntBytes...)

			actual := shared.GetDataBytesAtIndex(-2, logData)

			Expect(expected[:]).To(Equal(actual))
		})

		It("gets bytes for the third-to-last index in log data", func() {
			logData := hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c45dd6471a66616b6520696c6b0000000000000000000000000000000000000000000000007d7bee5fcfd8028cf7b00876c5b1421c800561a6000000000000000000000000a3e37186e017747dba34042e83e3f76ad3cce9b00000000000000000000000000f243e26db94b5426032e6dfa6007802dea2a61400000000000000000000000000000000000000000000000000000000000000000000000000000000075bcd15000000000000000000000000000000000000000000000000000000003ade68b1")
			addressBytes := common.HexToAddress("0x0F243E26db94B5426032E6DFA6007802Dea2a614").Bytes()
			// common.address.Bytes() returns [20]byte{}, need [32]byte{}
			expected := append(addressBytes, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}...)

			actual := shared.GetDataBytesAtIndex(-3, logData)

			Expect(expected[:]).To(Equal(actual))
		})

		It("converts values to rays", func() {
			rayOne := shared.ConvertToRay("123456789012345678901234567890")
			Expect(rayOne).To(Equal("123.456789012345680589533003513"))

			rayTwo := shared.ConvertToRay("1234567890123456790123567890")
			Expect(rayTwo).To(Equal("1.234567890123456912476740399"))
		})

		It("converts values to wads", func() {
			wadOne := shared.ConvertToWad("12345678901234567890123")
			Expect(wadOne).To(Equal("12345.678901234567092615"))

			wadTwo := shared.ConvertToWad("1234567890123456789")
			Expect(wadTwo).To(Equal("1.234567890123456690"))
		})
	})

	Describe("getting hex without prefix", func() {
		It("returns bytes as hex without 0x prefix", func() {
			raw := common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000").Bytes()
			result := shared.GetHexWithoutPrefix(raw)
			Expect(result).To(Equal("4554480000000000000000000000000000000000000000000000000000000000"))
		})
	})
})