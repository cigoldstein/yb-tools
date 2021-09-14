// Code generated by go-swagger; DO NOT EDIT.

package runtime_configuration

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

// NewListKeysParams creates a new ListKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListKeysParams() *ListKeysParams {
	return &ListKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListKeysParamsWithTimeout creates a new ListKeysParams object
// with the ability to set a timeout on a request.
func NewListKeysParamsWithTimeout(timeout time.Duration) *ListKeysParams {
	return &ListKeysParams{
		timeout: timeout,
	}
}

// NewListKeysParamsWithContext creates a new ListKeysParams object
// with the ability to set a context for a request.
func NewListKeysParamsWithContext(ctx context.Context) *ListKeysParams {
	return &ListKeysParams{
		Context: ctx,
	}
}

// NewListKeysParamsWithHTTPClient creates a new ListKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewListKeysParamsWithHTTPClient(client *http.Client) *ListKeysParams {
	return &ListKeysParams{
		HTTPClient: client,
	}
}

/* ListKeysParams contains all the parameters to send to the API endpoint
   for the list keys operation.

   Typically these are written to a http.Request.
*/
type ListKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKeysParams) WithDefaults() *ListKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list keys params
func (o *ListKeysParams) WithTimeout(timeout time.Duration) *ListKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list keys params
func (o *ListKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list keys params
func (o *ListKeysParams) WithContext(ctx context.Context) *ListKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list keys params
func (o *ListKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list keys params
func (o *ListKeysParams) WithHTTPClient(client *http.Client) *ListKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list keys params
func (o *ListKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
