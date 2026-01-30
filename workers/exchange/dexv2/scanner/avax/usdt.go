package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDTPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xbb4646a764358ee93c2a9c4a147d5aDEd527ab73")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xe3bA3d5e3F98eefF5e9EDdD5Bd20E476202770da")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x746C8cb107508033dFb7F6B0E94ec7600Efc99a7")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x7e5E4b677c2a682B6d2e95Ae3ec07ae1Ea7D3aB5")}) //png

	return pairs
}

func GetUSDT() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "USDT",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetUSDTPairs(),
	}
}
