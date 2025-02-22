// Code generated by go-swagger; DO NOT EDIT.

package third_party_credential

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

// NewDeleteThirdPartyLoginPlatformCredentialV3Params creates a new DeleteThirdPartyLoginPlatformCredentialV3Params object
// with the default values initialized.
func NewDeleteThirdPartyLoginPlatformCredentialV3Params() *DeleteThirdPartyLoginPlatformCredentialV3Params {
	var ()
	return &DeleteThirdPartyLoginPlatformCredentialV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithTimeout creates a new DeleteThirdPartyLoginPlatformCredentialV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithTimeout(timeout time.Duration) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	var ()
	return &DeleteThirdPartyLoginPlatformCredentialV3Params{

		timeout: timeout,
	}
}

// NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithContext creates a new DeleteThirdPartyLoginPlatformCredentialV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithContext(ctx context.Context) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	var ()
	return &DeleteThirdPartyLoginPlatformCredentialV3Params{

		Context: ctx,
	}
}

// NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithHTTPClient creates a new DeleteThirdPartyLoginPlatformCredentialV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteThirdPartyLoginPlatformCredentialV3ParamsWithHTTPClient(client *http.Client) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	var ()
	return &DeleteThirdPartyLoginPlatformCredentialV3Params{
		HTTPClient: client,
	}
}

/*DeleteThirdPartyLoginPlatformCredentialV3Params contains all the parameters to send to the API endpoint
for the delete third party login platform credential v3 operation typically these are written to a http.Request
*/
type DeleteThirdPartyLoginPlatformCredentialV3Params struct {

	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*PlatformID
	  Platform ID

	*/
	PlatformID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WithTimeout(timeout time.Duration) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WithContext(ctx context.Context) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WithHTTPClient(client *http.Client) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WithNamespace(namespace string) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPlatformID adds the platformID to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WithPlatformID(platformID string) *DeleteThirdPartyLoginPlatformCredentialV3Params {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the delete third party login platform credential v3 params
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) SetPlatformID(platformID string) {
	o.PlatformID = platformID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteThirdPartyLoginPlatformCredentialV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param platformId
	if err := r.SetPathParam("platformId", o.PlatformID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
