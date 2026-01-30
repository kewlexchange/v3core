package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetSHRAPPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x4b00BDec69f68b572271979930884C3823a12B75")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb8e84E5103eEfFaf1D08654B21db038E18AD75d0")}) //png

	return pairs
}

func GetSHRAP() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "SHRAP",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetSHRAPPairs(),
	}
}
