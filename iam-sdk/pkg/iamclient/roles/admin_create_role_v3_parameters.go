// Code generated by go-swagger; DO NOT EDIT.

package roles

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

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewAdminCreateRoleV3Params creates a new AdminCreateRoleV3Params object
// with the default values initialized.
func NewAdminCreateRoleV3Params() *AdminCreateRoleV3Params {
	var ()
	return &AdminCreateRoleV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateRoleV3ParamsWithTimeout creates a new AdminCreateRoleV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminCreateRoleV3ParamsWithTimeout(timeout time.Duration) *AdminCreateRoleV3Params {
	var ()
	return &AdminCreateRoleV3Params{

		timeout: timeout,
	}
}

// NewAdminCreateRoleV3ParamsWithContext creates a new AdminCreateRoleV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminCreateRoleV3ParamsWithContext(ctx context.Context) *AdminCreateRoleV3Params {
	var ()
	return &AdminCreateRoleV3Params{

		Context: ctx,
	}
}

// NewAdminCreateRoleV3ParamsWithHTTPClient creates a new AdminCreateRoleV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminCreateRoleV3ParamsWithHTTPClient(client *http.Client) *AdminCreateRoleV3Params {
	var ()
	return &AdminCreateRoleV3Params{
		HTTPClient: client,
	}
}

/*AdminCreateRoleV3Params contains all the parameters to send to the API endpoint
for the admin create role v3 operation typically these are written to a http.Request
*/
type AdminCreateRoleV3Params struct {

	/*Body*/
	Body *iamclientmodels.ModelRoleCreateV3Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin create role v3 params
func (o *AdminCreateRoleV3Params) WithTimeout(timeout time.Duration) *AdminCreateRoleV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create role v3 params
func (o *AdminCreateRoleV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create role v3 params
func (o *AdminCreateRoleV3Params) WithContext(ctx context.Context) *AdminCreateRoleV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create role v3 params
func (o *AdminCreateRoleV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create role v3 params
func (o *AdminCreateRoleV3Params) WithHTTPClient(client *http.Client) *AdminCreateRoleV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create role v3 params
func (o *AdminCreateRoleV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin create role v3 params
func (o *AdminCreateRoleV3Params) WithBody(body *iamclientmodels.ModelRoleCreateV3Request) *AdminCreateRoleV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin create role v3 params
func (o *AdminCreateRoleV3Params) SetBody(body *iamclientmodels.ModelRoleCreateV3Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateRoleV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
