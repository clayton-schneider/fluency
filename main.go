package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/clayton-schneider/fluency/models"
	"github.com/clayton-schneider/fluency/view/components"
	"github.com/clayton-schneider/fluency/view/pages"
)

func main() {
	db := models.CreateDB()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Dashboard(db.Students).Render(r.Context(), w)
	})

	r.Get("/students/{studentId}/charts", func(w http.ResponseWriter, r *http.Request) {

		studentId, err := strconv.Atoi(chi.URLParam(r, "studentId"))

		if err != nil {
			// UPDATE:this means that not a number was passed
			fmt.Println("Must pass an integer")
		}

		charts := models.DBFindStudentCharts(db, studentId)

		components.ChartCardList(charts).Render(r.Context(), w)
	})

	r.Route("/charts", func(r chi.Router) {

		r.Get("/{chartId}", func(w http.ResponseWriter, r *http.Request) {

			chartId, err := strconv.Atoi(chi.URLParam(r, "chartId"))

			if err != nil {
				// UPDATE:this means that not a number was passed
				fmt.Println("Must pass an integer")
			}

			chart := models.DBFindChart(db, chartId)
			measures := models.DBFindChartMeasurements(db, chartId)

			pages.ViewChart(chart, measures).Render(r.Context(), w)

		})

	})

	http.ListenAndServe(":42069", r)
}
