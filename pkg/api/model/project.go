package model

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

// Project is the struct to hold a project
type Project struct {
	CommonFields
	// TODO create index in migration
	Name        string `gorm:"unique;not null;index:project_name" json:"name"`
	Description string `json:"description"`
}

// Format formats the input
func (m *Project) Format() {
	if m.Name != "" {
		m.Name = html.EscapeString(strings.TrimSpace(m.Name))
	}
	if m.Description != "" {
		m.Description = strings.TrimSpace(m.Description)
	}
}

// Validate validates the input, use after Format
func (m *Project) Validate(db *gorm.DB) error {
	if m.Name == "" {
		return errors.New("Project Name cannot be empty")
	}
	// TODO check if unique is enforced, if so then remove below, and *gorm.DB from func param
	res := db.Where("name = ?", m.Name).Take(&Project{})
	if res.RowsAffected != 0 {
		return errors.New("Project Name already exists")
	}
	return nil
}

// Create creates the current instance in DB
func (m *Project) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

// GetAll gets all instances from DB
func (m *Project) GetAll(db *gorm.DB) (*[]Project, error) {
	projects := []Project{}
	err := db.Find(&projects).Error
	if err != nil {
		return &[]Project{}, err
	}
	return &projects, err

}

// GetByID gets one instance by ID from DB
func (m *Project) GetByID(db *gorm.DB, id uint64) (*Project, error) {
	err := db.Where("id = ?", id).Take(&m).Error
	if err != nil {
		return &Project{}, err
	}
	return m, err
}

// GetByName gets one instance by Name from DB
func (m *Project) GetByName(db *gorm.DB, name string) (*Project, error) {
	err := db.Where("name = ?", name).Take(&m).Error
	if err != nil {
		return &Project{}, err
	}
	return m, err
}

// UpdateByID updates one instance by ID with the value from (m *Project)
func (m *Project) UpdateByID(db *gorm.DB, id uint64) (*Project, error) {
	updateMap := make(map[string]interface{})
	var err error
	if m.Name != "" {
		updateMap["name"] = m.Name
	}
	if m.Description != "" {
		updateMap["description"] = m.Description
	}
	if len(updateMap) == 0 {
		return &Project{}, errors.New("Updated name or description needs to be provided")

	}
	err = db.Model(&Project{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return &Project{}, err
	}
	// get the updated project
	err = db.Model(&Project{}).Where("id = ?", id).Take(&m).Error
	if err != nil {
		return &Project{}, err
	}
	return m, err
}

// DeleteByID deletes the current instance from DB
func (m *Project) DeleteByID(db *gorm.DB, id uint64) error {
	return db.Model(&m).Where("id = ?", id).Take(&m).Delete(&m).Error
}

// DeleteByName deletes the current instance from DB
func (m *Project) DeleteByName(db *gorm.DB, name string) error {
	return db.Model(&m).Where("name = ?", name).Take(&m).Delete(&m).Error
}
