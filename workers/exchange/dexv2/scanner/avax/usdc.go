package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDCPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xa65Ae059810FE62FaE83F828545d0A4aF0B57819")})

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf4003F4efBE8691B60249E6afbD307aBE7758adb")}) //joe

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x0e0100Ab771E9288e0Aa97e11557E6654C3a9665")}) //pangolin

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x6239aE4D661379b71A90c4c79f0a95297342e391")}) //uni

	return pairs
}

func GetUSDC() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "USDC",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetUSDCPairs(),
	}
}
