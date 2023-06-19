package orderdetailmodel

import (
	"TKPM-Go/common"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Product_Origin struct {
	UID 		*common.UID `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (j *Product_Origin) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Product_Origin
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*j = img
	return nil
}

func (j *Product_Origin) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Product_Origins []Product_Origin

func (j *Product_Origins) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Product_Origin
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*j = img
	return nil
}

func (j *Product_Origins) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
