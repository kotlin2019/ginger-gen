package schema

import "strings"

const (
	cUnderScore = "_"
)

// 转换转换下划线，用于包命名
func Convert2PackageName(tableName string) string {
	return strings.ToLower(strings.Replace(tableName, cUnderScore, "", -1))
}

// 转换下划线字符串为驼峰格式，用于go 结构体命名
func ConvertUnderScoreToCamel(name string) string {
	arr := strings.Split(name, cUnderScore)
	for i := 0; i < len(arr); i++ {
		arr[i] = lintName(strings.Title(arr[i]))
	}
	return strings.Join(arr, "")
}
