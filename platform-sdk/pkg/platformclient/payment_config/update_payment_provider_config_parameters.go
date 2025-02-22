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

// NewUpdatePaymentProviderConfigParams creates a new UpdatePaymentProviderConfigParams object
// with the default values initialized.
func NewUpdatePaymentProviderConfigParams() *UpdatePaymentProviderConfigParams {
	var ()
	return &UpdatePaymentProviderConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePaymentProviderConfigParamsWithTimeout creates a new UpdatePaymentProviderConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePaymentProviderConfigParamsWithTimeout(timeout time.Duration) *UpdatePaymentProviderConfigParams {
	var ()
	return &UpdatePaymentProviderConfigParams{

		timeout: timeout,
	}
}

// NewUpdatePaymentProviderConfigParamsWithContext creates a new UpdatePaymentProviderConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePaymentProviderConfigParamsWithContext(ctx context.Context) *UpdatePaymentProviderConfigParams {
	var ()
	return &UpdatePaymentProviderConfigParams{

		Context: ctx,
	}
}

// NewUpdatePaymentProviderConfigParamsWithHTTPClient creates a new UpdatePaymentProviderConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePaymentProviderConfigParamsWithHTTPClient(client *http.Client) *UpdatePaymentProviderConfigParams {
	var ()
	return &UpdatePaymentProviderConfigParams{
		HTTPClient: client,
	}
}

/*UpdatePaymentProviderConfigParams contains all the parameters to send to the API endpoint
for the update payment provider config operation typically these are written to a http.Request
*/
type UpdatePaymentProviderConfigParams struct {

	/*Body*/
	Body *platformclientmodels.PaymentProviderConfigEdit
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) WithTimeout(timeout time.Duration) *UpdatePaymentProviderConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) WithContext(ctx context.Context) *UpdatePaymentProviderConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) WithHTTPClient(client *http.Client) *UpdatePaymentProviderConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) WithBody(body *platformclientmodels.PaymentProviderConfigEdit) *UpdatePaymentProviderConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) SetBody(body *platformclientmodels.PaymentProviderConfigEdit) {
	o.Body = body
}

// WithID adds the id to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) WithID(id string) *UpdatePaymentProviderConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update payment provider config params
func (o *UpdatePaymentProviderConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePaymentProviderConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
