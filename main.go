package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/clayton-schneider/fluency/models"
	"github.com/clayton-schneider/fluency/view/pages"
)




func main() {
	chartData := models.Chart{
		Date:    "2024-01-01",
		Student: "clayton",
		Skill:   "addition",
		Measurements:    []models.Measurements {
			{ 
			Id: 1,
			Acceleration: 32.0, 
			Deceleration: 1.0, 
			Duration: "1m30s",
			}, 
			{ 
			Id: 2,
			Acceleration: 30.0, 
			Deceleration: 2.0, 
			Duration: "2m00s",
			}, 
		},
	}

	charts := []models.Chart{chartData}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Route("/chart", func(r chi.Router) {
		r.Get("/{chartId}", func(w http.ResponseWriter, r *http.Request) {
			chartId, err := strconv.Atoi(chi.URLParam(r, "chartId"))

			if err != nil {
				// UPDATE:this means that not a number was passed
				fmt.Println("Must pass an integer")
			}

			pages.ViewChart(charts[chartId-1]).Render(r.Context(), w)
		})
	})


	http.HandleFunc("/create-chart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			pages.CreateChart().Render(r.Context(), w)
		}

		if r.Method == "POST" {
			r.ParseForm()

			for key, val := range r.Form {
				fmt.Printf("%s, %s\n", key, val)
			}
			pages.ViewChart(chartData).Render(r.Context(), w)
		}
	})

	http.HandleFunc("/measurements/1/edit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			pages.EditMeasurement(chartData.Measurements[0]).Render(r.Context(), w)
		}
	})

	http.HandleFunc("/measurements/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			pages.MeasurementInTable(chartData.Measurements[0]).Render(r.Context(), w)
		}
	})

	http.ListenAndServe(":42069", r)
}
