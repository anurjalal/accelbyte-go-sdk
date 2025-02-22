// Code generated by go-swagger; DO NOT EDIT.

package payment_config

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewTestWxPayConfigParams creates a new TestWxPayConfigParams object
// with the default values initialized.
func NewTestWxPayConfigParams() *TestWxPayConfigParams {
	var ()
	return &TestWxPayConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestWxPayConfigParamsWithTimeout creates a new TestWxPayConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestWxPayConfigParamsWithTimeout(timeout time.Duration) *TestWxPayConfigParams {
	var ()
	return &TestWxPayConfigParams{

		timeout: timeout,
	}
}

// NewTestWxPayConfigParamsWithContext creates a new TestWxPayConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestWxPayConfigParamsWithContext(ctx context.Context) *TestWxPayConfigParams {
	var ()
	return &TestWxPayConfigParams{

		Context: ctx,
	}
}

// NewTestWxPayConfigParamsWithHTTPClient creates a new TestWxPayConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestWxPayConfigParamsWithHTTPClient(client *http.Client) *TestWxPayConfigParams {
	var ()
	return &TestWxPayConfigParams{
		HTTPClient: client,
	}
}

/*TestWxPayConfigParams contains all the parameters to send to the API endpoint
for the test wx pay config operation typically these are written to a http.Request
*/
type TestWxPayConfigParams struct {

	/*Body*/
	Body *platformclientmodels.WxPayConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test wx pay config params
func (o *TestWxPayConfigParams) WithTimeout(timeout time.Duration) *TestWxPayConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test wx pay config params
func (o *TestWxPayConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test wx pay config params
func (o *TestWxPayConfigParams) WithContext(ctx context.Context) *TestWxPayConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test wx pay config params
func (o *TestWxPayConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test wx pay config params
func (o *TestWxPayConfigParams) WithHTTPClient(client *http.Client) *TestWxPayConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test wx pay config params
func (o *TestWxPayConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test wx pay config params
func (o *TestWxPayConfigParams) WithBody(body *platformclientmodels.WxPayConfigRequest) *TestWxPayConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test wx pay config params
func (o *TestWxPayConfigParams) SetBody(body *platformclientmodels.WxPayConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestWxPayConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
