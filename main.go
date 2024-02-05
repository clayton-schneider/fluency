package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/clayton-schneider/fluency/models"
	"github.com/clayton-schneider/fluency/view/components"
	"github.com/clayton-schneider/fluency/view/pages"
)

func main() {
	db := models.CreateDB()

	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("public"))

	r.Handle("/public/*", http.StripPrefix("/public/", fs))

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

	r.Route("/measurements", func (r chi.Router) {

		r.Get("/{measurementId}/edit", func(w http.ResponseWriter, r *http.Request) {

			measurementId, err := strconv.Atoi(chi.URLParam(r, "measurementId"))

			if err != nil {
				// UPDATE WITH BETTER ERROR HANDLING
				fmt.Println("Must pass int")
			}

			measurement := models.DBFindChartMeasurement(db, measurementId)
			pages.EditMeasurement(measurement).Render(r.Context(), w)
		})

		r.Get("/{measurementId}", func(w http.ResponseWriter, r *http.Request) {
			measurementId, err := strconv.Atoi(chi.URLParam(r, "measurementId"))

			if err != nil {
				// UPDATE WITH BETTER ERROR HANDLING
				fmt.Println("Must pass int")
			}

			measurement := models.DBFindChartMeasurement(db, measurementId)
			pages.MeasurementInTable(measurement).Render(r.Context(), w)
		})	
	
		r.Put("/{measurementId}", func(w http.ResponseWriter, r *http.Request) {
			measurementId, err := strconv.Atoi(chi.URLParam(r, "measurementId"))

			if err != nil {
				// UPDATE WITH BETTER ERROR HANDLING
				fmt.Println("Must pass int")
			}

			measurement := models.DBFindChartMeasurement(db, measurementId)
			r.ParseForm()

			ac, _ := strconv.ParseFloat(r.FormValue("acceleration"), 32)
			measurement.Acceleration = ac
			dc, _ := strconv.ParseFloat(r.FormValue("deceleration"), 32)
			measurement.Deceleration = dc
			measurement.Duration = r.FormValue("duration")
			
			models.DBUpdateMeasurement(db, measurement)

			pages.MeasurementInTable(measurement).Render(r.Context(), w)
		})	

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
	}
	http.ListenAndServe(":" + port, r)
}
