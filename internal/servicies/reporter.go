package servicies

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"
	"github.com/conflux-fans/cmctool/pkg/cmcsdk/types"
	"github.com/conflux-fans/cmctool/pkg/email"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type Reporter struct {
	AllTokenMarketPairs map[string]*types.MarketPairsResp
	PosRewards          map[common.Hash]decimal.Decimal
}

func NewReporter(allTokenMarketPairs map[string]*types.MarketPairsResp, posRewards map[common.Hash]decimal.Decimal) *Reporter {
	return &Reporter{
		AllTokenMarketPairs: allTokenMarketPairs,
		PosRewards:          posRewards,
	}
}

func (r *Reporter) WriteToExcel() (string, error) {
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	excelPath := fmt.Sprintf("./out/volume-and-pos-%v.xlsx", yesterday)

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
	if err := r.writeTokenMarket(f); err != nil {
		return "", err
	}

	if err := r.writePosRewards(f); err != nil {
		return "", err
	}

	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(0)

	// 保存 Excel 文件
	if err := f.SaveAs(excelPath); err != nil {
		return "", err
	}

	logrus.Infoln("[Reporter] Excel file created successfully.")
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
	f.SetCellValue(sheetName, "A1", "POS 地址")
	f.SetCellValue(sheetName, "B1", "奖励(CFX)")

	row := 2
	for posAddress, reward := range r.PosRewards {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(row), posAddress.Hex())
		f.SetCellValue(sheetName, "B"+strconv.Itoa(row), reward.String())
		row++
	}
	return nil
}

func (r *Reporter) SendMail(excelPath string) error {
	sender := configs.Get().Mail.Sender
	receivers := configs.Get().Mail.Receivers

	m := email.
		NewMailer(sender.Host, sender.Port, sender.Username, sender.Password).
		SetOption(new(email.MailerOptions).WithSenderName(sender.FromAlias))
	if err := m.Send(receivers, "交易量与POS奖励统计", "", []string{excelPath}); err != nil {
		return err
	}
	logrus.Info("[Reporter] Send mail done")
	logrus.Info("[Reporter] === Collect Volumes Done ===")

	return nil
}
