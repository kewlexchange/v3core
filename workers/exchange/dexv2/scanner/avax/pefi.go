package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetPEFIPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb78c8238bD907c42BE45AeBdB4A8C8a5D7B49755")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x494Dd9f783dAF777D3fb4303da4de795953592d0")}) //png

	return pairs
}

func GetPEFI() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "GetPEFI",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetPEFIPairs(),
	}
}
