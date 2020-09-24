package database_service

import "github.com/astaxie/beego/orm"

type DbVersion struct {
	DbVersion string
}

//获取mysql的版本
func GetMysqlVersion() string {
	var dbVersion DbVersion
	error := orm.NewOrm().Raw("select VERSION() as db_version").QueryRow(&dbVersion)
	if error != nil{
		return "not found."
	}
	return dbVersion.DbVersion
}
