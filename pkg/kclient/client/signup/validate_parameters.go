// Code generated by go-swagger; DO NOT EDIT.

package signup

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

// NewValidateParams creates a new ValidateParams object
// with the default values initialized.
func NewValidateParams() *ValidateParams {
	var ()
	return &ValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateParamsWithTimeout creates a new ValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateParamsWithTimeout(timeout time.Duration) *ValidateParams {
	var ()
	return &ValidateParams{

		timeout: timeout,
	}
}

// NewValidateParamsWithContext creates a new ValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateParamsWithContext(ctx context.Context) *ValidateParams {
	var ()
	return &ValidateParams{

		Context: ctx,
	}
}

// NewValidateParamsWithHTTPClient creates a new ValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateParamsWithHTTPClient(client *http.Client) *ValidateParams {
	var ()
	return &ValidateParams{
		HTTPClient: client,
	}
}

/*ValidateParams contains all the parameters to send to the API endpoint
for the validate operation typically these are written to a http.Request
*/
type ValidateParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate params
func (o *ValidateParams) WithTimeout(timeout time.Duration) *ValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate params
func (o *ValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate params
func (o *ValidateParams) WithContext(ctx context.Context) *ValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate params
func (o *ValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate params
func (o *ValidateParams) WithHTTPClient(client *http.Client) *ValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate params
func (o *ValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the validate params
func (o *ValidateParams) WithID(id string) *ValidateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate params
func (o *ValidateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
