// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/*LogoutOK handles this case with default header values.

A successful response.
*/
type LogoutOK struct {
	Payload models.AccountLogoutReply
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/account/logout][%d] logoutOK  %+v", 200, o.Payload)
}

func (o *LogoutOK) GetPayload() models.AccountLogoutReply {
	return o.Payload
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
