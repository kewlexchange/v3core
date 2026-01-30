package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetC98Pairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x38fd42c46Cb8Db714034dF920f6663b31Bb63DDe")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x92247860A03F48d5c6425c7CA35CDcFCB1013AA1")}) //KEWL

	return pairs
}

func GetC98() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "C98",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetC98Pairs(),
	}
}
