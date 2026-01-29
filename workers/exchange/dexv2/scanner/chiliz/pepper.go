package chiliz

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetPepperPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5F3efAB95224dBb5490E8DDc8D2C1dAAd4c0db37")}) //KAYEN

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x59ae0cff65e648fecec8e539a6f6c89c337a48f1")}) //DIVISWAP

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd83bb0b94132e7e72d7d5bfb481378ca770fc614")}) //KEWL
	return pairs
}

func GetPEPPER() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.ChilizChainId,
		Token:       "PEPPER",
		NativeToken: constants.WETH_MAP[constants.ChilizChainId][0],
		Pairs:       GetPepperPairs(),
	}
}
