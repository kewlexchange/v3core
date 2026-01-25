package chiliz

import (
	"core/models"
	"core/utils"
)

func GetPepperPairs() (pairs []models.Pair) {

	pairs = append(pairs, models.Pair{
		Pair: "0x5F3efAB95224dBb5490E8DDc8D2C1dAAd4c0db37"})

	pairs = append(pairs, models.Pair{
		Pair: "0x59ae0cff65e648fecec8e539a6f6c89c337a48f1"})

	pairs = append(pairs, models.Pair{
		Pair: "0xd83bb0b94132e7e72d7d5bfb481378ca770fc614"})
	return pairs
}

func GetPEPPER() (token models.ScanParams) {
	return models.ScanParams{
		Token:       "PEPPER",
		NativeToken: *utils.AddressFromHex("0x677F7e16C7Dd57be1D4C8aD1244883214953DC47"),
		Pairs:       GetPepperPairs(),
	}
}
