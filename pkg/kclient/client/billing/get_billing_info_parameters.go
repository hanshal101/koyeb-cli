// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewGetBillingInfoParams creates a new GetBillingInfoParams object
// with the default values initialized.
func NewGetBillingInfoParams() *GetBillingInfoParams {

	return &GetBillingInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingInfoParamsWithTimeout creates a new GetBillingInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillingInfoParamsWithTimeout(timeout time.Duration) *GetBillingInfoParams {

	return &GetBillingInfoParams{

		timeout: timeout,
	}
}

// NewGetBillingInfoParamsWithContext creates a new GetBillingInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillingInfoParamsWithContext(ctx context.Context) *GetBillingInfoParams {

	return &GetBillingInfoParams{

		Context: ctx,
	}
}

// NewGetBillingInfoParamsWithHTTPClient creates a new GetBillingInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillingInfoParamsWithHTTPClient(client *http.Client) *GetBillingInfoParams {

	return &GetBillingInfoParams{
		HTTPClient: client,
	}
}

/*GetBillingInfoParams contains all the parameters to send to the API endpoint
for the get billing info operation typically these are written to a http.Request
*/
type GetBillingInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get billing info params
func (o *GetBillingInfoParams) WithTimeout(timeout time.Duration) *GetBillingInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing info params
func (o *GetBillingInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing info params
func (o *GetBillingInfoParams) WithContext(ctx context.Context) *GetBillingInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing info params
func (o *GetBillingInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing info params
func (o *GetBillingInfoParams) WithHTTPClient(client *http.Client) *GetBillingInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing info params
func (o *GetBillingInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
