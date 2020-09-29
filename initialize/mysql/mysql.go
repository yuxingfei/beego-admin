package mysql

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//注册mysql
func init()  {
	err := orm.RegisterDriver("mysql",orm.DRMySQL)
	if err != nil{
		beego.Error("mysql register driver error:",err)
	}

	//dataSource := "root:root@tcp(127.0.0.1:3306)/test"
	mysqlConfig,err := beego.AppConfig.GetSection("mysql")
	if err != nil{
		beego.Error("mysql get config fail! error:",err)
	}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysqlConfig["username"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["database"],
		)

	err = orm.RegisterDataBase("default","mysql",dataSource)
	if err != nil{
		beego.Error("mysql register database error:",err)
	}
}
