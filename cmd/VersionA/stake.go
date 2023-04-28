// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BN254HashToG1G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254HashToG1G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254HashToG1G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254HashToG1G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"publicKeyHash\",\"type\":\"bytes32\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"publicKeyHash\",\"type\":\"bytes32\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WAIT_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254HashToG1.G2Point\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"hashToPoint\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"result\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"point\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"scalarMul\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254HashToG1.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254HashToG1.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254HashToG1.G2Point\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"withdraw90Percent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254HashToG1.G2Point\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254HashToG1.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"withdrawWaitFor1day\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_Stake *StakeCaller) DEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_Stake *StakeSession) DEPOSITAMOUNT() (*big.Int, error) {
	return _Stake.Contract.DEPOSITAMOUNT(&_Stake.CallOpts)
}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_Stake *StakeCallerSession) DEPOSITAMOUNT() (*big.Int, error) {
	return _Stake.Contract.DEPOSITAMOUNT(&_Stake.CallOpts)
}

// WAITTIME is a free data retrieval call binding the contract method 0x388aef5c.
//
// Solidity: function WAIT_TIME() view returns(uint256)
func (_Stake *StakeCaller) WAITTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "WAIT_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WAITTIME is a free data retrieval call binding the contract method 0x388aef5c.
//
// Solidity: function WAIT_TIME() view returns(uint256)
func (_Stake *StakeSession) WAITTIME() (*big.Int, error) {
	return _Stake.Contract.WAITTIME(&_Stake.CallOpts)
}

// WAITTIME is a free data retrieval call binding the contract method 0x388aef5c.
//
// Solidity: function WAIT_TIME() view returns(uint256)
func (_Stake *StakeCallerSession) WAITTIME() (*big.Int, error) {
	return _Stake.Contract.WAITTIME(&_Stake.CallOpts)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(uint256[2] result)
func (_Stake *StakeCaller) HashToPoint(opts *bind.CallOpts, data []byte) ([2]*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "hashToPoint", data)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(uint256[2] result)
func (_Stake *StakeSession) HashToPoint(data []byte) ([2]*big.Int, error) {
	return _Stake.Contract.HashToPoint(&_Stake.CallOpts, data)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(uint256[2] result)
func (_Stake *StakeCallerSession) HashToPoint(data []byte) ([2]*big.Int, error) {
	return _Stake.Contract.HashToPoint(&_Stake.CallOpts, data)
}

// PublicKeys is a free data retrieval call binding the contract method 0xf03dc5ef.
//
// Solidity: function publicKeys(bytes32 ) view returns(bool exists, uint256 timestamp)
func (_Stake *StakeCaller) PublicKeys(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "publicKeys", arg0)

	outstruct := new(struct {
		Exists    bool
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicKeys is a free data retrieval call binding the contract method 0xf03dc5ef.
//
// Solidity: function publicKeys(bytes32 ) view returns(bool exists, uint256 timestamp)
func (_Stake *StakeSession) PublicKeys(arg0 [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
}, error) {
	return _Stake.Contract.PublicKeys(&_Stake.CallOpts, arg0)
}

// PublicKeys is a free data retrieval call binding the contract method 0xf03dc5ef.
//
// Solidity: function publicKeys(bytes32 ) view returns(bool exists, uint256 timestamp)
func (_Stake *StakeCallerSession) PublicKeys(arg0 [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
}, error) {
	return _Stake.Contract.PublicKeys(&_Stake.CallOpts, arg0)
}

// ScalarMul is a free data retrieval call binding the contract method 0x40a05867.
//
// Solidity: function scalarMul((uint256,uint256) point, uint256 scalar) view returns((uint256,uint256))
func (_Stake *StakeCaller) ScalarMul(opts *bind.CallOpts, point BN254HashToG1G1Point, scalar *big.Int) (BN254HashToG1G1Point, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "scalarMul", point, scalar)

	if err != nil {
		return *new(BN254HashToG1G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254HashToG1G1Point)).(*BN254HashToG1G1Point)

	return out0, err

}

// ScalarMul is a free data retrieval call binding the contract method 0x40a05867.
//
// Solidity: function scalarMul((uint256,uint256) point, uint256 scalar) view returns((uint256,uint256))
func (_Stake *StakeSession) ScalarMul(point BN254HashToG1G1Point, scalar *big.Int) (BN254HashToG1G1Point, error) {
	return _Stake.Contract.ScalarMul(&_Stake.CallOpts, point, scalar)
}

// ScalarMul is a free data retrieval call binding the contract method 0x40a05867.
//
// Solidity: function scalarMul((uint256,uint256) point, uint256 scalar) view returns((uint256,uint256))
func (_Stake *StakeCallerSession) ScalarMul(point BN254HashToG1G1Point, scalar *big.Int) (BN254HashToG1G1Point, error) {
	return _Stake.Contract.ScalarMul(&_Stake.CallOpts, point, scalar)
}

// Verify is a free data retrieval call binding the contract method 0xcfe86c17.
//
// Solidity: function verify((uint256,uint256) message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeCaller) Verify(opts *bind.CallOpts, message BN254HashToG1G1Point, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "verify", message, pubKey, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xcfe86c17.
//
// Solidity: function verify((uint256,uint256) message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeSession) Verify(message BN254HashToG1G1Point, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	return _Stake.Contract.Verify(&_Stake.CallOpts, message, pubKey, signature)
}

// Verify is a free data retrieval call binding the contract method 0xcfe86c17.
//
// Solidity: function verify((uint256,uint256) message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeCallerSession) Verify(message BN254HashToG1G1Point, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	return _Stake.Contract.Verify(&_Stake.CallOpts, message, pubKey, signature)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xcb5abf1d.
//
// Solidity: function verifyMessage(bytes message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeCaller) VerifyMessage(opts *bind.CallOpts, message []byte, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "verifyMessage", message, pubKey, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0xcb5abf1d.
//
// Solidity: function verifyMessage(bytes message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeSession) VerifyMessage(message []byte, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	return _Stake.Contract.VerifyMessage(&_Stake.CallOpts, message, pubKey, signature)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xcb5abf1d.
//
// Solidity: function verifyMessage(bytes message, (uint256[2],uint256[2]) pubKey, (uint256,uint256) signature) view returns(bool)
func (_Stake *StakeCallerSession) VerifyMessage(message []byte, pubKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (bool, error) {
	return _Stake.Contract.VerifyMessage(&_Stake.CallOpts, message, pubKey, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xa7218907.
//
// Solidity: function deposit((uint256[2],uint256[2]) publicKey) payable returns()
func (_Stake *StakeTransactor) Deposit(opts *bind.TransactOpts, publicKey BN254HashToG1G2Point) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "deposit", publicKey)
}

// Deposit is a paid mutator transaction binding the contract method 0xa7218907.
//
// Solidity: function deposit((uint256[2],uint256[2]) publicKey) payable returns()
func (_Stake *StakeSession) Deposit(publicKey BN254HashToG1G2Point) (*types.Transaction, error) {
	return _Stake.Contract.Deposit(&_Stake.TransactOpts, publicKey)
}

// Deposit is a paid mutator transaction binding the contract method 0xa7218907.
//
// Solidity: function deposit((uint256[2],uint256[2]) publicKey) payable returns()
func (_Stake *StakeTransactorSession) Deposit(publicKey BN254HashToG1G2Point) (*types.Transaction, error) {
	return _Stake.Contract.Deposit(&_Stake.TransactOpts, publicKey)
}

// Withdraw90Percent is a paid mutator transaction binding the contract method 0xb3f4058e.
//
// Solidity: function withdraw90Percent((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeTransactor) Withdraw90Percent(opts *bind.TransactOpts, publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "withdraw90Percent", publicKey, signature)
}

// Withdraw90Percent is a paid mutator transaction binding the contract method 0xb3f4058e.
//
// Solidity: function withdraw90Percent((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeSession) Withdraw90Percent(publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.Contract.Withdraw90Percent(&_Stake.TransactOpts, publicKey, signature)
}

// Withdraw90Percent is a paid mutator transaction binding the contract method 0xb3f4058e.
//
// Solidity: function withdraw90Percent((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeTransactorSession) Withdraw90Percent(publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.Contract.Withdraw90Percent(&_Stake.TransactOpts, publicKey, signature)
}

// WithdrawWaitFor1day is a paid mutator transaction binding the contract method 0xe7147708.
//
// Solidity: function withdrawWaitFor1day((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeTransactor) WithdrawWaitFor1day(opts *bind.TransactOpts, publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "withdrawWaitFor1day", publicKey, signature)
}

// WithdrawWaitFor1day is a paid mutator transaction binding the contract method 0xe7147708.
//
// Solidity: function withdrawWaitFor1day((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeSession) WithdrawWaitFor1day(publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.Contract.WithdrawWaitFor1day(&_Stake.TransactOpts, publicKey, signature)
}

// WithdrawWaitFor1day is a paid mutator transaction binding the contract method 0xe7147708.
//
// Solidity: function withdrawWaitFor1day((uint256[2],uint256[2]) publicKey, (uint256,uint256) signature) returns()
func (_Stake *StakeTransactorSession) WithdrawWaitFor1day(publicKey BN254HashToG1G2Point, signature BN254HashToG1G1Point) (*types.Transaction, error) {
	return _Stake.Contract.WithdrawWaitFor1day(&_Stake.TransactOpts, publicKey, signature)
}

// StakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Stake contract.
type StakeDepositIterator struct {
	Event *StakeDeposit // Event containing the contract specifics and raw log

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
func (it *StakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeDeposit)
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
		it.Event = new(StakeDeposit)
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
func (it *StakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeDeposit represents a Deposit event raised by the Stake contract.
type StakeDeposit struct {
	PublicKeyHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x743d3067919fbf32a17803b5c0a9dd39d48691e1376031b691cc173f593b0933.
//
// Solidity: event Deposit(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) FilterDeposit(opts *bind.FilterOpts, publicKeyHash [][32]byte) (*StakeDepositIterator, error) {

	var publicKeyHashRule []interface{}
	for _, publicKeyHashItem := range publicKeyHash {
		publicKeyHashRule = append(publicKeyHashRule, publicKeyHashItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Deposit", publicKeyHashRule)
	if err != nil {
		return nil, err
	}
	return &StakeDepositIterator{contract: _Stake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x743d3067919fbf32a17803b5c0a9dd39d48691e1376031b691cc173f593b0933.
//
// Solidity: event Deposit(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakeDeposit, publicKeyHash [][32]byte) (event.Subscription, error) {

	var publicKeyHashRule []interface{}
	for _, publicKeyHashItem := range publicKeyHash {
		publicKeyHashRule = append(publicKeyHashRule, publicKeyHashItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Deposit", publicKeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeDeposit)
				if err := _Stake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x743d3067919fbf32a17803b5c0a9dd39d48691e1376031b691cc173f593b0933.
//
// Solidity: event Deposit(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) ParseDeposit(log types.Log) (*StakeDeposit, error) {
	event := new(StakeDeposit)
	if err := _Stake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Stake contract.
type StakeWithdrawIterator struct {
	Event *StakeWithdraw // Event containing the contract specifics and raw log

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
func (it *StakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeWithdraw)
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
		it.Event = new(StakeWithdraw)
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
func (it *StakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeWithdraw represents a Withdraw event raised by the Stake contract.
type StakeWithdraw struct {
	PublicKeyHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb038367bf72c9197d4f513be6ec7694358739492ca08fef497aac5dce4805e8e.
//
// Solidity: event Withdraw(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) FilterWithdraw(opts *bind.FilterOpts, publicKeyHash [][32]byte) (*StakeWithdrawIterator, error) {

	var publicKeyHashRule []interface{}
	for _, publicKeyHashItem := range publicKeyHash {
		publicKeyHashRule = append(publicKeyHashRule, publicKeyHashItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Withdraw", publicKeyHashRule)
	if err != nil {
		return nil, err
	}
	return &StakeWithdrawIterator{contract: _Stake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb038367bf72c9197d4f513be6ec7694358739492ca08fef497aac5dce4805e8e.
//
// Solidity: event Withdraw(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakeWithdraw, publicKeyHash [][32]byte) (event.Subscription, error) {

	var publicKeyHashRule []interface{}
	for _, publicKeyHashItem := range publicKeyHash {
		publicKeyHashRule = append(publicKeyHashRule, publicKeyHashItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Withdraw", publicKeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeWithdraw)
				if err := _Stake.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xb038367bf72c9197d4f513be6ec7694358739492ca08fef497aac5dce4805e8e.
//
// Solidity: event Withdraw(bytes32 indexed publicKeyHash)
func (_Stake *StakeFilterer) ParseWithdraw(log types.Log) (*StakeWithdraw, error) {
	event := new(StakeWithdraw)
	if err := _Stake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
