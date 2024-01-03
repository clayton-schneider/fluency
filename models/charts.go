package models

type Chart struct {
	Date string
	Student string
	Skill string
	Measurements []Measurements
}

type Measurements struct {
		Id int
		Acceleration float32
		Deceleration float32
		Duration string 
}
