// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SinglePaymentChannelABI is the input ABI used to generate the binding from.
const SinglePaymentChannelABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"startDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startChallengePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proof\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_originator\",\"type\":\"address\"}],\"name\":\"VerifyValidityOfMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proof\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"Challenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proof\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8[2]\"},{\"name\":\"_r\",\"type\":\"bytes32[2]\"},{\"name\":\"_s\",\"type\":\"bytes32[2]\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"CloseChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengePeriodLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"TimeoutClose\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bob\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPaymentProof\",\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"FinalizeChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bob\",\"type\":\"address\"}],\"name\":\"OpenChannel\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDeposited\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alice\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SinglePaymentChannelBin is the compiled bytecode used for deploying new contracts.
const SinglePaymentChannelBin = `0x60806040526103846004556201518060065534801561001d57600080fd5b5060008054600160a060020a03191633179055610f968061003f6000396000f3006080604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630b97bc8681146100d45780631310d31f146100fb57806322b55cba146101105780633fc747271461015a57806357ad3492146101845780636314e869146101a957806370dea79a146101be5780637a55d8f4146101d3578063c09cec77146101e8578063c5d330d814610219578063cf62ef6b14610247578063da02940c1461025c578063f86ccd4114610272578063fb47e3a214610287575b600080fd5b3480156100e057600080fd5b506100e961029c565b60408051918252519081900360200190f35b34801561010757600080fd5b506100e96102a2565b34801561011c57600080fd5b5061014660043560ff6024351660443560643560843560a435600160a060020a0360c435166102a8565b604080519115158252519081900360200190f35b34801561016657600080fd5b5061014660043560ff6024351660443560643560843560a435610696565b34801561019057600080fd5b506101466004356024606460a460e435610104356108cc565b3480156101b557600080fd5b506100e9610a76565b3480156101ca57600080fd5b506100e9610a7c565b3480156101df57600080fd5b50610146610a82565b3480156101f457600080fd5b506101fd610bc9565b60408051600160a060020a039092168252519081900360200190f35b34801561022557600080fd5b5061022e610bd8565b6040805192835260208301919091528051918290030190f35b34801561025357600080fd5b50610146610be1565b610270600160a060020a0360043516610d1b565b005b34801561027e57600080fd5b506100e9610f55565b34801561029357600080fd5b506101fd610f5b565b60025481565b60035481565b6000606060008060006040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509350838c6040516020018083805190602001908083835b6020831061031d5780518252601f1990920191602091820191016102fe565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193945092839250908401908083835b6020831061037d5780518252601f19909201916020918201910161035e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506001838c8c8c604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610423573d6000803e3d6000fd5b5050604051601f190151925050600160a060020a0380831690871614610493576040805160e560020a62461bcd02815260206004820152601c60248201527f7369676e6572206973206e6f7420746865206f726967696e61746f7200000000604482015290519081900360640190fd5b604080516c010000000000000000000000003002602080830191909152603482018b905260548083018b905283518084039091018152607490920192839052815191929182918401908083835b602083106104ff5780518252601f1990920191602091820191016104e0565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120935050508c82149050610585576040805160e560020a62461bcd02815260206004820152601860248201527f5468652070726f6f6620646f6573206e6f74206d617463680000000000000000604482015290519081900360640190fd5b600554881115610605576040805160e560020a62461bcd02815260206004820152602660248201527f76616c7565206578636565647320776861742077617320616d6f756e7444657060448201527f6f73697465640000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6007548711610684576040805160e560020a62461bcd02815260206004820152602260248201527f6e6f6e6365206973206e6f742067726561746572207468616e20746865206c6160448201527f7374000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060019b9a5050505050505050505050565b6000806003541115156106f3576040805160e560020a62461bcd02815260206004820152601e60248201527f6368616e6e656c206973206e6f7420696e20636c6f7365642073746174650000604482015290519081900360640190fd5b60045460035442910111610751576040805160e560020a62461bcd02815260206004820152601e60248201527f6368616c6c656e676520706572696f6420686173206e6f7420656e6465640000604482015290519081900360640190fd5b600054600160a060020a031633141561080957610788878787878787600160009054906101000a9004600160a060020a03166102a8565b1515610804576040805160e560020a62461bcd02815260206004820152602f60248201527f70726f6f66207468617420626f62207369676e65642074686973206d6573736160448201527f6765206973206e6f742076616c69640000000000000000000000000000000000606482015290519081900360840190fd5b6108a8565b61082c8787878787876000809054906101000a9004600160a060020a03166102a8565b15156108a8576040805160e560020a62461bcd02815260206004820152603160248201527f70726f6f66207468617420616c696365207369676e65642074686973206d657360448201527f73616765206973206e6f742076616c6964000000000000000000000000000000606482015290519081900360840190fd5b50604080518082019091528181526020018290526007556008555060019392505050565b6003546000901561094d576040805160e560020a62461bcd02815260206004820152602760248201527f63616e6e6f7420636c6f736520746865206368616e6e656c206d756c7469706c60448201527f652074696d657300000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600054610974908890883560ff169088359088359088908890600160a060020a03166102a8565b15156109ca576040805160e560020a62461bcd02815260206004820152601a60248201527f616c69636527732070726f6f66206973206e6f742076616c6964000000000000604482015290519081900360640190fd5b6001546109f890889060208981013560ff169189820135918901359088908890600160a060020a03166102a8565b1515610a4e576040805160e560020a62461bcd02815260206004820152601860248201527f626f6227732070726f6f66206973206e6f742076616c69640000000000000000604482015290519081900360640190fd5b5060408051808201909152818152602001829052600755600855505042600355506001919050565b60045481565b60065481565b600654600254600091014211610b08576040805160e560020a62461bcd02815260206004820152602b60248201527f74696d656f7574206f6e20746865206368616e6e656c20686173206e6f74206260448201527f65656e2072656163686564000000000000000000000000000000000000000000606482015290519081900360840190fd5b60035415610b86576040805160e560020a62461bcd02815260206004820152602260248201527f746865206368616e6e656c20697320696e2074686520636c6f7365642073746160448201527f7465000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600554604051600160a060020a039092169281156108fc029290818181858888f19350505050158015610bc1573d6000803e3d6000fd5b506001905090565b600154600160a060020a031681565b60075460085482565b600080600354111515610c3e576040805160e560020a62461bcd02815260206004820152601e60248201527f6368616e6e656c206973206e6f7420696e20636c6f7365642073746174650000604482015290519081900360640190fd5b60045460035442910110610c9c576040805160e560020a62461bcd02815260206004820152601e60248201527f6368616c6c656e676520706572696f6420686173206e6f7420656e6465640000604482015290519081900360640190fd5b600154600854604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610cd8573d6000803e3d6000fd5b5060008054600854600554604051600160a060020a039093169391900380156108fc02929091818181858888f19350505050158015610bc1573d6000803e3d6000fd5b60003411610d99576040805160e560020a62461bcd02815260206004820152602560248201527f796f75206d7573742073656e6420657468657220746f206f70656e206120636860448201527f616e6e656c000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a03163314610dfb576040805160e560020a62461bcd02815260206004820152601d60248201527f6f6e6c7920616c6963652063616e206f70656e2061206368616e6e656c000000604482015290519081900360640190fd5b600160a060020a0381161515610e81576040805160e560020a62461bcd02815260206004820152602560248201527f626f62277320616464726573732063616e6e6f7420626520746865203020616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60025415610eff576040805160e560020a62461bcd02815260206004820152602360248201527f796f752063616e6e6f742072656f70656e2061207061796d656e74206368616e60448201527f6e656c0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055426002553460058190556040805180820190915260008082526020909101829052600755600855565b60055481565b600054600160a060020a0316815600a165627a7a72305820c7c3696f8881618dd35982f9dc8c670131db3854514ffa26f962e19be9a27eb30029`

// DeploySinglePaymentChannel deploys a new Ethereum contract, binding an instance of SinglePaymentChannel to it.
func DeploySinglePaymentChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SinglePaymentChannel, error) {
	parsed, err := abi.JSON(strings.NewReader(SinglePaymentChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SinglePaymentChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SinglePaymentChannel{SinglePaymentChannelCaller: SinglePaymentChannelCaller{contract: contract}, SinglePaymentChannelTransactor: SinglePaymentChannelTransactor{contract: contract}, SinglePaymentChannelFilterer: SinglePaymentChannelFilterer{contract: contract}}, nil
}

// SinglePaymentChannel is an auto generated Go binding around an Ethereum contract.
type SinglePaymentChannel struct {
	SinglePaymentChannelCaller     // Read-only binding to the contract
	SinglePaymentChannelTransactor // Write-only binding to the contract
	SinglePaymentChannelFilterer   // Log filterer for contract events
}

// SinglePaymentChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SinglePaymentChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinglePaymentChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SinglePaymentChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinglePaymentChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SinglePaymentChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinglePaymentChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SinglePaymentChannelSession struct {
	Contract     *SinglePaymentChannel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SinglePaymentChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SinglePaymentChannelCallerSession struct {
	Contract *SinglePaymentChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SinglePaymentChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SinglePaymentChannelTransactorSession struct {
	Contract     *SinglePaymentChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SinglePaymentChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SinglePaymentChannelRaw struct {
	Contract *SinglePaymentChannel // Generic contract binding to access the raw methods on
}

// SinglePaymentChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SinglePaymentChannelCallerRaw struct {
	Contract *SinglePaymentChannelCaller // Generic read-only contract binding to access the raw methods on
}

// SinglePaymentChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SinglePaymentChannelTransactorRaw struct {
	Contract *SinglePaymentChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSinglePaymentChannel creates a new instance of SinglePaymentChannel, bound to a specific deployed contract.
func NewSinglePaymentChannel(address common.Address, backend bind.ContractBackend) (*SinglePaymentChannel, error) {
	contract, err := bindSinglePaymentChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SinglePaymentChannel{SinglePaymentChannelCaller: SinglePaymentChannelCaller{contract: contract}, SinglePaymentChannelTransactor: SinglePaymentChannelTransactor{contract: contract}, SinglePaymentChannelFilterer: SinglePaymentChannelFilterer{contract: contract}}, nil
}

// NewSinglePaymentChannelCaller creates a new read-only instance of SinglePaymentChannel, bound to a specific deployed contract.
func NewSinglePaymentChannelCaller(address common.Address, caller bind.ContractCaller) (*SinglePaymentChannelCaller, error) {
	contract, err := bindSinglePaymentChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SinglePaymentChannelCaller{contract: contract}, nil
}

// NewSinglePaymentChannelTransactor creates a new write-only instance of SinglePaymentChannel, bound to a specific deployed contract.
func NewSinglePaymentChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*SinglePaymentChannelTransactor, error) {
	contract, err := bindSinglePaymentChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SinglePaymentChannelTransactor{contract: contract}, nil
}

// NewSinglePaymentChannelFilterer creates a new log filterer instance of SinglePaymentChannel, bound to a specific deployed contract.
func NewSinglePaymentChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*SinglePaymentChannelFilterer, error) {
	contract, err := bindSinglePaymentChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SinglePaymentChannelFilterer{contract: contract}, nil
}

// bindSinglePaymentChannel binds a generic wrapper to an already deployed contract.
func bindSinglePaymentChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SinglePaymentChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SinglePaymentChannel *SinglePaymentChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SinglePaymentChannel.Contract.SinglePaymentChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SinglePaymentChannel *SinglePaymentChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.SinglePaymentChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SinglePaymentChannel *SinglePaymentChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.SinglePaymentChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SinglePaymentChannel *SinglePaymentChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SinglePaymentChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SinglePaymentChannel *SinglePaymentChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SinglePaymentChannel *SinglePaymentChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.contract.Transact(opts, method, params...)
}

// VerifyValidityOfMessage is a free data retrieval call binding the contract method 0x22b55cba.
//
// Solidity: function VerifyValidityOfMessage(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256, _originator address) constant returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) VerifyValidityOfMessage(opts *bind.CallOpts, _proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int, _originator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "VerifyValidityOfMessage", _proof, _v, _r, _s, _value, _nonce, _originator)
	return *ret0, err
}

// VerifyValidityOfMessage is a free data retrieval call binding the contract method 0x22b55cba.
//
// Solidity: function VerifyValidityOfMessage(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256, _originator address) constant returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelSession) VerifyValidityOfMessage(_proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int, _originator common.Address) (bool, error) {
	return _SinglePaymentChannel.Contract.VerifyValidityOfMessage(&_SinglePaymentChannel.CallOpts, _proof, _v, _r, _s, _value, _nonce, _originator)
}

// VerifyValidityOfMessage is a free data retrieval call binding the contract method 0x22b55cba.
//
// Solidity: function VerifyValidityOfMessage(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256, _originator address) constant returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) VerifyValidityOfMessage(_proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int, _originator common.Address) (bool, error) {
	return _SinglePaymentChannel.Contract.VerifyValidityOfMessage(&_SinglePaymentChannel.CallOpts, _proof, _v, _r, _s, _value, _nonce, _originator)
}

// Alice is a free data retrieval call binding the contract method 0xfb47e3a2.
//
// Solidity: function alice() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) Alice(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "alice")
	return *ret0, err
}

// Alice is a free data retrieval call binding the contract method 0xfb47e3a2.
//
// Solidity: function alice() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelSession) Alice() (common.Address, error) {
	return _SinglePaymentChannel.Contract.Alice(&_SinglePaymentChannel.CallOpts)
}

// Alice is a free data retrieval call binding the contract method 0xfb47e3a2.
//
// Solidity: function alice() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) Alice() (common.Address, error) {
	return _SinglePaymentChannel.Contract.Alice(&_SinglePaymentChannel.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) AmountDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "amountDeposited")
	return *ret0, err
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) AmountDeposited() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.AmountDeposited(&_SinglePaymentChannel.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) AmountDeposited() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.AmountDeposited(&_SinglePaymentChannel.CallOpts)
}

// Bob is a free data retrieval call binding the contract method 0xc09cec77.
//
// Solidity: function bob() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) Bob(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "bob")
	return *ret0, err
}

// Bob is a free data retrieval call binding the contract method 0xc09cec77.
//
// Solidity: function bob() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelSession) Bob() (common.Address, error) {
	return _SinglePaymentChannel.Contract.Bob(&_SinglePaymentChannel.CallOpts)
}

// Bob is a free data retrieval call binding the contract method 0xc09cec77.
//
// Solidity: function bob() constant returns(address)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) Bob() (common.Address, error) {
	return _SinglePaymentChannel.Contract.Bob(&_SinglePaymentChannel.CallOpts)
}

// ChallengePeriodLength is a free data retrieval call binding the contract method 0x6314e869.
//
// Solidity: function challengePeriodLength() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) ChallengePeriodLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "challengePeriodLength")
	return *ret0, err
}

// ChallengePeriodLength is a free data retrieval call binding the contract method 0x6314e869.
//
// Solidity: function challengePeriodLength() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) ChallengePeriodLength() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.ChallengePeriodLength(&_SinglePaymentChannel.CallOpts)
}

// ChallengePeriodLength is a free data retrieval call binding the contract method 0x6314e869.
//
// Solidity: function challengePeriodLength() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) ChallengePeriodLength() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.ChallengePeriodLength(&_SinglePaymentChannel.CallOpts)
}

// LastPaymentProof is a free data retrieval call binding the contract method 0xc5d330d8.
//
// Solidity: function lastPaymentProof() constant returns(nonce uint256, value uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) LastPaymentProof(opts *bind.CallOpts) (struct {
	Nonce *big.Int
	Value *big.Int
}, error) {
	ret := new(struct {
		Nonce *big.Int
		Value *big.Int
	})
	out := ret
	err := _SinglePaymentChannel.contract.Call(opts, out, "lastPaymentProof")
	return *ret, err
}

// LastPaymentProof is a free data retrieval call binding the contract method 0xc5d330d8.
//
// Solidity: function lastPaymentProof() constant returns(nonce uint256, value uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) LastPaymentProof() (struct {
	Nonce *big.Int
	Value *big.Int
}, error) {
	return _SinglePaymentChannel.Contract.LastPaymentProof(&_SinglePaymentChannel.CallOpts)
}

// LastPaymentProof is a free data retrieval call binding the contract method 0xc5d330d8.
//
// Solidity: function lastPaymentProof() constant returns(nonce uint256, value uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) LastPaymentProof() (struct {
	Nonce *big.Int
	Value *big.Int
}, error) {
	return _SinglePaymentChannel.Contract.LastPaymentProof(&_SinglePaymentChannel.CallOpts)
}

// StartChallengePeriod is a free data retrieval call binding the contract method 0x1310d31f.
//
// Solidity: function startChallengePeriod() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) StartChallengePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "startChallengePeriod")
	return *ret0, err
}

// StartChallengePeriod is a free data retrieval call binding the contract method 0x1310d31f.
//
// Solidity: function startChallengePeriod() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) StartChallengePeriod() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.StartChallengePeriod(&_SinglePaymentChannel.CallOpts)
}

// StartChallengePeriod is a free data retrieval call binding the contract method 0x1310d31f.
//
// Solidity: function startChallengePeriod() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) StartChallengePeriod() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.StartChallengePeriod(&_SinglePaymentChannel.CallOpts)
}

// StartDate is a free data retrieval call binding the contract method 0x0b97bc86.
//
// Solidity: function startDate() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) StartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "startDate")
	return *ret0, err
}

// StartDate is a free data retrieval call binding the contract method 0x0b97bc86.
//
// Solidity: function startDate() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) StartDate() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.StartDate(&_SinglePaymentChannel.CallOpts)
}

// StartDate is a free data retrieval call binding the contract method 0x0b97bc86.
//
// Solidity: function startDate() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) StartDate() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.StartDate(&_SinglePaymentChannel.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SinglePaymentChannel.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelSession) Timeout() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.Timeout(&_SinglePaymentChannel.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_SinglePaymentChannel *SinglePaymentChannelCallerSession) Timeout() (*big.Int, error) {
	return _SinglePaymentChannel.Contract.Timeout(&_SinglePaymentChannel.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x3fc74727.
//
// Solidity: function Challenge(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactor) Challenge(opts *bind.TransactOpts, _proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.contract.Transact(opts, "Challenge", _proof, _v, _r, _s, _value, _nonce)
}

// Challenge is a paid mutator transaction binding the contract method 0x3fc74727.
//
// Solidity: function Challenge(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelSession) Challenge(_proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.Challenge(&_SinglePaymentChannel.TransactOpts, _proof, _v, _r, _s, _value, _nonce)
}

// Challenge is a paid mutator transaction binding the contract method 0x3fc74727.
//
// Solidity: function Challenge(_proof bytes32, _v uint8, _r bytes32, _s bytes32, _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactorSession) Challenge(_proof [32]byte, _v uint8, _r [32]byte, _s [32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.Challenge(&_SinglePaymentChannel.TransactOpts, _proof, _v, _r, _s, _value, _nonce)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x57ad3492.
//
// Solidity: function CloseChannel(_proof bytes32, _v uint8[2], _r bytes32[2], _s bytes32[2], _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactor) CloseChannel(opts *bind.TransactOpts, _proof [32]byte, _v [2]uint8, _r [2][32]byte, _s [2][32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.contract.Transact(opts, "CloseChannel", _proof, _v, _r, _s, _value, _nonce)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x57ad3492.
//
// Solidity: function CloseChannel(_proof bytes32, _v uint8[2], _r bytes32[2], _s bytes32[2], _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelSession) CloseChannel(_proof [32]byte, _v [2]uint8, _r [2][32]byte, _s [2][32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.CloseChannel(&_SinglePaymentChannel.TransactOpts, _proof, _v, _r, _s, _value, _nonce)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x57ad3492.
//
// Solidity: function CloseChannel(_proof bytes32, _v uint8[2], _r bytes32[2], _s bytes32[2], _value uint256, _nonce uint256) returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactorSession) CloseChannel(_proof [32]byte, _v [2]uint8, _r [2][32]byte, _s [2][32]byte, _value *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.CloseChannel(&_SinglePaymentChannel.TransactOpts, _proof, _v, _r, _s, _value, _nonce)
}

// FinalizeChannel is a paid mutator transaction binding the contract method 0xcf62ef6b.
//
// Solidity: function FinalizeChannel() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactor) FinalizeChannel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SinglePaymentChannel.contract.Transact(opts, "FinalizeChannel")
}

// FinalizeChannel is a paid mutator transaction binding the contract method 0xcf62ef6b.
//
// Solidity: function FinalizeChannel() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelSession) FinalizeChannel() (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.FinalizeChannel(&_SinglePaymentChannel.TransactOpts)
}

// FinalizeChannel is a paid mutator transaction binding the contract method 0xcf62ef6b.
//
// Solidity: function FinalizeChannel() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactorSession) FinalizeChannel() (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.FinalizeChannel(&_SinglePaymentChannel.TransactOpts)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xda02940c.
//
// Solidity: function OpenChannel(_bob address) returns()
func (_SinglePaymentChannel *SinglePaymentChannelTransactor) OpenChannel(opts *bind.TransactOpts, _bob common.Address) (*types.Transaction, error) {
	return _SinglePaymentChannel.contract.Transact(opts, "OpenChannel", _bob)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xda02940c.
//
// Solidity: function OpenChannel(_bob address) returns()
func (_SinglePaymentChannel *SinglePaymentChannelSession) OpenChannel(_bob common.Address) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.OpenChannel(&_SinglePaymentChannel.TransactOpts, _bob)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xda02940c.
//
// Solidity: function OpenChannel(_bob address) returns()
func (_SinglePaymentChannel *SinglePaymentChannelTransactorSession) OpenChannel(_bob common.Address) (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.OpenChannel(&_SinglePaymentChannel.TransactOpts, _bob)
}

// TimeoutClose is a paid mutator transaction binding the contract method 0x7a55d8f4.
//
// Solidity: function TimeoutClose() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactor) TimeoutClose(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SinglePaymentChannel.contract.Transact(opts, "TimeoutClose")
}

// TimeoutClose is a paid mutator transaction binding the contract method 0x7a55d8f4.
//
// Solidity: function TimeoutClose() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelSession) TimeoutClose() (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.TimeoutClose(&_SinglePaymentChannel.TransactOpts)
}

// TimeoutClose is a paid mutator transaction binding the contract method 0x7a55d8f4.
//
// Solidity: function TimeoutClose() returns(bool)
func (_SinglePaymentChannel *SinglePaymentChannelTransactorSession) TimeoutClose() (*types.Transaction, error) {
	return _SinglePaymentChannel.Contract.TimeoutClose(&_SinglePaymentChannel.TransactOpts)
}
