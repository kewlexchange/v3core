package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBMONPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x3C2b7B578Dd2175A1c3524Aa0D515106282Bf108")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xcc1fCabD1546bC6d9b543DC47E19E299CC5bcd4E")}) //KEWL

	return pairs
}

func GetBMON() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BMON",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBMONPairs(),
	}
}
