package storage

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetToursList(dbConn *sqlx.DB) ([]entities.Tour, error) {
	rows, err := dbConn.Queryx("SELECT * FROM Tours")
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve tours from database: %w", err)
	}
	defer rows.Close()

	var tours []entities.Tour
	for rows.Next() {
		var tour entities.Tour

		if err = rows.StructScan(&tour); err != nil {
			return nil, fmt.Errorf("Failed to scan tour row: %w", err)
		}
		tours = append(tours, tour)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over tour rows: %w", err)
	}

	return tours, nil
}

func GetTour(dbConn *sqlx.DB, tourID int) (entities.Tour, error) {
	var tour entities.Tour

	err := dbConn.Get(&tour, "SELECT * FROM Tours WHERE id = ?", tourID)
	if err != nil {
		return entities.Tour{}, fmt.Errorf("Failed to retrieve tour with ID %d from database: %w", tourID, err)
	}

	return tour, nil
}

func SaveNewTour(dbConn *sqlx.DB, tour entities.Tour) error {
	_, err := dbConn.Exec(`
	INSERT INTO Tours (title, description, price, transport, start_date, end_date, status)
	VALUES (?, ?, ?, ?, ?, ?, ?)
`, tour.Title, tour.Description, tour.Price, tour.Transport, tour.StartDate, tour.EndDate, tour.Status)
	if err != nil {
		return fmt.Errorf("Failed to insert new tour into database: %w", err)
	}

	return nil
}

func RemoveTour(dbConn *sqlx.DB, tourID int) error {
	_, err := dbConn.Exec("DELETE FROM Tours WHERE id = ?", tourID)
	if err != nil {
		return fmt.Errorf("Failed to remove tour with ID %d from database: %w", tourID, err)
	}
	return nil
}
