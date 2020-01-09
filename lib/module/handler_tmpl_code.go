package module

type HandlerParam struct {
	Name      string
	Type      string
	StructTag string
	Comment   string
}

type UpdateArgs struct {
	Name    string
	GetName string
}

type ModuleHandlerTmplData struct {
	RootPackageName string
	ModuleName      string
	CamelModuleName string
	CreateArgs      string
	UpdateArgs      []UpdateArgs
	GetListParams   []HandlerParam
	GetOneParams    []HandlerParam
	CreateParams    []HandlerParam
	UpdateParams    []HandlerParam
	DeleteParams    []HandlerParam
}

//
const ModuleHandlerTmplCode = `
package handler

import(
	"github.com/gin-gonic/gin"
	"{{ .RootPackageName }}/common"
	"{{ .RootPackageName }}/model"
	"{{ .RootPackageName }}/model/{{ .ModuleName }}Model"
	"{{ .RootPackageName }}/util/e"
	"strings"
)


/*
列表
For Example:
Content-Type:"application/json"

{
	"page":1,
	"count":10,
	"order":{
		"name":-1,
		"create_at":1,
	},
	filter:[
		["name","like","%f%"],
		["create_at",">","1000"],
		["status","=",1]
	]
}

filter过滤条件：
"=",">","<","=","<=",">=","!=","<>"
"in","not in","like","not like","between","not between"

*/
type Get{{ .CamelModuleName }}ListParams struct {
	{{- range .GetListParams }}
		{{ .Name }} {{ .Type }} {{ .StructTag }} {{ .Comment }}
	{{- end}}
}

func Get{{ .CamelModuleName }}List(c *gin.Context) {
	// validate request params
	form := new(Get{{ .CamelModuleName }}ListParams)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err.Error())
		return
	}

	// set limit offset for paging
	var offset uint
	if form.Page > 1 {
		offset = (form.Page - 1) * form.Count
	}

	// generate order fields
	var sorts []string
	var orderFields string
	for k, v := range form.Order {
		if v == -1 {
			// 倒序
			sorts = append(sorts, k+" desc")
		} else if v == 1 {
			// 正序
			sorts = append(sorts, k+" asc")

		}
	}
	orderFields = strings.Join(sorts, ",")

	// generate filter fields
	filterMap, err := model.ConvertFilterToMap(form.Filter)
	if !e.Eh(err){
		common.ResponseInvalidParam(c,err.Error())
	}

	// TODO set selectFields
	var selectFields []string

	{{ .ModuleName }}List, err := {{ .ModuleName }}Model.Get{{ .CamelModuleName }}List(offset, form.Count, orderFields, filterMap, selectFields)
	if !e.Eh(err) {
		common.ResponseModelError(c, err.Error())
		return
	}

	// TODO 输出时请过滤字段

	common.ResponseOk(c, {{ .ModuleName }}List)
}


/*
获取
*/
type Get{{ .CamelModuleName }}Params struct {
	{{- range .GetOneParams }}
		{{ .Name }} {{ .Type }} {{ .StructTag }}  {{ .Comment }}
	{{- end}}
}

func Get{{ .CamelModuleName }}(c *gin.Context) {
	// validate request params
	form := new(Get{{ .CamelModuleName }}Params)
	if err := c.ShouldBindQuery(form); err != nil {
		common.ResponseInvalidParam(c,err.Error())
		return
	}

	// TODO set selectFields
	var selectFields []string

	adminInfo, err := {{ .ModuleName }}Model.Get{{ .CamelModuleName}}InfoById(form.Id, selectFields)
	if !e.Eh(err) {
		common.ResponseModelError(c,err.Error())
		return
	}

	// TODO 输出时请过滤字段

	common.ResponseOk(c, adminInfo)
}


/*
创建
*/
type Create{{ .CamelModuleName }}Params struct {
	{{- range .CreateParams }}
		{{ .Name }} {{ .Type }} {{ .StructTag }} {{ .Comment }}
	{{- end}}
}

func Create{{ .CamelModuleName }}(c *gin.Context) {
	// validate request params
	form := new(Create{{ .CamelModuleName }}Params)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err.Error())
		return
	}

	// TODO 此处代码自动生成，请删减不必要的字段

	menuId := {{ .ModuleName }}Model.Create{{ .CamelModuleName }}({{ .CreateArgs }})
	if menuId == -1 {
		common.ResponseModelError(c,"create fail,please try again")
		return
	}

	common.ResponseOk(c, gin.H{"result": "ok"})
}



/*
更新字段

// TODO 此处代码自动生成，请删减不必要的字段

*/
type Update{{ .CamelModuleName }}Params struct {
	{{- range .UpdateParams }}
		{{ .Name }} {{ .Type }} {{ .StructTag }} {{ .Comment }}
	{{- end}}
}

func Update{{ .CamelModuleName }}(c *gin.Context) {
	// validate request params
	form := new(Update{{ .CamelModuleName }}Params)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err.Error())
		return
	}

	data := map[string]interface{}{
		{{- range .UpdateArgs }}
			"{{ .Name }}": form.{{ .GetName }},
		{{- end}}
		"update_at":time.Now().Unix(),
	}
	_, err := {{ .ModuleName }}Model.Update{{ .CamelModuleName }}Field(form.Id, data)
	if !e.Eh(err) {
		common.ResponseModelError(c, err.Error())
		return
	}
}

/*
删除
*/
type Delete{{ .CamelModuleName }}Params struct {
	{{- range .DeleteParams }}
		{{ .Name }} {{ .Type }} {{ .StructTag }} {{ .Comment }}
	{{- end}}
}

func Delete{{ .CamelModuleName }}(c *gin.Context) {
	// validate request params
	form := new(Delete{{ .CamelModuleName }}Params)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err.Error())
		return
	}

	deleteId := {{ .ModuleName }}Model.Delete{{ .CamelModuleName }}(form.Id)
	if deleteId <= 0 {
		common.ResponseModelError(c,"delete admin fail")
	}

	common.ResponseOk(c, gin.H{"result": "ok"})
}


`
