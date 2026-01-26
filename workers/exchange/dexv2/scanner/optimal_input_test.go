package scanner

import (
	"math/big"
	"testing"
)

func bi(s string) *big.Int {
	v, _ := new(big.Int).SetString(s, 10)
	return v
}

func TestOptimalInputExact_REAL_RESERVES(t *testing.T) {

	/*
		CHEAP (KEWL)
		Price ≈ 42,948,462 PEPPER / CHZ

		CHZ  reserve (R1): 224.750608137612311029
		PEPPER reserve (R2): 9,652,693,102.492406127377066615
	*/

	R1 := bi("224750608137612311029")
	R2 := bi("9652693102492406127377066615")

	/*
		EXPENSIVE (FANX)
		Price ≈ 44,769,334 PEPPER / CHZ

		CHZ  reserve (R4): 4,345,731.336352111551097012
		PEPPER reserve (R3): 194,555,498,510,861.208995752022437008
	*/

	R4 := bi("4345731336352111551097012")
	R3 := bi("194555498510861208995752022437008")

	optimal := OptimalInputExact(R1, R2, R3, R4)

	t.Log("Optimal Input (raw):", optimal.String())

	if optimal.Sign() == 0 {
		t.Log("NO ARB (optimal = 0) — matematiksel olarak doğru sonuç")
	} else {
		t.Log("ARB FOUND — optimal input:", optimal.String())
	}

	// Güvenlik: negatif çıkamaz
	if optimal.Sign() < 0 {
		t.Fatal("Optimal input negatif çıktı — FORMÜL HATALI")
	}
}

func TestOptimalInputExact_WITH_ARB(t *testing.T) {
	// CHZ ve PEPPER rezervlerini (wei cinsinden)
	R1 := bi("1000000000000000000000")        // 1000 CHZ (Cheap CHZ)
	R2 := bi("40000000000000000000000000000") // 40B PEPPER (Cheap PEPPER)

	R4 := bi("100000000000000000000")        // 100 CHZ (Expensive CHZ)
	R3 := bi("3000000000000000000000000000") // 3B PEPPER (Expensive PEPPER)

	optimal := OptimalInputExact(R1, R2, R3, R4)

	t.Log("Optimal Input (raw):", optimal.String())

	if optimal.Sign() == 0 {
		t.Fatal("ARB OLMASI GEREKIRKEN optimal = 0 ÇIKTI ❌")
	} else {
		t.Log("ARB FOUND ✅ optimal input:", optimal.String())
	}
}
