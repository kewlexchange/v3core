package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetDPETPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2B34ded9E8D60BF1F11fEAA3a30dff79C2338FB2")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8d54AE6A4E9247E35f4274B274906a99e9aF52f1")}) //KEWL

	return pairs
}

func GetDPET() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "DPET",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetDPETPairs(),
	}
}
