package chiliz

import (
	"core/constants"
	"core/models"

	"github.com/ethereum/go-ethereum/common"
)

func GetChilizCycle() (cycle []models.Cycle) {
	cycle = []models.Cycle{
		{ //CHZINU -> PEPPER -> CHZ
			InputToken: constants.WETH_MAP[constants.ChilizChainId],
			Hops: []models.Hop{ //CHZINU -> PEPPER -> CHZ
				{
					Pair:        common.HexToAddress("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0"), //CHZINUxwCHZ
					OutputToken: common.HexToAddress("0xf3928e7871eb136dd6648ad08aeef6b6ea893001"), //CHZINU
				},
				{
					Pair:        common.HexToAddress("0xd7716a59066a431d703f3fd9dd9ab1c5f694282f"), //CHZINUxPEPPER
					OutputToken: common.HexToAddress("0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67"), //PEPPER
				},
				{
					Pair:        common.HexToAddress("0x5f3efab95224dbb5490e8ddc8d2c1daad4c0db37"), //PEPPERxCHZ
					OutputToken: constants.WETH_MAP[constants.ChilizChainId],                       //CHZ
				},
			},
		},
		{ //CHZINU -> PSG -> CHZ
			InputToken: constants.WETH_MAP[constants.ChilizChainId],
			Hops: []models.Hop{
				{
					Pair:        common.HexToAddress("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0"), //CHZINUxwCHZ
					OutputToken: common.HexToAddress("0xf3928e7871eb136dd6648ad08aeef6b6ea893001"), //CHZINU
				},
				{
					Pair:        common.HexToAddress("0x2112edcd1fb2026d46f09085ce26d1fdc0d8c467"), //CHZINUxPSG
					OutputToken: common.HexToAddress("0x476ef844b3e8318b3bc887a7db07a1a0fede5557"), //PSG
				},
				{
					Pair:        common.HexToAddress("0xea844079241c84fae62648c380a38b913d86e7cf"), //PSGxCHZ
					OutputToken: constants.WETH_MAP[constants.ChilizChainId],                       //CHZ
				},
			},
		},

		{
			InputToken: constants.WETH_MAP[constants.ChilizChainId],
			Hops: []models.Hop{ //CHZINU -> PEPPERxCHZINU FAN -> DSWAP PEPPERxCHZ
				{
					Pair:        common.HexToAddress("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0"), //CHZINUxwCHZ
					OutputToken: common.HexToAddress("0xf3928e7871eb136dd6648ad08aeef6b6ea893001"), //CHZINU
				},
				{
					Pair:        common.HexToAddress("0xd7716a59066a431d703f3fd9dd9ab1c5f694282f"), //CHZINUxPEPPER
					OutputToken: common.HexToAddress("0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67"), //PEPPER
				},
				{
					Pair:        common.HexToAddress("0x59ae0cff65e648fecec8e539a6f6c89c337a48f1"), //DSWAP PEPPER CHZ
					OutputToken: constants.WETH_MAP[constants.ChilizChainId],                       //CHZ
				},
			},
		},

		{
			InputToken: constants.WETH_MAP[constants.ChilizChainId],
			Hops: []models.Hop{ //CHZINU -> PEPPERxCHZINU -> FANX PEPPER CHZ
				{
					Pair:        common.HexToAddress("0xf9168b29f8534a449b7eb796fac8c60fcaed5af0"), //CHZINUxwCHZ
					OutputToken: common.HexToAddress("0xf3928e7871eb136dd6648ad08aeef6b6ea893001"), //CHZINU
				},
				{
					Pair:        common.HexToAddress("0xd7716a59066a431d703f3fd9dd9ab1c5f694282f"), //CHZINUxPEPPER
					OutputToken: common.HexToAddress("0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67"), //PEPPER
				},
				{
					Pair:        common.HexToAddress("0x5f3efab95224dbb5490e8ddc8d2c1daad4c0db37"), //FANX PEPPER CHZ
					OutputToken: constants.WETH_MAP[constants.ChilizChainId],                       //CHZ
				},
			},
		},
	}
	return cycle
}
