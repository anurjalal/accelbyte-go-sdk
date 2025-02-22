// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewGetTopicByTopicNameParams creates a new GetTopicByTopicNameParams object
// with the default values initialized.
func NewGetTopicByTopicNameParams() *GetTopicByTopicNameParams {
	var ()
	return &GetTopicByTopicNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTopicByTopicNameParamsWithTimeout creates a new GetTopicByTopicNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTopicByTopicNameParamsWithTimeout(timeout time.Duration) *GetTopicByTopicNameParams {
	var ()
	return &GetTopicByTopicNameParams{

		timeout: timeout,
	}
}

// NewGetTopicByTopicNameParamsWithContext creates a new GetTopicByTopicNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTopicByTopicNameParamsWithContext(ctx context.Context) *GetTopicByTopicNameParams {
	var ()
	return &GetTopicByTopicNameParams{

		Context: ctx,
	}
}

// NewGetTopicByTopicNameParamsWithHTTPClient creates a new GetTopicByTopicNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTopicByTopicNameParamsWithHTTPClient(client *http.Client) *GetTopicByTopicNameParams {
	var ()
	return &GetTopicByTopicNameParams{
		HTTPClient: client,
	}
}

/*GetTopicByTopicNameParams contains all the parameters to send to the API endpoint
for the get topic by topic name operation typically these are written to a http.Request
*/
type GetTopicByTopicNameParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*Topic
	  topic name

	*/
	Topic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get topic by topic name params
func (o *GetTopicByTopicNameParams) WithTimeout(timeout time.Duration) *GetTopicByTopicNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get topic by topic name params
func (o *GetTopicByTopicNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get topic by topic name params
func (o *GetTopicByTopicNameParams) WithContext(ctx context.Context) *GetTopicByTopicNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get topic by topic name params
func (o *GetTopicByTopicNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get topic by topic name params
func (o *GetTopicByTopicNameParams) WithHTTPClient(client *http.Client) *GetTopicByTopicNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get topic by topic name params
func (o *GetTopicByTopicNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get topic by topic name params
func (o *GetTopicByTopicNameParams) WithNamespace(namespace string) *GetTopicByTopicNameParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get topic by topic name params
func (o *GetTopicByTopicNameParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTopic adds the topic to the get topic by topic name params
func (o *GetTopicByTopicNameParams) WithTopic(topic string) *GetTopicByTopicNameParams {
	o.SetTopic(topic)
	return o
}

// SetTopic adds the topic to the get topic by topic name params
func (o *GetTopicByTopicNameParams) SetTopic(topic string) {
	o.Topic = topic
}

// WriteToRequest writes these params to a swagger request
func (o *GetTopicByTopicNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param topic
	if err := r.SetPathParam("topic", o.Topic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
