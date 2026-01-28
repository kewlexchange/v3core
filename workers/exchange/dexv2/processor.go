package dexv2

import (
	"core/workers/exchange/dexv2/contracts/flash"
	"strings"
	"sync"
)

var (
	flashLock   sync.Mutex
	activeFlash = make(map[string]bool)
)

func flashKey(p flash.SwapFlashParams) string {
	var b strings.Builder
	b.Grow(256)

	b.WriteString(p.DX.String())
	b.WriteByte('|')
	b.WriteString(p.Profit.String())
	b.WriteByte('|')
	b.WriteString(p.Mid.String())
	b.WriteByte('|')
	b.WriteString(p.Out.String())
	b.WriteByte('|')
	b.WriteString(p.Borrow.Hex())
	b.WriteByte('|')
	b.WriteString(p.Output.Hex())
	b.WriteByte('|')

	for _, a := range p.Path {
		b.WriteString(a.Hex())
		b.WriteByte('>')
	}

	return b.String()
}

func tryLockFlash(p flash.SwapFlashParams) bool {
	key := flashKey(p)

	flashLock.Lock()
	defer flashLock.Unlock()

	if activeFlash[key] {
		return false
	}

	activeFlash[key] = true
	return true
}

func unlockFlash(p flash.SwapFlashParams) {
	key := flashKey(p)

	flashLock.Lock()
	delete(activeFlash, key)
	flashLock.Unlock()
}
