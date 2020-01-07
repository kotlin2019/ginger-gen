package schema

import (
	"io"
)

// schema转换struct的公用调用方法，用于schema命令的直接调用
func GenSchemaStruct(w io.Writer, schemaArgs *SchemaArgs) (string, error) {
	db, err := GetDBInstance(&DBConfig{
		Host:     schemaArgs.Host,
		User:     schemaArgs.User,
		Password: schemaArgs.Password,
		Port:     schemaArgs.Port,
		DBName:   "information_schema",
	})
	if nil != err {
		return "", err
	}

	// 获取数据库指定表结构
	cols, err := readTableStruct(db, schemaArgs.Table, schemaArgs.Database)
	if nil != err {
		return "", err
	}
	// 创建go struct 源码
	r, structName, err := createStructSourceCode(cols, schemaArgs.Table)
	if nil != err {
		return "", err
	}
	_, err = io.Copy(w, r)
	return structName, err
}

// 从数据库表导出字段的方法，用于module命令的调用
func GetSchemaField(schemaArgs *SchemaArgs) (ColumnSlice, error) {
	db, err := GetDBInstance(&DBConfig{
		Host:     schemaArgs.Host,
		User:     schemaArgs.User,
		Password: schemaArgs.Password,
		Port:     schemaArgs.Port,
		DBName:   "information_schema",
	})
	if nil != err {
		return nil, err
	}

	// 获取数据库指定表结构
	cols, err := readTableStruct(db, schemaArgs.Table, schemaArgs.Database)
	if nil != err {
		return nil, err
	}

	return cols, nil
}
