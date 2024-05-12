package entities

type Order struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"customer_name" db:"customer_name"`
	Email    string `json:"email" db:"email"`
	TourInfo Tour   `json:"tour_info" db:"-"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}
