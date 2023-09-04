package internal

import (
	"fmt"
	"log"

	"github.com/samar2170/portfolio-manager-v4/pkg/utils/structs"
	"github.com/xuri/excelize/v2"
)

func CreateTradeTemplate() error {
	var err error
	// f, err := excelize.OpenFile("assets/trade-template.xlsx", excelize.Options{})
	f := excelize.NewFile()
	// if err != nil {
	// 	return err
	// }
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	var str StockTradeRequest
	var btr BondTradeRequest
	var mtr MutualFundTradeRequest
	var etr ETSTradeRequest

	var names []string
	var types []string
	names, types = createRowFromApiRequest(str)
	fmt.Println(names, types)
	fmt.Println(len(names))
	for i, name := range names {
		fmt.Println(i, name)
	}

	fmt.Printf("%T\n", names)
	fmt.Printf("%T\n", types)
	f.NewSheet("Stock")
	err = f.SetSheetRow("Stock", "A1", &names)
	if err != nil {
		return err
	}
	err = f.SetSheetRow("Stock", "A2", &types)
	if err != nil {
		return err
	}
	names, types = createRowFromApiRequest(btr)
	f.NewSheet("Bond")
	f.SetSheetRow("Bond", "A1", &names)
	f.SetSheetRow("Bond", "A2", &types)

	names, types = createRowFromApiRequest(mtr)
	f.NewSheet("MutualFund")
	f.SetSheetRow("MutualFund", "A1", &names)
	f.SetSheetRow("MutualFund", "A2", &types)

	names, types = createRowFromApiRequest(etr)
	f.NewSheet("ETS")

	err = f.SetSheetRow("ETS", "A1", &names)
	if err != nil {
		return err
	}
	f.SetSheetRow("ETS", "A2", &types)

	if err := f.SaveAs("assets/trade-template.xlsx"); err != nil {
		return err
	}
	return nil
}

func createRowFromApiRequest(t interface{}) (names []string, types []string) {
	s := structs.New(t)
	m := s.MapWithType()
	for k, v := range m {
		names = append(names, k)
		types = append(types, v)
	}
	return
}

func TestExcelize() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
