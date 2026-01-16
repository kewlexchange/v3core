// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kewl

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ArbitrageArbitrageParam is an auto generated low-level Go binding around an user-defined struct.
type ArbitrageArbitrageParam struct {
	Amount      *big.Int
	Buy         common.Address
	Sell        common.Address
	InputToken  common.Address
	OutputToken common.Address
}

// BountyClaimParam is an auto generated low-level Go binding around an user-defined struct.
type BountyClaimParam struct {
	BountyId *big.Int
	TaskId   *big.Int
	Params   string
}

// BountyInfo is an auto generated low-level Go binding around an user-defined struct.
type BountyInfo struct {
	Valid               bool
	CanUserClaim        bool
	BountyId            *big.Int
	BountyType          uint8
	RewardAmount        *big.Int
	TokenAmount         *big.Int
	TotalClaims         *big.Int
	CreatedAt           *big.Int
	NextReward          *big.Int
	UserAvailableReward *big.Int
	UserTotalReward     *big.Int
	UserLastClaimDate   *big.Int
	VerifyParam         *big.Int
	BountyName          string
	BountyDescription   string
	BountyToken         common.Address
}

// BountyTBountyPair is an auto generated low-level Go binding around an user-defined struct.
type BountyTBountyPair struct {
	InputPrice  *big.Int
	OutputPrice *big.Int
	Reserve0    *big.Int
	Reserve1    *big.Int
	Price0      *big.Int
	Price1      *big.Int
	KLast       *big.Int
	TotalSupply *big.Int
	UserBalance *big.Int
	Pair        common.Address
	Input       common.Address
	Output      common.Address
}

// BountyUserInfo is an auto generated low-level Go binding around an user-defined struct.
type BountyUserInfo struct {
	Valid      bool
	Registered bool
	UserId     *big.Int
	Lastaccess *big.Int
	Total      *big.Int
	Wallet     common.Address
	Referral   common.Address
	Avatar     string
	Cover      string
	Name       string
	Bio        string
	Twitter    string
	Telegram   string
	Instagram  string
	Youtube    string
	Facebook   string
	Discord    string
	Tiktok     string
	Website    string
	Geohash    string
	Followers  []common.Address
	Followings []common.Address
	Referrals  []common.Address
}

// LibLeaderBoardScoreInfo is an auto generated low-level Go binding around an user-defined struct.
type LibLeaderBoardScoreInfo struct {
	TotalBaseVolume       *big.Int
	TotalQuoteVolume      *big.Int
	UserBaseVolume        *big.Int
	UserQuoteVolume       *big.Int
	UserDailyBaseVolume   *big.Int
	UserDailyQuoteVolume  *big.Int
	TotalDailyBaseVolume  *big.Int
	TotalDailyQuoteVolume *big.Int
	UserScore             *big.Int
}

// LibLeaderBoardUserInfo is an auto generated low-level Go binding around an user-defined struct.
type LibLeaderBoardUserInfo struct {
	Name     string
	Telegram string
	Twitter  string
	User     common.Address
}

// LibLimitOrdersLimitOrderPairInfo is an auto generated low-level Go binding around an user-defined struct.
type LibLimitOrdersLimitOrderPairInfo struct {
	Valid              bool
	Base               common.Address
	Quote              common.Address
	NextOrderId        *big.Int
	BaseDecimals       uint8
	QuoteDecimals      uint8
	PriceDecimals      uint8
	TickSpacing        *big.Int
	LastPrice          *big.Int
	PriceMin           *big.Int
	PriceMax           *big.Int
	Change             *big.Int
	LastPriceTimestamp *big.Int
	BaseVolume         *big.Int
	QuoteVolume        *big.Int
	MinBuy             *big.Int
	MinSell            *big.Int
	MaxBuy             *big.Int
	MaxSell            *big.Int
	PairId             [32]byte
}

// LibLimitOrdersLimitOrderParam is an auto generated low-level Go binding around an user-defined struct.
type LibLimitOrdersLimitOrderParam struct {
	Kind       uint8
	Token0     common.Address
	Token1     common.Address
	Price      *big.Int
	Amount     *big.Int
	Deadline   *big.Int
	Entrypoint []*big.Int
}

// LibLimitOrdersOrder is an auto generated low-level Go binding around an user-defined struct.
type LibLimitOrdersOrder struct {
	Id             *big.Int
	Sequence       *big.Int
	Price          *big.Int
	Amount         *big.Int
	Total          *big.Int
	Filled         *big.Int
	Remaining      *big.Int
	RemainingValue *big.Int
	PriceTotal     *big.Int
	Matched        *big.Int
	Trader         common.Address
	Kind           uint8
	Status         uint8
	CreatedAt      *big.Int
	UpdatedAt      *big.Int
	CancelledAt    *big.Int
	FilledAt       *big.Int
	ClaimedAt      *big.Int
}

// LibLimitOrdersPriceLevel is an auto generated low-level Go binding around an user-defined struct.
type LibLimitOrdersPriceLevel struct {
	Index          *big.Int
	Price          *big.Int
	BaseLiquidity  *big.Int
	QuoteLiquidity *big.Int
	OrderCount     *big.Int
	Exists         bool
}

// OldArbitrageArbitrageParam is an auto generated low-level Go binding around an user-defined struct.
type OldArbitrageArbitrageParam struct {
	Amount      *big.Int
	Buy         common.Address
	Sell        common.Address
	InputToken  common.Address
	OutputToken common.Address
}

// ReflectionLiqudityPool is an auto generated low-level Go binding around an user-defined struct.
type ReflectionLiqudityPool struct {
	Side  bool
	Pair  common.Address
	Input common.Address
}

// ReflectionSwapConfig is an auto generated low-level Go binding around an user-defined struct.
type ReflectionSwapConfig struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Weth9     common.Address
	Input     common.Address
	Receiver  common.Address
}

// RouterTradeStats is an auto generated low-level Go binding around an user-defined struct.
type RouterTradeStats struct {
	TotalTradeBase       *big.Int
	TotalTradeQuote      *big.Int
	TotalDailyTradeBase  *big.Int
	TotalDailyTradeQuote *big.Int
	BaseVolume           []*big.Int
	QuoteVolume          []*big.Int
	BaseVolumeDaily      []*big.Int
	QuoteVolumeDaily     []*big.Int
}

// TPair is an auto generated low-level Go binding around an user-defined struct.
type TPair struct {
	Valid        bool
	Index        *big.Int
	ReserveBase  *big.Int
	ReserveQuote *big.Int
	Pair         common.Address
	Base         TToken
	Quote        TToken
}

// TToken is an auto generated low-level Go binding around an user-defined struct.
type TToken struct {
	Decimals *big.Int
	Token    common.Address
	Name     string
	Symbol   string
}

// TradePairInfo is an auto generated low-level Go binding around an user-defined struct.
type TradePairInfo struct {
	Valid          bool
	Flag           bool
	Reserve0       *big.Int
	Reserve1       *big.Int
	Amount0Out     *big.Int
	Amount1Out     *big.Int
	Token0Decimals *big.Int
	Token1Decimals *big.Int
	Token0         common.Address
	Token1         common.Address
	Pair           common.Address
	Router         common.Address
	Weth           common.Address
}

// TradePairInput is an auto generated low-level Go binding around an user-defined struct.
type TradePairInput struct {
	Flag   bool
	Router common.Address
	Pair   common.Address
	Input  common.Address
	Weth   common.Address
	Amount *big.Int
}

// TradeRouter is an auto generated low-level Go binding around an user-defined struct.
type TradeRouter struct {
	Router common.Address
	Weth   common.Address
	Flag   bool
	Stable bool
}

// TradeSwapParam is an auto generated low-level Go binding around an user-defined struct.
type TradeSwapParam struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Weth9     common.Address
	Wrapper   common.Address
	Pair      common.Address
	Input     common.Address
	Flag      bool
}

// KewlMetaData contains all meta data concerning the Kewl contract.
var KewlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIPAIR\",\"name\":\"buy\",\"type\":\"address\"},{\"internalType\":\"contractIPAIR\",\"name\":\"sell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"}],\"internalType\":\"structArbitrage.ArbitrageParam\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"arbitrage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIPAIR\",\"name\":\"buy\",\"type\":\"address\"},{\"internalType\":\"contractIPAIR\",\"name\":\"sell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"}],\"internalType\":\"structArbitrage.ArbitrageParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"arbitrageAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"fetchAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBOUNTY_TYPE\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canUserClaim\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bountyId\",\"type\":\"uint256\"},{\"internalType\":\"enumBOUNTY_TYPE\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userAvailableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userTotalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLastClaimDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifyParam\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bountyName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bountyDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bountyToken\",\"type\":\"address\"}],\"internalType\":\"structBountyInfo\",\"name\":\"bounty\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"calculateBountyReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bountyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"internalType\":\"structBountyClaimParam\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"claimReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchBounties\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canUserClaim\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bountyId\",\"type\":\"uint256\"},{\"internalType\":\"enumBOUNTY_TYPE\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userAvailableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userTotalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLastClaimDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifyParam\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bountyName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bountyDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bountyToken\",\"type\":\"address\"}],\"internalType\":\"structBountyInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"fetchBountiesInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canUserClaim\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bountyId\",\"type\":\"uint256\"},{\"internalType\":\"enumBOUNTY_TYPE\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userAvailableReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userTotalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLastClaimDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verifyParam\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bountyName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bountyDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bountyToken\",\"type\":\"address\"}],\"internalType\":\"structBountyInfo[]\",\"name\":\"_bounties\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastaccess\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"instagram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"youtube\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"facebook\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discord\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tiktok\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"geohash\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"followers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"followings\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"referrals\",\"type\":\"address[]\"}],\"internalType\":\"structBountyUserInfo\",\"name\":\"bountyUserInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"fetchJackPotInfo\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"jackpotAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getBountyAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"getPairAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"getPriceOfAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inputPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kLast\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"},{\"internalType\":\"contractIPAIR\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"internalType\":\"structBounty.TBountyPair\",\"name\":\"pairInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initBounties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"jackpot\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bountyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kwlToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethToken\",\"type\":\"address\"}],\"name\":\"setBountyTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"setDisableBountyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncBountyPairs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boundtyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bountyToken\",\"type\":\"address\"}],\"name\":\"updateBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"upDown\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PairAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair\",\"name\":\"pairInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth9\",\"type\":\"address\"}],\"name\":\"setWETH9Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PairAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair\",\"name\":\"pairInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth9\",\"type\":\"address\"}],\"name\":\"setWETH9Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fetchLeaderBoardUsers\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getLeaderboardUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userDailyBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userDailyQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userScore\",\"type\":\"uint256\"}],\"internalType\":\"structLibLeaderBoard.ScoreInfo\",\"name\":\"scoreInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fetchLeaderBoardUsers\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getLeaderboardUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userDailyBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userDailyQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyBaseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyQuoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userScore\",\"type\":\"uint256\"}],\"internalType\":\"structLibLeaderBoard.ScoreInfo\",\"name\":\"scoreInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structLibLeaderBoard.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlreadyCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountPriceOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPair\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTick\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTickSpacing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrderOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderIsNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotCancellable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TickStillHasLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"LimitOrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"kind\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pairHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TickInserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pairHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tick\",\"type\":\"uint256\"}],\"name\":\"TickRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"allLimitPairs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLimitPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibLimitOrders.OrderKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"entrypoint\",\"type\":\"uint256[]\"}],\"internalType\":\"structLibLimitOrders.LimitOrderParam\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"createPaidPair\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"createPairOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"fetchUserOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"matched\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelledAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAt\",\"type\":\"uint256\"}],\"internalType\":\"structLibLimitOrders.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPairStats\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"quoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMax\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastPriceTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSell\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"}],\"internalType\":\"structLibLimitOrders.LimitOrderPairInfo[]\",\"name\":\"pairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getLimitPair\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"quoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMax\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastPriceTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSell\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"}],\"internalType\":\"structLibLimitOrders.LimitOrderPairInfo\",\"name\":\"pair\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"}],\"name\":\"getLimitPairByPairId\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"quoteDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceMax\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastPriceTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSell\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"}],\"internalType\":\"structLibLimitOrders.LimitOrderPairInfo\",\"name\":\"pair\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPairId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pairId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"orderBook\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structLibLimitOrders.PriceLevel[]\",\"name\":\"orderbook\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setListingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setVaultConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"}],\"internalType\":\"structOldArbitrage.ArbitrageParam\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"arbitrage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"}],\"internalType\":\"structOldArbitrage.ArbitrageParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"arbitrageAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"fetchAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structReflection.SwapConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"side\",\"type\":\"bool\"},{\"internalType\":\"contractIPAIR\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"input\",\"type\":\"address\"}],\"internalType\":\"structReflection.LiqudityPool[]\",\"name\":\"pools\",\"type\":\"tuple[]\"}],\"name\":\"reflect\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTradeStatsForMultipleTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"totalTrades\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"individualTrades\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startOfDay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"getTradeStatsForMultipleUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalTradeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTradeQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyTradeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDailyTradeQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"baseVolume\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quoteVolume\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"baseVolumeDaily\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quoteVolumeDaily\",\"type\":\"uint256[]\"}],\"internalType\":\"structRouter.TradeStats\",\"name\":\"tradeStats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"getTradeStatsForMultipleUserWithoutTraders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTradeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTradeQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"baseVolume\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quoteVolume\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTradeStatsForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTrades\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"individualTrades\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"pairFor\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTradeStatsForMultipleTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"totalTrades\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"individualTrades\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"getTradeStatsForMultipleUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTradeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTradeQuote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"traders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"baseVolume\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quoteVolume\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"getTradeStatsForMultipleUserWithoutTraders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTradeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTradeQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"baseVolume\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quoteVolume\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTradeStatsForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTrades\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"individualTrades\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"pairFor\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveQuote\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTToken\",\"name\":\"quote\",\"type\":\"tuple\"}],\"internalType\":\"structTPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"fetchAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTrade.PairInput\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchPair\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0Decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Decimals\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"contractIPAIR\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"internalType\":\"structTrade.PairInfo\",\"name\":\"pairInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structTrade.Router[]\",\"name\":\"routers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fetchPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0Decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Decimals\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"contractIPAIR\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"internalType\":\"structTrade.PairInfo[]\",\"name\":\"pairInfo\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pairlist\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0Decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Decimals\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"contractIPAIR\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"internalType\":\"structTrade.PairInfo[]\",\"name\":\"pairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"weth9\",\"type\":\"address[]\"}],\"name\":\"setWrappedAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"internalType\":\"structTrade.SwapParam\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"internalType\":\"structTrade.SwapParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"swapAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structTrade.Router\",\"name\":\"router\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"testPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawCustomERC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawCustomETHAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"name\":\"withdrawERC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"banned_\",\"type\":\"bool\"}],\"name\":\"banAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getContractTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled_\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isKEWL\",\"type\":\"bool\"}],\"name\":\"setKEWL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCoins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrappedTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrappedTokenExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUnderlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newWrappedToken\",\"type\":\"address\"}],\"name\":\"MappingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unwrappedAmount\",\"type\":\"uint256\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrappedAmount\",\"type\":\"uint256\"}],\"name\":\"Wrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"WrappedTokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"createWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"getUnderlyingToWrapped\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"underlyingToWrapped\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedToUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"wrappedTokenFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrappedTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrappedTokenExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUnderlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newWrappedToken\",\"type\":\"address\"}],\"name\":\"MappingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unwrappedAmount\",\"type\":\"uint256\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrappedAmount\",\"type\":\"uint256\"}],\"name\":\"Wrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"WrappedTokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"createWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingToWrapped\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"wrappedTokenFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KewlABI is the input ABI used to generate the binding from.
// Deprecated: Use KewlMetaData.ABI instead.
var KewlABI = KewlMetaData.ABI

// Kewl is an auto generated Go binding around an Ethereum contract.
type Kewl struct {
	KewlCaller     // Read-only binding to the contract
	KewlTransactor // Write-only binding to the contract
	KewlFilterer   // Log filterer for contract events
}

// KewlCaller is an auto generated read-only Go binding around an Ethereum contract.
type KewlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KewlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KewlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KewlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KewlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KewlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KewlSession struct {
	Contract     *Kewl             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KewlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KewlCallerSession struct {
	Contract *KewlCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KewlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KewlTransactorSession struct {
	Contract     *KewlTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KewlRaw is an auto generated low-level Go binding around an Ethereum contract.
type KewlRaw struct {
	Contract *Kewl // Generic contract binding to access the raw methods on
}

// KewlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KewlCallerRaw struct {
	Contract *KewlCaller // Generic read-only contract binding to access the raw methods on
}

// KewlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KewlTransactorRaw struct {
	Contract *KewlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKewl creates a new instance of Kewl, bound to a specific deployed contract.
func NewKewl(address common.Address, backend bind.ContractBackend) (*Kewl, error) {
	contract, err := bindKewl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kewl{KewlCaller: KewlCaller{contract: contract}, KewlTransactor: KewlTransactor{contract: contract}, KewlFilterer: KewlFilterer{contract: contract}}, nil
}

// NewKewlCaller creates a new read-only instance of Kewl, bound to a specific deployed contract.
func NewKewlCaller(address common.Address, caller bind.ContractCaller) (*KewlCaller, error) {
	contract, err := bindKewl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KewlCaller{contract: contract}, nil
}

// NewKewlTransactor creates a new write-only instance of Kewl, bound to a specific deployed contract.
func NewKewlTransactor(address common.Address, transactor bind.ContractTransactor) (*KewlTransactor, error) {
	contract, err := bindKewl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KewlTransactor{contract: contract}, nil
}

// NewKewlFilterer creates a new log filterer instance of Kewl, bound to a specific deployed contract.
func NewKewlFilterer(address common.Address, filterer bind.ContractFilterer) (*KewlFilterer, error) {
	contract, err := bindKewl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KewlFilterer{contract: contract}, nil
}

// bindKewl binds a generic wrapper to an already deployed contract.
func bindKewl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KewlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kewl *KewlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kewl.Contract.KewlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kewl *KewlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.Contract.KewlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kewl *KewlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kewl.Contract.KewlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kewl *KewlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kewl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kewl *KewlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kewl *KewlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kewl.Contract.contract.Transact(opts, method, params...)
}

// AllLimitPairs is a free data retrieval call binding the contract method 0xa76aac69.
//
// Solidity: function allLimitPairs(uint256 index) view returns(bytes32)
func (_Kewl *KewlCaller) AllLimitPairs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "allLimitPairs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllLimitPairs is a free data retrieval call binding the contract method 0xa76aac69.
//
// Solidity: function allLimitPairs(uint256 index) view returns(bytes32)
func (_Kewl *KewlSession) AllLimitPairs(index *big.Int) ([32]byte, error) {
	return _Kewl.Contract.AllLimitPairs(&_Kewl.CallOpts, index)
}

// AllLimitPairs is a free data retrieval call binding the contract method 0xa76aac69.
//
// Solidity: function allLimitPairs(uint256 index) view returns(bytes32)
func (_Kewl *KewlCallerSession) AllLimitPairs(index *big.Int) ([32]byte, error) {
	return _Kewl.Contract.AllLimitPairs(&_Kewl.CallOpts, index)
}

// AllLimitPairsLength is a free data retrieval call binding the contract method 0x758b4842.
//
// Solidity: function allLimitPairsLength() view returns(uint256)
func (_Kewl *KewlCaller) AllLimitPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "allLimitPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllLimitPairsLength is a free data retrieval call binding the contract method 0x758b4842.
//
// Solidity: function allLimitPairsLength() view returns(uint256)
func (_Kewl *KewlSession) AllLimitPairsLength() (*big.Int, error) {
	return _Kewl.Contract.AllLimitPairsLength(&_Kewl.CallOpts)
}

// AllLimitPairsLength is a free data retrieval call binding the contract method 0x758b4842.
//
// Solidity: function allLimitPairsLength() view returns(uint256)
func (_Kewl *KewlCallerSession) AllLimitPairsLength() (*big.Int, error) {
	return _Kewl.Contract.AllLimitPairsLength(&_Kewl.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 index) view returns(address)
func (_Kewl *KewlCaller) AllPairs(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "allPairs", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 index) view returns(address)
func (_Kewl *KewlSession) AllPairs(index *big.Int) (common.Address, error) {
	return _Kewl.Contract.AllPairs(&_Kewl.CallOpts, index)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 index) view returns(address)
func (_Kewl *KewlCallerSession) AllPairs(index *big.Int) (common.Address, error) {
	return _Kewl.Contract.AllPairs(&_Kewl.CallOpts, index)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Kewl *KewlCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Kewl *KewlSession) AllPairsLength() (*big.Int, error) {
	return _Kewl.Contract.AllPairsLength(&_Kewl.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Kewl *KewlCallerSession) AllPairsLength() (*big.Int, error) {
	return _Kewl.Contract.AllPairsLength(&_Kewl.CallOpts)
}

// CalculateBountyReward is a free data retrieval call binding the contract method 0x07c143e1.
//
// Solidity: function calculateBountyReward((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address) bounty, address account) view returns(uint256 totalReward, uint256 availableReward, uint256 lastClaimDate)
func (_Kewl *KewlCaller) CalculateBountyReward(opts *bind.CallOpts, bounty BountyInfo, account common.Address) (struct {
	TotalReward     *big.Int
	AvailableReward *big.Int
	LastClaimDate   *big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "calculateBountyReward", bounty, account)

	outstruct := new(struct {
		TotalReward     *big.Int
		AvailableReward *big.Int
		LastClaimDate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AvailableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastClaimDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateBountyReward is a free data retrieval call binding the contract method 0x07c143e1.
//
// Solidity: function calculateBountyReward((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address) bounty, address account) view returns(uint256 totalReward, uint256 availableReward, uint256 lastClaimDate)
func (_Kewl *KewlSession) CalculateBountyReward(bounty BountyInfo, account common.Address) (struct {
	TotalReward     *big.Int
	AvailableReward *big.Int
	LastClaimDate   *big.Int
}, error) {
	return _Kewl.Contract.CalculateBountyReward(&_Kewl.CallOpts, bounty, account)
}

// CalculateBountyReward is a free data retrieval call binding the contract method 0x07c143e1.
//
// Solidity: function calculateBountyReward((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address) bounty, address account) view returns(uint256 totalReward, uint256 availableReward, uint256 lastClaimDate)
func (_Kewl *KewlCallerSession) CalculateBountyReward(bounty BountyInfo, account common.Address) (struct {
	TotalReward     *big.Int
	AvailableReward *big.Int
	LastClaimDate   *big.Int
}, error) {
	return _Kewl.Contract.CalculateBountyReward(&_Kewl.CallOpts, bounty, account)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlSession) FeeTo() (common.Address, error) {
	return _Kewl.Contract.FeeTo(&_Kewl.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlCallerSession) FeeTo() (common.Address, error) {
	return _Kewl.Contract.FeeTo(&_Kewl.CallOpts)
}

// FeeTo0 is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlCaller) FeeTo0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "feeTo0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo0 is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlSession) FeeTo0() (common.Address, error) {
	return _Kewl.Contract.FeeTo0(&_Kewl.CallOpts)
}

// FeeTo0 is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Kewl *KewlCallerSession) FeeTo0() (common.Address, error) {
	return _Kewl.Contract.FeeTo0(&_Kewl.CallOpts)
}

// FetchAmount is a free data retrieval call binding the contract method 0xa6ed1ad6.
//
// Solidity: function fetchAmount(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) FetchAmount(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchAmount", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchAmount is a free data retrieval call binding the contract method 0xa6ed1ad6.
//
// Solidity: function fetchAmount(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) FetchAmount(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmount(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchAmount is a free data retrieval call binding the contract method 0xa6ed1ad6.
//
// Solidity: function fetchAmount(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) FetchAmount(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmount(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchAmountOut is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) FetchAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchAmountOut is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) FetchAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchAmountOut is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) FetchAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchAmountOut0 is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) FetchAmountOut0(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchAmountOut0", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchAmountOut0 is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) FetchAmountOut0(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmountOut0(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchAmountOut0 is a free data retrieval call binding the contract method 0x0029d898.
//
// Solidity: function fetchAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) FetchAmountOut0(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.FetchAmountOut0(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// FetchBounties is a free data retrieval call binding the contract method 0x0f25c633.
//
// Solidity: function fetchBounties() view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[])
func (_Kewl *KewlCaller) FetchBounties(opts *bind.CallOpts) ([]BountyInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchBounties")

	if err != nil {
		return *new([]BountyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BountyInfo)).(*[]BountyInfo)

	return out0, err

}

// FetchBounties is a free data retrieval call binding the contract method 0x0f25c633.
//
// Solidity: function fetchBounties() view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[])
func (_Kewl *KewlSession) FetchBounties() ([]BountyInfo, error) {
	return _Kewl.Contract.FetchBounties(&_Kewl.CallOpts)
}

// FetchBounties is a free data retrieval call binding the contract method 0x0f25c633.
//
// Solidity: function fetchBounties() view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[])
func (_Kewl *KewlCallerSession) FetchBounties() ([]BountyInfo, error) {
	return _Kewl.Contract.FetchBounties(&_Kewl.CallOpts)
}

// FetchBountiesInfo is a free data retrieval call binding the contract method 0x4b1b6f51.
//
// Solidity: function fetchBountiesInfo(address _account) view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[] _bounties, (bool,bool,uint256,uint256,uint256,address,address,string,string,string,string,string,string,string,string,string,string,string,string,string,address[],address[],address[]) bountyUserInfo)
func (_Kewl *KewlCaller) FetchBountiesInfo(opts *bind.CallOpts, _account common.Address) (struct {
	Bounties       []BountyInfo
	BountyUserInfo BountyUserInfo
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchBountiesInfo", _account)

	outstruct := new(struct {
		Bounties       []BountyInfo
		BountyUserInfo BountyUserInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bounties = *abi.ConvertType(out[0], new([]BountyInfo)).(*[]BountyInfo)
	outstruct.BountyUserInfo = *abi.ConvertType(out[1], new(BountyUserInfo)).(*BountyUserInfo)

	return *outstruct, err

}

// FetchBountiesInfo is a free data retrieval call binding the contract method 0x4b1b6f51.
//
// Solidity: function fetchBountiesInfo(address _account) view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[] _bounties, (bool,bool,uint256,uint256,uint256,address,address,string,string,string,string,string,string,string,string,string,string,string,string,string,address[],address[],address[]) bountyUserInfo)
func (_Kewl *KewlSession) FetchBountiesInfo(_account common.Address) (struct {
	Bounties       []BountyInfo
	BountyUserInfo BountyUserInfo
}, error) {
	return _Kewl.Contract.FetchBountiesInfo(&_Kewl.CallOpts, _account)
}

// FetchBountiesInfo is a free data retrieval call binding the contract method 0x4b1b6f51.
//
// Solidity: function fetchBountiesInfo(address _account) view returns((bool,bool,uint256,uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,address)[] _bounties, (bool,bool,uint256,uint256,uint256,address,address,string,string,string,string,string,string,string,string,string,string,string,string,string,address[],address[],address[]) bountyUserInfo)
func (_Kewl *KewlCallerSession) FetchBountiesInfo(_account common.Address) (struct {
	Bounties       []BountyInfo
	BountyUserInfo BountyUserInfo
}, error) {
	return _Kewl.Contract.FetchBountiesInfo(&_Kewl.CallOpts, _account)
}

// FetchJackPotInfo is a free data retrieval call binding the contract method 0x18e47bf1.
//
// Solidity: function fetchJackPotInfo(uint256 limit) view returns(address[] receivers, uint256 jackpotAmount)
func (_Kewl *KewlCaller) FetchJackPotInfo(opts *bind.CallOpts, limit *big.Int) (struct {
	Receivers     []common.Address
	JackpotAmount *big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchJackPotInfo", limit)

	outstruct := new(struct {
		Receivers     []common.Address
		JackpotAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receivers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.JackpotAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FetchJackPotInfo is a free data retrieval call binding the contract method 0x18e47bf1.
//
// Solidity: function fetchJackPotInfo(uint256 limit) view returns(address[] receivers, uint256 jackpotAmount)
func (_Kewl *KewlSession) FetchJackPotInfo(limit *big.Int) (struct {
	Receivers     []common.Address
	JackpotAmount *big.Int
}, error) {
	return _Kewl.Contract.FetchJackPotInfo(&_Kewl.CallOpts, limit)
}

// FetchJackPotInfo is a free data retrieval call binding the contract method 0x18e47bf1.
//
// Solidity: function fetchJackPotInfo(uint256 limit) view returns(address[] receivers, uint256 jackpotAmount)
func (_Kewl *KewlCallerSession) FetchJackPotInfo(limit *big.Int) (struct {
	Receivers     []common.Address
	JackpotAmount *big.Int
}, error) {
	return _Kewl.Contract.FetchJackPotInfo(&_Kewl.CallOpts, limit)
}

// FetchLeaderBoardUsers is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlCaller) FetchLeaderBoardUsers(opts *bind.CallOpts) ([]LibLeaderBoardUserInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchLeaderBoardUsers")

	if err != nil {
		return *new([]LibLeaderBoardUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibLeaderBoardUserInfo)).(*[]LibLeaderBoardUserInfo)

	return out0, err

}

// FetchLeaderBoardUsers is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlSession) FetchLeaderBoardUsers() ([]LibLeaderBoardUserInfo, error) {
	return _Kewl.Contract.FetchLeaderBoardUsers(&_Kewl.CallOpts)
}

// FetchLeaderBoardUsers is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlCallerSession) FetchLeaderBoardUsers() ([]LibLeaderBoardUserInfo, error) {
	return _Kewl.Contract.FetchLeaderBoardUsers(&_Kewl.CallOpts)
}

// FetchLeaderBoardUsers0 is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlCaller) FetchLeaderBoardUsers0(opts *bind.CallOpts) ([]LibLeaderBoardUserInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchLeaderBoardUsers0")

	if err != nil {
		return *new([]LibLeaderBoardUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibLeaderBoardUserInfo)).(*[]LibLeaderBoardUserInfo)

	return out0, err

}

// FetchLeaderBoardUsers0 is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlSession) FetchLeaderBoardUsers0() ([]LibLeaderBoardUserInfo, error) {
	return _Kewl.Contract.FetchLeaderBoardUsers0(&_Kewl.CallOpts)
}

// FetchLeaderBoardUsers0 is a free data retrieval call binding the contract method 0x742c2afa.
//
// Solidity: function fetchLeaderBoardUsers() view returns((string,string,string,address)[])
func (_Kewl *KewlCallerSession) FetchLeaderBoardUsers0() ([]LibLeaderBoardUserInfo, error) {
	return _Kewl.Contract.FetchLeaderBoardUsers0(&_Kewl.CallOpts)
}

// FetchPair is a free data retrieval call binding the contract method 0xba6c0c09.
//
// Solidity: function fetchPair((bool,address,address,address,address,uint256) params) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address) pairInfo)
func (_Kewl *KewlCaller) FetchPair(opts *bind.CallOpts, params TradePairInput) (TradePairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchPair", params)

	if err != nil {
		return *new(TradePairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TradePairInfo)).(*TradePairInfo)

	return out0, err

}

// FetchPair is a free data retrieval call binding the contract method 0xba6c0c09.
//
// Solidity: function fetchPair((bool,address,address,address,address,uint256) params) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address) pairInfo)
func (_Kewl *KewlSession) FetchPair(params TradePairInput) (TradePairInfo, error) {
	return _Kewl.Contract.FetchPair(&_Kewl.CallOpts, params)
}

// FetchPair is a free data retrieval call binding the contract method 0xba6c0c09.
//
// Solidity: function fetchPair((bool,address,address,address,address,uint256) params) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address) pairInfo)
func (_Kewl *KewlCallerSession) FetchPair(params TradePairInput) (TradePairInfo, error) {
	return _Kewl.Contract.FetchPair(&_Kewl.CallOpts, params)
}

// FetchPairs is a free data retrieval call binding the contract method 0xf80cdc6d.
//
// Solidity: function fetchPairs((address,address,bool,bool)[] routers, address wrapper, address base, address quote, uint256 amount) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairInfo)
func (_Kewl *KewlCaller) FetchPairs(opts *bind.CallOpts, routers []TradeRouter, wrapper common.Address, base common.Address, quote common.Address, amount *big.Int) ([]TradePairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchPairs", routers, wrapper, base, quote, amount)

	if err != nil {
		return *new([]TradePairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TradePairInfo)).(*[]TradePairInfo)

	return out0, err

}

// FetchPairs is a free data retrieval call binding the contract method 0xf80cdc6d.
//
// Solidity: function fetchPairs((address,address,bool,bool)[] routers, address wrapper, address base, address quote, uint256 amount) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairInfo)
func (_Kewl *KewlSession) FetchPairs(routers []TradeRouter, wrapper common.Address, base common.Address, quote common.Address, amount *big.Int) ([]TradePairInfo, error) {
	return _Kewl.Contract.FetchPairs(&_Kewl.CallOpts, routers, wrapper, base, quote, amount)
}

// FetchPairs is a free data retrieval call binding the contract method 0xf80cdc6d.
//
// Solidity: function fetchPairs((address,address,bool,bool)[] routers, address wrapper, address base, address quote, uint256 amount) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairInfo)
func (_Kewl *KewlCallerSession) FetchPairs(routers []TradeRouter, wrapper common.Address, base common.Address, quote common.Address, amount *big.Int) ([]TradePairInfo, error) {
	return _Kewl.Contract.FetchPairs(&_Kewl.CallOpts, routers, wrapper, base, quote, amount)
}

// FetchUserOrders is a free data retrieval call binding the contract method 0xc8b6d2a8.
//
// Solidity: function fetchUserOrders(bytes32 pairId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256)[] orders)
func (_Kewl *KewlCaller) FetchUserOrders(opts *bind.CallOpts, pairId [32]byte, user common.Address) ([]LibLimitOrdersOrder, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "fetchUserOrders", pairId, user)

	if err != nil {
		return *new([]LibLimitOrdersOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibLimitOrdersOrder)).(*[]LibLimitOrdersOrder)

	return out0, err

}

// FetchUserOrders is a free data retrieval call binding the contract method 0xc8b6d2a8.
//
// Solidity: function fetchUserOrders(bytes32 pairId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256)[] orders)
func (_Kewl *KewlSession) FetchUserOrders(pairId [32]byte, user common.Address) ([]LibLimitOrdersOrder, error) {
	return _Kewl.Contract.FetchUserOrders(&_Kewl.CallOpts, pairId, user)
}

// FetchUserOrders is a free data retrieval call binding the contract method 0xc8b6d2a8.
//
// Solidity: function fetchUserOrders(bytes32 pairId, address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256)[] orders)
func (_Kewl *KewlCallerSession) FetchUserOrders(pairId [32]byte, user common.Address) ([]LibLimitOrdersOrder, error) {
	return _Kewl.Contract.FetchUserOrders(&_Kewl.CallOpts, pairId, user)
}

// GetAllPairStats is a free data retrieval call binding the contract method 0x65eec7ad.
//
// Solidity: function getAllPairStats() view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] pairs)
func (_Kewl *KewlCaller) GetAllPairStats(opts *bind.CallOpts) ([]LibLimitOrdersLimitOrderPairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAllPairStats")

	if err != nil {
		return *new([]LibLimitOrdersLimitOrderPairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibLimitOrdersLimitOrderPairInfo)).(*[]LibLimitOrdersLimitOrderPairInfo)

	return out0, err

}

// GetAllPairStats is a free data retrieval call binding the contract method 0x65eec7ad.
//
// Solidity: function getAllPairStats() view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] pairs)
func (_Kewl *KewlSession) GetAllPairStats() ([]LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetAllPairStats(&_Kewl.CallOpts)
}

// GetAllPairStats is a free data retrieval call binding the contract method 0x65eec7ad.
//
// Solidity: function getAllPairStats() view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] pairs)
func (_Kewl *KewlCallerSession) GetAllPairStats() ([]LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetAllPairStats(&_Kewl.CallOpts)
}

// GetAllPairs is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlCaller) GetAllPairs(opts *bind.CallOpts) ([]TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAllPairs")

	if err != nil {
		return *new([]TPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]TPair)).(*[]TPair)

	return out0, err

}

// GetAllPairs is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlSession) GetAllPairs() ([]TPair, error) {
	return _Kewl.Contract.GetAllPairs(&_Kewl.CallOpts)
}

// GetAllPairs is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlCallerSession) GetAllPairs() ([]TPair, error) {
	return _Kewl.Contract.GetAllPairs(&_Kewl.CallOpts)
}

// GetAllPairs0 is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlCaller) GetAllPairs0(opts *bind.CallOpts) ([]TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAllPairs0")

	if err != nil {
		return *new([]TPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]TPair)).(*[]TPair)

	return out0, err

}

// GetAllPairs0 is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlSession) GetAllPairs0() ([]TPair, error) {
	return _Kewl.Contract.GetAllPairs0(&_Kewl.CallOpts)
}

// GetAllPairs0 is a free data retrieval call binding the contract method 0xf800ece9.
//
// Solidity: function getAllPairs() view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string))[])
func (_Kewl *KewlCallerSession) GetAllPairs0() ([]TPair, error) {
	return _Kewl.Contract.GetAllPairs0(&_Kewl.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountIn(&_Kewl.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountIn(&_Kewl.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn0 is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlCaller) GetAmountIn0(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountIn0", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn0 is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlSession) GetAmountIn0(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountIn0(&_Kewl.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn0 is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Kewl *KewlCallerSession) GetAmountIn0(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountIn0(&_Kewl.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut0 is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) GetAmountOut0(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountOut0", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut0 is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) GetAmountOut0(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountOut0(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut0 is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) GetAmountOut0(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetAmountOut0(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsIn(&_Kewl.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsIn(&_Kewl.CallOpts, amountOut, path)
}

// GetAmountsIn0 is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCaller) GetAmountsIn0(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountsIn0", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn0 is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlSession) GetAmountsIn0(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsIn0(&_Kewl.CallOpts, amountOut, path)
}

// GetAmountsIn0 is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCallerSession) GetAmountsIn0(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsIn0(&_Kewl.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsOut(&_Kewl.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsOut(&_Kewl.CallOpts, amountIn, path)
}

// GetAmountsOut0 is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCaller) GetAmountsOut0(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getAmountsOut0", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut0 is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlSession) GetAmountsOut0(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsOut0(&_Kewl.CallOpts, amountIn, path)
}

// GetAmountsOut0 is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Kewl *KewlCallerSession) GetAmountsOut0(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Kewl.Contract.GetAmountsOut0(&_Kewl.CallOpts, amountIn, path)
}

// GetBountyAmountOut is a free data retrieval call binding the contract method 0x76a4fbef.
//
// Solidity: function getBountyAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCaller) GetBountyAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getBountyAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBountyAmountOut is a free data retrieval call binding the contract method 0x76a4fbef.
//
// Solidity: function getBountyAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlSession) GetBountyAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetBountyAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetBountyAmountOut is a free data retrieval call binding the contract method 0x76a4fbef.
//
// Solidity: function getBountyAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Kewl *KewlCallerSession) GetBountyAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Kewl.Contract.GetBountyAmountOut(&_Kewl.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Kewl *KewlCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Kewl *KewlSession) GetContractBalance() (*big.Int, error) {
	return _Kewl.Contract.GetContractBalance(&_Kewl.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Kewl *KewlCallerSession) GetContractBalance() (*big.Int, error) {
	return _Kewl.Contract.GetContractBalance(&_Kewl.CallOpts)
}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address _tokenAddress) view returns(uint256)
func (_Kewl *KewlCaller) GetContractTokenBalance(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getContractTokenBalance", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address _tokenAddress) view returns(uint256)
func (_Kewl *KewlSession) GetContractTokenBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _Kewl.Contract.GetContractTokenBalance(&_Kewl.CallOpts, _tokenAddress)
}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address _tokenAddress) view returns(uint256)
func (_Kewl *KewlCallerSession) GetContractTokenBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _Kewl.Contract.GetContractTokenBalance(&_Kewl.CallOpts, _tokenAddress)
}

// GetLeaderboardUserInfo is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlCaller) GetLeaderboardUserInfo(opts *bind.CallOpts, user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getLeaderboardUserInfo", user, baseToken, quoteToken)

	outstruct := new(struct {
		UserInfo  LibLeaderBoardUserInfo
		ScoreInfo LibLeaderBoardScoreInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserInfo = *abi.ConvertType(out[0], new(LibLeaderBoardUserInfo)).(*LibLeaderBoardUserInfo)
	outstruct.ScoreInfo = *abi.ConvertType(out[1], new(LibLeaderBoardScoreInfo)).(*LibLeaderBoardScoreInfo)

	return *outstruct, err

}

// GetLeaderboardUserInfo is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlSession) GetLeaderboardUserInfo(user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	return _Kewl.Contract.GetLeaderboardUserInfo(&_Kewl.CallOpts, user, baseToken, quoteToken)
}

// GetLeaderboardUserInfo is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlCallerSession) GetLeaderboardUserInfo(user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	return _Kewl.Contract.GetLeaderboardUserInfo(&_Kewl.CallOpts, user, baseToken, quoteToken)
}

// GetLeaderboardUserInfo0 is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlCaller) GetLeaderboardUserInfo0(opts *bind.CallOpts, user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getLeaderboardUserInfo0", user, baseToken, quoteToken)

	outstruct := new(struct {
		UserInfo  LibLeaderBoardUserInfo
		ScoreInfo LibLeaderBoardScoreInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserInfo = *abi.ConvertType(out[0], new(LibLeaderBoardUserInfo)).(*LibLeaderBoardUserInfo)
	outstruct.ScoreInfo = *abi.ConvertType(out[1], new(LibLeaderBoardScoreInfo)).(*LibLeaderBoardScoreInfo)

	return *outstruct, err

}

// GetLeaderboardUserInfo0 is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlSession) GetLeaderboardUserInfo0(user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	return _Kewl.Contract.GetLeaderboardUserInfo0(&_Kewl.CallOpts, user, baseToken, quoteToken)
}

// GetLeaderboardUserInfo0 is a free data retrieval call binding the contract method 0x90078e98.
//
// Solidity: function getLeaderboardUserInfo(address user, address baseToken, address quoteToken) view returns((string,string,string,address) userInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) scoreInfo)
func (_Kewl *KewlCallerSession) GetLeaderboardUserInfo0(user common.Address, baseToken common.Address, quoteToken common.Address) (struct {
	UserInfo  LibLeaderBoardUserInfo
	ScoreInfo LibLeaderBoardScoreInfo
}, error) {
	return _Kewl.Contract.GetLeaderboardUserInfo0(&_Kewl.CallOpts, user, baseToken, quoteToken)
}

// GetLimitPair is a free data retrieval call binding the contract method 0x1c3ec7ef.
//
// Solidity: function getLimitPair(uint256 index) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlCaller) GetLimitPair(opts *bind.CallOpts, index *big.Int) (LibLimitOrdersLimitOrderPairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getLimitPair", index)

	if err != nil {
		return *new(LibLimitOrdersLimitOrderPairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibLimitOrdersLimitOrderPairInfo)).(*LibLimitOrdersLimitOrderPairInfo)

	return out0, err

}

// GetLimitPair is a free data retrieval call binding the contract method 0x1c3ec7ef.
//
// Solidity: function getLimitPair(uint256 index) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlSession) GetLimitPair(index *big.Int) (LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetLimitPair(&_Kewl.CallOpts, index)
}

// GetLimitPair is a free data retrieval call binding the contract method 0x1c3ec7ef.
//
// Solidity: function getLimitPair(uint256 index) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlCallerSession) GetLimitPair(index *big.Int) (LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetLimitPair(&_Kewl.CallOpts, index)
}

// GetLimitPairByPairId is a free data retrieval call binding the contract method 0xc4a903fa.
//
// Solidity: function getLimitPairByPairId(bytes32 pairId) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlCaller) GetLimitPairByPairId(opts *bind.CallOpts, pairId [32]byte) (LibLimitOrdersLimitOrderPairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getLimitPairByPairId", pairId)

	if err != nil {
		return *new(LibLimitOrdersLimitOrderPairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibLimitOrdersLimitOrderPairInfo)).(*LibLimitOrdersLimitOrderPairInfo)

	return out0, err

}

// GetLimitPairByPairId is a free data retrieval call binding the contract method 0xc4a903fa.
//
// Solidity: function getLimitPairByPairId(bytes32 pairId) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlSession) GetLimitPairByPairId(pairId [32]byte) (LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetLimitPairByPairId(&_Kewl.CallOpts, pairId)
}

// GetLimitPairByPairId is a free data retrieval call binding the contract method 0xc4a903fa.
//
// Solidity: function getLimitPairByPairId(bytes32 pairId) view returns((bool,address,address,uint256,uint8,uint8,uint8,int24,uint256,uint256,uint256,int256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) pair)
func (_Kewl *KewlCallerSession) GetLimitPairByPairId(pairId [32]byte) (LibLimitOrdersLimitOrderPairInfo, error) {
	return _Kewl.Contract.GetLimitPairByPairId(&_Kewl.CallOpts, pairId)
}

// GetListingFee is a free data retrieval call binding the contract method 0xb8fe6abe.
//
// Solidity: function getListingFee() view returns(uint256)
func (_Kewl *KewlCaller) GetListingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getListingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingFee is a free data retrieval call binding the contract method 0xb8fe6abe.
//
// Solidity: function getListingFee() view returns(uint256)
func (_Kewl *KewlSession) GetListingFee() (*big.Int, error) {
	return _Kewl.Contract.GetListingFee(&_Kewl.CallOpts)
}

// GetListingFee is a free data retrieval call binding the contract method 0xb8fe6abe.
//
// Solidity: function getListingFee() view returns(uint256)
func (_Kewl *KewlCallerSession) GetListingFee() (*big.Int, error) {
	return _Kewl.Contract.GetListingFee(&_Kewl.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_Kewl *KewlCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_Kewl *KewlSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Kewl.Contract.GetPair(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_Kewl *KewlCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Kewl.Contract.GetPair(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairAddress is a free data retrieval call binding the contract method 0xb4f69a2f.
//
// Solidity: function getPairAddress(address input, address output) view returns(address)
func (_Kewl *KewlCaller) GetPairAddress(opts *bind.CallOpts, input common.Address, output common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPairAddress", input, output)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPairAddress is a free data retrieval call binding the contract method 0xb4f69a2f.
//
// Solidity: function getPairAddress(address input, address output) view returns(address)
func (_Kewl *KewlSession) GetPairAddress(input common.Address, output common.Address) (common.Address, error) {
	return _Kewl.Contract.GetPairAddress(&_Kewl.CallOpts, input, output)
}

// GetPairAddress is a free data retrieval call binding the contract method 0xb4f69a2f.
//
// Solidity: function getPairAddress(address input, address output) view returns(address)
func (_Kewl *KewlCallerSession) GetPairAddress(input common.Address, output common.Address) (common.Address, error) {
	return _Kewl.Contract.GetPairAddress(&_Kewl.CallOpts, input, output)
}

// GetPairId is a free data retrieval call binding the contract method 0x3f92a339.
//
// Solidity: function getPairId(address tokenA, address tokenB) pure returns(bytes32)
func (_Kewl *KewlCaller) GetPairId(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPairId", tokenA, tokenB)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPairId is a free data retrieval call binding the contract method 0x3f92a339.
//
// Solidity: function getPairId(address tokenA, address tokenB) pure returns(bytes32)
func (_Kewl *KewlSession) GetPairId(tokenA common.Address, tokenB common.Address) ([32]byte, error) {
	return _Kewl.Contract.GetPairId(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairId is a free data retrieval call binding the contract method 0x3f92a339.
//
// Solidity: function getPairId(address tokenA, address tokenB) pure returns(bytes32)
func (_Kewl *KewlCallerSession) GetPairId(tokenA common.Address, tokenB common.Address) ([32]byte, error) {
	return _Kewl.Contract.GetPairId(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlCaller) GetPairInfo(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPairInfo", tokenA, tokenB)

	if err != nil {
		return *new(TPair), err
	}

	out0 := *abi.ConvertType(out[0], new(TPair)).(*TPair)

	return out0, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlSession) GetPairInfo(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.GetPairInfo(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlCallerSession) GetPairInfo(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.GetPairInfo(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairInfo0 is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlCaller) GetPairInfo0(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPairInfo0", tokenA, tokenB)

	if err != nil {
		return *new(TPair), err
	}

	out0 := *abi.ConvertType(out[0], new(TPair)).(*TPair)

	return out0, err

}

// GetPairInfo0 is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlSession) GetPairInfo0(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.GetPairInfo0(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPairInfo0 is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)) pairInfo)
func (_Kewl *KewlCallerSession) GetPairInfo0(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.GetPairInfo0(&_Kewl.CallOpts, tokenA, tokenB)
}

// GetPriceOfAsset is a free data retrieval call binding the contract method 0x21d5403a.
//
// Solidity: function getPriceOfAsset(address user, uint256 amount, address input, address output) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address) pairInfo)
func (_Kewl *KewlCaller) GetPriceOfAsset(opts *bind.CallOpts, user common.Address, amount *big.Int, input common.Address, output common.Address) (BountyTBountyPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getPriceOfAsset", user, amount, input, output)

	if err != nil {
		return *new(BountyTBountyPair), err
	}

	out0 := *abi.ConvertType(out[0], new(BountyTBountyPair)).(*BountyTBountyPair)

	return out0, err

}

// GetPriceOfAsset is a free data retrieval call binding the contract method 0x21d5403a.
//
// Solidity: function getPriceOfAsset(address user, uint256 amount, address input, address output) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address) pairInfo)
func (_Kewl *KewlSession) GetPriceOfAsset(user common.Address, amount *big.Int, input common.Address, output common.Address) (BountyTBountyPair, error) {
	return _Kewl.Contract.GetPriceOfAsset(&_Kewl.CallOpts, user, amount, input, output)
}

// GetPriceOfAsset is a free data retrieval call binding the contract method 0x21d5403a.
//
// Solidity: function getPriceOfAsset(address user, uint256 amount, address input, address output) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address) pairInfo)
func (_Kewl *KewlCallerSession) GetPriceOfAsset(user common.Address, amount *big.Int, input common.Address, output common.Address) (BountyTBountyPair, error) {
	return _Kewl.Contract.GetPriceOfAsset(&_Kewl.CallOpts, user, amount, input, output)
}

// GetReservesByPairAddresses is a free data retrieval call binding the contract method 0xc12d83ca.
//
// Solidity: function getReservesByPairAddresses(address[] pairlist) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairs)
func (_Kewl *KewlCaller) GetReservesByPairAddresses(opts *bind.CallOpts, pairlist []common.Address) ([]TradePairInfo, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getReservesByPairAddresses", pairlist)

	if err != nil {
		return *new([]TradePairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TradePairInfo)).(*[]TradePairInfo)

	return out0, err

}

// GetReservesByPairAddresses is a free data retrieval call binding the contract method 0xc12d83ca.
//
// Solidity: function getReservesByPairAddresses(address[] pairlist) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairs)
func (_Kewl *KewlSession) GetReservesByPairAddresses(pairlist []common.Address) ([]TradePairInfo, error) {
	return _Kewl.Contract.GetReservesByPairAddresses(&_Kewl.CallOpts, pairlist)
}

// GetReservesByPairAddresses is a free data retrieval call binding the contract method 0xc12d83ca.
//
// Solidity: function getReservesByPairAddresses(address[] pairlist) view returns((bool,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address)[] pairs)
func (_Kewl *KewlCallerSession) GetReservesByPairAddresses(pairlist []common.Address) ([]TradePairInfo, error) {
	return _Kewl.Contract.GetReservesByPairAddresses(&_Kewl.CallOpts, pairlist)
}

// GetTradeStatsForMultipleTokens is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleTokens(opts *bind.CallOpts, tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleTokens", tokens, user)

	outstruct := new(struct {
		TotalTrades      []*big.Int
		IndividualTrades []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTrades = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.IndividualTrades = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTradeStatsForMultipleTokens is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlSession) GetTradeStatsForMultipleTokens(tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleTokens(&_Kewl.CallOpts, tokens, user)
}

// GetTradeStatsForMultipleTokens is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleTokens(tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleTokens(&_Kewl.CallOpts, tokens, user)
}

// GetTradeStatsForMultipleTokens0 is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleTokens0(opts *bind.CallOpts, tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleTokens0", tokens, user)

	outstruct := new(struct {
		TotalTrades      []*big.Int
		IndividualTrades []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTrades = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.IndividualTrades = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTradeStatsForMultipleTokens0 is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlSession) GetTradeStatsForMultipleTokens0(tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleTokens0(&_Kewl.CallOpts, tokens, user)
}

// GetTradeStatsForMultipleTokens0 is a free data retrieval call binding the contract method 0x3caa9eb6.
//
// Solidity: function getTradeStatsForMultipleTokens(address[] tokens, address user) view returns(uint256[] totalTrades, uint256[] individualTrades)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleTokens0(tokens []common.Address, user common.Address) (struct {
	TotalTrades      []*big.Int
	IndividualTrades []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleTokens0(&_Kewl.CallOpts, tokens, user)
}

// GetTradeStatsForMultipleUser is a free data retrieval call binding the contract method 0x05e6d07d.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, uint256 startOfDay, address[] users) view returns((uint256,uint256,uint256,uint256,uint256[],uint256[],uint256[],uint256[]) tradeStats)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleUser(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, startOfDay *big.Int, users []common.Address) (RouterTradeStats, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleUser", baseToken, quoteToken, startOfDay, users)

	if err != nil {
		return *new(RouterTradeStats), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterTradeStats)).(*RouterTradeStats)

	return out0, err

}

// GetTradeStatsForMultipleUser is a free data retrieval call binding the contract method 0x05e6d07d.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, uint256 startOfDay, address[] users) view returns((uint256,uint256,uint256,uint256,uint256[],uint256[],uint256[],uint256[]) tradeStats)
func (_Kewl *KewlSession) GetTradeStatsForMultipleUser(baseToken common.Address, quoteToken common.Address, startOfDay *big.Int, users []common.Address) (RouterTradeStats, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUser(&_Kewl.CallOpts, baseToken, quoteToken, startOfDay, users)
}

// GetTradeStatsForMultipleUser is a free data retrieval call binding the contract method 0x05e6d07d.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, uint256 startOfDay, address[] users) view returns((uint256,uint256,uint256,uint256,uint256[],uint256[],uint256[],uint256[]) tradeStats)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleUser(baseToken common.Address, quoteToken common.Address, startOfDay *big.Int, users []common.Address) (RouterTradeStats, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUser(&_Kewl.CallOpts, baseToken, quoteToken, startOfDay, users)
}

// GetTradeStatsForMultipleUser0 is a free data retrieval call binding the contract method 0x96a8a4c8.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, address[] traders, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleUser0(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	Traders         []common.Address
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleUser0", baseToken, quoteToken, users)

	outstruct := new(struct {
		TotalTradeBase  *big.Int
		TotalTradeQuote *big.Int
		Traders         []common.Address
		BaseVolume      []*big.Int
		QuoteVolume     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTradeBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTradeQuote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Traders = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.BaseVolume = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.QuoteVolume = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTradeStatsForMultipleUser0 is a free data retrieval call binding the contract method 0x96a8a4c8.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, address[] traders, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlSession) GetTradeStatsForMultipleUser0(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	Traders         []common.Address
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUser0(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForMultipleUser0 is a free data retrieval call binding the contract method 0x96a8a4c8.
//
// Solidity: function getTradeStatsForMultipleUser(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, address[] traders, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleUser0(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	Traders         []common.Address
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUser0(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForMultipleUserWithoutTraders is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleUserWithoutTraders(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleUserWithoutTraders", baseToken, quoteToken, users)

	outstruct := new(struct {
		TotalTradeBase  *big.Int
		TotalTradeQuote *big.Int
		BaseVolume      []*big.Int
		QuoteVolume     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTradeBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTradeQuote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BaseVolume = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.QuoteVolume = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTradeStatsForMultipleUserWithoutTraders is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlSession) GetTradeStatsForMultipleUserWithoutTraders(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUserWithoutTraders(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForMultipleUserWithoutTraders is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleUserWithoutTraders(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUserWithoutTraders(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForMultipleUserWithoutTraders0 is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCaller) GetTradeStatsForMultipleUserWithoutTraders0(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForMultipleUserWithoutTraders0", baseToken, quoteToken, users)

	outstruct := new(struct {
		TotalTradeBase  *big.Int
		TotalTradeQuote *big.Int
		BaseVolume      []*big.Int
		QuoteVolume     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTradeBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTradeQuote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BaseVolume = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.QuoteVolume = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTradeStatsForMultipleUserWithoutTraders0 is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlSession) GetTradeStatsForMultipleUserWithoutTraders0(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUserWithoutTraders0(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForMultipleUserWithoutTraders0 is a free data retrieval call binding the contract method 0x0b405ab2.
//
// Solidity: function getTradeStatsForMultipleUserWithoutTraders(address baseToken, address quoteToken, address[] users) view returns(uint256 totalTradeBase, uint256 totalTradeQuote, uint256[] baseVolume, uint256[] quoteVolume)
func (_Kewl *KewlCallerSession) GetTradeStatsForMultipleUserWithoutTraders0(baseToken common.Address, quoteToken common.Address, users []common.Address) (struct {
	TotalTradeBase  *big.Int
	TotalTradeQuote *big.Int
	BaseVolume      []*big.Int
	QuoteVolume     []*big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForMultipleUserWithoutTraders0(&_Kewl.CallOpts, baseToken, quoteToken, users)
}

// GetTradeStatsForToken is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlCaller) GetTradeStatsForToken(opts *bind.CallOpts, token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForToken", token, user)

	outstruct := new(struct {
		TotalTrades      *big.Int
		IndividualTrades *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTrades = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IndividualTrades = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTradeStatsForToken is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlSession) GetTradeStatsForToken(token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForToken(&_Kewl.CallOpts, token, user)
}

// GetTradeStatsForToken is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlCallerSession) GetTradeStatsForToken(token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForToken(&_Kewl.CallOpts, token, user)
}

// GetTradeStatsForToken0 is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlCaller) GetTradeStatsForToken0(opts *bind.CallOpts, token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getTradeStatsForToken0", token, user)

	outstruct := new(struct {
		TotalTrades      *big.Int
		IndividualTrades *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalTrades = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IndividualTrades = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTradeStatsForToken0 is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlSession) GetTradeStatsForToken0(token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForToken0(&_Kewl.CallOpts, token, user)
}

// GetTradeStatsForToken0 is a free data retrieval call binding the contract method 0x1fd23dd3.
//
// Solidity: function getTradeStatsForToken(address token, address user) view returns(uint256 totalTrades, uint256 individualTrades)
func (_Kewl *KewlCallerSession) GetTradeStatsForToken0(token common.Address, user common.Address) (struct {
	TotalTrades      *big.Int
	IndividualTrades *big.Int
}, error) {
	return _Kewl.Contract.GetTradeStatsForToken0(&_Kewl.CallOpts, token, user)
}

// GetUnderlyingToWrapped is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlying) view returns(address)
func (_Kewl *KewlCaller) GetUnderlyingToWrapped(opts *bind.CallOpts, underlying common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getUnderlyingToWrapped", underlying)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUnderlyingToWrapped is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlying) view returns(address)
func (_Kewl *KewlSession) GetUnderlyingToWrapped(underlying common.Address) (common.Address, error) {
	return _Kewl.Contract.GetUnderlyingToWrapped(&_Kewl.CallOpts, underlying)
}

// GetUnderlyingToWrapped is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlying) view returns(address)
func (_Kewl *KewlCallerSession) GetUnderlyingToWrapped(underlying common.Address) (common.Address, error) {
	return _Kewl.Contract.GetUnderlyingToWrapped(&_Kewl.CallOpts, underlying)
}

// GetUnderlyingToWrapped0 is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlyingToken) view returns(address)
func (_Kewl *KewlCaller) GetUnderlyingToWrapped0(opts *bind.CallOpts, underlyingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "getUnderlyingToWrapped0", underlyingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUnderlyingToWrapped0 is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlyingToken) view returns(address)
func (_Kewl *KewlSession) GetUnderlyingToWrapped0(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.GetUnderlyingToWrapped0(&_Kewl.CallOpts, underlyingToken)
}

// GetUnderlyingToWrapped0 is a free data retrieval call binding the contract method 0x1127c7ff.
//
// Solidity: function getUnderlyingToWrapped(address underlyingToken) view returns(address)
func (_Kewl *KewlCallerSession) GetUnderlyingToWrapped0(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.GetUnderlyingToWrapped0(&_Kewl.CallOpts, underlyingToken)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Kewl *KewlCaller) IsAuthorized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "isAuthorized", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Kewl *KewlSession) IsAuthorized(account common.Address) (bool, error) {
	return _Kewl.Contract.IsAuthorized(&_Kewl.CallOpts, account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address account) view returns(bool)
func (_Kewl *KewlCallerSession) IsAuthorized(account common.Address) (bool, error) {
	return _Kewl.Contract.IsAuthorized(&_Kewl.CallOpts, account)
}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlCaller) IsWrapped(opts *bind.CallOpts, token common.Address) (common.Address, bool, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "isWrapped", token)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlSession) IsWrapped(token common.Address) (common.Address, bool, error) {
	return _Kewl.Contract.IsWrapped(&_Kewl.CallOpts, token)
}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlCallerSession) IsWrapped(token common.Address) (common.Address, bool, error) {
	return _Kewl.Contract.IsWrapped(&_Kewl.CallOpts, token)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlCaller) IsWrapped0(opts *bind.CallOpts, token common.Address) (common.Address, bool, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "isWrapped0", token)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// IsWrapped0 is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlSession) IsWrapped0(token common.Address) (common.Address, bool, error) {
	return _Kewl.Contract.IsWrapped0(&_Kewl.CallOpts, token)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address token) view returns(address, bool)
func (_Kewl *KewlCallerSession) IsWrapped0(token common.Address) (common.Address, bool, error) {
	return _Kewl.Contract.IsWrapped0(&_Kewl.CallOpts, token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kewl *KewlCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kewl *KewlSession) Name() (string, error) {
	return _Kewl.Contract.Name(&_Kewl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kewl *KewlCallerSession) Name() (string, error) {
	return _Kewl.Contract.Name(&_Kewl.CallOpts)
}

// OrderBook is a free data retrieval call binding the contract method 0x0bb431b6.
//
// Solidity: function orderBook(bytes32 pairId, uint256 start, uint256 limit) view returns((uint256,uint256,uint256,uint256,uint256,bool)[] orderbook)
func (_Kewl *KewlCaller) OrderBook(opts *bind.CallOpts, pairId [32]byte, start *big.Int, limit *big.Int) ([]LibLimitOrdersPriceLevel, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "orderBook", pairId, start, limit)

	if err != nil {
		return *new([]LibLimitOrdersPriceLevel), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibLimitOrdersPriceLevel)).(*[]LibLimitOrdersPriceLevel)

	return out0, err

}

// OrderBook is a free data retrieval call binding the contract method 0x0bb431b6.
//
// Solidity: function orderBook(bytes32 pairId, uint256 start, uint256 limit) view returns((uint256,uint256,uint256,uint256,uint256,bool)[] orderbook)
func (_Kewl *KewlSession) OrderBook(pairId [32]byte, start *big.Int, limit *big.Int) ([]LibLimitOrdersPriceLevel, error) {
	return _Kewl.Contract.OrderBook(&_Kewl.CallOpts, pairId, start, limit)
}

// OrderBook is a free data retrieval call binding the contract method 0x0bb431b6.
//
// Solidity: function orderBook(bytes32 pairId, uint256 start, uint256 limit) view returns((uint256,uint256,uint256,uint256,uint256,bool)[] orderbook)
func (_Kewl *KewlCallerSession) OrderBook(pairId [32]byte, start *big.Int, limit *big.Int) ([]LibLimitOrdersPriceLevel, error) {
	return _Kewl.Contract.OrderBook(&_Kewl.CallOpts, pairId, start, limit)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kewl *KewlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kewl *KewlSession) Owner() (common.Address, error) {
	return _Kewl.Contract.Owner(&_Kewl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kewl *KewlCallerSession) Owner() (common.Address, error) {
	return _Kewl.Contract.Owner(&_Kewl.CallOpts)
}

// PairFor is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "pairFor", tokenA, tokenB)

	if err != nil {
		return *new(TPair), err
	}

	out0 := *abi.ConvertType(out[0], new(TPair)).(*TPair)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlSession) PairFor(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.PairFor(&_Kewl.CallOpts, tokenA, tokenB)
}

// PairFor is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlCallerSession) PairFor(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.PairFor(&_Kewl.CallOpts, tokenA, tokenB)
}

// PairFor0 is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlCaller) PairFor0(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (TPair, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "pairFor0", tokenA, tokenB)

	if err != nil {
		return *new(TPair), err
	}

	out0 := *abi.ConvertType(out[0], new(TPair)).(*TPair)

	return out0, err

}

// PairFor0 is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlSession) PairFor0(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.PairFor0(&_Kewl.CallOpts, tokenA, tokenB)
}

// PairFor0 is a free data retrieval call binding the contract method 0x96ed28f9.
//
// Solidity: function pairFor(address tokenA, address tokenB) view returns((bool,uint256,uint256,uint256,address,(uint256,address,string,string),(uint256,address,string,string)))
func (_Kewl *KewlCallerSession) PairFor0(tokenA common.Address, tokenB common.Address) (TPair, error) {
	return _Kewl.Contract.PairFor0(&_Kewl.CallOpts, tokenA, tokenB)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kewl *KewlCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kewl *KewlSession) Symbol() (string, error) {
	return _Kewl.Contract.Symbol(&_Kewl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kewl *KewlCallerSession) Symbol() (string, error) {
	return _Kewl.Contract.Symbol(&_Kewl.CallOpts)
}

// TestPair is a free data retrieval call binding the contract method 0x0dd4409e.
//
// Solidity: function testPair((address,address,bool,bool) router, address base, address quote) view returns(address)
func (_Kewl *KewlCaller) TestPair(opts *bind.CallOpts, router TradeRouter, base common.Address, quote common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "testPair", router, base, quote)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TestPair is a free data retrieval call binding the contract method 0x0dd4409e.
//
// Solidity: function testPair((address,address,bool,bool) router, address base, address quote) view returns(address)
func (_Kewl *KewlSession) TestPair(router TradeRouter, base common.Address, quote common.Address) (common.Address, error) {
	return _Kewl.Contract.TestPair(&_Kewl.CallOpts, router, base, quote)
}

// TestPair is a free data retrieval call binding the contract method 0x0dd4409e.
//
// Solidity: function testPair((address,address,bool,bool) router, address base, address quote) view returns(address)
func (_Kewl *KewlCallerSession) TestPair(router TradeRouter, base common.Address, quote common.Address) (common.Address, error) {
	return _Kewl.Contract.TestPair(&_Kewl.CallOpts, router, base, quote)
}

// UnderlyingToWrapped is a free data retrieval call binding the contract method 0xb4ec3372.
//
// Solidity: function underlyingToWrapped(address ) view returns(address)
func (_Kewl *KewlCaller) UnderlyingToWrapped(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "underlyingToWrapped", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToWrapped is a free data retrieval call binding the contract method 0xb4ec3372.
//
// Solidity: function underlyingToWrapped(address ) view returns(address)
func (_Kewl *KewlSession) UnderlyingToWrapped(arg0 common.Address) (common.Address, error) {
	return _Kewl.Contract.UnderlyingToWrapped(&_Kewl.CallOpts, arg0)
}

// UnderlyingToWrapped is a free data retrieval call binding the contract method 0xb4ec3372.
//
// Solidity: function underlyingToWrapped(address ) view returns(address)
func (_Kewl *KewlCallerSession) UnderlyingToWrapped(arg0 common.Address) (common.Address, error) {
	return _Kewl.Contract.UnderlyingToWrapped(&_Kewl.CallOpts, arg0)
}

// WrappedToUnderlying is a free data retrieval call binding the contract method 0x55579d32.
//
// Solidity: function wrappedToUnderlying(address ) view returns(address)
func (_Kewl *KewlCaller) WrappedToUnderlying(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "wrappedToUnderlying", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedToUnderlying is a free data retrieval call binding the contract method 0x55579d32.
//
// Solidity: function wrappedToUnderlying(address ) view returns(address)
func (_Kewl *KewlSession) WrappedToUnderlying(arg0 common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedToUnderlying(&_Kewl.CallOpts, arg0)
}

// WrappedToUnderlying is a free data retrieval call binding the contract method 0x55579d32.
//
// Solidity: function wrappedToUnderlying(address ) view returns(address)
func (_Kewl *KewlCallerSession) WrappedToUnderlying(arg0 common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedToUnderlying(&_Kewl.CallOpts, arg0)
}

// WrappedTokenFor is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address wrappedToken)
func (_Kewl *KewlCaller) WrappedTokenFor(opts *bind.CallOpts, underlyingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "wrappedTokenFor", underlyingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenFor is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address wrappedToken)
func (_Kewl *KewlSession) WrappedTokenFor(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedTokenFor(&_Kewl.CallOpts, underlyingToken)
}

// WrappedTokenFor is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address wrappedToken)
func (_Kewl *KewlCallerSession) WrappedTokenFor(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedTokenFor(&_Kewl.CallOpts, underlyingToken)
}

// WrappedTokenFor0 is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address)
func (_Kewl *KewlCaller) WrappedTokenFor0(opts *bind.CallOpts, underlyingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Kewl.contract.Call(opts, &out, "wrappedTokenFor0", underlyingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenFor0 is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address)
func (_Kewl *KewlSession) WrappedTokenFor0(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedTokenFor0(&_Kewl.CallOpts, underlyingToken)
}

// WrappedTokenFor0 is a free data retrieval call binding the contract method 0xc7bc7c1c.
//
// Solidity: function wrappedTokenFor(address underlyingToken) view returns(address)
func (_Kewl *KewlCallerSession) WrappedTokenFor0(underlyingToken common.Address) (common.Address, error) {
	return _Kewl.Contract.WrappedTokenFor0(&_Kewl.CallOpts, underlyingToken)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidity(&_Kewl.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidity(&_Kewl.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlTransactor) AddLiquidity0(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "addLiquidity0", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlSession) AddLiquidity0(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidity0(&_Kewl.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Kewl *KewlTransactorSession) AddLiquidity0(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidity0(&_Kewl.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidityETH(&_Kewl.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidityETH(&_Kewl.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH0 is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlTransactor) AddLiquidityETH0(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "addLiquidityETH0", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH0 is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlSession) AddLiquidityETH0(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidityETH0(&_Kewl.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH0 is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Kewl *KewlTransactorSession) AddLiquidityETH0(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.AddLiquidityETH0(&_Kewl.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlTransactor) Arbitrage(opts *bind.TransactOpts, params ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "arbitrage", params)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlSession) Arbitrage(params ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.Arbitrage(&_Kewl.TransactOpts, params)
}

// Arbitrage is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlTransactorSession) Arbitrage(params ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.Arbitrage(&_Kewl.TransactOpts, params)
}

// Arbitrage0 is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlTransactor) Arbitrage0(opts *bind.TransactOpts, params OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "arbitrage0", params)
}

// Arbitrage0 is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlSession) Arbitrage0(params OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.Arbitrage0(&_Kewl.TransactOpts, params)
}

// Arbitrage0 is a paid mutator transaction binding the contract method 0xb6536d3f.
//
// Solidity: function arbitrage((uint256,address,address,address,address) params) payable returns()
func (_Kewl *KewlTransactorSession) Arbitrage0(params OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.Arbitrage0(&_Kewl.TransactOpts, params)
}

// ArbitrageAll is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlTransactor) ArbitrageAll(opts *bind.TransactOpts, params []ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "arbitrageAll", params)
}

// ArbitrageAll is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlSession) ArbitrageAll(params []ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.ArbitrageAll(&_Kewl.TransactOpts, params)
}

// ArbitrageAll is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlTransactorSession) ArbitrageAll(params []ArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.ArbitrageAll(&_Kewl.TransactOpts, params)
}

// ArbitrageAll0 is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlTransactor) ArbitrageAll0(opts *bind.TransactOpts, params []OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "arbitrageAll0", params)
}

// ArbitrageAll0 is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlSession) ArbitrageAll0(params []OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.ArbitrageAll0(&_Kewl.TransactOpts, params)
}

// ArbitrageAll0 is a paid mutator transaction binding the contract method 0xf748c85b.
//
// Solidity: function arbitrageAll((uint256,address,address,address,address)[] params) payable returns()
func (_Kewl *KewlTransactorSession) ArbitrageAll0(params []OldArbitrageArbitrageParam) (*types.Transaction, error) {
	return _Kewl.Contract.ArbitrageAll0(&_Kewl.TransactOpts, params)
}

// BanAsset is a paid mutator transaction binding the contract method 0x22ac0fb4.
//
// Solidity: function banAsset(address _token, bool banned_) returns()
func (_Kewl *KewlTransactor) BanAsset(opts *bind.TransactOpts, _token common.Address, banned_ bool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "banAsset", _token, banned_)
}

// BanAsset is a paid mutator transaction binding the contract method 0x22ac0fb4.
//
// Solidity: function banAsset(address _token, bool banned_) returns()
func (_Kewl *KewlSession) BanAsset(_token common.Address, banned_ bool) (*types.Transaction, error) {
	return _Kewl.Contract.BanAsset(&_Kewl.TransactOpts, _token, banned_)
}

// BanAsset is a paid mutator transaction binding the contract method 0x22ac0fb4.
//
// Solidity: function banAsset(address _token, bool banned_) returns()
func (_Kewl *KewlTransactorSession) BanAsset(_token common.Address, banned_ bool) (*types.Transaction, error) {
	return _Kewl.Contract.BanAsset(&_Kewl.TransactOpts, _token, banned_)
}

// Cancel is a paid mutator transaction binding the contract method 0x069455a0.
//
// Solidity: function cancel(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlTransactor) Cancel(opts *bind.TransactOpts, pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "cancel", pairId, orderId)
}

// Cancel is a paid mutator transaction binding the contract method 0x069455a0.
//
// Solidity: function cancel(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlSession) Cancel(pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Cancel(&_Kewl.TransactOpts, pairId, orderId)
}

// Cancel is a paid mutator transaction binding the contract method 0x069455a0.
//
// Solidity: function cancel(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlTransactorSession) Cancel(pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Cancel(&_Kewl.TransactOpts, pairId, orderId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlTransactor) Claim(opts *bind.TransactOpts, pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "claim", pairId, orderId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlSession) Claim(pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim(&_Kewl.TransactOpts, pairId, orderId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 pairId, uint256 orderId) payable returns()
func (_Kewl *KewlTransactorSession) Claim(pairId [32]byte, orderId *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim(&_Kewl.TransactOpts, pairId, orderId)
}

// Claim0 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) returns(bool)
func (_Kewl *KewlTransactor) Claim0(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "claim0", token, receiver, amount)
}

// Claim0 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) returns(bool)
func (_Kewl *KewlSession) Claim0(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim0(&_Kewl.TransactOpts, token, receiver, amount)
}

// Claim0 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) returns(bool)
func (_Kewl *KewlTransactorSession) Claim0(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim0(&_Kewl.TransactOpts, token, receiver, amount)
}

// Claim1 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) payable returns(bool)
func (_Kewl *KewlTransactor) Claim1(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "claim1", token, receiver, amount)
}

// Claim1 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) payable returns(bool)
func (_Kewl *KewlSession) Claim1(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim1(&_Kewl.TransactOpts, token, receiver, amount)
}

// Claim1 is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address token, address receiver, uint256 amount) payable returns(bool)
func (_Kewl *KewlTransactorSession) Claim1(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Claim1(&_Kewl.TransactOpts, token, receiver, amount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xcf007095.
//
// Solidity: function claimReward((uint256,uint256,string) params) payable returns(bool)
func (_Kewl *KewlTransactor) ClaimReward(opts *bind.TransactOpts, params BountyClaimParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "claimReward", params)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xcf007095.
//
// Solidity: function claimReward((uint256,uint256,string) params) payable returns(bool)
func (_Kewl *KewlSession) ClaimReward(params BountyClaimParam) (*types.Transaction, error) {
	return _Kewl.Contract.ClaimReward(&_Kewl.TransactOpts, params)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xcf007095.
//
// Solidity: function claimReward((uint256,uint256,string) params) payable returns(bool)
func (_Kewl *KewlTransactorSession) ClaimReward(params BountyClaimParam) (*types.Transaction, error) {
	return _Kewl.Contract.ClaimReward(&_Kewl.TransactOpts, params)
}

// Create is a paid mutator transaction binding the contract method 0x247ce1c4.
//
// Solidity: function create((uint8,address,address,uint256,uint256,uint256,uint256[]) params) payable returns()
func (_Kewl *KewlTransactor) Create(opts *bind.TransactOpts, params LibLimitOrdersLimitOrderParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "create", params)
}

// Create is a paid mutator transaction binding the contract method 0x247ce1c4.
//
// Solidity: function create((uint8,address,address,uint256,uint256,uint256,uint256[]) params) payable returns()
func (_Kewl *KewlSession) Create(params LibLimitOrdersLimitOrderParam) (*types.Transaction, error) {
	return _Kewl.Contract.Create(&_Kewl.TransactOpts, params)
}

// Create is a paid mutator transaction binding the contract method 0x247ce1c4.
//
// Solidity: function create((uint8,address,address,uint256,uint256,uint256,uint256[]) params) payable returns()
func (_Kewl *KewlTransactorSession) Create(params LibLimitOrdersLimitOrderParam) (*types.Transaction, error) {
	return _Kewl.Contract.Create(&_Kewl.TransactOpts, params)
}

// CreatePaidPair is a paid mutator transaction binding the contract method 0x779a88a9.
//
// Solidity: function createPaidPair(address base, address quote) payable returns()
func (_Kewl *KewlTransactor) CreatePaidPair(opts *bind.TransactOpts, base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createPaidPair", base, quote)
}

// CreatePaidPair is a paid mutator transaction binding the contract method 0x779a88a9.
//
// Solidity: function createPaidPair(address base, address quote) payable returns()
func (_Kewl *KewlSession) CreatePaidPair(base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePaidPair(&_Kewl.TransactOpts, base, quote)
}

// CreatePaidPair is a paid mutator transaction binding the contract method 0x779a88a9.
//
// Solidity: function createPaidPair(address base, address quote) payable returns()
func (_Kewl *KewlTransactorSession) CreatePaidPair(base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePaidPair(&_Kewl.TransactOpts, base, quote)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePair(&_Kewl.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePair(&_Kewl.TransactOpts, tokenA, tokenB)
}

// CreatePair0 is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlTransactor) CreatePair0(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createPair0", tokenA, tokenB)
}

// CreatePair0 is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlSession) CreatePair0(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePair0(&_Kewl.TransactOpts, tokenA, tokenB)
}

// CreatePair0 is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns()
func (_Kewl *KewlTransactorSession) CreatePair0(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePair0(&_Kewl.TransactOpts, tokenA, tokenB)
}

// CreatePairOwner is a paid mutator transaction binding the contract method 0x09ecd276.
//
// Solidity: function createPairOwner(address base, address quote) returns()
func (_Kewl *KewlTransactor) CreatePairOwner(opts *bind.TransactOpts, base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createPairOwner", base, quote)
}

// CreatePairOwner is a paid mutator transaction binding the contract method 0x09ecd276.
//
// Solidity: function createPairOwner(address base, address quote) returns()
func (_Kewl *KewlSession) CreatePairOwner(base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePairOwner(&_Kewl.TransactOpts, base, quote)
}

// CreatePairOwner is a paid mutator transaction binding the contract method 0x09ecd276.
//
// Solidity: function createPairOwner(address base, address quote) returns()
func (_Kewl *KewlTransactorSession) CreatePairOwner(base common.Address, quote common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreatePairOwner(&_Kewl.TransactOpts, base, quote)
}

// CreateWrappedToken is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlTransactor) CreateWrappedToken(opts *bind.TransactOpts, underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createWrappedToken", underlyingToken)
}

// CreateWrappedToken is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlSession) CreateWrappedToken(underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreateWrappedToken(&_Kewl.TransactOpts, underlyingToken)
}

// CreateWrappedToken is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlTransactorSession) CreateWrappedToken(underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreateWrappedToken(&_Kewl.TransactOpts, underlyingToken)
}

// CreateWrappedToken0 is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlTransactor) CreateWrappedToken0(opts *bind.TransactOpts, underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "createWrappedToken0", underlyingToken)
}

// CreateWrappedToken0 is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlSession) CreateWrappedToken0(underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreateWrappedToken0(&_Kewl.TransactOpts, underlyingToken)
}

// CreateWrappedToken0 is a paid mutator transaction binding the contract method 0x4f9215c4.
//
// Solidity: function createWrappedToken(address underlyingToken) returns(address)
func (_Kewl *KewlTransactorSession) CreateWrappedToken0(underlyingToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.CreateWrappedToken0(&_Kewl.TransactOpts, underlyingToken)
}

// InitBounties is a paid mutator transaction binding the contract method 0xfe9a14e3.
//
// Solidity: function initBounties() returns()
func (_Kewl *KewlTransactor) InitBounties(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "initBounties")
}

// InitBounties is a paid mutator transaction binding the contract method 0xfe9a14e3.
//
// Solidity: function initBounties() returns()
func (_Kewl *KewlSession) InitBounties() (*types.Transaction, error) {
	return _Kewl.Contract.InitBounties(&_Kewl.TransactOpts)
}

// InitBounties is a paid mutator transaction binding the contract method 0xfe9a14e3.
//
// Solidity: function initBounties() returns()
func (_Kewl *KewlTransactorSession) InitBounties() (*types.Transaction, error) {
	return _Kewl.Contract.InitBounties(&_Kewl.TransactOpts)
}

// Jackpot is a paid mutator transaction binding the contract method 0x627fa511.
//
// Solidity: function jackpot(uint256 amount) payable returns()
func (_Kewl *KewlTransactor) Jackpot(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "jackpot", amount)
}

// Jackpot is a paid mutator transaction binding the contract method 0x627fa511.
//
// Solidity: function jackpot(uint256 amount) payable returns()
func (_Kewl *KewlSession) Jackpot(amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Jackpot(&_Kewl.TransactOpts, amount)
}

// Jackpot is a paid mutator transaction binding the contract method 0x627fa511.
//
// Solidity: function jackpot(uint256 amount) payable returns()
func (_Kewl *KewlTransactorSession) Jackpot(amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Jackpot(&_Kewl.TransactOpts, amount)
}

// Reflect is a paid mutator transaction binding the contract method 0x4cb4bed2.
//
// Solidity: function reflect((uint256,uint256,address,address,address) config, (bool,address,address)[] pools) payable returns()
func (_Kewl *KewlTransactor) Reflect(opts *bind.TransactOpts, config ReflectionSwapConfig, pools []ReflectionLiqudityPool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "reflect", config, pools)
}

// Reflect is a paid mutator transaction binding the contract method 0x4cb4bed2.
//
// Solidity: function reflect((uint256,uint256,address,address,address) config, (bool,address,address)[] pools) payable returns()
func (_Kewl *KewlSession) Reflect(config ReflectionSwapConfig, pools []ReflectionLiqudityPool) (*types.Transaction, error) {
	return _Kewl.Contract.Reflect(&_Kewl.TransactOpts, config, pools)
}

// Reflect is a paid mutator transaction binding the contract method 0x4cb4bed2.
//
// Solidity: function reflect((uint256,uint256,address,address,address) config, (bool,address,address)[] pools) payable returns()
func (_Kewl *KewlTransactorSession) Reflect(config ReflectionSwapConfig, pools []ReflectionLiqudityPool) (*types.Transaction, error) {
	return _Kewl.Contract.Reflect(&_Kewl.TransactOpts, config, pools)
}

// Register is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlTransactor) Register(opts *bind.TransactOpts, userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "register", userInfo)
}

// Register is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlSession) Register(userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.Contract.Register(&_Kewl.TransactOpts, userInfo)
}

// Register is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlTransactorSession) Register(userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.Contract.Register(&_Kewl.TransactOpts, userInfo)
}

// Register0 is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlTransactor) Register0(opts *bind.TransactOpts, userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "register0", userInfo)
}

// Register0 is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlSession) Register0(userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.Contract.Register0(&_Kewl.TransactOpts, userInfo)
}

// Register0 is a paid mutator transaction binding the contract method 0x442593f5.
//
// Solidity: function register((string,string,string,address) userInfo) payable returns()
func (_Kewl *KewlTransactorSession) Register0(userInfo LibLeaderBoardUserInfo) (*types.Transaction, error) {
	return _Kewl.Contract.Register0(&_Kewl.TransactOpts, userInfo)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidity(&_Kewl.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidity(&_Kewl.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlTransactor) RemoveLiquidity0(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "removeLiquidity0", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlSession) RemoveLiquidity0(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidity0(&_Kewl.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Kewl *KewlTransactorSession) RemoveLiquidity0(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidity0(&_Kewl.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidityETH(&_Kewl.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidityETH(&_Kewl.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH0 is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlTransactor) RemoveLiquidityETH0(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "removeLiquidityETH0", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH0 is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlSession) RemoveLiquidityETH0(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidityETH0(&_Kewl.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH0 is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Kewl *KewlTransactorSession) RemoveLiquidityETH0(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.RemoveLiquidityETH0(&_Kewl.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kewl *KewlTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kewl *KewlSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kewl.Contract.RenounceOwnership(&_Kewl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kewl *KewlTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kewl.Contract.RenounceOwnership(&_Kewl.TransactOpts)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address token) returns()
func (_Kewl *KewlTransactor) SetBaseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setBaseToken", token)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address token) returns()
func (_Kewl *KewlSession) SetBaseToken(token common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetBaseToken(&_Kewl.TransactOpts, token)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address token) returns()
func (_Kewl *KewlTransactorSession) SetBaseToken(token common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetBaseToken(&_Kewl.TransactOpts, token)
}

// SetBountyTokens is a paid mutator transaction binding the contract method 0x6e95605a.
//
// Solidity: function setBountyTokens(address bountyToken, address kwlToken, address wethToken) returns()
func (_Kewl *KewlTransactor) SetBountyTokens(opts *bind.TransactOpts, bountyToken common.Address, kwlToken common.Address, wethToken common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setBountyTokens", bountyToken, kwlToken, wethToken)
}

// SetBountyTokens is a paid mutator transaction binding the contract method 0x6e95605a.
//
// Solidity: function setBountyTokens(address bountyToken, address kwlToken, address wethToken) returns()
func (_Kewl *KewlSession) SetBountyTokens(bountyToken common.Address, kwlToken common.Address, wethToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetBountyTokens(&_Kewl.TransactOpts, bountyToken, kwlToken, wethToken)
}

// SetBountyTokens is a paid mutator transaction binding the contract method 0x6e95605a.
//
// Solidity: function setBountyTokens(address bountyToken, address kwlToken, address wethToken) returns()
func (_Kewl *KewlTransactorSession) SetBountyTokens(bountyToken common.Address, kwlToken common.Address, wethToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetBountyTokens(&_Kewl.TransactOpts, bountyToken, kwlToken, wethToken)
}

// SetDisableBountyStatus is a paid mutator transaction binding the contract method 0x4f0e73e1.
//
// Solidity: function setDisableBountyStatus(bool isEnabled) returns()
func (_Kewl *KewlTransactor) SetDisableBountyStatus(opts *bind.TransactOpts, isEnabled bool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setDisableBountyStatus", isEnabled)
}

// SetDisableBountyStatus is a paid mutator transaction binding the contract method 0x4f0e73e1.
//
// Solidity: function setDisableBountyStatus(bool isEnabled) returns()
func (_Kewl *KewlSession) SetDisableBountyStatus(isEnabled bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetDisableBountyStatus(&_Kewl.TransactOpts, isEnabled)
}

// SetDisableBountyStatus is a paid mutator transaction binding the contract method 0x4f0e73e1.
//
// Solidity: function setDisableBountyStatus(bool isEnabled) returns()
func (_Kewl *KewlTransactorSession) SetDisableBountyStatus(isEnabled bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetDisableBountyStatus(&_Kewl.TransactOpts, isEnabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool enabled_) returns()
func (_Kewl *KewlTransactor) SetEnabled(opts *bind.TransactOpts, enabled_ bool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setEnabled", enabled_)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool enabled_) returns()
func (_Kewl *KewlSession) SetEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetEnabled(&_Kewl.TransactOpts, enabled_)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool enabled_) returns()
func (_Kewl *KewlTransactorSession) SetEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetEnabled(&_Kewl.TransactOpts, enabled_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetFeeTo(&_Kewl.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetFeeTo(&_Kewl.TransactOpts, _feeTo)
}

// SetFeeTo0 is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlTransactor) SetFeeTo0(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setFeeTo0", _feeTo)
}

// SetFeeTo0 is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlSession) SetFeeTo0(_feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetFeeTo0(&_Kewl.TransactOpts, _feeTo)
}

// SetFeeTo0 is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Kewl *KewlTransactorSession) SetFeeTo0(_feeTo common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetFeeTo0(&_Kewl.TransactOpts, _feeTo)
}

// SetKEWL is a paid mutator transaction binding the contract method 0x49bf0ab1.
//
// Solidity: function setKEWL(address _address, bool _isKEWL) returns()
func (_Kewl *KewlTransactor) SetKEWL(opts *bind.TransactOpts, _address common.Address, _isKEWL bool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setKEWL", _address, _isKEWL)
}

// SetKEWL is a paid mutator transaction binding the contract method 0x49bf0ab1.
//
// Solidity: function setKEWL(address _address, bool _isKEWL) returns()
func (_Kewl *KewlSession) SetKEWL(_address common.Address, _isKEWL bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetKEWL(&_Kewl.TransactOpts, _address, _isKEWL)
}

// SetKEWL is a paid mutator transaction binding the contract method 0x49bf0ab1.
//
// Solidity: function setKEWL(address _address, bool _isKEWL) returns()
func (_Kewl *KewlTransactorSession) SetKEWL(_address common.Address, _isKEWL bool) (*types.Transaction, error) {
	return _Kewl.Contract.SetKEWL(&_Kewl.TransactOpts, _address, _isKEWL)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _fee) returns()
func (_Kewl *KewlTransactor) SetListingFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setListingFee", _fee)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _fee) returns()
func (_Kewl *KewlSession) SetListingFee(_fee *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SetListingFee(&_Kewl.TransactOpts, _fee)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _fee) returns()
func (_Kewl *KewlTransactorSession) SetListingFee(_fee *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SetListingFee(&_Kewl.TransactOpts, _fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Kewl *KewlTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Kewl *KewlSession) SetName(name_ string) (*types.Transaction, error) {
	return _Kewl.Contract.SetName(&_Kewl.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_Kewl *KewlTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _Kewl.Contract.SetName(&_Kewl.TransactOpts, name_)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router_) returns()
func (_Kewl *KewlTransactor) SetRouter(opts *bind.TransactOpts, router_ common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setRouter", router_)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router_) returns()
func (_Kewl *KewlSession) SetRouter(router_ common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetRouter(&_Kewl.TransactOpts, router_)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router_) returns()
func (_Kewl *KewlTransactorSession) SetRouter(router_ common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetRouter(&_Kewl.TransactOpts, router_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_Kewl *KewlTransactor) SetSymbol(opts *bind.TransactOpts, symbol_ string) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setSymbol", symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_Kewl *KewlSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _Kewl.Contract.SetSymbol(&_Kewl.TransactOpts, symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_Kewl *KewlTransactorSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _Kewl.Contract.SetSymbol(&_Kewl.TransactOpts, symbol_)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0xf6ad77a7.
//
// Solidity: function setVaultConfig(address _vault, address _weth) returns()
func (_Kewl *KewlTransactor) SetVaultConfig(opts *bind.TransactOpts, _vault common.Address, _weth common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setVaultConfig", _vault, _weth)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0xf6ad77a7.
//
// Solidity: function setVaultConfig(address _vault, address _weth) returns()
func (_Kewl *KewlSession) SetVaultConfig(_vault common.Address, _weth common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetVaultConfig(&_Kewl.TransactOpts, _vault, _weth)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0xf6ad77a7.
//
// Solidity: function setVaultConfig(address _vault, address _weth) returns()
func (_Kewl *KewlTransactorSession) SetVaultConfig(_vault common.Address, _weth common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetVaultConfig(&_Kewl.TransactOpts, _vault, _weth)
}

// SetWETH9Address is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlTransactor) SetWETH9Address(opts *bind.TransactOpts, _weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setWETH9Address", _weth9)
}

// SetWETH9Address is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlSession) SetWETH9Address(_weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWETH9Address(&_Kewl.TransactOpts, _weth9)
}

// SetWETH9Address is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlTransactorSession) SetWETH9Address(_weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWETH9Address(&_Kewl.TransactOpts, _weth9)
}

// SetWETH9Address0 is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlTransactor) SetWETH9Address0(opts *bind.TransactOpts, _weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setWETH9Address0", _weth9)
}

// SetWETH9Address0 is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlSession) SetWETH9Address0(_weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWETH9Address0(&_Kewl.TransactOpts, _weth9)
}

// SetWETH9Address0 is a paid mutator transaction binding the contract method 0x45f00bff.
//
// Solidity: function setWETH9Address(address _weth9) returns()
func (_Kewl *KewlTransactorSession) SetWETH9Address0(_weth9 common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWETH9Address0(&_Kewl.TransactOpts, _weth9)
}

// SetWrappedAssets is a paid mutator transaction binding the contract method 0x82a8b976.
//
// Solidity: function setWrappedAssets(address[] weth9) returns()
func (_Kewl *KewlTransactor) SetWrappedAssets(opts *bind.TransactOpts, weth9 []common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "setWrappedAssets", weth9)
}

// SetWrappedAssets is a paid mutator transaction binding the contract method 0x82a8b976.
//
// Solidity: function setWrappedAssets(address[] weth9) returns()
func (_Kewl *KewlSession) SetWrappedAssets(weth9 []common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWrappedAssets(&_Kewl.TransactOpts, weth9)
}

// SetWrappedAssets is a paid mutator transaction binding the contract method 0x82a8b976.
//
// Solidity: function setWrappedAssets(address[] weth9) returns()
func (_Kewl *KewlTransactorSession) SetWrappedAssets(weth9 []common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.SetWrappedAssets(&_Kewl.TransactOpts, weth9)
}

// Swap is a paid mutator transaction binding the contract method 0xddaa9bf0.
//
// Solidity: function swap((uint256,uint256,address,address,address,address,bool) params) payable returns()
func (_Kewl *KewlTransactor) Swap(opts *bind.TransactOpts, params TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swap", params)
}

// Swap is a paid mutator transaction binding the contract method 0xddaa9bf0.
//
// Solidity: function swap((uint256,uint256,address,address,address,address,bool) params) payable returns()
func (_Kewl *KewlSession) Swap(params TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.Contract.Swap(&_Kewl.TransactOpts, params)
}

// Swap is a paid mutator transaction binding the contract method 0xddaa9bf0.
//
// Solidity: function swap((uint256,uint256,address,address,address,address,bool) params) payable returns()
func (_Kewl *KewlTransactorSession) Swap(params TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.Contract.Swap(&_Kewl.TransactOpts, params)
}

// SwapAll is a paid mutator transaction binding the contract method 0xea9483f7.
//
// Solidity: function swapAll((uint256,uint256,address,address,address,address,bool)[] params) payable returns()
func (_Kewl *KewlTransactor) SwapAll(opts *bind.TransactOpts, params []TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapAll", params)
}

// SwapAll is a paid mutator transaction binding the contract method 0xea9483f7.
//
// Solidity: function swapAll((uint256,uint256,address,address,address,address,bool)[] params) payable returns()
func (_Kewl *KewlSession) SwapAll(params []TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.Contract.SwapAll(&_Kewl.TransactOpts, params)
}

// SwapAll is a paid mutator transaction binding the contract method 0xea9483f7.
//
// Solidity: function swapAll((uint256,uint256,address,address,address,address,bool)[] params) payable returns()
func (_Kewl *KewlTransactorSession) SwapAll(params []TradeSwapParam) (*types.Transaction, error) {
	return _Kewl.Contract.SwapAll(&_Kewl.TransactOpts, params)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapETHForExactTokens", amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapETHForExactTokens(&_Kewl.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapETHForExactTokens(&_Kewl.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens0 is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapETHForExactTokens0(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapETHForExactTokens0", amountOut, path, to, deadline)
}

// SwapETHForExactTokens0 is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapETHForExactTokens0(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapETHForExactTokens0(&_Kewl.TransactOpts, amountOut, path, to, deadline)
}

// SwapETHForExactTokens0 is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapETHForExactTokens0(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapETHForExactTokens0(&_Kewl.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokens(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokens(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens0 is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactETHForTokens0(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactETHForTokens0", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens0 is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactETHForTokens0(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokens0(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens0 is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactETHForTokens0(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokens0(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens0(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens0", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlSession) SwapExactETHForTokensSupportingFeeOnTransferTokens0(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Kewl *KewlTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens0(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETH(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETH(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH0 is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactTokensForETH0(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForETH0", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH0 is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactTokensForETH0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETH0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH0 is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactTokensForETH0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETH0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens0(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens0", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlSession) SwapExactTokensForETHSupportingFeeOnTransferTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens0 is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapExactTokensForTokens0(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForTokens0", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens0 is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapExactTokensForTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens0 is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapExactTokensForTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens0(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens0", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens0 is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Kewl *KewlTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens0(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens0(&_Kewl.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactETH(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactETH(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH0 is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapTokensForExactETH0(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapTokensForExactETH0", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH0 is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapTokensForExactETH0(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactETH0(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH0 is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapTokensForExactETH0(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactETH0(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactTokens(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactTokens(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens0 is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactor) SwapTokensForExactTokens0(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "swapTokensForExactTokens0", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens0 is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlSession) SwapTokensForExactTokens0(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactTokens0(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens0 is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Kewl *KewlTransactorSession) SwapTokensForExactTokens0(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.SwapTokensForExactTokens0(&_Kewl.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SyncBountyPairs is a paid mutator transaction binding the contract method 0x8ca990db.
//
// Solidity: function syncBountyPairs() returns()
func (_Kewl *KewlTransactor) SyncBountyPairs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "syncBountyPairs")
}

// SyncBountyPairs is a paid mutator transaction binding the contract method 0x8ca990db.
//
// Solidity: function syncBountyPairs() returns()
func (_Kewl *KewlSession) SyncBountyPairs() (*types.Transaction, error) {
	return _Kewl.Contract.SyncBountyPairs(&_Kewl.TransactOpts)
}

// SyncBountyPairs is a paid mutator transaction binding the contract method 0x8ca990db.
//
// Solidity: function syncBountyPairs() returns()
func (_Kewl *KewlTransactorSession) SyncBountyPairs() (*types.Transaction, error) {
	return _Kewl.Contract.SyncBountyPairs(&_Kewl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kewl *KewlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kewl *KewlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.TransferOwnership(&_Kewl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kewl *KewlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.TransferOwnership(&_Kewl.TransactOpts, newOwner)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlTransactor) Unwrap(opts *bind.TransactOpts, account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "unwrap", account, wrappedToken, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlSession) Unwrap(account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Unwrap(&_Kewl.TransactOpts, account, wrappedToken, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlTransactorSession) Unwrap(account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Unwrap(&_Kewl.TransactOpts, account, wrappedToken, amount)
}

// Unwrap0 is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlTransactor) Unwrap0(opts *bind.TransactOpts, account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "unwrap0", account, wrappedToken, amount)
}

// Unwrap0 is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlSession) Unwrap0(account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Unwrap0(&_Kewl.TransactOpts, account, wrappedToken, amount)
}

// Unwrap0 is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address account, address wrappedToken, uint256 amount) returns()
func (_Kewl *KewlTransactorSession) Unwrap0(account common.Address, wrappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Unwrap0(&_Kewl.TransactOpts, account, wrappedToken, amount)
}

// UpdateBounty is a paid mutator transaction binding the contract method 0x3e9db954.
//
// Solidity: function updateBounty(uint256 boundtyId, bool valid, uint256 rewardAmount, address bountyToken) returns()
func (_Kewl *KewlTransactor) UpdateBounty(opts *bind.TransactOpts, boundtyId *big.Int, valid bool, rewardAmount *big.Int, bountyToken common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "updateBounty", boundtyId, valid, rewardAmount, bountyToken)
}

// UpdateBounty is a paid mutator transaction binding the contract method 0x3e9db954.
//
// Solidity: function updateBounty(uint256 boundtyId, bool valid, uint256 rewardAmount, address bountyToken) returns()
func (_Kewl *KewlSession) UpdateBounty(boundtyId *big.Int, valid bool, rewardAmount *big.Int, bountyToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.UpdateBounty(&_Kewl.TransactOpts, boundtyId, valid, rewardAmount, bountyToken)
}

// UpdateBounty is a paid mutator transaction binding the contract method 0x3e9db954.
//
// Solidity: function updateBounty(uint256 boundtyId, bool valid, uint256 rewardAmount, address bountyToken) returns()
func (_Kewl *KewlTransactorSession) UpdateBounty(boundtyId *big.Int, valid bool, rewardAmount *big.Int, bountyToken common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.UpdateBounty(&_Kewl.TransactOpts, boundtyId, valid, rewardAmount, bountyToken)
}

// Vote is a paid mutator transaction binding the contract method 0x4b9f5c98.
//
// Solidity: function vote(bool upDown) payable returns()
func (_Kewl *KewlTransactor) Vote(opts *bind.TransactOpts, upDown bool) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "vote", upDown)
}

// Vote is a paid mutator transaction binding the contract method 0x4b9f5c98.
//
// Solidity: function vote(bool upDown) payable returns()
func (_Kewl *KewlSession) Vote(upDown bool) (*types.Transaction, error) {
	return _Kewl.Contract.Vote(&_Kewl.TransactOpts, upDown)
}

// Vote is a paid mutator transaction binding the contract method 0x4b9f5c98.
//
// Solidity: function vote(bool upDown) payable returns()
func (_Kewl *KewlTransactorSession) Vote(upDown bool) (*types.Transaction, error) {
	return _Kewl.Contract.Vote(&_Kewl.TransactOpts, upDown)
}

// WithdrawCoins is a paid mutator transaction binding the contract method 0xf9f95a0f.
//
// Solidity: function withdrawCoins(uint256 amount) returns(bool)
func (_Kewl *KewlTransactor) WithdrawCoins(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawCoins", amount)
}

// WithdrawCoins is a paid mutator transaction binding the contract method 0xf9f95a0f.
//
// Solidity: function withdrawCoins(uint256 amount) returns(bool)
func (_Kewl *KewlSession) WithdrawCoins(amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCoins(&_Kewl.TransactOpts, amount)
}

// WithdrawCoins is a paid mutator transaction binding the contract method 0xf9f95a0f.
//
// Solidity: function withdrawCoins(uint256 amount) returns(bool)
func (_Kewl *KewlTransactorSession) WithdrawCoins(amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCoins(&_Kewl.TransactOpts, amount)
}

// WithdrawCustomERC is a paid mutator transaction binding the contract method 0x87907b8e.
//
// Solidity: function withdrawCustomERC(address _tokenAddr, uint256 _amount, address _to) returns()
func (_Kewl *KewlTransactor) WithdrawCustomERC(opts *bind.TransactOpts, _tokenAddr common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawCustomERC", _tokenAddr, _amount, _to)
}

// WithdrawCustomERC is a paid mutator transaction binding the contract method 0x87907b8e.
//
// Solidity: function withdrawCustomERC(address _tokenAddr, uint256 _amount, address _to) returns()
func (_Kewl *KewlSession) WithdrawCustomERC(_tokenAddr common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCustomERC(&_Kewl.TransactOpts, _tokenAddr, _amount, _to)
}

// WithdrawCustomERC is a paid mutator transaction binding the contract method 0x87907b8e.
//
// Solidity: function withdrawCustomERC(address _tokenAddr, uint256 _amount, address _to) returns()
func (_Kewl *KewlTransactorSession) WithdrawCustomERC(_tokenAddr common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCustomERC(&_Kewl.TransactOpts, _tokenAddr, _amount, _to)
}

// WithdrawCustomETHAmount is a paid mutator transaction binding the contract method 0x89c13b69.
//
// Solidity: function withdrawCustomETHAmount(uint256 amount, address receiver) returns()
func (_Kewl *KewlTransactor) WithdrawCustomETHAmount(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawCustomETHAmount", amount, receiver)
}

// WithdrawCustomETHAmount is a paid mutator transaction binding the contract method 0x89c13b69.
//
// Solidity: function withdrawCustomETHAmount(uint256 amount, address receiver) returns()
func (_Kewl *KewlSession) WithdrawCustomETHAmount(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCustomETHAmount(&_Kewl.TransactOpts, amount, receiver)
}

// WithdrawCustomETHAmount is a paid mutator transaction binding the contract method 0x89c13b69.
//
// Solidity: function withdrawCustomETHAmount(uint256 amount, address receiver) returns()
func (_Kewl *KewlTransactorSession) WithdrawCustomETHAmount(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawCustomETHAmount(&_Kewl.TransactOpts, amount, receiver)
}

// WithdrawERC is a paid mutator transaction binding the contract method 0xe95164f5.
//
// Solidity: function withdrawERC(address _tokenAddr) returns()
func (_Kewl *KewlTransactor) WithdrawERC(opts *bind.TransactOpts, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawERC", _tokenAddr)
}

// WithdrawERC is a paid mutator transaction binding the contract method 0xe95164f5.
//
// Solidity: function withdrawERC(address _tokenAddr) returns()
func (_Kewl *KewlSession) WithdrawERC(_tokenAddr common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawERC(&_Kewl.TransactOpts, _tokenAddr)
}

// WithdrawERC is a paid mutator transaction binding the contract method 0xe95164f5.
//
// Solidity: function withdrawERC(address _tokenAddr) returns()
func (_Kewl *KewlTransactorSession) WithdrawERC(_tokenAddr common.Address) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawERC(&_Kewl.TransactOpts, _tokenAddr)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Kewl *KewlTransactor) WithdrawETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawETH")
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Kewl *KewlSession) WithdrawETH() (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawETH(&_Kewl.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Kewl *KewlTransactorSession) WithdrawETH() (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawETH(&_Kewl.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address _tokenAddress, uint256 amount) returns(bool)
func (_Kewl *KewlTransactor) WithdrawTokens(opts *bind.TransactOpts, _tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "withdrawTokens", _tokenAddress, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address _tokenAddress, uint256 amount) returns(bool)
func (_Kewl *KewlSession) WithdrawTokens(_tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawTokens(&_Kewl.TransactOpts, _tokenAddress, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address _tokenAddress, uint256 amount) returns(bool)
func (_Kewl *KewlTransactorSession) WithdrawTokens(_tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.WithdrawTokens(&_Kewl.TransactOpts, _tokenAddress, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlTransactor) Wrap(opts *bind.TransactOpts, account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "wrap", account, underlyingToken, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlSession) Wrap(account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Wrap(&_Kewl.TransactOpts, account, underlyingToken, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlTransactorSession) Wrap(account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Wrap(&_Kewl.TransactOpts, account, underlyingToken, amount)
}

// Wrap0 is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlTransactor) Wrap0(opts *bind.TransactOpts, account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.contract.Transact(opts, "wrap0", account, underlyingToken, amount)
}

// Wrap0 is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlSession) Wrap0(account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Wrap0(&_Kewl.TransactOpts, account, underlyingToken, amount)
}

// Wrap0 is a paid mutator transaction binding the contract method 0x62355638.
//
// Solidity: function wrap(address account, address underlyingToken, uint256 amount) returns(address wrappedToken)
func (_Kewl *KewlTransactorSession) Wrap0(account common.Address, underlyingToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kewl.Contract.Wrap0(&_Kewl.TransactOpts, account, underlyingToken, amount)
}

// KewlLimitOrderCreatedIterator is returned from FilterLimitOrderCreated and is used to iterate over the raw logs and unpacked data for LimitOrderCreated events raised by the Kewl contract.
type KewlLimitOrderCreatedIterator struct {
	Event *KewlLimitOrderCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlLimitOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlLimitOrderCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlLimitOrderCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlLimitOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlLimitOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlLimitOrderCreated represents a LimitOrderCreated event raised by the Kewl contract.
type KewlLimitOrderCreated struct {
	PairId    [32]byte
	OrderId   *big.Int
	Trader    common.Address
	Kind      uint8
	Price     *big.Int
	Amount    *big.Int
	Total     *big.Int
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLimitOrderCreated is a free log retrieval operation binding the contract event 0x883e4c8b21425cd4f47c50ddc9d73d810d53d27345b4acb1293f53ebf11701b9.
//
// Solidity: event LimitOrderCreated(bytes32 indexed pairId, uint256 indexed orderId, address indexed trader, uint8 kind, uint256 price, uint256 amount, uint256 total, uint256 createdAt)
func (_Kewl *KewlFilterer) FilterLimitOrderCreated(opts *bind.FilterOpts, pairId [][32]byte, orderId []*big.Int, trader []common.Address) (*KewlLimitOrderCreatedIterator, error) {

	var pairIdRule []interface{}
	for _, pairIdItem := range pairId {
		pairIdRule = append(pairIdRule, pairIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "LimitOrderCreated", pairIdRule, orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &KewlLimitOrderCreatedIterator{contract: _Kewl.contract, event: "LimitOrderCreated", logs: logs, sub: sub}, nil
}

// WatchLimitOrderCreated is a free log subscription operation binding the contract event 0x883e4c8b21425cd4f47c50ddc9d73d810d53d27345b4acb1293f53ebf11701b9.
//
// Solidity: event LimitOrderCreated(bytes32 indexed pairId, uint256 indexed orderId, address indexed trader, uint8 kind, uint256 price, uint256 amount, uint256 total, uint256 createdAt)
func (_Kewl *KewlFilterer) WatchLimitOrderCreated(opts *bind.WatchOpts, sink chan<- *KewlLimitOrderCreated, pairId [][32]byte, orderId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var pairIdRule []interface{}
	for _, pairIdItem := range pairId {
		pairIdRule = append(pairIdRule, pairIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "LimitOrderCreated", pairIdRule, orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlLimitOrderCreated)
				if err := _Kewl.contract.UnpackLog(event, "LimitOrderCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLimitOrderCreated is a log parse operation binding the contract event 0x883e4c8b21425cd4f47c50ddc9d73d810d53d27345b4acb1293f53ebf11701b9.
//
// Solidity: event LimitOrderCreated(bytes32 indexed pairId, uint256 indexed orderId, address indexed trader, uint8 kind, uint256 price, uint256 amount, uint256 total, uint256 createdAt)
func (_Kewl *KewlFilterer) ParseLimitOrderCreated(log types.Log) (*KewlLimitOrderCreated, error) {
	event := new(KewlLimitOrderCreated)
	if err := _Kewl.contract.UnpackLog(event, "LimitOrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlMappingChangedIterator is returned from FilterMappingChanged and is used to iterate over the raw logs and unpacked data for MappingChanged events raised by the Kewl contract.
type KewlMappingChangedIterator struct {
	Event *KewlMappingChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlMappingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlMappingChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlMappingChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlMappingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlMappingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlMappingChanged represents a MappingChanged event raised by the Kewl contract.
type KewlMappingChanged struct {
	UnderlyingToken    common.Address
	WrappedToken       common.Address
	NewUnderlyingToken common.Address
	NewWrappedToken    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMappingChanged is a free log retrieval operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) FilterMappingChanged(opts *bind.FilterOpts) (*KewlMappingChangedIterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "MappingChanged")
	if err != nil {
		return nil, err
	}
	return &KewlMappingChangedIterator{contract: _Kewl.contract, event: "MappingChanged", logs: logs, sub: sub}, nil
}

// WatchMappingChanged is a free log subscription operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) WatchMappingChanged(opts *bind.WatchOpts, sink chan<- *KewlMappingChanged) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "MappingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlMappingChanged)
				if err := _Kewl.contract.UnpackLog(event, "MappingChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMappingChanged is a log parse operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) ParseMappingChanged(log types.Log) (*KewlMappingChanged, error) {
	event := new(KewlMappingChanged)
	if err := _Kewl.contract.UnpackLog(event, "MappingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlMappingChanged0Iterator is returned from FilterMappingChanged0 and is used to iterate over the raw logs and unpacked data for MappingChanged0 events raised by the Kewl contract.
type KewlMappingChanged0Iterator struct {
	Event *KewlMappingChanged0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlMappingChanged0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlMappingChanged0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlMappingChanged0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlMappingChanged0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlMappingChanged0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlMappingChanged0 represents a MappingChanged0 event raised by the Kewl contract.
type KewlMappingChanged0 struct {
	UnderlyingToken    common.Address
	WrappedToken       common.Address
	NewUnderlyingToken common.Address
	NewWrappedToken    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMappingChanged0 is a free log retrieval operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) FilterMappingChanged0(opts *bind.FilterOpts) (*KewlMappingChanged0Iterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "MappingChanged0")
	if err != nil {
		return nil, err
	}
	return &KewlMappingChanged0Iterator{contract: _Kewl.contract, event: "MappingChanged0", logs: logs, sub: sub}, nil
}

// WatchMappingChanged0 is a free log subscription operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) WatchMappingChanged0(opts *bind.WatchOpts, sink chan<- *KewlMappingChanged0) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "MappingChanged0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlMappingChanged0)
				if err := _Kewl.contract.UnpackLog(event, "MappingChanged0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMappingChanged0 is a log parse operation binding the contract event 0x78039ef781848572230b0b6190bd823265154a51ed620591886a3b1f57c30809.
//
// Solidity: event MappingChanged(address underlyingToken, address wrappedToken, address newUnderlyingToken, address newWrappedToken)
func (_Kewl *KewlFilterer) ParseMappingChanged0(log types.Log) (*KewlMappingChanged0, error) {
	event := new(KewlMappingChanged0)
	if err := _Kewl.contract.UnpackLog(event, "MappingChanged0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Kewl contract.
type KewlOrderCancelledIterator struct {
	Event *KewlOrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOrderCancelled represents a OrderCancelled event raised by the Kewl contract.
type KewlOrderCancelled struct {
	OrderId *big.Int
	Trader  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xc0362da6f2ff36b382b34aec0814f6b3cdf89f5ef282a1d1f114d0c0b036d596.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed trader)
func (_Kewl *KewlFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId []*big.Int, trader []common.Address) (*KewlOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OrderCancelled", orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &KewlOrderCancelledIterator{contract: _Kewl.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xc0362da6f2ff36b382b34aec0814f6b3cdf89f5ef282a1d1f114d0c0b036d596.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed trader)
func (_Kewl *KewlFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *KewlOrderCancelled, orderId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OrderCancelled", orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOrderCancelled)
				if err := _Kewl.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0xc0362da6f2ff36b382b34aec0814f6b3cdf89f5ef282a1d1f114d0c0b036d596.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, address indexed trader)
func (_Kewl *KewlFilterer) ParseOrderCancelled(log types.Log) (*KewlOrderCancelled, error) {
	event := new(KewlOrderCancelled)
	if err := _Kewl.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the Kewl contract.
type KewlOrderMatchedIterator struct {
	Event *KewlOrderMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOrderMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOrderMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOrderMatched represents a OrderMatched event raised by the Kewl contract.
type KewlOrderMatched struct {
	PairId    [32]byte
	Kind      bool
	Price     *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0x27d34d69da38551391ec9cc4c62037d5821893aa0b6187dd6b7c6b16c025a256.
//
// Solidity: event OrderMatched(bytes32 indexed pairId, bool kind, uint256 price, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) FilterOrderMatched(opts *bind.FilterOpts, pairId [][32]byte) (*KewlOrderMatchedIterator, error) {

	var pairIdRule []interface{}
	for _, pairIdItem := range pairId {
		pairIdRule = append(pairIdRule, pairIdItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OrderMatched", pairIdRule)
	if err != nil {
		return nil, err
	}
	return &KewlOrderMatchedIterator{contract: _Kewl.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0x27d34d69da38551391ec9cc4c62037d5821893aa0b6187dd6b7c6b16c025a256.
//
// Solidity: event OrderMatched(bytes32 indexed pairId, bool kind, uint256 price, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *KewlOrderMatched, pairId [][32]byte) (event.Subscription, error) {

	var pairIdRule []interface{}
	for _, pairIdItem := range pairId {
		pairIdRule = append(pairIdRule, pairIdItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OrderMatched", pairIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOrderMatched)
				if err := _Kewl.contract.UnpackLog(event, "OrderMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderMatched is a log parse operation binding the contract event 0x27d34d69da38551391ec9cc4c62037d5821893aa0b6187dd6b7c6b16c025a256.
//
// Solidity: event OrderMatched(bytes32 indexed pairId, bool kind, uint256 price, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) ParseOrderMatched(log types.Log) (*KewlOrderMatched, error) {
	event := new(KewlOrderMatched)
	if err := _Kewl.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kewl contract.
type KewlOwnershipTransferredIterator struct {
	Event *KewlOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred represents a OwnershipTransferred event raised by the Kewl contract.
type KewlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferredIterator{contract: _Kewl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred(log types.Log) (*KewlOwnershipTransferred, error) {
	event := new(KewlOwnershipTransferred)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred0Iterator is returned from FilterOwnershipTransferred0 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred0 events raised by the Kewl contract.
type KewlOwnershipTransferred0Iterator struct {
	Event *KewlOwnershipTransferred0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred0 represents a OwnershipTransferred0 event raised by the Kewl contract.
type KewlOwnershipTransferred0 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred0 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred0(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred0Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred0Iterator{contract: _Kewl.contract, event: "OwnershipTransferred0", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred0 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred0(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred0, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred0)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred0 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred0(log types.Log) (*KewlOwnershipTransferred0, error) {
	event := new(KewlOwnershipTransferred0)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred1Iterator is returned from FilterOwnershipTransferred1 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred1 events raised by the Kewl contract.
type KewlOwnershipTransferred1Iterator struct {
	Event *KewlOwnershipTransferred1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred1 represents a OwnershipTransferred1 event raised by the Kewl contract.
type KewlOwnershipTransferred1 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred1 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred1(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred1Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred1", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred1Iterator{contract: _Kewl.contract, event: "OwnershipTransferred1", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred1 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred1(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred1, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred1", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred1)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred1 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred1(log types.Log) (*KewlOwnershipTransferred1, error) {
	event := new(KewlOwnershipTransferred1)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred10Iterator is returned from FilterOwnershipTransferred10 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred10 events raised by the Kewl contract.
type KewlOwnershipTransferred10Iterator struct {
	Event *KewlOwnershipTransferred10 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred10Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred10)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred10)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred10Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred10Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred10 represents a OwnershipTransferred10 event raised by the Kewl contract.
type KewlOwnershipTransferred10 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred10 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred10(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred10Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred10", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred10Iterator{contract: _Kewl.contract, event: "OwnershipTransferred10", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred10 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred10(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred10, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred10", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred10)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred10", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred10 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred10(log types.Log) (*KewlOwnershipTransferred10, error) {
	event := new(KewlOwnershipTransferred10)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred10", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred11Iterator is returned from FilterOwnershipTransferred11 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred11 events raised by the Kewl contract.
type KewlOwnershipTransferred11Iterator struct {
	Event *KewlOwnershipTransferred11 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred11Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred11)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred11)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred11Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred11Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred11 represents a OwnershipTransferred11 event raised by the Kewl contract.
type KewlOwnershipTransferred11 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred11 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred11(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred11Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred11", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred11Iterator{contract: _Kewl.contract, event: "OwnershipTransferred11", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred11 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred11(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred11, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred11", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred11)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred11", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred11 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred11(log types.Log) (*KewlOwnershipTransferred11, error) {
	event := new(KewlOwnershipTransferred11)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred11", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred2Iterator is returned from FilterOwnershipTransferred2 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred2 events raised by the Kewl contract.
type KewlOwnershipTransferred2Iterator struct {
	Event *KewlOwnershipTransferred2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred2 represents a OwnershipTransferred2 event raised by the Kewl contract.
type KewlOwnershipTransferred2 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred2 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred2(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred2Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred2", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred2Iterator{contract: _Kewl.contract, event: "OwnershipTransferred2", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred2 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred2(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred2, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred2", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred2)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred2 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred2(log types.Log) (*KewlOwnershipTransferred2, error) {
	event := new(KewlOwnershipTransferred2)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred3Iterator is returned from FilterOwnershipTransferred3 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred3 events raised by the Kewl contract.
type KewlOwnershipTransferred3Iterator struct {
	Event *KewlOwnershipTransferred3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred3 represents a OwnershipTransferred3 event raised by the Kewl contract.
type KewlOwnershipTransferred3 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred3 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred3(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred3Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred3", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred3Iterator{contract: _Kewl.contract, event: "OwnershipTransferred3", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred3 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred3(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred3, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred3", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred3)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred3 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred3(log types.Log) (*KewlOwnershipTransferred3, error) {
	event := new(KewlOwnershipTransferred3)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred4Iterator is returned from FilterOwnershipTransferred4 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred4 events raised by the Kewl contract.
type KewlOwnershipTransferred4Iterator struct {
	Event *KewlOwnershipTransferred4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred4 represents a OwnershipTransferred4 event raised by the Kewl contract.
type KewlOwnershipTransferred4 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred4 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred4(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred4Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred4", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred4Iterator{contract: _Kewl.contract, event: "OwnershipTransferred4", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred4 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred4(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred4, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred4", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred4)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred4 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred4(log types.Log) (*KewlOwnershipTransferred4, error) {
	event := new(KewlOwnershipTransferred4)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred5Iterator is returned from FilterOwnershipTransferred5 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred5 events raised by the Kewl contract.
type KewlOwnershipTransferred5Iterator struct {
	Event *KewlOwnershipTransferred5 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred5Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred5)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred5)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred5Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred5Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred5 represents a OwnershipTransferred5 event raised by the Kewl contract.
type KewlOwnershipTransferred5 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred5 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred5(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred5Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred5", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred5Iterator{contract: _Kewl.contract, event: "OwnershipTransferred5", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred5 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred5(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred5, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred5", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred5)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred5", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred5 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred5(log types.Log) (*KewlOwnershipTransferred5, error) {
	event := new(KewlOwnershipTransferred5)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred5", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred6Iterator is returned from FilterOwnershipTransferred6 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred6 events raised by the Kewl contract.
type KewlOwnershipTransferred6Iterator struct {
	Event *KewlOwnershipTransferred6 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred6Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred6)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred6)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred6Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred6Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred6 represents a OwnershipTransferred6 event raised by the Kewl contract.
type KewlOwnershipTransferred6 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred6 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred6(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred6Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred6", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred6Iterator{contract: _Kewl.contract, event: "OwnershipTransferred6", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred6 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred6(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred6, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred6", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred6)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred6", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred6 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred6(log types.Log) (*KewlOwnershipTransferred6, error) {
	event := new(KewlOwnershipTransferred6)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred6", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred7Iterator is returned from FilterOwnershipTransferred7 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred7 events raised by the Kewl contract.
type KewlOwnershipTransferred7Iterator struct {
	Event *KewlOwnershipTransferred7 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred7Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred7)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred7)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred7Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred7Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred7 represents a OwnershipTransferred7 event raised by the Kewl contract.
type KewlOwnershipTransferred7 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred7 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred7(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred7Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred7", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred7Iterator{contract: _Kewl.contract, event: "OwnershipTransferred7", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred7 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred7(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred7, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred7", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred7)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred7", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred7 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred7(log types.Log) (*KewlOwnershipTransferred7, error) {
	event := new(KewlOwnershipTransferred7)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred7", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred8Iterator is returned from FilterOwnershipTransferred8 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred8 events raised by the Kewl contract.
type KewlOwnershipTransferred8Iterator struct {
	Event *KewlOwnershipTransferred8 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred8Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred8)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred8)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred8Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred8Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred8 represents a OwnershipTransferred8 event raised by the Kewl contract.
type KewlOwnershipTransferred8 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred8 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred8(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred8Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred8", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred8Iterator{contract: _Kewl.contract, event: "OwnershipTransferred8", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred8 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred8(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred8, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred8", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred8)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred8", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred8 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred8(log types.Log) (*KewlOwnershipTransferred8, error) {
	event := new(KewlOwnershipTransferred8)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred8", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlOwnershipTransferred9Iterator is returned from FilterOwnershipTransferred9 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred9 events raised by the Kewl contract.
type KewlOwnershipTransferred9Iterator struct {
	Event *KewlOwnershipTransferred9 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlOwnershipTransferred9Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlOwnershipTransferred9)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlOwnershipTransferred9)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlOwnershipTransferred9Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlOwnershipTransferred9Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlOwnershipTransferred9 represents a OwnershipTransferred9 event raised by the Kewl contract.
type KewlOwnershipTransferred9 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred9 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) FilterOwnershipTransferred9(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KewlOwnershipTransferred9Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "OwnershipTransferred9", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KewlOwnershipTransferred9Iterator{contract: _Kewl.contract, event: "OwnershipTransferred9", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred9 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) WatchOwnershipTransferred9(opts *bind.WatchOpts, sink chan<- *KewlOwnershipTransferred9, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "OwnershipTransferred9", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlOwnershipTransferred9)
				if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred9", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred9 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kewl *KewlFilterer) ParseOwnershipTransferred9(log types.Log) (*KewlOwnershipTransferred9, error) {
	event := new(KewlOwnershipTransferred9)
	if err := _Kewl.contract.UnpackLog(event, "OwnershipTransferred9", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Kewl contract.
type KewlPairCreatedIterator struct {
	Event *KewlPairCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlPairCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlPairCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlPairCreated represents a PairCreated event raised by the Kewl contract.
type KewlPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 index)
func (_Kewl *KewlFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*KewlPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &KewlPairCreatedIterator{contract: _Kewl.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 index)
func (_Kewl *KewlFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *KewlPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlPairCreated)
				if err := _Kewl.contract.UnpackLog(event, "PairCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 index)
func (_Kewl *KewlFilterer) ParsePairCreated(log types.Log) (*KewlPairCreated, error) {
	event := new(KewlPairCreated)
	if err := _Kewl.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlPairCreated0Iterator is returned from FilterPairCreated0 and is used to iterate over the raw logs and unpacked data for PairCreated0 events raised by the Kewl contract.
type KewlPairCreated0Iterator struct {
	Event *KewlPairCreated0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlPairCreated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlPairCreated0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlPairCreated0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlPairCreated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlPairCreated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlPairCreated0 represents a PairCreated0 event raised by the Kewl contract.
type KewlPairCreated0 struct {
	Token0 common.Address
	Token1 common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated0 is a free log retrieval operation binding the contract event 0xf4e6903dddf8034e1689a73259c54f96ff67410886872805530f383ededfec23.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1)
func (_Kewl *KewlFilterer) FilterPairCreated0(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*KewlPairCreated0Iterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "PairCreated0", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &KewlPairCreated0Iterator{contract: _Kewl.contract, event: "PairCreated0", logs: logs, sub: sub}, nil
}

// WatchPairCreated0 is a free log subscription operation binding the contract event 0xf4e6903dddf8034e1689a73259c54f96ff67410886872805530f383ededfec23.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1)
func (_Kewl *KewlFilterer) WatchPairCreated0(opts *bind.WatchOpts, sink chan<- *KewlPairCreated0, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "PairCreated0", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlPairCreated0)
				if err := _Kewl.contract.UnpackLog(event, "PairCreated0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePairCreated0 is a log parse operation binding the contract event 0xf4e6903dddf8034e1689a73259c54f96ff67410886872805530f383ededfec23.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1)
func (_Kewl *KewlFilterer) ParsePairCreated0(log types.Log) (*KewlPairCreated0, error) {
	event := new(KewlPairCreated0)
	if err := _Kewl.contract.UnpackLog(event, "PairCreated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Kewl contract.
type KewlRewardClaimedIterator struct {
	Event *KewlRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlRewardClaimed represents a RewardClaimed event raised by the Kewl contract.
type KewlRewardClaimed struct {
	User       common.Address
	BountyType uint8
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x8ac7271c35529ed01701c9be2f0dc971b641f52c58f9e68739da0c2421a0fc1d.
//
// Solidity: event RewardClaimed(address user, uint8 bountyType, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*KewlRewardClaimedIterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &KewlRewardClaimedIterator{contract: _Kewl.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x8ac7271c35529ed01701c9be2f0dc971b641f52c58f9e68739da0c2421a0fc1d.
//
// Solidity: event RewardClaimed(address user, uint8 bountyType, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *KewlRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlRewardClaimed)
				if err := _Kewl.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x8ac7271c35529ed01701c9be2f0dc971b641f52c58f9e68739da0c2421a0fc1d.
//
// Solidity: event RewardClaimed(address user, uint8 bountyType, uint256 amount, uint256 timestamp)
func (_Kewl *KewlFilterer) ParseRewardClaimed(log types.Log) (*KewlRewardClaimed, error) {
	event := new(KewlRewardClaimed)
	if err := _Kewl.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlTickInsertedIterator is returned from FilterTickInserted and is used to iterate over the raw logs and unpacked data for TickInserted events raised by the Kewl contract.
type KewlTickInsertedIterator struct {
	Event *KewlTickInserted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlTickInsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlTickInserted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlTickInserted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlTickInsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlTickInsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlTickInserted represents a TickInserted event raised by the Kewl contract.
type KewlTickInserted struct {
	PairHash [32]byte
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTickInserted is a free log retrieval operation binding the contract event 0x290de96699e86cd69f943008041c9edb1f0634fae543364e39953629b89f6b16.
//
// Solidity: event TickInserted(bytes32 indexed pairHash, uint256 indexed price)
func (_Kewl *KewlFilterer) FilterTickInserted(opts *bind.FilterOpts, pairHash [][32]byte, price []*big.Int) (*KewlTickInsertedIterator, error) {

	var pairHashRule []interface{}
	for _, pairHashItem := range pairHash {
		pairHashRule = append(pairHashRule, pairHashItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "TickInserted", pairHashRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &KewlTickInsertedIterator{contract: _Kewl.contract, event: "TickInserted", logs: logs, sub: sub}, nil
}

// WatchTickInserted is a free log subscription operation binding the contract event 0x290de96699e86cd69f943008041c9edb1f0634fae543364e39953629b89f6b16.
//
// Solidity: event TickInserted(bytes32 indexed pairHash, uint256 indexed price)
func (_Kewl *KewlFilterer) WatchTickInserted(opts *bind.WatchOpts, sink chan<- *KewlTickInserted, pairHash [][32]byte, price []*big.Int) (event.Subscription, error) {

	var pairHashRule []interface{}
	for _, pairHashItem := range pairHash {
		pairHashRule = append(pairHashRule, pairHashItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "TickInserted", pairHashRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlTickInserted)
				if err := _Kewl.contract.UnpackLog(event, "TickInserted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTickInserted is a log parse operation binding the contract event 0x290de96699e86cd69f943008041c9edb1f0634fae543364e39953629b89f6b16.
//
// Solidity: event TickInserted(bytes32 indexed pairHash, uint256 indexed price)
func (_Kewl *KewlFilterer) ParseTickInserted(log types.Log) (*KewlTickInserted, error) {
	event := new(KewlTickInserted)
	if err := _Kewl.contract.UnpackLog(event, "TickInserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlTickRemovedIterator is returned from FilterTickRemoved and is used to iterate over the raw logs and unpacked data for TickRemoved events raised by the Kewl contract.
type KewlTickRemovedIterator struct {
	Event *KewlTickRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlTickRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlTickRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlTickRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlTickRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlTickRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlTickRemoved represents a TickRemoved event raised by the Kewl contract.
type KewlTickRemoved struct {
	PairHash [32]byte
	Tick     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTickRemoved is a free log retrieval operation binding the contract event 0x352aa89aaa0293de12e9e097406eaebe01e90ee5c46109abe746dbf77a584a8b.
//
// Solidity: event TickRemoved(bytes32 indexed pairHash, uint256 indexed tick)
func (_Kewl *KewlFilterer) FilterTickRemoved(opts *bind.FilterOpts, pairHash [][32]byte, tick []*big.Int) (*KewlTickRemovedIterator, error) {

	var pairHashRule []interface{}
	for _, pairHashItem := range pairHash {
		pairHashRule = append(pairHashRule, pairHashItem)
	}
	var tickRule []interface{}
	for _, tickItem := range tick {
		tickRule = append(tickRule, tickItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "TickRemoved", pairHashRule, tickRule)
	if err != nil {
		return nil, err
	}
	return &KewlTickRemovedIterator{contract: _Kewl.contract, event: "TickRemoved", logs: logs, sub: sub}, nil
}

// WatchTickRemoved is a free log subscription operation binding the contract event 0x352aa89aaa0293de12e9e097406eaebe01e90ee5c46109abe746dbf77a584a8b.
//
// Solidity: event TickRemoved(bytes32 indexed pairHash, uint256 indexed tick)
func (_Kewl *KewlFilterer) WatchTickRemoved(opts *bind.WatchOpts, sink chan<- *KewlTickRemoved, pairHash [][32]byte, tick []*big.Int) (event.Subscription, error) {

	var pairHashRule []interface{}
	for _, pairHashItem := range pairHash {
		pairHashRule = append(pairHashRule, pairHashItem)
	}
	var tickRule []interface{}
	for _, tickItem := range tick {
		tickRule = append(tickRule, tickItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "TickRemoved", pairHashRule, tickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlTickRemoved)
				if err := _Kewl.contract.UnpackLog(event, "TickRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTickRemoved is a log parse operation binding the contract event 0x352aa89aaa0293de12e9e097406eaebe01e90ee5c46109abe746dbf77a584a8b.
//
// Solidity: event TickRemoved(bytes32 indexed pairHash, uint256 indexed tick)
func (_Kewl *KewlFilterer) ParseTickRemoved(log types.Log) (*KewlTickRemoved, error) {
	event := new(KewlTickRemoved)
	if err := _Kewl.contract.UnpackLog(event, "TickRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlUnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the Kewl contract.
type KewlUnwrapIterator struct {
	Event *KewlUnwrap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlUnwrap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlUnwrap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlUnwrap represents a Unwrap event raised by the Kewl contract.
type KewlUnwrap struct {
	Account         common.Address
	WrappedToken    common.Address
	Amount          *big.Int
	UnwrappedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) FilterUnwrap(opts *bind.FilterOpts) (*KewlUnwrapIterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return &KewlUnwrapIterator{contract: _Kewl.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *KewlUnwrap) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlUnwrap)
				if err := _Kewl.contract.UnpackLog(event, "Unwrap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnwrap is a log parse operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) ParseUnwrap(log types.Log) (*KewlUnwrap, error) {
	event := new(KewlUnwrap)
	if err := _Kewl.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlUnwrap0Iterator is returned from FilterUnwrap0 and is used to iterate over the raw logs and unpacked data for Unwrap0 events raised by the Kewl contract.
type KewlUnwrap0Iterator struct {
	Event *KewlUnwrap0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlUnwrap0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlUnwrap0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlUnwrap0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlUnwrap0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlUnwrap0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlUnwrap0 represents a Unwrap0 event raised by the Kewl contract.
type KewlUnwrap0 struct {
	Account         common.Address
	WrappedToken    common.Address
	Amount          *big.Int
	UnwrappedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnwrap0 is a free log retrieval operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) FilterUnwrap0(opts *bind.FilterOpts) (*KewlUnwrap0Iterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "Unwrap0")
	if err != nil {
		return nil, err
	}
	return &KewlUnwrap0Iterator{contract: _Kewl.contract, event: "Unwrap0", logs: logs, sub: sub}, nil
}

// WatchUnwrap0 is a free log subscription operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) WatchUnwrap0(opts *bind.WatchOpts, sink chan<- *KewlUnwrap0) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "Unwrap0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlUnwrap0)
				if err := _Kewl.contract.UnpackLog(event, "Unwrap0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnwrap0 is a log parse operation binding the contract event 0x844bdb9a38d1fd821344dfafdefc02959dbbcddd8e8f158cdebb56b561cb6702.
//
// Solidity: event Unwrap(address account, address wrappedToken, uint256 amount, uint256 unwrappedAmount)
func (_Kewl *KewlFilterer) ParseUnwrap0(log types.Log) (*KewlUnwrap0, error) {
	event := new(KewlUnwrap0)
	if err := _Kewl.contract.UnpackLog(event, "Unwrap0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlWrapIterator is returned from FilterWrap and is used to iterate over the raw logs and unpacked data for Wrap events raised by the Kewl contract.
type KewlWrapIterator struct {
	Event *KewlWrap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlWrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlWrap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlWrap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlWrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlWrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlWrap represents a Wrap event raised by the Kewl contract.
type KewlWrap struct {
	Account         common.Address
	UnderlyingToken common.Address
	Amount          *big.Int
	WrappedToken    common.Address
	WrappedAmount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWrap is a free log retrieval operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) FilterWrap(opts *bind.FilterOpts) (*KewlWrapIterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "Wrap")
	if err != nil {
		return nil, err
	}
	return &KewlWrapIterator{contract: _Kewl.contract, event: "Wrap", logs: logs, sub: sub}, nil
}

// WatchWrap is a free log subscription operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) WatchWrap(opts *bind.WatchOpts, sink chan<- *KewlWrap) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "Wrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlWrap)
				if err := _Kewl.contract.UnpackLog(event, "Wrap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrap is a log parse operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) ParseWrap(log types.Log) (*KewlWrap, error) {
	event := new(KewlWrap)
	if err := _Kewl.contract.UnpackLog(event, "Wrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlWrap0Iterator is returned from FilterWrap0 and is used to iterate over the raw logs and unpacked data for Wrap0 events raised by the Kewl contract.
type KewlWrap0Iterator struct {
	Event *KewlWrap0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlWrap0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlWrap0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlWrap0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlWrap0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlWrap0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlWrap0 represents a Wrap0 event raised by the Kewl contract.
type KewlWrap0 struct {
	Account         common.Address
	UnderlyingToken common.Address
	Amount          *big.Int
	WrappedToken    common.Address
	WrappedAmount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWrap0 is a free log retrieval operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) FilterWrap0(opts *bind.FilterOpts) (*KewlWrap0Iterator, error) {

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "Wrap0")
	if err != nil {
		return nil, err
	}
	return &KewlWrap0Iterator{contract: _Kewl.contract, event: "Wrap0", logs: logs, sub: sub}, nil
}

// WatchWrap0 is a free log subscription operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) WatchWrap0(opts *bind.WatchOpts, sink chan<- *KewlWrap0) (event.Subscription, error) {

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "Wrap0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlWrap0)
				if err := _Kewl.contract.UnpackLog(event, "Wrap0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrap0 is a log parse operation binding the contract event 0x73beda72ae240fb6754898b37890af8aaee5488791c693a8c89a5656d82695ec.
//
// Solidity: event Wrap(address account, address underlyingToken, uint256 amount, address wrappedToken, uint256 wrappedAmount)
func (_Kewl *KewlFilterer) ParseWrap0(log types.Log) (*KewlWrap0, error) {
	event := new(KewlWrap0)
	if err := _Kewl.contract.UnpackLog(event, "Wrap0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlWrappedTokenCreatedIterator is returned from FilterWrappedTokenCreated and is used to iterate over the raw logs and unpacked data for WrappedTokenCreated events raised by the Kewl contract.
type KewlWrappedTokenCreatedIterator struct {
	Event *KewlWrappedTokenCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlWrappedTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlWrappedTokenCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlWrappedTokenCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlWrappedTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlWrappedTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlWrappedTokenCreated represents a WrappedTokenCreated event raised by the Kewl contract.
type KewlWrappedTokenCreated struct {
	UnderlyingToken common.Address
	WrappedToken    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWrappedTokenCreated is a free log retrieval operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) FilterWrappedTokenCreated(opts *bind.FilterOpts, underlyingToken []common.Address, wrappedToken []common.Address) (*KewlWrappedTokenCreatedIterator, error) {

	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "WrappedTokenCreated", underlyingTokenRule, wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &KewlWrappedTokenCreatedIterator{contract: _Kewl.contract, event: "WrappedTokenCreated", logs: logs, sub: sub}, nil
}

// WatchWrappedTokenCreated is a free log subscription operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) WatchWrappedTokenCreated(opts *bind.WatchOpts, sink chan<- *KewlWrappedTokenCreated, underlyingToken []common.Address, wrappedToken []common.Address) (event.Subscription, error) {

	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "WrappedTokenCreated", underlyingTokenRule, wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlWrappedTokenCreated)
				if err := _Kewl.contract.UnpackLog(event, "WrappedTokenCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrappedTokenCreated is a log parse operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) ParseWrappedTokenCreated(log types.Log) (*KewlWrappedTokenCreated, error) {
	event := new(KewlWrappedTokenCreated)
	if err := _Kewl.contract.UnpackLog(event, "WrappedTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KewlWrappedTokenCreated0Iterator is returned from FilterWrappedTokenCreated0 and is used to iterate over the raw logs and unpacked data for WrappedTokenCreated0 events raised by the Kewl contract.
type KewlWrappedTokenCreated0Iterator struct {
	Event *KewlWrappedTokenCreated0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KewlWrappedTokenCreated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KewlWrappedTokenCreated0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KewlWrappedTokenCreated0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KewlWrappedTokenCreated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KewlWrappedTokenCreated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KewlWrappedTokenCreated0 represents a WrappedTokenCreated0 event raised by the Kewl contract.
type KewlWrappedTokenCreated0 struct {
	UnderlyingToken common.Address
	WrappedToken    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWrappedTokenCreated0 is a free log retrieval operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) FilterWrappedTokenCreated0(opts *bind.FilterOpts, underlyingToken []common.Address, wrappedToken []common.Address) (*KewlWrappedTokenCreated0Iterator, error) {

	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Kewl.contract.FilterLogs(opts, "WrappedTokenCreated0", underlyingTokenRule, wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &KewlWrappedTokenCreated0Iterator{contract: _Kewl.contract, event: "WrappedTokenCreated0", logs: logs, sub: sub}, nil
}

// WatchWrappedTokenCreated0 is a free log subscription operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) WatchWrappedTokenCreated0(opts *bind.WatchOpts, sink chan<- *KewlWrappedTokenCreated0, underlyingToken []common.Address, wrappedToken []common.Address) (event.Subscription, error) {

	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Kewl.contract.WatchLogs(opts, "WrappedTokenCreated0", underlyingTokenRule, wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KewlWrappedTokenCreated0)
				if err := _Kewl.contract.UnpackLog(event, "WrappedTokenCreated0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrappedTokenCreated0 is a log parse operation binding the contract event 0x47cdbd9ccb21e8c42b64c67901cd6369386d9c327365e6af6d4e479add31095e.
//
// Solidity: event WrappedTokenCreated(address indexed underlyingToken, address indexed wrappedToken)
func (_Kewl *KewlFilterer) ParseWrappedTokenCreated0(log types.Log) (*KewlWrappedTokenCreated0, error) {
	event := new(KewlWrappedTokenCreated0)
	if err := _Kewl.contract.UnpackLog(event, "WrappedTokenCreated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
