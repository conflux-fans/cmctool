// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pospooldx

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
	Available   *big.Int
	Reward      *big.Int
	TotalReward *big.Int
}

// PoSPoolUserShot is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolUserShot struct {
	Available       *big.Int
	AccRewardPerCfx *big.Int
	BlockNumber     *big.Int
}

// PoSPoolUserSummary is an auto generated low-level Go binding around an user-defined struct.
type PoSPoolUserSummary struct {
	Votes         *big.Int
	Available     *big.Int
	Locked        *big.Int
	Unlocked      *big.Int
	ClaimedReward *big.Int
	CurrentReward *big.Int
}

// VotePowerQueueQueueNode is an auto generated low-level Go binding around an user-defined struct.
type VotePowerQueueQueueNode struct {
	VotePower *big.Int
	EndBlock  *big.Int
}

// PosPoolDxABI is the input ABI used to generate the binding from.
const PosPoolDxABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"DecreaseStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"IncreaseStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"}],\"name\":\"WithdrawStake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"_userShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accRewardPerCfx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_freeAddress\",\"type\":\"address\"}],\"name\":\"addToFeeFreeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"internalType\":\"structParamsControl.Vote[]\",\"name\":\"vote_data\",\"type\":\"tuple[]\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"decreaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getUserRewardInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlockNumber\",\"type\":\"uint256\"}],\"name\":\"lockForVotePower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramsControl\",\"outputs\":[{\"internalType\":\"contractParamsControl\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolShot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.PoolShot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.PoolSummary\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolUserShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"posAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[2]\",\"name\":\"blsPubKeyProof\",\"type\":\"bytes[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_freeAddress\",\"type\":\"address\"}],\"name\":\"removeFromFeeFreeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votes\",\"type\":\"uint64\"}],\"name\":\"restakePosVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"restakeUserStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"}],\"name\":\"retireUserStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"setCfxCountOfOneVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"name\":\"setLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setParamsControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setPoolName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ratio\",\"type\":\"uint64\"}],\"name\":\"setPoolUserShareRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"name\":\"setUnlockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_votingEscrow\",\"type\":\"address\"}],\"name\":\"setVotingEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"stakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"userInQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userInQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIVotingEscrow.LockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userOutQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"userOutQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structVotePowerQueue.QueueNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserShareRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userShot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerCfx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.UserShot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"locked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentReward\",\"type\":\"uint256\"}],\"internalType\":\"structPoSPool.UserSummary\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userVotePower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEscrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPoolProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PosPoolDx is an auto generated Go binding around an Conflux contract.
type PosPoolDx struct {
	PosPoolDxCaller     // Read-only binding to the contract
	PosPoolDxTransactor // Write-only binding to the contract
	PosPoolDxFilterer   // Log filterer for contract events
}

// PosPoolDxCaller is an auto generated read-only Go binding around an Conflux contract.
type PosPoolDxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolDxBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type PosPoolDxBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolDxTransactor is an auto generated write-only Go binding around an Conflux contract.
type PosPoolDxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolDxBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type PosPoolDxBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolDxFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type PosPoolDxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PosPoolDxSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type PosPoolDxSession struct {
	Contract     *PosPoolDx        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PosPoolDxCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type PosPoolDxCallerSession struct {
	Contract *PosPoolDxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PosPoolDxTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type PosPoolDxTransactorSession struct {
	Contract     *PosPoolDxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PosPoolDxRaw is an auto generated low-level Go binding around an Conflux contract.
type PosPoolDxRaw struct {
	Contract *PosPoolDx // Generic contract binding to access the raw methods on
}

// PosPoolDxCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type PosPoolDxCallerRaw struct {
	Contract *PosPoolDxCaller // Generic read-only contract binding to access the raw methods on
}

// PosPoolDxTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type PosPoolDxTransactorRaw struct {
	Contract *PosPoolDxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPosPoolDx creates a new instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDx(address types.Address, backend bind.ContractBackend) (*PosPoolDx, error) {
	contract, err := bindPosPoolDx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PosPoolDx{PosPoolDxCaller: PosPoolDxCaller{contract: contract}, PosPoolDxTransactor: PosPoolDxTransactor{contract: contract}, PosPoolDxFilterer: PosPoolDxFilterer{contract: contract}}, nil
}

// NewPosPoolDxCaller creates a new read-only instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDxCaller(address types.Address, caller bind.ContractCaller) (*PosPoolDxCaller, error) {
	contract, err := bindPosPoolDx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxCaller{contract: contract}, nil
}

// NewPosPoolDxTransactor creates a new write-only instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDxTransactor(address types.Address, transactor bind.ContractTransactor) (*PosPoolDxTransactor, error) {
	contract, err := bindPosPoolDx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxTransactor{contract: contract}, nil
}

// NewPosPoolDxFilterer creates a new log filterer instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDxFilterer(address types.Address, filterer bind.ContractFilterer) (*PosPoolDxFilterer, error) {
	contract, err := bindPosPoolDx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxFilterer{contract: contract}, nil
}

// NewPosPoolDxCaller creates a new read-only instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDxBulkCaller(address types.Address, caller bind.ContractCaller) (*PosPoolDxBulkCaller, error) {
	contract, err := bindPosPoolDx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxBulkCaller{contract: contract}, nil
}

// NewPosPoolDxBulkTransactor creates a new write-only instance of PosPoolDx, bound to a specific deployed contract.
func NewPosPoolDxBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*PosPoolDxBulkTransactor, error) {
	contract, err := bindPosPoolDx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxBulkTransactor{contract: contract}, nil
}

// bindPosPoolDx binds a generic wrapper to an already deployed contract.
func bindPosPoolDx(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PosPoolDxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PosPoolDx *PosPoolDxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PosPoolDx.Contract.PosPoolDxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PosPoolDx *PosPoolDxRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.PosPoolDxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PosPoolDx *PosPoolDxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.PosPoolDxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PosPoolDx *PosPoolDxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PosPoolDx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PosPoolDx *PosPoolDxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PosPoolDx *PosPoolDxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPoolDx *PosPoolDxCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "VERSION")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPoolDx *PosPoolDxBulkCaller) VERSION(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "VERSION")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "VERSION")
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
func (_PosPoolDx *PosPoolDxSession) VERSION() (string, error) {
	return _PosPoolDx.Contract.VERSION(&_PosPoolDx.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PosPoolDx *PosPoolDxCallerSession) VERSION() (string, error) {
	return _PosPoolDx.Contract.VERSION(&_PosPoolDx.CallOpts)
}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) PoolLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "_poolLockPeriod")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolLockPeriod(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "_poolLockPeriod")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "_poolLockPeriod")
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
func (_PosPoolDx *PosPoolDxSession) PoolLockPeriod() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolLockPeriod(&_PosPoolDx.CallOpts)
}

// PoolLockPeriod is a free data retrieval call binding the contract method 0xb4064a25.
//
// Solidity: function _poolLockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) PoolLockPeriod() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolLockPeriod(&_PosPoolDx.CallOpts)
}

// PoolRegistered is a free data retrieval call binding the contract method 0x71086584.
//
// Solidity: function _poolRegistered() view returns(bool)
func (_PosPoolDx *PosPoolDxCaller) PoolRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "_poolRegistered")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// PoolRegistered is a free data retrieval call binding the contract method 0x71086584.
//
// Solidity: function _poolRegistered() view returns(bool)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolRegistered(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "_poolRegistered")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "_poolRegistered")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// PoolRegistered is a free data retrieval call binding the contract method 0x71086584.
//
// Solidity: function _poolRegistered() view returns(bool)
func (_PosPoolDx *PosPoolDxSession) PoolRegistered() (bool, error) {
	return _PosPoolDx.Contract.PoolRegistered(&_PosPoolDx.CallOpts)
}

// PoolRegistered is a free data retrieval call binding the contract method 0x71086584.
//
// Solidity: function _poolRegistered() view returns(bool)
func (_PosPoolDx *PosPoolDxCallerSession) PoolRegistered() (bool, error) {
	return _PosPoolDx.Contract.PoolRegistered(&_PosPoolDx.CallOpts)
}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) PoolUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "_poolUnlockPeriod")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolUnlockPeriod(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "_poolUnlockPeriod")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "_poolUnlockPeriod")
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
func (_PosPoolDx *PosPoolDxSession) PoolUnlockPeriod() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolUnlockPeriod(&_PosPoolDx.CallOpts)
}

// PoolUnlockPeriod is a free data retrieval call binding the contract method 0x9fc95e03.
//
// Solidity: function _poolUnlockPeriod() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) PoolUnlockPeriod() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolUnlockPeriod(&_PosPoolDx.CallOpts)
}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) UserShareRatio(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "_userShareRatio", user)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) UserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "_userShareRatio", user)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "_userShareRatio")
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
// Solidity: function _userShareRatio(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxSession) UserShareRatio(user common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.UserShareRatio(&_PosPoolDx.CallOpts, user)
}

// UserShareRatio is a free data retrieval call binding the contract method 0xbd0b2dad.
//
// Solidity: function _userShareRatio(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) UserShareRatio(user common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.UserShareRatio(&_PosPoolDx.CallOpts, user)
}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) AccRewardPerCfx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "accRewardPerCfx")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) AccRewardPerCfx(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "accRewardPerCfx")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "accRewardPerCfx")
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
func (_PosPoolDx *PosPoolDxSession) AccRewardPerCfx() (*big.Int, error) {
	return _PosPoolDx.Contract.AccRewardPerCfx(&_PosPoolDx.CallOpts)
}

// AccRewardPerCfx is a free data retrieval call binding the contract method 0xf660776b.
//
// Solidity: function accRewardPerCfx() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) AccRewardPerCfx() (*big.Int, error) {
	return _PosPoolDx.Contract.AccRewardPerCfx(&_PosPoolDx.CallOpts)
}

// GetUserReward is a free data retrieval call binding the contract method 0xc75ebb82.
//
// Solidity: function getUserReward(address depositor) view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) GetUserReward(opts *bind.CallOpts, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "getUserReward", depositor)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// GetUserReward is a free data retrieval call binding the contract method 0xc75ebb82.
//
// Solidity: function getUserReward(address depositor) view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) GetUserReward(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, depositor common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "getUserReward", depositor)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "getUserReward")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetUserReward is a free data retrieval call binding the contract method 0xc75ebb82.
//
// Solidity: function getUserReward(address depositor) view returns(uint256)
func (_PosPoolDx *PosPoolDxSession) GetUserReward(depositor common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.GetUserReward(&_PosPoolDx.CallOpts, depositor)
}

// GetUserReward is a free data retrieval call binding the contract method 0xc75ebb82.
//
// Solidity: function getUserReward(address depositor) view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) GetUserReward(depositor common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.GetUserReward(&_PosPoolDx.CallOpts, depositor)
}

// GetUserRewardInfo is a free data retrieval call binding the contract method 0x69517310.
//
// Solidity: function getUserRewardInfo(address depositor) view returns(uint256, uint256, uint256)
func (_PosPoolDx *PosPoolDxCaller) GetUserRewardInfo(opts *bind.CallOpts, depositor common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "getUserRewardInfo", depositor)

	if __err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, __err

}

// GetUserRewardInfo is a free data retrieval call binding the contract method 0x69517310.
//
// Solidity: function getUserRewardInfo(address depositor) view returns(uint256, uint256, uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) GetUserRewardInfo(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, depositor common.Address) (**big.Int, **big.Int, **big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "getUserRewardInfo", depositor)

	out0 := new(*big.Int)
	out1 := new(*big.Int)
	out2 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "getUserRewardInfo")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		*out1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
		*out2 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, out1, out2, __err

}

// GetUserRewardInfo is a free data retrieval call binding the contract method 0x69517310.
//
// Solidity: function getUserRewardInfo(address depositor) view returns(uint256, uint256, uint256)
func (_PosPoolDx *PosPoolDxSession) GetUserRewardInfo(depositor common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _PosPoolDx.Contract.GetUserRewardInfo(&_PosPoolDx.CallOpts, depositor)
}

// GetUserRewardInfo is a free data retrieval call binding the contract method 0x69517310.
//
// Solidity: function getUserRewardInfo(address depositor) view returns(uint256, uint256, uint256)
func (_PosPoolDx *PosPoolDxCallerSession) GetUserRewardInfo(depositor common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _PosPoolDx.Contract.GetUserRewardInfo(&_PosPoolDx.CallOpts, depositor)
}

// GetUserShareRatio is a free data retrieval call binding the contract method 0xf4a49874.
//
// Solidity: function getUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) GetUserShareRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "getUserShareRatio")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// GetUserShareRatio is a free data retrieval call binding the contract method 0xf4a49874.
//
// Solidity: function getUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) GetUserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "getUserShareRatio")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "getUserShareRatio")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetUserShareRatio is a free data retrieval call binding the contract method 0xf4a49874.
//
// Solidity: function getUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxSession) GetUserShareRatio() (*big.Int, error) {
	return _PosPoolDx.Contract.GetUserShareRatio(&_PosPoolDx.CallOpts)
}

// GetUserShareRatio is a free data retrieval call binding the contract method 0xf4a49874.
//
// Solidity: function getUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) GetUserShareRatio() (*big.Int, error) {
	return _PosPoolDx.Contract.GetUserShareRatio(&_PosPoolDx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPoolDx *PosPoolDxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPoolDx *PosPoolDxBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "owner")
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
func (_PosPoolDx *PosPoolDxSession) Owner() (common.Address, error) {
	return _PosPoolDx.Contract.Owner(&_PosPoolDx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PosPoolDx *PosPoolDxCallerSession) Owner() (common.Address, error) {
	return _PosPoolDx.Contract.Owner(&_PosPoolDx.CallOpts)
}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPoolDx *PosPoolDxCaller) ParamsControl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "paramsControl")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPoolDx *PosPoolDxBulkCaller) ParamsControl(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "paramsControl")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "paramsControl")
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
func (_PosPoolDx *PosPoolDxSession) ParamsControl() (common.Address, error) {
	return _PosPoolDx.Contract.ParamsControl(&_PosPoolDx.CallOpts)
}

// ParamsControl is a free data retrieval call binding the contract method 0x52bf4d5e.
//
// Solidity: function paramsControl() view returns(address)
func (_PosPoolDx *PosPoolDxCallerSession) ParamsControl() (common.Address, error) {
	return _PosPoolDx.Contract.ParamsControl(&_PosPoolDx.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PosPoolDx *PosPoolDxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "paused")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PosPoolDx *PosPoolDxBulkCaller) Paused(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "paused")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "paused")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PosPoolDx *PosPoolDxSession) Paused() (bool, error) {
	return _PosPoolDx.Contract.Paused(&_PosPoolDx.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PosPoolDx *PosPoolDxCallerSession) Paused() (bool, error) {
	return _PosPoolDx.Contract.Paused(&_PosPoolDx.CallOpts)
}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) PoolAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "poolAPY")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolAPY(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "poolAPY")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "poolAPY")
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
func (_PosPoolDx *PosPoolDxSession) PoolAPY() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolAPY(&_PosPoolDx.CallOpts)
}

// PoolAPY is a free data retrieval call binding the contract method 0x82df7a3b.
//
// Solidity: function poolAPY() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) PoolAPY() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolAPY(&_PosPoolDx.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPoolDx *PosPoolDxCaller) PoolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "poolName")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolName(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "poolName")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "poolName")
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
func (_PosPoolDx *PosPoolDxSession) PoolName() (string, error) {
	return _PosPoolDx.Contract.PoolName(&_PosPoolDx.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_PosPoolDx *PosPoolDxCallerSession) PoolName() (string, error) {
	return _PosPoolDx.Contract.PoolName(&_PosPoolDx.CallOpts)
}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCaller) PoolShot(opts *bind.CallOpts) (PoSPoolPoolShot, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "poolShot")

	if __err != nil {
		return *new(PoSPoolPoolShot), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolPoolShot)).(*PoSPoolPoolShot)

	return out0, __err

}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxBulkCaller) PoolShot(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*PoSPoolPoolShot, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "poolShot")

	out0 := new(PoSPoolPoolShot)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "poolShot")
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
func (_PosPoolDx *PosPoolDxSession) PoolShot() (PoSPoolPoolShot, error) {
	return _PosPoolDx.Contract.PoolShot(&_PosPoolDx.CallOpts)
}

// PoolShot is a free data retrieval call binding the contract method 0x7be47601.
//
// Solidity: function poolShot() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCallerSession) PoolShot() (PoSPoolPoolShot, error) {
	return _PosPoolDx.Contract.PoolShot(&_PosPoolDx.CallOpts)
}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCaller) PoolSummary(opts *bind.CallOpts) (PoSPoolPoolSummary, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "poolSummary")

	if __err != nil {
		return *new(PoSPoolPoolSummary), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolPoolSummary)).(*PoSPoolPoolSummary)

	return out0, __err

}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxBulkCaller) PoolSummary(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*PoSPoolPoolSummary, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "poolSummary")

	out0 := new(PoSPoolPoolSummary)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "poolSummary")
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
func (_PosPoolDx *PosPoolDxSession) PoolSummary() (PoSPoolPoolSummary, error) {
	return _PosPoolDx.Contract.PoolSummary(&_PosPoolDx.CallOpts)
}

// PoolSummary is a free data retrieval call binding the contract method 0x7571fcf6.
//
// Solidity: function poolSummary() view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCallerSession) PoolSummary() (PoSPoolPoolSummary, error) {
	return _PosPoolDx.Contract.PoolSummary(&_PosPoolDx.CallOpts)
}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) PoolUserShareRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "poolUserShareRatio")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) PoolUserShareRatio(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "poolUserShareRatio")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "poolUserShareRatio")
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
func (_PosPoolDx *PosPoolDxSession) PoolUserShareRatio() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolUserShareRatio(&_PosPoolDx.CallOpts)
}

// PoolUserShareRatio is a free data retrieval call binding the contract method 0x67a1aa4f.
//
// Solidity: function poolUserShareRatio() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) PoolUserShareRatio() (*big.Int, error) {
	return _PosPoolDx.Contract.PoolUserShareRatio(&_PosPoolDx.CallOpts)
}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPoolDx *PosPoolDxCaller) PosAddress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "posAddress")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPoolDx *PosPoolDxBulkCaller) PosAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "posAddress")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "posAddress")
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
func (_PosPoolDx *PosPoolDxSession) PosAddress() ([32]byte, error) {
	return _PosPoolDx.Contract.PosAddress(&_PosPoolDx.CallOpts)
}

// PosAddress is a free data retrieval call binding the contract method 0x1764618c.
//
// Solidity: function posAddress() view returns(bytes32)
func (_PosPoolDx *PosPoolDxCallerSession) PosAddress() ([32]byte, error) {
	return _PosPoolDx.Contract.PosAddress(&_PosPoolDx.CallOpts)
}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPoolDx *PosPoolDxCaller) StakerAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "stakerAddress", i)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPoolDx *PosPoolDxBulkCaller) StakerAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, i *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "stakerAddress", i)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "stakerAddress")
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
func (_PosPoolDx *PosPoolDxSession) StakerAddress(i *big.Int) (common.Address, error) {
	return _PosPoolDx.Contract.StakerAddress(&_PosPoolDx.CallOpts, i)
}

// StakerAddress is a free data retrieval call binding the contract method 0x94067045.
//
// Solidity: function stakerAddress(uint256 i) view returns(address)
func (_PosPoolDx *PosPoolDxCallerSession) StakerAddress(i *big.Int) (common.Address, error) {
	return _PosPoolDx.Contract.StakerAddress(&_PosPoolDx.CallOpts, i)
}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) StakerNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "stakerNumber")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) StakerNumber(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "stakerNumber")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "stakerNumber")
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
func (_PosPoolDx *PosPoolDxSession) StakerNumber() (*big.Int, error) {
	return _PosPoolDx.Contract.StakerNumber(&_PosPoolDx.CallOpts)
}

// StakerNumber is a free data retrieval call binding the contract method 0x9aafd5b8.
//
// Solidity: function stakerNumber() view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) StakerNumber() (*big.Int, error) {
	return _PosPoolDx.Contract.StakerNumber(&_PosPoolDx.CallOpts)
}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCaller) UserInQueue(opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userInQueue", account, offset, limit)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxBulkCaller) UserInQueue(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userInQueue", account, offset, limit)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userInQueue")
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
func (_PosPoolDx *PosPoolDxSession) UserInQueue(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserInQueue(&_PosPoolDx.CallOpts, account, offset, limit)
}

// UserInQueue is a free data retrieval call binding the contract method 0xa6698002.
//
// Solidity: function userInQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCallerSession) UserInQueue(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserInQueue(&_PosPoolDx.CallOpts, account, offset, limit)
}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCaller) UserInQueue0(opts *bind.CallOpts, account common.Address) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userInQueue0", account)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxBulkCaller) UserInQueue0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userInQueue0", account)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userInQueue0")
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
func (_PosPoolDx *PosPoolDxSession) UserInQueue0(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserInQueue0(&_PosPoolDx.CallOpts, account)
}

// UserInQueue0 is a free data retrieval call binding the contract method 0xe2af7928.
//
// Solidity: function userInQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCallerSession) UserInQueue0(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserInQueue0(&_PosPoolDx.CallOpts, account)
}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPoolDx *PosPoolDxCaller) UserLockInfo(opts *bind.CallOpts, user common.Address) (IVotingEscrowLockInfo, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userLockInfo", user)

	if __err != nil {
		return *new(IVotingEscrowLockInfo), __err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingEscrowLockInfo)).(*IVotingEscrowLockInfo)

	return out0, __err

}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPoolDx *PosPoolDxBulkCaller) UserLockInfo(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (*IVotingEscrowLockInfo, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userLockInfo", user)

	out0 := new(IVotingEscrowLockInfo)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userLockInfo")
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
func (_PosPoolDx *PosPoolDxSession) UserLockInfo(user common.Address) (IVotingEscrowLockInfo, error) {
	return _PosPoolDx.Contract.UserLockInfo(&_PosPoolDx.CallOpts, user)
}

// UserLockInfo is a free data retrieval call binding the contract method 0x4a42fc89.
//
// Solidity: function userLockInfo(address user) view returns((uint256,uint256))
func (_PosPoolDx *PosPoolDxCallerSession) UserLockInfo(user common.Address) (IVotingEscrowLockInfo, error) {
	return _PosPoolDx.Contract.UserLockInfo(&_PosPoolDx.CallOpts, user)
}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCaller) UserOutQueue(opts *bind.CallOpts, account common.Address) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userOutQueue", account)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxBulkCaller) UserOutQueue(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userOutQueue", account)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userOutQueue")
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
func (_PosPoolDx *PosPoolDxSession) UserOutQueue(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserOutQueue(&_PosPoolDx.CallOpts, account)
}

// UserOutQueue is a free data retrieval call binding the contract method 0x24a3bbe2.
//
// Solidity: function userOutQueue(address account) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCallerSession) UserOutQueue(account common.Address) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserOutQueue(&_PosPoolDx.CallOpts, account)
}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCaller) UserOutQueue0(opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userOutQueue0", account, offset, limit)

	if __err != nil {
		return *new([]VotePowerQueueQueueNode), __err
	}

	out0 := *abi.ConvertType(out[0], new([]VotePowerQueueQueueNode)).(*[]VotePowerQueueQueueNode)

	return out0, __err

}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxBulkCaller) UserOutQueue0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, offset uint64, limit uint64) (*[]VotePowerQueueQueueNode, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userOutQueue0", account, offset, limit)

	out0 := new([]VotePowerQueueQueueNode)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userOutQueue0")
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
func (_PosPoolDx *PosPoolDxSession) UserOutQueue0(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserOutQueue0(&_PosPoolDx.CallOpts, account, offset, limit)
}

// UserOutQueue0 is a free data retrieval call binding the contract method 0x336f8e0e.
//
// Solidity: function userOutQueue(address account, uint64 offset, uint64 limit) view returns((uint256,uint256)[])
func (_PosPoolDx *PosPoolDxCallerSession) UserOutQueue0(account common.Address, offset uint64, limit uint64) ([]VotePowerQueueQueueNode, error) {
	return _PosPoolDx.Contract.UserOutQueue0(&_PosPoolDx.CallOpts, account, offset, limit)
}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address user) view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCaller) UserShot(opts *bind.CallOpts, user common.Address) (PoSPoolUserShot, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userShot", user)

	if __err != nil {
		return *new(PoSPoolUserShot), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolUserShot)).(*PoSPoolUserShot)

	return out0, __err

}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address user) view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxBulkCaller) UserShot(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (*PoSPoolUserShot, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userShot", user)

	out0 := new(PoSPoolUserShot)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userShot")
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
// Solidity: function userShot(address user) view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxSession) UserShot(user common.Address) (PoSPoolUserShot, error) {
	return _PosPoolDx.Contract.UserShot(&_PosPoolDx.CallOpts, user)
}

// UserShot is a free data retrieval call binding the contract method 0x9d46992b.
//
// Solidity: function userShot(address user) view returns((uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCallerSession) UserShot(user common.Address) (PoSPoolUserShot, error) {
	return _PosPoolDx.Contract.UserShot(&_PosPoolDx.CallOpts, user)
}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCaller) UserSummary(opts *bind.CallOpts, user common.Address) (PoSPoolUserSummary, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userSummary", user)

	if __err != nil {
		return *new(PoSPoolUserSummary), __err
	}

	out0 := *abi.ConvertType(out[0], new(PoSPoolUserSummary)).(*PoSPoolUserSummary)

	return out0, __err

}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxBulkCaller) UserSummary(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (*PoSPoolUserSummary, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userSummary", user)

	out0 := new(PoSPoolUserSummary)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userSummary")
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
// Solidity: function userSummary(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxSession) UserSummary(user common.Address) (PoSPoolUserSummary, error) {
	return _PosPoolDx.Contract.UserSummary(&_PosPoolDx.CallOpts, user)
}

// UserSummary is a free data retrieval call binding the contract method 0xe3d5b233.
//
// Solidity: function userSummary(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_PosPoolDx *PosPoolDxCallerSession) UserSummary(user common.Address) (PoSPoolUserSummary, error) {
	return _PosPoolDx.Contract.UserSummary(&_PosPoolDx.CallOpts, user)
}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxCaller) UserVotePower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "userVotePower", user)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxBulkCaller) UserVotePower(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, user common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "userVotePower", user)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "userVotePower")
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
func (_PosPoolDx *PosPoolDxSession) UserVotePower(user common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.UserVotePower(&_PosPoolDx.CallOpts, user)
}

// UserVotePower is a free data retrieval call binding the contract method 0x3cd5c824.
//
// Solidity: function userVotePower(address user) view returns(uint256)
func (_PosPoolDx *PosPoolDxCallerSession) UserVotePower(user common.Address) (*big.Int, error) {
	return _PosPoolDx.Contract.UserVotePower(&_PosPoolDx.CallOpts, user)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPoolDx *PosPoolDxCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _PosPoolDx.contract.Call(opts, &out, "votingEscrow")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPoolDx *PosPoolDxBulkCaller) VotingEscrow(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _PosPoolDx.contract.GenRequest(opts, "votingEscrow")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _PosPoolDx.contract.DecodeOutput(&out, rawOut, "votingEscrow")
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
func (_PosPoolDx *PosPoolDxSession) VotingEscrow() (common.Address, error) {
	return _PosPoolDx.Contract.VotingEscrow(&_PosPoolDx.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_PosPoolDx *PosPoolDxCallerSession) VotingEscrow() (common.Address, error) {
	return _PosPoolDx.Contract.VotingEscrow(&_PosPoolDx.CallOpts)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxTransactor) AddToFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "addToFeeFreeWhiteList", _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxBulkTransactor) AddToFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "addToFeeFreeWhiteList", _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxSession) AddToFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.AddToFeeFreeWhiteList(&_PosPoolDx.TransactOpts, _freeAddress)
}

// AddToFeeFreeWhiteList is a paid mutator transaction binding the contract method 0x0e589c0a.
//
// Solidity: function addToFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxTransactorSession) AddToFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.AddToFeeFreeWhiteList(&_PosPoolDx.TransactOpts, _freeAddress)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPoolDx *PosPoolDxTransactor) CastVote(opts *bind.TransactOpts, vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "castVote", vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) CastVote(opts *bind.TransactOpts, vote_round uint64, vote_data []ParamsControlVote) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "castVote", vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPoolDx *PosPoolDxSession) CastVote(vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.CastVote(&_PosPoolDx.TransactOpts, vote_round, vote_data)
}

// CastVote is a paid mutator transaction binding the contract method 0x0bf64c90.
//
// Solidity: function castVote(uint64 vote_round, (uint16,uint256[3])[] vote_data) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) CastVote(vote_round uint64, vote_data []ParamsControlVote) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.CastVote(&_PosPoolDx.TransactOpts, vote_round, vote_data)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x8a623d86.
//
// Solidity: function claimAllReward() returns()
func (_PosPoolDx *PosPoolDxTransactor) ClaimAllReward(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "claimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x8a623d86.
//
// Solidity: function claimAllReward() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) ClaimAllReward(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "claimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x8a623d86.
//
// Solidity: function claimAllReward() returns()
func (_PosPoolDx *PosPoolDxSession) ClaimAllReward() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.ClaimAllReward(&_PosPoolDx.TransactOpts)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x8a623d86.
//
// Solidity: function claimAllReward() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) ClaimAllReward() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.ClaimAllReward(&_PosPoolDx.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxTransactor) ClaimReward(opts *bind.TransactOpts, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "claimReward", amount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) ClaimReward(opts *bind.TransactOpts, amount *big.Int) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "claimReward", amount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxSession) ClaimReward(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.ClaimReward(&_PosPoolDx.TransactOpts, amount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) ClaimReward(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.ClaimReward(&_PosPoolDx.TransactOpts, amount)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxTransactor) DecreaseStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "decreaseStake", votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) DecreaseStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "decreaseStake", votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxSession) DecreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.DecreaseStake(&_PosPoolDx.TransactOpts, votePower)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0xbeb99c1e.
//
// Solidity: function decreaseStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) DecreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.DecreaseStake(&_PosPoolDx.TransactOpts, votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPoolDx *PosPoolDxTransactor) IncreaseStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "increaseStake", votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) IncreaseStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "increaseStake", votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPoolDx *PosPoolDxSession) IncreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.IncreaseStake(&_PosPoolDx.TransactOpts, votePower)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x09fecf7f.
//
// Solidity: function increaseStake(uint64 votePower) payable returns()
func (_PosPoolDx *PosPoolDxTransactorSession) IncreaseStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.IncreaseStake(&_PosPoolDx.TransactOpts, votePower)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPoolDx *PosPoolDxTransactor) Initialize(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) Initialize(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPoolDx *PosPoolDxSession) Initialize() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.Initialize(&_PosPoolDx.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) Initialize() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.Initialize(&_PosPoolDx.TransactOpts)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPoolDx *PosPoolDxTransactor) LockForVotePower(opts *bind.TransactOpts, amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "lockForVotePower", amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) LockForVotePower(opts *bind.TransactOpts, amount *big.Int, unlockBlockNumber *big.Int) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "lockForVotePower", amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPoolDx *PosPoolDxSession) LockForVotePower(amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.LockForVotePower(&_PosPoolDx.TransactOpts, amount, unlockBlockNumber)
}

// LockForVotePower is a paid mutator transaction binding the contract method 0x9d1b07ec.
//
// Solidity: function lockForVotePower(uint256 amount, uint256 unlockBlockNumber) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) LockForVotePower(amount *big.Int, unlockBlockNumber *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.LockForVotePower(&_PosPoolDx.TransactOpts, amount, unlockBlockNumber)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_PosPoolDx *PosPoolDxTransactor) PauseContract(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) PauseContract(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_PosPoolDx *PosPoolDxSession) PauseContract() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.PauseContract(&_PosPoolDx.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) PauseContract() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.PauseContract(&_PosPoolDx.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 identifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPoolDx *PosPoolDxTransactor) Register(opts *bind.TransactOpts, identifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "register", identifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 identifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) Register(opts *bind.TransactOpts, identifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "register", identifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 identifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPoolDx *PosPoolDxSession) Register(identifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.Register(&_PosPoolDx.TransactOpts, identifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// Register is a paid mutator transaction binding the contract method 0xe335b451.
//
// Solidity: function register(bytes32 identifier, uint64 votePower, bytes blsPubKey, bytes vrfPubKey, bytes[2] blsPubKeyProof) payable returns()
func (_PosPoolDx *PosPoolDxTransactorSession) Register(identifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.Register(&_PosPoolDx.TransactOpts, identifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxTransactor) RemoveFromFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "removeFromFeeFreeWhiteList", _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxBulkTransactor) RemoveFromFeeFreeWhiteList(opts *bind.TransactOpts, _freeAddress common.Address) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "removeFromFeeFreeWhiteList", _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxSession) RemoveFromFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RemoveFromFeeFreeWhiteList(&_PosPoolDx.TransactOpts, _freeAddress)
}

// RemoveFromFeeFreeWhiteList is a paid mutator transaction binding the contract method 0xada9344f.
//
// Solidity: function removeFromFeeFreeWhiteList(address _freeAddress) returns(bool)
func (_PosPoolDx *PosPoolDxTransactorSession) RemoveFromFeeFreeWhiteList(_freeAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RemoveFromFeeFreeWhiteList(&_PosPoolDx.TransactOpts, _freeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPoolDx *PosPoolDxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPoolDx *PosPoolDxSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RenounceOwnership(&_PosPoolDx.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RenounceOwnership(&_PosPoolDx.TransactOpts)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x3a7f64c9.
//
// Solidity: function restakePosVote(uint64 votes) returns()
func (_PosPoolDx *PosPoolDxTransactor) RestakePosVote(opts *bind.TransactOpts, votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "restakePosVote", votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x3a7f64c9.
//
// Solidity: function restakePosVote(uint64 votes) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) RestakePosVote(opts *bind.TransactOpts, votes uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "restakePosVote", votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x3a7f64c9.
//
// Solidity: function restakePosVote(uint64 votes) returns()
func (_PosPoolDx *PosPoolDxSession) RestakePosVote(votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RestakePosVote(&_PosPoolDx.TransactOpts, votes)
}

// RestakePosVote is a paid mutator transaction binding the contract method 0x3a7f64c9.
//
// Solidity: function restakePosVote(uint64 votes) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) RestakePosVote(votes uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RestakePosVote(&_PosPoolDx.TransactOpts, votes)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x43d02a41.
//
// Solidity: function restakeUserStake(address depositor) returns()
func (_PosPoolDx *PosPoolDxTransactor) RestakeUserStake(opts *bind.TransactOpts, depositor common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "restakeUserStake", depositor)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x43d02a41.
//
// Solidity: function restakeUserStake(address depositor) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) RestakeUserStake(opts *bind.TransactOpts, depositor common.Address) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "restakeUserStake", depositor)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x43d02a41.
//
// Solidity: function restakeUserStake(address depositor) returns()
func (_PosPoolDx *PosPoolDxSession) RestakeUserStake(depositor common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RestakeUserStake(&_PosPoolDx.TransactOpts, depositor)
}

// RestakeUserStake is a paid mutator transaction binding the contract method 0x43d02a41.
//
// Solidity: function restakeUserStake(address depositor) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) RestakeUserStake(depositor common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RestakeUserStake(&_PosPoolDx.TransactOpts, depositor)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0x290e155f.
//
// Solidity: function retireUserStake(address depositor, uint64 endBlockNumber) returns()
func (_PosPoolDx *PosPoolDxTransactor) RetireUserStake(opts *bind.TransactOpts, depositor common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "retireUserStake", depositor, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0x290e155f.
//
// Solidity: function retireUserStake(address depositor, uint64 endBlockNumber) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) RetireUserStake(opts *bind.TransactOpts, depositor common.Address, endBlockNumber uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "retireUserStake", depositor, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0x290e155f.
//
// Solidity: function retireUserStake(address depositor, uint64 endBlockNumber) returns()
func (_PosPoolDx *PosPoolDxSession) RetireUserStake(depositor common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RetireUserStake(&_PosPoolDx.TransactOpts, depositor, endBlockNumber)
}

// RetireUserStake is a paid mutator transaction binding the contract method 0x290e155f.
//
// Solidity: function retireUserStake(address depositor, uint64 endBlockNumber) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) RetireUserStake(depositor common.Address, endBlockNumber uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.RetireUserStake(&_PosPoolDx.TransactOpts, depositor, endBlockNumber)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetCfxCountOfOneVote(opts *bind.TransactOpts, count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setCfxCountOfOneVote", count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetCfxCountOfOneVote(opts *bind.TransactOpts, count *big.Int) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setCfxCountOfOneVote", count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPoolDx *PosPoolDxSession) SetCfxCountOfOneVote(count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetCfxCountOfOneVote(&_PosPoolDx.TransactOpts, count)
}

// SetCfxCountOfOneVote is a paid mutator transaction binding the contract method 0x96b6b59d.
//
// Solidity: function setCfxCountOfOneVote(uint256 count) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetCfxCountOfOneVote(count *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetCfxCountOfOneVote(&_PosPoolDx.TransactOpts, count)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetLockPeriod(opts *bind.TransactOpts, period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setLockPeriod", period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetLockPeriod(opts *bind.TransactOpts, period uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setLockPeriod", period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxSession) SetLockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetLockPeriod(&_PosPoolDx.TransactOpts, period)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0xe5e20ccf.
//
// Solidity: function setLockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetLockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetLockPeriod(&_PosPoolDx.TransactOpts, period)
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPoolDx *PosPoolDxTransactor) SetParamsControl(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setParamsControl")
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetParamsControl(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setParamsControl")
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPoolDx *PosPoolDxSession) SetParamsControl() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetParamsControl(&_PosPoolDx.TransactOpts)
}

// SetParamsControl is a paid mutator transaction binding the contract method 0xa5fd7f23.
//
// Solidity: function setParamsControl() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetParamsControl() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetParamsControl(&_PosPoolDx.TransactOpts)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetPoolName(opts *bind.TransactOpts, name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setPoolName", name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetPoolName(opts *bind.TransactOpts, name string) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setPoolName", name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPoolDx *PosPoolDxSession) SetPoolName(name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetPoolName(&_PosPoolDx.TransactOpts, name)
}

// SetPoolName is a paid mutator transaction binding the contract method 0x72b65424.
//
// Solidity: function setPoolName(string name) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetPoolName(name string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetPoolName(&_PosPoolDx.TransactOpts, name)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetPoolUserShareRatio(opts *bind.TransactOpts, ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setPoolUserShareRatio", ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetPoolUserShareRatio(opts *bind.TransactOpts, ratio uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setPoolUserShareRatio", ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPoolDx *PosPoolDxSession) SetPoolUserShareRatio(ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetPoolUserShareRatio(&_PosPoolDx.TransactOpts, ratio)
}

// SetPoolUserShareRatio is a paid mutator transaction binding the contract method 0xc3b8eba0.
//
// Solidity: function setPoolUserShareRatio(uint64 ratio) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetPoolUserShareRatio(ratio uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetPoolUserShareRatio(&_PosPoolDx.TransactOpts, ratio)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetUnlockPeriod(opts *bind.TransactOpts, period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setUnlockPeriod", period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetUnlockPeriod(opts *bind.TransactOpts, period uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setUnlockPeriod", period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxSession) SetUnlockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetUnlockPeriod(&_PosPoolDx.TransactOpts, period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0xbe84d35d.
//
// Solidity: function setUnlockPeriod(uint64 period) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetUnlockPeriod(period uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetUnlockPeriod(&_PosPoolDx.TransactOpts, period)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPoolDx *PosPoolDxTransactor) SetVotingEscrow(opts *bind.TransactOpts, _votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "setVotingEscrow", _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) SetVotingEscrow(opts *bind.TransactOpts, _votingEscrow common.Address) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "setVotingEscrow", _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPoolDx *PosPoolDxSession) SetVotingEscrow(_votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetVotingEscrow(&_PosPoolDx.TransactOpts, _votingEscrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0xce5ec92e.
//
// Solidity: function setVotingEscrow(address _votingEscrow) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) SetVotingEscrow(_votingEscrow common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.SetVotingEscrow(&_PosPoolDx.TransactOpts, _votingEscrow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPoolDx *PosPoolDxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPoolDx *PosPoolDxSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.TransferOwnership(&_PosPoolDx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.TransferOwnership(&_PosPoolDx.TransactOpts, newOwner)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_PosPoolDx *PosPoolDxTransactor) UnpauseContract(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "unpauseContract")
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) UnpauseContract(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "unpauseContract")
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_PosPoolDx *PosPoolDxSession) UnpauseContract() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.UnpauseContract(&_PosPoolDx.TransactOpts)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_PosPoolDx *PosPoolDxTransactorSession) UnpauseContract() (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.UnpauseContract(&_PosPoolDx.TransactOpts)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x6858a497.
//
// Solidity: function withdrawPoolProfit(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxTransactor) WithdrawPoolProfit(opts *bind.TransactOpts, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "withdrawPoolProfit", amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x6858a497.
//
// Solidity: function withdrawPoolProfit(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) WithdrawPoolProfit(opts *bind.TransactOpts, amount *big.Int) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "withdrawPoolProfit", amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x6858a497.
//
// Solidity: function withdrawPoolProfit(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxSession) WithdrawPoolProfit(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.WithdrawPoolProfit(&_PosPoolDx.TransactOpts, amount)
}

// WithdrawPoolProfit is a paid mutator transaction binding the contract method 0x6858a497.
//
// Solidity: function withdrawPoolProfit(uint256 amount) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) WithdrawPoolProfit(amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.WithdrawPoolProfit(&_PosPoolDx.TransactOpts, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxTransactor) WithdrawStake(opts *bind.TransactOpts, votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.contract.Transact(opts, "withdrawStake", votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxBulkTransactor) WithdrawStake(opts *bind.TransactOpts, votePower uint64) types.UnsignedTransaction {
	return _PosPoolDx.contract.GenUnsignedTransaction(opts, "withdrawStake", votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxSession) WithdrawStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.WithdrawStake(&_PosPoolDx.TransactOpts, votePower)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x2ce7777f.
//
// Solidity: function withdrawStake(uint64 votePower) returns()
func (_PosPoolDx *PosPoolDxTransactorSession) WithdrawStake(votePower uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _PosPoolDx.Contract.WithdrawStake(&_PosPoolDx.TransactOpts, votePower)
}

// PosPoolDxClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the PosPoolDx contract.
type PosPoolDxClaimRewardIterator struct {
	Event *PosPoolDxClaimReward // Event containing the contract specifics and raw log

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
func (it *PosPoolDxClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxClaimReward)
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
		it.Event = new(PosPoolDxClaimReward)
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
func (it *PosPoolDxClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxClaimReward represents a ClaimReward event raised by the PosPoolDx contract.
type PosPoolDxClaimReward struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// PosPoolDxClaimRewardOrChainReorg represents a ClaimReward subscription event raised by the PosPoolDx contract.
type PosPoolDxClaimRewardOrChainReorg struct {
	Event      *PosPoolDxClaimReward
	ChainReorg *types.ChainReorg
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_PosPoolDx *PosPoolDxFilterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address) (*PosPoolDxClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxClaimRewardIterator{contract: _PosPoolDx.contract, event: "ClaimReward", logs: logs}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_PosPoolDx *PosPoolDxFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *PosPoolDxClaimRewardOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxClaimRewardOrChainReorg)
				event.Event = new(PosPoolDxClaimReward)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "ClaimReward", *log.Log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_PosPoolDx *PosPoolDxFilterer) ParseClaimReward(log types.Log) (*PosPoolDxClaimReward, error) {
	event := new(PosPoolDxClaimReward)
	if err := _PosPoolDx.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxDecreaseStakeIterator is returned from FilterDecreaseStake and is used to iterate over the raw logs and unpacked data for DecreaseStake events raised by the PosPoolDx contract.
type PosPoolDxDecreaseStakeIterator struct {
	Event *PosPoolDxDecreaseStake // Event containing the contract specifics and raw log

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
func (it *PosPoolDxDecreaseStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxDecreaseStake)
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
		it.Event = new(PosPoolDxDecreaseStake)
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
func (it *PosPoolDxDecreaseStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxDecreaseStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxDecreaseStake represents a DecreaseStake event raised by the PosPoolDx contract.
type PosPoolDxDecreaseStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolDxDecreaseStakeOrChainReorg represents a DecreaseStake subscription event raised by the PosPoolDx contract.
type PosPoolDxDecreaseStakeOrChainReorg struct {
	Event      *PosPoolDxDecreaseStake
	ChainReorg *types.ChainReorg
}

// FilterDecreaseStake is a free log retrieval operation binding the contract event 0xde64b14a21e5111d183f8e908405a25086d072032964325008132e0569b16346.
//
// Solidity: event DecreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) FilterDecreaseStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolDxDecreaseStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "DecreaseStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxDecreaseStakeIterator{contract: _PosPoolDx.contract, event: "DecreaseStake", logs: logs}, nil
}

// WatchDecreaseStake is a free log subscription operation binding the contract event 0xde64b14a21e5111d183f8e908405a25086d072032964325008132e0569b16346.
//
// Solidity: event DecreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) WatchDecreaseStake(opts *bind.WatchOpts, sink chan<- *PosPoolDxDecreaseStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "DecreaseStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxDecreaseStakeOrChainReorg)
				event.Event = new(PosPoolDxDecreaseStake)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "DecreaseStake", *log.Log); err != nil {
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

// ParseDecreaseStake is a log parse operation binding the contract event 0xde64b14a21e5111d183f8e908405a25086d072032964325008132e0569b16346.
//
// Solidity: event DecreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) ParseDecreaseStake(log types.Log) (*PosPoolDxDecreaseStake, error) {
	event := new(PosPoolDxDecreaseStake)
	if err := _PosPoolDx.contract.UnpackLog(event, "DecreaseStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxIncreaseStakeIterator is returned from FilterIncreaseStake and is used to iterate over the raw logs and unpacked data for IncreaseStake events raised by the PosPoolDx contract.
type PosPoolDxIncreaseStakeIterator struct {
	Event *PosPoolDxIncreaseStake // Event containing the contract specifics and raw log

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
func (it *PosPoolDxIncreaseStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxIncreaseStake)
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
		it.Event = new(PosPoolDxIncreaseStake)
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
func (it *PosPoolDxIncreaseStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxIncreaseStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxIncreaseStake represents a IncreaseStake event raised by the PosPoolDx contract.
type PosPoolDxIncreaseStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolDxIncreaseStakeOrChainReorg represents a IncreaseStake subscription event raised by the PosPoolDx contract.
type PosPoolDxIncreaseStakeOrChainReorg struct {
	Event      *PosPoolDxIncreaseStake
	ChainReorg *types.ChainReorg
}

// FilterIncreaseStake is a free log retrieval operation binding the contract event 0xbffd49014f772f01f8219427603e630b60bbbf2faf3c706da305dc9b6b0b3aba.
//
// Solidity: event IncreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) FilterIncreaseStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolDxIncreaseStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "IncreaseStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxIncreaseStakeIterator{contract: _PosPoolDx.contract, event: "IncreaseStake", logs: logs}, nil
}

// WatchIncreaseStake is a free log subscription operation binding the contract event 0xbffd49014f772f01f8219427603e630b60bbbf2faf3c706da305dc9b6b0b3aba.
//
// Solidity: event IncreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) WatchIncreaseStake(opts *bind.WatchOpts, sink chan<- *PosPoolDxIncreaseStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "IncreaseStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxIncreaseStakeOrChainReorg)
				event.Event = new(PosPoolDxIncreaseStake)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "IncreaseStake", *log.Log); err != nil {
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

// ParseIncreaseStake is a log parse operation binding the contract event 0xbffd49014f772f01f8219427603e630b60bbbf2faf3c706da305dc9b6b0b3aba.
//
// Solidity: event IncreaseStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) ParseIncreaseStake(log types.Log) (*PosPoolDxIncreaseStake, error) {
	event := new(PosPoolDxIncreaseStake)
	if err := _PosPoolDx.contract.UnpackLog(event, "IncreaseStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PosPoolDx contract.
type PosPoolDxInitializedIterator struct {
	Event *PosPoolDxInitialized // Event containing the contract specifics and raw log

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
func (it *PosPoolDxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxInitialized)
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
		it.Event = new(PosPoolDxInitialized)
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
func (it *PosPoolDxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxInitialized represents a Initialized event raised by the PosPoolDx contract.
type PosPoolDxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// PosPoolDxInitializedOrChainReorg represents a Initialized subscription event raised by the PosPoolDx contract.
type PosPoolDxInitializedOrChainReorg struct {
	Event      *PosPoolDxInitialized
	ChainReorg *types.ChainReorg
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PosPoolDx *PosPoolDxFilterer) FilterInitialized(opts *bind.FilterOpts) (*PosPoolDxInitializedIterator, error) {

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PosPoolDxInitializedIterator{contract: _PosPoolDx.contract, event: "Initialized", logs: logs}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PosPoolDx *PosPoolDxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PosPoolDxInitializedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxInitializedOrChainReorg)
				event.Event = new(PosPoolDxInitialized)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "Initialized", *log.Log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PosPoolDx *PosPoolDxFilterer) ParseInitialized(log types.Log) (*PosPoolDxInitialized, error) {
	event := new(PosPoolDxInitialized)
	if err := _PosPoolDx.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PosPoolDx contract.
type PosPoolDxOwnershipTransferredIterator struct {
	Event *PosPoolDxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PosPoolDxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxOwnershipTransferred)
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
		it.Event = new(PosPoolDxOwnershipTransferred)
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
func (it *PosPoolDxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxOwnershipTransferred represents a OwnershipTransferred event raised by the PosPoolDx contract.
type PosPoolDxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// PosPoolDxOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the PosPoolDx contract.
type PosPoolDxOwnershipTransferredOrChainReorg struct {
	Event      *PosPoolDxOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PosPoolDx *PosPoolDxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PosPoolDxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxOwnershipTransferredIterator{contract: _PosPoolDx.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PosPoolDx *PosPoolDxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PosPoolDxOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxOwnershipTransferredOrChainReorg)
				event.Event = new(PosPoolDxOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
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
func (_PosPoolDx *PosPoolDxFilterer) ParseOwnershipTransferred(log types.Log) (*PosPoolDxOwnershipTransferred, error) {
	event := new(PosPoolDxOwnershipTransferred)
	if err := _PosPoolDx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PosPoolDx contract.
type PosPoolDxPausedIterator struct {
	Event *PosPoolDxPaused // Event containing the contract specifics and raw log

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
func (it *PosPoolDxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxPaused)
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
		it.Event = new(PosPoolDxPaused)
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
func (it *PosPoolDxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxPaused represents a Paused event raised by the PosPoolDx contract.
type PosPoolDxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// PosPoolDxPausedOrChainReorg represents a Paused subscription event raised by the PosPoolDx contract.
type PosPoolDxPausedOrChainReorg struct {
	Event      *PosPoolDxPaused
	ChainReorg *types.ChainReorg
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PosPoolDx *PosPoolDxFilterer) FilterPaused(opts *bind.FilterOpts) (*PosPoolDxPausedIterator, error) {

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PosPoolDxPausedIterator{contract: _PosPoolDx.contract, event: "Paused", logs: logs}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PosPoolDx *PosPoolDxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PosPoolDxPausedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxPausedOrChainReorg)
				event.Event = new(PosPoolDxPaused)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "Paused", *log.Log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PosPoolDx *PosPoolDxFilterer) ParsePaused(log types.Log) (*PosPoolDxPaused, error) {
	event := new(PosPoolDxPaused)
	if err := _PosPoolDx.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxRatioUpdatedIterator is returned from FilterRatioUpdated and is used to iterate over the raw logs and unpacked data for RatioUpdated events raised by the PosPoolDx contract.
type PosPoolDxRatioUpdatedIterator struct {
	Event *PosPoolDxRatioUpdated // Event containing the contract specifics and raw log

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
func (it *PosPoolDxRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxRatioUpdated)
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
		it.Event = new(PosPoolDxRatioUpdated)
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
func (it *PosPoolDxRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxRatioUpdated represents a RatioUpdated event raised by the PosPoolDx contract.
type PosPoolDxRatioUpdated struct {
	Ratio *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// PosPoolDxRatioUpdatedOrChainReorg represents a RatioUpdated subscription event raised by the PosPoolDx contract.
type PosPoolDxRatioUpdatedOrChainReorg struct {
	Event      *PosPoolDxRatioUpdated
	ChainReorg *types.ChainReorg
}

// FilterRatioUpdated is a free log retrieval operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 ratio)
func (_PosPoolDx *PosPoolDxFilterer) FilterRatioUpdated(opts *bind.FilterOpts) (*PosPoolDxRatioUpdatedIterator, error) {

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "RatioUpdated")
	if err != nil {
		return nil, err
	}
	return &PosPoolDxRatioUpdatedIterator{contract: _PosPoolDx.contract, event: "RatioUpdated", logs: logs}, nil
}

// WatchRatioUpdated is a free log subscription operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 ratio)
func (_PosPoolDx *PosPoolDxFilterer) WatchRatioUpdated(opts *bind.WatchOpts, sink chan<- *PosPoolDxRatioUpdatedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "RatioUpdated")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxRatioUpdatedOrChainReorg)
				event.Event = new(PosPoolDxRatioUpdated)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "RatioUpdated", *log.Log); err != nil {
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

// ParseRatioUpdated is a log parse operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 ratio)
func (_PosPoolDx *PosPoolDxFilterer) ParseRatioUpdated(log types.Log) (*PosPoolDxRatioUpdated, error) {
	event := new(PosPoolDxRatioUpdated)
	if err := _PosPoolDx.contract.UnpackLog(event, "RatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PosPoolDx contract.
type PosPoolDxUnpausedIterator struct {
	Event *PosPoolDxUnpaused // Event containing the contract specifics and raw log

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
func (it *PosPoolDxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxUnpaused)
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
		it.Event = new(PosPoolDxUnpaused)
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
func (it *PosPoolDxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxUnpaused represents a Unpaused event raised by the PosPoolDx contract.
type PosPoolDxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// PosPoolDxUnpausedOrChainReorg represents a Unpaused subscription event raised by the PosPoolDx contract.
type PosPoolDxUnpausedOrChainReorg struct {
	Event      *PosPoolDxUnpaused
	ChainReorg *types.ChainReorg
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PosPoolDx *PosPoolDxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PosPoolDxUnpausedIterator, error) {

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PosPoolDxUnpausedIterator{contract: _PosPoolDx.contract, event: "Unpaused", logs: logs}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PosPoolDx *PosPoolDxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PosPoolDxUnpausedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxUnpausedOrChainReorg)
				event.Event = new(PosPoolDxUnpaused)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "Unpaused", *log.Log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PosPoolDx *PosPoolDxFilterer) ParseUnpaused(log types.Log) (*PosPoolDxUnpaused, error) {
	event := new(PosPoolDxUnpaused)
	if err := _PosPoolDx.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PosPoolDxWithdrawStakeIterator is returned from FilterWithdrawStake and is used to iterate over the raw logs and unpacked data for WithdrawStake events raised by the PosPoolDx contract.
type PosPoolDxWithdrawStakeIterator struct {
	Event *PosPoolDxWithdrawStake // Event containing the contract specifics and raw log

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
func (it *PosPoolDxWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PosPoolDxWithdrawStake)
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
		it.Event = new(PosPoolDxWithdrawStake)
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
func (it *PosPoolDxWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PosPoolDxWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PosPoolDxWithdrawStake represents a WithdrawStake event raised by the PosPoolDx contract.
type PosPoolDxWithdrawStake struct {
	User      common.Address
	VotePower *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// PosPoolDxWithdrawStakeOrChainReorg represents a WithdrawStake subscription event raised by the PosPoolDx contract.
type PosPoolDxWithdrawStakeOrChainReorg struct {
	Event      *PosPoolDxWithdrawStake
	ChainReorg *types.ChainReorg
}

// FilterWithdrawStake is a free log retrieval operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) FilterWithdrawStake(opts *bind.FilterOpts, user []common.Address) (*PosPoolDxWithdrawStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, err := _PosPoolDx.contract.FilterLogs(opts, "WithdrawStake", userRule)
	if err != nil {
		return nil, err
	}
	return &PosPoolDxWithdrawStakeIterator{contract: _PosPoolDx.contract, event: "WithdrawStake", logs: logs}, nil
}

// WatchWithdrawStake is a free log subscription operation binding the contract event 0x141ef67c4a6d3ec2adfb2f66d33c2b11de5b4f34344757554d430570b18a92ec.
//
// Solidity: event WithdrawStake(address indexed user, uint256 votePower)
func (_PosPoolDx *PosPoolDxFilterer) WatchWithdrawStake(opts *bind.WatchOpts, sink chan<- *PosPoolDxWithdrawStakeOrChainReorg, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PosPoolDx.contract.WatchLogs(opts, "WithdrawStake", userRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PosPoolDxWithdrawStakeOrChainReorg)
				event.Event = new(PosPoolDxWithdrawStake)

				if log.ChainReorg == nil {
					if err := _PosPoolDx.contract.UnpackLog(event.Event, "WithdrawStake", *log.Log); err != nil {
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
func (_PosPoolDx *PosPoolDxFilterer) ParseWithdrawStake(log types.Log) (*PosPoolDxWithdrawStake, error) {
	event := new(PosPoolDxWithdrawStake)
	if err := _PosPoolDx.contract.UnpackLog(event, "WithdrawStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
