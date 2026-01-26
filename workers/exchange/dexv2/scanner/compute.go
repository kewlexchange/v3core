package scanner

import (
	"core/models"
	"math/big"
)

// Fee sabitleri (Uniswap benzeri %0.3 komisyon)
var (
	feeMul = big.NewInt(997)
	feeDiv = big.NewInt(1000)
)

// FlashSwap, flash swap fırsatını hesaplayıp detayları yazdırır.
func FlashSwap(scan models.ScanParams, pairs []models.TradingPair) {

}
