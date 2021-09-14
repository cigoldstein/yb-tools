// Code generated by go-swagger; DO NOT EDIT.

package backup_schedule_management

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

// NewListBackupSchedulesParams creates a new ListBackupSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBackupSchedulesParams() *ListBackupSchedulesParams {
	return &ListBackupSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBackupSchedulesParamsWithTimeout creates a new ListBackupSchedulesParams object
// with the ability to set a timeout on a request.
func NewListBackupSchedulesParamsWithTimeout(timeout time.Duration) *ListBackupSchedulesParams {
	return &ListBackupSchedulesParams{
		timeout: timeout,
	}
}

// NewListBackupSchedulesParamsWithContext creates a new ListBackupSchedulesParams object
// with the ability to set a context for a request.
func NewListBackupSchedulesParamsWithContext(ctx context.Context) *ListBackupSchedulesParams {
	return &ListBackupSchedulesParams{
		Context: ctx,
	}
}

// NewListBackupSchedulesParamsWithHTTPClient creates a new ListBackupSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBackupSchedulesParamsWithHTTPClient(client *http.Client) *ListBackupSchedulesParams {
	return &ListBackupSchedulesParams{
		HTTPClient: client,
	}
}

/* ListBackupSchedulesParams contains all the parameters to send to the API endpoint
   for the list backup schedules operation.

   Typically these are written to a http.Request.
*/
type ListBackupSchedulesParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list backup schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackupSchedulesParams) WithDefaults() *ListBackupSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list backup schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackupSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list backup schedules params
func (o *ListBackupSchedulesParams) WithTimeout(timeout time.Duration) *ListBackupSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list backup schedules params
func (o *ListBackupSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list backup schedules params
func (o *ListBackupSchedulesParams) WithContext(ctx context.Context) *ListBackupSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list backup schedules params
func (o *ListBackupSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list backup schedules params
func (o *ListBackupSchedulesParams) WithHTTPClient(client *http.Client) *ListBackupSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list backup schedules params
func (o *ListBackupSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the list backup schedules params
func (o *ListBackupSchedulesParams) WithCUUID(cUUID strfmt.UUID) *ListBackupSchedulesParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the list backup schedules params
func (o *ListBackupSchedulesParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBackupSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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