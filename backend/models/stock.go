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

func GetStock(userID uint) ([]Stock, error) {
	var stocks []Stock
	if err := DB.Where("user_id = ?", userID).Find(&stocks).Error; err != nil {
		return stocks, err
	}

	return stocks, nil
}

func (s *Stock) UpdateStock() (*Stock, error) {
	var stock Stock
	if err := DB.Where("id = ? AND user_id = ?", s.ID, s.UserID).First(&stock).Error; err != nil {
		return &Stock{}, err
	}

	if err := DB.Model(&stock).Updates(&s).Error; err != nil {
		return &Stock{}, err
	}

	return s, nil
}

func (s *Stock) DeleteStock() (*Stock, error) {
	var stock Stock
	if err := DB.Where("id = ? AND user_id = ?", s.ID, s.UserID).First(&stock).Error; err != nil {
		return &Stock{}, err
	}

	if err := DB.Delete(&stock).Error; err != nil {
		return &Stock{}, err
	}

	return s, nil
}

func GetStockByID(uid uint) (Stock, error) {
	var s Stock
	if err := DB.First(&s, uid).Error; err != nil {
		return s, errors.New("Stock not found!")
	}

	return s, nil
}
