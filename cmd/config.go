package cmd

import (
	"fmt"
	"github.com/gofuncchan/ginger-cli/util"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

// 根据config 配置文件生成配置参数解析代码
var ConfigCommand = cli.Command{
	Name:        "config",
	Usage:       "generate config parse code for yaml file ",
	UsageText:   "ginger-cli config [sub-command] [option]",
	Description: "generate config init code for yaml config file",
	Subcommands: []cli.Command{
		subCommandParse,
	},
}

var subCommandParse = cli.Command{
	Name:        "parse",
	UsageText:   "",
	Description: "generate config parse code for yaml",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "path, p", Value: "config",Usage:"yaml config file path"},
		cli.StringFlag{Name:"env, e",Value:"debug",Usage:"what env directory"},
		cli.StringFlag{Name:"file, f",Usage:"yaml file name"},
	},
	Action: subCommandParseAction,
}

func subCommandParseAction(c *cli.Context) error {
	configPath := c.String("path")
	envDir := c.String("env")
	fileName := c.String("file")

	target := configPath + "/" + envDir + "/" +fileName + ".yaml"

	baseConfFile, err := ioutil.ReadFile(target)
	if err != nil {
		return err
	}

	configMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal(baseConfFile, configMap)
	if err != nil {
		return err
	}
	// fmt.Println(configMap)

	structName := util.CamelString(fileName)

	structCode,err := convertMapToStructCode(structName,configMap)
	if err != nil {
		return err
	}


	outputStr := fmt.Sprintf(`This code is generated with ginger-cli.

Please copy the following code to /config/parse_config.go:
%s
var %sConf %s

func Init(){
// ...

	%sConfFile, err := ioutil.ReadFile(confPath + "/%s.yaml")
	common.EF(err)
	err = yaml.Unmarshal(%sConfFile, &%sConf)
	common.EF(err)

// ...
}
`,structCode,structName,structName,structName,fileName,structName,structName)

	util.OutputInfo("Generate Successful",outputStr )

	return nil
}

func convertMapToStructCode(structName string,configMap map[interface{}]interface{}) (structCode string,err error) {

	var fieldList string
	for k,v := range configMap {
		fieldName := k.(string)
		fieldType := reflect.TypeOf(v).Name()
		fieldTag := "`yaml:\""+fieldName +"\"`"
		line := fmt.Sprintf(`%s %s %s
	`,util.CamelString(fieldName),fieldType,fieldTag)
		fieldList += line
	}

	structCode = fmt.Sprintf(`
type %s struct{
	%s
}
`,structName,fieldList)


	return structCode,nil
}