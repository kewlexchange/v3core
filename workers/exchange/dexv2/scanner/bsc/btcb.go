package bsc

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetBTCBPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x61EB789d75A95CAa3fF50ed7E47b96c132fEc082")}) //PANCAKE

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xCe4FAfaC517ABFB1F550a8C6713421F2505A259B")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xC7e9d76ba11099AF3F330ff829c5F442d571e057")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc3C16CaE0cf35615b8716a9A8BB0b1257CdbCA53")}) //KEWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x58521373474810915b02FE968D1BCBe35Fc61E09")}) //KEWL

	return pairs
}

func GetBTCB() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.BSCChainId,
		Token:       "BTCB",
		NativeToken: constants.WETH_MAP[constants.BSCChainId][0],
		Pairs:       GetBTCBPairs(),
	}
}
