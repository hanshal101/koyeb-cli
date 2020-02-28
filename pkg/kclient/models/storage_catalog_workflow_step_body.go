// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageCatalogWorkflowStepBody storage catalog workflow step body
// swagger:model storageCatalogWorkflowStepBody
type StorageCatalogWorkflowStepBody struct {

	// after
	After string `json:"after,omitempty"`

	// if
	If string `json:"if,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters map[string]StorageServiceParam `json:"parameters,omitempty"`

	// use
	Use string `json:"use,omitempty"`

	// with
	With map[string]interface{} `json:"with,omitempty"`
}

// Validate validates this storage catalog workflow step body
func (m *StorageCatalogWorkflowStepBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCatalogWorkflowStepBody) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for k := range m.Parameters {

		if err := validate.Required("parameters"+"."+k, "body", m.Parameters[k]); err != nil {
			return err
		}
		if val, ok := m.Parameters[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCatalogWorkflowStepBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCatalogWorkflowStepBody) UnmarshalBinary(b []byte) error {
	var res StorageCatalogWorkflowStepBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
