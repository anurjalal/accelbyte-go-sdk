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
)

// NewFeatureItemParams creates a new FeatureItemParams object
// with the default values initialized.
func NewFeatureItemParams() *FeatureItemParams {
	var ()
	return &FeatureItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFeatureItemParamsWithTimeout creates a new FeatureItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFeatureItemParamsWithTimeout(timeout time.Duration) *FeatureItemParams {
	var ()
	return &FeatureItemParams{

		timeout: timeout,
	}
}

// NewFeatureItemParamsWithContext creates a new FeatureItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewFeatureItemParamsWithContext(ctx context.Context) *FeatureItemParams {
	var ()
	return &FeatureItemParams{

		Context: ctx,
	}
}

// NewFeatureItemParamsWithHTTPClient creates a new FeatureItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFeatureItemParamsWithHTTPClient(client *http.Client) *FeatureItemParams {
	var ()
	return &FeatureItemParams{
		HTTPClient: client,
	}
}

/*FeatureItemParams contains all the parameters to send to the API endpoint
for the feature item operation typically these are written to a http.Request
*/
type FeatureItemParams struct {

	/*Feature*/
	Feature string
	/*ItemID*/
	ItemID string
	/*Namespace*/
	Namespace string
	/*StoreID*/
	StoreID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the feature item params
func (o *FeatureItemParams) WithTimeout(timeout time.Duration) *FeatureItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the feature item params
func (o *FeatureItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the feature item params
func (o *FeatureItemParams) WithContext(ctx context.Context) *FeatureItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the feature item params
func (o *FeatureItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the feature item params
func (o *FeatureItemParams) WithHTTPClient(client *http.Client) *FeatureItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the feature item params
func (o *FeatureItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeature adds the feature to the feature item params
func (o *FeatureItemParams) WithFeature(feature string) *FeatureItemParams {
	o.SetFeature(feature)
	return o
}

// SetFeature adds the feature to the feature item params
func (o *FeatureItemParams) SetFeature(feature string) {
	o.Feature = feature
}

// WithItemID adds the itemID to the feature item params
func (o *FeatureItemParams) WithItemID(itemID string) *FeatureItemParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the feature item params
func (o *FeatureItemParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithNamespace adds the namespace to the feature item params
func (o *FeatureItemParams) WithNamespace(namespace string) *FeatureItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the feature item params
func (o *FeatureItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the feature item params
func (o *FeatureItemParams) WithStoreID(storeID string) *FeatureItemParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the feature item params
func (o *FeatureItemParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *FeatureItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param feature
	if err := r.SetPathParam("feature", o.Feature); err != nil {
		return err
	}

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param storeId
	qrStoreID := o.StoreID
	qStoreID := qrStoreID
	if qStoreID != "" {
		if err := r.SetQueryParam("storeId", qStoreID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
