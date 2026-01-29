package chiliz

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetComplexPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd7716a59066a431d703f3fd9dd9ab1c5f694282f")}) //CHZINUxPEPPER

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x2112edcd1fb2026d46f09085ce26d1fdc0d8c467")}) //CHZINUxPSG

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0")}) //CHZINUxCHZ FANX

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x5f3efab95224dbb5490e8ddc8d2c1daad4c0db37")}) //PEPPERxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5")}) //USDCxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x14a634bf2d5be1c6ad7790d958e748174d8a2d43")}) //USDTxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xea844079241c84fae62648c380a38b913d86e7cf")}) //PSGxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf4f4524e3840f5116af077f0d8b36e2d809ade98")}) // BARxCHZ

	//DIVISWAP
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x59ae0cff65e648fecec8e539a6f6c89c337a48f1")}) //PEPPERxCHZ DSWAP

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xb0a8310f11be8dfeea4e200b9935b815f3faa2fa")}) //DSWAPxCHZ DSWAP

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x3159a90f80fa4aeccc044923b7a504a98417145d")}) //DSWAPxPEPPER DSWAP

	//KEWL
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xf4a49ce9f40d1818b81321f54371a245a98a8618")}) //wCHZxCHZINU

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xd7457eb347e877a6803104edcc7754148d090ea9")}) // KWLxBAR

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x316d8b195a6736c1251e9495b4653e9d5c65c8f1")}) // KWLxUSDC

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x22ce65b596c41d9893d4e1c91737e52a93c8ca0d")}) // KWLxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xaac2bb13fe5735e5683da057f60d6648af2c8092")}) // TBTxCHZ

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x15f809db562c071d5600e6e4d48fd0798cd4b6a1")}) // TBTxKWL

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xddf175F2688EbcBca1dA4B87D6e944a88f55a034")}) // KWLxCHZINU

	return pairs
}

func GetComplex() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.ChilizChainId,
		Token:       "CHZ",
		NativeToken: constants.WETH_MAP[constants.ChilizChainId][0],
		Pairs:       GetComplexPairs(),
	}
}
