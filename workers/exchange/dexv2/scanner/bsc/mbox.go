package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetMBOXPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8FA59693458289914dB0097F5F366d771B7a7C3F")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x76e1B3B2B15A4Ff61aB3E245d6b98ae808DEe6e1")}) //KEWL

	return pairs
}

func GetMBOX() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "MBOX",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetMBOXPairs(),
	}
}
