// Code generated by go-swagger; DO NOT EDIT.

package file_upload

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

// NewPublicGeneratedUploadURLParams creates a new PublicGeneratedUploadURLParams object
// with the default values initialized.
func NewPublicGeneratedUploadURLParams() *PublicGeneratedUploadURLParams {
	var ()
	return &PublicGeneratedUploadURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGeneratedUploadURLParamsWithTimeout creates a new PublicGeneratedUploadURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGeneratedUploadURLParamsWithTimeout(timeout time.Duration) *PublicGeneratedUploadURLParams {
	var ()
	return &PublicGeneratedUploadURLParams{

		timeout: timeout,
	}
}

// NewPublicGeneratedUploadURLParamsWithContext creates a new PublicGeneratedUploadURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGeneratedUploadURLParamsWithContext(ctx context.Context) *PublicGeneratedUploadURLParams {
	var ()
	return &PublicGeneratedUploadURLParams{

		Context: ctx,
	}
}

// NewPublicGeneratedUploadURLParamsWithHTTPClient creates a new PublicGeneratedUploadURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGeneratedUploadURLParamsWithHTTPClient(client *http.Client) *PublicGeneratedUploadURLParams {
	var ()
	return &PublicGeneratedUploadURLParams{
		HTTPClient: client,
	}
}

/*PublicGeneratedUploadURLParams contains all the parameters to send to the API endpoint
for the public generated upload Url operation typically these are written to a http.Request
*/
type PublicGeneratedUploadURLParams struct {

	/*FileType
	  one of the these types: jpeg, jpg, png, bmp, gif, mp3, bin

	*/
	FileType string
	/*Folder
	  the name of folder where the file will be uploaded, must be between 1-256 characters, all characters allowed no whitespace

	*/
	Folder string
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithTimeout(timeout time.Duration) *PublicGeneratedUploadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithContext(ctx context.Context) *PublicGeneratedUploadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithHTTPClient(client *http.Client) *PublicGeneratedUploadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileType adds the fileType to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithFileType(fileType string) *PublicGeneratedUploadURLParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetFileType(fileType string) {
	o.FileType = fileType
}

// WithFolder adds the folder to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithFolder(folder string) *PublicGeneratedUploadURLParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetFolder(folder string) {
	o.Folder = folder
}

// WithNamespace adds the namespace to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) WithNamespace(namespace string) *PublicGeneratedUploadURLParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public generated upload Url params
func (o *PublicGeneratedUploadURLParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGeneratedUploadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param fileType
	qrFileType := o.FileType
	qFileType := qrFileType
	if qFileType != "" {
		if err := r.SetQueryParam("fileType", qFileType); err != nil {
			return err
		}
	}

	// path param folder
	if err := r.SetPathParam("folder", o.Folder); err != nil {
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
