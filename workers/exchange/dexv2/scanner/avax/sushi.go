package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetSushiPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xF62381AFFdfd27Dba91A1Ea2aCf57d426E28c341")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x84fe1a84C0448d1dC4199a40fe53DB6A49ed6037")}) //png
	return pairs
}

func GetSUSHI() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "SUSHI",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetSushiPairs(),
	}
}
