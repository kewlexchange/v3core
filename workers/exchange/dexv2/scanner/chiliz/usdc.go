package chiliz

import (
	"core/models"
	"core/utils"
)

func GetUSDCPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5")}) //KAYEN

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xFdf741bdEa532c13b673c66942869b413E60F2C8")}) //KEWL
	return pairs
}

func GetUSDC() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     88888,
		Token:       "USDC",
		NativeToken: *utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Pairs:       GetUSDCPairs(),
	}
}
