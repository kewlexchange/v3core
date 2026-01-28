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
		ChainID:     88888,
		Token:       "CHZINU",
		NativeToken: *utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Pairs:       GetCHZINUPairs(),
	}
}

func GetCHZINUComplexPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd7716a59066a431d703f3fd9dd9ab1c5f694282f")}) //CHZINUxPEPPER

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2112edcd1fb2026d46f09085ce26d1fdc0d8c467")}) //CHZINUxPSG

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5f3efab95224dbb5490e8ddc8d2c1daad4c0db37")}) //PEPPERxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5")}) //USDCxCHZ

	return pairs
}
