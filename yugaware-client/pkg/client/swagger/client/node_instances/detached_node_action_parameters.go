// Code generated by go-swagger; DO NOT EDIT.

package node_instances

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

// NewDetachedNodeActionParams creates a new DetachedNodeActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetachedNodeActionParams() *DetachedNodeActionParams {
	return &DetachedNodeActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetachedNodeActionParamsWithTimeout creates a new DetachedNodeActionParams object
// with the ability to set a timeout on a request.
func NewDetachedNodeActionParamsWithTimeout(timeout time.Duration) *DetachedNodeActionParams {
	return &DetachedNodeActionParams{
		timeout: timeout,
	}
}

// NewDetachedNodeActionParamsWithContext creates a new DetachedNodeActionParams object
// with the ability to set a context for a request.
func NewDetachedNodeActionParamsWithContext(ctx context.Context) *DetachedNodeActionParams {
	return &DetachedNodeActionParams{
		Context: ctx,
	}
}

// NewDetachedNodeActionParamsWithHTTPClient creates a new DetachedNodeActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetachedNodeActionParamsWithHTTPClient(client *http.Client) *DetachedNodeActionParams {
	return &DetachedNodeActionParams{
		HTTPClient: client,
	}
}

/* DetachedNodeActionParams contains all the parameters to send to the API endpoint
   for the detached node action operation.

   Typically these are written to a http.Request.
*/
type DetachedNodeActionParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// InstanceIP.
	InstanceIP string

	// PUUID.
	//
	// Format: uuid
	PUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detached node action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachedNodeActionParams) WithDefaults() *DetachedNodeActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detached node action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachedNodeActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detached node action params
func (o *DetachedNodeActionParams) WithTimeout(timeout time.Duration) *DetachedNodeActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detached node action params
func (o *DetachedNodeActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detached node action params
func (o *DetachedNodeActionParams) WithContext(ctx context.Context) *DetachedNodeActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detached node action params
func (o *DetachedNodeActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detached node action params
func (o *DetachedNodeActionParams) WithHTTPClient(client *http.Client) *DetachedNodeActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detached node action params
func (o *DetachedNodeActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the detached node action params
func (o *DetachedNodeActionParams) WithCUUID(cUUID strfmt.UUID) *DetachedNodeActionParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the detached node action params
func (o *DetachedNodeActionParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithInstanceIP adds the instanceIP to the detached node action params
func (o *DetachedNodeActionParams) WithInstanceIP(instanceIP string) *DetachedNodeActionParams {
	o.SetInstanceIP(instanceIP)
	return o
}

// SetInstanceIP adds the instanceIp to the detached node action params
func (o *DetachedNodeActionParams) SetInstanceIP(instanceIP string) {
	o.InstanceIP = instanceIP
}

// WithPUUID adds the pUUID to the detached node action params
func (o *DetachedNodeActionParams) WithPUUID(pUUID strfmt.UUID) *DetachedNodeActionParams {
	o.SetPUUID(pUUID)
	return o
}

// SetPUUID adds the pUuid to the detached node action params
func (o *DetachedNodeActionParams) SetPUUID(pUUID strfmt.UUID) {
	o.PUUID = pUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachedNodeActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param instanceIP
	if err := r.SetPathParam("instanceIP", o.InstanceIP); err != nil {
		return err
	}

	// path param pUUID
	if err := r.SetPathParam("pUUID", o.PUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}