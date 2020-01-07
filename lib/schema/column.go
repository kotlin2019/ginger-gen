package schema

import "github.com/gofuncchan/ginger-gen/util"

var (
	TypeWrappers = []typeWrapper{i64TypeWrapper, byteTypeWrapper, intTypeWrapper, float64TypeWrapper, stringTypeWrapper, timeTypeWrapper}
)

// Column stands for a column of a table
type Column struct {
	Name    string `json:"COLUMN_NAME"`
	Type    string `json:"COLUMN_TYPE"`
	Comment string `json:"COLUMN_COMMENT"`
}

// GetType returns which built in type the column should be in generated go code
func (c *Column) GetType() (string, error) {
	t := getType(c.Type)
	if "" == t {
		return "", errUnknownType(c.Name, c.Type)
	}
	return t, nil
}

// GetName returns the Cammel Name of the struct
func (c *Column) GetName() string {
	return util.CamelString(c.Name)
}

func getType(t string) string {
	for _, wrapper := range TypeWrappers {
		typer := wrapper(t)
		if typer.Match() {
			return typer.Type()
		}
	}
	return ""
}
