package app

import (
	"go/build"
	"html/template"
	"log"
	"net/http"
)

const basePkg = "github.com/nexneo/topturls"


var (
	tpl      *template.Template
	rootPath string
)

func init() {
	appSetup()
	routerSetup()
}

func appSetup() {
	p, err := build.Default.Import(basePkg, "", build.FindOnly)
	if err != nil {
		log.Fatalf("Couldn't find staic: %v", err)
	}
	rootPath = p.Dir
	templatePath := rootPath + "/templates/*.html"
	tpl = template.Must(template.ParseGlob(templatePath))
}

func Start() {
	//start http server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
