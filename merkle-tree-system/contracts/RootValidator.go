// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootValidator

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RootValidatorABI is the input ABI used to generate the binding from.
const RootValidatorABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limeRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"nodes\",\"type\":\"bytes32[]\"},{\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"name\":\"verifyDataInState\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootValidator is an auto generated Go binding around an Ethereum contract.
type RootValidator struct {
	RootValidatorCaller     // Read-only binding to the contract
	RootValidatorTransactor // Write-only binding to the contract
	RootValidatorFilterer   // Log filterer for contract events
}

// RootValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootValidatorSession struct {
	Contract     *RootValidator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootValidatorCallerSession struct {
	Contract *RootValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RootValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootValidatorTransactorSession struct {
	Contract     *RootValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RootValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootValidatorRaw struct {
	Contract *RootValidator // Generic contract binding to access the raw methods on
}

// RootValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootValidatorCallerRaw struct {
	Contract *RootValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// RootValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootValidatorTransactorRaw struct {
	Contract *RootValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootValidator creates a new instance of RootValidator, bound to a specific deployed contract.
func NewRootValidator(address common.Address, backend bind.ContractBackend) (*RootValidator, error) {
	contract, err := bindRootValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootValidator{RootValidatorCaller: RootValidatorCaller{contract: contract}, RootValidatorTransactor: RootValidatorTransactor{contract: contract}, RootValidatorFilterer: RootValidatorFilterer{contract: contract}}, nil
}

// NewRootValidatorCaller creates a new read-only instance of RootValidator, bound to a specific deployed contract.
func NewRootValidatorCaller(address common.Address, caller bind.ContractCaller) (*RootValidatorCaller, error) {
	contract, err := bindRootValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootValidatorCaller{contract: contract}, nil
}

// NewRootValidatorTransactor creates a new write-only instance of RootValidator, bound to a specific deployed contract.
func NewRootValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RootValidatorTransactor, error) {
	contract, err := bindRootValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootValidatorTransactor{contract: contract}, nil
}

// NewRootValidatorFilterer creates a new log filterer instance of RootValidator, bound to a specific deployed contract.
func NewRootValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RootValidatorFilterer, error) {
	contract, err := bindRootValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootValidatorFilterer{contract: contract}, nil
}

// bindRootValidator binds a generic wrapper to an already deployed contract.
func bindRootValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootValidator *RootValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootValidator.Contract.RootValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootValidator *RootValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootValidator.Contract.RootValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootValidator *RootValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootValidator.Contract.RootValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootValidator *RootValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootValidator *RootValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootValidator *RootValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootValidator.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootValidator *RootValidatorCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootValidator.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootValidator *RootValidatorSession) IsOwner() (bool, error) {
	return _RootValidator.Contract.IsOwner(&_RootValidator.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootValidator *RootValidatorCallerSession) IsOwner() (bool, error) {
	return _RootValidator.Contract.IsOwner(&_RootValidator.CallOpts)
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_RootValidator *RootValidatorCaller) LimeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootValidator.contract.Call(opts, out, "limeRoot")
	return *ret0, err
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_RootValidator *RootValidatorSession) LimeRoot() ([32]byte, error) {
	return _RootValidator.Contract.LimeRoot(&_RootValidator.CallOpts)
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_RootValidator *RootValidatorCallerSession) LimeRoot() ([32]byte, error) {
	return _RootValidator.Contract.LimeRoot(&_RootValidator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootValidator *RootValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootValidator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootValidator *RootValidatorSession) Owner() (common.Address, error) {
	return _RootValidator.Contract.Owner(&_RootValidator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootValidator *RootValidatorCallerSession) Owner() (common.Address, error) {
	return _RootValidator.Contract.Owner(&_RootValidator.CallOpts)
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_RootValidator *RootValidatorCaller) VerifyDataInState(opts *bind.CallOpts, data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootValidator.contract.Call(opts, out, "verifyDataInState", data, nodes, leafIndex)
	return *ret0, err
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_RootValidator *RootValidatorSession) VerifyDataInState(data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	return _RootValidator.Contract.VerifyDataInState(&_RootValidator.CallOpts, data, nodes, leafIndex)
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_RootValidator *RootValidatorCallerSession) VerifyDataInState(data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	return _RootValidator.Contract.VerifyDataInState(&_RootValidator.CallOpts, data, nodes, leafIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootValidator *RootValidatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootValidator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootValidator *RootValidatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootValidator.Contract.RenounceOwnership(&_RootValidator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootValidator *RootValidatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootValidator.Contract.RenounceOwnership(&_RootValidator.TransactOpts)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_RootValidator *RootValidatorTransactor) SetRoot(opts *bind.TransactOpts, merkleRoot [32]byte) (*types.Transaction, error) {
	return _RootValidator.contract.Transact(opts, "setRoot", merkleRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_RootValidator *RootValidatorSession) SetRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _RootValidator.Contract.SetRoot(&_RootValidator.TransactOpts, merkleRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_RootValidator *RootValidatorTransactorSession) SetRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _RootValidator.Contract.SetRoot(&_RootValidator.TransactOpts, merkleRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RootValidator *RootValidatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RootValidator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RootValidator *RootValidatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootValidator.Contract.TransferOwnership(&_RootValidator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RootValidator *RootValidatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootValidator.Contract.TransferOwnership(&_RootValidator.TransactOpts, newOwner)
}

// RootValidatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RootValidator contract.
type RootValidatorOwnershipTransferredIterator struct {
	Event *RootValidatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RootValidatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootValidatorOwnershipTransferred)
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
		it.Event = new(RootValidatorOwnershipTransferred)
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
func (it *RootValidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootValidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootValidatorOwnershipTransferred represents a OwnershipTransferred event raised by the RootValidator contract.
type RootValidatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RootValidator *RootValidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RootValidatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootValidator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RootValidatorOwnershipTransferredIterator{contract: _RootValidator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RootValidator *RootValidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RootValidatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootValidator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootValidatorOwnershipTransferred)
				if err := _RootValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
