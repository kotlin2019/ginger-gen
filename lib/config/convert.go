package config

import (
	"fmt"
	"github.com/gofuncchan/ginger-gen/util"
	"reflect"
)

// 嵌入map 转struct代码的容器
var InnerStructCodeList []string

// 转换map为struct代码
func ConvertMapToStructCode(structName string, configMap map[interface{}]interface{}) (structCode string, err error) {

	var fieldList string
	for k, v := range configMap {
		var line string
		fieldName := k.(string)
		fieldType := reflect.TypeOf(v)
		fmt.Println("key:", k, "fieldType:", fieldType)
		if fieldType.String() == `map[interface {}]interface {}` {
			// 转换map
			innerMapStructName := k.(string)
			innerMap := v.(map[interface{}]interface{})
			innerStructCode, err := ConvertMapToStructCode(util.CamelString(innerMapStructName), innerMap)
			InnerStructCodeList = append(InnerStructCodeList, innerStructCode)
			if err != nil {
				return "", err
			}
			fieldTag := "`yaml:\"" + fieldName + "\"`"
			line = fmt.Sprintf("%s %s \n\t", util.CamelString(innerMapStructName), fieldTag)
		} else if fieldType.String() == `[]interface {}` {
			// 转换slice,可转换[]string/[]int/[]bool/[]map[interface{}]interface{}
			innerSlice := v.([]interface{})
			switch innerSlice[0].(type) {
			case string:
				fieldTag := "`yaml:\"" + fieldName + "\"`"
				line = fmt.Sprintf("%s []string %s \n\t", util.CamelString(fieldName), fieldTag)
			case int:
				fieldTag := "`yaml:\"" + fieldName + "\"`"
				line = fmt.Sprintf("%s []int %s \n\t", util.CamelString(fieldName), fieldTag)
			case bool:
				fieldTag := "`yaml:\"" + fieldName + "\"`"
				line = fmt.Sprintf("%s []bool %s \n\t", util.CamelString(fieldName), fieldTag)
			case map[interface{}]interface{}:
				innerSliceElementStructCode, err := ConverSliceElementToStructCode(fieldName, innerSlice)
				if err != nil {
					return "", err
				}
				InnerStructCodeList = append(InnerStructCodeList, innerSliceElementStructCode)
				fieldTag := "`yaml:\"" + fieldName + "\"`"
				line = fmt.Sprintf("%s []%s %s\n\t", util.CamelString(fieldName)+"s", util.CamelString(fieldName), fieldTag)
			}
		} else {
			// 一般字段值
			fieldTag := "`yaml:\"" + fieldName + "\"`"
			line = fmt.Sprintf("%s %s %s \n\t", util.CamelString(fieldName), fieldType, fieldTag)
		}

		fieldList += line
	}

	structCode = fmt.Sprintf(`
type %s struct{
	%s
}
`, structName, fieldList)

	return structCode, nil
}

// 转换slice元素为struct代码
func ConverSliceElementToStructCode(sliceKeyName string, configSlice []interface{}) (structCode string, err error) {
	element := configSlice[0]
	var innerStructCode string
	elementType := reflect.TypeOf(element)
	if elementType.String() == `map[interface {}]interface {}` {
		elementMap := element.(map[interface{}]interface{})
		innerStructCode, err = ConvertMapToStructCode(util.CamelString(sliceKeyName), elementMap)
		if err != nil {
			return "", err
		}
	}
	return innerStructCode, nil
}
