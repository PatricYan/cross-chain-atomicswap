// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/seeleteam/go-seele/core/types"
)

// HashedTimelockABI is the input ABI used to generate the binding from.
const HashedTimelockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_hashlock\",\"type\":\"bytes32\"},{\"name\":\"_timelock\",\"type\":\"uint256\"}],\"name\":\"newContract\",\"outputs\":[{\"name\":\"contractId\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractId\",\"type\":\"bytes32\"},{\"name\":\"_preimage\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractId\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractId\",\"type\":\"bytes32\"}],\"name\":\"getContract\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"hashlock\",\"type\":\"bytes32\"},{\"name\":\"timelock\",\"type\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"bool\"},{\"name\":\"refunded\",\"type\":\"bool\"},{\"name\":\"preimage\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"hashlock\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"LogHTLCNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractId\",\"type\":\"bytes32\"}],\"name\":\"LogHTLCWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractId\",\"type\":\"bytes32\"}],\"name\":\"LogHTLCRefund\",\"type\":\"event\"}]"

// HashedTimelockBin is the compiled bytecode used for deploying new contracts.
const HashedTimelockBin = `0x608060405234801561001057600080fd5b50610bde806100206000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663335ef5bd811461006657806363615149146100925780637249fbb6146100c1578063e16c7d98146100d9575b600080fd5b610080600160a060020a0360043516602435604435610142565b60408051918252519081900360200190f35b34801561009e57600080fd5b506100ad6004356024356104c5565b604080519115158252519081900360200190f35b3480156100cd57600080fd5b506100ad60043561083e565b3480156100e557600080fd5b506100f1600435610b0c565b60408051600160a060020a03998a1681529790981660208801528688019590955260608601939093526080850191909152151560a0840152151560c083015260e08201529051908190036101000190f35b600034811061019b576040805160e560020a62461bcd02815260206004820152601560248201527f6d73672e76616c7565206d757374206265203e20300000000000000000000000604482015290519081900360640190fd5b81428111610219576040805160e560020a62461bcd02815260206004820152602360248201527f74696d656c6f636b2074696d65206d75737420626520696e207468652066757460448201527f7572650000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b604080516c01000000000000000000000000338102602080840191909152600160a060020a03891690910260348301523460488301526068820187905260888083018790528351808403909101815260a8909201928390528151600293918291908401908083835b602083106102a05780518252601f199092019160209182019101610281565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156102df573d6000803e3d6000fd5b5050506040513d60208110156102f457600080fd5b5051915061030182610b95565b1561030b57600080fd5b6101006040519081016040528033600160a060020a0316815260200186600160a060020a0316815260200134815260200185600019168152602001848152602001600015158152602001600015158152602001600060010260001916815250600080846000191660001916815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a0316021790555060208201518160010160006101000a815481600160a060020a030219169083600160a060020a0316021790555060408201518160020155606082015181600301906000191690556080820151816004015560a08201518160050160006101000a81548160ff02191690831515021790555060c08201518160050160016101000a81548160ff02191690831515021790555060e0820151816006019060001916905590505084600160a060020a031633600160a060020a031683600019167f329a8316ed9c3b2299597538371c2944c5026574e803b1ec31d6113e1cd67bde348888604051808481526020018360001916600019168152602001828152602001935050505060405180910390a4509392505050565b600080836104d281610b95565b1515610528576040805160e560020a62461bcd02815260206004820152601960248201527f636f6e7472616374496420646f6573206e6f7420657869737400000000000000604482015290519081900360640190fd5b60408051602080820187905282518083038201815291830192839052815188938893600293909282918401908083835b602083106105775780518252601f199092019160209182019101610558565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156105b6573d6000803e3d6000fd5b5050506040513d60208110156105cb57600080fd5b505160008381526020819052604090206003015414610634576040805160e560020a62461bcd02815260206004820152601c60248201527f686173686c6f636b206861736820646f6573206e6f74206d6174636800000000604482015290519081900360640190fd5b6000878152602081905260409020600101548790600160a060020a031633146106a7576040805160e560020a62461bcd02815260206004820152601a60248201527f776974686472617761626c653a206e6f74207265636569766572000000000000604482015290519081900360640190fd5b60008181526020819052604090206005015460ff1615610711576040805160e560020a62461bcd02815260206004820152601f60248201527f776974686472617761626c653a20616c72656164792077697468647261776e00604482015290519081900360640190fd5b600081815260208190526040902060040154421061079f576040805160e560020a62461bcd02815260206004820152603160248201527f776974686472617761626c653a2074696d656c6f636b2074696d65206d75737460448201527f20626520696e2074686520667574757265000000000000000000000000000000606482015290519081900360840190fd5b600088815260208190526040808220600681018a905560058101805460ff1916600190811790915581015460028201549251919850600160a060020a03169282156108fc02929190818181858888f19350505050158015610804573d6000803e3d6000fd5b5060405188907fd6fd4c8e45bf0c70693141c7ce46451b6a6a28ac8386fca2ba914044e0e2391690600090a2506001979650505050505050565b6000808261084b81610b95565b15156108a1576040805160e560020a62461bcd02815260206004820152601960248201527f636f6e7472616374496420646f6573206e6f7420657869737400000000000000604482015290519081900360640190fd5b6000848152602081905260409020548490600160a060020a03163314610911576040805160e560020a62461bcd02815260206004820152601660248201527f726566756e6461626c653a206e6f742073656e64657200000000000000000000604482015290519081900360640190fd5b600081815260208190526040902060050154610100900460ff1615610980576040805160e560020a62461bcd02815260206004820152601c60248201527f726566756e6461626c653a20616c726561647920726566756e64656400000000604482015290519081900360640190fd5b60008181526020819052604090206005015460ff16156109ea576040805160e560020a62461bcd02815260206004820152601d60248201527f726566756e6461626c653a20616c72656164792077697468647261776e000000604482015290519081900360640190fd5b600081815260208190526040902060040154421015610a79576040805160e560020a62461bcd02815260206004820152602360248201527f726566756e6461626c653a2074696d656c6f636b206e6f74207965742070617360448201527f7365640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008581526020819052604080822060058101805461ff001916610100179055805460028201549251919650600160a060020a03169282156108fc02929190818181858888f19350505050158015610ad5573d6000803e3d6000fd5b5060405185907f989b3a845197c9aec15f8982bbb30b5da714050e662a7a287bb1a94c81e2e70e90600090a2506001949350505050565b6000806000806000806000806000610b238a610b95565b1515610b2e57610b89565b50505060008781526020819052604090208054600182015460028301546003840154600485015460058601546006870154600160a060020a039687169d50959094169a509198509650945060ff808216945061010090910416915b50919395975091939597565b600090815260208190526040902054600160a060020a03161515905600a165627a7a723058206ae87e62dfff3a6c14b5a0e7f82bf8658c82e68a35b897d310bbdf23038b8fee0029`

// DeployHashedTimelock deploys a new Ethereum contract, binding an instance of HashedTimelock to it.
func DeployHashedTimelock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashedTimelock, error) {
	parsed, err := abi.JSON(strings.NewReader(HashedTimelockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashedTimelockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashedTimelock{HashedTimelockCaller: HashedTimelockCaller{contract: contract}, HashedTimelockTransactor: HashedTimelockTransactor{contract: contract}, HashedTimelockFilterer: HashedTimelockFilterer{contract: contract}}, nil
}

// HashedTimelock is an auto generated Go binding around an Ethereum contract.
type HashedTimelock struct {
	HashedTimelockCaller     // Read-only binding to the contract
	HashedTimelockTransactor // Write-only binding to the contract
	HashedTimelockFilterer   // Log filterer for contract events
}

// HashedTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashedTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashedTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashedTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashedTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashedTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashedTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashedTimelockSession struct {
	Contract     *HashedTimelock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashedTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashedTimelockCallerSession struct {
	Contract *HashedTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HashedTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashedTimelockTransactorSession struct {
	Contract     *HashedTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HashedTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashedTimelockRaw struct {
	Contract *HashedTimelock // Generic contract binding to access the raw methods on
}

// HashedTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashedTimelockCallerRaw struct {
	Contract *HashedTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// HashedTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashedTimelockTransactorRaw struct {
	Contract *HashedTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashedTimelock creates a new instance of HashedTimelock, bound to a specific deployed contract.
func NewHashedTimelock(address common.Address, backend bind.ContractBackend) (*HashedTimelock, error) {
	contract, err := bindHashedTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashedTimelock{HashedTimelockCaller: HashedTimelockCaller{contract: contract}, HashedTimelockTransactor: HashedTimelockTransactor{contract: contract}, HashedTimelockFilterer: HashedTimelockFilterer{contract: contract}}, nil
}

// NewHashedTimelockCaller creates a new read-only instance of HashedTimelock, bound to a specific deployed contract.
func NewHashedTimelockCaller(address common.Address, caller bind.ContractCaller) (*HashedTimelockCaller, error) {
	contract, err := bindHashedTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockCaller{contract: contract}, nil
}

// NewHashedTimelockTransactor creates a new write-only instance of HashedTimelock, bound to a specific deployed contract.
func NewHashedTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*HashedTimelockTransactor, error) {
	contract, err := bindHashedTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockTransactor{contract: contract}, nil
}

// NewHashedTimelockFilterer creates a new log filterer instance of HashedTimelock, bound to a specific deployed contract.
func NewHashedTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*HashedTimelockFilterer, error) {
	contract, err := bindHashedTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockFilterer{contract: contract}, nil
}

// bindHashedTimelock binds a generic wrapper to an already deployed contract.
func bindHashedTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashedTimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashedTimelock *HashedTimelockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashedTimelock.Contract.HashedTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashedTimelock *HashedTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashedTimelock.Contract.HashedTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashedTimelock *HashedTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashedTimelock.Contract.HashedTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashedTimelock *HashedTimelockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashedTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashedTimelock *HashedTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashedTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashedTimelock *HashedTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashedTimelock.Contract.contract.Transact(opts, method, params...)
}

// GetContract is a free data retrieval call binding the contract method 0xe16c7d98.
//
// Solidity: function getContract(_contractId bytes32) constant returns(sender address, receiver address, amount uint256, hashlock bytes32, timelock uint256, withdrawn bool, refunded bool, preimage bytes32)
func (_HashedTimelock *HashedTimelockCaller) GetContract(opts *bind.CallOpts, _contractId [32]byte) (struct {
	Sender    common.Address
	Receiver  common.Address
	Amount    *big.Int
	Hashlock  [32]byte
	Timelock  *big.Int
	Withdrawn bool
	Refunded  bool
	Preimage  [32]byte
}, error) {
	ret := new(struct {
		Sender    common.Address
		Receiver  common.Address
		Amount    *big.Int
		Hashlock  [32]byte
		Timelock  *big.Int
		Withdrawn bool
		Refunded  bool
		Preimage  [32]byte
	})
	out := ret
	err := _HashedTimelock.contract.Call(opts, out, "getContract", _contractId)
	return *ret, err
}

// GetContract is a free data retrieval call binding the contract method 0xe16c7d98.
//
// Solidity: function getContract(_contractId bytes32) constant returns(sender address, receiver address, amount uint256, hashlock bytes32, timelock uint256, withdrawn bool, refunded bool, preimage bytes32)
func (_HashedTimelock *HashedTimelockSession) GetContract(_contractId [32]byte) (struct {
	Sender    common.Address
	Receiver  common.Address
	Amount    *big.Int
	Hashlock  [32]byte
	Timelock  *big.Int
	Withdrawn bool
	Refunded  bool
	Preimage  [32]byte
}, error) {
	return _HashedTimelock.Contract.GetContract(&_HashedTimelock.CallOpts, _contractId)
}

// GetContract is a free data retrieval call binding the contract method 0xe16c7d98.
//
// Solidity: function getContract(_contractId bytes32) constant returns(sender address, receiver address, amount uint256, hashlock bytes32, timelock uint256, withdrawn bool, refunded bool, preimage bytes32)
func (_HashedTimelock *HashedTimelockCallerSession) GetContract(_contractId [32]byte) (struct {
	Sender    common.Address
	Receiver  common.Address
	Amount    *big.Int
	Hashlock  [32]byte
	Timelock  *big.Int
	Withdrawn bool
	Refunded  bool
	Preimage  [32]byte
}, error) {
	return _HashedTimelock.Contract.GetContract(&_HashedTimelock.CallOpts, _contractId)
}

// NewContract is a paid mutator transaction binding the contract method 0x335ef5bd.
//
// Solidity: function newContract(_receiver address, _hashlock bytes32, _timelock uint256) returns(contractId bytes32)
func (_HashedTimelock *HashedTimelockTransactor) NewContract(opts *bind.TransactOpts, _receiver common.Address, _hashlock [32]byte, _timelock *big.Int) (*types.Transaction, error) {
	return _HashedTimelock.contract.Transact(opts, "newContract", _receiver, _hashlock, _timelock)
}

// NewContract is a paid mutator transaction binding the contract method 0x335ef5bd.
//
// Solidity: function newContract(_receiver address, _hashlock bytes32, _timelock uint256) returns(contractId bytes32)
func (_HashedTimelock *HashedTimelockSession) NewContract(_receiver common.Address, _hashlock [32]byte, _timelock *big.Int) (*types.Transaction, error) {
	return _HashedTimelock.Contract.NewContract(&_HashedTimelock.TransactOpts, _receiver, _hashlock, _timelock)
}

// NewContract is a paid mutator transaction binding the contract method 0x335ef5bd.
//
// Solidity: function newContract(_receiver address, _hashlock bytes32, _timelock uint256) returns(contractId bytes32)
func (_HashedTimelock *HashedTimelockTransactorSession) NewContract(_receiver common.Address, _hashlock [32]byte, _timelock *big.Int) (*types.Transaction, error) {
	return _HashedTimelock.Contract.NewContract(&_HashedTimelock.TransactOpts, _receiver, _hashlock, _timelock)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(_contractId bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockTransactor) Refund(opts *bind.TransactOpts, _contractId [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.contract.Transact(opts, "refund", _contractId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(_contractId bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockSession) Refund(_contractId [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.Contract.Refund(&_HashedTimelock.TransactOpts, _contractId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(_contractId bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockTransactorSession) Refund(_contractId [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.Contract.Refund(&_HashedTimelock.TransactOpts, _contractId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x63615149.
//
// Solidity: function withdraw(_contractId bytes32, _preimage bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockTransactor) Withdraw(opts *bind.TransactOpts, _contractId [32]byte, _preimage [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.contract.Transact(opts, "withdraw", _contractId, _preimage)
}

// Withdraw is a paid mutator transaction binding the contract method 0x63615149.
//
// Solidity: function withdraw(_contractId bytes32, _preimage bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockSession) Withdraw(_contractId [32]byte, _preimage [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.Contract.Withdraw(&_HashedTimelock.TransactOpts, _contractId, _preimage)
}

// Withdraw is a paid mutator transaction binding the contract method 0x63615149.
//
// Solidity: function withdraw(_contractId bytes32, _preimage bytes32) returns(bool)
func (_HashedTimelock *HashedTimelockTransactorSession) Withdraw(_contractId [32]byte, _preimage [32]byte) (*types.Transaction, error) {
	return _HashedTimelock.Contract.Withdraw(&_HashedTimelock.TransactOpts, _contractId, _preimage)
}

// HashedTimelockLogHTLCNewIterator is returned from FilterLogHTLCNew and is used to iterate over the raw logs and unpacked data for LogHTLCNew events raised by the HashedTimelock contract.
type HashedTimelockLogHTLCNewIterator struct {
	Event *HashedTimelockLogHTLCNew // Event containing the contract specifics and raw log

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
func (it *HashedTimelockLogHTLCNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashedTimelockLogHTLCNew)
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
		it.Event = new(HashedTimelockLogHTLCNew)
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
func (it *HashedTimelockLogHTLCNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashedTimelockLogHTLCNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashedTimelockLogHTLCNew represents a LogHTLCNew event raised by the HashedTimelock contract.
type HashedTimelockLogHTLCNew struct {
	ContractId [32]byte
	Sender     common.Address
	Receiver   common.Address
	Amount     *big.Int
	Hashlock   [32]byte
	Timelock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogHTLCNew is a free log retrieval operation binding the contract event 0x329a8316ed9c3b2299597538371c2944c5026574e803b1ec31d6113e1cd67bde.
//
// Solidity: e LogHTLCNew(contractId indexed bytes32, sender indexed address, receiver indexed address, amount uint256, hashlock bytes32, timelock uint256)
func (_HashedTimelock *HashedTimelockFilterer) FilterLogHTLCNew(opts *bind.FilterOpts, contractId [][32]byte, sender []common.Address, receiver []common.Address) (*HashedTimelockLogHTLCNewIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _HashedTimelock.contract.FilterLogs(opts, "LogHTLCNew", contractIdRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockLogHTLCNewIterator{contract: _HashedTimelock.contract, event: "LogHTLCNew", logs: logs, sub: sub}, nil
}

// WatchLogHTLCNew is a free log subscription operation binding the contract event 0x329a8316ed9c3b2299597538371c2944c5026574e803b1ec31d6113e1cd67bde.
//
// Solidity: e LogHTLCNew(contractId indexed bytes32, sender indexed address, receiver indexed address, amount uint256, hashlock bytes32, timelock uint256)
func (_HashedTimelock *HashedTimelockFilterer) WatchLogHTLCNew(opts *bind.WatchOpts, sink chan<- *HashedTimelockLogHTLCNew, contractId [][32]byte, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _HashedTimelock.contract.WatchLogs(opts, "LogHTLCNew", contractIdRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashedTimelockLogHTLCNew)
				if err := _HashedTimelock.contract.UnpackLog(event, "LogHTLCNew", log); err != nil {
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

// HashedTimelockLogHTLCRefundIterator is returned from FilterLogHTLCRefund and is used to iterate over the raw logs and unpacked data for LogHTLCRefund events raised by the HashedTimelock contract.
type HashedTimelockLogHTLCRefundIterator struct {
	Event *HashedTimelockLogHTLCRefund // Event containing the contract specifics and raw log

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
func (it *HashedTimelockLogHTLCRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashedTimelockLogHTLCRefund)
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
		it.Event = new(HashedTimelockLogHTLCRefund)
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
func (it *HashedTimelockLogHTLCRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashedTimelockLogHTLCRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashedTimelockLogHTLCRefund represents a LogHTLCRefund event raised by the HashedTimelock contract.
type HashedTimelockLogHTLCRefund struct {
	ContractId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogHTLCRefund is a free log retrieval operation binding the contract event 0x989b3a845197c9aec15f8982bbb30b5da714050e662a7a287bb1a94c81e2e70e.
//
// Solidity: e LogHTLCRefund(contractId indexed bytes32)
func (_HashedTimelock *HashedTimelockFilterer) FilterLogHTLCRefund(opts *bind.FilterOpts, contractId [][32]byte) (*HashedTimelockLogHTLCRefundIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _HashedTimelock.contract.FilterLogs(opts, "LogHTLCRefund", contractIdRule)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockLogHTLCRefundIterator{contract: _HashedTimelock.contract, event: "LogHTLCRefund", logs: logs, sub: sub}, nil
}

// WatchLogHTLCRefund is a free log subscription operation binding the contract event 0x989b3a845197c9aec15f8982bbb30b5da714050e662a7a287bb1a94c81e2e70e.
//
// Solidity: e LogHTLCRefund(contractId indexed bytes32)
func (_HashedTimelock *HashedTimelockFilterer) WatchLogHTLCRefund(opts *bind.WatchOpts, sink chan<- *HashedTimelockLogHTLCRefund, contractId [][32]byte) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _HashedTimelock.contract.WatchLogs(opts, "LogHTLCRefund", contractIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashedTimelockLogHTLCRefund)
				if err := _HashedTimelock.contract.UnpackLog(event, "LogHTLCRefund", log); err != nil {
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

// HashedTimelockLogHTLCWithdrawIterator is returned from FilterLogHTLCWithdraw and is used to iterate over the raw logs and unpacked data for LogHTLCWithdraw events raised by the HashedTimelock contract.
type HashedTimelockLogHTLCWithdrawIterator struct {
	Event *HashedTimelockLogHTLCWithdraw // Event containing the contract specifics and raw log

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
func (it *HashedTimelockLogHTLCWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashedTimelockLogHTLCWithdraw)
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
		it.Event = new(HashedTimelockLogHTLCWithdraw)
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
func (it *HashedTimelockLogHTLCWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashedTimelockLogHTLCWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashedTimelockLogHTLCWithdraw represents a LogHTLCWithdraw event raised by the HashedTimelock contract.
type HashedTimelockLogHTLCWithdraw struct {
	ContractId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogHTLCWithdraw is a free log retrieval operation binding the contract event 0xd6fd4c8e45bf0c70693141c7ce46451b6a6a28ac8386fca2ba914044e0e23916.
//
// Solidity: e LogHTLCWithdraw(contractId indexed bytes32)
func (_HashedTimelock *HashedTimelockFilterer) FilterLogHTLCWithdraw(opts *bind.FilterOpts, contractId [][32]byte) (*HashedTimelockLogHTLCWithdrawIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _HashedTimelock.contract.FilterLogs(opts, "LogHTLCWithdraw", contractIdRule)
	if err != nil {
		return nil, err
	}
	return &HashedTimelockLogHTLCWithdrawIterator{contract: _HashedTimelock.contract, event: "LogHTLCWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogHTLCWithdraw is a free log subscription operation binding the contract event 0xd6fd4c8e45bf0c70693141c7ce46451b6a6a28ac8386fca2ba914044e0e23916.
//
// Solidity: e LogHTLCWithdraw(contractId indexed bytes32)
func (_HashedTimelock *HashedTimelockFilterer) WatchLogHTLCWithdraw(opts *bind.WatchOpts, sink chan<- *HashedTimelockLogHTLCWithdraw, contractId [][32]byte) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _HashedTimelock.contract.WatchLogs(opts, "LogHTLCWithdraw", contractIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashedTimelockLogHTLCWithdraw)
				if err := _HashedTimelock.contract.UnpackLog(event, "LogHTLCWithdraw", log); err != nil {
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
