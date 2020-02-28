// Code generated by go-swagger; DO NOT EDIT.

package deliveries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// UpdateDelivery2Reader is a Reader for the UpdateDelivery2 structure.
type UpdateDelivery2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDelivery2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDelivery2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDelivery2OK creates a UpdateDelivery2OK with default headers values
func NewUpdateDelivery2OK() *UpdateDelivery2OK {
	return &UpdateDelivery2OK{}
}

/*UpdateDelivery2OK handles this case with default header values.

A successful response.
*/
type UpdateDelivery2OK struct {
	Payload *models.StorageDeliveryReply
}

func (o *UpdateDelivery2OK) Error() string {
	return fmt.Sprintf("[PATCH /v1/deliveries/{id}][%d] updateDelivery2OK  %+v", 200, o.Payload)
}

func (o *UpdateDelivery2OK) GetPayload() *models.StorageDeliveryReply {
	return o.Payload
}

func (o *UpdateDelivery2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageDeliveryReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
