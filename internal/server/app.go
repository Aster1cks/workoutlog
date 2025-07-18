package server

import (
	"log"

	"github.com/Aster1cks/workoutlog/internal/database"
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Workouts    *database.WorkoutModel
}
