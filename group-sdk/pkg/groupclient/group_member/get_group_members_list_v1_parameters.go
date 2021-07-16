// Code generated by go-swagger; DO NOT EDIT.

package group_member

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

// NewGetGroupMembersListV1Params creates a new GetGroupMembersListV1Params object
// with the default values initialized.
func NewGetGroupMembersListV1Params() *GetGroupMembersListV1Params {
	var ()
	return &GetGroupMembersListV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupMembersListV1ParamsWithTimeout creates a new GetGroupMembersListV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupMembersListV1ParamsWithTimeout(timeout time.Duration) *GetGroupMembersListV1Params {
	var ()
	return &GetGroupMembersListV1Params{

		timeout: timeout,
	}
}

// NewGetGroupMembersListV1ParamsWithContext creates a new GetGroupMembersListV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupMembersListV1ParamsWithContext(ctx context.Context) *GetGroupMembersListV1Params {
	var ()
	return &GetGroupMembersListV1Params{

		Context: ctx,
	}
}

// NewGetGroupMembersListV1ParamsWithHTTPClient creates a new GetGroupMembersListV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupMembersListV1ParamsWithHTTPClient(client *http.Client) *GetGroupMembersListV1Params {
	var ()
	return &GetGroupMembersListV1Params{
		HTTPClient: client,
	}
}

/*GetGroupMembersListV1Params contains all the parameters to send to the API endpoint
for the get group members list v1 operation typically these are written to a http.Request
*/
type GetGroupMembersListV1Params struct {

	/*GroupID
	  Group ID

	*/
	GroupID string
	/*Limit
	  size of displayed data

	*/
	Limit *int64
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Offset
	  The start position that points to query data

	*/
	Offset *int64
	/*Order
	  Sort group members list by User ID with ascending as default order. Set it 'desc' for descending order

	*/
	Order *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithTimeout(timeout time.Duration) *GetGroupMembersListV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithContext(ctx context.Context) *GetGroupMembersListV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithHTTPClient(client *http.Client) *GetGroupMembersListV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithGroupID(groupID string) *GetGroupMembersListV1Params {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithLimit(limit *int64) *GetGroupMembersListV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithNamespace(namespace string) *GetGroupMembersListV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithOffset(offset *int64) *GetGroupMembersListV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrder adds the order to the get group members list v1 params
func (o *GetGroupMembersListV1Params) WithOrder(order *string) *GetGroupMembersListV1Params {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get group members list v1 params
func (o *GetGroupMembersListV1Params) SetOrder(order *string) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupMembersListV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
