package database

import (
	"database/sql"

	"github.com/Aster1cks/workoutlog/internal/errdef"
)

func (w *WorkoutModel) GetAll() ([]*Workoutentry, error) {
	stmt := "SELECT id, workout_type, duration_minutes, notes, date FROM workouts"
	rows, err := w.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workoutsTable := []*Workoutentry{}

	for rows.Next() {
		tw := &Workoutentry{}
		err = rows.Scan(&tw.ID, &tw.WorkoutType, &tw.Duration, &tw.Notes, &tw.Date)
		if err == sql.ErrNoRows {
			return nil, errdef.ErrNoRecord
		} else if err != nil {
			return nil, err
		}
		workoutsTable = append(workoutsTable, tw)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workoutsTable, nil
}

func (w *WorkoutModel) DeleteEntry(id int) (*Workoutentry, error) {
	stmt := "DELETE FROM workouts WHERE id = $1 RETURNING *"
	tw := &Workoutentry{}
	err := w.DB.QueryRow(stmt, id).Scan(&tw.ID, &tw.WorkoutType, &tw.Duration, &tw.Notes, &tw.Date)
	if err == sql.ErrNoRows {
		return nil, errdef.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return tw, nil
}

func (w *WorkoutModel) AddEntry(wType string, duration int, notes string) (int, error) {
	stmt := `
	INSERT INTO workouts (workout_type, duration_minutes, notes, date)
	VALUES ($1, $2, $3, NOW()) RETURNING id
	`
	var id int
	err := w.DB.QueryRow(stmt, wType, duration, notes).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (w *WorkoutModel) EditEntry(wType string, duration int, notes string, id int) error {
	stmt := `UPDATE workouts
    SET workout_type = $1, duration_minutes = $2, notes = $3
    WHERE id = $4;`

	_, err := w.DB.Exec(stmt, wType, duration, notes, id)
	//err := w.DB.QueryRow(stmt, wType, duration, notes).Scan(&id)
	if err == sql.ErrNoRows {
		return errdef.ErrNoRecord
	} else if err != nil {
		return err
	}
	return nil
}

func (w *WorkoutModel) EntryByID(id int) (*Workoutentry, error) {
	stmt := `SELECT id, workout_type, duration_minutes, notes, date FROM workouts WHERE id = $1`
	tw := &Workoutentry{}

	err := w.DB.QueryRow(stmt, id).Scan(&tw.ID, &tw.WorkoutType, &tw.Duration, &tw.Notes, &tw.Date)
	if err == sql.ErrNoRows {
		return nil, errdef.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return tw, nil

}
