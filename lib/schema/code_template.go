package schema

import (
	"bytes"
	"fmt"
	"io"
)

func AddImportContent(packageName string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package schema

import(
	"time"
)
	`)))
}

const codeTemplate = `

const {{ .StructName }}TableName = "{{ .TableName }}"

// {{ .StructName }} is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}} struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }}
{{- end}}
}

func(*{{ .StructName }})TableName() string {
	return {{ .StructName }}TableName
}

`
