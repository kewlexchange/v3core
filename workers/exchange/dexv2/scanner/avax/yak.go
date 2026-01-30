package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetYAKPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb5c9e891AF3063004a441BA4FaB4cA3D6DEb5626")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd2F01cd87A43962fD93C21e07c1a420714Cc94C9")}) //png

	return pairs
}

func GetYAK() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "YAK",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetYAKPairs(),
	}
}
