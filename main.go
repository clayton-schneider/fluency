package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/clayton-schneider/fluency/view/components"
	"github.com/clayton-schneider/fluency/view/pages"
)


func main() {
	hPage := components.Page()


	http.Handle("/", templ.Handler(hPage))

	http.HandleFunc("/create-chart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			pages.CreateChart().Render(r.Context(), w)
		}

		if r.Method == "POST" {
			components.Page().Render(r.Context(), w)			
		}
	})

		


	http.ListenAndServe(":42069", nil)
}
