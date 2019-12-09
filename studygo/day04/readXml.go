package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	content, err := ioutil.ReadFile(`./要处理的/testsuite-deep.xml`)
	ErrHandler(err)
	var first TestsuiteFirst
	err2 := xml.Unmarshal(content, &first)

	ErrHandler(err2)
	//得到xml中的内容

	exportToExcel(first.TestsuiteSecond)

	//生成excel

}
func exportToExcel(testsuiteSecond []Testsuite) {
	f := excelize.NewFile()
	sheetIndex := 1
	//设置sheet 模板
	//注意range和使用下标的区别 range为静态而使用下标为动态，即在循环里边对切片的长度进行变化时会影响循环次数
	for i := 0; i < len(testsuiteSecond); i++ {
		//每一个testsuite就是一个sheet
		testsuite := testsuiteSecond[i]
		sheetName := "Sheet" + strconv.Itoa(sheetIndex)
		if len(testsuite.TestsuiteThird) > 0 {
			fmt.Print("执行前长度\t")
			fmt.Println(len(testsuiteSecond))
			testsuiteSecond = append(testsuiteSecond, testsuite.TestsuiteThird...)
			fmt.Print("执行后长度\t")
			fmt.Println(len(testsuiteSecond))
			//对于这种不规则的，只对子树进行循环
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
			f.SetCellValue(sheetName, "A"+strconv.Itoa(colposition), testcase.Name)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(colposition), testcase.Summary)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(colposition), testcase.Preconditions)
			for _, step := range testcase.Steps.Step {
				f.SetCellValue(sheetName, "D"+strconv.Itoa(colposition), step.StepNumber)
				f.SetCellValue(sheetName, "E"+strconv.Itoa(colposition), step.Action)
				f.SetCellValue(sheetName, "F"+strconv.Itoa(colposition), step.Expectedresults)
				colposition++
			}
			colposition++
		}

		fmt.Println("----------执行次数-----------")
		sheetIndex++
	}

	fileName := "生成的测试用例结果.xlsx"
	err := f.SaveAs(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("生成成功，请查看《%s》文件", fileName)
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

//TestsuiteFirst xml最顶层
type TestsuiteFirst struct {
	NodeOrder       string      `xml:"node_order"`
	Details         string      `xml:"details"`
	TestsuiteSecond []Testsuite `xml:"testsuite"`
}

//Testsuite xml第二层
type Testsuite struct {
	Nodeorder      string      `xml:"node_order"`
	Details        string      `xml:"details"`
	Testcases      []Testcase  `xml:"testcase"`
	Name           string      `xml:"name,attr"`
	TestsuiteThird []Testsuite `xml:"testsuite"`
}

// //Testsuite xml第三层
// type Testsuite struct {
// 	Nodeorder string     `xml:"node_order"`
// 	Details   string     `xml:"details"`
// 	Testcases []Testcase `xml:"testcase"`
// 	Name      string     `xml:"name,attr"`
// }

//Testcase 结构体为xml文件对应的实体类
type Testcase struct {
	NodeOrder             string `xml:"node_order"`
	Externalid            string `xml:"externalid"`
	Version               string `xml:"version"`
	Summary               string `xml:"summary"`
	Preconditions         string `xml:"preconditions"`
	ExecutionType         string `xml:"excution_type"`
	Importance            string `xml:"importance"`
	EstimatedExecDuration string `xml:"estimated_exec_duration"`
	Status                string `xml:"status"`
	Steps                 Steps  `xml:"steps"`
	Name                  string `xml:"name,attr"`
	Internalid            string `xml:"internalid,attr"`
}

//Steps testcase中的步骤
type Steps struct {
	Step []Step `xml:"step"`
}

//Step Steps里的元素
type Step struct {
	StepNumber      string `xml:"step_number"`
	Action          string `xml:"actions"`
	Expectedresults string `xml:"expectedresults"`
	ExecutionType   string `xml:"execution_type"`
}
