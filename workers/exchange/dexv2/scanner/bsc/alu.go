package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetALUPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x3C7e7122f4DDcfC5b5c31b5c735Ce2Ae3b015856")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xca7172132824954667AAD12600450097D04D7b29")}) //KEWL

	return pairs
}

func GetALU() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "ALU",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetALUPairs(),
	}
}
