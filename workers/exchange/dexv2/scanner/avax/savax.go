package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetSavaxPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4b946c91C2B1a7d7C40FB3C130CdfBaf8389094d")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4E9A38F05c38106C1cf5c145Df24959ec50ff70D")}) //png

	return pairs
}

func GetSAVAX() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "SAVAX",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetSavaxPairs(),
	}
}
