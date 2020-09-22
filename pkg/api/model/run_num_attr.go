package model

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"gorm.io/gorm"
)

// RunNumAttr is Numeric Attribute for Model Runs, e.g. metrics & hyper parameters
type RunNumAttr struct {
	CommonFields
	ModelRunID uint64   `gorm:"not null" json:"model_run_id"`
	Name       string   `gorm:"not null" json:"name"`
	Category   string   `gorm:"not null" json:"category"`
	Value      *float32 `gorm:"not null" json:"value"`
}

// Format formats the input
func (r *RunNumAttr) Format() {
	if r.Name != "" {
		r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	}
}

// Validate validates the input, use after Format
func (r *RunNumAttr) Validate(db *gorm.DB) error {
	if r.ModelRunID == 0 {
		return errors.New("Model Run ID cannot be empty")
	}
	if r.Name == "" {
		return errors.New("Name cannot be empty")
	}
	res := db.Where("id = ?", r.ModelRunID).Take(&ModelRun{})
	if res.RowsAffected == 0 {
		return fmt.Errorf("Model Run with id %d doesn't exist", r.ModelRunID)
	}
	return nil
}

// Create creates a RunNumAttr for the specified model run
func (r *RunNumAttr) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

// Get gets all RunNumAttrs based on input for the specified model run from DB
// modelRunID is required, the rest is optional
func (r *RunNumAttr) Get(db *gorm.DB, modelRunID uint64, name string, category string) (*[]RunNumAttr, error) {
	// build the query
	sqlArgs := make(map[string]interface{})
	sqlArgs["model_run_id"] = modelRunID
	var sqlQuery string = "SELECT * FROM run_num_attrs WHERE model_run_id = @model_run_id"

	if name != "" {
		sqlArgs["name"] = name
		sqlQuery = sqlQuery + " AND name = @name"
	}

	if category != "" {
		sqlArgs["category"] = category
		sqlQuery = sqlQuery + " AND category = @category"
	}

	// run the query
	attrs := []RunNumAttr{}
	err := db.Raw(sqlQuery, sqlArgs).Find(&attrs).Error
	if err == nil && len(attrs) == 0 {
		err = errors.New("no record found")
	}
	return &attrs, err
}
