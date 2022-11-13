package models

import (
	"errors"
	"time"
)

type Bond struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Symbol    string    `json:"symbol"`
	Quantity    string    `json:"quantity"`
}

func (s *Bond) CreateBond() (*Bond, error) {
	if err := DB.Create(&s).Error; err != nil {
		return &Bond{}, err
	}

	return s, nil
}

// TODO - Update Bond Route
func (s *Bond) UpdateBond() (*Bond, error) {
	return &Bond{}, nil
}

// TODO - Delete Bond Route
func (u *Bond) DeleteBond() (*Bond, error) {
	return &Bond{}, nil
}

func GetBondByID(uid uint) (Bond, error) {
	var s Bond
	if err := DB.First(&s, uid).Error; err != nil {
		return s, errors.New("Bond not found!")
	}

	return s, nil
}
