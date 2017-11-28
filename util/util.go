package util

import (
	"github.com/tealeg/xlsx"
)

func WriteXlsFile(filename string, headers []string, items [][]string) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	row = sheet.AddRow()

	for _, header := range headers {
		addCell(row, header)
	}

	for _, item := range items {
		row = sheet.AddRow()

		for _, field := range item {
			addCell(row, field)
		}
	}

	err = file.Save(filename)
	if err != nil {
		return err
	}

	return nil
}

// addCell adds new cell to xls file
func addCell(row *xlsx.Row, value string) {
	cell := row.AddCell()

	cell.Value = value
}
