package model

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// ModelRun is the struct to hold the training run done to the ml model
type ModelRun struct {
	CommonFields
	ModelID uint64 `gorm:"not null" json:"model_id"`
	Name    string `json:"name"`
}

// FormatAndValidate formats the input then validates it
func (r *ModelRun) FormatAndValidate(db *gorm.DB) error {
	if r.Name != "" {
		r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	}
	return nil
}

// Create creates the current instance in DB
func (r *ModelRun) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

// GetAll gets all instances from DB
func (r *ModelRun) GetAll(db *gorm.DB) (*[]ModelRun, error) {
	runs := []ModelRun{}
	err := db.Find(&runs).Error
	if err != nil {
		return &[]ModelRun{}, err
	}
	return &runs, err

}

// GetByID gets one instance by ID from DB
func (r *ModelRun) GetByID(db *gorm.DB, id uint64) (*ModelRun, error) {
	err := db.Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &ModelRun{}, err
	}
	return r, err
}

// DeleteByID deletes the current instance from DB
func (r *ModelRun) DeleteByID(db *gorm.DB, id uint64) error {
	return db.Model(&r).Where("id = ?", id).Take(&r).Delete(&r).Error
}
