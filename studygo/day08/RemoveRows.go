package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	//  stepOne();
	readUnit()

}

func printErr(err error) {
	if err != nil {
		// fmt.Println(err)
	}
}

// 如果没内容返回true,如果有返回false
func hasContent(cv string) bool {
	if cv == "" {
		return true
	}
	return false

}

//去掉第三和第四行
func stepOne() {
	f, err := excelize.OpenFile("./附件3-2020年省级专项资金绩效目标表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range f.GetSheetMap() {

		// var nocontent = false
		// fmt.Println(name)
		content, _ := f.GetCellValue(name, "A3")
		var flag = hasContent(content)
		content1, _ := f.GetCellValue(name, "A4")
		var flag1 = hasContent(content1)
		
		//获取所有单位名称start 第二步
		// fmt.Println(f.GetCellValue(name,"D3"))
		// end
		if flag {
			err := f.RemoveRow(name, 3)
			println(err)

		}
		if flag1 {
			err1 := f.RemoveRow(name, 3)
			println(err1)
		}

	}
	f.SaveAs("result.xlsx")
}

func readUnit() {
	f, _ := excelize.OpenFile("./result.xlsx")
	
	for _, name := range f.GetSheetMap() {
		unit, _ := f.GetCellValue(name, "D3")
		if unit != "" {
			fmt.Println(name + "\t" + unit)
		}
	}
}
