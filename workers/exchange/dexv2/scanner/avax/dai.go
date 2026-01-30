package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetDAIPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x87Dee1cC9FFd464B79e058ba20387c1984aed86a")}) //png

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xbA09679Ab223C6bdaf44D45Ba2d7279959289AB0")}) //png

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x55CF10BFbC6a9deAeB3c7ec0dd96D3C1179CB948")}) //png

	return pairs
}

func GetDAI() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "DAI",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetDAIPairs(),
	}
}
