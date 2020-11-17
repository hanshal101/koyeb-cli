// Code generated by go-swagger; DO NOT EDIT.

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// StackCancelBuildReader is a Reader for the StackCancelBuild structure.
type StackCancelBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackCancelBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackCancelBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackCancelBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackCancelBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackCancelBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStackCancelBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStackCancelBuildOK creates a StackCancelBuildOK with default headers values
func NewStackCancelBuildOK() *StackCancelBuildOK {
	return &StackCancelBuildOK{}
}

/*StackCancelBuildOK handles this case with default header values.

A successful response.
*/
type StackCancelBuildOK struct {
	Payload models.StorageCancelBuildReply
}

func (o *StackCancelBuildOK) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/build/cancel][%d] stackCancelBuildOK  %+v", 200, o.Payload)
}

func (o *StackCancelBuildOK) GetPayload() models.StorageCancelBuildReply {
	return o.Payload
}

func (o *StackCancelBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCancelBuildBadRequest creates a StackCancelBuildBadRequest with default headers values
func NewStackCancelBuildBadRequest() *StackCancelBuildBadRequest {
	return &StackCancelBuildBadRequest{}
}

/*StackCancelBuildBadRequest handles this case with default header values.

Validation error
*/
type StackCancelBuildBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *StackCancelBuildBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/build/cancel][%d] stackCancelBuildBadRequest  %+v", 400, o.Payload)
}

func (o *StackCancelBuildBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *StackCancelBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCancelBuildForbidden creates a StackCancelBuildForbidden with default headers values
func NewStackCancelBuildForbidden() *StackCancelBuildForbidden {
	return &StackCancelBuildForbidden{}
}

/*StackCancelBuildForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type StackCancelBuildForbidden struct {
	Payload *models.CommonError
}

func (o *StackCancelBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/build/cancel][%d] stackCancelBuildForbidden  %+v", 403, o.Payload)
}

func (o *StackCancelBuildForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *StackCancelBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCancelBuildNotFound creates a StackCancelBuildNotFound with default headers values
func NewStackCancelBuildNotFound() *StackCancelBuildNotFound {
	return &StackCancelBuildNotFound{}
}

/*StackCancelBuildNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type StackCancelBuildNotFound struct {
	Payload *models.CommonError
}

func (o *StackCancelBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/build/cancel][%d] stackCancelBuildNotFound  %+v", 404, o.Payload)
}

func (o *StackCancelBuildNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *StackCancelBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCancelBuildDefault creates a StackCancelBuildDefault with default headers values
func NewStackCancelBuildDefault(code int) *StackCancelBuildDefault {
	return &StackCancelBuildDefault{
		_statusCode: code,
	}
}

/*StackCancelBuildDefault handles this case with default header values.

An unexpected error response
*/
type StackCancelBuildDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the stack cancel build default response
func (o *StackCancelBuildDefault) Code() int {
	return o._statusCode
}

func (o *StackCancelBuildDefault) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions/{sha}/build/cancel][%d] stack_CancelBuild default  %+v", o._statusCode, o.Payload)
}

func (o *StackCancelBuildDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *StackCancelBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
