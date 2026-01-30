package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBSWPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8CA3fF14A52b080C54A6d1a405eecA02959d39fE")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x46492B26639Df0cda9b2769429845cb991591E0A")}) //KEWL

	return pairs
}

func GetBSW() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BSW",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBSWPairs(),
	}
}
