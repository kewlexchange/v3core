package constants

import (
	"core/models"
	"core/utils"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var ChilizChainId = models.Chiliz
var ChilizChain = models.ChainInfoMap[models.Chiliz]
var BSCChainId = models.BSC
var BSCChain = models.ChainInfoMap[models.BSC]
var AVAXChainID = models.Avalanche
var AVAXChain = models.ChainInfoMap[models.Avalanche]

var DEXExchanges = []models.Exchange{
	{
		Name:        "FANX_CHZ",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xE2918AA38088878546c1A18F2F9b1BC83297fdD3"),
		Factory:     utils.AddressFromHex("0xE2918AA38088878546c1A18F2F9b1BC83297fdD3"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5"),
		ChainID:     &ChilizChainId,
		RPC:         &ChilizChain.RPC,
	},
	{
		Name:        "KEWL",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		Factory:     utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5"),

		ChainID: &ChilizChainId,
		RPC:     &ChilizChain.RPC,
	},
	{
		Name:        "DIVISWAP_CHZ",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		Factory:     utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5"),

		ChainID: &ChilizChainId,
		RPC:     &ChilizChain.RPC,
	},

	{
		Name:        "DIVISWAP_BSC",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		Factory:     utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xff6d22b9bd32afb60ac717248c77be1ea16107a5"),

		ChainID: &ChilizChainId,
		RPC:     &ChilizChain.RPC,
	},

	{
		Name:        "KEWL",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0x5636A64B835F4E3821C798fdA16E0bA106357646"),
		Factory:     utils.AddressFromHex("0x5636A64B835F4E3821C798fdA16E0bA106357646"),
		NativeToken: utils.AddressFromHex("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xd99c7f6c65857ac913a8f880a4cb84032ab2fc5b"),
		ChainID:     &BSCChainId,
		RPC:         &BSCChain.RPC,
	},

	{
		Name:        "KEWL",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		Factory:     utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		NativeToken: utils.AddressFromHex("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		StablePair:  utils.AddressFromHex("0xf4003F4efBE8691B60249E6afbD307aBE7758adb"),
		ChainID:     &AVAXChainID,
		RPC:         &AVAXChain.RPC,
	},
}

var FlashContractMap = map[models.ChainID]common.Address{
	models.Chiliz:    common.HexToAddress("0x48b68970abC5de47c6B0526f704b8A95eFeF8aF8"),
	models.BSC:       common.HexToAddress("0x48b68970abC5de47c6B0526f704b8A95eFeF8aF8"),
	models.Avalanche: common.HexToAddress("0x48b68970abC5de47c6B0526f704b8A95eFeF8aF8"),
}

func GetExchangeByName(name string, chainId models.ChainID) (*models.Exchange, error) {
	for i := range DEXExchanges {
		if DEXExchanges[i].Name == name && *DEXExchanges[i].ChainID == chainId {
			return &DEXExchanges[i], nil
		}
	}
	return nil, fmt.Errorf("exchange with name %s not found", name)
}
