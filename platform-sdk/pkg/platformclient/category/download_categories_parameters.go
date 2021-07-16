// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewDownloadCategoriesParams creates a new DownloadCategoriesParams object
// with the default values initialized.
func NewDownloadCategoriesParams() *DownloadCategoriesParams {
	var ()
	return &DownloadCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadCategoriesParamsWithTimeout creates a new DownloadCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadCategoriesParamsWithTimeout(timeout time.Duration) *DownloadCategoriesParams {
	var ()
	return &DownloadCategoriesParams{

		timeout: timeout,
	}
}

// NewDownloadCategoriesParamsWithContext creates a new DownloadCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadCategoriesParamsWithContext(ctx context.Context) *DownloadCategoriesParams {
	var ()
	return &DownloadCategoriesParams{

		Context: ctx,
	}
}

// NewDownloadCategoriesParamsWithHTTPClient creates a new DownloadCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadCategoriesParamsWithHTTPClient(client *http.Client) *DownloadCategoriesParams {
	var ()
	return &DownloadCategoriesParams{
		HTTPClient: client,
	}
}

/*DownloadCategoriesParams contains all the parameters to send to the API endpoint
for the download categories operation typically these are written to a http.Request
*/
type DownloadCategoriesParams struct {

	/*Language*/
	Language *string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*StoreID
	  default is published store id

	*/
	StoreID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download categories params
func (o *DownloadCategoriesParams) WithTimeout(timeout time.Duration) *DownloadCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download categories params
func (o *DownloadCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download categories params
func (o *DownloadCategoriesParams) WithContext(ctx context.Context) *DownloadCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download categories params
func (o *DownloadCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download categories params
func (o *DownloadCategoriesParams) WithHTTPClient(client *http.Client) *DownloadCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download categories params
func (o *DownloadCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the download categories params
func (o *DownloadCategoriesParams) WithLanguage(language *string) *DownloadCategoriesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the download categories params
func (o *DownloadCategoriesParams) SetLanguage(language *string) {
	o.Language = language
}

// WithNamespace adds the namespace to the download categories params
func (o *DownloadCategoriesParams) WithNamespace(namespace string) *DownloadCategoriesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the download categories params
func (o *DownloadCategoriesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the download categories params
func (o *DownloadCategoriesParams) WithStoreID(storeID *string) *DownloadCategoriesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the download categories params
func (o *DownloadCategoriesParams) SetStoreID(storeID *string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.StoreID != nil {

		// query param storeId
		var qrStoreID string
		if o.StoreID != nil {
			qrStoreID = *o.StoreID
		}
		qStoreID := qrStoreID
		if qStoreID != "" {
			if err := r.SetQueryParam("storeId", qStoreID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
