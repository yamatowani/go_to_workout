package main

import (
	"time"
)

type Workout struct {
	Exercise string
	Weight float64
	Sets int
	Reps []int
	Date time.Time
}
