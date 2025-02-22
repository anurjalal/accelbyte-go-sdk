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

// NewAdminDeleteProfanityFilterParams creates a new AdminDeleteProfanityFilterParams object
// with the default values initialized.
func NewAdminDeleteProfanityFilterParams() *AdminDeleteProfanityFilterParams {
	var ()
	return &AdminDeleteProfanityFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteProfanityFilterParamsWithTimeout creates a new AdminDeleteProfanityFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeleteProfanityFilterParamsWithTimeout(timeout time.Duration) *AdminDeleteProfanityFilterParams {
	var ()
	return &AdminDeleteProfanityFilterParams{

		timeout: timeout,
	}
}

// NewAdminDeleteProfanityFilterParamsWithContext creates a new AdminDeleteProfanityFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeleteProfanityFilterParamsWithContext(ctx context.Context) *AdminDeleteProfanityFilterParams {
	var ()
	return &AdminDeleteProfanityFilterParams{

		Context: ctx,
	}
}

// NewAdminDeleteProfanityFilterParamsWithHTTPClient creates a new AdminDeleteProfanityFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeleteProfanityFilterParamsWithHTTPClient(client *http.Client) *AdminDeleteProfanityFilterParams {
	var ()
	return &AdminDeleteProfanityFilterParams{
		HTTPClient: client,
	}
}

/*AdminDeleteProfanityFilterParams contains all the parameters to send to the API endpoint
for the admin delete profanity filter operation typically these are written to a http.Request
*/
type AdminDeleteProfanityFilterParams struct {

	/*Body
	  request

	*/
	Body *lobbyclientmodels.ModelsAdminDeleteProfanityFilterRequest
	/*List
	  list

	*/
	List string
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithTimeout(timeout time.Duration) *AdminDeleteProfanityFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithContext(ctx context.Context) *AdminDeleteProfanityFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithHTTPClient(client *http.Client) *AdminDeleteProfanityFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithBody(body *lobbyclientmodels.ModelsAdminDeleteProfanityFilterRequest) *AdminDeleteProfanityFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetBody(body *lobbyclientmodels.ModelsAdminDeleteProfanityFilterRequest) {
	o.Body = body
}

// WithList adds the list to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithList(list string) *AdminDeleteProfanityFilterParams {
	o.SetList(list)
	return o
}

// SetList adds the list to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetList(list string) {
	o.List = list
}

// WithNamespace adds the namespace to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) WithNamespace(namespace string) *AdminDeleteProfanityFilterParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin delete profanity filter params
func (o *AdminDeleteProfanityFilterParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteProfanityFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param list
	if err := r.SetPathParam("list", o.List); err != nil {
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
