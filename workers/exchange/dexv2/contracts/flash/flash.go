// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flash

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

// IDiamondReadableFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondReadableFacet struct {
	Target    common.Address
	Selectors [][4]byte
}

// IDiamondWritableInternalFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondWritableInternalFacetCut struct {
	Target    common.Address
	Action    uint8
	Selectors [][4]byte
}

// SwapFlashParams is an auto generated low-level Go binding around an user-defined struct.
type SwapFlashParams struct {
	DX     *big.Int
	Profit *big.Int
	Mid    *big.Int
	Out    *big.Int
	Borrow common.Address
	Output common.Address
	Path   []common.Address
}

// FlashMetaData contains all meta data concerning the Flash contract.
var FlashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DiamondWritable__InvalidInitializationParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__RemoveTargetNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__ReplaceTargetIsIdentical\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__SelectorAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__SelectorIsImmutable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__SelectorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__SelectorNotSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DiamondWritable__TargetHasNoCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC165Base__InvalidInterfaceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Proxy__ImplementationIsNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeOwnable__NotNomineeOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondWritableInternal.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondWritableInternal.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondWritableInternal.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondWritableInternal.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondReadable.Facet[]\",\"name\":\"diamondFacets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fallbackAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nomineeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackAddress\",\"type\":\"address\"}],\"name\":\"setFallbackAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"KayenCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"internalType\":\"structSwap.FlashParams\",\"name\":\"flashParams\",\"type\":\"tuple\"}],\"name\":\"handleFlash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"internalType\":\"structSwap.FlashParams[]\",\"name\":\"flashParams\",\"type\":\"tuple[]\"}],\"name\":\"handleFlashEx\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"internalType\":\"structSwap.FlashParams\",\"name\":\"flashParams\",\"type\":\"tuple\"}],\"name\":\"handleSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlashABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashMetaData.ABI instead.
var FlashABI = FlashMetaData.ABI

// Flash is an auto generated Go binding around an Ethereum contract.
type Flash struct {
	FlashCaller     // Read-only binding to the contract
	FlashTransactor // Write-only binding to the contract
	FlashFilterer   // Log filterer for contract events
}

// FlashCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashSession struct {
	Contract     *Flash            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashCallerSession struct {
	Contract *FlashCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FlashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashTransactorSession struct {
	Contract     *FlashTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashRaw struct {
	Contract *Flash // Generic contract binding to access the raw methods on
}

// FlashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashCallerRaw struct {
	Contract *FlashCaller // Generic read-only contract binding to access the raw methods on
}

// FlashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashTransactorRaw struct {
	Contract *FlashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlash creates a new instance of Flash, bound to a specific deployed contract.
func NewFlash(address common.Address, backend bind.ContractBackend) (*Flash, error) {
	contract, err := bindFlash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flash{FlashCaller: FlashCaller{contract: contract}, FlashTransactor: FlashTransactor{contract: contract}, FlashFilterer: FlashFilterer{contract: contract}}, nil
}

// NewFlashCaller creates a new read-only instance of Flash, bound to a specific deployed contract.
func NewFlashCaller(address common.Address, caller bind.ContractCaller) (*FlashCaller, error) {
	contract, err := bindFlash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashCaller{contract: contract}, nil
}

// NewFlashTransactor creates a new write-only instance of Flash, bound to a specific deployed contract.
func NewFlashTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashTransactor, error) {
	contract, err := bindFlash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashTransactor{contract: contract}, nil
}

// NewFlashFilterer creates a new log filterer instance of Flash, bound to a specific deployed contract.
func NewFlashFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashFilterer, error) {
	contract, err := bindFlash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashFilterer{contract: contract}, nil
}

// bindFlash binds a generic wrapper to an already deployed contract.
func bindFlash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flash *FlashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flash.Contract.FlashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flash *FlashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flash.Contract.FlashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flash *FlashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flash.Contract.FlashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flash *FlashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flash *FlashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flash *FlashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flash.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector) view returns(address facet)
func (_Flash *FlashCaller) FacetAddress(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "facetAddress", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector) view returns(address facet)
func (_Flash *FlashSession) FacetAddress(selector [4]byte) (common.Address, error) {
	return _Flash.Contract.FacetAddress(&_Flash.CallOpts, selector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector) view returns(address facet)
func (_Flash *FlashCallerSession) FacetAddress(selector [4]byte) (common.Address, error) {
	return _Flash.Contract.FacetAddress(&_Flash.CallOpts, selector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_Flash *FlashCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_Flash *FlashSession) FacetAddresses() ([]common.Address, error) {
	return _Flash.Contract.FacetAddresses(&_Flash.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_Flash *FlashCallerSession) FacetAddresses() ([]common.Address, error) {
	return _Flash.Contract.FacetAddresses(&_Flash.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] selectors)
func (_Flash *FlashCaller) FacetFunctionSelectors(opts *bind.CallOpts, facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "facetFunctionSelectors", facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] selectors)
func (_Flash *FlashSession) FacetFunctionSelectors(facet common.Address) ([][4]byte, error) {
	return _Flash.Contract.FacetFunctionSelectors(&_Flash.CallOpts, facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] selectors)
func (_Flash *FlashCallerSession) FacetFunctionSelectors(facet common.Address) ([][4]byte, error) {
	return _Flash.Contract.FacetFunctionSelectors(&_Flash.CallOpts, facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] diamondFacets)
func (_Flash *FlashCaller) Facets(opts *bind.CallOpts) ([]IDiamondReadableFacet, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondReadableFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondReadableFacet)).(*[]IDiamondReadableFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] diamondFacets)
func (_Flash *FlashSession) Facets() ([]IDiamondReadableFacet, error) {
	return _Flash.Contract.Facets(&_Flash.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] diamondFacets)
func (_Flash *FlashCallerSession) Facets() ([]IDiamondReadableFacet, error) {
	return _Flash.Contract.Facets(&_Flash.CallOpts)
}

// GetFallbackAddress is a free data retrieval call binding the contract method 0x2c408059.
//
// Solidity: function getFallbackAddress() view returns(address fallbackAddress)
func (_Flash *FlashCaller) GetFallbackAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "getFallbackAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackAddress is a free data retrieval call binding the contract method 0x2c408059.
//
// Solidity: function getFallbackAddress() view returns(address fallbackAddress)
func (_Flash *FlashSession) GetFallbackAddress() (common.Address, error) {
	return _Flash.Contract.GetFallbackAddress(&_Flash.CallOpts)
}

// GetFallbackAddress is a free data retrieval call binding the contract method 0x2c408059.
//
// Solidity: function getFallbackAddress() view returns(address fallbackAddress)
func (_Flash *FlashCallerSession) GetFallbackAddress() (common.Address, error) {
	return _Flash.Contract.GetFallbackAddress(&_Flash.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flash *FlashCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flash *FlashSession) Name() (string, error) {
	return _Flash.Contract.Name(&_Flash.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flash *FlashCallerSession) Name() (string, error) {
	return _Flash.Contract.Name(&_Flash.CallOpts)
}

// NomineeOwner is a free data retrieval call binding the contract method 0x8ab5150a.
//
// Solidity: function nomineeOwner() view returns(address)
func (_Flash *FlashCaller) NomineeOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "nomineeOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NomineeOwner is a free data retrieval call binding the contract method 0x8ab5150a.
//
// Solidity: function nomineeOwner() view returns(address)
func (_Flash *FlashSession) NomineeOwner() (common.Address, error) {
	return _Flash.Contract.NomineeOwner(&_Flash.CallOpts)
}

// NomineeOwner is a free data retrieval call binding the contract method 0x8ab5150a.
//
// Solidity: function nomineeOwner() view returns(address)
func (_Flash *FlashCallerSession) NomineeOwner() (common.Address, error) {
	return _Flash.Contract.NomineeOwner(&_Flash.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flash *FlashCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flash *FlashSession) Owner() (common.Address, error) {
	return _Flash.Contract.Owner(&_Flash.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flash *FlashCallerSession) Owner() (common.Address, error) {
	return _Flash.Contract.Owner(&_Flash.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flash *FlashCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flash *FlashSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flash.Contract.SupportsInterface(&_Flash.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flash *FlashCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flash.Contract.SupportsInterface(&_Flash.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flash *FlashCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flash.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flash *FlashSession) Symbol() (string, error) {
	return _Flash.Contract.Symbol(&_Flash.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flash *FlashCallerSession) Symbol() (string, error) {
	return _Flash.Contract.Symbol(&_Flash.CallOpts)
}

// KayenCall is a paid mutator transaction binding the contract method 0xc5233bde.
//
// Solidity: function KayenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashTransactor) KayenCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "KayenCall", sender, amount0, amount1, data)
}

// KayenCall is a paid mutator transaction binding the contract method 0xc5233bde.
//
// Solidity: function KayenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashSession) KayenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.KayenCall(&_Flash.TransactOpts, sender, amount0, amount1, data)
}

// KayenCall is a paid mutator transaction binding the contract method 0xc5233bde.
//
// Solidity: function KayenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashTransactorSession) KayenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.KayenCall(&_Flash.TransactOpts, sender, amount0, amount1, data)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Flash *FlashTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Flash *FlashSession) AcceptOwnership() (*types.Transaction, error) {
	return _Flash.Contract.AcceptOwnership(&_Flash.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Flash *FlashTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Flash.Contract.AcceptOwnership(&_Flash.TransactOpts)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data) returns()
func (_Flash *FlashTransactor) DiamondCut(opts *bind.TransactOpts, facetCuts []IDiamondWritableInternalFacetCut, target common.Address, data []byte) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "diamondCut", facetCuts, target, data)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data) returns()
func (_Flash *FlashSession) DiamondCut(facetCuts []IDiamondWritableInternalFacetCut, target common.Address, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.DiamondCut(&_Flash.TransactOpts, facetCuts, target, data)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data) returns()
func (_Flash *FlashTransactorSession) DiamondCut(facetCuts []IDiamondWritableInternalFacetCut, target common.Address, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.DiamondCut(&_Flash.TransactOpts, facetCuts, target, data)
}

// HandleFlash is a paid mutator transaction binding the contract method 0x38501393.
//
// Solidity: function handleFlash((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashTransactor) HandleFlash(opts *bind.TransactOpts, flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "handleFlash", flashParams)
}

// HandleFlash is a paid mutator transaction binding the contract method 0x38501393.
//
// Solidity: function handleFlash((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashSession) HandleFlash(flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleFlash(&_Flash.TransactOpts, flashParams)
}

// HandleFlash is a paid mutator transaction binding the contract method 0x38501393.
//
// Solidity: function handleFlash((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashTransactorSession) HandleFlash(flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleFlash(&_Flash.TransactOpts, flashParams)
}

// HandleFlashEx is a paid mutator transaction binding the contract method 0x4b45e7b7.
//
// Solidity: function handleFlashEx((uint256,uint256,uint256,uint256,address,address,address[])[] flashParams) payable returns()
func (_Flash *FlashTransactor) HandleFlashEx(opts *bind.TransactOpts, flashParams []SwapFlashParams) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "handleFlashEx", flashParams)
}

// HandleFlashEx is a paid mutator transaction binding the contract method 0x4b45e7b7.
//
// Solidity: function handleFlashEx((uint256,uint256,uint256,uint256,address,address,address[])[] flashParams) payable returns()
func (_Flash *FlashSession) HandleFlashEx(flashParams []SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleFlashEx(&_Flash.TransactOpts, flashParams)
}

// HandleFlashEx is a paid mutator transaction binding the contract method 0x4b45e7b7.
//
// Solidity: function handleFlashEx((uint256,uint256,uint256,uint256,address,address,address[])[] flashParams) payable returns()
func (_Flash *FlashTransactorSession) HandleFlashEx(flashParams []SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleFlashEx(&_Flash.TransactOpts, flashParams)
}

// HandleSwap is a paid mutator transaction binding the contract method 0x1039fb0c.
//
// Solidity: function handleSwap((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashTransactor) HandleSwap(opts *bind.TransactOpts, flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "handleSwap", flashParams)
}

// HandleSwap is a paid mutator transaction binding the contract method 0x1039fb0c.
//
// Solidity: function handleSwap((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashSession) HandleSwap(flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleSwap(&_Flash.TransactOpts, flashParams)
}

// HandleSwap is a paid mutator transaction binding the contract method 0x1039fb0c.
//
// Solidity: function handleSwap((uint256,uint256,uint256,uint256,address,address,address[]) flashParams) payable returns()
func (_Flash *FlashTransactorSession) HandleSwap(flashParams SwapFlashParams) (*types.Transaction, error) {
	return _Flash.Contract.HandleSwap(&_Flash.TransactOpts, flashParams)
}

// SetFallbackAddress is a paid mutator transaction binding the contract method 0x91423765.
//
// Solidity: function setFallbackAddress(address fallbackAddress) returns()
func (_Flash *FlashTransactor) SetFallbackAddress(opts *bind.TransactOpts, fallbackAddress common.Address) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "setFallbackAddress", fallbackAddress)
}

// SetFallbackAddress is a paid mutator transaction binding the contract method 0x91423765.
//
// Solidity: function setFallbackAddress(address fallbackAddress) returns()
func (_Flash *FlashSession) SetFallbackAddress(fallbackAddress common.Address) (*types.Transaction, error) {
	return _Flash.Contract.SetFallbackAddress(&_Flash.TransactOpts, fallbackAddress)
}

// SetFallbackAddress is a paid mutator transaction binding the contract method 0x91423765.
//
// Solidity: function setFallbackAddress(address fallbackAddress) returns()
func (_Flash *FlashTransactorSession) SetFallbackAddress(fallbackAddress common.Address) (*types.Transaction, error) {
	return _Flash.Contract.SetFallbackAddress(&_Flash.TransactOpts, fallbackAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Flash *FlashTransactor) TransferOwnership(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "transferOwnership", account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Flash *FlashSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Flash.Contract.TransferOwnership(&_Flash.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Flash *FlashTransactorSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Flash.Contract.TransferOwnership(&_Flash.TransactOpts, account)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.UniswapV2Call(&_Flash.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Flash *FlashTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Flash.Contract.UniswapV2Call(&_Flash.TransactOpts, sender, amount0, amount1, data)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x50431ce4.
//
// Solidity: function withdrawNative() returns()
func (_Flash *FlashTransactor) WithdrawNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "withdrawNative")
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x50431ce4.
//
// Solidity: function withdrawNative() returns()
func (_Flash *FlashSession) WithdrawNative() (*types.Transaction, error) {
	return _Flash.Contract.WithdrawNative(&_Flash.TransactOpts)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x50431ce4.
//
// Solidity: function withdrawNative() returns()
func (_Flash *FlashTransactorSession) WithdrawNative() (*types.Transaction, error) {
	return _Flash.Contract.WithdrawNative(&_Flash.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_Flash *FlashTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Flash.contract.Transact(opts, "withdrawToken", token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_Flash *FlashSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _Flash.Contract.WithdrawToken(&_Flash.TransactOpts, token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_Flash *FlashTransactorSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _Flash.Contract.WithdrawToken(&_Flash.TransactOpts, token)
}

// FlashDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the Flash contract.
type FlashDiamondCutIterator struct {
	Event *FlashDiamondCut // Event containing the contract specifics and raw log

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
func (it *FlashDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashDiamondCut)
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
		it.Event = new(FlashDiamondCut)
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
func (it *FlashDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashDiamondCut represents a DiamondCut event raised by the Flash contract.
type FlashDiamondCut struct {
	FacetCuts []IDiamondWritableInternalFacetCut
	Target    common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data)
func (_Flash *FlashFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*FlashDiamondCutIterator, error) {

	logs, sub, err := _Flash.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &FlashDiamondCutIterator{contract: _Flash.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data)
func (_Flash *FlashFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *FlashDiamondCut) (event.Subscription, error) {

	logs, sub, err := _Flash.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashDiamondCut)
				if err := _Flash.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facetCuts, address target, bytes data)
func (_Flash *FlashFilterer) ParseDiamondCut(log types.Log) (*FlashDiamondCut, error) {
	event := new(FlashDiamondCut)
	if err := _Flash.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Flash contract.
type FlashOwnershipTransferredIterator struct {
	Event *FlashOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FlashOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashOwnershipTransferred)
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
		it.Event = new(FlashOwnershipTransferred)
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
func (it *FlashOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashOwnershipTransferred represents a OwnershipTransferred event raised by the Flash contract.
type FlashOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flash *FlashFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlashOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flash.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlashOwnershipTransferredIterator{contract: _Flash.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flash *FlashFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlashOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flash.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashOwnershipTransferred)
				if err := _Flash.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Flash *FlashFilterer) ParseOwnershipTransferred(log types.Log) (*FlashOwnershipTransferred, error) {
	event := new(FlashOwnershipTransferred)
	if err := _Flash.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashOwnershipTransferred0Iterator is returned from FilterOwnershipTransferred0 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred0 events raised by the Flash contract.
type FlashOwnershipTransferred0Iterator struct {
	Event *FlashOwnershipTransferred0 // Event containing the contract specifics and raw log

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
func (it *FlashOwnershipTransferred0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashOwnershipTransferred0)
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
		it.Event = new(FlashOwnershipTransferred0)
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
func (it *FlashOwnershipTransferred0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashOwnershipTransferred0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashOwnershipTransferred0 represents a OwnershipTransferred0 event raised by the Flash contract.
type FlashOwnershipTransferred0 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred0 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flash *FlashFilterer) FilterOwnershipTransferred0(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlashOwnershipTransferred0Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flash.contract.FilterLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlashOwnershipTransferred0Iterator{contract: _Flash.contract, event: "OwnershipTransferred0", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred0 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flash *FlashFilterer) WatchOwnershipTransferred0(opts *bind.WatchOpts, sink chan<- *FlashOwnershipTransferred0, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flash.contract.WatchLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashOwnershipTransferred0)
				if err := _Flash.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
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
func (_Flash *FlashFilterer) ParseOwnershipTransferred0(log types.Log) (*FlashOwnershipTransferred0, error) {
	event := new(FlashOwnershipTransferred0)
	if err := _Flash.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
