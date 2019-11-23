# ginger-cli
A client for ginger scaffold

### Install 
`go get -u github.com/gofuncchan/ginger-cli`

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
   ginger_cli init - ginger app init.

USAGE:
   ginger-cli init [--name|-n] [project_name]

DESCRIPTION:
   The init command create a new gin application in current directory，this command will generate some necessary folders and files,which make up project scaffold.

OPTIONS:
   --name value, -n value  project name
   -g                      git init

```

This command will pull `github.com/gofuncchan/ginger` scaffold, it create project default directory.

Tips:

- ginger use go module to manage package,you just do like this: `go mod tidy`,it will download dependency package to local.
- you can also use vendor like this `go mod vendor`, it will copy downloaded package to current vendor directory.
- use `-g` option to git init your project.
  
#### Generate code
ginger-cli provide many command to help you generate code,like:

- handler function code
- sql builder code for mysql
- biz model function code for mysql 
- mongo repository function
- redis cache function

Use `ginger-cli [command] -h` to see command detail.

