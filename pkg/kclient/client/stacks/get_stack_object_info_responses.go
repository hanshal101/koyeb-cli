// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// GetStackObjectInfoReader is a Reader for the GetStackObjectInfo structure.
type GetStackObjectInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackObjectInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStackObjectInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStackObjectInfoOK creates a GetStackObjectInfoOK with default headers values
func NewGetStackObjectInfoOK() *GetStackObjectInfoOK {
	return &GetStackObjectInfoOK{}
}

/*GetStackObjectInfoOK handles this case with default header values.

A successful response.
*/
type GetStackObjectInfoOK struct {
	Payload *models.StorageGetStackObjectInfoReply
}

func (o *GetStackObjectInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/info/{path}][%d] getStackObjectInfoOK  %+v", 200, o.Payload)
}

func (o *GetStackObjectInfoOK) GetPayload() *models.StorageGetStackObjectInfoReply {
	return o.Payload
}

func (o *GetStackObjectInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageGetStackObjectInfoReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
