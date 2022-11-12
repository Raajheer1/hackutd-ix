package models

import (
	"errors"
	"time"
)

type Stock struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Ticker    string    `json:"ticker"`
	Shares    string    `json:"shares"`
}

func (s *Stock) CreateStock() (*Stock, error) {
	if err := DB.Create(&s).Error; err != nil {
		return &Stock{}, err
	}

	return s, nil
}

// TODO - Update Stock Route
func (s *Stock) UpdateStock() (*Stock, error) {
	return &Stock{}, nil
}

// TODO - Delete Stock Route
func (u *Stock) DeleteStock() (*Stock, error) {
	return &Stock{}, nil
}

func GetStockByID(uid uint) (Stock, error) {
	var s Stock
	if err := DB.First(&s, uid).Error; err != nil {
		return s, errors.New("Stock not found!")
	}

	return s, nil
}
