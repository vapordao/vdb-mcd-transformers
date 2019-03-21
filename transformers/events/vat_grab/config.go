package vat_grab

import (
	shared_t "github.com/vulcanize/vulcanizedb/libraries/shared/transformer"

	"github.com/vulcanize/mcd_transformers/transformers/shared/constants"
)

func GetVatGrabConfig() shared_t.EventTransformerConfig {
	return shared_t.EventTransformerConfig{
		TransformerName:     constants.VatGrabLabel,
		ContractAddresses:   []string{constants.OldVatContractAddress()},
		ContractAbi:         constants.OldVatABI(),
		Topic:               constants.GetVatGrabSignature(),
		StartingBlockNumber: constants.OldVatDeploymentBlock(),
		EndingBlockNumber:   -1,
	}
}
