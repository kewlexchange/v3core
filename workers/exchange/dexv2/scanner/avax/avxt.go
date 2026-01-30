package avax

import (
	"core/constants"
	"core/models"
	"core/utils"
)

func GetAVXTPairs() (pairs []models.TradingPair) {

	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x7fF6A435d129fe0EE0104b7936d5476803692853")}) //png
	pairs = append(pairs, models.TradingPair{
		Pair: *utils.AddressFromHex("0x792055e49a6421F7544c5479eCC380bad62Bc7EE")}) //png

	return pairs
}

func GetAVXT() (token models.ScanParams) {
	return models.ScanParams{
		ChainID:     constants.AVAXChainID,
		Token:       "AVXT",
		NativeToken: constants.WETH_MAP[constants.AVAXChainID][0],
		Pairs:       GetAVXTPairs(),
	}
}
