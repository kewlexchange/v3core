package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBABYDOGEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc736cA3d9b1E90Af4230BD8F9626528B3D4e0Ee0")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x719B97aA936223C31d4263eEB2b33DCBdCB0F666")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xbed24d194AfE26169a1F4a419E381dd4e99f9C1f")}) //KEWL

	return pairs
}

func GetBABYDOGE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BABYDOGE",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBABYDOGEPairs(),
	}
}
