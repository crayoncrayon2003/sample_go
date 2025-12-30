package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

/*
-------------------------------------------------

	Excelファイルを生成する関数

-------------------------------------------------
*/
func writeExcel(filename string) error {
	f := excelize.NewFile()

	sheet := "Sheet1"
	f.SetSheetName("Sheet1", sheet)

	// ヘッダ
	headers := []string{
		"ID", "Name", "Status", "Description",
	}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// データ行
	data := [][]interface{}{
		{1000, "SampleName", "Active", "Description sample"},
		{1001, "SampleName2", "Inactive", "Another sample"},
	}

	for r, row := range data {
		for c, v := range row {
			cell, _ := excelize.CoordinatesToCellName(c+1, r+2)
			f.SetCellValue(sheet, cell, v)
		}
	}

	return f.SaveAs(filename)
}

/*
-------------------------------------------------

	Excelファイルを読み込む関数

-------------------------------------------------
*/
func readExcel(filename string) error {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for i, row := range rows {
		fmt.Printf("row %d\n", i+1)
		for j, col := range row {
			fmt.Printf("  [%02d] %q\n", j, col)
		}
		fmt.Println("--------------------------------")
	}
	return nil
}

/*
-------------------------------------------------

	main

-------------------------------------------------
*/
func main() {
	filename := "sample.xlsx"

	// Excel生成
	if err := writeExcel(filename); err != nil {
		fmt.Println("write error:", err)
		return
	}

	// Excel読込
	if err := readExcel(filename); err != nil {
		fmt.Println("read error:", err)
		return
	}
}
