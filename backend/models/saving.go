package models

import (
	"errors"
	"time"
)

type Saving struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Amount    string    `json:"amount"`
}

func (s *Saving) CreateSaving() (*Saving, error) {
	if err := DB.Create(&s).Error; err != nil {
		return &Saving{}, err
	}

	return s, nil
}

// TODO - Update Saving Route
func (s *Saving) UpdateSaving() (*Saving, error) {
	return &Saving{}, nil
}

// TODO - Delete Saving Route
func (u *Saving) DeleteSaving() (*Saving, error) {
	return &Saving{}, nil
}

func GetSavingByID(uid uint) (Saving, error) {
	var s Saving
	if err := DB.First(&s, uid).Error; err != nil {
		return s, errors.New("Saving not found!")
	}

	return s, nil
}
