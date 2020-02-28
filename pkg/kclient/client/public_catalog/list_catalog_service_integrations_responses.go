// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ListCatalogServiceIntegrationsReader is a Reader for the ListCatalogServiceIntegrations structure.
type ListCatalogServiceIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCatalogServiceIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCatalogServiceIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCatalogServiceIntegrationsOK creates a ListCatalogServiceIntegrationsOK with default headers values
func NewListCatalogServiceIntegrationsOK() *ListCatalogServiceIntegrationsOK {
	return &ListCatalogServiceIntegrationsOK{}
}

/*ListCatalogServiceIntegrationsOK handles this case with default header values.

A successful response.
*/
type ListCatalogServiceIntegrationsOK struct {
	Payload *models.StorageListCatalogIntegrationsReply
}

func (o *ListCatalogServiceIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/public/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListCatalogServiceIntegrationsOK) GetPayload() *models.StorageListCatalogIntegrationsReply {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListCatalogIntegrationsReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
