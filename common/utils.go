package common

import (
	"regexp"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var templateFuncMap = template.FuncMap{
	"CamelCase":   CamelCase,
	"UcCamelCase": UcCamelCase,
	"SnakeCase":   SnakeCase,
	"UpperCase":   UpperCase,
	"UcFirst":     UcFirst,
	"LcFirst":     LcFirst,
}

func CamelCase(s string) string {

	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	s = strings.ReplaceAll(s, "_", " ")

	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

	s = strings.ReplaceAll(s, " ", "")

	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

func UcCamelCase(s string) string {
	return UcFirst(CamelCase(s))
}

func SnakeCase(input string) string {
	snakeCase := strings.ReplaceAll(input, " ", "_")
	return strings.ToLower(snakeCase)
}

func UpperCase(input string) string {
	return strings.ToUpper(input)
}

func UcFirst(s string) string {
	return cases.Title(language.Tag{}, cases.NoLower).String(s)
}

func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func NewTemplate() *template.Template {
	temp := template.New("").Funcs(templateFuncMap)
	return temp
}
