// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/cisco-open/kubei/api/client/models"
)

// NewPutRuntimeScheduleScanStartParams creates a new PutRuntimeScheduleScanStartParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRuntimeScheduleScanStartParams() *PutRuntimeScheduleScanStartParams {
	return &PutRuntimeScheduleScanStartParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRuntimeScheduleScanStartParamsWithTimeout creates a new PutRuntimeScheduleScanStartParams object
// with the ability to set a timeout on a request.
func NewPutRuntimeScheduleScanStartParamsWithTimeout(timeout time.Duration) *PutRuntimeScheduleScanStartParams {
	return &PutRuntimeScheduleScanStartParams{
		timeout: timeout,
	}
}

// NewPutRuntimeScheduleScanStartParamsWithContext creates a new PutRuntimeScheduleScanStartParams object
// with the ability to set a context for a request.
func NewPutRuntimeScheduleScanStartParamsWithContext(ctx context.Context) *PutRuntimeScheduleScanStartParams {
	return &PutRuntimeScheduleScanStartParams{
		Context: ctx,
	}
}

// NewPutRuntimeScheduleScanStartParamsWithHTTPClient creates a new PutRuntimeScheduleScanStartParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRuntimeScheduleScanStartParamsWithHTTPClient(client *http.Client) *PutRuntimeScheduleScanStartParams {
	return &PutRuntimeScheduleScanStartParams{
		HTTPClient: client,
	}
}

/* PutRuntimeScheduleScanStartParams contains all the parameters to send to the API endpoint
   for the put runtime schedule scan start operation.

   Typically these are written to a http.Request.
*/
type PutRuntimeScheduleScanStartParams struct {

	// Body.
	Body *models.RuntimeScheduleScanConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put runtime schedule scan start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRuntimeScheduleScanStartParams) WithDefaults() *PutRuntimeScheduleScanStartParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put runtime schedule scan start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRuntimeScheduleScanStartParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) WithTimeout(timeout time.Duration) *PutRuntimeScheduleScanStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) WithContext(ctx context.Context) *PutRuntimeScheduleScanStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) WithHTTPClient(client *http.Client) *PutRuntimeScheduleScanStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) WithBody(body *models.RuntimeScheduleScanConfig) *PutRuntimeScheduleScanStartParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put runtime schedule scan start params
func (o *PutRuntimeScheduleScanStartParams) SetBody(body *models.RuntimeScheduleScanConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutRuntimeScheduleScanStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
