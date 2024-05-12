package storage

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func SaveOrder(dbConn *sqlx.DB, order entities.Order) error {
	_, err := dbConn.Exec(`
        INSERT INTO Orders (order_name, email, tour_id)
        VALUES (?, ?, ?)
    `, order.Name, order.Email, order.TourInfo.Id)
	if err != nil {
		return fmt.Errorf("Failed to insert order into database: %w", err)
	}

	return nil
}
