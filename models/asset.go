package models

type Asset struct {
	ContractAddress string
	Currency        Currency
	TradingPairs    []TradingPair
}
