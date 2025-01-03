package servicies

import (
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/internal/constants"
	"github.com/conflux-fans/cmctool/internal/contracts/pospool"
	"github.com/conflux-fans/cmctool/internal/contracts/pospooldx"
	"github.com/conflux-fans/cmctool/pkg/cfxutil"
	"github.com/conflux-fans/cmctool/pkg/timeutils"
	"github.com/conflux-fans/go-scan-sdk/client"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PosRewardFetcher struct {
	posValidatorsByScan     []*configs.PosValidatorByScan
	posValidatorsByContract []*configs.PosValidatorByContract
	scanClient              *client.Client
	cspaceClient            *sdk.Client
}

func NewPosRewardFetcher(posValidatorsByScan []*configs.PosValidatorByScan, posValidatorsByContract []*configs.PosValidatorByContract) *PosRewardFetcher {
	return &PosRewardFetcher{
		posValidatorsByScan:     posValidatorsByScan,
		posValidatorsByContract: posValidatorsByContract,
		scanClient:              client.NewClient(configs.Get().Server.Scan),
		cspaceClient:            sdk.MustNewClient(configs.Get().Server.CoreSpaceNode),
	}
}

type PosValidatorByScanWithResult[T any] struct {
	*configs.PosValidatorByScan
	Time time.Time
	Data T
}

func (p *PosRewardFetcher) GetPosRewardsByScan() ([]*PosValidatorByScanWithResult[decimal.Decimal], error) {
	logrus.Info("[PosRewardFetcher] get pos rewards")
	results := make([]*PosValidatorByScanWithResult[decimal.Decimal], 0)
	for _, posAddress := range p.posValidatorsByScan {
		posOverView, err := p.scanClient.GetPosAccountOverview(posAddress.PosAddress)
		if err != nil {
			return nil, err
		}
		rewardInWei, err := decimal.NewFromString(posOverView.TotalReward)
		if err != nil {
			return nil, err
		}

		rewardInEth := rewardInWei.Div(decimal.New(1, 18))
		results = append(results, &PosValidatorByScanWithResult[decimal.Decimal]{
			PosValidatorByScan: posAddress,
			Time:               time.Now(),
			Data:               rewardInEth,
		})
	}
	return results, nil
}

type PosValidatorByContractWithResult[T any] struct {
	*configs.PosValidatorByContract
	Time  time.Time
	Epoch *big.Int
	Data  T
}

func (p *PosRewardFetcher) GetPosRewardsByContract() ([]*PosValidatorByContractWithResult[decimal.Decimal], error) {
	results := make([]*PosValidatorByContractWithResult[decimal.Decimal], 0)
	monthStart := timeutils.GetMonthStart(timeutils.GetChinaLocation())
	epochNumber, epochTime, err := cfxutil.SearchEpochOnTime(p.cspaceClient, monthStart)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to search epoch on time")
	}

	for _, v := range p.posValidatorsByContract {
		interest, err := p.GetPosReward(v, epochNumber)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to get pos reward: %s", v.Name)
		}

		interestInEth := decimal.NewFromBigInt(interest, 0).Div(decimal.New(1, 18))
		results = append(results, &PosValidatorByContractWithResult[decimal.Decimal]{
			PosValidatorByContract: v,
			Time:                   epochTime,
			Epoch:                  epochNumber,
			Data:                   interestInEth,
		})
	}

	return results, nil
}

func (p *PosRewardFetcher) GetPosReward(v *configs.PosValidatorByContract, epochNumber *big.Int) (*big.Int, error) {

	opt := &bind.CallOpts{
		EpochNumber: types.NewEpochNumberBig(epochNumber),
	}

	switch v.Method {
	case constants.POS_REWARD_METHOD_USERINTEREST:
		posPoolContract, err := pospool.NewPosPool(v.PowAddress, p.cspaceClient)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to new pos pool contract: %s", v.Name)
		}
		interest, err := posPoolContract.UserInterest(opt, v.QueryUser.MustGetCommonAddress())
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to get user interest: %s", v.Name)
		}
		return interest, nil
	case constants.POS_REWARD_METHOD_GETUSERREWARD:
		posPoolContract, err := pospooldx.NewPosPoolDx(v.PowAddress, p.cspaceClient)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to new pos pool contract: %s", v.Name)
		}
		interest, err := posPoolContract.GetUserReward(opt, v.QueryUser.MustGetCommonAddress())
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to get user interest: %s", v.Name)
		}
		return interest, nil
	default:
		return nil, errors.New("invalid method")
	}
}

func (p *PosRewardFetcher) FetchAndMail() error {
	logrus.Info("[PosRewardFetcher] === Start Collect Pos Rewards ===")
	posRewardsByContractResult, err := p.GetPosRewardsByContract()
	if err != nil {
		logrus.Errorf("[PosRewardFetcher] Collect Pos Rewards by contract failed: %v", err)
		return err
	}
	logrus.Info("[PosRewardFetcher] === Collect Pos Rewards by contract done ===")

	posRewardsByScanResult, err := p.GetPosRewardsByScan()
	if err != nil {
		logrus.Errorf("[PosRewardFetcher] Collect Pos Rewards by scan failed: %v", err)
		return err
	}
	logrus.Info("[PosRewardFetcher] === Collect Pos Rewards by scan done ===")

	reporter := NewReporter(nil, posRewardsByScanResult, posRewardsByContractResult)
	excelPath, err := reporter.WriteToExcel(&WriteEnables{
		IncludePosRewards: true,
	})
	if err != nil {
		logrus.Errorf("[PosRewardFetcher] Write to excel failed: %v", err)
		return err
	}

	err = reporter.SendMail(excelPath, configs.Get().Mail.Receivers.PosReward)
	if err != nil {
		logrus.Errorf("[PosRewardFetcher] Send mail failed: %v", err)
		return err
	}
	logrus.Info("[PosRewardFetcher] === Send mail Done ===")
	return nil
}
