package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	//"time"

	_ "github.com/go-sql-driver/mysql"
)

// sql 查询
type Post struct {
	postID     string
	postCode   string
	postName   string
	postSort   string
	status     string
	createTime string
}

func (p *Post) toString() {
	fmt.Print("{postID=" + p.postID + ",")
	fmt.Print("postCode=" + p.postCode + ",")
	fmt.Print("postName=" + p.postName + ",")
	fmt.Print("postSort=" + p.postSort + ",")
	fmt.Print("status=" + p.status + ",")
	fmt.Print("createTime=" + p.createTime + "}\n")
}

func main() {
	//sql查询
	//search()
	MyUpdate()
}

func search() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ruoyi?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	//fmt.Println(db.Ping())  检查是否连接成功数据库
	postList := make([]Post, 0)
	rows, err := db.Query("select post_id,post_code,post_name,post_sort,status,create_Time from sys_post")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	for rows.Next() {
		var post Post
		e := rows.Scan(&post.postCode, &post.postID, &post.postName, &post.postSort, &post.status, &post.createTime)
		if e != nil {
			fmt.Println(json.Marshal(post))
		}
		postList = append(postList, post)
		i++
	}
	rows.Close()
	postList[0].toString()
}

func MyUpdate(){
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ruoyi?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	result,err1 :=db.Exec("update sys_post set remark='test go ' where post_id=1")
	if err1 != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)

}
