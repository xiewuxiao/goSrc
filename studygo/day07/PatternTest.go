package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"regexp"
)

func main() {
	file, err := ioutil.ReadFile("ios.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(file))
	getPatternString(string(file))
}

func getPatternString(old string)  {
	//默认是贪婪模式；在量词后面直接加上一个问号？就是非贪婪模式。
	replaceReg := "Host\\:.*"
	reg := regexp.MustCompile(replaceReg)
	//fmt.Printf("%q\n", reg.FindAllString(old, -1))
	//return reg
	patternStrings := reg.FindAllString(old, -1)
	file, err := os.Create("result.txt")
    if err != nil {
        fmt.Println("create file err", err)
    }
    
    
	for i:=0;i<len(patternStrings);i++{
		
		
		file.WriteString(strings.ReplaceAll(patternStrings[i],"Host: ",""))
	}
	file.Close()
}