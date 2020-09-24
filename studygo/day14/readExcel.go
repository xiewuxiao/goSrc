package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"database/sql"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	// f, err := excelize.OpenFile("./【处理】附件3-2020年省级专项资金绩效目标表.xlsx")
	f, err := excelize.OpenFile("./【处理】附件2-2020年省级部门预算归类项目绩效目标表-new.xlsx")
	//【处理】附件2-2020年省级部门预算归类项目绩效目标表.xlsx
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(firstTarget)
	// fmt.Println(firstTarget1)
	var colFirst = []string{"B", "C", "D", "E"}
	// var index = 10
	var targetapp Targetapp
	var list = make([]Targetapp, 0)
	for _, sheetName := range f.GetSheetMap() {

		rows, _ := f.Rows(sheetName)
		if sheetName != "目录" {
			targetapp.ProjectName, _ = f.GetCellValue(sheetName, "D3")
			targetapp.ProjectName = trim(targetapp.ProjectName)
			targetapp.UnitName, _ = f.GetCellValue(sheetName, "D4")
			targetapp.UnitName = trim(targetapp.UnitName)
			targetapp.TotalAmt, _ = f.GetCellValue(sheetName, "E5")
			targetapp.TotalAmt = trim(targetapp.TotalAmt)
			targetapp.TotalAmt_cz, _ = f.GetCellValue(sheetName, "E6")
			targetapp.TotalAmt_cz = trim(targetapp.TotalAmt_cz)
			targetapp.TotalAmt_other, _ = f.GetCellValue(sheetName, "E7")
			targetapp.TotalAmt_other = trim(targetapp.TotalAmt_other)
			targetapp.Alltarget, _ = f.GetCellValue(sheetName, "B9")
			targetapp.Alltarget = trim(targetapp.Alltarget)

			var targets []Target
			for index := 11; index <= rows.TotalRow; index++ {
				var target Target
				for _, col := range colFirst {

					celvalue, _ := f.GetCellValue(sheetName, col+strconv.Itoa(index))
					if celvalue != "" || strings.Index(celvalue, "备注") == -1 {
						targetapp.Id = sheetName
						switch col {
						case colFirst[0]:
							target.FirsttTarget = trim(celvalue)
						case colFirst[1]:
							target.SecondTarget = trim(celvalue)
						case colFirst[2]:
							target.ThirdTarget = trim(celvalue)
						case colFirst[3]:
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
			//targetapp.toString()

			list = append(list, targetapp)
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
	Id             string `json:"id"`
	ProjectName    string `json:"projectName"`
	UnitName       string `json:"unitName"`
	TotalAmt       string `json:"totalAmt"`
	TotalAmt_cz    string `json:"totalAmt_cz"`
	TotalAmt_other string `json:"totalAmt_other"`
	Alltarget      string `json:"alltarget"`
	Targets        []Target
}
type Target struct {
	TargetappId      string `json:"targetappId"`
	FirsttTarget     string `json:"firsttTarget"`
	SecondTarget     string `json:"secondTarget"`
	ThirdTarget      string `json:"thirdTarget"`
	ThirdTargetValue string `json:"thirdTargetValue"`
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

	db, err := sql.Open("postgres", "postgres://dbu610000cz:Yg127329798@43.254.3.235:25308/postgres?sslmode=verify-full")
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

		fmt.Println(targetapps[i].Id)
		for t := 0; t < len(targetapps[i].Targets); t++ {
			_, err1 := db.Exec("insert into t_target(targetappId,firsttarget,secondtarget,thirdtarget,thirdtargetvalue)  values(?,?,?,?,?)", targetapps[i].Targets[t].TargetappId, targetapps[i].Targets[t].FirsttTarget, targetapps[i].Targets[t].SecondTarget, targetapps[i].Targets[t].ThirdTarget, targetapps[i].Targets[t].ThirdTargetValue)
			if err1 != nil {
				fmt.Printf("插入指标表失败:[%v]", err.Error())
				return
			}
		}
		_, err := db.Exec("insert into t_targetapp(id,projectname,unitname,totalamt,totalamt_cz,totalamt_other,alltarget)  values(?,?,?,?,?,?,?)", targetapps[i].Id, targetapps[i].ProjectName, targetapps[i].UnitName, targetapps[i].TotalAmt, targetapps[i].TotalAmt_cz, targetapps[i].TotalAmt_other, targetapps[i].Alltarget)
		if err != nil {
			fmt.Printf("插入目标表失败:[%v]", err.Error())
			return
		}
		// fmt.Println(result)
	}
	conn.Commit()

}

func save_for_mysql(targetapps []Targetapp) {

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

		fmt.Println(targetapps[i].Id)
		for t := 0; t < len(targetapps[i].Targets); t++ {
			_, err1 := db.Exec("insert into t_target(targetappId,firsttarget,secondtarget,thirdtarget,thirdtargetvalue)  values(?,?,?,?,?)", targetapps[i].Targets[t].TargetappId, targetapps[i].Targets[t].FirsttTarget, targetapps[i].Targets[t].SecondTarget, targetapps[i].Targets[t].ThirdTarget, targetapps[i].Targets[t].ThirdTargetValue)
			if err1 != nil {
				fmt.Printf("插入指标表失败:[%v]", err.Error())
				return
			}
		}
		_, err := db.Exec("insert into t_targetapp(id,projectname,unitname,totalamt,totalamt_cz,totalamt_other,alltarget)  values(?,?,?,?,?,?,?)", targetapps[i].Id, targetapps[i].ProjectName, targetapps[i].UnitName, targetapps[i].TotalAmt, targetapps[i].TotalAmt_cz, targetapps[i].TotalAmt_other, targetapps[i].Alltarget)
		if err != nil {
			fmt.Printf("插入目标表失败:[%v]", err.Error())
			return
		}
		// fmt.Println(result)
	}
	conn.Commit()

}
