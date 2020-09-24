package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("./江苏省厅指标.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	
	for _, name := range f.GetSheetMap() {
		// fmt.Println(index, name)
		//输入所有项目类型
		 projectType,err := f.GetCellValue(name,"A1")
		 if err !=nil{
			 fmt.Println(err)
		 }
		 fmt.Println(projectType)
		// rows, _ := f.GetRows(f.GetSheetName(index))
		// for _, row := range rows {
		// 	for _, colCell := range row {
		// 		fmt.Print(colCell, "\t")
		// 	}
		// 	fmt.Println()
		// }
	}
}
