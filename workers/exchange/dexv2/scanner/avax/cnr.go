package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetCNRPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2b35C076034e947eDA57Da1C2a5136f4cec5e456")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xC04dE3796716ae5A6788b75DC0d4a1ecE06092d9")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xFad01bF209AaD1E2Dfe98F289d8C8f63BAE71A3d")}) //png

	return pairs
}

func GetCNR() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "CNR",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetCNRPairs(),
	}
}
