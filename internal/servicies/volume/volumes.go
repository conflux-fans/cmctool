package volume

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/conflux-fans/cmctool/internal/configs"
	sdk "github.com/conflux-fans/cmctool/pkg/cmcsdk"
	"github.com/conflux-fans/cmctool/pkg/cmcsdk/types"
	"github.com/conflux-fans/cmctool/pkg/common"
	"github.com/conflux-fans/cmctool/pkg/email"
	"github.com/robfig/cron/v3"
	"github.com/xuri/excelize/v2"
)

var (
	client *sdk.Client
)

func init() {
	client = sdk.NewClient()
}

func Loop() {
	c := cron.New()

	var entryId cron.EntryID
	entryId, _ = c.AddFunc("@daily", func() {
		common.Retry(3, time.Minute, collectVolumeAndMail)
		nextTime := c.Entry(cron.EntryID(entryId)).Next
		log.Printf("Next task execution time: %v\n", nextTime)
	})

	c.Start()
}

func collectVolumeAndMail() error {
	log.Print("=== Start Collect Volumes ===")
	spots, err := collectSpotVloumes()
	if err != nil {
		return err
	}
	log.Print("Get spots done")

	perps, err := collectPerpVolumes()
	if err != nil {
		return err
	}
	log.Print("Get perps done")

	all := spots
	for k, v := range perps {
		all[k] = v
	}

	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	excelPath := fmt.Sprintf("./out/volumes-%v.xlsx", yesterday)

	if err = writeToExcel(excelPath, spots); err != nil {
		return err
	}
	log.Print("Save excel done")

	sender := configs.Get().Mail.Sender
	receivers := configs.Get().Mail.Receivers

	m := email.
		NewMailer(sender.Host, sender.Port, sender.Username, sender.Password).
		SetOption(new(email.MailerOptions).WithSenderName(sender.FromAlias))
	if err = m.Send(receivers, "交易量统计", "", []string{excelPath}); err != nil {
		return err
	}
	log.Print("Send mail done")
	log.Print("=== Collect Volumes Done ===")

	return nil
}

func collectSpotVloumes() (spots map[string]*types.MarketPairsResp, err error) {
	// time.Sleep(time.Second)
	btcSpot, err := client.MarketPairLatest(types.MarketPairQueryParams{"bitcoin", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get btc spot done")

	time.Sleep(time.Second)
	ethSpot, err := client.MarketPairLatest(types.MarketPairQueryParams{"ethereum", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get eth spot done")

	time.Sleep(time.Second)
	cfxSpot, err := client.MarketPairLatest(types.MarketPairQueryParams{"conflux-network", "spot", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get cfx spot done")

	spots = map[string]*types.MarketPairsResp{
		"BTC-Spot": btcSpot,
		"ETH-Spot": ethSpot,
		"CFX-Spot": cfxSpot,
	}
	return
}

func collectPerpVolumes() (perps map[string]*types.MarketPairsResp, err error) {
	time.Sleep(time.Second)
	btcPerp, err := client.MarketPairLatest(types.MarketPairQueryParams{"bitcoin", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get btc perp done")

	time.Sleep(time.Second)
	ethPerp, err := client.MarketPairLatest(types.MarketPairQueryParams{"ethereum", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get eth perp done")

	time.Sleep(time.Second)
	cfxPerp, err := client.MarketPairLatest(types.MarketPairQueryParams{"conflux-network", "perpetual", 1, 100})
	if err != nil {
		return nil, err
	}
	log.Println("get cfx perp done")

	perps = map[string]*types.MarketPairsResp{
		"BTC-Perp": btcPerp,
		"ETH-Perp": ethPerp,
		"CFX-Perp": cfxPerp,
	}
	return
}

func writeToExcel(excelPath string, allTokenMarketPairs map[string]*types.MarketPairsResp) error {
	dir := filepath.Dir(excelPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	for mpName, tokenResult := range allTokenMarketPairs {
		// 创建一个新的工作表
		_, err := f.NewSheet(mpName)
		if err != nil {
			return err
		}

		// 设置表头
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

		// 写入结构体数组数据
		for i, mpLine := range tokenResult.Data.MarketPairs {
			row := i + 2 // 从第二行开始写入数据
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

	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(0)

	// 保存 Excel 文件
	if err := f.SaveAs(excelPath); err != nil {
		return err
	}

	log.Println("Excel file created successfully.")
	return nil
}
