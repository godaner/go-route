package go_util

import (
	"html/template"
	"io"
)

func WriteTemplate(w  io.Writer,templ string){
	t, err := template.New("webpage").Parse(templ)
	CheckErr(err)
	err = t.Execute(w, nil)
	CheckErr(err)
}