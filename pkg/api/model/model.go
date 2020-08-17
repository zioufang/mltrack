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

// Create creates the current instance in DB
func (m *Model) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// GetAll gets all instances from DB
func (m *Model) GetAll(db *gorm.DB) (*[]Model, error) {
	models := []Model{}
	err := db.Find([]Model{}).Error
	if err != nil {
		return &[]Model{}, err
	}
	return &models, err

}

// GetByID gets one instance by ID from DB
func (m *Model) GetByID(db *gorm.DB, uid uint) (*Model, error) {
	err := db.Where("id = ?", uid).Take(&m).Error
	if err != nil {
		return &Model{}, err
	}
	return m, err
}

// Delete deletes the current instance from DB
func (m *Model) Delete(db *gorm.DB) error {
	return db.Delete(&m).Error
}
