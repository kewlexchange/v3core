package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetREEFPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xfbbd096A99e95D6808b918A5C0863ed9989EBd41")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd63b5CecB1f40d626307B92706Df357709D05827")}) //KEWL

	return pairs
}

func GetREEF() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "REEF",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetREEFPairs(),
	}
}
