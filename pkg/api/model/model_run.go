package model

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"gorm.io/gorm"
)

// ModelRun is the struct to hold the training run done to the ml model
type ModelRun struct {
	CommonFields
	ModelID uint64 `gorm:"not null" json:"model_id"`
	Name    string `json:"name"`
}

// Format formats the input
func (r *ModelRun) Format() {
	if r.Name != "" {
		r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	}
}

// Validate validates the input, use after Format
func (r *ModelRun) Validate(db *gorm.DB) error {
	if r.ModelID == 0 {
		return errors.New("Model ID cannot be empty")
	}
	res := db.Where("id = ?", r.ModelID).Take(&Model{})
	if res.RowsAffected == 0 {
		return fmt.Errorf("Model with id %d doesn't exist", r.ModelID)
	}
	return nil
}

// Create creates a model run for the specified model
func (r *ModelRun) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

// GetByModelID gets all model runs for the specified model from DB
func (r *ModelRun) GetByModelID(db *gorm.DB, modelID uint64) (*[]ModelRun, error) {
	runs := []ModelRun{}
	err := db.Where("model_id = ?", modelID).Find(&runs).Error
	// if project_id is not found returns an error, as Find doesn't return err with o result unlike Take
	if err == nil && len(runs) == 0 {
		err = errors.New("no record found with given project_id")
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
