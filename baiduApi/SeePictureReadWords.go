package main

import (
	"fmt"
	"io/ioutil"

	token "com.todaytech.ben/baiduApi/token"
)

const PICTURE_FOLD = "./图片/"
const WORDS_FOLD = "./文字/"

func main() {
	var accessToken = token.GetToken()
	fmt.Println("获取到的token为：\n" + accessToken)
	if accessToken == "" {
		return
	}
	readFiles(accessToken)

}
func readFiles(accessToken string) {
	files, err := ioutil.ReadDir(PICTURE_FOLD)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		var imageBase64Content = token.ReadFileByBytes(PICTURE_FOLD, file.Name())
		file.Name()
		token.SeePirctureGetWords(accessToken, imageBase64Content, file.Name(), WORDS_FOLD)
	}

}
