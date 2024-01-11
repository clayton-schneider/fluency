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
	chartOne := models.Chart{
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
	chartTwo := models.Chart{
		Date:    "2024-01-02",
		Student: "carly",
		Skill:   "subtraction",
		Measurements:    []models.Measurements {
			{ 
			Id: 1,
			Acceleration: 10.0, 
			Deceleration: 4.0, 
			Duration: "10m20s",
			}, 
			{ 
			Id: 2,
			Acceleration: 20.0, 
			Deceleration: 2.0, 
			Duration: "2m00s",
			}, 
		},
	}

	charts := []models.Chart{chartOne, chartTwo}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Route("/charts", func(r chi.Router) {

		// LIST CHARTS
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("List Charts"))
		})
		
		// GET - SHOW CHART PAGE
		r.Get("/{chartId}", func(w http.ResponseWriter, r *http.Request) {
			chartId, err := strconv.Atoi(chi.URLParam(r, "chartId"))

			if err != nil {
				// UPDATE:this means that not a number was passed
				fmt.Println("Must pass an integer")
			}

			pages.ViewChart(charts[chartId-1]).Render(r.Context(), w)
		})
	})

	http.ListenAndServe(":42069", r)
}
