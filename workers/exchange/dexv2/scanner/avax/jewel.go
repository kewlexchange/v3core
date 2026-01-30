package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetJEWELPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x91F80068c9ac3837e9A4F0bCacca12095A29BeaD")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x6a086fC546E1EB4Bf5bd2852BFd5d7592234BCbc")}) //png

	return pairs
}

func GetJEWEL() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "JEWEL",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetJEWELPairs(),
	}
}
