// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// OrganizationUpdateOrganizationReader is a Reader for the OrganizationUpdateOrganization structure.
type OrganizationUpdateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationUpdateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationUpdateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationUpdateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationUpdateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationUpdateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationUpdateOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationUpdateOrganizationOK creates a OrganizationUpdateOrganizationOK with default headers values
func NewOrganizationUpdateOrganizationOK() *OrganizationUpdateOrganizationOK {
	return &OrganizationUpdateOrganizationOK{}
}

/*OrganizationUpdateOrganizationOK handles this case with default header values.

A successful response.
*/
type OrganizationUpdateOrganizationOK struct {
	Payload *models.AccountUpdateOrganizationReply
}

func (o *OrganizationUpdateOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] organizationUpdateOrganizationOK  %+v", 200, o.Payload)
}

func (o *OrganizationUpdateOrganizationOK) GetPayload() *models.AccountUpdateOrganizationReply {
	return o.Payload
}

func (o *OrganizationUpdateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountUpdateOrganizationReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateOrganizationBadRequest creates a OrganizationUpdateOrganizationBadRequest with default headers values
func NewOrganizationUpdateOrganizationBadRequest() *OrganizationUpdateOrganizationBadRequest {
	return &OrganizationUpdateOrganizationBadRequest{}
}

/*OrganizationUpdateOrganizationBadRequest handles this case with default header values.

Validation error
*/
type OrganizationUpdateOrganizationBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *OrganizationUpdateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] organizationUpdateOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationUpdateOrganizationBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *OrganizationUpdateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateOrganizationForbidden creates a OrganizationUpdateOrganizationForbidden with default headers values
func NewOrganizationUpdateOrganizationForbidden() *OrganizationUpdateOrganizationForbidden {
	return &OrganizationUpdateOrganizationForbidden{}
}

/*OrganizationUpdateOrganizationForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type OrganizationUpdateOrganizationForbidden struct {
	Payload *models.CommonError
}

func (o *OrganizationUpdateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] organizationUpdateOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationUpdateOrganizationForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationUpdateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateOrganizationNotFound creates a OrganizationUpdateOrganizationNotFound with default headers values
func NewOrganizationUpdateOrganizationNotFound() *OrganizationUpdateOrganizationNotFound {
	return &OrganizationUpdateOrganizationNotFound{}
}

/*OrganizationUpdateOrganizationNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type OrganizationUpdateOrganizationNotFound struct {
	Payload *models.CommonError
}

func (o *OrganizationUpdateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] organizationUpdateOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationUpdateOrganizationNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationUpdateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationUpdateOrganizationDefault creates a OrganizationUpdateOrganizationDefault with default headers values
func NewOrganizationUpdateOrganizationDefault(code int) *OrganizationUpdateOrganizationDefault {
	return &OrganizationUpdateOrganizationDefault{
		_statusCode: code,
	}
}

/*OrganizationUpdateOrganizationDefault handles this case with default header values.

An unexpected error response
*/
type OrganizationUpdateOrganizationDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the organization update organization default response
func (o *OrganizationUpdateOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationUpdateOrganizationDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/{id}][%d] organization_UpdateOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationUpdateOrganizationDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *OrganizationUpdateOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
