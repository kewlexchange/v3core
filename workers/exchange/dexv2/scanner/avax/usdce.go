package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDCEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xA389f9430876455C36478DeEa9769B7Ca4E3DDB1")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xbd918Ed441767fe7924e99F6a0E0B568ac1970D9")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x6c508C68398793247d9E03B506bbC6C7535b5a90")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4ed65dAB34d5FD4b1eb384432027CE47E90E1185")}) //png
	return pairs
}

func GetUSDCE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "USDCE",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetUSDCEPairs(),
	}
}
