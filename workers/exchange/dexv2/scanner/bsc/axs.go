package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetAXSPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x19058558Bbc66C2Dd97c5cA8a189d350A34e4423")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xC2d00De94795e60FB76Bc37d899170996cBdA436")}) //KEWL

	return pairs
}

func GetAXS() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "AXS",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetAXSPairs(),
	}
}
