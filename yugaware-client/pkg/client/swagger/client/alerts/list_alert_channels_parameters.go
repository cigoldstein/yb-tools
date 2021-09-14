// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewListAlertChannelsParams creates a new ListAlertChannelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAlertChannelsParams() *ListAlertChannelsParams {
	return &ListAlertChannelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAlertChannelsParamsWithTimeout creates a new ListAlertChannelsParams object
// with the ability to set a timeout on a request.
func NewListAlertChannelsParamsWithTimeout(timeout time.Duration) *ListAlertChannelsParams {
	return &ListAlertChannelsParams{
		timeout: timeout,
	}
}

// NewListAlertChannelsParamsWithContext creates a new ListAlertChannelsParams object
// with the ability to set a context for a request.
func NewListAlertChannelsParamsWithContext(ctx context.Context) *ListAlertChannelsParams {
	return &ListAlertChannelsParams{
		Context: ctx,
	}
}

// NewListAlertChannelsParamsWithHTTPClient creates a new ListAlertChannelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAlertChannelsParamsWithHTTPClient(client *http.Client) *ListAlertChannelsParams {
	return &ListAlertChannelsParams{
		HTTPClient: client,
	}
}

/* ListAlertChannelsParams contains all the parameters to send to the API endpoint
   for the list alert channels operation.

   Typically these are written to a http.Request.
*/
type ListAlertChannelsParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list alert channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlertChannelsParams) WithDefaults() *ListAlertChannelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list alert channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlertChannelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list alert channels params
func (o *ListAlertChannelsParams) WithTimeout(timeout time.Duration) *ListAlertChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alert channels params
func (o *ListAlertChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alert channels params
func (o *ListAlertChannelsParams) WithContext(ctx context.Context) *ListAlertChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alert channels params
func (o *ListAlertChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alert channels params
func (o *ListAlertChannelsParams) WithHTTPClient(client *http.Client) *ListAlertChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alert channels params
func (o *ListAlertChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the list alert channels params
func (o *ListAlertChannelsParams) WithCUUID(cUUID strfmt.UUID) *ListAlertChannelsParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the list alert channels params
func (o *ListAlertChannelsParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlertChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
