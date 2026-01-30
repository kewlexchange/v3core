package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetSWAPPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xbd640457A0A2DE67fE18D9ad6756Dd2d501F08B6")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5BE4063911D88fd07122671C9F3c94693846787c")}) //png

	return pairs
}

func GetSWAP() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "SWAP",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetSWAPPairs(),
	}
}
