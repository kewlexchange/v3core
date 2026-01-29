package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetCAKEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x0eD7e52944161450477ee417DE9Cd3a859b14fD0")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xB9600e90414a9C8C128C78d4078784ecDfB03E49")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x3d94d03eb9ea2D4726886aB8Ac9fc0F18355Fd13")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5a918dEcB4FC88cE5C2f54eca3d4EA6fcb259AFc")}) //KEWL

	return pairs
}

func GetCAKE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "CAKE",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetUSDCPairs(),
	}
}
