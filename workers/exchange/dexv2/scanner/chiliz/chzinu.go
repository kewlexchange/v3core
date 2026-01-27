package chiliz

import (
	"core/models"
	"core/utils"
)

func GetCHZINUPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0")}) //KAYEN

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf4a49ce9f40d1818b81321f54371a245a98a8618")}) //KEWL
	return pairs
}

func GetCHZINU() (token models.ScanParams) {
	return models.ScanParams{
		Token:       "CHZINU",
		NativeToken: *utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Pairs:       GetCHZINUPairs(),
	}
}
