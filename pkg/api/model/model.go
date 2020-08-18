package model

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// Model is the struct to hold machine learning model
type Model struct {
	gorm.Model
	Name string
}

// FormatAndValidate formats the input then validates it
func (m *Model) FormatAndValidate() error {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))
	if m.Name == "" {
		return errors.New("Model Name cannot be empty")
	}
	return nil
}

// CreateModel creates the current instance in DB
func (m *Model) CreateModel(db *gorm.DB) error {
	return db.Create(&m).Error
}

// GetAllModels gets all instances from DB
func (m *Model) GetAllModels(db *gorm.DB) (*[]Model, error) {
	models := []Model{}
	err := db.Find(&models).Error
	if err != nil {
		return &[]Model{}, err
	}
	return &models, err

}

// GetModelByID gets one instance by ID from DB
func (m *Model) GetModelByID(db *gorm.DB, uid uint64) (*Model, error) {
	err := db.Where("id = ?", uid).Take(&m).Error
	if err != nil {
		return &Model{}, err
	}
	return m, err
}

// DeleteModelByID deletes the current instance from DB
func (m *Model) DeleteModelByID(db *gorm.DB, uid uint64) error {
	return db.Model(&m).Where("id = ?", uid).Take(&m).Delete(&m).Error
}
