package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetDOGEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xac109C8025F272414fd9e2faA805a583708A017f")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x27F190037Ab0Dcf89a5f79ecA5179c3Dbe474302")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x1eF315fa08e0E1B116D97E3dFE0aF292Ed8b7f02")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x302DbdACD7BCDCf288bd6342F9ffa5F31E8E6744")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb3432500334e8b08f12A66916912456Aad1c78C9")}) //KEWL

	return pairs
}

func GetDOGE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "DOGE",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetDOGEPairs(),
	}
}
