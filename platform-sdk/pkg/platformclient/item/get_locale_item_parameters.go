// Code generated by go-swagger; DO NOT EDIT.

package item

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
	"github.com/go-openapi/swag"
)

// NewGetLocaleItemParams creates a new GetLocaleItemParams object
// with the default values initialized.
func NewGetLocaleItemParams() *GetLocaleItemParams {
	var (
		activeOnlyDefault     = bool(true)
		populateBundleDefault = bool(false)
	)
	return &GetLocaleItemParams{
		ActiveOnly:     &activeOnlyDefault,
		PopulateBundle: &populateBundleDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocaleItemParamsWithTimeout creates a new GetLocaleItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLocaleItemParamsWithTimeout(timeout time.Duration) *GetLocaleItemParams {
	var (
		activeOnlyDefault     = bool(true)
		populateBundleDefault = bool(false)
	)
	return &GetLocaleItemParams{
		ActiveOnly:     &activeOnlyDefault,
		PopulateBundle: &populateBundleDefault,

		timeout: timeout,
	}
}

// NewGetLocaleItemParamsWithContext creates a new GetLocaleItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLocaleItemParamsWithContext(ctx context.Context) *GetLocaleItemParams {
	var (
		activeOnlyDefault     = bool(true)
		populateBundleDefault = bool(false)
	)
	return &GetLocaleItemParams{
		ActiveOnly:     &activeOnlyDefault,
		PopulateBundle: &populateBundleDefault,

		Context: ctx,
	}
}

// NewGetLocaleItemParamsWithHTTPClient creates a new GetLocaleItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLocaleItemParamsWithHTTPClient(client *http.Client) *GetLocaleItemParams {
	var (
		activeOnlyDefault     = bool(true)
		populateBundleDefault = bool(false)
	)
	return &GetLocaleItemParams{
		ActiveOnly:     &activeOnlyDefault,
		PopulateBundle: &populateBundleDefault,
		HTTPClient:     client,
	}
}

/*GetLocaleItemParams contains all the parameters to send to the API endpoint
for the get locale item operation typically these are written to a http.Request
*/
type GetLocaleItemParams struct {

	/*ActiveOnly*/
	ActiveOnly *bool
	/*ItemID*/
	ItemID string
	/*Language*/
	Language *string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*PopulateBundle
	  whether populate bundled items if it's a bundle

	*/
	PopulateBundle *bool
	/*Region*/
	Region *string
	/*StoreID
	  default is published store id

	*/
	StoreID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get locale item params
func (o *GetLocaleItemParams) WithTimeout(timeout time.Duration) *GetLocaleItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get locale item params
func (o *GetLocaleItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get locale item params
func (o *GetLocaleItemParams) WithContext(ctx context.Context) *GetLocaleItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get locale item params
func (o *GetLocaleItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get locale item params
func (o *GetLocaleItemParams) WithHTTPClient(client *http.Client) *GetLocaleItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get locale item params
func (o *GetLocaleItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the get locale item params
func (o *GetLocaleItemParams) WithActiveOnly(activeOnly *bool) *GetLocaleItemParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the get locale item params
func (o *GetLocaleItemParams) SetActiveOnly(activeOnly *bool) {
	o.ActiveOnly = activeOnly
}

// WithItemID adds the itemID to the get locale item params
func (o *GetLocaleItemParams) WithItemID(itemID string) *GetLocaleItemParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get locale item params
func (o *GetLocaleItemParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithLanguage adds the language to the get locale item params
func (o *GetLocaleItemParams) WithLanguage(language *string) *GetLocaleItemParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get locale item params
func (o *GetLocaleItemParams) SetLanguage(language *string) {
	o.Language = language
}

// WithNamespace adds the namespace to the get locale item params
func (o *GetLocaleItemParams) WithNamespace(namespace string) *GetLocaleItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get locale item params
func (o *GetLocaleItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPopulateBundle adds the populateBundle to the get locale item params
func (o *GetLocaleItemParams) WithPopulateBundle(populateBundle *bool) *GetLocaleItemParams {
	o.SetPopulateBundle(populateBundle)
	return o
}

// SetPopulateBundle adds the populateBundle to the get locale item params
func (o *GetLocaleItemParams) SetPopulateBundle(populateBundle *bool) {
	o.PopulateBundle = populateBundle
}

// WithRegion adds the region to the get locale item params
func (o *GetLocaleItemParams) WithRegion(region *string) *GetLocaleItemParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get locale item params
func (o *GetLocaleItemParams) SetRegion(region *string) {
	o.Region = region
}

// WithStoreID adds the storeID to the get locale item params
func (o *GetLocaleItemParams) WithStoreID(storeID *string) *GetLocaleItemParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get locale item params
func (o *GetLocaleItemParams) SetStoreID(storeID *string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocaleItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveOnly != nil {

		// query param activeOnly
		var qrActiveOnly bool
		if o.ActiveOnly != nil {
			qrActiveOnly = *o.ActiveOnly
		}
		qActiveOnly := swag.FormatBool(qrActiveOnly)
		if qActiveOnly != "" {
			if err := r.SetQueryParam("activeOnly", qActiveOnly); err != nil {
				return err
			}
		}

	}

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

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

	if o.PopulateBundle != nil {

		// query param populateBundle
		var qrPopulateBundle bool
		if o.PopulateBundle != nil {
			qrPopulateBundle = *o.PopulateBundle
		}
		qPopulateBundle := swag.FormatBool(qrPopulateBundle)
		if qPopulateBundle != "" {
			if err := r.SetQueryParam("populateBundle", qPopulateBundle); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

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
