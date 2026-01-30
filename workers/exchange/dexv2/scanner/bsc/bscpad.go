package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBSCPADPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xba01662E978DE7d67F8FfC937726215eb8995d17")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xAeEF66a27028e4a00D5906E934c55DAb65291f4A")}) //KEWL

	return pairs
}

func GetBSCPAD() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BSCPAD",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBSCPADPairs(),
	}
}
