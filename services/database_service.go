package services

import (
	"github.com/beego/beego/v2/client/orm"
)

// DbVersion struct
type DbVersion struct {
	DbVersion string
}

// DatabaseService struct
type DatabaseService struct {
}

// GetMysqlVersion 获取mysql的版本
func (*DatabaseService) GetMysqlVersion() string {
	var dbVersion DbVersion
	err := orm.NewOrm().Raw("select VERSION() as db_version").QueryRow(&dbVersion)
	if err != nil {
		return "not found."
	}
	return dbVersion.DbVersion
}

// GetTableStatus 获取所有数据表的状态
func (ds *DatabaseService) GetTableStatus() ([]map[string]string, int) {
	var maps []orm.Params
	var resultMaps []map[string]string
	o := orm.NewOrm()
	affectRows, err := o.Raw("SHOW TABLE STATUS").Values(&maps)

	if affectRows > 0 && err == nil {
		for _, item := range maps {
			resultMaps = append(resultMaps, map[string]string{
				"name":        ds.nil2String(item["Name"]),
				"comment":     ds.nil2String(item["Comment"]),
				"engine":      ds.nil2String(item["Engine"]),
				"collation":   ds.nil2String(item["Collation"]),
				"data_length": ds.nil2String(item["Data_length"]),
				"create_time": ds.nil2String(item["Create_time"]),
				"update_time": ds.nil2String(item["Update_time"]),
			})
		}
	}

	return resultMaps, int(affectRows)
}

// OptimizeTable 优化数据表
func (*DatabaseService) OptimizeTable(tableName string) bool {
	o := orm.NewOrm()
	_, err := o.Raw("OPTIMIZE TABLE `" + tableName + "`").Exec()
	if err == nil {
		return true
	}
	return false
}

// RepairTable 修复数据表
func (*DatabaseService) RepairTable(tableName string) bool {
	o := orm.NewOrm()
	_, err := o.Raw("REPAIR TABLE `" + tableName + "`").Exec()
	if err == nil {
		return true
	}
	return false
}

// GetFullColumnsFromTable 获取数据表的所有字段
func (ds *DatabaseService) GetFullColumnsFromTable(tableName string) []map[string]string {
	var maps []orm.Params
	var resultMaps []map[string]string
	o := orm.NewOrm()
	affectRows, err := o.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`").Values(&maps)

	if affectRows > 0 && err == nil {
		for _, item := range maps {
			resultMaps = append(resultMaps, map[string]string{
				"name":       ds.nil2String(item["Field"]),
				"type":       ds.nil2String(item["Type"]),
				"collation":  ds.nil2String(item["Collation"]),
				"null":       ds.nil2String(item["Null"]),
				"key":        ds.nil2String(item["Key"]),
				"default":    ds.nil2String(item["Default"]),
				"extra":      ds.nil2String(item["Extra"]),
				"privileges": ds.nil2String(item["Privileges"]),
				"comment":    ds.nil2String(item["Comment"]),
			})
		}
	}

	return resultMaps
}

// nil2String interface 转换 为string
func (*DatabaseService) nil2String(val interface{}) string {
	if val == nil {
		return ""
	}
	return val.(string)
}
