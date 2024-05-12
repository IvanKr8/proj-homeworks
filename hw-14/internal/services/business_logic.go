package services

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/entities"
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/storage"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func GetTourStatus(currentTime time.Time, startDate time.Time, endDate time.Time) string {
	if currentTime.Before(startDate) {
		return "Скоро"
	} else if currentTime.After(endDate) {
		return "Завершений"
	} else {
		return "Триває"
	}
}

func BuyTour(dbConn *sqlx.DB, tourID int, order entities.Order) error {
	tour, err := storage.GetTour(dbConn, tourID)
	if err != nil {
		return err
	}

	order.TourInfo = tour
	err = storage.SaveOrder(dbConn, order)
	if err != nil {
		return err
	}
	SendTourPurchaseEmail(order)
	return nil
}

func SendTourPurchaseEmail(order entities.Order) {
	log.Printf(`
Шановний(а) %s,

Дякуємо за вашу покупку тура в нашому туристичному агентстві!

Ми отримали ваше замовлення і раді повідомити вам, що процес оплати успішно завершено.

Деталі вашого замовлення:
- Тур: %s
- Дата виїзду: %s
- Загальна вартість: %f
`, order.Name, order.TourInfo.Title, order.TourInfo.StartDate, order.TourInfo.Price)
}
