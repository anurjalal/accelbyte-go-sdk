// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// NewQueryUserIAPOrdersParams creates a new QueryUserIAPOrdersParams object
// with the default values initialized.
func NewQueryUserIAPOrdersParams() *QueryUserIAPOrdersParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &QueryUserIAPOrdersParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryUserIAPOrdersParamsWithTimeout creates a new QueryUserIAPOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryUserIAPOrdersParamsWithTimeout(timeout time.Duration) *QueryUserIAPOrdersParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &QueryUserIAPOrdersParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewQueryUserIAPOrdersParamsWithContext creates a new QueryUserIAPOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryUserIAPOrdersParamsWithContext(ctx context.Context) *QueryUserIAPOrdersParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &QueryUserIAPOrdersParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewQueryUserIAPOrdersParamsWithHTTPClient creates a new QueryUserIAPOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryUserIAPOrdersParamsWithHTTPClient(client *http.Client) *QueryUserIAPOrdersParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &QueryUserIAPOrdersParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*QueryUserIAPOrdersParams contains all the parameters to send to the API endpoint
for the query user i a p orders operation typically these are written to a http.Request
*/
type QueryUserIAPOrdersParams struct {

	/*EndTime
	  end time is exclusive, using ISO 8601 format e.g. yyyy-MM-dd'T'HH:mm:ssZZ

	*/
	EndTime *string
	/*Limit*/
	Limit *int32
	/*Namespace*/
	Namespace string
	/*Offset*/
	Offset *int32
	/*ProductID*/
	ProductID *string
	/*StartTime
	  start time is inclusive, using ISO 8601 format e.g. yyyy-MM-dd'T'HH:mm:ssZZ

	*/
	StartTime *string
	/*Status*/
	Status *string
	/*Type*/
	Type *string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithTimeout(timeout time.Duration) *QueryUserIAPOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithContext(ctx context.Context) *QueryUserIAPOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithHTTPClient(client *http.Client) *QueryUserIAPOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithEndTime(endTime *string) *QueryUserIAPOrdersParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithLimit adds the limit to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithLimit(limit *int32) *QueryUserIAPOrdersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithNamespace(namespace string) *QueryUserIAPOrdersParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithOffset(offset *int32) *QueryUserIAPOrdersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithProductID adds the productID to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithProductID(productID *string) *QueryUserIAPOrdersParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetProductID(productID *string) {
	o.ProductID = productID
}

// WithStartTime adds the startTime to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithStartTime(startTime *string) *QueryUserIAPOrdersParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WithStatus adds the status to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithStatus(status *string) *QueryUserIAPOrdersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithType(typeVar *string) *QueryUserIAPOrdersParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUserID adds the userID to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) WithUserID(userID string) *QueryUserIAPOrdersParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the query user i a p orders params
func (o *QueryUserIAPOrdersParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryUserIAPOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime string
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {
			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
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
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.ProductID != nil {

		// query param productId
		var qrProductID string
		if o.ProductID != nil {
			qrProductID = *o.ProductID
		}
		qProductID := qrProductID
		if qProductID != "" {
			if err := r.SetQueryParam("productId", qProductID); err != nil {
				return err
			}
		}

	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime string
		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {
			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

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
