package main

import (
	"fmt"
	"os"
	"strconv"

	"com.todaytech.ben/studygo/day10/simplemath"
)

//Usage 消息提示
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("The commands are :\n\tAddition of two values.\n\t Square root of a non-negative value.")
}

func main() {
	args := os.Args
	fmt.Println(len(args))
	if args == nil || len(args) < 4 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE:calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE:calc add <integer1><integer2> ")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("result:", ret)
	}
}
