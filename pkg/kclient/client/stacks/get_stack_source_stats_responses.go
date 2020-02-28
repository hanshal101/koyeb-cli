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

// GetStackSourceStatsReader is a Reader for the GetStackSourceStats structure.
type GetStackSourceStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackSourceStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStackSourceStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStackSourceStatsOK creates a GetStackSourceStatsOK with default headers values
func NewGetStackSourceStatsOK() *GetStackSourceStatsOK {
	return &GetStackSourceStatsOK{}
}

/*GetStackSourceStatsOK handles this case with default header values.

A successful response.
*/
type GetStackSourceStatsOK struct {
	Payload models.StorageStackSourceStatsReply
}

func (o *GetStackSourceStatsOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{stack_id}/sources/{id}/stats][%d] getStackSourceStatsOK  %+v", 200, o.Payload)
}

func (o *GetStackSourceStatsOK) GetPayload() models.StorageStackSourceStatsReply {
	return o.Payload
}

func (o *GetStackSourceStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
