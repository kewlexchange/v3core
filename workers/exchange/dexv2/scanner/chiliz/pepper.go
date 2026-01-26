package chiliz

import (
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
		Token:       "PEPPER",
		NativeToken: *utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Pairs:       GetPepperPairs(),
	}
}
