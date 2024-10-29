package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var log = "workout_log.json"

func addLog(exercise string, weight float64, sets int, reps []int, date string) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	workout := Workout{
		Exercise: exercise,
		Weight:   weight,
		Sets:     sets,
		Reps:     reps,
		Date:     parsedDate,
	}

	workouts := loadLogs()
	workouts = append(workouts, workout)

	file, err := os.Create(log)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(workouts); err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	fmt.Println("Workout logged successfully")
}

func loadLogs() []Workout {
	var workouts []Workout
	file, err := os.Open(log)
	if err != nil {
		return workouts
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&workouts)
	return workouts
}
