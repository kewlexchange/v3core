package blockchains

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
)

var SwapMethodIDs = map[string]struct{}{
	"38ed1739": {}, // swapExactTokensForTokens
	"8803dbee": {}, // swapTokensForExactTokens
	"7ff36ab5": {}, // swapExactETHForTokens
	"18cbafe5": {}, // swapExactTokensForETH
	"4a25d94a": {}, // swapTokensForExactETH
	"fb3bdb41": {}, // swapETHForExactTokens
	"0000189a": {}, // handleOps (örnek)
	"24856bc3": {}, //execute

}

func IsSwapMethod(tx *types.Transaction) bool {
	data := tx.Data()
	if len(data) < 4 {
		return false
	}
	methodID := hex.EncodeToString(data[:4])
	_, ok := SwapMethodIDs[methodID]
	return ok
}
