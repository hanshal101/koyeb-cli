// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBillingInfo(params *CreateBillingInfoParams) (*CreateBillingInfoOK, error)

	GetBillingInfo(params *GetBillingInfoParams) (*GetBillingInfoOK, error)

	UpdateBilling(params *UpdateBillingParams) (*UpdateBillingOK, error)

	UpdateBilling2(params *UpdateBilling2Params) (*UpdateBilling2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBillingInfo create billing info API
*/
func (a *Client) CreateBillingInfo(params *CreateBillingInfoParams) (*CreateBillingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBillingInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBillingInfo",
		Method:             "POST",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBillingInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBillingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBillingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBillingInfo get billing info API
*/
func (a *Client) GetBillingInfo(params *GetBillingInfoParams) (*GetBillingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBillingInfo",
		Method:             "GET",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBillingInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBillingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBillingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBilling update billing API
*/
func (a *Client) UpdateBilling(params *UpdateBillingParams) (*UpdateBillingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBillingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateBilling",
		Method:             "PUT",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBillingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBillingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBilling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBilling2 update billing2 API
*/
func (a *Client) UpdateBilling2(params *UpdateBilling2Params) (*UpdateBilling2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBilling2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateBilling2",
		Method:             "PATCH",
		PathPattern:        "/v1/account/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBilling2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBilling2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBilling2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
