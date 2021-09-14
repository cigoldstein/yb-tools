// Code generated by go-swagger; DO NOT EDIT.

package customer_tasks

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

// NewTasksListParams creates a new TasksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTasksListParams() *TasksListParams {
	return &TasksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTasksListParamsWithTimeout creates a new TasksListParams object
// with the ability to set a timeout on a request.
func NewTasksListParamsWithTimeout(timeout time.Duration) *TasksListParams {
	return &TasksListParams{
		timeout: timeout,
	}
}

// NewTasksListParamsWithContext creates a new TasksListParams object
// with the ability to set a context for a request.
func NewTasksListParamsWithContext(ctx context.Context) *TasksListParams {
	return &TasksListParams{
		Context: ctx,
	}
}

// NewTasksListParamsWithHTTPClient creates a new TasksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTasksListParamsWithHTTPClient(client *http.Client) *TasksListParams {
	return &TasksListParams{
		HTTPClient: client,
	}
}

/* TasksListParams contains all the parameters to send to the API endpoint
   for the tasks list operation.

   Typically these are written to a http.Request.
*/
type TasksListParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// UUUID.
	//
	// Format: uuid
	UUUID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tasks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TasksListParams) WithDefaults() *TasksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tasks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TasksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tasks list params
func (o *TasksListParams) WithTimeout(timeout time.Duration) *TasksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tasks list params
func (o *TasksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tasks list params
func (o *TasksListParams) WithContext(ctx context.Context) *TasksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tasks list params
func (o *TasksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tasks list params
func (o *TasksListParams) WithHTTPClient(client *http.Client) *TasksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tasks list params
func (o *TasksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the tasks list params
func (o *TasksListParams) WithCUUID(cUUID strfmt.UUID) *TasksListParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the tasks list params
func (o *TasksListParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUUUID adds the uUUID to the tasks list params
func (o *TasksListParams) WithUUUID(uUUID *strfmt.UUID) *TasksListParams {
	o.SetUUUID(uUUID)
	return o
}

// SetUUUID adds the uUuid to the tasks list params
func (o *TasksListParams) SetUUUID(uUUID *strfmt.UUID) {
	o.UUUID = uUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TasksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if o.UUUID != nil {

		// query param uUUID
		var qrUUUID strfmt.UUID

		if o.UUUID != nil {
			qrUUUID = *o.UUUID
		}
		qUUUID := qrUUUID.String()
		if qUUUID != "" {

			if err := r.SetQueryParam("uUUID", qUUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}