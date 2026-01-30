package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetTSDPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2d16af2D7f1edB4bC5DBAdF3ffF04670B4BcD0BB")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x67E395B6ACd948931eeE8F52C7c1Fe537E7f1a7a")}) //png

	return pairs
}

func GetTSD() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "TSD",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetTSDPairs(),
	}
}
