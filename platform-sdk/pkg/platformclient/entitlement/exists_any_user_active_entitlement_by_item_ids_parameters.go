// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// NewExistsAnyUserActiveEntitlementByItemIdsParams creates a new ExistsAnyUserActiveEntitlementByItemIdsParams object
// with the default values initialized.
func NewExistsAnyUserActiveEntitlementByItemIdsParams() *ExistsAnyUserActiveEntitlementByItemIdsParams {
	var ()
	return &ExistsAnyUserActiveEntitlementByItemIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExistsAnyUserActiveEntitlementByItemIdsParamsWithTimeout creates a new ExistsAnyUserActiveEntitlementByItemIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExistsAnyUserActiveEntitlementByItemIdsParamsWithTimeout(timeout time.Duration) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	var ()
	return &ExistsAnyUserActiveEntitlementByItemIdsParams{

		timeout: timeout,
	}
}

// NewExistsAnyUserActiveEntitlementByItemIdsParamsWithContext creates a new ExistsAnyUserActiveEntitlementByItemIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewExistsAnyUserActiveEntitlementByItemIdsParamsWithContext(ctx context.Context) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	var ()
	return &ExistsAnyUserActiveEntitlementByItemIdsParams{

		Context: ctx,
	}
}

// NewExistsAnyUserActiveEntitlementByItemIdsParamsWithHTTPClient creates a new ExistsAnyUserActiveEntitlementByItemIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExistsAnyUserActiveEntitlementByItemIdsParamsWithHTTPClient(client *http.Client) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	var ()
	return &ExistsAnyUserActiveEntitlementByItemIdsParams{
		HTTPClient: client,
	}
}

/*ExistsAnyUserActiveEntitlementByItemIdsParams contains all the parameters to send to the API endpoint
for the exists any user active entitlement by item ids operation typically these are written to a http.Request
*/
type ExistsAnyUserActiveEntitlementByItemIdsParams struct {

	/*ItemIds*/
	ItemIds []string
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithTimeout(timeout time.Duration) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithContext(ctx context.Context) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithHTTPClient(client *http.Client) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemIds adds the itemIds to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithItemIds(itemIds []string) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetItemIds(itemIds)
	return o
}

// SetItemIds adds the itemIds to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetItemIds(itemIds []string) {
	o.ItemIds = itemIds
}

// WithNamespace adds the namespace to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithNamespace(namespace string) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WithUserID(userID string) *ExistsAnyUserActiveEntitlementByItemIdsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the exists any user active entitlement by item ids params
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ExistsAnyUserActiveEntitlementByItemIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesItemIds := o.ItemIds

	joinedItemIds := swag.JoinByFormat(valuesItemIds, "multi")
	// query array param itemIds
	if err := r.SetQueryParam("itemIds", joinedItemIds...); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
