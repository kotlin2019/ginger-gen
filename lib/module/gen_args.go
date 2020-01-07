package module

import (
	libSchema "github.com/gofuncchan/ginger-gen/lib/schema"
	"strings"
)

func GenCreateArgs(cols libSchema.ColumnSlice) string {
	args := make([]string, 0)
	for _, col := range cols {
		if col.Name == "id" || col.Name == "create_at" || col.Name == "update_at" {
			continue
		}
		args = append(args, "form."+col.GetName())
	}

	return strings.Join(args, ",")
}

func GenCreateModelArgs(cols libSchema.ColumnSlice) string {
	args := make([]string, 0)
	for _, col := range cols {
		if col.Name == "id" || col.Name == "create_at" || col.Name == "update_at" {
			continue
		}

		getType, err := col.GetType()
		if nil != err {
			continue
		}

		arg := col.Name + " " + getType
		args = append(args, arg)

	}

	return strings.Join(args, ",")
}

func GenCreateArgsForMap(cols libSchema.ColumnSlice) []ArgsForMap {

	args := make([]ArgsForMap, 0)
	for _, col := range cols {
		if col.Name == "id" || col.Name == "create_at" || col.Name == "update_at" {
			continue
		}

		col := ArgsForMap{
			Name:    col.Name,
		}
		args = append(args, col)
	}

	return args
}

func GenUpdateArgs(cols libSchema.ColumnSlice) []UpdateArgs {

	args := make([]UpdateArgs, 0)
	for _, col := range cols {
		if col.Name == "id" || col.Name == "create_at" || col.Name == "update_at" {
			continue
		}

		col := UpdateArgs{
			Name:    col.Name,
			GetName: col.GetName(),
		}
		args = append(args, col)
	}

	return args
}
