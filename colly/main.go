package main

import (
	"embed"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

type Crow struct {
	gorm.Model
	Title       string `gorm:"type:varchar(500)" json:"title"`
	Url         string `gorm:"type:varchar(255)" json:"url"`
	Data        string `gorm:"type:varchar(255)" json:"data"`
	Description string `gorm:"type:longtext" json:"description"`
}

var (
	SqlName  string
	SqlPwd   string
	Address  string
	HttpPort string
	DbName   string
)

//go:embed "config.ini"
var config embed.FS

func init() {
	data, err := config.ReadFile("config.ini")
	file, err := ini.Load(data)
	if err != nil {
		fmt.Println("配置文件错误", err)
	}
	LoadData(file)
}

func LoadData(file *ini.File) {
	SqlName = file.Section("database").Key("SqlName").MustString("")
	SqlPwd = file.Section("database").Key("SqlPwd").MustString("")
	HttpPort = file.Section("database").Key("HttpPort").MustString("")
	Address = file.Section("database").Key("Address").MustString("")
	DbName = file.Section("database").Key("dbName").MustString("")
}

func initSql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", SqlName, SqlPwd, Address, HttpPort, DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}
	//迁移表结构。在初次使用。其余时间注释
	db.AutoMigrate(&Crow{})
}

func main() {
	var data []Crow
	initSql()
	c := colly.NewCollector()
	c.OnXML("//item", func(element *colly.XMLElement) {
		strArr := strings.Split(element.Text, "\n")
		//fmt.Println(strArr[1], strArr[2], strArr[5],strArr[4])
		t, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", strings.TrimSpace(strArr[5]))
		//fmt.Printf("%T", t.UnixNano()/1e6)
		data = append(data, Crow{Title: strings.TrimSpace(strArr[1]), Url: strings.TrimSpace(strArr[2]), Data: fmt.Sprintf("%v", t.UnixNano()/1e6), Description: strings.TrimSpace(strArr[4])})
		//fmt.Println(data)
		//for i := 1; i < len(data); i++ {
		//	fmt.Printf("%+v\n", data[i])
		//}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.oschina.net/news/rss")
	db.Create(&data)
}
