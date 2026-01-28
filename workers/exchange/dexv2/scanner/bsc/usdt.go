package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDTPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x9Ec271C041a18aA7beF070A1F196eea1D06Ab7cb")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x38153dAE67B364DC2639717b5458461598762E0a")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8840C6252e2e86e545deFb6da98B2a0E26d8C1BA")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")}) //KEWL

	return pairs
}

func GetUSDT() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "USDT",
		NativeToken: *utils.AddressFromHex("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		Pairs:       GetUSDCPairs(),
	}
}
