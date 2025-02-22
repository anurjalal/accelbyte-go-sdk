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

// NewGeneratedUploadURLParams creates a new GeneratedUploadURLParams object
// with the default values initialized.
func NewGeneratedUploadURLParams() *GeneratedUploadURLParams {
	var ()
	return &GeneratedUploadURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGeneratedUploadURLParamsWithTimeout creates a new GeneratedUploadURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeneratedUploadURLParamsWithTimeout(timeout time.Duration) *GeneratedUploadURLParams {
	var ()
	return &GeneratedUploadURLParams{

		timeout: timeout,
	}
}

// NewGeneratedUploadURLParamsWithContext creates a new GeneratedUploadURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGeneratedUploadURLParamsWithContext(ctx context.Context) *GeneratedUploadURLParams {
	var ()
	return &GeneratedUploadURLParams{

		Context: ctx,
	}
}

// NewGeneratedUploadURLParamsWithHTTPClient creates a new GeneratedUploadURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeneratedUploadURLParamsWithHTTPClient(client *http.Client) *GeneratedUploadURLParams {
	var ()
	return &GeneratedUploadURLParams{
		HTTPClient: client,
	}
}

/*GeneratedUploadURLParams contains all the parameters to send to the API endpoint
for the generated upload Url operation typically these are written to a http.Request
*/
type GeneratedUploadURLParams struct {

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

// WithTimeout adds the timeout to the generated upload Url params
func (o *GeneratedUploadURLParams) WithTimeout(timeout time.Duration) *GeneratedUploadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generated upload Url params
func (o *GeneratedUploadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generated upload Url params
func (o *GeneratedUploadURLParams) WithContext(ctx context.Context) *GeneratedUploadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generated upload Url params
func (o *GeneratedUploadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generated upload Url params
func (o *GeneratedUploadURLParams) WithHTTPClient(client *http.Client) *GeneratedUploadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generated upload Url params
func (o *GeneratedUploadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileType adds the fileType to the generated upload Url params
func (o *GeneratedUploadURLParams) WithFileType(fileType string) *GeneratedUploadURLParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the generated upload Url params
func (o *GeneratedUploadURLParams) SetFileType(fileType string) {
	o.FileType = fileType
}

// WithFolder adds the folder to the generated upload Url params
func (o *GeneratedUploadURLParams) WithFolder(folder string) *GeneratedUploadURLParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the generated upload Url params
func (o *GeneratedUploadURLParams) SetFolder(folder string) {
	o.Folder = folder
}

// WithNamespace adds the namespace to the generated upload Url params
func (o *GeneratedUploadURLParams) WithNamespace(namespace string) *GeneratedUploadURLParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the generated upload Url params
func (o *GeneratedUploadURLParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GeneratedUploadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
