package utils

import (
	"html/template"
)

func HtmlX(x string) template.HTML {
	return template.HTML(x)
}
