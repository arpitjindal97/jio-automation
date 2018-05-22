package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var tabledata = []TableItem{}

var dataModel *CustomTableModel

func OpenFile(filePath string) {

	fmt.Println(filePath)

	xlFile, _ := xlsx.OpenFile(filePath)

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()

				item := TableItem{text, "", "", "", ""}

				dataModel.add(item)

				break
			}
		}
		break
	}
}

func SetModalInstance(m *CustomTableModel) {
	dataModel = m
}
