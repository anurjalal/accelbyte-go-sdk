// Code generated by go-swagger; DO NOT EDIT.

package payment_dedicated

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

// NewRefundPaymentOrderByDedicatedParams creates a new RefundPaymentOrderByDedicatedParams object
// with the default values initialized.
func NewRefundPaymentOrderByDedicatedParams() *RefundPaymentOrderByDedicatedParams {
	var ()
	return &RefundPaymentOrderByDedicatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefundPaymentOrderByDedicatedParamsWithTimeout creates a new RefundPaymentOrderByDedicatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefundPaymentOrderByDedicatedParamsWithTimeout(timeout time.Duration) *RefundPaymentOrderByDedicatedParams {
	var ()
	return &RefundPaymentOrderByDedicatedParams{

		timeout: timeout,
	}
}

// NewRefundPaymentOrderByDedicatedParamsWithContext creates a new RefundPaymentOrderByDedicatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefundPaymentOrderByDedicatedParamsWithContext(ctx context.Context) *RefundPaymentOrderByDedicatedParams {
	var ()
	return &RefundPaymentOrderByDedicatedParams{

		Context: ctx,
	}
}

// NewRefundPaymentOrderByDedicatedParamsWithHTTPClient creates a new RefundPaymentOrderByDedicatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefundPaymentOrderByDedicatedParamsWithHTTPClient(client *http.Client) *RefundPaymentOrderByDedicatedParams {
	var ()
	return &RefundPaymentOrderByDedicatedParams{
		HTTPClient: client,
	}
}

/*RefundPaymentOrderByDedicatedParams contains all the parameters to send to the API endpoint
for the refund payment order by dedicated operation typically these are written to a http.Request
*/
type RefundPaymentOrderByDedicatedParams struct {

	/*Body*/
	Body *platformclientmodels.PaymentOrderRefund
	/*Namespace*/
	Namespace string
	/*PaymentOrderNo*/
	PaymentOrderNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithTimeout(timeout time.Duration) *RefundPaymentOrderByDedicatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithContext(ctx context.Context) *RefundPaymentOrderByDedicatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithHTTPClient(client *http.Client) *RefundPaymentOrderByDedicatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithBody(body *platformclientmodels.PaymentOrderRefund) *RefundPaymentOrderByDedicatedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetBody(body *platformclientmodels.PaymentOrderRefund) {
	o.Body = body
}

// WithNamespace adds the namespace to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithNamespace(namespace string) *RefundPaymentOrderByDedicatedParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPaymentOrderNo adds the paymentOrderNo to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) WithPaymentOrderNo(paymentOrderNo string) *RefundPaymentOrderByDedicatedParams {
	o.SetPaymentOrderNo(paymentOrderNo)
	return o
}

// SetPaymentOrderNo adds the paymentOrderNo to the refund payment order by dedicated params
func (o *RefundPaymentOrderByDedicatedParams) SetPaymentOrderNo(paymentOrderNo string) {
	o.PaymentOrderNo = paymentOrderNo
}

// WriteToRequest writes these params to a swagger request
func (o *RefundPaymentOrderByDedicatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentOrderNo
	if err := r.SetPathParam("paymentOrderNo", o.PaymentOrderNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
