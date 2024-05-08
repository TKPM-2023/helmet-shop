package orderdetailmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orgball2608/helmet-shop-be/common"
)

type ProductOrigin struct {
	UID         *common.UID    `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      *common.Images `json:"images"`
}

func (j *ProductOrigin) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img ProductOrigin
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*j = img
	return nil
}

func (j *ProductOrigin) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type ProductOrigins []ProductOrigin

func (j *ProductOrigins) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []ProductOrigin
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*j = img
	return nil
}

func (j *ProductOrigins) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
