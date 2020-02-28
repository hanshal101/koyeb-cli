// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateUserRequestUserUpdateBody update user request user update body
// swagger:model UpdateUserRequestUserUpdateBody
type UpdateUserRequestUserUpdateBody struct {

	// current password
	CurrentPassword string `json:"current_password,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// firstname
	Firstname string `json:"firstname,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// lastname
	Lastname string `json:"lastname,omitempty"`

	// newsletter subscribed
	NewsletterSubscribed bool `json:"newsletter_subscribed,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this update user request user update body
func (m *UpdateUserRequestUserUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserRequestUserUpdateBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserRequestUserUpdateBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserRequestUserUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
