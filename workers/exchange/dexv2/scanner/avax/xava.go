package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetXAVAPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x72c3438cf1c915EcF5D9F17A6eD346B273d5bF71")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x42152bDD72dE8d6767FE3B4E17a221D6985E8B25")}) //png

	return pairs
}

func GetXAVA() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "XAVA",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetXAVAPairs(),
	}
}
