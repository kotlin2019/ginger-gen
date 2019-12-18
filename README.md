# ginger-gen
A Code Generator for ginger scaffold

### Install 
`go get -u github.com/gofuncchan/ginger-gen`

### Commands:

```
   init     ginger app init.
   mysql    generate dao code
   handler  generate handler function code
   model    generate biz logic model function code for mysql builder
   repo     generate repo file and function code for mongodb repository
   cache    generate cache function code
   config   generate config parse code for yaml file 
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Usage

#### init project
```
NAME:
   ginger-gen init - ginger app init.

USAGE:
   ginger-gen init [--name|-n] [project_name]

DESCRIPTION:
   The init command create a new gin application in current directory，this command will generate some necessary folders and files,which make up project scaffold.

OPTIONS:
   --name value, -n value  project name
   -g                      git init

```

init 命令会拉取 `github.com/gofuncchan/ginger` 项目脚手架, 默认当前目录创建项目。

Tips:

- 由于ginger使用go module 管理依赖包，默认配置你的config就可以启动，`go build main.go`;
- 默认根包为`github.com/gofuncchan/ginger`,如需更改请自行全局replace替换并修改go.mod文件；
- 一旦替换根包，则该工具的多数命令都带-r 参数设置你的自定义根包，以便生成的代码与你的项目一致；
- 如ide不能识别包，请使用`go mod tidy`和`go mod vendor`做本地化处理，
- init 时可使用`-g` 选项自动git初始化
  
#### Generate code
ginger-gen 提供多个命令方便你自动生成脚手架代码，如下：

- handler 控制器相关代码
- 基于didi/gendry 的sql builder，生成相关表的基本curd代码
- 基于mysql的业务模型层代码
- 基于mgo的数据存储层代码
- 基于redis缓存层的代码
- 根据yaml文件自动生成解析逻辑代码

Use `ginger-gen [command] -h` to see command detail.

