package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetArenaPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc010A798F9eED73395A9C72520298ba33a48a2B7")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2552077d373A1117DeAe531f38b791CFBeFC379f")}) //png

	return pairs
}

func GetARENA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "ARENA",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetArenaPairs(),
	}
}
