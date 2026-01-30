package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetRACAPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xC016F93D1b11878804c345e93C8588794480CD83")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x9C3d4Fb14D3A021aee4Fd763506B1F71d509Dc90")}) //KEWL

	return pairs
}

func GetRACA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "RACA",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetRACAPairs(),
	}
}
