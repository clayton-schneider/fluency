package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/clayton-schneider/fluency/view/components"
	"github.com/clayton-schneider/fluency/view/pages"
)


func main() {
	hPage := components.Page()

	createChart := pages.CreateChart()

	http.Handle("/", templ.Handler(hPage))

	http.Handle("/create-chart", templ.Handler(createChart))

	http.ListenAndServe(":42069", nil)
}
