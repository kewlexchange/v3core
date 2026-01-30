package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetETHPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x74E4716E431f45807DCF19f284c7aA99F18a4fbc")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x21fbF92Bbe1903d0af6e8b21c9Ee5B43565Bd3b9")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5bf6941f029424674bb93A43b79fc46bF4A67c21")}) //KEWL

	return pairs
}

func GetETH() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "ETH",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetETHPairs(),
	}
}
