// Code generated by go-swagger; DO NOT EDIT.

package instance_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetGCPTypesParams creates a new GetGCPTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGCPTypesParams() *GetGCPTypesParams {
	return &GetGCPTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGCPTypesParamsWithTimeout creates a new GetGCPTypesParams object
// with the ability to set a timeout on a request.
func NewGetGCPTypesParamsWithTimeout(timeout time.Duration) *GetGCPTypesParams {
	return &GetGCPTypesParams{
		timeout: timeout,
	}
}

// NewGetGCPTypesParamsWithContext creates a new GetGCPTypesParams object
// with the ability to set a context for a request.
func NewGetGCPTypesParamsWithContext(ctx context.Context) *GetGCPTypesParams {
	return &GetGCPTypesParams{
		Context: ctx,
	}
}

// NewGetGCPTypesParamsWithHTTPClient creates a new GetGCPTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGCPTypesParamsWithHTTPClient(client *http.Client) *GetGCPTypesParams {
	return &GetGCPTypesParams{
		HTTPClient: client,
	}
}

/* GetGCPTypesParams contains all the parameters to send to the API endpoint
   for the get g c p types operation.

   Typically these are written to a http.Request.
*/
type GetGCPTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get g c p types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGCPTypesParams) WithDefaults() *GetGCPTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get g c p types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGCPTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get g c p types params
func (o *GetGCPTypesParams) WithTimeout(timeout time.Duration) *GetGCPTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get g c p types params
func (o *GetGCPTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get g c p types params
func (o *GetGCPTypesParams) WithContext(ctx context.Context) *GetGCPTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get g c p types params
func (o *GetGCPTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get g c p types params
func (o *GetGCPTypesParams) WithHTTPClient(client *http.Client) *GetGCPTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get g c p types params
func (o *GetGCPTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGCPTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
