package schema

import (
	"database/sql"
	"github.com/gofuncchan/ginger-forge/lib/util"
	"io"
)

// schema转换struct的公用调用方法，用于schema和dao命令的直接调用
func GenSchema(w io.Writer, schemaArgs *SchemaArgs) (string, error) {
	db, err := util.GetDBInstance(&util.DBConfig{
		Host:     schemaArgs.Host,
		User:     schemaArgs.User,
		Password: schemaArgs.Password,
		Port:     schemaArgs.Port,
		DBName:   "information_schema",
	})
	if nil != err {
		return "", err
	}
	return genSchemaToStruct(w, db, schemaArgs.Table, schemaArgs.Database)
}

// 将数据库表结构转化为go结构体代码
func genSchemaToStruct(w io.Writer, db *sql.DB, tableName, dbName string) (string, error) {
	// 获取数据库指定表结构
	cols, err := readTableStruct(db, tableName, dbName)
	if nil != err {
		return "", err
	}
	// 创建go struct 源码
	r, structName, err := createStructSourceCode(cols, tableName)
	if nil != err {
		return "", err
	}
	_, err = io.Copy(w, r)
	return structName, err
}
