// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageUpdateConnectorReply storage update connector reply
//
// swagger:model storageUpdateConnectorReply
type StorageUpdateConnectorReply struct {

	// connector
	Connector *StorageConnector `json:"connector,omitempty"`
}

// Validate validates this storage update connector reply
func (m *StorageUpdateConnectorReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUpdateConnectorReply) validateConnector(formats strfmt.Registry) error {

	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if m.Connector != nil {
		if err := m.Connector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageUpdateConnectorReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUpdateConnectorReply) UnmarshalBinary(b []byte) error {
	var res StorageUpdateConnectorReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
