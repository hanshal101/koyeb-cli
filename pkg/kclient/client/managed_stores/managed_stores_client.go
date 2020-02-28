// Code generated by go-swagger; DO NOT EDIT.

package managed_stores

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new managed stores API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for managed stores API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteManagedStore(params *DeleteManagedStoreParams) (*DeleteManagedStoreOK, error)

	GetManagedStore(params *GetManagedStoreParams) (*GetManagedStoreOK, error)

	ListManagedStores(params *ListManagedStoresParams) (*ListManagedStoresOK, error)

	NewManagedStore(params *NewManagedStoreParams) (*NewManagedStoreOK, error)

	UpdateManagedStore(params *UpdateManagedStoreParams) (*UpdateManagedStoreOK, error)

	UpdateManagedStore2(params *UpdateManagedStore2Params) (*UpdateManagedStore2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteManagedStore delete managed store API
*/
func (a *Client) DeleteManagedStore(params *DeleteManagedStoreParams) (*DeleteManagedStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagedStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteManagedStore",
		Method:             "DELETE",
		PathPattern:        "/v1/managed_stores/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteManagedStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteManagedStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteManagedStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetManagedStore get managed store API
*/
func (a *Client) GetManagedStore(params *GetManagedStoreParams) (*GetManagedStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagedStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManagedStore",
		Method:             "GET",
		PathPattern:        "/v1/managed_stores/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetManagedStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManagedStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetManagedStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListManagedStores list managed stores API
*/
func (a *Client) ListManagedStores(params *ListManagedStoresParams) (*ListManagedStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListManagedStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListManagedStores",
		Method:             "GET",
		PathPattern:        "/v1/managed_stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListManagedStoresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListManagedStoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListManagedStores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NewManagedStore new managed store API
*/
func (a *Client) NewManagedStore(params *NewManagedStoreParams) (*NewManagedStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNewManagedStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NewManagedStore",
		Method:             "POST",
		PathPattern:        "/v1/managed_stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NewManagedStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NewManagedStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NewManagedStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateManagedStore update managed store API
*/
func (a *Client) UpdateManagedStore(params *UpdateManagedStoreParams) (*UpdateManagedStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateManagedStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateManagedStore",
		Method:             "PUT",
		PathPattern:        "/v1/managed_stores/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateManagedStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateManagedStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateManagedStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateManagedStore2 update managed store2 API
*/
func (a *Client) UpdateManagedStore2(params *UpdateManagedStore2Params) (*UpdateManagedStore2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateManagedStore2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateManagedStore2",
		Method:             "PATCH",
		PathPattern:        "/v1/managed_stores/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateManagedStore2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateManagedStore2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateManagedStore2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
