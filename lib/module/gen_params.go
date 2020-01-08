package module

import (
	"fmt"
	libSchema "github.com/gofuncchan/ginger-gen/lib/schema"
)

// 生成handler GetList params struct
func GenGetListParams() []HandlerParam {
	params := []HandlerParam{
		{
			Name:      "page",
			Type:      "uint",
			StructTag: "`json:\"page\" validate:\"required,numeric\"`",
			Comment:	"// 页数",
		},
		{
			Name:      "Count",
			Type:      "uint",
			StructTag: "`json:\"count\" validate:\"required,numeric\"`",
			Comment:	"// 每次获取数目",
		},
		{
			Name:      "Order",
			Type:      "map[string]int",
			StructTag: "`json:\"order\"`",
			Comment:	"// 排序",
		},
	}
	return params
}

// 生成handler GetOne params struct
func GenGetOneParams() []HandlerParam {
	params := make([]HandlerParam, 0)
	params = append(params, HandlerParam{
		Name:      "Id",
		Type:      "int64",
		StructTag: "`form:\"id\" validate:\"required,numeric\"`",
		Comment:	"// 根据id获取",
	})
	return params
}

// 生成handler Create params struct
func GenCreateParams(cols libSchema.ColumnSlice) []HandlerParam {
	params := make([]HandlerParam, len(cols))
	for idx, col := range cols {
		if col.Name == "id" || col.Name == "create_at" || col.Name == "update_at" {
			continue
		}
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		col := HandlerParam{
			Name: col.GetName(),
			Type: colType,
			Comment: "// "+col.Comment,
		}

		switch colType {
		case libSchema.CTypeInt, libSchema.CTypeInt64, libSchema.CTypeUInt, libSchema.CTypeInt8, libSchema.CTypeUInt64, libSchema.CTypeTime, libSchema.CTypeFloat64:
			col.StructTag = fmt.Sprintf("`json:\"%s\" validate:\"required,numeric\"`", col.Name)
		default:
			col.StructTag = fmt.Sprintf("`json:\"%s\" validate:\"required\"`", col.Name)
		}

		params[idx] = col
	}

	return params
}

// 生成handler Update params struct
func GenUpdateParams(cols libSchema.ColumnSlice) []HandlerParam {

	params := make([]HandlerParam, len(cols))
	for idx, col := range cols {
		if col.Name == "create_at" || col.Name == "update_at" {
			continue
		}
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		col := HandlerParam{
			Name: col.GetName(),
			Type: colType,
			Comment: "// "+col.Comment,
		}

		switch colType {
		case libSchema.CTypeInt, libSchema.CTypeInt64, libSchema.CTypeUInt, libSchema.CTypeInt8, libSchema.CTypeUInt64, libSchema.CTypeTime, libSchema.CTypeFloat64:
			col.StructTag = fmt.Sprintf("`json:\"%s\" validate:\"required,numeric\"`", col.Name)
		default:
			col.StructTag = fmt.Sprintf("`json:\"%s\" validate:\"required\"`", col.Name)
		}

		params[idx] = col
	}

	return params
}

// 生成handler Delete params struct
func GenDeleteParams() []HandlerParam {
	params := make([]HandlerParam, 0)
	params = append(params, HandlerParam{
		Name:      "Id",
		Type:      "int64",
		StructTag: "`json:\"id\" validate:\"required,numeric\"`",
		Comment: 	"// 根据id删除",
	})
	return params
}
