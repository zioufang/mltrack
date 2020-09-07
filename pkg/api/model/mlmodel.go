package model

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

// Model is the struct to hold a machine learning model
type Model struct {
	CommonFields
	// TODO create index in migration
	Name        string `gorm:"unique;not null;index:model_name" json:"name"`
	ProjectID   uint64 `json:"project_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

// Format formats the input
func (m *Model) Format() {
	if m.Name != "" {
		m.Name = html.EscapeString(strings.TrimSpace(m.Name))
	}
	if m.Description != "" {
		m.Description = strings.TrimSpace(m.Description)
	}
	if m.Status != "" {
		m.Status = strings.TrimSpace(m.Status)
	}
}

// Validate validates the input, use after Format
func (m *Model) Validate(db *gorm.DB) error {
	if m.Name == "" {
		return errors.New("Model Name cannot be empty")
	}
	if m.ProjectID == 0 {
		return errors.New("Project ID cannot be empty")
	}
	res := db.Where("id = ?", m.ProjectID).Take(&Project{})
	if res.RowsAffected == 0 {
		return fmt.Errorf("Project with id %d doesn't exist", m.ProjectID)
	}
	// TODO check if unique is enforced, if so then remove below, and *gorm.DB from func param
	res = db.Where("name = ?", m.Name).Take(&Model{})
	if res.RowsAffected != 0 {
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

// GetByProjectID gets a list of model based on project_id
func (m *Model) GetByProjectID(db *gorm.DB, projectID uint64) (*[]Model, error) {
	models := []Model{}
	err := db.Where("project_id = ?", projectID).Find(&models).Error
	// if project_id is not found returns an error, as Find doesn't return err with o result unlike Take
	if err == nil && len(models) == 0 {
		err = errors.New("no record found with given project_id")
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

// UpdateByID updates one instance by ID with the value from (m *Model)
// project_id cannot be updated
func (m *Model) UpdateByID(db *gorm.DB, id uint64) (*Model, error) {
	updateMap := make(map[string]interface{})
	var err error
	if m.Name != "" {
		updateMap["name"] = m.Name
	}
	if m.Name != "" {
		updateMap["status"] = m.Status
	}
	if m.Description != "" {
		updateMap["description"] = m.Description
	}
	if m.ProjectID != 0 {
		return &Model{}, errors.New("project_id cannot be updated")
	}
	if len(updateMap) == 0 {
		return &Model{}, errors.New("Nothing is provided for update")

	}
	err = db.Model(&Model{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return &Model{}, err
	}
	// get the updated project
	err = db.Model(&Model{}).Where("id = ?", id).Take(&m).Error
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
