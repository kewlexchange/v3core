package constants

import (
	"core/models"

	"github.com/ethereum/go-ethereum/common"
)

var WETH_MAP = map[models.ChainID]common.Address{
	models.Chiliz:    common.HexToAddress("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
	models.BSC:       common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
	models.Avalanche: common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
}
