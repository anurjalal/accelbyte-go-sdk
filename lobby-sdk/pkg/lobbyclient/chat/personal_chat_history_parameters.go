// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// NewPersonalChatHistoryParams creates a new PersonalChatHistoryParams object
// with the default values initialized.
func NewPersonalChatHistoryParams() *PersonalChatHistoryParams {
	var ()
	return &PersonalChatHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPersonalChatHistoryParamsWithTimeout creates a new PersonalChatHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPersonalChatHistoryParamsWithTimeout(timeout time.Duration) *PersonalChatHistoryParams {
	var ()
	return &PersonalChatHistoryParams{

		timeout: timeout,
	}
}

// NewPersonalChatHistoryParamsWithContext creates a new PersonalChatHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPersonalChatHistoryParamsWithContext(ctx context.Context) *PersonalChatHistoryParams {
	var ()
	return &PersonalChatHistoryParams{

		Context: ctx,
	}
}

// NewPersonalChatHistoryParamsWithHTTPClient creates a new PersonalChatHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPersonalChatHistoryParamsWithHTTPClient(client *http.Client) *PersonalChatHistoryParams {
	var ()
	return &PersonalChatHistoryParams{
		HTTPClient: client,
	}
}

/*PersonalChatHistoryParams contains all the parameters to send to the API endpoint
for the personal chat history operation typically these are written to a http.Request
*/
type PersonalChatHistoryParams struct {

	/*FriendID
	  user ID that receive the chat

	*/
	FriendID string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the personal chat history params
func (o *PersonalChatHistoryParams) WithTimeout(timeout time.Duration) *PersonalChatHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the personal chat history params
func (o *PersonalChatHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the personal chat history params
func (o *PersonalChatHistoryParams) WithContext(ctx context.Context) *PersonalChatHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the personal chat history params
func (o *PersonalChatHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the personal chat history params
func (o *PersonalChatHistoryParams) WithHTTPClient(client *http.Client) *PersonalChatHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the personal chat history params
func (o *PersonalChatHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFriendID adds the friendID to the personal chat history params
func (o *PersonalChatHistoryParams) WithFriendID(friendID string) *PersonalChatHistoryParams {
	o.SetFriendID(friendID)
	return o
}

// SetFriendID adds the friendId to the personal chat history params
func (o *PersonalChatHistoryParams) SetFriendID(friendID string) {
	o.FriendID = friendID
}

// WithNamespace adds the namespace to the personal chat history params
func (o *PersonalChatHistoryParams) WithNamespace(namespace string) *PersonalChatHistoryParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the personal chat history params
func (o *PersonalChatHistoryParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the personal chat history params
func (o *PersonalChatHistoryParams) WithUserID(userID string) *PersonalChatHistoryParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the personal chat history params
func (o *PersonalChatHistoryParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PersonalChatHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param friendId
	if err := r.SetPathParam("friendId", o.FriendID); err != nil {
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
