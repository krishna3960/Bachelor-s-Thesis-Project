// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"SwitchMotorState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"update\",\"type\":\"uint64\"}],\"name\":\"Update\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMotorState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160481b031916905561014d806100306000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80636f10c3ab14610051578063828fac601461006d578063b452edf21461008e578063b642e3ea146100a9575b600080fd5b60005460ff165b60405190151581526020015b60405180910390f35b600054610100900467ffffffffffffffff165b604051908152602001610064565b6000805460ff8082161560ff19909216821790925516610058565b6100806100b73660046100e6565b6000805468ffffffffffffffff00191661010067ffffffffffffffff9384168102919091179182905590041690565b6000602082840312156100f857600080fd5b813567ffffffffffffffff8116811461011057600080fd5b939250505056fea264697066735822122003d2a4680cb12470c70b2d94a40aca878912f70c8c78dfbf304e6d22141bd15364736f6c634300080f0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetMotorState is a free data retrieval call binding the contract method 0x6f10c3ab.
//
// Solidity: function getMotorState() view returns(bool)
func (_Api *ApiCaller) GetMotorState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getMotorState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMotorState is a free data retrieval call binding the contract method 0x6f10c3ab.
//
// Solidity: function getMotorState() view returns(bool)
func (_Api *ApiSession) GetMotorState() (bool, error) {
	return _Api.Contract.GetMotorState(&_Api.CallOpts)
}

// GetMotorState is a free data retrieval call binding the contract method 0x6f10c3ab.
//
// Solidity: function getMotorState() view returns(bool)
func (_Api *ApiCallerSession) GetMotorState() (bool, error) {
	return _Api.Contract.GetMotorState(&_Api.CallOpts)
}

// GetOpenPercentage is a free data retrieval call binding the contract method 0x828fac60.
//
// Solidity: function getOpenPercentage() view returns(uint256)
func (_Api *ApiCaller) GetOpenPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOpenPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOpenPercentage is a free data retrieval call binding the contract method 0x828fac60.
//
// Solidity: function getOpenPercentage() view returns(uint256)
func (_Api *ApiSession) GetOpenPercentage() (*big.Int, error) {
	return _Api.Contract.GetOpenPercentage(&_Api.CallOpts)
}

// GetOpenPercentage is a free data retrieval call binding the contract method 0x828fac60.
//
// Solidity: function getOpenPercentage() view returns(uint256)
func (_Api *ApiCallerSession) GetOpenPercentage() (*big.Int, error) {
	return _Api.Contract.GetOpenPercentage(&_Api.CallOpts)
}

// SwitchMotorState is a paid mutator transaction binding the contract method 0xb452edf2.
//
// Solidity: function SwitchMotorState() returns(bool)
func (_Api *ApiTransactor) SwitchMotorState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "SwitchMotorState")
}

// SwitchMotorState is a paid mutator transaction binding the contract method 0xb452edf2.
//
// Solidity: function SwitchMotorState() returns(bool)
func (_Api *ApiSession) SwitchMotorState() (*types.Transaction, error) {
	return _Api.Contract.SwitchMotorState(&_Api.TransactOpts)
}

// SwitchMotorState is a paid mutator transaction binding the contract method 0xb452edf2.
//
// Solidity: function SwitchMotorState() returns(bool)
func (_Api *ApiTransactorSession) SwitchMotorState() (*types.Transaction, error) {
	return _Api.Contract.SwitchMotorState(&_Api.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xb642e3ea.
//
// Solidity: function Update(uint64 update) returns(uint256)
func (_Api *ApiTransactor) Update(opts *bind.TransactOpts, update uint64) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Update", update)
}

// Update is a paid mutator transaction binding the contract method 0xb642e3ea.
//
// Solidity: function Update(uint64 update) returns(uint256)
func (_Api *ApiSession) Update(update uint64) (*types.Transaction, error) {
	return _Api.Contract.Update(&_Api.TransactOpts, update)
}

// Update is a paid mutator transaction binding the contract method 0xb642e3ea.
//
// Solidity: function Update(uint64 update) returns(uint256)
func (_Api *ApiTransactorSession) Update(update uint64) (*types.Transaction, error) {
	return _Api.Contract.Update(&_Api.TransactOpts, update)
}
