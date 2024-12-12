package servicies

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/pkg/cmcsdk/types"
	"github.com/conflux-fans/cmctool/pkg/email"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type Reporter struct {
	AllTokenMarketPairs        map[string]*types.MarketPairsResp
	PosRewardsByScanResult     []*PosValidatorByScanWithResult
	PosRewardsByContractResult []*PosValidatorByContractWithResult
}

func NewReporter(allTokenMarketPairs map[string]*types.MarketPairsResp, posRewardsByScanResult []*PosValidatorByScanWithResult, posRewardsByContractResult []*PosValidatorByContractWithResult) *Reporter {
	return &Reporter{
		AllTokenMarketPairs:        allTokenMarketPairs,
		PosRewardsByScanResult:     posRewardsByScanResult,
		PosRewardsByContractResult: posRewardsByContractResult,
	}
}

// func (r *Reporter) WriteAllToExcel() (string, error) {
// 	return r.writeToExcel(true, true)
// }

// func (r *Reporter) WriteVolumeToExcel() (string, error) {
// 	return r.writeToExcel(true, false)
// }

// func (r *Reporter) WritePosRewardsToExcel() (string, error) {
// 	return r.writeToExcel(false, true)
// }

type WriteEnables struct {
	IncludeTokenMarket bool
	IncludePosRewards  bool
}

func (r *Reporter) WriteToExcel(enables *WriteEnables) (string, error) {
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	excelPath := "./out/"
	if enables.IncludeTokenMarket {
		excelPath += "交易量-"
	}
	if enables.IncludePosRewards {
		excelPath += "POS奖励-"
	}
	excelPath += yesterday + ".xlsx"

	dir := filepath.Dir(excelPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 创建一个新的工作表
	// 设置表头
	// 写入结构体数组数据
	// 从第二行开始写入数据
	if enables.IncludeTokenMarket {
		if err := r.writeTokenMarket(f); err != nil {
			return "", err
		}
	}

	if enables.IncludePosRewards {
		if err := r.writePosRewards(f); err != nil {
			return "", err
		}
	}

	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(0)

	// 保存 Excel 文件
	if err := f.SaveAs(excelPath); err != nil {
		return "", err
	}

	logrus.WithField("excelPath", excelPath).WithField("includePosRewards", enables.IncludePosRewards).WithField("includeTokenMarket", enables.IncludeTokenMarket).Infoln("[Reporter] Excel file created successfully.")
	return excelPath, nil
}

func (r *Reporter) writeTokenMarket(f *excelize.File) error {
	for mpName, tokenResult := range r.AllTokenMarketPairs {

		_, err := f.NewSheet(mpName)
		if err != nil {
			return err
		}

		f.SetCellValue(mpName, "A1", "交易所")
		f.SetCellValue(mpName, "B1", "类别")
		f.SetCellValue(mpName, "C1", "币对")
		f.SetCellValue(mpName, "D1", "链接")
		f.SetCellValue(mpName, "E1", "币价")
		f.SetCellValue(mpName, "F1", "交易量")
		f.SetCellValue(mpName, "G1", "交易量占比")
		f.SetCellValue(mpName, "H1", "-2%深度")
		f.SetCellValue(mpName, "I1", "+2%深度")
		f.SetCellValue(mpName, "J1", "更新时间")
		f.SetCellValue(mpName, "K1", "交易所类型")

		for i, mpLine := range tokenResult.Data.MarketPairs {
			row := i + 2
			f.SetCellValue(mpName, "A"+strconv.Itoa(row), mpLine.ExchangeName)
			f.SetCellValue(mpName, "B"+strconv.Itoa(row), mpLine.Category)
			f.SetCellValue(mpName, "C"+strconv.Itoa(row), mpLine.MarketPair)
			f.SetCellValue(mpName, "D"+strconv.Itoa(row), mpLine.MarketURL)
			f.SetCellValue(mpName, "E"+strconv.Itoa(row), mpLine.Price)
			f.SetCellValue(mpName, "F"+strconv.Itoa(row), mpLine.VolumeUsd)
			f.SetCellValue(mpName, "G"+strconv.Itoa(row), mpLine.VolumePercent)
			f.SetCellValue(mpName, "H"+strconv.Itoa(row), mpLine.DepthUsdNegativeTwo)
			f.SetCellValue(mpName, "I"+strconv.Itoa(row), mpLine.DepthUsdPositiveTwo)
			f.SetCellValue(mpName, "J"+strconv.Itoa(row), mpLine.LastUpdated)
			f.SetCellValue(mpName, "K"+strconv.Itoa(row), mpLine.CenterType)
		}
	}
	return nil
}

func (r *Reporter) writePosRewards(f *excelize.File) error {
	sheetName := "POS奖励"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Validator")
	f.SetCellValue(sheetName, "B1", "POS 地址")
	f.SetCellValue(sheetName, "C1", "POW 地址")
	f.SetCellValue(sheetName, "D1", "查询用户")
	f.SetCellValue(sheetName, "E1", "奖励(CFX)")

	row := 2
	for _, reward := range r.PosRewardsByScanResult {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(row), reward.Name)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(row), reward.PosAddress.Hex())
		f.SetCellValue(sheetName, "C"+strconv.Itoa(row), reward.PowAddress.String())
		f.SetCellValue(sheetName, "E"+strconv.Itoa(row), reward.Reward.String())
		row++
	}

	for _, reward := range r.PosRewardsByContractResult {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(row), reward.Name)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(row), reward.PosAddress.Hex())
		f.SetCellValue(sheetName, "C"+strconv.Itoa(row), reward.PowAddress.String())
		f.SetCellValue(sheetName, "D"+strconv.Itoa(row), reward.QueryUser.String())
		f.SetCellValue(sheetName, "E"+strconv.Itoa(row), reward.Reward.String())
		row++
	}

	return nil
}

func (r *Reporter) SendMail(excelPath string, receivers []string) error {
	sender := configs.Get().Mail.Sender
	// receivers := configs.Get().Mail.Receivers

	m := email.
		NewMailer(sender.Host, sender.Port, sender.Username, sender.Password).
		SetOption(new(email.MailerOptions).WithSenderName(sender.FromAlias))
	if err := m.Send(receivers, filepath.Base(excelPath), "", []string{excelPath}); err != nil {
		return err
	}
	logrus.Info("[Reporter] Send mail done")
	logrus.Info("[Reporter] === Collect Done ===")

	return nil
}
