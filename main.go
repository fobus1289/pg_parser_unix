package main

import (
	pgparser "pg_parser/pg_parser"
)

func main() {

	const sqlQuery = `
	/* multiline comment
	* with nesting
	*/
	CREATE TABLE schedule2 (
		id3 JSONB default ((1+2+(1+1))+((3+4)*5))
	);
	`

	pgparser.ParseToJson(sqlQuery)
}

func writeStruct() {

	const tmp = `
	type {{UcCamelCase .Name}} struct {
	{{range .Columns}}
		{{if .Nullable}}
		{{UcCamelCase .Name}} *{{.DataType}} ` + "`json:\"{{CamelCase .Name}},opitempty\"`" + `
		{{ else }}
		{{UcCamelCase .Name}} {{.DataType}} ` + "`json:\"{{CamelCase .Name}}\"`" + `
		{{end}}
	{{end}}
	}
	`

}

type Model struct {
	Id int `json:"id"`
}
