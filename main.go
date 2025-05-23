package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin" //导入gin框架
	"gorm.io/driver/mysql"     //mysql驱动包
	"gorm.io/gorm"             //快速操作数据库框架
)

type News struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/", Index)
	r.GET("/about", AboutMe)
	r.GET("/test", Test)
	r.POST("/cms", InsertArtticle)
	r.DELETE("/cms/:id", DeleteArtticle)
	r.Run() //默认8080端口
}
func Index(c *gin.Context) {
	c.String(200, "欢迎光临我的首页yy") //在客户端输出Sting
	c.JSON(200, "欢迎光临我的首页")     //json
}
func AboutMe(c *gin.Context) {
	c.String(200, "这是一个使用gin框架开发的网站")
}
func Test(c *gin.Context) {
	c.JSON(200, gin.H{"user": "alice", "age:": 10})
}

// 自定义 InsertArtticle 函数，用于添加一篇文章进数据库中
func InsertArtticle(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/cms?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用gorm打开连接
	if err != nil {
		panic("数据库连接失败")
	}
	var news News
	err = c.ShouldBind(&news)
	if err != nil {
		panic(err.Error())
	}
	err = db.Create(&news).Error
	if err != nil {
		c.String(500, err.Error())
		return
	}
	str := fmt.Sprintf("添加成功,记录id为:%d", news.Id)
	c.String(200, str)
}
func DeleteArtticle(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/cms?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //使用gorm打开连接
	if err != nil {
		panic("数据库连接失败")
	}
	id, _ := strconv.Atoi(c.Param("id"))
	news := News{Id: uint(id)}
	err = db.Delete(&news).Error
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "删除成功")
}
