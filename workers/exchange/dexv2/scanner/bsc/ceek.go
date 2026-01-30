package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetCEEKPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x046A9B3A9b743340eE2Bc4C6dDD35543E237C6c2")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xC115389F942483DA3805F74EF5e6462fF4b139D9")}) //KEWL

	return pairs
}

func GetCEEK() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "CEEK",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetCEEKPairs(),
	}
}
