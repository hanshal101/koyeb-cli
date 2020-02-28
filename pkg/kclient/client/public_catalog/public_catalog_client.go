// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCatalogIntegration(params *GetCatalogIntegrationParams) (*GetCatalogIntegrationOK, error)

	GetCatalogService(params *GetCatalogServiceParams) (*GetCatalogServiceOK, error)

	ListCatalogIntegrations(params *ListCatalogIntegrationsParams) (*ListCatalogIntegrationsOK, error)

	ListCatalogServiceIntegrations(params *ListCatalogServiceIntegrationsParams) (*ListCatalogServiceIntegrationsOK, error)

	ListCatalogServicesMixin0(params *ListCatalogServicesMixin0Params) (*ListCatalogServicesMixin0OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCatalogIntegration get catalog integration API
*/
func (a *Client) GetCatalogIntegration(params *GetCatalogIntegrationParams) (*GetCatalogIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogIntegrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogIntegration",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/integrations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogIntegrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCatalogIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCatalogService get catalog service API
*/
func (a *Client) GetCatalogService(params *GetCatalogServiceParams) (*GetCatalogServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogService",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCatalogService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogIntegrations list catalog integrations API
*/
func (a *Client) ListCatalogIntegrations(params *ListCatalogIntegrationsParams) (*ListCatalogIntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogIntegrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogIntegrations",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogIntegrationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogIntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogIntegrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogServiceIntegrations list catalog service integrations API
*/
func (a *Client) ListCatalogServiceIntegrations(params *ListCatalogServiceIntegrationsParams) (*ListCatalogServiceIntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogServiceIntegrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogServiceIntegrations",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services/{id}/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogServiceIntegrationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogServiceIntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogServiceIntegrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogServicesMixin0 list catalog services mixin0 API
*/
func (a *Client) ListCatalogServicesMixin0(params *ListCatalogServicesMixin0Params) (*ListCatalogServicesMixin0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogServicesMixin0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogServicesMixin0",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogServicesMixin0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogServicesMixin0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogServicesMixin0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
