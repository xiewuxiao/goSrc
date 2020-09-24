package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"database/sql"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// f, err := excelize.OpenFile("./【处理】附件3-2020年省级专项资金绩效目标表.xlsx")
	f, err := excelize.OpenFile("./【处理】附件1-2020年省级部门整体绩效目标表.xlsx")
	//【处理】附件2-2020年省级部门预算归类项目绩效目标表.xlsx
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(firstTarget)
	// fmt.Println(firstTarget1)
	var colTask = []string{"B", "D","F", "G", "H"}
	var colTarget = []string{"B", "D", "F", "G", "H"}
	// var index = 10

	var list = make([]Targetapp, 0)
	for _, sheetName := range f.GetSheetMap() {
		var targetapp Targetapp
		rows, _ := f.Rows(sheetName)
		if sheetName != "目录" {
			// if sheetName == "Sheet1" {

			targetapp.UnitName, _ = f.GetCellValue(sheetName, "D3")
			targetapp.UnitName = trim(targetapp.UnitName)

			if sheetName=="71" {
				fmt.Print()
			}
			var tasks []Task
			var ti = 6
			for ; ; ti++ {
				var task Task
				cellB3, _ := f.GetCellValue(sheetName, colTask[0]+strconv.Itoa(ti))
				if cellB3 == "金额合计" || cellB3 == "" {
					targetapp.TotalAmt, _ = f.GetCellValue(sheetName, "F"+strconv.Itoa(ti))
					targetapp.TotalCzAmt, _ = f.GetCellValue(sheetName, "G"+strconv.Itoa(ti))
					targetapp.TotalOtherAmt, _ = f.GetCellValue(sheetName, "H"+strconv.Itoa(ti))
					break
				}

				for _, col := range colTask {

					celvalue, _ := f.GetCellValue(sheetName, col+strconv.Itoa(ti))

					targetapp.Id = sheetName
					switch col {
					case colTask[0]:
						task.TaskName = trim(celvalue)
					case colTask[1]:
						task.TaskContext = trim(celvalue)
					case colTask[2]:
						task.TaskAmt = trim(celvalue)
					case colTask[3]:
						task.TaskCzAmt = trim(celvalue)
					case colTask[4]:
						task.TaskOtherAmt = trim(celvalue)

					}
					task.Ind = ti - 5

				}
				tasks = append(tasks, task)

			}
			targetapp.Alltarget, _ = f.GetCellValue(sheetName, colTask[0]+strconv.Itoa(ti+1))

			targetapp.Alltarget = trim(targetapp.Alltarget)

			var targets []Target
			for ti = ti + 3; ti <= rows.TotalRow; ti++ {
				var target Target
				for _, col := range colTarget {
					

					celvalue, _ := f.GetCellValue(sheetName, col+strconv.Itoa(ti))
					if celvalue != "" || strings.Index(celvalue, "备注") == -1 {
						targetapp.Id = sheetName
						switch col {
						case colTarget[0]:
							target.FirsttTarget = trim(celvalue)
						case colTarget[1]:
							target.SecondTarget = trim(celvalue)
						case colTarget[2]:
							target.ThirdTarget = trim(celvalue)
						case colTarget[3]:
							target.ThirdTargetValue = trim(celvalue)

						}
						if target != (Target{}) {
							target.TargetappId = targetapp.Id
						}
					}

				}
				if target != (Target{}) {
					targets = append(targets, target)
				}

			}
			targetapp.Targets = targets
			targetapp.Tasks = tasks

			if targetapp.Id != "" {
				// targetapp.toString()
				list = append(list, targetapp)
			}
			// fmt.Println("共有%D个部门整体支出", len(list))
		}
	}

	save(list)
}
func trim(s string) string {
	temp1 := strings.ReplaceAll(s, " ", "")
	temp := strings.ReplaceAll(temp1, "\n", "")
	return temp
}

type Targetapp struct {
	Id string `json:"id"`
	// ProjectName    string   `json:"projectName"`
	UnitName      string `json:"unitName"`
	TotalAmt      string `json:totalamt`
	TotalCzAmt    string `json:totalczamt`
	TotalOtherAmt string `json:totalotheramt`
	Alltarget     string `json:"alltarget"`
	Tasks         []Task `json:"tasks"`
	Targets       []Target
}
type Target struct {
	TargetappId      string `json:"targetappId"`
	FirsttTarget     string `json:"firsttTarget"`
	SecondTarget     string `json:"secondTarget"`
	ThirdTarget      string `json:"thirdTarget"`
	ThirdTargetValue string `json:"thirdTargetValue"`
}
type Task struct {
	TaskName     string `json:"taskname"`
	TaskContext  string `json:"taskcontext"`
	TaskAmt      string `json:taskamt`
	TaskCzAmt    string `json:taskczamt`
	TaskOtherAmt string `json:taskotheramt`
	Ind          int    `json:ind`
}

func TaskstoString(ts []Task) string {
	data, _ := json.Marshal(ts)
	return string(data)
}

func (t Targetapp) toString() {
	// fmt.Println(t.projectName)
	// fmt.Println(t.unitName)
	// fmt.Println(t.totalAmt + "--" + t.totalAmt_cz + "---" + t.totalAmt_other)
	// fmt.Println(t.alltarget)
	// fmt.Println(t.targets)
	data, err := json.Marshal(t)
	if err != nil {
		// fmt.Println(err)
	}
	fmt.Println(string(data))
}

func save(targetapps []Targetapp) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pems?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
	conn, err := db.Begin()
	if err != nil {
		return
	}
	//fmt.Println(db.Ping())  检查是否连接成功数据库

	for i := 0; i < len(targetapps); i++ {

		fmt.Printf(targetapps[i].Id+"\t")
		for t := 0; t < len(targetapps[i].Targets); t++ {
			_, err1 := db.Exec("insert into t_target_1(targetappId,firsttarget,secondtarget,thirdtarget,thirdtargetvalue)  values(?,?,?,?,?)", targetapps[i].Targets[t].TargetappId, targetapps[i].Targets[t].FirsttTarget, targetapps[i].Targets[t].SecondTarget, targetapps[i].Targets[t].ThirdTarget, targetapps[i].Targets[t].ThirdTargetValue)
			if err1 != nil {
				fmt.Printf("插入指标表失败:[%v]", err.Error())
				return
			}
		}
		_, err := db.Exec("insert into t_targetapp_1(id,unitname,totalamt,totalamt_cz,totalamt_other,alltarget,taskcontext)  values(?,?,?,?,?,?,?)", targetapps[i].Id, targetapps[i].UnitName, targetapps[i].TotalAmt, targetapps[i].TotalCzAmt, targetapps[i].TotalOtherAmt, targetapps[i].Alltarget, TaskstoString(targetapps[i].Tasks))
		if err != nil {
			fmt.Printf("插入目标表失败:[%v]", err.Error())
			return
		}
		// fmt.Println(result)
	}
	conn.Commit()

}
