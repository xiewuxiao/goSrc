// 使用go语言读取文件
// 该种方式直接将文件内容输出为字符串
package main

import (
	"bufio"
	"regexp"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "time"
	// "unsafe"
	// "runtime/pprof"
	// "runtime"
)

func main() {
	
	// t1 :=time.Now()
	readFileType1()
	// elapsed1 :=time.Since(t1)
	// t2 :=time.Now()
	// readFileType2()
	// elapsed2 :=time.Since(t2)
	
	
	// fmt.Printf("使用第一种试耗时%s\n",elapsed1)
	// fmt.Printf("使用第二种试耗时%s\n",elapsed2)
	// time.Sleep(30*time.Second)
	
}

//该方式通过读取ios.txt文件，将全部内容加载进内存，然后再将内容输出
func readFileType1() {
	fileBytes, err := ioutil.ReadFile("ios.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(fileBytes))
	// fileContent :=string(fileBytes)
	regStr := "Host\\:.*"
	reg :=regexp.MustCompile(regStr)
	allPattern :=reg.FindAll(fileBytes,-1)
	
	resultFile,err :=os.Create("result.txt")
	if err!=nil {
		fmt.Println(resultFile)
	}
	// fmt.Println(len(allPattern))
	for si:=0;si<len(allPattern);si++{
		resultFile.Write(allPattern[si])
	}
	resultFile.Close()

}
func readFileType2() {
	f, err := os.Open("ios.txt")
	if err != nil {
		fmt.Println(err)
	}
	bfreader := bufio.NewReader(f)
	for {
		line, _, err := bfreader.ReadLine()
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			continue
		}
		fmt.Println(string(line))
		

	}
	fmt.Println("第二种方式：")

}
