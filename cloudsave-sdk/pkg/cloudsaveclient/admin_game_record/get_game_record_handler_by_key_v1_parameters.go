// Code generated by go-swagger; DO NOT EDIT.

package admin_game_record

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

// NewGetGameRecordHandlerByKeyV1Params creates a new GetGameRecordHandlerByKeyV1Params object
// with the default values initialized.
func NewGetGameRecordHandlerByKeyV1Params() *GetGameRecordHandlerByKeyV1Params {
	var ()
	return &GetGameRecordHandlerByKeyV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGameRecordHandlerByKeyV1ParamsWithTimeout creates a new GetGameRecordHandlerByKeyV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGameRecordHandlerByKeyV1ParamsWithTimeout(timeout time.Duration) *GetGameRecordHandlerByKeyV1Params {
	var ()
	return &GetGameRecordHandlerByKeyV1Params{

		timeout: timeout,
	}
}

// NewGetGameRecordHandlerByKeyV1ParamsWithContext creates a new GetGameRecordHandlerByKeyV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetGameRecordHandlerByKeyV1ParamsWithContext(ctx context.Context) *GetGameRecordHandlerByKeyV1Params {
	var ()
	return &GetGameRecordHandlerByKeyV1Params{

		Context: ctx,
	}
}

// NewGetGameRecordHandlerByKeyV1ParamsWithHTTPClient creates a new GetGameRecordHandlerByKeyV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGameRecordHandlerByKeyV1ParamsWithHTTPClient(client *http.Client) *GetGameRecordHandlerByKeyV1Params {
	var ()
	return &GetGameRecordHandlerByKeyV1Params{
		HTTPClient: client,
	}
}

/*GetGameRecordHandlerByKeyV1Params contains all the parameters to send to the API endpoint
for the get game record handler by key v1 operation typically these are written to a http.Request
*/
type GetGameRecordHandlerByKeyV1Params struct {

	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) WithTimeout(timeout time.Duration) *GetGameRecordHandlerByKeyV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) WithContext(ctx context.Context) *GetGameRecordHandlerByKeyV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) WithHTTPClient(client *http.Client) *GetGameRecordHandlerByKeyV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) WithKey(key string) *GetGameRecordHandlerByKeyV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) WithNamespace(namespace string) *GetGameRecordHandlerByKeyV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get game record handler by key v1 params
func (o *GetGameRecordHandlerByKeyV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetGameRecordHandlerByKeyV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
