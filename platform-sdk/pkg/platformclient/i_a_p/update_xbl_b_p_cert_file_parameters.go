// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// NewUpdateXblBPCertFileParams creates a new UpdateXblBPCertFileParams object
// with the default values initialized.
func NewUpdateXblBPCertFileParams() *UpdateXblBPCertFileParams {
	var ()
	return &UpdateXblBPCertFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateXblBPCertFileParamsWithTimeout creates a new UpdateXblBPCertFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateXblBPCertFileParamsWithTimeout(timeout time.Duration) *UpdateXblBPCertFileParams {
	var ()
	return &UpdateXblBPCertFileParams{

		timeout: timeout,
	}
}

// NewUpdateXblBPCertFileParamsWithContext creates a new UpdateXblBPCertFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateXblBPCertFileParamsWithContext(ctx context.Context) *UpdateXblBPCertFileParams {
	var ()
	return &UpdateXblBPCertFileParams{

		Context: ctx,
	}
}

// NewUpdateXblBPCertFileParamsWithHTTPClient creates a new UpdateXblBPCertFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateXblBPCertFileParamsWithHTTPClient(client *http.Client) *UpdateXblBPCertFileParams {
	var ()
	return &UpdateXblBPCertFileParams{
		HTTPClient: client,
	}
}

/*UpdateXblBPCertFileParams contains all the parameters to send to the API endpoint
for the update xbl b p cert file operation typically these are written to a http.Request
*/
type UpdateXblBPCertFileParams struct {

	/*File*/
	File runtime.NamedReadCloser
	/*Namespace*/
	Namespace string
	/*Password*/
	Password *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithTimeout(timeout time.Duration) *UpdateXblBPCertFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithContext(ctx context.Context) *UpdateXblBPCertFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithHTTPClient(client *http.Client) *UpdateXblBPCertFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithFile(file runtime.NamedReadCloser) *UpdateXblBPCertFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithNamespace adds the namespace to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithNamespace(namespace string) *UpdateXblBPCertFileParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPassword adds the password to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) WithPassword(password *string) *UpdateXblBPCertFileParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the update xbl b p cert file params
func (o *UpdateXblBPCertFileParams) SetPassword(password *string) {
	o.Password = password
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateXblBPCertFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {

			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}

		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
