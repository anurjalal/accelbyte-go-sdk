// Code generated by go-swagger; DO NOT EDIT.

package o_auth

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

// NewPlatformTokenGrantParams creates a new PlatformTokenGrantParams object
// with the default values initialized.
func NewPlatformTokenGrantParams() *PlatformTokenGrantParams {
	var ()
	return &PlatformTokenGrantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlatformTokenGrantParamsWithTimeout creates a new PlatformTokenGrantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlatformTokenGrantParamsWithTimeout(timeout time.Duration) *PlatformTokenGrantParams {
	var ()
	return &PlatformTokenGrantParams{

		timeout: timeout,
	}
}

// NewPlatformTokenGrantParamsWithContext creates a new PlatformTokenGrantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlatformTokenGrantParamsWithContext(ctx context.Context) *PlatformTokenGrantParams {
	var ()
	return &PlatformTokenGrantParams{

		Context: ctx,
	}
}

// NewPlatformTokenGrantParamsWithHTTPClient creates a new PlatformTokenGrantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlatformTokenGrantParamsWithHTTPClient(client *http.Client) *PlatformTokenGrantParams {
	var ()
	return &PlatformTokenGrantParams{
		HTTPClient: client,
	}
}

/*PlatformTokenGrantParams contains all the parameters to send to the API endpoint
for the platform token grant operation typically these are written to a http.Request
*/
type PlatformTokenGrantParams struct {
	DeviceID *string
	/*DeviceID
	  DeviceID (Used on grant type 'password' to track login history) ex. 90252d14544846d79f367148e3f9a3d9

	*/
	Device_id *string
	/*Device_id
	  Device/hardware identifier

	*/
	Namespace *string
	/*Namespace
	  Delegated namespace token grant

	*/
	PlatformID string
	/*PlatformID
	  Platform ID to login with

	*/
	PlatformToken *string
	/*PlatformToken
	  Token from platform auth

	*/

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the platform token grant params
func (o *PlatformTokenGrantParams) WithTimeout(timeout time.Duration) *PlatformTokenGrantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the platform token grant params
func (o *PlatformTokenGrantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the platform token grant params
func (o *PlatformTokenGrantParams) WithContext(ctx context.Context) *PlatformTokenGrantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the platform token grant params
func (o *PlatformTokenGrantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the platform token grant params
func (o *PlatformTokenGrantParams) WithHTTPClient(client *http.Client) *PlatformTokenGrantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the platform token grant params
func (o *PlatformTokenGrantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the platform token grant params
func (o *PlatformTokenGrantParams) WithDeviceID(deviceID *string) *PlatformTokenGrantParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the platform token grant params
func (o *PlatformTokenGrantParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithNamespace adds the namespace to the platform token grant params
func (o *PlatformTokenGrantParams) WithNamespace(namespace *string) *PlatformTokenGrantParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the platform token grant params
func (o *PlatformTokenGrantParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithPlatformID adds the platformID to the platform token grant params
func (o *PlatformTokenGrantParams) WithPlatformID(platformID string) *PlatformTokenGrantParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the platform token grant params
func (o *PlatformTokenGrantParams) SetPlatformID(platformID string) {
	o.PlatformID = platformID
}

// WithPlatformToken adds the platformToken to the platform token grant params
func (o *PlatformTokenGrantParams) WithPlatformToken(platformToken *string) *PlatformTokenGrantParams {
	o.SetPlatformToken(platformToken)
	return o
}

// SetPlatformToken adds the platformToken to the platform token grant params
func (o *PlatformTokenGrantParams) SetPlatformToken(platformToken *string) {
	o.PlatformToken = platformToken
}

// WriteToRequest writes these params to a swagger request
func (o *PlatformTokenGrantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeviceID != nil {

		// header param Device-Id
		if err := r.SetHeaderParam("Device-Id", *o.DeviceID); err != nil {
			return err
		}

	}

	if o.DeviceID != nil {

		// form param device_id
		var frDeviceID string
		if o.DeviceID != nil {
			frDeviceID = *o.DeviceID
		}
		fDeviceID := frDeviceID
		if fDeviceID != "" {
			if err := r.SetFormParam("device_id", fDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Namespace != nil {

		// form param namespace
		var frNamespace string
		if o.Namespace != nil {
			frNamespace = *o.Namespace
		}
		fNamespace := frNamespace
		if fNamespace != "" {
			if err := r.SetFormParam("namespace", fNamespace); err != nil {
				return err
			}
		}

	}

	// path param platformId
	if err := r.SetPathParam("platformId", o.PlatformID); err != nil {
		return err
	}

	if o.PlatformToken != nil {

		// form param platform_token
		var frPlatformToken string
		if o.PlatformToken != nil {
			frPlatformToken = *o.PlatformToken
		}
		fPlatformToken := frPlatformToken
		if fPlatformToken != "" {
			if err := r.SetFormParam("platform_token", fPlatformToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
