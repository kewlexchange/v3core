package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetCOQPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x41aB86EEcBd110a82cA602D032a461f453066F1E")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0xc4D6309D1a68685B67de637489C0091024C7fA4F")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x0CB193E49a14A455F764fE7996997599ba1edaEB")}) //png

	return pairs
}

func GetCOQ() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "COQ",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID],
		Pairs:       GetCOQPairs(),
	}
}
