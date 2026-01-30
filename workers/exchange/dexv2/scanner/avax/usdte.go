package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetUSDTEPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xeD8CBD9F0cE3C6986b22002F03c6475CEb7a6256")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xe28984e1EE8D431346D32BeC9Ec800Efb643eef4")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x09657b445dF5BF0141e3EF0f5276a329fc01DE01")}) //png

	return pairs
}

func GetUSDTE() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "USDTE",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetUSDTEPairs(),
	}
}
