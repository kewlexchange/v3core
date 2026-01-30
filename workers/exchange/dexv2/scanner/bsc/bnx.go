package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBNXPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5f37d6817a4D6DF8E2A9B4115f58c93805d9da63")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5cb2EA1302C65e1B46D10AAC0f1Aba031Fc3B435")}) //KEWL

	return pairs
}

func GetBNX() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BNX",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBNXPairs(),
	}
}
