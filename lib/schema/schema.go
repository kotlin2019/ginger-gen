package schema

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"text/template"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

const (
	cDefaultTable = "COLUMNS"
	// cTimeFormat   = "2006-01-02 15:04:05"
)

type ColumnSlice []Column

// 读取数据库schema表结构表，获取表结构的列信息
func readTableStruct(db *sql.DB, tableName string, dbName string) (ColumnSlice, error) {
	var where = map[string]interface{}{
		"TABLE_NAME":   tableName,
		"TABLE_SCHEMA": dbName,
	}
	var selectFields = []string{"COLUMN_NAME", "COLUMN_TYPE", "COLUMN_COMMENT"}
	cond, vals, err := builder.BuildSelect(cDefaultTable, where, selectFields)
	if nil != err {
		return nil, err
	}
	rows, err := db.Query(cond, vals...)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var ts ColumnSlice
	scanner.SetTagName("json")
	err = scanner.Scan(rows, &ts)
	if nil != err {
		return nil, err
	}
	return ts, nil
}

// 根据表列结构信息生成go结构体源码
func createStructSourceCode(cols ColumnSlice, tableName string) (io.Reader, string, error) {
	structName := ConvertUnderScoreToCamel(tableName)
	fillData := sourceCode{
		packageName: tableName,
		StructName:  structName,
		TableName:   tableName,
		FieldList:   make([]sourceColumn, len(cols)),
	}
	for idx, col := range cols {
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		fillData.FieldList[idx] = sourceColumn{
			Name:      col.GetName(),
			Type:      colType,
			StructTag: fmt.Sprintf("`ddb:\"%s\" json:\"%s\"`", col.Name, col.Name),
		}
	}
	var buff bytes.Buffer
	err := template.Must(template.New("struct").Parse(codeTemplate)).Execute(&buff, fillData)
	if nil != err {
		return nil, "", err
	}
	return &buff, structName, nil
}

// 源码信息结构体
type sourceCode struct {
	packageName string
	StructName  string
	TableName   string
	FieldList   []sourceColumn
}

// 表列信息结构体
type sourceColumn struct {
	Name      string
	Type      string
	StructTag string
}
