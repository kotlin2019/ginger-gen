package module

import (
	"fmt"
	"github.com/gofuncchan/ginger-gen/util"
)

func PrintHandlerTips(moduleName string) string {
	header := "binding this handler function to app router. \n"
	footer := "According to the http method what you need,copy the code to router/router.go.\n"
	examples := "For example:\n"

	camelModuleName := util.CamelString(moduleName)
	examples += fmt.Sprintf(`
		r.POST("/%ss", handler.Get%sList)
		// 创建
		r.POST("/%s", handler.Create%s)
		// 获取
		r.GET("/%s/:id", handler.Get%s)
		// 修改
		r.PUT("/%s", handler.Update%s)
		// 删除
		r.DELETE("/%s", handler.Delete%s)
	`, moduleName, camelModuleName, moduleName, camelModuleName, moduleName, camelModuleName, moduleName, camelModuleName, moduleName, camelModuleName)

	return header + examples + footer

}
