// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package supply

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

// VaultDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultDepositInfo struct {
	Token          common.Address
	TotalAmount    *big.Int
	SupplyAmount   *big.Int
	MinAmount      *big.Int
	SplitTime      *big.Int
	LastSupplyTime *big.Int
	Reciever       common.Address
}

// SupplyMetaData contains all meta data concerning the Supply contract.
var SupplyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"splitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSupplyTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structVault.DepositInfo\",\"name\":\"depositInfo\",\"type\":\"tuple\"}],\"name\":\"TokensDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"splitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSupplyTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structVault.DepositInfo\",\"name\":\"depositInfo\",\"type\":\"tuple\"}],\"name\":\"TokensSupplied\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"splitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSupplyTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"name\":\"addDepositInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"splitTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSupplyTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refillTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"supplyTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SupplyABI is the input ABI used to generate the binding from.
// Deprecated: Use SupplyMetaData.ABI instead.
var SupplyABI = SupplyMetaData.ABI

// Supply is an auto generated Go binding around an Ethereum contract.
type Supply struct {
	SupplyCaller     // Read-only binding to the contract
	SupplyTransactor // Write-only binding to the contract
	SupplyFilterer   // Log filterer for contract events
}

// SupplyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupplyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupplyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupplyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupplySession struct {
	Contract     *Supply           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupplyCallerSession struct {
	Contract *SupplyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SupplyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupplyTransactorSession struct {
	Contract     *SupplyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupplyRaw struct {
	Contract *Supply // Generic contract binding to access the raw methods on
}

// SupplyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupplyCallerRaw struct {
	Contract *SupplyCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupplyTransactorRaw struct {
	Contract *SupplyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupply creates a new instance of Supply, bound to a specific deployed contract.
func NewSupply(address common.Address, backend bind.ContractBackend) (*Supply, error) {
	contract, err := bindSupply(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Supply{SupplyCaller: SupplyCaller{contract: contract}, SupplyTransactor: SupplyTransactor{contract: contract}, SupplyFilterer: SupplyFilterer{contract: contract}}, nil
}

// NewSupplyCaller creates a new read-only instance of Supply, bound to a specific deployed contract.
func NewSupplyCaller(address common.Address, caller bind.ContractCaller) (*SupplyCaller, error) {
	contract, err := bindSupply(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyCaller{contract: contract}, nil
}

// NewSupplyTransactor creates a new write-only instance of Supply, bound to a specific deployed contract.
func NewSupplyTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyTransactor, error) {
	contract, err := bindSupply(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyTransactor{contract: contract}, nil
}

// NewSupplyFilterer creates a new log filterer instance of Supply, bound to a specific deployed contract.
func NewSupplyFilterer(address common.Address, filterer bind.ContractFilterer) (*SupplyFilterer, error) {
	contract, err := bindSupply(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupplyFilterer{contract: contract}, nil
}

// bindSupply binds a generic wrapper to an already deployed contract.
func bindSupply(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SupplyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supply *SupplyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supply.Contract.SupplyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supply *SupplyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supply.Contract.SupplyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supply *SupplyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supply.Contract.SupplyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supply *SupplyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supply.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supply *SupplyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supply.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supply *SupplyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supply.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a free data retrieval call binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address ) view returns(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever)
func (_Supply *SupplyCaller) Deposit(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token          common.Address
	TotalAmount    *big.Int
	SupplyAmount   *big.Int
	MinAmount      *big.Int
	SplitTime      *big.Int
	LastSupplyTime *big.Int
	Reciever       common.Address
}, error) {
	var out []interface{}
	err := _Supply.contract.Call(opts, &out, "deposit", arg0)

	outstruct := new(struct {
		Token          common.Address
		TotalAmount    *big.Int
		SupplyAmount   *big.Int
		MinAmount      *big.Int
		SplitTime      *big.Int
		LastSupplyTime *big.Int
		Reciever       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SupplyAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SplitTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastSupplyTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Reciever = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Deposit is a free data retrieval call binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address ) view returns(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever)
func (_Supply *SupplySession) Deposit(arg0 common.Address) (struct {
	Token          common.Address
	TotalAmount    *big.Int
	SupplyAmount   *big.Int
	MinAmount      *big.Int
	SplitTime      *big.Int
	LastSupplyTime *big.Int
	Reciever       common.Address
}, error) {
	return _Supply.Contract.Deposit(&_Supply.CallOpts, arg0)
}

// Deposit is a free data retrieval call binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address ) view returns(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever)
func (_Supply *SupplyCallerSession) Deposit(arg0 common.Address) (struct {
	Token          common.Address
	TotalAmount    *big.Int
	SupplyAmount   *big.Int
	MinAmount      *big.Int
	SplitTime      *big.Int
	LastSupplyTime *big.Int
	Reciever       common.Address
}, error) {
	return _Supply.Contract.Deposit(&_Supply.CallOpts, arg0)
}

// AddDepositInfo is a paid mutator transaction binding the contract method 0x39b63797.
//
// Solidity: function addDepositInfo(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever) payable returns()
func (_Supply *SupplyTransactor) AddDepositInfo(opts *bind.TransactOpts, token common.Address, totalAmount *big.Int, supplyAmount *big.Int, minAmount *big.Int, splitTime *big.Int, lastSupplyTime *big.Int, reciever common.Address) (*types.Transaction, error) {
	return _Supply.contract.Transact(opts, "addDepositInfo", token, totalAmount, supplyAmount, minAmount, splitTime, lastSupplyTime, reciever)
}

// AddDepositInfo is a paid mutator transaction binding the contract method 0x39b63797.
//
// Solidity: function addDepositInfo(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever) payable returns()
func (_Supply *SupplySession) AddDepositInfo(token common.Address, totalAmount *big.Int, supplyAmount *big.Int, minAmount *big.Int, splitTime *big.Int, lastSupplyTime *big.Int, reciever common.Address) (*types.Transaction, error) {
	return _Supply.Contract.AddDepositInfo(&_Supply.TransactOpts, token, totalAmount, supplyAmount, minAmount, splitTime, lastSupplyTime, reciever)
}

// AddDepositInfo is a paid mutator transaction binding the contract method 0x39b63797.
//
// Solidity: function addDepositInfo(address token, uint256 totalAmount, uint256 supplyAmount, uint256 minAmount, uint256 splitTime, uint256 lastSupplyTime, address reciever) payable returns()
func (_Supply *SupplyTransactorSession) AddDepositInfo(token common.Address, totalAmount *big.Int, supplyAmount *big.Int, minAmount *big.Int, splitTime *big.Int, lastSupplyTime *big.Int, reciever common.Address) (*types.Transaction, error) {
	return _Supply.Contract.AddDepositInfo(&_Supply.TransactOpts, token, totalAmount, supplyAmount, minAmount, splitTime, lastSupplyTime, reciever)
}

// RefillTokens is a paid mutator transaction binding the contract method 0x7ef5b6ea.
//
// Solidity: function refillTokens(uint256 amount) payable returns()
func (_Supply *SupplyTransactor) RefillTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Supply.contract.Transact(opts, "refillTokens", amount)
}

// RefillTokens is a paid mutator transaction binding the contract method 0x7ef5b6ea.
//
// Solidity: function refillTokens(uint256 amount) payable returns()
func (_Supply *SupplySession) RefillTokens(amount *big.Int) (*types.Transaction, error) {
	return _Supply.Contract.RefillTokens(&_Supply.TransactOpts, amount)
}

// RefillTokens is a paid mutator transaction binding the contract method 0x7ef5b6ea.
//
// Solidity: function refillTokens(uint256 amount) payable returns()
func (_Supply *SupplyTransactorSession) RefillTokens(amount *big.Int) (*types.Transaction, error) {
	return _Supply.Contract.RefillTokens(&_Supply.TransactOpts, amount)
}

// SupplyTokens is a paid mutator transaction binding the contract method 0x587775e8.
//
// Solidity: function supplyTokens(address supplier, uint256 timestamp) returns()
func (_Supply *SupplyTransactor) SupplyTokens(opts *bind.TransactOpts, supplier common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Supply.contract.Transact(opts, "supplyTokens", supplier, timestamp)
}

// SupplyTokens is a paid mutator transaction binding the contract method 0x587775e8.
//
// Solidity: function supplyTokens(address supplier, uint256 timestamp) returns()
func (_Supply *SupplySession) SupplyTokens(supplier common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Supply.Contract.SupplyTokens(&_Supply.TransactOpts, supplier, timestamp)
}

// SupplyTokens is a paid mutator transaction binding the contract method 0x587775e8.
//
// Solidity: function supplyTokens(address supplier, uint256 timestamp) returns()
func (_Supply *SupplyTransactorSession) SupplyTokens(supplier common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Supply.Contract.SupplyTokens(&_Supply.TransactOpts, supplier, timestamp)
}

// WithrawBalance is a paid mutator transaction binding the contract method 0x354a4fae.
//
// Solidity: function withrawBalance(address to) returns()
func (_Supply *SupplyTransactor) WithrawBalance(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Supply.contract.Transact(opts, "withrawBalance", to)
}

// WithrawBalance is a paid mutator transaction binding the contract method 0x354a4fae.
//
// Solidity: function withrawBalance(address to) returns()
func (_Supply *SupplySession) WithrawBalance(to common.Address) (*types.Transaction, error) {
	return _Supply.Contract.WithrawBalance(&_Supply.TransactOpts, to)
}

// WithrawBalance is a paid mutator transaction binding the contract method 0x354a4fae.
//
// Solidity: function withrawBalance(address to) returns()
func (_Supply *SupplyTransactorSession) WithrawBalance(to common.Address) (*types.Transaction, error) {
	return _Supply.Contract.WithrawBalance(&_Supply.TransactOpts, to)
}

// SupplyTokensDepositedIterator is returned from FilterTokensDeposited and is used to iterate over the raw logs and unpacked data for TokensDeposited events raised by the Supply contract.
type SupplyTokensDepositedIterator struct {
	Event *SupplyTokensDeposited // Event containing the contract specifics and raw log

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
func (it *SupplyTokensDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyTokensDeposited)
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
		it.Event = new(SupplyTokensDeposited)
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
func (it *SupplyTokensDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyTokensDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyTokensDeposited represents a TokensDeposited event raised by the Supply contract.
type SupplyTokensDeposited struct {
	From        common.Address
	DepositInfo VaultDepositInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensDeposited is a free log retrieval operation binding the contract event 0xa2654d7426f73f09f0bb8794594945ac999a6df3bb0e0d9074ade2182a846842.
//
// Solidity: event TokensDeposited(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) FilterTokensDeposited(opts *bind.FilterOpts, from []common.Address) (*SupplyTokensDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Supply.contract.FilterLogs(opts, "TokensDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &SupplyTokensDepositedIterator{contract: _Supply.contract, event: "TokensDeposited", logs: logs, sub: sub}, nil
}

// WatchTokensDeposited is a free log subscription operation binding the contract event 0xa2654d7426f73f09f0bb8794594945ac999a6df3bb0e0d9074ade2182a846842.
//
// Solidity: event TokensDeposited(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) WatchTokensDeposited(opts *bind.WatchOpts, sink chan<- *SupplyTokensDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Supply.contract.WatchLogs(opts, "TokensDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyTokensDeposited)
				if err := _Supply.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
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

// ParseTokensDeposited is a log parse operation binding the contract event 0xa2654d7426f73f09f0bb8794594945ac999a6df3bb0e0d9074ade2182a846842.
//
// Solidity: event TokensDeposited(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) ParseTokensDeposited(log types.Log) (*SupplyTokensDeposited, error) {
	event := new(SupplyTokensDeposited)
	if err := _Supply.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyTokensSuppliedIterator is returned from FilterTokensSupplied and is used to iterate over the raw logs and unpacked data for TokensSupplied events raised by the Supply contract.
type SupplyTokensSuppliedIterator struct {
	Event *SupplyTokensSupplied // Event containing the contract specifics and raw log

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
func (it *SupplyTokensSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyTokensSupplied)
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
		it.Event = new(SupplyTokensSupplied)
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
func (it *SupplyTokensSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyTokensSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyTokensSupplied represents a TokensSupplied event raised by the Supply contract.
type SupplyTokensSupplied struct {
	From        common.Address
	DepositInfo VaultDepositInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensSupplied is a free log retrieval operation binding the contract event 0xd1498ab2aa1fa793cf31190894fae1c1da4c11e5ffe7b2eab816a3f2da24294c.
//
// Solidity: event TokensSupplied(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) FilterTokensSupplied(opts *bind.FilterOpts, from []common.Address) (*SupplyTokensSuppliedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Supply.contract.FilterLogs(opts, "TokensSupplied", fromRule)
	if err != nil {
		return nil, err
	}
	return &SupplyTokensSuppliedIterator{contract: _Supply.contract, event: "TokensSupplied", logs: logs, sub: sub}, nil
}

// WatchTokensSupplied is a free log subscription operation binding the contract event 0xd1498ab2aa1fa793cf31190894fae1c1da4c11e5ffe7b2eab816a3f2da24294c.
//
// Solidity: event TokensSupplied(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) WatchTokensSupplied(opts *bind.WatchOpts, sink chan<- *SupplyTokensSupplied, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Supply.contract.WatchLogs(opts, "TokensSupplied", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyTokensSupplied)
				if err := _Supply.contract.UnpackLog(event, "TokensSupplied", log); err != nil {
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

// ParseTokensSupplied is a log parse operation binding the contract event 0xd1498ab2aa1fa793cf31190894fae1c1da4c11e5ffe7b2eab816a3f2da24294c.
//
// Solidity: event TokensSupplied(address indexed from, (address,uint256,uint256,uint256,uint256,uint256,address) depositInfo)
func (_Supply *SupplyFilterer) ParseTokensSupplied(log types.Log) (*SupplyTokensSupplied, error) {
	event := new(SupplyTokensSupplied)
	if err := _Supply.contract.UnpackLog(event, "TokensSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
