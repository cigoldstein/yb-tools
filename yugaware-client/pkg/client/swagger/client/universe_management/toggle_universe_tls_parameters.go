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

// NewToggleUniverseTLSParams creates a new ToggleUniverseTLSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewToggleUniverseTLSParams() *ToggleUniverseTLSParams {
	return &ToggleUniverseTLSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewToggleUniverseTLSParamsWithTimeout creates a new ToggleUniverseTLSParams object
// with the ability to set a timeout on a request.
func NewToggleUniverseTLSParamsWithTimeout(timeout time.Duration) *ToggleUniverseTLSParams {
	return &ToggleUniverseTLSParams{
		timeout: timeout,
	}
}

// NewToggleUniverseTLSParamsWithContext creates a new ToggleUniverseTLSParams object
// with the ability to set a context for a request.
func NewToggleUniverseTLSParamsWithContext(ctx context.Context) *ToggleUniverseTLSParams {
	return &ToggleUniverseTLSParams{
		Context: ctx,
	}
}

// NewToggleUniverseTLSParamsWithHTTPClient creates a new ToggleUniverseTLSParams object
// with the ability to set a custom HTTPClient for a request.
func NewToggleUniverseTLSParamsWithHTTPClient(client *http.Client) *ToggleUniverseTLSParams {
	return &ToggleUniverseTLSParams{
		HTTPClient: client,
	}
}

/* ToggleUniverseTLSParams contains all the parameters to send to the API endpoint
   for the toggle universe TLS operation.

   Typically these are written to a http.Request.
*/
type ToggleUniverseTLSParams struct {

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

// WithDefaults hydrates default values in the toggle universe TLS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleUniverseTLSParams) WithDefaults() *ToggleUniverseTLSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the toggle universe TLS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleUniverseTLSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) WithTimeout(timeout time.Duration) *ToggleUniverseTLSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) WithContext(ctx context.Context) *ToggleUniverseTLSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) WithHTTPClient(client *http.Client) *ToggleUniverseTLSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) WithCUUID(cUUID strfmt.UUID) *ToggleUniverseTLSParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) WithUniUUID(uniUUID strfmt.UUID) *ToggleUniverseTLSParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the toggle universe TLS params
func (o *ToggleUniverseTLSParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleUniverseTLSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
