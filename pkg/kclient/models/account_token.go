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

// AccountToken account token
// swagger:model accountToken
type AccountToken struct {

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this account token
func (m *AccountToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountToken) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(m.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires", "body", "date-time", m.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountToken) UnmarshalBinary(b []byte) error {
	var res AccountToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
