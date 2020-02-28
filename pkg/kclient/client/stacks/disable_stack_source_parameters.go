// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDisableStackSourceParams creates a new DisableStackSourceParams object
// with the default values initialized.
func NewDisableStackSourceParams() *DisableStackSourceParams {
	var ()
	return &DisableStackSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableStackSourceParamsWithTimeout creates a new DisableStackSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableStackSourceParamsWithTimeout(timeout time.Duration) *DisableStackSourceParams {
	var ()
	return &DisableStackSourceParams{

		timeout: timeout,
	}
}

// NewDisableStackSourceParamsWithContext creates a new DisableStackSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableStackSourceParamsWithContext(ctx context.Context) *DisableStackSourceParams {
	var ()
	return &DisableStackSourceParams{

		Context: ctx,
	}
}

// NewDisableStackSourceParamsWithHTTPClient creates a new DisableStackSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableStackSourceParamsWithHTTPClient(client *http.Client) *DisableStackSourceParams {
	var ()
	return &DisableStackSourceParams{
		HTTPClient: client,
	}
}

/*DisableStackSourceParams contains all the parameters to send to the API endpoint
for the disable stack source operation typically these are written to a http.Request
*/
type DisableStackSourceParams struct {

	/*ID*/
	ID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable stack source params
func (o *DisableStackSourceParams) WithTimeout(timeout time.Duration) *DisableStackSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable stack source params
func (o *DisableStackSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable stack source params
func (o *DisableStackSourceParams) WithContext(ctx context.Context) *DisableStackSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable stack source params
func (o *DisableStackSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable stack source params
func (o *DisableStackSourceParams) WithHTTPClient(client *http.Client) *DisableStackSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable stack source params
func (o *DisableStackSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the disable stack source params
func (o *DisableStackSourceParams) WithID(id string) *DisableStackSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the disable stack source params
func (o *DisableStackSourceParams) SetID(id string) {
	o.ID = id
}

// WithStackID adds the stackID to the disable stack source params
func (o *DisableStackSourceParams) WithStackID(stackID string) *DisableStackSourceParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the disable stack source params
func (o *DisableStackSourceParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DisableStackSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
