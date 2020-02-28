// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageCatalogIntegrationBody storage catalog integration body
// swagger:model storageCatalogIntegrationBody
type StorageCatalogIntegrationBody struct {

	// author
	Author string `json:"author,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters map[string]StorageServiceParam `json:"parameters,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// required services
	RequiredServices []*StorageCatalogServiceLightBody `json:"required_services"`

	// service
	Service *StorageCatalogServiceLightBody `json:"service,omitempty"`

	// short description
	ShortDescription string `json:"short_description,omitempty"`

	// status
	Status StorageCatalogItemStatus `json:"status,omitempty"`

	// steps
	Steps []*StorageCatalogWorkflowStepBody `json:"steps"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type StorageIntegrationType `json:"type,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this storage catalog integration body
func (m *StorageCatalogIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCatalogIntegrationBody) validateParameters(formats strfmt.Registry) error {

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

func (m *StorageCatalogIntegrationBody) validateRequiredServices(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredServices) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredServices); i++ {
		if swag.IsZero(m.RequiredServices[i]) { // not required
			continue
		}

		if m.RequiredServices[i] != nil {
			if err := m.RequiredServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageCatalogIntegrationBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *StorageCatalogIntegrationBody) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *StorageCatalogIntegrationBody) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageCatalogIntegrationBody) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCatalogIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCatalogIntegrationBody) UnmarshalBinary(b []byte) error {
	var res StorageCatalogIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
