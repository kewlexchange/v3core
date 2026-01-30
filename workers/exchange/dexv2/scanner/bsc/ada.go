package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetADAPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x28415ff2C35b65B9E5c7de82126b4015ab9d031F")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8860922Eb2795aB0D57363653Dd7EBf18D7c0A42")}) //KEWL

	return pairs
}

func GetADA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "ADA",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetADAPairs(),
	}
}
