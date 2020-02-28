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

// NewEnableStackSourceParams creates a new EnableStackSourceParams object
// with the default values initialized.
func NewEnableStackSourceParams() *EnableStackSourceParams {
	var ()
	return &EnableStackSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableStackSourceParamsWithTimeout creates a new EnableStackSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableStackSourceParamsWithTimeout(timeout time.Duration) *EnableStackSourceParams {
	var ()
	return &EnableStackSourceParams{

		timeout: timeout,
	}
}

// NewEnableStackSourceParamsWithContext creates a new EnableStackSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableStackSourceParamsWithContext(ctx context.Context) *EnableStackSourceParams {
	var ()
	return &EnableStackSourceParams{

		Context: ctx,
	}
}

// NewEnableStackSourceParamsWithHTTPClient creates a new EnableStackSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableStackSourceParamsWithHTTPClient(client *http.Client) *EnableStackSourceParams {
	var ()
	return &EnableStackSourceParams{
		HTTPClient: client,
	}
}

/*EnableStackSourceParams contains all the parameters to send to the API endpoint
for the enable stack source operation typically these are written to a http.Request
*/
type EnableStackSourceParams struct {

	/*ID*/
	ID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable stack source params
func (o *EnableStackSourceParams) WithTimeout(timeout time.Duration) *EnableStackSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable stack source params
func (o *EnableStackSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable stack source params
func (o *EnableStackSourceParams) WithContext(ctx context.Context) *EnableStackSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable stack source params
func (o *EnableStackSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable stack source params
func (o *EnableStackSourceParams) WithHTTPClient(client *http.Client) *EnableStackSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable stack source params
func (o *EnableStackSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the enable stack source params
func (o *EnableStackSourceParams) WithID(id string) *EnableStackSourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enable stack source params
func (o *EnableStackSourceParams) SetID(id string) {
	o.ID = id
}

// WithStackID adds the stackID to the enable stack source params
func (o *EnableStackSourceParams) WithStackID(stackID string) *EnableStackSourceParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the enable stack source params
func (o *EnableStackSourceParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableStackSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
