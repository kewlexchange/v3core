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
	ABI: "[{\"inputs\":[],\"name\":\"Ownable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Ownable__NotTransitiveOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"internalType\":\"structSwap.FlashParams\",\"name\":\"flashParams\",\"type\":\"tuple\"}],\"name\":\"handleFlash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"internalType\":\"structSwap.FlashParams[]\",\"name\":\"flashParams\",\"type\":\"tuple[]\"}],\"name\":\"handleFlashEx\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
