package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var dataModel *CustomTableModel
var username = ""
var password = ""

func OpenFile(filePath string) {

	fmt.Println(filePath)

	xlFile, _ := xlsx.OpenFile(filePath)

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()

				if username == "" {
					username = text
				} else if password == "" {
					password = text
				} else {

					item := TableItem{text, "", "", ""}

					dataModel.add(item)
				}
				break
			}
		}
		break
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

	for index, item := range dataModel.modelData {
		item := SearchIMEI(item.IMEI, wd)
		dataModel.edit(index, item)
	}

}
