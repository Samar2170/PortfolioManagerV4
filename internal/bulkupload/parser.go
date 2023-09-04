package bulkupload

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func parseBulkUploadTradeSheet(buId uint) error {
	var err error
	var stockRows [][]string
	var bondRows [][]string
	var mfRows [][]string
	var etsRows [][]string
	bus, err := GetBulkUploadSheetByID(buId)
	if err != nil {
		return err
	}
	sheet, err := excelize.OpenFile(bus.Path)
	if err != nil {
		return err
	}
	defer func() {
		if err := sheet.Close(); err != nil {
			log.Println(err)
		}
	}()
	stockRows, err = sheet.GetRows("Stock")
	if err != nil {
		return err
	}
	bondRows, err = sheet.GetRows("Bond")
	if err != nil {
		return err
	}
	mfRows, err = sheet.GetRows("MF")
	if err != nil {
		return err
	}
	etsRows, err = sheet.GetRows("ETS")
	if err != nil {
		return err
	}
	fmt.Println(stockRows)
	fmt.Println(bondRows)
	fmt.Print(mfRows)
	fmt.Print(etsRows)
	return nil
}
