package scanner

import (
	"core/models"
	"fmt"
)

// FlashSwap, flash swap fırsatını hesaplayıp detayları yazdırır.
func FlashSwap(scan models.ScanParams, tradingPairs []models.TradingPair) {

	for _, tradingPair := range tradingPairs {

		fmt.Println(tradingPair)
	}

}
