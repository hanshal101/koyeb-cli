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

// StorageDeliveryBodyStatus storage delivery body status
// swagger:model storageDeliveryBodyStatus
type StorageDeliveryBodyStatus string

const (

	// StorageDeliveryBodyStatusDISABLED captures enum value "DISABLED"
	StorageDeliveryBodyStatusDISABLED StorageDeliveryBodyStatus = "DISABLED"

	// StorageDeliveryBodyStatusENABLED captures enum value "ENABLED"
	StorageDeliveryBodyStatusENABLED StorageDeliveryBodyStatus = "ENABLED"

	// StorageDeliveryBodyStatusLOCKED captures enum value "LOCKED"
	StorageDeliveryBodyStatusLOCKED StorageDeliveryBodyStatus = "LOCKED"

	// StorageDeliveryBodyStatusERROR captures enum value "ERROR"
	StorageDeliveryBodyStatusERROR StorageDeliveryBodyStatus = "ERROR"
)

// for schema
var storageDeliveryBodyStatusEnum []interface{}

func init() {
	var res []StorageDeliveryBodyStatus
	if err := json.Unmarshal([]byte(`["DISABLED","ENABLED","LOCKED","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageDeliveryBodyStatusEnum = append(storageDeliveryBodyStatusEnum, v)
	}
}

func (m StorageDeliveryBodyStatus) validateStorageDeliveryBodyStatusEnum(path, location string, value StorageDeliveryBodyStatus) error {
	if err := validate.Enum(path, location, value, storageDeliveryBodyStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this storage delivery body status
func (m StorageDeliveryBodyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStorageDeliveryBodyStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
