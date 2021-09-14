// Code generated by go-swagger; DO NOT EDIT.

package availability_zones

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

// NewListOfAZParams creates a new ListOfAZParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOfAZParams() *ListOfAZParams {
	return &ListOfAZParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOfAZParamsWithTimeout creates a new ListOfAZParams object
// with the ability to set a timeout on a request.
func NewListOfAZParamsWithTimeout(timeout time.Duration) *ListOfAZParams {
	return &ListOfAZParams{
		timeout: timeout,
	}
}

// NewListOfAZParamsWithContext creates a new ListOfAZParams object
// with the ability to set a context for a request.
func NewListOfAZParamsWithContext(ctx context.Context) *ListOfAZParams {
	return &ListOfAZParams{
		Context: ctx,
	}
}

// NewListOfAZParamsWithHTTPClient creates a new ListOfAZParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOfAZParamsWithHTTPClient(client *http.Client) *ListOfAZParams {
	return &ListOfAZParams{
		HTTPClient: client,
	}
}

/* ListOfAZParams contains all the parameters to send to the API endpoint
   for the list of a z operation.

   Typically these are written to a http.Request.
*/
type ListOfAZParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// PUUID.
	//
	// Format: uuid
	PUUID strfmt.UUID

	// RUUID.
	//
	// Format: uuid
	RUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list of a z params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOfAZParams) WithDefaults() *ListOfAZParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list of a z params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOfAZParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list of a z params
func (o *ListOfAZParams) WithTimeout(timeout time.Duration) *ListOfAZParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list of a z params
func (o *ListOfAZParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list of a z params
func (o *ListOfAZParams) WithContext(ctx context.Context) *ListOfAZParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list of a z params
func (o *ListOfAZParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list of a z params
func (o *ListOfAZParams) WithHTTPClient(client *http.Client) *ListOfAZParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list of a z params
func (o *ListOfAZParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the list of a z params
func (o *ListOfAZParams) WithCUUID(cUUID strfmt.UUID) *ListOfAZParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the list of a z params
func (o *ListOfAZParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithPUUID adds the pUUID to the list of a z params
func (o *ListOfAZParams) WithPUUID(pUUID strfmt.UUID) *ListOfAZParams {
	o.SetPUUID(pUUID)
	return o
}

// SetPUUID adds the pUuid to the list of a z params
func (o *ListOfAZParams) SetPUUID(pUUID strfmt.UUID) {
	o.PUUID = pUUID
}

// WithRUUID adds the rUUID to the list of a z params
func (o *ListOfAZParams) WithRUUID(rUUID strfmt.UUID) *ListOfAZParams {
	o.SetRUUID(rUUID)
	return o
}

// SetRUUID adds the rUuid to the list of a z params
func (o *ListOfAZParams) SetRUUID(rUUID strfmt.UUID) {
	o.RUUID = rUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOfAZParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param pUUID
	if err := r.SetPathParam("pUUID", o.PUUID.String()); err != nil {
		return err
	}

	// path param rUUID
	if err := r.SetPathParam("rUUID", o.RUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
