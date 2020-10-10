package services

import "github.com/astaxie/beego/orm"

type DbVersion struct {
	DbVersion string
}

type DatabaseService struct {
}

//获取mysql的版本
func (*DatabaseService) GetMysqlVersion() string {
	var dbVersion DbVersion
	error := orm.NewOrm().Raw("select VERSION() as db_version").QueryRow(&dbVersion)
	if error != nil {
		return "not found."
	}
	return dbVersion.DbVersion
}
