package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetWETHairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xFE15c2695F1F920da45C30AAE47d11dE51007AF9")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x7c05d54fc5CB6e4Ad87c6f5db3b807C94bB89c52")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xBb060b8d9224e98A0C15B3CEE01fBc58EF4ecA85")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2FdE1c280a623950b10b6483B9a0C23549c9B515")}) //png
	return pairs
}

func GetWETH() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "WETH",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetWETHairs(),
	}
}
