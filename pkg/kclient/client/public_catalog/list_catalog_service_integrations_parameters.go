// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

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

// NewListCatalogServiceIntegrationsParams creates a new ListCatalogServiceIntegrationsParams object
// with the default values initialized.
func NewListCatalogServiceIntegrationsParams() *ListCatalogServiceIntegrationsParams {
	var ()
	return &ListCatalogServiceIntegrationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCatalogServiceIntegrationsParamsWithTimeout creates a new ListCatalogServiceIntegrationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCatalogServiceIntegrationsParamsWithTimeout(timeout time.Duration) *ListCatalogServiceIntegrationsParams {
	var ()
	return &ListCatalogServiceIntegrationsParams{

		timeout: timeout,
	}
}

// NewListCatalogServiceIntegrationsParamsWithContext creates a new ListCatalogServiceIntegrationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCatalogServiceIntegrationsParamsWithContext(ctx context.Context) *ListCatalogServiceIntegrationsParams {
	var ()
	return &ListCatalogServiceIntegrationsParams{

		Context: ctx,
	}
}

// NewListCatalogServiceIntegrationsParamsWithHTTPClient creates a new ListCatalogServiceIntegrationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCatalogServiceIntegrationsParamsWithHTTPClient(client *http.Client) *ListCatalogServiceIntegrationsParams {
	var ()
	return &ListCatalogServiceIntegrationsParams{
		HTTPClient: client,
	}
}

/*ListCatalogServiceIntegrationsParams contains all the parameters to send to the API endpoint
for the list catalog service integrations operation typically these are written to a http.Request
*/
type ListCatalogServiceIntegrationsParams struct {

	/*ID*/
	ID string
	/*Limit*/
	Limit *string
	/*Name*/
	Name *string
	/*Offset*/
	Offset *string
	/*Tags*/
	Tags *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithTimeout(timeout time.Duration) *ListCatalogServiceIntegrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithContext(ctx context.Context) *ListCatalogServiceIntegrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithHTTPClient(client *http.Client) *ListCatalogServiceIntegrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithID(id string) *ListCatalogServiceIntegrationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithLimit(limit *string) *ListCatalogServiceIntegrationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithName(name *string) *ListCatalogServiceIntegrationsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithOffset(offset *string) *ListCatalogServiceIntegrationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithTags adds the tags to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithTags(tags *string) *ListCatalogServiceIntegrationsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithType adds the typeVar to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) WithType(typeVar *string) *ListCatalogServiceIntegrationsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list catalog service integrations params
func (o *ListCatalogServiceIntegrationsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListCatalogServiceIntegrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Tags != nil {

		// query param tags
		var qrTags string
		if o.Tags != nil {
			qrTags = *o.Tags
		}
		qTags := qrTags
		if qTags != "" {
			if err := r.SetQueryParam("tags", qTags); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
