package httputil

import (
	"go-util/errorutil"
	"html/template"
	"io"
)

func WriteTemplate(w  io.Writer,templ string){
	t, err := template.New("webpage").Parse(templ)
	errorutil.CheckErr(err)
	err = t.Execute(w, nil)
	errorutil.CheckErr(err)
}