package constants

import (
	"core/models"
	"core/utils"
	"fmt"
)

var chilizChainId = int64(models.Chiliz)
var chilizChain = models.ChainInfoMap[models.Chiliz]

var DEXExchanges = []models.Exchange{
	{
		Name:        "FANX",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xE2918AA38088878546c1A18F2F9b1BC83297fdD3"),
		Factory:     utils.AddressFromHex("0xE2918AA38088878546c1A18F2F9b1BC83297fdD3"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		ChainID:     &chilizChainId,
		RPC:         &chilizChain.RPC,
	},
	{
		Name:        "KEWL",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		Factory:     utils.AddressFromHex("0xA0BB8f9865f732C277d0C162249A4F6c157ae9D0"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		ChainID:     &chilizChainId,
		RPC:         &chilizChain.RPC,
	},
	{
		Name:        "DIVISWAP",
		Kind:        models.ExchangeKindDEX,
		Protocol:    models.DexProtocolV2,
		Router:      utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		Factory:     utils.AddressFromHex("0xbdd9c322ecf401e09c9d2dca3be46a7e45d48bb1"),
		NativeToken: utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Multicall3:  utils.AddressFromHex("0xcA11bde05977b3631167028862bE2a173976CA11"),
		ChainID:     &chilizChainId,
		RPC:         &chilizChain.RPC,
	},
}

func GetExchangeByName(name string) (*models.Exchange, error) {
	for i := range DEXExchanges {
		if DEXExchanges[i].Name == name {
			return &DEXExchanges[i], nil
		}
	}
	return nil, fmt.Errorf("exchange with name %s not found", name)
}
