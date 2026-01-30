package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetAAVEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc3e6D9f7a3A5E3e223356383C7C055Ee3F26A34F")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5944f135e4F1E3fA2E5550d4B5170783868cc4fE")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5fBE35bF97f54EC7CA95EdBc8dfAeB21514739a7")}) //png

	return pairs
}

func GetAAAVE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "AAVE",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetAAVEPairs(),
	}
}
