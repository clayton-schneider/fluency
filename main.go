package main

import (
	"fmt"
	"net/http"

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
		},
	}

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

	http.HandleFunc("/charts/", func(w http.ResponseWriter, r *http.Request) {
		//f, err := os.ReadFile("./data/test.txt")

		//if err != nil {
		//	fmt.Println("file could not be read")
		//}

		//charts := strings.Split(string(f), "\n")

		//pathChar := strings.Split(r.URL.Path, "/charts/")
		//chartId := pathChar[1]

		pages.ViewChart(chartData).Render(r.Context(), w)

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

	http.ListenAndServe(":42069", nil)
}
