package entities

type Tour struct {
	Id          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	Transport   string  `db:"transport" json:"transport"`
	StartDate   string  `db:"start_date" json:"start_date"`
	EndDate     string  `db:"end_date" json:"end_date"`
	Status      string  `db:"status" json:"status"`
}

type Tours struct {
	Tours []Tour `db:"tours" json:"tours"`
}
