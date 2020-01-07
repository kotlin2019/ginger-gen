package module

type ArgsForMap struct {
	Name    string
}

type ModuleModelTmplData struct {
	RootPackageName  string
	ModuleName       string
	CamelModuleName  string
	CreateArgs       string
	CreateArgsForMap []ArgsForMap
}

const ModuleModelTmplCode = `
package {{ .ModuleName }}Model


import(
	"{{ .RootPackageName }}/dao/mysql"
	"{{ .RootPackageName }}/dao/mysql/schema"
    "{{ .RootPackageName }}/util/e"
	"time"
)


/*
获取列表
@param offset uint 起始位移
@param count uint 获取行数
@param sort string 字段排序 如"a desc,b asc,..."
@param filter map[string]interface{} 字段过滤条件
*/
func Get{{ .CamelModuleName }}List(offset, count uint, sort string, filter map[string]interface{}) ([]*schema.{{ .CamelModuleName }}, error) {
	where := make(map[string]interface{}, 0)
	if count != 0 {
		where["_limit"] = []uint{offset, count}
	}
	if sort != "" {
		where["_orderby"] = sort
	}
	if filter != nil {
		for k, v := range filter {
			where[k] = v
		}
	}

	// 接收查询结果
	{{ .ModuleName }}Results := make([]*schema.{{ .CamelModuleName }}, 0)
	err := mysql.GetMulti(schema.{{ .CamelModuleName }}TableName, where, nil, &{{ .ModuleName }}Results)
	if !e.Em(err) {
		return nil, err
	}

	return {{ .ModuleName }}Results, nil
}

/*
获取单个
*/
func Get{{ .CamelModuleName }}InfoById(id int64) (*schema.{{ .CamelModuleName }}, error) {
	where := map[string]interface{}{
		"id": id,
	}
	{{ .ModuleName }}Result := new(schema.{{ .CamelModuleName }})
	err := mysql.GetOne(schema.{{ .CamelModuleName }}TableName, where, nil, {{ .ModuleName }}Result)
	if !e.Em(err) {
		return nil, err
	}
	return {{ .ModuleName }}Result, nil
}

/*
创建
TODO 此处代码自动生成，请删减不必要的字段

*/
func Create{{ .CamelModuleName }}( {{ .CreateArgs }}) int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		{{- range .CreateArgsForMap }}
			"{{ .Name }}": {{ .Name }},
		{{- end}}
		"create_at":   time.Now(),
		"update_at":   time.Now(),
	})

	id, err := mysql.Insert(schema.{{ .CamelModuleName }}TableName, data)
	if !e.Em(err) {
		return -1
	}
	return id
}






/*
根据id修改角色字段信息
*/
func Update{{ .CamelModuleName }}Field(id int64,data map[string]interface{}) (int64,error) {
	where := map[string]interface{}{
		"id":id,
	}
	RowsAffected, err := mysql.Update(schema.{{ .CamelModuleName }}TableName, where, data)
	if !e.Em(err) {
		return -1, err
	}

	return RowsAffected,nil
}


/*
删除
*/
func Delete{{ .CamelModuleName }}(id int64) int64 {
	where := map[string]interface{}{
		"id": id,
	}

	// 删除表信息
	deleteId, err := mysql.Delete(schema.{{ .CamelModuleName }}TableName, where)
	if !e.Em(err) {
		return -1
	}

	return deleteId
}

`
