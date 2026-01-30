package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetCRAPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x140CAc5f0e05cBEc857e65353839FddD0D8482C1")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x960FA242468746C59BC32513E2E1e1c24FDFaF3F")}) //png

	return pairs
}

func GetCRA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "CRA",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetCRAPairs(),
	}
}
