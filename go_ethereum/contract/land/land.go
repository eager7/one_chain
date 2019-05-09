// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package land

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LandABI is the input ABI used to generate the binding from.
const LandABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proxyOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"initializedWith\",\"type\":\"bytes\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// LandBin is the compiled bytecode used for deploying new contracts.
const LandBin = `606060405233600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107a6806100956000396000f30060606040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063025313a214610114578063721d7d8e146101695780638da5cb5b146101be578063c987336c14610213578063f2fde38b1461028f575b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156100b457600080fd5b6101126000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000368080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050506102c8565b005b341561011f57600080fd5b610127610306565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561017457600080fd5b61017c61032c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101c957600080fd5b6101d1610351565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561021e57600080fd5b61028d600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610377565b005b341561029a57600080fd5b6102c6600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506105b1565b005b6102d182610767565b15156102dc57600080fd5b600080825160208401856127105a03f43d604051816000823e8260008114610302578282f35b8282fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103d357600080fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503073ffffffffffffffffffffffffffffffffffffffff1663439fab91826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561049b578082015181840152602081019050610480565b50505050905090810190601f1680156104c85780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15156104e657600080fd5b6102c65a03f115156104f757600080fd5b5050508173ffffffffffffffffffffffffffffffffffffffff167fe74baeef5988edac1159d9177ca52f0f3d68f624a1996f77467eb3ebfb316537826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610573578082015181840152602081019050610558565b50505050905090810190601f1680156105a05780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561060d57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561066a57600080fd5b7f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a180600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080823b9050600081119150509190505600a165627a7a723058209c5cf1525023f56b020c8e8faef83d3057389d0673a91d557b8dcd49d811dae90029`

// DeployLand deploys a new Ethereum contract, binding an instance of Land to it.
func DeployLand(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Land, error) {
	parsed, err := abi.JSON(strings.NewReader(LandABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LandBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Land{LandCaller: LandCaller{contract: contract}, LandTransactor: LandTransactor{contract: contract}, LandFilterer: LandFilterer{contract: contract}}, nil
}

// Land is an auto generated Go binding around an Ethereum contract.
type Land struct {
	LandCaller     // Read-only binding to the contract
	LandTransactor // Write-only binding to the contract
	LandFilterer   // Log filterer for contract events
}

// LandCaller is an auto generated read-only Go binding around an Ethereum contract.
type LandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LandSession struct {
	Contract     *Land             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LandCallerSession struct {
	Contract *LandCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LandTransactorSession struct {
	Contract     *LandTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandRaw is an auto generated low-level Go binding around an Ethereum contract.
type LandRaw struct {
	Contract *Land // Generic contract binding to access the raw methods on
}

// LandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LandCallerRaw struct {
	Contract *LandCaller // Generic read-only contract binding to access the raw methods on
}

// LandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LandTransactorRaw struct {
	Contract *LandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLand creates a new instance of Land, bound to a specific deployed contract.
func NewLand(address common.Address, backend bind.ContractBackend) (*Land, error) {
	contract, err := bindLand(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Land{LandCaller: LandCaller{contract: contract}, LandTransactor: LandTransactor{contract: contract}, LandFilterer: LandFilterer{contract: contract}}, nil
}

// NewLandCaller creates a new read-only instance of Land, bound to a specific deployed contract.
func NewLandCaller(address common.Address, caller bind.ContractCaller) (*LandCaller, error) {
	contract, err := bindLand(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LandCaller{contract: contract}, nil
}

// NewLandTransactor creates a new write-only instance of Land, bound to a specific deployed contract.
func NewLandTransactor(address common.Address, transactor bind.ContractTransactor) (*LandTransactor, error) {
	contract, err := bindLand(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LandTransactor{contract: contract}, nil
}

// NewLandFilterer creates a new log filterer instance of Land, bound to a specific deployed contract.
func NewLandFilterer(address common.Address, filterer bind.ContractFilterer) (*LandFilterer, error) {
	contract, err := bindLand(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LandFilterer{contract: contract}, nil
}

// bindLand binds a generic wrapper to an already deployed contract.
func bindLand(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LandABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Land.Contract.LandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Land.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.contract.Transact(opts, method, params...)
}

// CurrentContract is a free data retrieval call binding the contract method 0x721d7d8e.
//
// Solidity: function currentContract() constant returns(address)
func (_Land *LandCaller) CurrentContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Land.contract.Call(opts, out, "currentContract")
	return *ret0, err
}

// CurrentContract is a free data retrieval call binding the contract method 0x721d7d8e.
//
// Solidity: function currentContract() constant returns(address)
func (_Land *LandSession) CurrentContract() (common.Address, error) {
	return _Land.Contract.CurrentContract(&_Land.CallOpts)
}

// CurrentContract is a free data retrieval call binding the contract method 0x721d7d8e.
//
// Solidity: function currentContract() constant returns(address)
func (_Land *LandCallerSession) CurrentContract() (common.Address, error) {
	return _Land.Contract.CurrentContract(&_Land.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Land *LandCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Land.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Land *LandSession) Owner() (common.Address, error) {
	return _Land.Contract.Owner(&_Land.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Land *LandCallerSession) Owner() (common.Address, error) {
	return _Land.Contract.Owner(&_Land.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() constant returns(address)
func (_Land *LandCaller) ProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Land.contract.Call(opts, out, "proxyOwner")
	return *ret0, err
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() constant returns(address)
func (_Land *LandSession) ProxyOwner() (common.Address, error) {
	return _Land.Contract.ProxyOwner(&_Land.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() constant returns(address)
func (_Land *LandCallerSession) ProxyOwner() (common.Address, error) {
	return _Land.Contract.ProxyOwner(&_Land.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Land *LandTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Land *LandSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Land.Contract.TransferOwnership(&_Land.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Land *LandTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Land.Contract.TransferOwnership(&_Land.TransactOpts, _newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc987336c.
//
// Solidity: function upgrade(address newContract, bytes data) returns()
func (_Land *LandTransactor) Upgrade(opts *bind.TransactOpts, newContract common.Address, data []byte) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "upgrade", newContract, data)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc987336c.
//
// Solidity: function upgrade(address newContract, bytes data) returns()
func (_Land *LandSession) Upgrade(newContract common.Address, data []byte) (*types.Transaction, error) {
	return _Land.Contract.Upgrade(&_Land.TransactOpts, newContract, data)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc987336c.
//
// Solidity: function upgrade(address newContract, bytes data) returns()
func (_Land *LandTransactorSession) Upgrade(newContract common.Address, data []byte) (*types.Transaction, error) {
	return _Land.Contract.Upgrade(&_Land.TransactOpts, newContract, data)
}

// LandOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the Land contract.
type LandOwnerUpdateIterator struct {
	Event *LandOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *LandOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandOwnerUpdate)
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
		it.Event = new(LandOwnerUpdate)
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
func (it *LandOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandOwnerUpdate represents a OwnerUpdate event raised by the Land contract.
type LandOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address _prevOwner, address _newOwner)
func (_Land *LandFilterer) FilterOwnerUpdate(opts *bind.FilterOpts) (*LandOwnerUpdateIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return &LandOwnerUpdateIterator{contract: _Land.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address _prevOwner, address _newOwner)
func (_Land *LandFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *LandOwnerUpdate) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandOwnerUpdate)
				if err := _Land.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// LandUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the Land contract.
type LandUpgradeIterator struct {
	Event *LandUpgrade // Event containing the contract specifics and raw log

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
func (it *LandUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandUpgrade)
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
		it.Event = new(LandUpgrade)
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
func (it *LandUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandUpgrade represents a Upgrade event raised by the Land contract.
type LandUpgrade struct {
	NewContract     common.Address
	InitializedWith []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0xe74baeef5988edac1159d9177ca52f0f3d68f624a1996f77467eb3ebfb316537.
//
// Solidity: event Upgrade(address indexed newContract, bytes initializedWith)
func (_Land *LandFilterer) FilterUpgrade(opts *bind.FilterOpts, newContract []common.Address) (*LandUpgradeIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "Upgrade", newContractRule)
	if err != nil {
		return nil, err
	}
	return &LandUpgradeIterator{contract: _Land.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0xe74baeef5988edac1159d9177ca52f0f3d68f624a1996f77467eb3ebfb316537.
//
// Solidity: event Upgrade(address indexed newContract, bytes initializedWith)
func (_Land *LandFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *LandUpgrade, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "Upgrade", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandUpgrade)
				if err := _Land.contract.UnpackLog(event, "Upgrade", log); err != nil {
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
