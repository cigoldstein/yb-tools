// Code generated by go-swagger; DO NOT EDIT.

package universe_cluster_mutations

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

// NewCreateReadOnlyClusterParams creates a new CreateReadOnlyClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateReadOnlyClusterParams() *CreateReadOnlyClusterParams {
	return &CreateReadOnlyClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReadOnlyClusterParamsWithTimeout creates a new CreateReadOnlyClusterParams object
// with the ability to set a timeout on a request.
func NewCreateReadOnlyClusterParamsWithTimeout(timeout time.Duration) *CreateReadOnlyClusterParams {
	return &CreateReadOnlyClusterParams{
		timeout: timeout,
	}
}

// NewCreateReadOnlyClusterParamsWithContext creates a new CreateReadOnlyClusterParams object
// with the ability to set a context for a request.
func NewCreateReadOnlyClusterParamsWithContext(ctx context.Context) *CreateReadOnlyClusterParams {
	return &CreateReadOnlyClusterParams{
		Context: ctx,
	}
}

// NewCreateReadOnlyClusterParamsWithHTTPClient creates a new CreateReadOnlyClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateReadOnlyClusterParamsWithHTTPClient(client *http.Client) *CreateReadOnlyClusterParams {
	return &CreateReadOnlyClusterParams{
		HTTPClient: client,
	}
}

/* CreateReadOnlyClusterParams contains all the parameters to send to the API endpoint
   for the create read only cluster operation.

   Typically these are written to a http.Request.
*/
type CreateReadOnlyClusterParams struct {

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

// WithDefaults hydrates default values in the create read only cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReadOnlyClusterParams) WithDefaults() *CreateReadOnlyClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create read only cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReadOnlyClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create read only cluster params
func (o *CreateReadOnlyClusterParams) WithTimeout(timeout time.Duration) *CreateReadOnlyClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create read only cluster params
func (o *CreateReadOnlyClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create read only cluster params
func (o *CreateReadOnlyClusterParams) WithContext(ctx context.Context) *CreateReadOnlyClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create read only cluster params
func (o *CreateReadOnlyClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create read only cluster params
func (o *CreateReadOnlyClusterParams) WithHTTPClient(client *http.Client) *CreateReadOnlyClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create read only cluster params
func (o *CreateReadOnlyClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the create read only cluster params
func (o *CreateReadOnlyClusterParams) WithCUUID(cUUID strfmt.UUID) *CreateReadOnlyClusterParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the create read only cluster params
func (o *CreateReadOnlyClusterParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the create read only cluster params
func (o *CreateReadOnlyClusterParams) WithUniUUID(uniUUID strfmt.UUID) *CreateReadOnlyClusterParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the create read only cluster params
func (o *CreateReadOnlyClusterParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReadOnlyClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
