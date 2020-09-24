package main

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"

	// "database/sql"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	step1()
	step2()
	step3()
}

func step1() {
	f, err := excelize.OpenFile("./指标库0813.xlsx")
	// firststd :=excelize.NewFile()
	//第一步：先把通用类的指标复制单独成一个
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheetName := range f.GetSheetMap() {

		targetNo, _ := f.GetCellValue(sheetName, "A5")
		if strings.HasPrefix(targetNo, "1") {
			f.DeleteSheet(sheetName)
		}
	}
	f.SaveAs("非通用类.xlsx")
	f1, _ := excelize.OpenFile("./指标库0813.xlsx")
	for _, sheetName1 := range f1.GetSheetMap() {

		targetNo1, _ := f1.GetCellValue(sheetName1, "A5")
		if !strings.HasPrefix(targetNo1, "1") {
			f1.DeleteSheet(sheetName1)
		}
	}
	f1.SaveAs("通用类.xlsx")
}

//循环《通用类》将内容复制到另外一个excel
func step2() {
	f, err := excelize.OpenFile("./通用类.xlsx")
	f1, err := excelize.OpenFile("./temp2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	f1sheetindex := 1
	for _, sheetName := range f.GetSheetMap() {
		rows, _ := f.Rows(sheetName)
		f1.NewSheet(sheetName)
		f1.CopySheet(f1.GetSheetIndex("Sheet1"), f1.GetSheetIndex(sheetName))
		f1sheetindex++
		for ti := 5; ti <= rows.TotalRow; ti++ {

			// f1.GetSheetName(f1index);

			a5, _ := f.GetCellValue(sheetName, "A"+strconv.Itoa(ti))
			if a5 == "" {
				break
			}
			f1.SetCellValue(sheetName, "A"+strconv.Itoa(ti-1), a5)
			b5, _ := f.GetCellValue(sheetName, "B"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "B"+strconv.Itoa(ti-1), b5)
			c5, _ := f.GetCellValue(sheetName, "C"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "C"+strconv.Itoa(ti-1), c5)
			d5, _ := f.GetCellValue(sheetName, "D"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "D"+strconv.Itoa(ti-1), d5)
			e5, _ := f.GetCellValue(sheetName, "E"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "E"+strconv.Itoa(ti-1), e5)
			h5, _ := f.GetCellValue(sheetName, "H"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "F"+strconv.Itoa(ti-1), h5)
			i5, _ := f.GetCellValue(sheetName, "I"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "H"+strconv.Itoa(ti-1), i5)
			f5, _ := f.GetCellValue(sheetName, "F"+strconv.Itoa(ti))
			g5, _ := f.GetCellValue(sheetName, "G"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "I"+strconv.Itoa(ti-1), f5+" "+g5)
			f1.SetCellValue(sheetName, "J"+strconv.Itoa(ti-1), "当年实际完成值")
			f1.SetCellValue(sheetName, "M"+strconv.Itoa(ti-1), "省,市,县区")
			f1.SetCellValue(sheetName, "O"+strconv.Itoa(ti-1), sheetName)
			// f1.SetSheetVisible(true)
		}

	}
	f1.SaveAs("【已处理】通用类-v2.xlsx")
}

func step3() {
	f, err := excelize.OpenFile("./非通用类.xlsx")
	f1, err := excelize.OpenFile("./temp3.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	f1sheetindex := 1
	for _, sheetName := range f.GetSheetMap() {
		rows, _ := f.Rows(sheetName)
		f1.NewSheet(sheetName)
		f1.CopySheet(f1.GetSheetIndex("Sheet1"), f1.GetSheetIndex(sheetName))
		f1sheetindex++
		for ti := 5; ti <= rows.TotalRow; ti++ {

			// f1.GetSheetName(f1index);

			a5, _ := f.GetCellValue(sheetName, "A"+strconv.Itoa(ti))
			if a5 == "" {
				break
			}
			a51, _ := f.GetCellValue(sheetName, "B"+strconv.Itoa(ti))
			a52, _ := f.GetCellValue(sheetName, "C"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "A"+strconv.Itoa(ti-1), a5+a51+a52)
			d5, _ := f.GetCellValue(sheetName, "E"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "B"+strconv.Itoa(ti-1), d5)
			e5, _ := f.GetCellValue(sheetName, "F"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "C"+strconv.Itoa(ti-1), e5)
			f5, _ := f.GetCellValue(sheetName, "G"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "D"+strconv.Itoa(ti-1), f5)
			g5, _ := f.GetCellValue(sheetName, "H"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "E"+strconv.Itoa(ti-1), g5)
			j5, _ := f.GetCellValue(sheetName, "K"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "F"+strconv.Itoa(ti-1), j5)
			k5, _ := f.GetCellValue(sheetName, "L"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "G"+strconv.Itoa(ti-1), k5)
			l5, _ := f.GetCellValue(sheetName, "M"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "L"+strconv.Itoa(ti-1), l5)
			h5, _ := f.GetCellValue(sheetName, "I"+strconv.Itoa(ti))
			i5, _ := f.GetCellValue(sheetName, "J"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "I"+strconv.Itoa(ti-1), h5+" "+i5)

			m5, _ := f.GetCellValue(sheetName, "N"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "J"+strconv.Itoa(ti-1), m5)

			n5, _ := f.GetCellValue(sheetName, "O"+strconv.Itoa(ti))
			f1.SetCellValue(sheetName, "L"+strconv.Itoa(ti-1), n5)

			f1.SetCellValue(sheetName, "M"+strconv.Itoa(ti-1), "省,市,县区")
			f1.SetCellValue(sheetName, "P"+strconv.Itoa(ti-1), sheetName)
			// f1.SetSheetVisible(true)
		}

	}
	f1.SaveAs("【已处理】非通用类-v2.xlsx")
}
