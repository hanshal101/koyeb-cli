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

// OrganizationGetOrganizationReader is a Reader for the OrganizationGetOrganization structure.
type OrganizationGetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationGetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationGetOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationGetOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationGetOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationGetOrganizationOK creates a OrganizationGetOrganizationOK with default headers values
func NewOrganizationGetOrganizationOK() *OrganizationGetOrganizationOK {
	return &OrganizationGetOrganizationOK{}
}

/*OrganizationGetOrganizationOK handles this case with default header values.

A successful response.
*/
type OrganizationGetOrganizationOK struct {
	Payload *models.AccountGetOrganizationReply
}

func (o *OrganizationGetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] organizationGetOrganizationOK  %+v", 200, o.Payload)
}

func (o *OrganizationGetOrganizationOK) GetPayload() *models.AccountGetOrganizationReply {
	return o.Payload
}

func (o *OrganizationGetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountGetOrganizationReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetOrganizationBadRequest creates a OrganizationGetOrganizationBadRequest with default headers values
func NewOrganizationGetOrganizationBadRequest() *OrganizationGetOrganizationBadRequest {
	return &OrganizationGetOrganizationBadRequest{}
}

/*OrganizationGetOrganizationBadRequest handles this case with default header values.

Validation error
*/
type OrganizationGetOrganizationBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *OrganizationGetOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] organizationGetOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationGetOrganizationBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *OrganizationGetOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetOrganizationForbidden creates a OrganizationGetOrganizationForbidden with default headers values
func NewOrganizationGetOrganizationForbidden() *OrganizationGetOrganizationForbidden {
	return &OrganizationGetOrganizationForbidden{}
}

/*OrganizationGetOrganizationForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type OrganizationGetOrganizationForbidden struct {
	Payload *models.CommonError
}

func (o *OrganizationGetOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] organizationGetOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationGetOrganizationForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGetOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetOrganizationNotFound creates a OrganizationGetOrganizationNotFound with default headers values
func NewOrganizationGetOrganizationNotFound() *OrganizationGetOrganizationNotFound {
	return &OrganizationGetOrganizationNotFound{}
}

/*OrganizationGetOrganizationNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type OrganizationGetOrganizationNotFound struct {
	Payload *models.CommonError
}

func (o *OrganizationGetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] organizationGetOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationGetOrganizationNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetOrganizationDefault creates a OrganizationGetOrganizationDefault with default headers values
func NewOrganizationGetOrganizationDefault(code int) *OrganizationGetOrganizationDefault {
	return &OrganizationGetOrganizationDefault{
		_statusCode: code,
	}
}

/*OrganizationGetOrganizationDefault handles this case with default header values.

An unexpected error response
*/
type OrganizationGetOrganizationDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the organization get organization default response
func (o *OrganizationGetOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationGetOrganizationDefault) Error() string {
	return fmt.Sprintf("[GET /v1/organizations/{id}][%d] organization_GetOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationGetOrganizationDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *OrganizationGetOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
