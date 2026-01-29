package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetJOEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x454E67025631C065d3cFAD6d71E6892f74487a15")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x134Ad631337E8Bf7E01bA641fB650070a2e0efa8")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb73c30C2741B8C62730B58B10CeAa55bdDdA7327")}) //png

	return pairs
}

func GetJOE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "JOE",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetJOEPairs(),
	}
}
