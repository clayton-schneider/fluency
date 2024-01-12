package main

import (
	"fmt"
	"net/http"
	//"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/clayton-schneider/fluency/models"
	"github.com/clayton-schneider/fluency/view/pages"
)




func main() {
	db := models.CreateDB()



	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Dashboard(db.Students).Render(r.Context(), w)
	})

	// r.Route("/charts", func(r chi.Router) {
	//
	// 	// LIST CHARTS
	// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("List Charts"))
	// 	})
	// 	
	// 	// GET - SHOW CHART PAGE
	// 	r.Get("/{chartId}", func(w http.ResponseWriter, r *http.Request) {
	// 		chartId, err := strconv.Atoi(chi.URLParam(r, "chartId"))
	//
	// 		if err != nil {
	// 			// UPDATE:this means that not a number was passed
	// 			fmt.Println("Must pass an integer")
	// 		}
	// 	})
	// })

	http.ListenAndServe(":42069", r)
}
