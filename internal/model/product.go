package model

import "time"

type Product struct {
	ID        uint64
	Name      string
	Details   string
	Price     float64
	ImageUrl  string
	CreatedAt time.Time
}

type Cart struct {
	ID        uint64
	Username  string
	Qty       int
	ProductID uint64
	Product   Product
}
