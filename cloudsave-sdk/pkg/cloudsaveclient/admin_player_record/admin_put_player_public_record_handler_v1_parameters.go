// Code generated by go-swagger; DO NOT EDIT.

package admin_player_record

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

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// NewAdminPutPlayerPublicRecordHandlerV1Params creates a new AdminPutPlayerPublicRecordHandlerV1Params object
// with the default values initialized.
func NewAdminPutPlayerPublicRecordHandlerV1Params() *AdminPutPlayerPublicRecordHandlerV1Params {
	var ()
	return &AdminPutPlayerPublicRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminPutPlayerPublicRecordHandlerV1ParamsWithTimeout creates a new AdminPutPlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminPutPlayerPublicRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *AdminPutPlayerPublicRecordHandlerV1Params {
	var ()
	return &AdminPutPlayerPublicRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewAdminPutPlayerPublicRecordHandlerV1ParamsWithContext creates a new AdminPutPlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminPutPlayerPublicRecordHandlerV1ParamsWithContext(ctx context.Context) *AdminPutPlayerPublicRecordHandlerV1Params {
	var ()
	return &AdminPutPlayerPublicRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewAdminPutPlayerPublicRecordHandlerV1ParamsWithHTTPClient creates a new AdminPutPlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminPutPlayerPublicRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *AdminPutPlayerPublicRecordHandlerV1Params {
	var ()
	return &AdminPutPlayerPublicRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*AdminPutPlayerPublicRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the admin put player public record handler v1 operation typically these are written to a http.Request
*/
type AdminPutPlayerPublicRecordHandlerV1Params struct {

	/*Body*/
	Body cloudsaveclientmodels.ModelsPlayerRecordRequest
	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*UserID
	  user ID who own the record

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithTimeout(timeout time.Duration) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithContext(ctx context.Context) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithHTTPClient(client *http.Client) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithBody(body cloudsaveclientmodels.ModelsPlayerRecordRequest) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetBody(body cloudsaveclientmodels.ModelsPlayerRecordRequest) {
	o.Body = body
}

// WithKey adds the key to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithKey(key string) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithNamespace(namespace string) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WithUserID(userID string) *AdminPutPlayerPublicRecordHandlerV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin put player public record handler v1 params
func (o *AdminPutPlayerPublicRecordHandlerV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminPutPlayerPublicRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userID
	if err := r.SetPathParam("userID", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
