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

// NewGetStackWorkflowParams creates a new GetStackWorkflowParams object
// with the default values initialized.
func NewGetStackWorkflowParams() *GetStackWorkflowParams {
	var ()
	return &GetStackWorkflowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackWorkflowParamsWithTimeout creates a new GetStackWorkflowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackWorkflowParamsWithTimeout(timeout time.Duration) *GetStackWorkflowParams {
	var ()
	return &GetStackWorkflowParams{

		timeout: timeout,
	}
}

// NewGetStackWorkflowParamsWithContext creates a new GetStackWorkflowParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackWorkflowParamsWithContext(ctx context.Context) *GetStackWorkflowParams {
	var ()
	return &GetStackWorkflowParams{

		Context: ctx,
	}
}

// NewGetStackWorkflowParamsWithHTTPClient creates a new GetStackWorkflowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackWorkflowParamsWithHTTPClient(client *http.Client) *GetStackWorkflowParams {
	var ()
	return &GetStackWorkflowParams{
		HTTPClient: client,
	}
}

/*GetStackWorkflowParams contains all the parameters to send to the API endpoint
for the get stack workflow operation typically these are written to a http.Request
*/
type GetStackWorkflowParams struct {

	/*ID*/
	ID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack workflow params
func (o *GetStackWorkflowParams) WithTimeout(timeout time.Duration) *GetStackWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack workflow params
func (o *GetStackWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack workflow params
func (o *GetStackWorkflowParams) WithContext(ctx context.Context) *GetStackWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack workflow params
func (o *GetStackWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack workflow params
func (o *GetStackWorkflowParams) WithHTTPClient(client *http.Client) *GetStackWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack workflow params
func (o *GetStackWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get stack workflow params
func (o *GetStackWorkflowParams) WithID(id string) *GetStackWorkflowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stack workflow params
func (o *GetStackWorkflowParams) SetID(id string) {
	o.ID = id
}

// WithStackID adds the stackID to the get stack workflow params
func (o *GetStackWorkflowParams) WithStackID(stackID string) *GetStackWorkflowParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get stack workflow params
func (o *GetStackWorkflowParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
