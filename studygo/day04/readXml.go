package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"com.todaytech.ben/studygo/day04/dto"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xmlFold := "./要处理的"
	outFold := "./程序生成的"
	files, err := ioutil.ReadDir(xmlFold)
	if err != nil {
		// fmt.Println(err)
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(xmlFold + string(os.PathSeparator) + file.Name())
		ErrHandler(err)
		var first dto.TestsuiteFirst
		err2 := xml.Unmarshal(content, &first)

		ErrHandler(err2)
		//得到xml中的内容
		if len(first.Testcases) > 0 {
			//把第一级的testcase加到第二级当中去
			secondTestSuite := &dto.Testsuite{}
			secondTestSuite.Testcases = append(secondTestSuite.Testcases, first.Testcases...)
			first.TestsuiteSecond = append(first.TestsuiteSecond, *secondTestSuite)
		}
		exportToExcel(first.TestsuiteSecond, outFold, file.Name())
	}

	//生成excel

}
func exportToExcel(testsuiteSecond []dto.Testsuite, outFold string, fileName string) {

	f := excelize.NewFile()
	sheetIndex := 1

	//设置sheet 模板
	//注意range和使用下标的区别 range为静态而使用下标为动态，即在循环里边对切片的长度进行变化时会影响循环次数
	for i := 0; i < len(testsuiteSecond); i++ {
		//每一个testsuite就是一个sheet

		testsuite := testsuiteSecond[i]
		sheetName := "Sheet" + strconv.Itoa(sheetIndex)
		if len(testsuite.TestsuiteThird) > 0 {
			// fmt.Print("执行前长度\t")
			// fmt.Println(len(testsuiteSecond))
			testsuiteSecond = append(testsuiteSecond, testsuite.TestsuiteThird...)
			// fmt.Print("执行后长度\t")
			// fmt.Println(len(testsuiteSecond))
			//对于这种不规则的，只对子树进行循环

		}
		if len(testsuite.Testcases) == 0 {
			continue
		}

		if sheetIndex != 1 {
			f.NewSheet(sheetName)
		}
		colposition := 1
		f.SetCellValue(sheetName, "A"+strconv.Itoa(colposition), "用例标题")

		f.SetCellValue(sheetName, "B"+strconv.Itoa(colposition), "摘要")

		f.SetCellValue(sheetName, "C"+strconv.Itoa(colposition), "前提")

		f.SetCellValue(sheetName, "D"+strconv.Itoa(colposition), "步骤编号")

		f.SetCellValue(sheetName, "E"+strconv.Itoa(colposition), "测试步骤")

		f.SetCellValue(sheetName, "F"+strconv.Itoa(colposition), "期望结果")
		colposition++
		for _, testcase := range testsuite.Testcases {

			f.SetCellValue(sheetName, "A"+strconv.Itoa(colposition), regReplaceAll(testcase.Name))
			f.SetCellValue(sheetName, "B"+strconv.Itoa(colposition), regReplaceAll(testcase.Summary))
			f.SetCellValue(sheetName, "C"+strconv.Itoa(colposition), regReplaceAll(testcase.Preconditions))
			for _, step := range testcase.Steps.Step {
				f.SetCellValue(sheetName, "D"+strconv.Itoa(colposition), regReplaceAll(step.StepNumber))
				f.SetCellValue(sheetName, "E"+strconv.Itoa(colposition), regReplaceAll(step.Action))
				f.SetCellValue(sheetName, "F"+strconv.Itoa(colposition), regReplaceAll(step.Expectedresults))
				colposition++
			}
			//colposition++

		}

		//fmt.Println("----------执行次数-----------")
		sheetIndex++
	}

	fileName = outFold + string(os.PathSeparator) + strings.Split(fileName, ".")[0] + ".xlsx"
	err := f.SaveAs(fileName)

	if err != nil {
		// fmt.Println(err)
	}
	fmt.Printf("%s文件生成成功\n", fileName)
	// Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	//f.SetActiveSheet(0)
	// Save xlsx file by the given path.

}

//ErrHandler 错误处理函数
func ErrHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func regReplaceAll(old string) string {
	//默认是贪婪模式；在量词后面直接加上一个问号？就是非贪婪模式。
	replaceReg := "<.*?>|&.*?;|\n|\t"
	// fmt.Print(old)
	reg := regexp.MustCompile(replaceReg)
	reStr := reg.ReplaceAllString(old, "")
	//去掉空格
	// reg1 := regexp.MustCompile("\r\n")
	// reStr1 := reg1.ReplaceAllString(reStr,"\r\n")
	// reStr1 = strings.ReplaceAll(reStr1,"\n",``)
	// reStr2 := strings.ReplaceAll(reStr1, "\r\n", ``)
	fmt.Print(reStr)

	// reStr2 :=strings.ReplaceAll(reStr1,"",``)
	// fmt.Println(string(reStr2))
	// if strings.Index(old,"已安装qq") >0 {
	// fmt.Print(strings.Index(reStr," "))
	// fmt.Print(reStr)
	// }
	return reStr
}
