package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetALOTPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x1D07Ab9cca358ECb5c669b7fbA6BD3900df741A5")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x100B6DbC2678930e1823260e61F6DB03d755A0E2")}) //png

	return pairs
}

func GetALOT() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "ALOT",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetALOTPairs(),
	}
}
