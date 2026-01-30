package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetONEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x9D2296e2Fe3CdBf2EB3e3e2CA8811BaFA42eeDFF")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xe73fe11863e4C3714EAFDee832a0987b33651f27")}) //KEWL

	return pairs
}

func GetONE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "ONE",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetONEPairs(),
	}
}
