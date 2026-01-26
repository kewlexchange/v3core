package sdk

import (
	"core/sdk/uniswap/constants"
	_ "core/sdk/uniswap/constants"
	"core/sdk/uniswap/entities"
	_ "core/sdk/uniswap/entities"
	_ "core/sdk/uniswap/number"
	_ "core/sdk/uniswap/utils"

	"github.com/ethereum/go-ethereum/common"
)

// Re-export types
type Token = entities.Token
type Pair = entities.Pair
type TokenAmount = entities.TokenAmount
type Trade = entities.Trade
type Route = entities.Route

// Re-export constructors
func NewToken(
	chainID constants.ChainID,
	address common.Address,
	decimals int,
	symbol string,
	name string,
) (*entities.Token, error) {
	//(chainID constants.ChainID, address common.Address, decimals int, symbol, name string)
	return entities.NewToken(chainID, address, decimals, symbol, name)
}

func NewPair(tokenAmountA, tokenAmountB *TokenAmount, address *common.Address) (*entities.Pair, error) {
	return entities.NewPair(tokenAmountA, tokenAmountB, address)
}
