package main

import (
	"fmt"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	postID     string `gorm:"column:post_id"`
	postCode   string `gorm:"column:post_code"`
	postName   string `gorm:"column:post_name"`
	postSort   string `gorm:"column:post_sort"`
	status     string `gorm:"column:status"`
	createTime string `gorm:"column:create_time"`
}

func (Post) TableName() string {
	return "sys_post"
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
	db, err := gorm.Open("mysql", "root:root@/ruoyi?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	//db.LogMode(true)
	defer db.Close()
	var post Post
	//fmt.Println(db.HasTable(post))

	db.Debug().Where("post_id = ?", "3").Find(&post)
	fmt.Println(post)
}
