package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDCPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x57c68F0e9E6bE81FB4D9a46e1910C5DAA9c4cfeb")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x06cd679121ec37b0a2fd673d4976b09d81791856")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4b56EE85949164519bBa2e454269dE156930a4C8")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xa439f6415A9DCAb607844732a675f76Dc5A0F9B9")}) //KEWL

	return pairs
}

func GetUSDC() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "USDC",
		NativeToken: constants.WETH_MAP[constants.BSCChainId],
		Pairs:       GetUSDCPairs(),
	}
}
