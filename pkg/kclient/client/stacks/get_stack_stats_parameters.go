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

// NewGetStackStatsParams creates a new GetStackStatsParams object
// with the default values initialized.
func NewGetStackStatsParams() *GetStackStatsParams {
	var ()
	return &GetStackStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackStatsParamsWithTimeout creates a new GetStackStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStackStatsParamsWithTimeout(timeout time.Duration) *GetStackStatsParams {
	var ()
	return &GetStackStatsParams{

		timeout: timeout,
	}
}

// NewGetStackStatsParamsWithContext creates a new GetStackStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStackStatsParamsWithContext(ctx context.Context) *GetStackStatsParams {
	var ()
	return &GetStackStatsParams{

		Context: ctx,
	}
}

// NewGetStackStatsParamsWithHTTPClient creates a new GetStackStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStackStatsParamsWithHTTPClient(client *http.Client) *GetStackStatsParams {
	var ()
	return &GetStackStatsParams{
		HTTPClient: client,
	}
}

/*GetStackStatsParams contains all the parameters to send to the API endpoint
for the get stack stats operation typically these are written to a http.Request
*/
type GetStackStatsParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stack stats params
func (o *GetStackStatsParams) WithTimeout(timeout time.Duration) *GetStackStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack stats params
func (o *GetStackStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack stats params
func (o *GetStackStatsParams) WithContext(ctx context.Context) *GetStackStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack stats params
func (o *GetStackStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack stats params
func (o *GetStackStatsParams) WithHTTPClient(client *http.Client) *GetStackStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack stats params
func (o *GetStackStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get stack stats params
func (o *GetStackStatsParams) WithID(id string) *GetStackStatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stack stats params
func (o *GetStackStatsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
