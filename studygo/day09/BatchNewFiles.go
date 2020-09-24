package main
import (
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(){
	for i:=52;i<=112;i++ {
		f:=excelize.NewFile()
		f.NewSheet("Sheet2")
		f.NewSheet("Sheet3")
		f.NewSheet("Sheet3")
		f.SaveAs(strconv.Itoa(i)+".xlsx")
		
	}
}