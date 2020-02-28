// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StorageStackBodyStatus storage stack body status
// swagger:model storageStackBodyStatus
type StorageStackBodyStatus string

const (

	// StorageStackBodyStatusDISABLED captures enum value "DISABLED"
	StorageStackBodyStatusDISABLED StorageStackBodyStatus = "DISABLED"

	// StorageStackBodyStatusENABLED captures enum value "ENABLED"
	StorageStackBodyStatusENABLED StorageStackBodyStatus = "ENABLED"

	// StorageStackBodyStatusLOCKED captures enum value "LOCKED"
	StorageStackBodyStatusLOCKED StorageStackBodyStatus = "LOCKED"

	// StorageStackBodyStatusERROR captures enum value "ERROR"
	StorageStackBodyStatusERROR StorageStackBodyStatus = "ERROR"
)

// for schema
var storageStackBodyStatusEnum []interface{}

func init() {
	var res []StorageStackBodyStatus
	if err := json.Unmarshal([]byte(`["DISABLED","ENABLED","LOCKED","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageStackBodyStatusEnum = append(storageStackBodyStatusEnum, v)
	}
}

func (m StorageStackBodyStatus) validateStorageStackBodyStatusEnum(path, location string, value StorageStackBodyStatus) error {
	if err := validate.Enum(path, location, value, storageStackBodyStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this storage stack body status
func (m StorageStackBodyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStorageStackBodyStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
