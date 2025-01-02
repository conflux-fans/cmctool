// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pospool

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// IVotingEscrowLockInfo is an auto generated low-level Go binding around an user-defined struct.
type IVotingEscrowLockInfo struct {
	Amount      *big.Int
	UnlockBlock *big.Int
}

// ParamsControlVote is an auto generated low-level Go binding around an user-defined struct.
type ParamsControlVote struct {
	TopicIndex uint16
	Votes      [3]*big.Int
}

// PoSPoolPoolShot is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolPoolShot struct {
	Available   *big.Int
	Balance     *big.Int
	BlockNumber *big.Int
}

// PoSPoolPoolSummary is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolPoolSummary struct {
	Available     *big.Int
	Interest      *big.Int
	TotalInterest *big.Int
}

// PoSPoolUserShot is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolUserShot struct {
	Available       *big.Int
	AccRewardPerCfx *big.Int
	BlockNumber     *big.Int
}

// PoSPoolUserSummary is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolUserSummary struct {
	Votes           *big.Int
	Available       *big.Int
	Locked          *big.Int
	Unlocked        *big.Int
	ClaimedInterest *big.Int
	CurrentInterest *big.Int
}

// VotePowerQueueQueueNode is an auto generated low-level Go binding around an user-defined struct.
type VotePowerQueueQueueNode struct {
	VotePower *big.Int
	EndBlock  *big.Int
}

// PosPoolABI is the input ABI used to generate the binding from.
const PosPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"DecreasePoSStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"IncreasePoSStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"RatioChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"WithdrawStake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolRegisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votes\",\"type\":\"uint64\"}],\"name\":\"_restakePosVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_restakeUserStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"}],\"name\":\"_retireUserStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"_userShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_withdrawPoolProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accRewardPerCfx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_freeAddress\",\"type\":\"address\"}],\"name\":\"addToFeeFreeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"internalType\":\"structParamsControl.Vote[]\",\"name\":\"vote_data\",\"type\":\"tuple[]\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAllInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"decreaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlockNumber\",\"type\":\"uint256\"}],\"name\":\"lockForVotePower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramsControl\",\"outputs\":[{\"internalType\":\"contractParamsControl\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolShot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.PoolShot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalInterest\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.PoolSummary\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolUserShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"posAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"indentifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[2]\",\"name\":\"blsPubKeyProof\",\"type\":\"bytes[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_freeAddress\",\"type\":\"address\"}],\"name\":\"removeFromFeeFreeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"setCfxCountOfOneVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"name\":\"setLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setParamsControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ratio\",\"type\":\"uint64\"}],\"name\":\"setPoolUserShareRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"name\":\"setUnlockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_votingEscrow\",\"type\":\"address\"}],\"name\":\"setVotingEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"stakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"userInQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userInQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"userInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIVotingEscrow.LockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userOutQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"userOutQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allUserShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userShot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerCfx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.UserShot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"locked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentInterest\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.UserSummary\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userVotePower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEscrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PosPool is an auto generated Go binding around an Conflux contract.
type PosPool struct {
	PosPoolCaller     // Read-only binding to the contract
	PosPoolTransactor // Write-only binding to the contract
	PosPoolFilterer   // Log filterer for contract events
}

// PosPoolCaller is an auto generated read-only Go binding around an Conflux contract.
type PosPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type PosPoolBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolTransactor is an auto generated write-only Go binding around an Conflux contract.
type PosPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type PosPoolBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type PosPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type PosPoolSession struct {
	Contract     *PosPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PosPoolCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type PosPoolCallerSession struct {
	Contract *PosPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PosPoolTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type PosPoolTransactorSession struct {
	Contract     *PosPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PosPoolRaw is an auto generated low-level Go binding around an Conflux contract.
type PosPoolRaw struct {
	Contract *PosPool // Generic contract binding to access the raw methods on
}

// PosPoolCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type PosPoolCallerRaw struct {
	Contract *PosPoolCaller // Generic read-only contract binding to access the raw methods on
}

// PosPoolTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type PosPoolTransactorRaw struct {
	Contract *PosPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPosPool creates a new instance of PosPool, bound to a specific deployed contract.
func NewPosPool(address types.Address, backend bind.ContractBackend) (*PosPool, error) {
	contract, err := bindPosPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PosPool{PosPoolCaller: PosPoolCaller{contract: contract}, PosPoolTransactor: PosPoolTransactor{contract: contract}, PosPoolFilterer: PosPoolFilterer{contract: contract}}, nil
}

// NewPosPoolCaller creates a new read-only instance of PosPool, bound to a specific deployed contract.
func NewPosPoolCaller(address types.Address, caller bind.ContractCaller) (*PosPoolCaller, error) {
	contract, err := bindPosPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolCaller{contract: contract}, nil
}

// NewPosPoolTransactor creates a new write-only instance of PosPool, bound to a specific deployed contract.
func NewPosPoolTransactor(address types.Address, transactor bind.ContractTransactor) (*PosPoolTransactor, error) {
	contract, err := bindPosPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolTransactor{contract: contract}, nil
}

// NewPosPoolFilterer creates a new log filterer instance of PosPool, bound to a specific deployed contract.
func NewPosPoolFilterer(address types.Address, filterer bind.ContractFilterer) (*PosPoolFilterer, error) {
	contract, err := bindPosPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PosPoolFilterer{contract: contract}, nil
}

// NewPosPoolCaller creates a new read-only instance of PosPool, bound to a specific deployed contract.
func NewPosPoolBulkCaller(address types.Address, caller bind.ContractCaller) (*PosPoolBulkCaller, error) {
	contract, err := bindPosPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolBulkCaller{contract: contract}, nil
}

// NewPosPoolBulkTransactor creates a new write-only instance of PosPool, bound to a specific deployed contract.
func NewPosPoolBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*PosPoolBulkTransactor, error) {
	contract, err := bindPosPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolBulkTransactor{contract: contract}, nil
}

// bindPosPool binds a generic wrapper to an already deployed contract.
func bindPosPool(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PosPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PosPool *PosPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PosPool.Contract.PosPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PosPool *PosPoolRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.PosPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PosPool *PosPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.PosPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PosPool *PosPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PosPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PosPool *PosPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PosPool *PosPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPool *PosPoolCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "VERSION")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPool *PosPoolBulkCaller) VERSION(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "VERSION")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "VERSION")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPool *PosPoolSession) VERSION() (string, error) {
	return _PosPool.Contract.VERSION(&_PosPool.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPool *PosPoolCallerSession) VERSION() (string, error) {
	return _PosPool.Contract.VERSION(&_PosPool.CallOpts)
}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPool *PosPoolCaller) PoolLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "_poolLockPeriod")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) PoolLockPeriod(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "_poolLockPeriod")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "_poolLockPeriod")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPool *PosPoolSession) PoolLockPeriod() (*big.Int, error) {
	return _PosPool.Contract.PoolLockPeriod(&_PosPool.CallOpts)
}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPool *PosPoolCallerSession) PoolLockPeriod() (*big.Int, error) {
	return _PosPool.Contract.PoolLockPeriod(&_PosPool.CallOpts)
}

// PoolRegisted is a free data retrieval call binding the contract method 0x4b1b45e6.
//
// Solidity: function _poolRegisted() view returns(bool)
func (_PosPool *PosPoolCaller) PoolRegisted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "_poolRegisted")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// PoolRegisted is a free data retrieval call binding the contract method 0x4b1b45e6.
//
// Solidity: function _poolRegisted() view returns(bool)
func (_PosPool *PosPoolBulkCaller) PoolRegisted(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "_poolRegisted")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "_poolRegisted")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolRegisted is a free data retrieval call binding the contract method 0x4b1b45e6.
//
// Solidity: function _poolRegisted() view returns(bool)
func (_PosPool *PosPoolSession) PoolRegisted() (bool, error) {
	return _PosPool.Contract.PoolRegisted(&_PosPool.CallOpts)
}

// PoolRegisted is a free data retrieval call binding the contract method 0x4b1b45e6.
//
// Solidity: function _poolRegisted() view returns(bool)
func (_PosPool *PosPoolCallerSession) PoolRegisted() (bool, error) {
	return _PosPool.Contract.PoolRegisted(&_PosPool.CallOpts)
}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPool *PosPoolCaller) PoolUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "_poolUnlockPeriod")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) PoolUnlockPeriod(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "_poolUnlockPeriod")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "_poolUnlockPeriod")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPool *PosPoolSession) PoolUnlockPeriod() (*big.Int, error) {
	return _PosPool.Contract.PoolUnlockPeriod(&_PosPool.CallOpts)
}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPool *PosPoolCallerSession) PoolUnlockPeriod() (*big.Int, error) {
	return _PosPool.Contract.PoolUnlockPeriod(&_PosPool.CallOpts)
}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address _user) view returns(uint256)
func (_PosPool *PosPoolCaller) UserShareRatio(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "_userShareRatio", _user)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address _user) view returns(uint256)
func (_PosPool *PosPoolBulkCaller) UserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _user common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "_userShareRatio", _user)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "_userShareRatio")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address _user) view returns(uint256)
func (_PosPool *PosPoolSession) UserShareRatio(_user common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserShareRatio(&_PosPool.CallOpts, _user)
}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address _user) view returns(uint256)
func (_PosPool *PosPoolCallerSession) UserShareRatio(_user common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserShareRatio(&_PosPool.CallOpts, _user)
}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPool *PosPoolCaller) AccRewardPerCfx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "accRewardPerCfx")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) AccRewardPerCfx(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "accRewardPerCfx")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "accRewardPerCfx")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPool *PosPoolSession) AccRewardPerCfx() (*big.Int, error) {
	return _PosPool.Contract.AccRewardPerCfx(&_PosPool.CallOpts)
}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPool *PosPoolCallerSession) AccRewardPerCfx() (*big.Int, error) {
	return _PosPool.Contract.AccRewardPerCfx(&_PosPool.CallOpts)
}

// AllUserShareRatio is a free data retrieval call binding the contract method 0x611888b1.
//
// Solidity: function allUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolCaller) AllUserShareRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "allUserShareRatio")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// AllUserShareRatio is a free data retrieval call binding the contract method 0x611888b1.
//
// Solidity: function allUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) AllUserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "allUserShareRatio")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "allUserShareRatio")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// AllUserShareRatio is a free data retrieval call binding the contract method 0x611888b1.
//
// Solidity: function allUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolSession) AllUserShareRatio() (*big.Int, error) {
	return _PosPool.Contract.AllUserShareRatio(&_PosPool.CallOpts)
}

// AllUserShareRatio is a free data retrieval call binding the contract method 0x611888b1.
//
// Solidity: function allUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolCallerSession) AllUserShareRatio() (*big.Int, error) {
	return _PosPool.Contract.AllUserShareRatio(&_PosPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPool *PosPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPool *PosPoolBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "owner")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPool *PosPoolSession) Owner() (common.Address, error) {
	return _PosPool.Contract.Owner(&_PosPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPool *PosPoolCallerSession) Owner() (common.Address, error) {
	return _PosPool.Contract.Owner(&_PosPool.CallOpts)
}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPool *PosPoolCaller) ParamsControl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "paramsControl")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPool *PosPoolBulkCaller) ParamsControl(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "paramsControl")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "paramsControl")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPool *PosPoolSession) ParamsControl() (common.Address, error) {
	return _PosPool.Contract.ParamsControl(&_PosPool.CallOpts)
}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPool *PosPoolCallerSession) ParamsControl() (common.Address, error) {
	return _PosPool.Contract.ParamsControl(&_PosPool.CallOpts)
}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPool *PosPoolCaller) PoolAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "poolAPY")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) PoolAPY(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "poolAPY")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "poolAPY")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPool *PosPoolSession) PoolAPY() (*big.Int, error) {
	return _PosPool.Contract.PoolAPY(&_PosPool.CallOpts)
}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPool *PosPoolCallerSession) PoolAPY() (*big.Int, error) {
	return _PosPool.Contract.PoolAPY(&_PosPool.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPool *PosPoolCaller) PoolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "poolName")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPool *PosPoolBulkCaller) PoolName(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "poolName")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "poolName")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPool *PosPoolSession) PoolName() (string, error) {
	return _PosPool.Contract.PoolName(&_PosPool.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPool *PosPoolCallerSession) PoolName() (string, error) {
	return _PosPool.Contract.PoolName(&_PosPool.CallOpts)
}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCaller) PoolShot(opts *bind.CallOpts) (PoSPoolPoolShot, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "poolShot")

	if __err != nil {
		return *new(PoSPoolPoolShot), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolPoolShot)).(*PoSPoolPoolShot)

	return out0, __err

}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolBulkCaller) PoolShot(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*PoSPoolPoolShot, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "poolShot")

	out0 := new(PoSPoolPoolShot)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "poolShot")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(PoSPoolPoolShot)).(*PoSPoolPoolShot)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolSession) PoolShot() (PoSPoolPoolShot, error) {
	return _PosPool.Contract.PoolShot(&_PosPool.CallOpts)
}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCallerSession) PoolShot() (PoSPoolPoolShot, error) {
	return _PosPool.Contract.PoolShot(&_PosPool.CallOpts)
}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCaller) PoolSummary(opts *bind.CallOpts) (PoSPoolPoolSummary, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "poolSummary")

	if __err != nil {
		return *new(PoSPoolPoolSummary), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolPoolSummary)).(*PoSPoolPoolSummary)

	return out0, __err

}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolBulkCaller) PoolSummary(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*PoSPoolPoolSummary, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "poolSummary")

	out0 := new(PoSPoolPoolSummary)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "poolSummary")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(PoSPoolPoolSummary)).(*PoSPoolPoolSummary)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolSession) PoolSummary() (PoSPoolPoolSummary, error) {
	return _PosPool.Contract.PoolSummary(&_PosPool.CallOpts)
}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCallerSession) PoolSummary() (PoSPoolPoolSummary, error) {
	return _PosPool.Contract.PoolSummary(&_PosPool.CallOpts)
}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolCaller) PoolUserShareRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "poolUserShareRatio")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) PoolUserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "poolUserShareRatio")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "poolUserShareRatio")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolSession) PoolUserShareRatio() (*big.Int, error) {
	return _PosPool.Contract.PoolUserShareRatio(&_PosPool.CallOpts)
}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPool *PosPoolCallerSession) PoolUserShareRatio() (*big.Int, error) {
	return _PosPool.Contract.PoolUserShareRatio(&_PosPool.CallOpts)
}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPool *PosPoolCaller) PosAddress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "posAddress")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPool *PosPoolBulkCaller) PosAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "posAddress")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "posAddress")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPool *PosPoolSession) PosAddress() ([32]byte, error) {
	return _PosPool.Contract.PosAddress(&_PosPool.CallOpts)
}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPool *PosPoolCallerSession) PosAddress() ([32]byte, error) {
	return _PosPool.Contract.PosAddress(&_PosPool.CallOpts)
}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPool *PosPoolCaller) StakerAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "stakerAddress", i)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPool *PosPoolBulkCaller) StakerAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, i *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "stakerAddress", i)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "stakerAddress")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPool *PosPoolSession) StakerAddress(i *big.Int) (common.Address, error) {
	return _PosPool.Contract.StakerAddress(&_PosPool.CallOpts, i)
}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPool *PosPoolCallerSession) StakerAddress(i *big.Int) (common.Address, error) {
	return _PosPool.Contract.StakerAddress(&_PosPool.CallOpts, i)
}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPool *PosPoolCaller) StakerNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "stakerNumber")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPool *PosPoolBulkCaller) StakerNumber(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "stakerNumber")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "stakerNumber")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPool *PosPoolSession) StakerNumber() (*big.Int, error) {
	return _PosPool.Contract.StakerNumber(&_PosPool.CallOpts)
}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPool *PosPoolCallerSession) StakerNumber() (*big.Int, error) {
	return _PosPool.Contract.StakerNumber(&_PosPool.CallOpts)
}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCaller) UserInQueue(opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userInQueue", account, offset, limit)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolBulkCaller) UserInQueue(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userInQueue", account, offset, limit)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userInQueue")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolSession) UserInQueue(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserInQueue(&_PosPool.CallOpts, account, offset, limit)
}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCallerSession) UserInQueue(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserInQueue(&_PosPool.CallOpts, account, offset, limit)
}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCaller) UserInQueue0(opts *bind.CallOpts, account common.Address) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userInQueue0", account)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolBulkCaller) UserInQueue0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userInQueue0", account)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userInQueue0")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolSession) UserInQueue0(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserInQueue0(&_PosPool.CallOpts, account)
}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCallerSession) UserInQueue0(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserInQueue0(&_PosPool.CallOpts, account)
}

// UserInterest is a free data retrieval call binding the contract method 0xd68076c3.
//
// Solidity: function userInterest(address _address) view returns(uint256)
func (_PosPool *PosPoolCaller) UserInterest(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userInterest", _address)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// UserInterest is a free data retrieval call binding the contract method 0xd68076c3.
//
// Solidity: function userInterest(address _address) view returns(uint256)
func (_PosPool *PosPoolBulkCaller) UserInterest(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _address common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userInterest", _address)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userInterest")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserInterest is a free data retrieval call binding the contract method 0xd68076c3.
//
// Solidity: function userInterest(address _address) view returns(uint256)
func (_PosPool *PosPoolSession) UserInterest(_address common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserInterest(&_PosPool.CallOpts, _address)
}

// UserInterest is a free data retrieval call binding the contract method 0xd68076c3.
//
// Solidity: function userInterest(address _address) view returns(uint256)
func (_PosPool *PosPoolCallerSession) UserInterest(_address common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserInterest(&_PosPool.CallOpts, _address)
}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPool *PosPoolCaller) UserLockInfo(opts *bind.CallOpts, user common.Address) (IVotingEscrowLockInfo, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userLockInfo", user)

	if __err != nil {
		return *new(IVotingEscrowLockInfo), __err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingEscrowLockInfo)).(*IVotingEscrowLockInfo)

	return out0, __err

}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPool *PosPoolBulkCaller) UserLockInfo(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (*IVotingEscrowLockInfo, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userLockInfo", user)

	out0 := new(IVotingEscrowLockInfo)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userLockInfo")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(IVotingEscrowLockInfo)).(*IVotingEscrowLockInfo)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPool *PosPoolSession) UserLockInfo(user common.Address) (IVotingEscrowLockInfo, error) {
	return _PosPool.Contract.UserLockInfo(&_PosPool.CallOpts, user)
}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPool *PosPoolCallerSession) UserLockInfo(user common.Address) (IVotingEscrowLockInfo, error) {
	return _PosPool.Contract.UserLockInfo(&_PosPool.CallOpts, user)
}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCaller) UserOutQueue(opts *bind.CallOpts, account common.Address) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userOutQueue", account)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolBulkCaller) UserOutQueue(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userOutQueue", account)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userOutQueue")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolSession) UserOutQueue(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserOutQueue(&_PosPool.CallOpts, account)
}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCallerSession) UserOutQueue(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserOutQueue(&_PosPool.CallOpts, account)
}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCaller) UserOutQueue0(opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userOutQueue0", account, offset, limit)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolBulkCaller) UserOutQueue0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userOutQueue0", account, offset, limit)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userOutQueue0")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolSession) UserOutQueue0(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserOutQueue0(&_PosPool.CallOpts, account, offset, limit)
}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPool *PosPoolCallerSession) UserOutQueue0(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPool.Contract.UserOutQueue0(&_PosPool.CallOpts, account, offset, limit)
}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address _user) view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCaller) UserShot(opts *bind.CallOpts, _user common.Address) (PoSPoolUserShot, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userShot", _user)

	if __err != nil {
		return *new(PoSPoolUserShot), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolUserShot)).(*PoSPoolUserShot)

	return out0, __err

}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address _user) view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolBulkCaller) UserShot(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _user common.Address) (*PoSPoolUserShot, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userShot", _user)

	out0 := new(PoSPoolUserShot)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userShot")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(PoSPoolUserShot)).(*PoSPoolUserShot)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address _user) view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolSession) UserShot(_user common.Address) (PoSPoolUserShot, error) {
	return _PosPool.Contract.UserShot(&_PosPool.CallOpts, _user)
}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address _user) view returns((uint256,uint256,uint256))
func (_PosPool *PosPoolCallerSession) UserShot(_user common.Address) (PoSPoolUserShot, error) {
	return _PosPool.Contract.UserShot(&_PosPool.CallOpts, _user)
}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPool *PosPoolCaller) UserSummary(opts *bind.CallOpts, _user common.Address) (PoSPoolUserSummary, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userSummary", _user)

	if __err != nil {
		return *new(PoSPoolUserSummary), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolUserSummary)).(*PoSPoolUserSummary)

	return out0, __err

}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPool *PosPoolBulkCaller) UserSummary(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _user common.Address) (*PoSPoolUserSummary, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userSummary", _user)

	out0 := new(PoSPoolUserSummary)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userSummary")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(PoSPoolUserSummary)).(*PoSPoolUserSummary)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPool *PosPoolSession) UserSummary(_user common.Address) (PoSPoolUserSummary, error) {
	return _PosPool.Contract.UserSummary(&_PosPool.CallOpts, _user)
}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address _user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPool *PosPoolCallerSession) UserSummary(_user common.Address) (PoSPoolUserSummary, error) {
	return _PosPool.Contract.UserSummary(&_PosPool.CallOpts, _user)
}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPool *PosPoolCaller) UserVotePower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "userVotePower", user)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPool *PosPoolBulkCaller) UserVotePower(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "userVotePower", user)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "userVotePower")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPool *PosPoolSession) UserVotePower(user common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserVotePower(&_PosPool.CallOpts, user)
}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPool *PosPoolCallerSession) UserVotePower(user common.Address) (*big.Int, error) {
	return _PosPool.Contract.UserVotePower(&_PosPool.CallOpts, user)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPool *PosPoolCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPool.contract.Call(opts, &out, "votingEscrow")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPool *PosPoolBulkCaller) VotingEscrow(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPool.contract.GenRequest(opts, "votingEscrow")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPool.contract.DecodeOutput(&out, rawOut, "votingEscrow")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPool *PosPoolSession) VotingEscrow() (common.Address, error) {
	return _PosPool.Contract.VotingEscrow(&_PosPool.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPool *PosPoolCallerSession) VotingEscrow() (common.Address, error) {
	return _PosPool.Contract.VotingEscrow(&_PosPool.CallOpts)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x1e941601.
//
// Solidity: function _restakePosVote(uint64 votes) returns()
func (_PosPool *PosPoolTransactor) RestakePosVote(opts *bind.TransactOpts, votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "_restakePosVote", votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x1e941601.
//
// Solidity: function _restakePosVote(uint64 votes) returns()
func (_PosPool *PosPoolBulkTransactor) RestakePosVote(opts *bind.TransactOpts, votes uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "_restakePosVote", votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x1e941601.
//
// Solidity: function _restakePosVote(uint64 votes) returns()
func (_PosPool *PosPoolSession) RestakePosVote(votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RestakePosVote(&_PosPool.TransactOpts, votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x1e941601.
//
// Solidity: function _restakePosVote(uint64 votes) returns()
func (_PosPool *PosPoolTransactorSession) RestakePosVote(votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RestakePosVote(&_PosPool.TransactOpts, votes)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x86eb624a.
//
// Solidity: function _restakeUserStake(address _addr) returns()
func (_PosPool *PosPoolTransactor) RestakeUserStake(opts *bind.TransactOpts, _addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "_restakeUserStake", _addr)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x86eb624a.
//
// Solidity: function _restakeUserStake(address _addr) returns()
func (_PosPool *PosPoolBulkTransactor) RestakeUserStake(opts *bind.TransactOpts, _addr common.Address) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "_restakeUserStake", _addr)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x86eb624a.
//
// Solidity: function _restakeUserStake(address _addr) returns()
func (_PosPool *PosPoolSession) RestakeUserStake(_addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RestakeUserStake(&_PosPool.TransactOpts, _addr)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x86eb624a.
//
// Solidity: function _restakeUserStake(address _addr) returns()
func (_PosPool *PosPoolTransactorSession) RestakeUserStake(_addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RestakeUserStake(&_PosPool.TransactOpts, _addr)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0xbb61be5f.
//
// Solidity: function _retireUserStake(address _addr, uint64 endBlockNumber) returns()
func (_PosPool *PosPoolTransactor) RetireUserStake(opts *bind.TransactOpts, _addr common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "_retireUserStake", _addr, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0xbb61be5f.
//
// Solidity: function _retireUserStake(address _addr, uint64 endBlockNumber) returns()
func (_PosPool *PosPoolBulkTransactor) RetireUserStake(opts *bind.TransactOpts, _addr common.Address, endBlockNumber uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "_retireUserStake", _addr, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0xbb61be5f.
//
// Solidity: function _retireUserStake(address _addr, uint64 endBlockNumber) returns()
func (_PosPool *PosPoolSession) RetireUserStake(_addr common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RetireUserStake(&_PosPool.TransactOpts, _addr, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0xbb61be5f.
//
// Solidity: function _retireUserStake(address _addr, uint64 endBlockNumber) returns()
func (_PosPool *PosPoolTransactorSession) RetireUserStake(_addr common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RetireUserStake(&_PosPool.TransactOpts, _addr, endBlockNumber)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x2ab7107b.
//
// Solidity: function _withdrawPoolProfit(uint256 amount) returns()
func (_PosPool *PosPoolTransactor) WithdrawPoolProfit(opts *bind.TransactOpts, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "_withdrawPoolProfit", amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x2ab7107b.
//
// Solidity: function _withdrawPoolProfit(uint256 amount) returns()
func (_PosPool *PosPoolBulkTransactor) WithdrawPoolProfit(opts *bind.TransactOpts, amount *big.Int) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "_withdrawPoolProfit", amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x2ab7107b.
//
// Solidity: function _withdrawPoolProfit(uint256 amount) returns()
func (_PosPool *PosPoolSession) WithdrawPoolProfit(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.WithdrawPoolProfit(&_PosPool.TransactOpts, amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x2ab7107b.
//
// Solidity: function _withdrawPoolProfit(uint256 amount) returns()
func (_PosPool *PosPoolTransactorSession) WithdrawPoolProfit(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.WithdrawPoolProfit(&_PosPool.TransactOpts, amount)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolTransactor) AddToFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "addToFeeFreeWhiteList", _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolBulkTransactor) AddToFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "addToFeeFreeWhiteList", _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolSession) AddToFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.AddToFeeFreeWhiteList(&_PosPool.TransactOpts, _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolTransactorSession) AddToFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.AddToFeeFreeWhiteList(&_PosPool.TransactOpts, _freeAddress)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPool *PosPoolTransactor) CastVote(opts *bind.TransactOpts, vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "castVote", vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPool *PosPoolBulkTransactor) CastVote(opts *bind.TransactOpts, vote_round uint64, vote_data []ParamsControlVote) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "castVote", vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPool *PosPoolSession) CastVote(vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.CastVote(&_PosPool.TransactOpts, vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPool *PosPoolTransactorSession) CastVote(vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.CastVote(&_PosPool.TransactOpts, vote_round, vote_data)
}

// ClaimAllInterest is a paid mutator transaction binding the contract method 0x8270512f.
//
// Solidity: function claimAllInterest() returns()
func (_PosPool *PosPoolTransactor) ClaimAllInterest(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "claimAllInterest")
}

// ClaimAllInterest is a paid mutator transaction binding the contract method 0x8270512f.
//
// Solidity: function claimAllInterest() returns()
func (_PosPool *PosPoolBulkTransactor) ClaimAllInterest(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "claimAllInterest")
}

// ClaimAllInterest is a paid mutator transaction binding the contract method 0x8270512f.
//
// Solidity: function claimAllInterest() returns()
func (_PosPool *PosPoolSession) ClaimAllInterest() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.ClaimAllInterest(&_PosPool.TransactOpts)
}

// ClaimAllInterest is a paid mutator transaction binding the contract method 0x8270512f.
//
// Solidity: function claimAllInterest() returns()
func (_PosPool *PosPoolTransactorSession) ClaimAllInterest() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.ClaimAllInterest(&_PosPool.TransactOpts)
}

// ClaimInterest is a paid mutator transaction binding the contract method 0x2eb375ea.
//
// Solidity: function claimInterest(uint256 amount) returns()
func (_PosPool *PosPoolTransactor) ClaimInterest(opts *bind.TransactOpts, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "claimInterest", amount)
}

// ClaimInterest is a paid mutator transaction binding the contract method 0x2eb375ea.
//
// Solidity: function claimInterest(uint256 amount) returns()
func (_PosPool *PosPoolBulkTransactor) ClaimInterest(opts *bind.TransactOpts, amount *big.Int) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "claimInterest", amount)
}

// ClaimInterest is a paid mutator transaction binding the contract method 0x2eb375ea.
//
// Solidity: function claimInterest(uint256 amount) returns()
func (_PosPool *PosPoolSession) ClaimInterest(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.ClaimInterest(&_PosPool.TransactOpts, amount)
}

// ClaimInterest is a paid mutator transaction binding the contract method 0x2eb375ea.
//
// Solidity: function claimInterest(uint256 amount) returns()
func (_PosPool *PosPoolTransactorSession) ClaimInterest(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.ClaimInterest(&_PosPool.TransactOpts, amount)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPool *PosPoolTransactor) DecreaseStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "decreaseStake", votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPool *PosPoolBulkTransactor) DecreaseStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "decreaseStake", votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPool *PosPoolSession) DecreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.DecreaseStake(&_PosPool.TransactOpts, votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPool *PosPoolTransactorSession) DecreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.DecreaseStake(&_PosPool.TransactOpts, votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPool *PosPoolTransactor) IncreaseStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "increaseStake", votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPool *PosPoolBulkTransactor) IncreaseStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "increaseStake", votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPool *PosPoolSession) IncreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.IncreaseStake(&_PosPool.TransactOpts, votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPool *PosPoolTransactorSession) IncreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.IncreaseStake(&_PosPool.TransactOpts, votePower)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPool *PosPoolTransactor) Initialize(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPool *PosPoolBulkTransactor) Initialize(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPool *PosPoolSession) Initialize() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.Initialize(&_PosPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPool *PosPoolTransactorSession) Initialize() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.Initialize(&_PosPool.TransactOpts)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPool *PosPoolTransactor) LockForVotePower(opts *bind.TransactOpts, amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "lockForVotePower", amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPool *PosPoolBulkTransactor) LockForVotePower(opts *bind.TransactOpts, amount *big.Int, unlockBlockNumber *big.Int) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "lockForVotePower", amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPool *PosPoolSession) LockForVotePower(amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.LockForVotePower(&_PosPool.TransactOpts, amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPool *PosPoolTransactorSession) LockForVotePower(amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.LockForVotePower(&_PosPool.TransactOpts, amount, unlockBlockNumber)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 indentifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPool *PosPoolTransactor) Register(opts *bind.TransactOpts, indentifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "register", indentifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 indentifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPool *PosPoolBulkTransactor) Register(opts *bind.TransactOpts, indentifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "register", indentifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 indentifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPool *PosPoolSession) Register(indentifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.Register(&_PosPool.TransactOpts, indentifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 indentifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPool *PosPoolTransactorSession) Register(indentifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.Register(&_PosPool.TransactOpts, indentifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolTransactor) RemoveFromFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "removeFromFeeFreeWhiteList", _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolBulkTransactor) RemoveFromFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "removeFromFeeFreeWhiteList", _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolSession) RemoveFromFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RemoveFromFeeFreeWhiteList(&_PosPool.TransactOpts, _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPool *PosPoolTransactorSession) RemoveFromFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RemoveFromFeeFreeWhiteList(&_PosPool.TransactOpts, _freeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPool *PosPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPool *PosPoolBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPool *PosPoolSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RenounceOwnership(&_PosPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPool *PosPoolTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.RenounceOwnership(&_PosPool.TransactOpts)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPool *PosPoolTransactor) SetCfxCountOfOneVote(opts *bind.TransactOpts, count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setCfxCountOfOneVote", count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPool *PosPoolBulkTransactor) SetCfxCountOfOneVote(opts *bind.TransactOpts, count *big.Int) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setCfxCountOfOneVote", count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPool *PosPoolSession) SetCfxCountOfOneVote(count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetCfxCountOfOneVote(&_PosPool.TransactOpts, count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPool *PosPoolTransactorSession) SetCfxCountOfOneVote(count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetCfxCountOfOneVote(&_PosPool.TransactOpts, count)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPool *PosPoolTransactor) SetLockPeriod(opts *bind.TransactOpts, period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setLockPeriod", period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPool *PosPoolBulkTransactor) SetLockPeriod(opts *bind.TransactOpts, period uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setLockPeriod", period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPool *PosPoolSession) SetLockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetLockPeriod(&_PosPool.TransactOpts, period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPool *PosPoolTransactorSession) SetLockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetLockPeriod(&_PosPool.TransactOpts, period)
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPool *PosPoolTransactor) SetParamsControl(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setParamsControl")
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPool *PosPoolBulkTransactor) SetParamsControl(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setParamsControl")
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPool *PosPoolSession) SetParamsControl() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetParamsControl(&_PosPool.TransactOpts)
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPool *PosPoolTransactorSession) SetParamsControl() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetParamsControl(&_PosPool.TransactOpts)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPool *PosPoolTransactor) SetPoolName(opts *bind.TransactOpts, name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setPoolName", name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPool *PosPoolBulkTransactor) SetPoolName(opts *bind.TransactOpts, name string) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setPoolName", name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPool *PosPoolSession) SetPoolName(name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetPoolName(&_PosPool.TransactOpts, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPool *PosPoolTransactorSession) SetPoolName(name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetPoolName(&_PosPool.TransactOpts, name)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPool *PosPoolTransactor) SetPoolUserShareRatio(opts *bind.TransactOpts, ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setPoolUserShareRatio", ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPool *PosPoolBulkTransactor) SetPoolUserShareRatio(opts *bind.TransactOpts, ratio uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setPoolUserShareRatio", ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPool *PosPoolSession) SetPoolUserShareRatio(ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetPoolUserShareRatio(&_PosPool.TransactOpts, ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPool *PosPoolTransactorSession) SetPoolUserShareRatio(ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetPoolUserShareRatio(&_PosPool.TransactOpts, ratio)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPool *PosPoolTransactor) SetUnlockPeriod(opts *bind.TransactOpts, period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setUnlockPeriod", period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPool *PosPoolBulkTransactor) SetUnlockPeriod(opts *bind.TransactOpts, period uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setUnlockPeriod", period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPool *PosPoolSession) SetUnlockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetUnlockPeriod(&_PosPool.TransactOpts, period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPool *PosPoolTransactorSession) SetUnlockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetUnlockPeriod(&_PosPool.TransactOpts, period)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPool *PosPoolTransactor) SetVotingEscrow(opts *bind.TransactOpts, _votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "setVotingEscrow", _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPool *PosPoolBulkTransactor) SetVotingEscrow(opts *bind.TransactOpts, _votingEscrow common.Address) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "setVotingEscrow", _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPool *PosPoolSession) SetVotingEscrow(_votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetVotingEscrow(&_PosPool.TransactOpts, _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPool *PosPoolTransactorSession) SetVotingEscrow(_votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.SetVotingEscrow(&_PosPool.TransactOpts, _votingEscrow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPool *PosPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPool *PosPoolBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPool *PosPoolSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.TransferOwnership(&_PosPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPool *PosPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.TransferOwnership(&_PosPool.TransactOpts, newOwner)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPool *PosPoolTransactor) WithdrawStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.contract.Transact(opts, "withdrawStake", votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPool *PosPoolBulkTransactor) WithdrawStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPool.contract.GenUnsignedTransaction(opts, "withdrawStake", votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPool *PosPoolSession) WithdrawStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.WithdrawStake(&_PosPool.TransactOpts, votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPool *PosPoolTransactorSession) WithdrawStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPool.Contract.WithdrawStake(&_PosPool.TransactOpts, votePower)
}

// PosPoolClaimInterestIterator is returned from FilterClaimInterest and is used to iterate over the raw logs and unpacked data for ClaimInterest events raised by the PosPool contract.
type PosPoolClaimInterestIterator struct {
	Event *PosPoolClaimInterest // Event containing the contract specifics and raw log

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
func (it *PosPoolClaimInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolClaimInterest)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolClaimInterest)
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
func (it *PosPoolClaimInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolClaimInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolClaimInterest represents a ClaimInterest event raised by the PosPool contract.
type PosPoolClaimInterest struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// PosPoolClaimInterestOrChainReorg represents a ClaimInterest subscription event raised by the PosPool contract.
type PosPoolClaimInterestOrChainReorg struct {
	Event      *PosPoolClaimInterest
	ChainReorg *types.ChainReorg
}

// FilterClaimInterest is a free log retrieval operation binding the contract event 0xbbf701f67a7d285bbd2b9b2dc7a5fee474e6dc3c059183385da6cc9de2be34a9.
//
// Solidity: event ClaimInterest(address indexed user, uint256 amount)
func (_PosPool *PosPoolFilterer) FilterClaimInterest(opts *bind.FilterOpts, user []common.Address) (*PosPoolClaimInterestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPool.contract.FilterLogs(opts, "ClaimInterest", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolClaimInterestIterator{contract: _PosPool.contract, event: "ClaimInterest", logs: logs}, nil
}

// WatchClaimInterest is a free log subscription operation binding the contract event 0xbbf701f67a7d285bbd2b9b2dc7a5fee474e6dc3c059183385da6cc9de2be34a9.
//
// Solidity: event ClaimInterest(address indexed user, uint256 amount)
func (_PosPool *PosPoolFilterer) WatchClaimInterest(opts *bind.WatchOpts, sink chan<- *PosPoolClaimInterestOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "ClaimInterest", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolClaimInterestOrChainReorg)
				event.Event = new(PosPoolClaimInterest)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "ClaimInterest", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseClaimInterest is a log parse operation binding the contract event 0xbbf701f67a7d285bbd2b9b2dc7a5fee474e6dc3c059183385da6cc9de2be34a9.
//
// Solidity: event ClaimInterest(address indexed user, uint256 amount)
func (_PosPool *PosPoolFilterer) ParseClaimInterest(log types.Log) (*PosPoolClaimInterest, error) {
	event := new(PosPoolClaimInterest)
	if err := _PosPool.contract.UnpackLog(event, "ClaimInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDecreasePoSStakeIterator is returned from FilterDecreasePoSStake and is used to iterate over the raw logs and unpacked data for DecreasePoSStake events raised by the PosPool contract.
type PosPoolDecreasePoSStakeIterator struct {
	Event *PosPoolDecreasePoSStake // Event containing the contract specifics and raw log

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
func (it *PosPoolDecreasePoSStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDecreasePoSStake)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolDecreasePoSStake)
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
func (it *PosPoolDecreasePoSStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDecreasePoSStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDecreasePoSStake represents a DecreasePoSStake event raised by the PosPool contract.
type PosPoolDecreasePoSStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolDecreasePoSStakeOrChainReorg represents a DecreasePoSStake subscription event raised by the PosPool contract.
type PosPoolDecreasePoSStakeOrChainReorg struct {
	Event      *PosPoolDecreasePoSStake
	ChainReorg *types.ChainReorg
}

// FilterDecreasePoSStake is a free log retrieval operation binding the contract event 0xd81524ca5b038be9f98b286808a6f46f99c8700bce9ab4793dc12a2dbaa4f2d8.
//
// Solidity: event DecreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) FilterDecreasePoSStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolDecreasePoSStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPool.contract.FilterLogs(opts, "DecreasePoSStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDecreasePoSStakeIterator{contract: _PosPool.contract, event: "DecreasePoSStake", logs: logs}, nil
}

// WatchDecreasePoSStake is a free log subscription operation binding the contract event 0xd81524ca5b038be9f98b286808a6f46f99c8700bce9ab4793dc12a2dbaa4f2d8.
//
// Solidity: event DecreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) WatchDecreasePoSStake(opts *bind.WatchOpts, sink chan<- *PosPoolDecreasePoSStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "DecreasePoSStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDecreasePoSStakeOrChainReorg)
				event.Event = new(PosPoolDecreasePoSStake)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "DecreasePoSStake", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseDecreasePoSStake is a log parse operation binding the contract event 0xd81524ca5b038be9f98b286808a6f46f99c8700bce9ab4793dc12a2dbaa4f2d8.
//
// Solidity: event DecreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) ParseDecreasePoSStake(log types.Log) (*PosPoolDecreasePoSStake, error) {
	event := new(PosPoolDecreasePoSStake)
	if err := _PosPool.contract.UnpackLog(event, "DecreasePoSStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolIncreasePoSStakeIterator is returned from FilterIncreasePoSStake and is used to iterate over the raw logs and unpacked data for IncreasePoSStake events raised by the PosPool contract.
type PosPoolIncreasePoSStakeIterator struct {
	Event *PosPoolIncreasePoSStake // Event containing the contract specifics and raw log

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
func (it *PosPoolIncreasePoSStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolIncreasePoSStake)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolIncreasePoSStake)
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
func (it *PosPoolIncreasePoSStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolIncreasePoSStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolIncreasePoSStake represents a IncreasePoSStake event raised by the PosPool contract.
type PosPoolIncreasePoSStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolIncreasePoSStakeOrChainReorg represents a IncreasePoSStake subscription event raised by the PosPool contract.
type PosPoolIncreasePoSStakeOrChainReorg struct {
	Event      *PosPoolIncreasePoSStake
	ChainReorg *types.ChainReorg
}

// FilterIncreasePoSStake is a free log retrieval operation binding the contract event 0xf5ebfcde71acb15a758233f0c9ab25632209399b90a9209a4f2379ec5fcee89f.
//
// Solidity: event IncreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) FilterIncreasePoSStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolIncreasePoSStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPool.contract.FilterLogs(opts, "IncreasePoSStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolIncreasePoSStakeIterator{contract: _PosPool.contract, event: "IncreasePoSStake", logs: logs}, nil
}

// WatchIncreasePoSStake is a free log subscription operation binding the contract event 0xf5ebfcde71acb15a758233f0c9ab25632209399b90a9209a4f2379ec5fcee89f.
//
// Solidity: event IncreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) WatchIncreasePoSStake(opts *bind.WatchOpts, sink chan<- *PosPoolIncreasePoSStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "IncreasePoSStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolIncreasePoSStakeOrChainReorg)
				event.Event = new(PosPoolIncreasePoSStake)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "IncreasePoSStake", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseIncreasePoSStake is a log parse operation binding the contract event 0xf5ebfcde71acb15a758233f0c9ab25632209399b90a9209a4f2379ec5fcee89f.
//
// Solidity: event IncreasePoSStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) ParseIncreasePoSStake(log types.Log) (*PosPoolIncreasePoSStake, error) {
	event := new(PosPoolIncreasePoSStake)
	if err := _PosPool.contract.UnpackLog(event, "IncreasePoSStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PosPool contract.
type PosPoolOwnershipTransferredIterator struct {
	Event *PosPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PosPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolOwnershipTransferred)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolOwnershipTransferred)
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
func (it *PosPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolOwnershipTransferred represents a OwnershipTransferred event raised by the PosPool contract.
type PosPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// PosPoolOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the PosPool contract.
type PosPoolOwnershipTransferredOrChainReorg struct {
	Event      *PosPoolOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PosPool *PosPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PosPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _PosPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolOwnershipTransferredIterator{contract: _PosPool.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PosPool *PosPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PosPoolOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolOwnershipTransferredOrChainReorg)
				event.Event = new(PosPoolOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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
func (_PosPool *PosPoolFilterer) ParseOwnershipTransferred(log types.Log) (*PosPoolOwnershipTransferred, error) {
	event := new(PosPoolOwnershipTransferred)
	if err := _PosPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolRatioChangedIterator is returned from FilterRatioChanged and is used to iterate over the raw logs and unpacked data for RatioChanged events raised by the PosPool contract.
type PosPoolRatioChangedIterator struct {
	Event *PosPoolRatioChanged // Event containing the contract specifics and raw log

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
func (it *PosPoolRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolRatioChanged)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolRatioChanged)
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
func (it *PosPoolRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolRatioChanged represents a RatioChanged event raised by the PosPool contract.
type PosPoolRatioChanged struct {
	Ratio *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// PosPoolRatioChangedOrChainReorg represents a RatioChanged subscription event raised by the PosPool contract.
type PosPoolRatioChangedOrChainReorg struct {
	Event      *PosPoolRatioChanged
	ChainReorg *types.ChainReorg
}

// FilterRatioChanged is a free log retrieval operation binding the contract event 0xe200219383a2dbe79fad4e1549a81b6057429299f33583641c5fdddcdc0b3797.
//
// Solidity: event RatioChanged(uint256 ratio)
func (_PosPool *PosPoolFilterer) FilterRatioChanged(opts *bind.FilterOpts) (*PosPoolRatioChangedIterator, error) {

	logs, err := _PosPool.contract.FilterLogs(opts, "RatioChanged")
	if err != nil {
		return nil, err
	}
	return &PosPoolRatioChangedIterator{contract: _PosPool.contract, event: "RatioChanged", logs: logs}, nil
}

// WatchRatioChanged is a free log subscription operation binding the contract event 0xe200219383a2dbe79fad4e1549a81b6057429299f33583641c5fdddcdc0b3797.
//
// Solidity: event RatioChanged(uint256 ratio)
func (_PosPool *PosPoolFilterer) WatchRatioChanged(opts *bind.WatchOpts, sink chan<- *PosPoolRatioChangedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "RatioChanged")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolRatioChangedOrChainReorg)
				event.Event = new(PosPoolRatioChanged)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "RatioChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseRatioChanged is a log parse operation binding the contract event 0xe200219383a2dbe79fad4e1549a81b6057429299f33583641c5fdddcdc0b3797.
//
// Solidity: event RatioChanged(uint256 ratio)
func (_PosPool *PosPoolFilterer) ParseRatioChanged(log types.Log) (*PosPoolRatioChanged, error) {
	event := new(PosPoolRatioChanged)
	if err := _PosPool.contract.UnpackLog(event, "RatioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolWithdrawStakeIterator is returned from FilterWithdrawStake and is used to iterate over the raw logs and unpacked data for WithdrawStake events raised by the PosPool contract.
type PosPoolWithdrawStakeIterator struct {
	Event *PosPoolWithdrawStake // Event containing the contract specifics and raw log

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
func (it *PosPoolWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolWithdrawStake)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PosPoolWithdrawStake)
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
func (it *PosPoolWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolWithdrawStake represents a WithdrawStake event raised by the PosPool contract.
type PosPoolWithdrawStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolWithdrawStakeOrChainReorg represents a WithdrawStake subscription event raised by the PosPool contract.
type PosPoolWithdrawStakeOrChainReorg struct {
	Event      *PosPoolWithdrawStake
	ChainReorg *types.ChainReorg
}

// FilterWithdrawStake is a free log retrieval operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) FilterWithdrawStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolWithdrawStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPool.contract.FilterLogs(opts, "WithdrawStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolWithdrawStakeIterator{contract: _PosPool.contract, event: "WithdrawStake", logs: logs}, nil
}

// WatchWithdrawStake is a free log subscription operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) WatchWithdrawStake(opts *bind.WatchOpts, sink chan<- *PosPoolWithdrawStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPool.contract.WatchLogs(opts, "WithdrawStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolWithdrawStakeOrChainReorg)
				event.Event = new(PosPoolWithdrawStake)

				if log.ChainReorg == nil {
					if err := _PosPool.contract.UnpackLog(event.Event, "WithdrawStake", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseWithdrawStake is a log parse operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address indexed user, uint256 votePower)
func (_PosPool *PosPoolFilterer) ParseWithdrawStake(log types.Log) (*PosPoolWithdrawStake, error) {
	event := new(PosPoolWithdrawStake)
	if err := _PosPool.contract.UnpackLog(event, "WithdrawStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
