package test_data

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared/constants"
)

// This file contains "shortcuts" to some configuration values useful for testing

func Cat100Address() string   { return checksum(constants.GetContractAddress("MCD_CAT_1_0_0")) }
func Cat110Address() string   { return checksum(constants.GetContractAddress("MCD_CAT_1_1_0")) }
func FlapV100Address() string { return checksum(constants.GetContractAddress("MCD_FLAP_1_0_0")) }
func FlapV109Address() string { return checksum(constants.GetContractAddress("MCD_FLAP_1_0_9")) }
func FlipV100Addresses() []string {
	var addressesResult []string
	flipAddresses := constants.GetContractAddresses([]string{
		"MCD_FLIP_BAT_A_1_0_0",
		"MCD_FLIP_BAT_A_1_0_9",
		"MCD_FLIP_ETH_A_1_0_0",
		"MCD_FLIP_ETH_A_1_0_9",
		"MCD_FLIP_KNC_A_1_0_8",
		"MCD_FLIP_KNC_A_1_0_9",
		"MCD_FLIP_MANA_A_1_0_9",
		"MCD_FLIP_SAI_1_0_0",
		"MCD_FLIP_TUSD_A_1_0_7",
		"MCD_FLIP_TUSD_A_1_0_9",
		"MCD_FLIP_USDC_A_1_0_4",
		"MCD_FLIP_USDC_A_1_0_9",
		"MCD_FLIP_USDC_B_1_0_7",
		"MCD_FLIP_USDC_B_1_0_9",
		"MCD_FLIP_WBTC_A_1_0_6",
		"MCD_FLIP_WBTC_A_1_0_9",
		"MCD_FLIP_ZRX_A_1_0_8",
		"MCD_FLIP_ZRX_A_1_0_9",
	})

	for _, address := range flipAddresses {
		addressesResult = append(addressesResult, checksum(address))
	}
	return addressesResult
}

func FlipV110Addresses() []string {
	var addressesResult []string
	flipAddresses := constants.GetContractAddresses([]string{
		"MCD_FLIP_AAVE_A_1_2_2",
		"MCD_FLIP_BAL_A_1_1_14",
		"MCD_FLIP_BAT_A_1_1_0",
		"MCD_FLIP_COMP_A_1_1_2",
		"MCD_FLIP_ETH_A_1_1_0",
		"MCD_FLIP_ETH_B_1_1_3",
		"MCD_FLIP_GUSD_A_1_1_5",
		"MCD_FLIP_KNC_A_1_1_0",
		"MCD_FLIP_LINK_A_1_1_2",
		"MCD_FLIP_LRC_A_1_1_2",
		"MCD_FLIP_MANA_A_1_1_0",
		"MCD_FLIP_PAXUSD_A_1_1_1",
		"MCD_FLIP_RENBTC_A_1_2_1",
		"MCD_FLIP_TUSD_A_1_1_0",
		"MCD_FLIP_UNI_A_1_2_1",
		"MCD_FLIP_UNIV2DAIETH_A_1_2_2",
		"MCD_FLIP_USDC_A_1_1_0",
		"MCD_FLIP_USDC_B_1_1_0",
		"MCD_FLIP_USDT_A_1_1_1",
		"MCD_FLIP_WBTC_A_1_1_0",
		"MCD_FLIP_YFI_A_1_1_14",
		"MCD_FLIP_ZRX_A_1_1_0",
	})

	for _, address := range flipAddresses {
		addressesResult = append(addressesResult, checksum(address))
	}
	return addressesResult
}
func FlipAaveAV122Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_AAVE_A_1_2_2"))
}
func FlipBalAV1114Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_BAL_A_1_1_14"))
}
func FlipBatV100Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_BAT_A_1_0_0"))
}
func FlipBatV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_BAT_A_1_0_9"))
}
func FlipBatV110Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_BAT_A_1_1_0"))
}
func FlipCompV112Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_COMP_A_1_1_2"))
}
func FlipEthAV100Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ETH_A_1_0_0"))
}
func FlipEthAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ETH_A_1_0_9"))
}
func FlipEthV110Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ETH_A_1_1_0"))
}
func FlipEthBV113Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ETH_B_1_1_3"))
}
func FlipGusdAV115Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_GUSD_A_1_1_5"))
}
func FlipKncAV108Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_KNC_A_1_0_8"))
}
func FlipKncAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_KNC_A_1_0_9"))
}
func FlipLinkV112Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_LINK_A_1_1_2"))
}
func FlipLrcV112Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_LRC_A_1_1_2"))
}
func FlipManaAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_MANA_A_1_0_9"))
}
func FlipPaxusdAV111Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_PAXUSD_A_1_1_1"))
}
func FlipRenbtcA121Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_RENBTC_A_1_2_1"))
}
func FlipTusdAV107Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_TUSD_A_1_0_7"))
}
func FlipTusdAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_TUSD_A_1_0_9"))
}
func FlipUniAV121Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_UNI_A_1_2_1"))
}
func FlipUniV2DaiEthAddress() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_UNIV2DAIETH_A_1_2_2"))
}
func FlipUsdcAV104Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_USDC_A_1_0_4"))
}
func FlipUsdcAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_USDC_A_1_0_9"))
}
func FlipUsdcBV107Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_USDC_B_1_0_7"))
}
func FlipUsdcBV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_USDC_B_1_0_9"))
}
func FlipUsdtAV111Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_USDT_A_1_1_1"))
}
func FlipWbtcAV106Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_WBTC_A_1_0_6"))
}
func FlipWbtcAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_WBTC_A_1_0_9"))
}
func FlipYfiAV1114Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_YFI_A_1_1_14"))
}
func FlipZrxAV108Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ZRX_A_1_0_8"))
}
func FlipZrxAV109Address() string {
	return checksum(constants.GetContractAddress("MCD_FLIP_ZRX_A_1_0_9"))
}
func FlopV101Address() string { return checksum(constants.GetContractAddress("MCD_FLOP_1_0_1")) }
func FlopV109Address() string { return checksum(constants.GetContractAddress("MCD_FLOP_1_0_9")) }
func JugAddress() string      { return checksum(constants.GetContractAddress("MCD_JUG")) }
func MedianAddresses() []string {
	var addressesResult []string
	medianAddresses := constants.GetContractAddresses([]string{
		"MEDIAN_AAVE_1_2_2",
		"MEDIAN_BAL_1_1_14",
		"MEDIAN_BAT_1_0_0",
		"MEDIAN_COMP_1_1_2",
		"MEDIAN_ETH_1_0_0",
		"MEDIAN_KNC_1_0_8",
		"MEDIAN_LINK_1_1_2",
		"MEDIAN_LRC_1_1_2",
		"MEDIAN_MANA_1_0_9",
		"MEDIAN_UNI_1_2_1",
		"MEDIAN_USDT_1_0_4",
		"MEDIAN_WBTC_1_0_6",
		"MEDIAN_YFI_1_1_14",
		"MEDIAN_ZRX_1_0_8",
	})

	for _, address := range medianAddresses {
		addressesResult = append(addressesResult, checksum(address))
	}
	return addressesResult
}
func MedianAaveAddress() string { return checksum(constants.GetContractAddress("MEDIAN_AAVE_1_2_2")) }
func MedianBalAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_BAL_1_1_14")) }
func MedianBatAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_BAT_1_0_0")) }
func MedianCompAddress() string {
	return checksum(constants.GetContractAddress("MEDIAN_COMP_1_1_2"))
}
func MedianEthAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_ETH_1_0_0")) }
func MedianKncAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_KNC_1_0_8")) }
func MedianLinkAddress() string { return checksum(constants.GetContractAddress("MEDIAN_LINK_1_1_2")) }
func MedianLrcAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_LRC_1_1_2")) }
func MedianManaAddress() string { return checksum(constants.GetContractAddress("MEDIAN_MANA_1_0_9")) }
func MedianUniAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_UNI_1_2_1")) }
func MedianUsdtAddress() string { return checksum(constants.GetContractAddress("MEDIAN_USDT_1_0_4")) }
func MedianWbtcAddress() string { return checksum(constants.GetContractAddress("MEDIAN_WBTC_1_0_6")) }
func MedianYfiAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_YFI_1_1_14")) }
func MedianZrxAddress() string  { return checksum(constants.GetContractAddress("MEDIAN_ZRX_1_0_8")) }
func OsmAddresses() []string {
	var addressesResult []string
	// Does not include OSM_USDC since that's actually just a DSValue contract right now, not an OSM
	osmAddresses := constants.GetContractAddresses([]string{
		"OSM_AAVE",
		"OSM_BAL",
		"OSM_BAT",
		"OSM_COMP",
		"OSM_ETH",
		"OSM_KNC",
		"OSM_LINK",
		"OSM_LRC",
		"OSM_MANA",
		"OSM_UNI",
		"OSM_USDT",
		"OSM_WBTC",
		"OSM_YFI",
		"OSM_ZRX",
	})

	for _, address := range osmAddresses {
		addressesResult = append(addressesResult, checksum(address))
	}
	return addressesResult
}
func OsmAaveAddress() string    { return checksum(constants.GetContractAddress("OSM_AAVE")) }
func OsmBalAddress() string     { return checksum(constants.GetContractAddress("OSM_BAL")) }
func OsmBatAddress() string     { return checksum(constants.GetContractAddress("OSM_BAT")) }
func OsmCompAddress() string    { return checksum(constants.GetContractAddress("OSM_COMP")) }
func OsmEthAddress() string     { return checksum(constants.GetContractAddress("OSM_ETH")) }
func OsmKncAddress() string     { return checksum(constants.GetContractAddress("OSM_KNC")) }
func OsmLinkAddress() string    { return checksum(constants.GetContractAddress("OSM_LINK")) }
func OsmLrcAddress() string     { return checksum(constants.GetContractAddress("OSM_LRC")) }
func OsmManaAddress() string    { return checksum(constants.GetContractAddress("OSM_MANA")) }
func OsmUniAddress() string     { return checksum(constants.GetContractAddress("OSM_UNI")) }
func OsmUsdtAddress() string    { return checksum(constants.GetContractAddress("OSM_USDT")) }
func OsmWbtcAddress() string    { return checksum(constants.GetContractAddress("OSM_WBTC")) }
func OsmYfiAddress() string     { return checksum(constants.GetContractAddress("OSM_YFI")) }
func OsmZrxAddress() string     { return checksum(constants.GetContractAddress("OSM_ZRX")) }
func PotAddress() string        { return checksum(constants.GetContractAddress("MCD_POT")) }
func SpotAddress() string       { return checksum(constants.GetContractAddress("MCD_SPOT")) }
func VatAddress() string        { return checksum(constants.GetContractAddress("MCD_VAT")) }
func VowAddress() string        { return checksum(constants.GetContractAddress("MCD_VOW")) }
func CdpManagerAddress() string { return checksum(constants.GetContractAddress("CDP_MANAGER")) }

func checksum(addressString string) string {
	return common.HexToAddress(addressString).Hex()
}
