package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBATPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xD72AE03Be5ce45bB6FdFE0AA0D81d7eef709dCC3")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xa099F0D90Eb17283Ae7d302d0a6ebb0CaCab39DF")}) //KEWL

	return pairs
}

func GetBAT() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BAT",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBATPairs(),
	}
}
