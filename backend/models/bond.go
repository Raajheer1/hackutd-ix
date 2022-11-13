package models

import (
	"errors"
	"time"
)

type Bond struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	UserID       uint      `json:"user_id"`
	CreatedAt    time.Time ``
	UpdatedAt    time.Time ``
	PurchaseDate string    `json:"purchase_date"`
	Maturity     string    `json:"maturity"`
}

func (b *Bond) CreateBond() (*Bond, error) {
	if err := DB.Create(&b).Error; err != nil {
		return &Bond{}, err
	}

	return b, nil
}

func GetBond(userID uint) ([]Bond, error) {
	var bonds []Bond
	if err := DB.Where("user_id = ?", userID).Find(&bonds).Error; err != nil {
		return bonds, err
	}

	return bonds, nil
}

func (b *Bond) UpdateBond() (*Bond, error) {
	var bond Bond
	if err := DB.Where("id = ? AND user_id = ?", b.ID, b.UserID).First(&bond).Error; err != nil {
		return &Bond{}, err
	}

	if err := DB.Model(&bond).Updates(&b).Error; err != nil {
		return &Bond{}, err
	}

	return b, nil
}

func (b *Bond) DeleteBond() (*Bond, error) {
	var bond Bond
	if err := DB.Where("id = ? AND user_id = ?", b.ID, b.UserID).First(&bond).Error; err != nil {
		return &Bond{}, err
	}

	if err := DB.Delete(&bond).Error; err != nil {
		return &Bond{}, err
	}

	return b, nil
}

func GetBondByID(uid uint) (Bond, error) {
	var b Bond
	if err := DB.First(&b, uid).Error; err != nil {
		return b, errors.New("Bond not found!")
	}

	return b, nil
}
