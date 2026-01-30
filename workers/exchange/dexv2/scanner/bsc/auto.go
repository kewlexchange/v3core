package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetAUTOPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5657cada6B6bbEF1DBd77aEDB1A6105cFAb4b836")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xa34E97d80b76315665687E36c0E8b7f6a611685F")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xeA3D46e1360dA830c72A1429F0bB2c90beAaE2aA")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xeA3D46e1360dA830c72A1429F0bB2c90beAaE2aA")}) //KEWL

	return pairs
}

func GetAUTO() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "AUTO",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetAUTOPairs(),
	}
}
