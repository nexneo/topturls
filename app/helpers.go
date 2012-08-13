package app

import (
	"net/http"
)

func (t Action) Render(response http.ResponseWriter, context interface{}) {
	tpl.Lookup("head.html").Execute(response, context)
	tpl.Lookup(string(t)).Execute(response, context)
	tpl.Lookup("tail.html").Execute(response, context)
}
