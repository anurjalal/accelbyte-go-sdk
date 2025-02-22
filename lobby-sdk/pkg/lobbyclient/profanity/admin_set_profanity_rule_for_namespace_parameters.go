// Code generated by go-swagger; DO NOT EDIT.

package profanity

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewAdminSetProfanityRuleForNamespaceParams creates a new AdminSetProfanityRuleForNamespaceParams object
// with the default values initialized.
func NewAdminSetProfanityRuleForNamespaceParams() *AdminSetProfanityRuleForNamespaceParams {
	var ()
	return &AdminSetProfanityRuleForNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminSetProfanityRuleForNamespaceParamsWithTimeout creates a new AdminSetProfanityRuleForNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminSetProfanityRuleForNamespaceParamsWithTimeout(timeout time.Duration) *AdminSetProfanityRuleForNamespaceParams {
	var ()
	return &AdminSetProfanityRuleForNamespaceParams{

		timeout: timeout,
	}
}

// NewAdminSetProfanityRuleForNamespaceParamsWithContext creates a new AdminSetProfanityRuleForNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminSetProfanityRuleForNamespaceParamsWithContext(ctx context.Context) *AdminSetProfanityRuleForNamespaceParams {
	var ()
	return &AdminSetProfanityRuleForNamespaceParams{

		Context: ctx,
	}
}

// NewAdminSetProfanityRuleForNamespaceParamsWithHTTPClient creates a new AdminSetProfanityRuleForNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminSetProfanityRuleForNamespaceParamsWithHTTPClient(client *http.Client) *AdminSetProfanityRuleForNamespaceParams {
	var ()
	return &AdminSetProfanityRuleForNamespaceParams{
		HTTPClient: client,
	}
}

/*AdminSetProfanityRuleForNamespaceParams contains all the parameters to send to the API endpoint
for the admin set profanity rule for namespace operation typically these are written to a http.Request
*/
type AdminSetProfanityRuleForNamespaceParams struct {

	/*Body
	  request

	*/
	Body *lobbyclientmodels.ModelsAdminSetProfanityRuleForNamespaceRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) WithTimeout(timeout time.Duration) *AdminSetProfanityRuleForNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) WithContext(ctx context.Context) *AdminSetProfanityRuleForNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) WithHTTPClient(client *http.Client) *AdminSetProfanityRuleForNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) WithBody(body *lobbyclientmodels.ModelsAdminSetProfanityRuleForNamespaceRequest) *AdminSetProfanityRuleForNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) SetBody(body *lobbyclientmodels.ModelsAdminSetProfanityRuleForNamespaceRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) WithNamespace(namespace string) *AdminSetProfanityRuleForNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin set profanity rule for namespace params
func (o *AdminSetProfanityRuleForNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminSetProfanityRuleForNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
