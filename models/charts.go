package models

type Student struct {
	Id   int
	Name string
}

func newStudent(id int, name string) Student {
	return Student{
		Id:   id,
		Name: name,
	}
}


type Chart struct {
	Id           int
	Date         string
	StudentId    int
	Skill        string
}

func newChart(id int, date string, studentId int, skill string) Chart {
	return Chart{
		Id: id,
		Date: date,
		StudentId: studentId,
		Skill: skill,
	}
}

type Measurement struct {
	Id           int
	ChartId		 int
	Acceleration float32
	Deceleration float32
	Duration     string
}

func newMeasurement(id int, chartId int, a float32, d float32, dur string) Measurement {
	return Measurement{
		Id: id,
		ChartId: chartId,
		Acceleration: a,
		Deceleration: d,
		Duration: dur,
	}
}

type DB struct {
	Students []Student
	Charts   []Chart
	Measurements []Measurement
}

func DBFindChart(db DB, id int) Chart {
	for _, chart := range db.Charts {
		if chart.Id == id {
			return chart
		}
	}

	return db.Charts[0]

}

func DBFindStudentCharts(db DB, studentId int) []Chart {
	var charts []Chart

	for _, chart := range db.Charts {
		if chart.StudentId == studentId {
			charts = append(charts, chart)
		}
	}
	// TODO - Handle the case if no charts are found

	return charts
}

func DBFindChartMeasurements(db DB, chartId int) []Measurement {
	var measures []Measurement

	for _, measure := range db.Measurements {
		if measure.ChartId == chartId {
			measures = append(measures, measure)
		}
	}

	// TODO - Handle the case of no measures
	return measures
}



func CreateDB() DB {

	var measurements []Measurement

	for i := 0; i < 10; i++ {
		measurements = append(measurements, newMeasurement(i, 0, float32(i * 3), float32(i / 2), "time"))
	}

	db := DB{
		Students: []Student{
			newStudent(1, "Clayton"),
			newStudent(2, "John"),
		},
		Charts: []Chart{
			newChart(0, "2023-11-10", 1, "addition"),
			newChart(1, "2023-11-15", 1, "subtraction"),
			newChart(2, "2023-11-10", 2, "addition"),
			newChart(3, "2023-11-15", 2, "subtraction"),
		},
		Measurements: measurements,
	}

	return db
}
