// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// UpdateCredentialReader is a Reader for the UpdateCredential structure.
type UpdateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCredentialOK creates a UpdateCredentialOK with default headers values
func NewUpdateCredentialOK() *UpdateCredentialOK {
	return &UpdateCredentialOK{}
}

/*UpdateCredentialOK handles this case with default header values.

A successful response.
*/
type UpdateCredentialOK struct {
	Payload *models.AccountCredentialReply
}

func (o *UpdateCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/{id}][%d] updateCredentialOK  %+v", 200, o.Payload)
}

func (o *UpdateCredentialOK) GetPayload() *models.AccountCredentialReply {
	return o.Payload
}

func (o *UpdateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountCredentialReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
