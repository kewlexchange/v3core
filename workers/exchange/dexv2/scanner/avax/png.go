package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetPNGPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x3dAF1C6268362214eBB064647555438c6f365F96")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd7538cABBf8605BdE1f4901B47B8D42c61DE0367")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc0ca74965738b8875F3aBc70d9DC59Bff104bC6c")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xa1708efD71E516A2D0Ebd7eC8D877C02d4d2De6d")}) //png

	return pairs
}

func GetPNG() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "PNG",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID],
		Pairs:       GetPNGPairs(),
	}
}
