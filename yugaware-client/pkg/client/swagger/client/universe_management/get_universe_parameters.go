// Code generated by go-swagger; DO NOT EDIT.

package universe_management

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

// NewGetUniverseParams creates a new GetUniverseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseParams() *GetUniverseParams {
	return &GetUniverseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseParamsWithTimeout creates a new GetUniverseParams object
// with the ability to set a timeout on a request.
func NewGetUniverseParamsWithTimeout(timeout time.Duration) *GetUniverseParams {
	return &GetUniverseParams{
		timeout: timeout,
	}
}

// NewGetUniverseParamsWithContext creates a new GetUniverseParams object
// with the ability to set a context for a request.
func NewGetUniverseParamsWithContext(ctx context.Context) *GetUniverseParams {
	return &GetUniverseParams{
		Context: ctx,
	}
}

// NewGetUniverseParamsWithHTTPClient creates a new GetUniverseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseParamsWithHTTPClient(client *http.Client) *GetUniverseParams {
	return &GetUniverseParams{
		HTTPClient: client,
	}
}

/* GetUniverseParams contains all the parameters to send to the API endpoint
   for the get universe operation.

   Typically these are written to a http.Request.
*/
type GetUniverseParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// UniUUID.
	//
	// Format: uuid
	UniUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseParams) WithDefaults() *GetUniverseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get universe params
func (o *GetUniverseParams) WithTimeout(timeout time.Duration) *GetUniverseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe params
func (o *GetUniverseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe params
func (o *GetUniverseParams) WithContext(ctx context.Context) *GetUniverseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe params
func (o *GetUniverseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe params
func (o *GetUniverseParams) WithHTTPClient(client *http.Client) *GetUniverseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe params
func (o *GetUniverseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get universe params
func (o *GetUniverseParams) WithCUUID(cUUID strfmt.UUID) *GetUniverseParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get universe params
func (o *GetUniverseParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the get universe params
func (o *GetUniverseParams) WithUniUUID(uniUUID strfmt.UUID) *GetUniverseParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the get universe params
func (o *GetUniverseParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param uniUUID
	if err := r.SetPathParam("uniUUID", o.UniUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}