package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetINJPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x1BdCebcA3b93af70b58C41272AEa2231754B23ca")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x078BE27872f928e177a9e3425A40F5f229785f09")}) //KEWL

	return pairs
}

func GetINJ() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "INJ",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetINJPairs(),
	}
}
