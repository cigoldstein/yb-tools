// Code generated by go-swagger; DO NOT EDIT.

package customer_management

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

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// NewMetricsParams creates a new MetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetricsParams() *MetricsParams {
	return &MetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetricsParamsWithTimeout creates a new MetricsParams object
// with the ability to set a timeout on a request.
func NewMetricsParamsWithTimeout(timeout time.Duration) *MetricsParams {
	return &MetricsParams{
		timeout: timeout,
	}
}

// NewMetricsParamsWithContext creates a new MetricsParams object
// with the ability to set a context for a request.
func NewMetricsParamsWithContext(ctx context.Context) *MetricsParams {
	return &MetricsParams{
		Context: ctx,
	}
}

// NewMetricsParamsWithHTTPClient creates a new MetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetricsParamsWithHTTPClient(client *http.Client) *MetricsParams {
	return &MetricsParams{
		HTTPClient: client,
	}
}

/* MetricsParams contains all the parameters to send to the API endpoint
   for the metrics operation.

   Typically these are written to a http.Request.
*/
type MetricsParams struct {

	/* Metrics.

	   Metrics to be added
	*/
	Metrics *models.MetricQueryParams

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetricsParams) WithDefaults() *MetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the metrics params
func (o *MetricsParams) WithTimeout(timeout time.Duration) *MetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrics params
func (o *MetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrics params
func (o *MetricsParams) WithContext(ctx context.Context) *MetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrics params
func (o *MetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrics params
func (o *MetricsParams) WithHTTPClient(client *http.Client) *MetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrics params
func (o *MetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetrics adds the metrics to the metrics params
func (o *MetricsParams) WithMetrics(metrics *models.MetricQueryParams) *MetricsParams {
	o.SetMetrics(metrics)
	return o
}

// SetMetrics adds the metrics to the metrics params
func (o *MetricsParams) SetMetrics(metrics *models.MetricQueryParams) {
	o.Metrics = metrics
}

// WithCUUID adds the cUUID to the metrics params
func (o *MetricsParams) WithCUUID(cUUID strfmt.UUID) *MetricsParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the metrics params
func (o *MetricsParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Metrics != nil {
		if err := r.SetBodyParam(o.Metrics); err != nil {
			return err
		}
	}

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}