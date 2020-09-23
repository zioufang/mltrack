package model

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"gorm.io/gorm"
)

// RunTag is tag for Model Runs
type RunTag struct {
	CommonFields
	ModelRunID uint64 `gorm:"not null" json:"model_run_id"`
	Key        string `gorm:"not null" json:"key"`
	Value      string `gorm:"not null" json:"value"`
}

// Format formats the input
func (r *RunTag) Format() {
	if r.Key != "" {
		r.Key = html.EscapeString(strings.TrimSpace(r.Key))
	}
}

// Validate validates the input, use after Format
func (r *RunTag) Validate(db *gorm.DB) error {
	if r.ModelRunID == 0 {
		return errors.New("Model Run ID cannot be empty")
	}
	if r.Key == "" {
		return errors.New("Key cannot be empty")
	}
	res := db.Where("id = ?", r.ModelRunID).Take(&ModelRun{})
	if res.RowsAffected == 0 {
		return fmt.Errorf("Model Run with id %d doesn't exist", r.ModelRunID)
	}
	res = db.Where("id = ? and key = ?", r.ModelRunID, r.Key).Take(&RunTag{})
	if res.RowsAffected > 0 {
		return fmt.Errorf("Key %s already exists for this model run", r.Key)
	}
	return nil
}

// Create creates a RunTag for the specified model run
func (r *RunTag) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

// Get gets all RunTags based on input for the specified model run from DB
// modelRunID is required, the rest is optional
func (r *RunTag) Get(db *gorm.DB, modelRunIDs []uint64, keys []string) (*[]RunTag, error) {
	// build the query
	sqlArgs := make(map[string]interface{})
	sqlArgs["model_run_ids"] = modelRunIDs
	var sqlQuery string = "SELECT * FROM run_tags WHERE model_run_id IN @model_run_ids"

	if len(keys) > 0 {
		sqlArgs["keys"] = keys
		sqlQuery = sqlQuery + " AND key IN @keys"
	}
	// run the query
	tags := []RunTag{}
	err := db.Raw(sqlQuery, sqlArgs).Find(&tags).Error
	if err == nil && len(tags) == 0 {
		err = errors.New("no record found")
	}
	return &tags, err
}
