package servicies

import (
	"strconv"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/pkg/common"
	"github.com/conflux-fans/go-scan-sdk/client"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type PosRewardHistoryFetcher struct {
	posValidatorsByContract []*configs.PosValidatorByContract
	scanClient              *client.Client
}

func NewPosRewardHistoryFetcher(posValidatorsByContract []*configs.PosValidatorByContract) *PosRewardHistoryFetcher {
	return &PosRewardHistoryFetcher{
		posValidatorsByContract: posValidatorsByContract,
		scanClient:              client.NewClient(configs.Get().Server.Scan),
	}
}

func (p *PosRewardHistoryFetcher) GetPosRewardHistory() ([]*PosValidatorByContractWithResult[*client.List[*client.PosAccountReward]], error) {
	results := make([]*PosValidatorByContractWithResult[*client.List[*client.PosAccountReward]], 0)
	for _, posValidator := range p.posValidatorsByContract {
		rewards, err := p.scanClient.GetPosAccountRewardIncomingHistory(posValidator.PosAddress)
		if err != nil {
			return nil, err
		}
		results = append(results, &PosValidatorByContractWithResult[*client.List[*client.PosAccountReward]]{
			PosValidatorByContract: posValidator,
			Data:                   rewards,
		})
	}
	return results, nil
}

func (p *PosRewardHistoryFetcher) WriteExcel(historyies []*PosValidatorByContractWithResult[*client.List[*client.PosAccountReward]]) (string, error) {
	f := excelize.NewFile()

	for _, history := range historyies {
		sheetName := history.Name
		common.Must(f.NewSheet(sheetName))
		common.MustNil(f.SetCellValue(sheetName, "A1", "时间"))
		common.MustNil(f.SetCellValue(sheetName, "B1", "奖励"))
		common.MustNil(f.SetCellValue(sheetName, "C1", "Epoch"))
		common.MustNil(f.SetCellValue(sheetName, "D1", "POW 区块哈希"))

		row := 2
		for _, reward := range history.Data.List {
			rewardInOnes, err := decimal.NewFromString(reward.Reward)
			if err != nil {
				return "", err
			}
			rewardInOnes = rewardInOnes.Div(decimal.New(1, 18))
			common.MustNil(f.SetCellValue(sheetName, "A"+strconv.Itoa(row), reward.CreatedAt.Format("2006-01-02 15:04:05")))
			common.MustNil(f.SetCellValue(sheetName, "B"+strconv.Itoa(row), rewardInOnes.String()))
			common.MustNil(f.SetCellValue(sheetName, "C"+strconv.Itoa(row), reward.Epoch))
			common.MustNil(f.SetCellValue(sheetName, "D"+strconv.Itoa(row), reward.PowBlockHash))
			row++
		}
		logrus.WithField("sheetName", sheetName).Infoln("[PosRewardHistoryFetcher] Write Excel")
	}

	common.MustNil(f.DeleteSheet("Sheet1"))
	f.SetActiveSheet(0)

	if err := f.SaveAs("pos_reward_history.xlsx"); err != nil {
		return "", err
	}

	return f.Path, nil
}

func (p *PosRewardHistoryFetcher) FetchAndWriteExcel() error {
	historyies, err := p.GetPosRewardHistory()
	if err != nil {
		return err
	}

	outPath, err := p.WriteExcel(historyies)
	if err != nil {
		return err
	}

	logrus.WithField("outPath", outPath).Infoln("[PosRewardHistoryFetcher] Excel file created successfully.")
	return nil
}
