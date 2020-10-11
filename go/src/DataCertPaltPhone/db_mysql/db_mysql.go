package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func Connect() {
	//定于config变量接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目的应用昵称：", appName)
	port, err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请检查后重试")
	}
	fmt.Println("应用监听端口：", port)

	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	db, err := sql.Open(driver, dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil { //err不为nil,表示连接数据库时出现错误，程序就在此中断就不用在执行了。
		//早发现，早解决
		//panic:让程序进入一种恐慌状态，程序会终止执行。
		panic("数据库连接打开失败，请重试")
	}
	Db = db
	fmt.Println(db)
}
