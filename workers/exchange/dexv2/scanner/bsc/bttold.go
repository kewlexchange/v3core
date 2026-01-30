package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBTTOLDPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x946696344e7d4346b223e1Cf77035a76690d6A73")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x6E28eb8244ACf16c2907d0d0AD7C52436B75099c")}) //KEWL

	return pairs
}

func GetBTTOLD() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BTTOLD",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBTTOLDPairs(),
	}
}
