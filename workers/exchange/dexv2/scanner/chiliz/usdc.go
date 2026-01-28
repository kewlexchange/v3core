package chiliz

import (
	"core/constants"
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
		ChainID:     constants.ChilizChainId,
		Token:       "USDC",
		NativeToken: constants.WETH_MAP[constants.ChilizChainId],
		Pairs:       GetUSDCPairs(),
	}
}
