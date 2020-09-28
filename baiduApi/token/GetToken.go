package token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	entity "com.todaytech.ben/baiduApi/entity"
)

func GetToken() string {
	URL, err := url.Parse("https://aip.baidubce.com/oauth/2.0/token")
	if err != nil {
		fmt.Printf("requestURL parse failed, err:[%s]", err.Error())
		return ""
	}

	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", "xDxVQ0mPENVuSZRkHHKysMsl")
	params.Set("client_secret", "X4OtnpmdC5Yj9wYZf1nCrGBvNaChbTM3")
	URL.RawQuery = params.Encode()

	requestURL := URL.String()
	// fmt.Printf("requestURL:[%s]\n", requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("请求错误:[%s]", err.Error())
		return ""
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)

	// fmt.Printf("返回码：%d\n", resp.StatusCode)
	// fmt.Printf("返回数据：%s\n", string(bodyContent))
	var token entity.Token
	json.Unmarshal([]byte(bodyContent), &token)
	// 遍历Json
	return token.ACCESS_TOKEN
}

func SeePirctureGetWords(token string, imageBase64Content string, fileName string, WORDS_FOLD string) {
	client := &http.Client{}

	var URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic?access_token=" + token
	paramData := url.Values{}
	//循环的部分主要在这里
	paramData.Set("image", imageBase64Content)
	//方式1
	// resp, err := client.PostForm(URL, paramData)
	//方式2
	requestPost, _ := http.NewRequest("POST", URL, strings.NewReader(paramData.Encode()))
	requestPost.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(requestPost)

	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Printf("请求错误:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("请求返回码为：%d\n", resp.StatusCode)
	fmt.Println(string(bodyContent))

	var result entity.Result
	json.Unmarshal([]byte(bodyContent), &result)
	////把文件写入文件
	os.Mkdir(WORDS_FOLD, os.FileMode(int(0777)))
	f, err := os.Create(WORDS_FOLD + strings.Split(fileName, ".")[0] + ".txt")

	defer f.Close()
	for word := range result.WORDS_RESULT {
		perline := result.WORDS_RESULT[word].WORDS
		f.WriteString(perline + "\n")
	}

}

/**
@param picture 图片的地址+文件名
**/
func ReadFileByBytes(PICTURE_FOLD string, picture string) string {
	data, _ := ioutil.ReadFile(PICTURE_FOLD + picture)
	base64Content := base64.StdEncoding.EncodeToString(data)
	return base64Content
}
