package model

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// Model is the struct to hold a machine learning model
type Model struct {
	CommonFields
	// TODO create index in migration
	Name        string `gorm:"unique;not null;index:model_name" json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// FormatAndValidate formats the input then validates it
func (m *Model) FormatAndValidate(db *gorm.DB) error {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))
	if m.Name == "" {
		return errors.New("Model Name cannot be empty")
	}
	// TODO check if unique is enforced, if so then remove below, and *gorm.DB from func param
	var count int
	db.Where("name = ?", m.Name).Take(&Model{}).Count(&count)
	if count != 0 {
		return errors.New("Model Name already exists")
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
	err := db.Find(&models).Error
	if err != nil {
		return &[]Model{}, err
	}
	return &models, err

}

// GetByID gets one instance by ID from DB
func (m *Model) GetByID(db *gorm.DB, id uint64) (*Model, error) {
	err := db.Where("id = ?", id).Take(&m).Error
	if err != nil {
		return &Model{}, err
	}
	return m, err
}

// GetByName gets one instance by Name from DB
func (m *Model) GetByName(db *gorm.DB, name string) (*Model, error) {
	err := db.Where("name = ?", name).Take(&m).Error
	if err != nil {
		return &Model{}, err
	}
	return m, err
}

// DeleteByID deletes the current instance from DB
func (m *Model) DeleteByID(db *gorm.DB, id uint64) error {
	return db.Model(&m).Where("id = ?", id).Take(&m).Delete(&m).Error
}

// DeleteByName deletes the current instance from DB
func (m *Model) DeleteByName(db *gorm.DB, name string) error {
	return db.Model(&m).Where("name = ?", name).Take(&m).Delete(&m).Error
}
