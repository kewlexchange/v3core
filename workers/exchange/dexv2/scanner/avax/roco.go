package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetROCOPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4a2cb99e8d91f82cf10fb97d43745a1f23e47caa")}) //png

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x8C28394Ed230cD6cAF0DAA0E51680fD57826DEE3")}) //png

	return pairs
}

func GetROCO() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "ROCO",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID],
		Pairs:       GetROCOPairs(),
	}
}
