package database

import (
	"database/sql"
	"time"
)

type Workoutentry struct {
	ID          int       `json:"id" db:"id"`
	WorkoutType string    `json:"workout_type" db:"workout_type"`
	Duration    int       `json:"duration_minutes" db:"duration_minutes"`
	Notes       string    `json:"notes" db:"notes"`
	Date        time.Time `json:"date" db:"date"`
}

type WorkoutModel struct {
	DB *sql.DB
}
