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

// NewSetUniverseHelm3CompatibleParams creates a new SetUniverseHelm3CompatibleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetUniverseHelm3CompatibleParams() *SetUniverseHelm3CompatibleParams {
	return &SetUniverseHelm3CompatibleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetUniverseHelm3CompatibleParamsWithTimeout creates a new SetUniverseHelm3CompatibleParams object
// with the ability to set a timeout on a request.
func NewSetUniverseHelm3CompatibleParamsWithTimeout(timeout time.Duration) *SetUniverseHelm3CompatibleParams {
	return &SetUniverseHelm3CompatibleParams{
		timeout: timeout,
	}
}

// NewSetUniverseHelm3CompatibleParamsWithContext creates a new SetUniverseHelm3CompatibleParams object
// with the ability to set a context for a request.
func NewSetUniverseHelm3CompatibleParamsWithContext(ctx context.Context) *SetUniverseHelm3CompatibleParams {
	return &SetUniverseHelm3CompatibleParams{
		Context: ctx,
	}
}

// NewSetUniverseHelm3CompatibleParamsWithHTTPClient creates a new SetUniverseHelm3CompatibleParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetUniverseHelm3CompatibleParamsWithHTTPClient(client *http.Client) *SetUniverseHelm3CompatibleParams {
	return &SetUniverseHelm3CompatibleParams{
		HTTPClient: client,
	}
}

/* SetUniverseHelm3CompatibleParams contains all the parameters to send to the API endpoint
   for the set universe helm3 compatible operation.

   Typically these are written to a http.Request.
*/
type SetUniverseHelm3CompatibleParams struct {

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

// WithDefaults hydrates default values in the set universe helm3 compatible params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUniverseHelm3CompatibleParams) WithDefaults() *SetUniverseHelm3CompatibleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set universe helm3 compatible params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUniverseHelm3CompatibleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) WithTimeout(timeout time.Duration) *SetUniverseHelm3CompatibleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) WithContext(ctx context.Context) *SetUniverseHelm3CompatibleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) WithHTTPClient(client *http.Client) *SetUniverseHelm3CompatibleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) WithCUUID(cUUID strfmt.UUID) *SetUniverseHelm3CompatibleParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) WithUniUUID(uniUUID strfmt.UUID) *SetUniverseHelm3CompatibleParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the set universe helm3 compatible params
func (o *SetUniverseHelm3CompatibleParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SetUniverseHelm3CompatibleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
