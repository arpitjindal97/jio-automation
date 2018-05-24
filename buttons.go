package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type TableItem struct {
	IMEI        string
	Distributor string
	Retailer    string
	JCNum       string
}

var dataModel *CustomTableModel
var username = ""
var password = ""
var excelFile string

func OpenFile(filePath string) {

	fmt.Println(filePath)
	excelFile = filePath

	xlsx, _ := excelize.OpenFile(excelFile)

	sheetName := xlsx.GetSheetName(1)

	rows := xlsx.GetRows(sheetName)

	for _, row := range rows {
		for _, cell := range row {

			if username == "" {
				username = cell
			} else if password == "" {
				password = cell
			} else {

				item := TableItem{cell, "", "", ""}

				dataModel.add(item)
			}
			break
		}
	}

}

func SetModalInstance(m *CustomTableModel) {
	dataModel = m
}

func StartProcess() {

	fmt.Println("Username : " + username)
	fmt.Println("Password: " + password)

	service, wd := SetupSelenium()
	defer service.Stop()

	xlsx, _ := excelize.OpenFile(excelFile)
	sheet := xlsx.GetSheetName(1)

	for index, item := range dataModel.modelData {
		item := SearchIMEI(item.IMEI, wd)
		dataModel.edit(index, item)

		xlsx.SetCellValue(sheet, "B"+strconv.Itoa(index+3), item.Distributor)
		xlsx.SetCellValue(sheet, "C"+strconv.Itoa(index+3), item.Retailer)
		xlsx.SetCellValue(sheet, "D"+strconv.Itoa(index+3), item.JCNum)
		xlsx.Save()
	}

}
