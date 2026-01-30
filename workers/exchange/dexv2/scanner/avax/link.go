package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetLinkPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x6F3a0C89f611Ef5dC9d96650324ac633D02265D3")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5875c368Cddd5FB9Bf2f410666ca5aad236DAbD4")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xF34B86ddaB4B7d0Fb5269b074d8Bd2DE397D2dDC")}) //png

	return pairs
}

func GetLINK() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "LINK",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetLinkPairs(),
	}
}
