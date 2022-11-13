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

func (v *Saving) CreateSaving() (*Saving, error) {
	if err := DB.Create(&v).Error; err != nil {
		return &Saving{}, err
	}

	return v, nil
}

func GetSaving(userID uint) ([]Saving, error) {
	var savings []Saving
	if err := DB.Where("user_id = ?", userID).Find(&savings).Error; err != nil {
		return savings, err
	}

	return savings, nil
}

func (v *Saving) UpdateSaving() (*Saving, error) {
	var saving Saving
	if err := DB.Where("id = ?", v.ID).First(&saving).Error; err != nil {
		return &Saving{}, err
	}

	if err := DB.Model(&saving).Updates(&v).Error; err != nil {
		return &Saving{}, err
	}

	return v, nil
}

func (v *Saving) DeleteSaving() (*Saving, error) {
	var saving Saving
	if err := DB.Where("id = ?", v.ID).First(&saving).Error; err != nil {
		return &Saving{}, err
	}

	if err := DB.Delete(&saving).Error; err != nil {
		return &Saving{}, err
	}

	return v, nil
}

func GetSavingByID(uid uint) (Saving, error) {
	var v Saving
	if err := DB.First(&v, uid).Error; err != nil {
		return v, errors.New("Saving not found!")
	}

	return v, nil
}
