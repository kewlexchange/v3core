package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetALPHAPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xACF47CBEaab5c8A6Ee99263cfE43995f89fB3206")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5ea2efFB38a84Fb95e7B7713f0c10BDA99b25A9F")}) //KEWL

	return pairs
}

func GetALPHA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "ALPHA",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetALPHAPairs(),
	}
}
