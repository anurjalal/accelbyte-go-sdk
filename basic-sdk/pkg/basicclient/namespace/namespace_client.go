// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new namespace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeNamespaceStatus(params *ChangeNamespaceStatusParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeNamespaceStatusOK, *ChangeNamespaceStatusBadRequest, *ChangeNamespaceStatusUnauthorized, *ChangeNamespaceStatusForbidden, *ChangeNamespaceStatusNotFound, *ChangeNamespaceStatusConflict, error)

	CreateNamespace(params *CreateNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNamespaceCreated, *CreateNamespaceBadRequest, *CreateNamespaceUnauthorized, *CreateNamespaceForbidden, *CreateNamespaceConflict, error)

	DeleteNamespace(params *DeleteNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNamespaceOK, *DeleteNamespaceBadRequest, *DeleteNamespaceUnauthorized, *DeleteNamespaceForbidden, *DeleteNamespaceNotFound, *DeleteNamespaceConflict, error)

	GetNamespace(params *GetNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespaceOK, *GetNamespaceBadRequest, *GetNamespaceUnauthorized, *GetNamespaceForbidden, *GetNamespaceNotFound, error)

	GetNamespacePublisher(params *GetNamespacePublisherParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespacePublisherOK, *GetNamespacePublisherBadRequest, *GetNamespacePublisherUnauthorized, *GetNamespacePublisherForbidden, *GetNamespacePublisherNotFound, error)

	GetNamespaces(params *GetNamespacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespacesOK, *GetNamespacesUnauthorized, *GetNamespacesForbidden, error)

	PublicGetNamespacePublisher(params *PublicGetNamespacePublisherParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetNamespacePublisherOK, *PublicGetNamespacePublisherBadRequest, *PublicGetNamespacePublisherUnauthorized, *PublicGetNamespacePublisherForbidden, *PublicGetNamespacePublisherNotFound, error)

	PublicGetNamespaces(params *PublicGetNamespacesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetNamespacesOK, *PublicGetNamespacesUnauthorized, error)

	UpdateNamespace(params *UpdateNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNamespaceOK, *UpdateNamespaceBadRequest, *UpdateNamespaceUnauthorized, *UpdateNamespaceForbidden, *UpdateNamespaceNotFound, *UpdateNamespaceConflict, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeNamespaceStatus changes namespace status

  Change a namespace status.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE:{namespace}:NAMESPACE"</b>, action=4 <b>(UPDATE)</b></li><li>Action code<i></i>: 11306</li><li><i>Returns</i>: updated namespace</li></ul>
*/
func (a *Client) ChangeNamespaceStatus(params *ChangeNamespaceStatusParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeNamespaceStatusOK, *ChangeNamespaceStatusBadRequest, *ChangeNamespaceStatusUnauthorized, *ChangeNamespaceStatusForbidden, *ChangeNamespaceStatusNotFound, *ChangeNamespaceStatusConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeNamespaceStatusParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeNamespaceStatus",
		Method:             "PATCH",
		PathPattern:        "/v1/admin/namespaces/{namespace}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeNamespaceStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ChangeNamespaceStatusOK:
		return v, nil, nil, nil, nil, nil, nil
	case *ChangeNamespaceStatusBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *ChangeNamespaceStatusUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *ChangeNamespaceStatusForbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *ChangeNamespaceStatusNotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *ChangeNamespaceStatusConflict:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateNamespace creates a namespace

  Create a namespace.<br>By default the namespace is enabled.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE"</b>, action=1 <b>(CREATE)</b></li><li><i>Action code</i>: 11301</li><li><i>Returns</i>: created namespace</li></ul>
*/
func (a *Client) CreateNamespace(params *CreateNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNamespaceCreated, *CreateNamespaceBadRequest, *CreateNamespaceUnauthorized, *CreateNamespaceForbidden, *CreateNamespaceConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNamespaceParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNamespace",
		Method:             "POST",
		PathPattern:        "/v1/admin/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateNamespaceCreated:
		return v, nil, nil, nil, nil, nil
	case *CreateNamespaceBadRequest:
		return nil, v, nil, nil, nil, nil
	case *CreateNamespaceUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *CreateNamespaceForbidden:
		return nil, nil, nil, v, nil, nil
	case *CreateNamespaceConflict:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteNamespace deletes a namespace

  Delete a namespace.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE:{namespace}:NAMESPACE"</b>, action=8 <b>(DELETE)</b></li><li><i>Action code</i>: 11307</li><li><i>Returns</i>: deleted namespace</li></ul>
*/
func (a *Client) DeleteNamespace(params *DeleteNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNamespaceOK, *DeleteNamespaceBadRequest, *DeleteNamespaceUnauthorized, *DeleteNamespaceForbidden, *DeleteNamespaceNotFound, *DeleteNamespaceConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNamespace",
		Method:             "DELETE",
		PathPattern:        "/v1/admin/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteNamespaceOK:
		return v, nil, nil, nil, nil, nil, nil
	case *DeleteNamespaceBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeleteNamespaceUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeleteNamespaceForbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *DeleteNamespaceNotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *DeleteNamespaceConflict:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetNamespace gets a namespace

  Get a namespace.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE:{namespace}:NAMESPACE"</b>, action=2 <b>(READ)</b></li><li><i>Action code</i>: 11304</li><li><i>Returns</i>: namespace</li></ul>
*/
func (a *Client) GetNamespace(params *GetNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespaceOK, *GetNamespaceBadRequest, *GetNamespaceUnauthorized, *GetNamespaceForbidden, *GetNamespaceNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNamespace",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetNamespaceOK:
		return v, nil, nil, nil, nil, nil
	case *GetNamespaceBadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetNamespaceUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetNamespaceForbidden:
		return nil, nil, nil, v, nil, nil
	case *GetNamespaceNotFound:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetNamespacePublisher gets namespace info related publisher namespace

  Get namespace info related publisher namespace.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE:{namespace}:NAMESPACE"</b>, action=2 <b>(READ)</b></li><li><i>Action code</i>: 11305</li><li><i>Returns</i>: Namespace info related publisher namespace</li></ul>
*/
func (a *Client) GetNamespacePublisher(params *GetNamespacePublisherParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespacePublisherOK, *GetNamespacePublisherBadRequest, *GetNamespacePublisherUnauthorized, *GetNamespacePublisherForbidden, *GetNamespacePublisherNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespacePublisherParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNamespacePublisher",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/publisher",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNamespacePublisherReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetNamespacePublisherOK:
		return v, nil, nil, nil, nil, nil
	case *GetNamespacePublisherBadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetNamespacePublisherUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetNamespacePublisherForbidden:
		return nil, nil, nil, v, nil, nil
	case *GetNamespacePublisherNotFound:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetNamespaces gets all namespaces

  Get all namespaces.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE"</b>, action=2 <b>(READ)</b></li><li><i>Action code</i>: 11303</li><li><i>Returns</i>: list of namespaces</li></ul>
*/
func (a *Client) GetNamespaces(params *GetNamespacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNamespacesOK, *GetNamespacesUnauthorized, *GetNamespacesForbidden, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespacesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNamespaces",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNamespacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetNamespacesOK:
		return v, nil, nil, nil
	case *GetNamespacesUnauthorized:
		return nil, v, nil, nil
	case *GetNamespacesForbidden:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetNamespacePublisher gets namespace info related publisher namespace

  Get namespace info related publisher namespace.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"NAMESPACE:{namespace}:NAMESPACE"</b>, action=2 <b>(READ)</b></li><li><i>Action code</i>: 11305</li><li><i>Returns</i>: Namespace info related publisher namespace</li></ul>
*/
func (a *Client) PublicGetNamespacePublisher(params *PublicGetNamespacePublisherParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetNamespacePublisherOK, *PublicGetNamespacePublisherBadRequest, *PublicGetNamespacePublisherUnauthorized, *PublicGetNamespacePublisherForbidden, *PublicGetNamespacePublisherNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetNamespacePublisherParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetNamespacePublisher",
		Method:             "GET",
		PathPattern:        "/v1/public/namespaces/{namespace}/publisher",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetNamespacePublisherReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetNamespacePublisherOK:
		return v, nil, nil, nil, nil, nil
	case *PublicGetNamespacePublisherBadRequest:
		return nil, v, nil, nil, nil, nil
	case *PublicGetNamespacePublisherUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *PublicGetNamespacePublisherForbidden:
		return nil, nil, nil, v, nil, nil
	case *PublicGetNamespacePublisherNotFound:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetNamespaces gets all namespaces

  Get all namespaces.<br>Other detail info: <ul><li><i>Required permission</i>: login user</li><li><i>Action code</i>: 11303</li><li><i>Returns</i>: list of namespaces</li></ul>
*/
func (a *Client) PublicGetNamespaces(params *PublicGetNamespacesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetNamespacesOK, *PublicGetNamespacesUnauthorized, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetNamespacesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetNamespaces",
		Method:             "GET",
		PathPattern:        "/v1/public/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetNamespacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetNamespacesOK:
		return v, nil, nil
	case *PublicGetNamespacesUnauthorized:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateNamespace updates namespace basic info

  Update namespace basic info.<br>Other detail info: <ul><li><i>Required permission</i>: resource=<b>"ADMIN:NAMESPACE:{namespace}:NAMESPACE"</b>, action=4 <b>(UPDATE)</b></li><li><i>Action code</i>: 11302</li><li><i>Returns</i>: updated namespace</li></ul>
*/
func (a *Client) UpdateNamespace(params *UpdateNamespaceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNamespaceOK, *UpdateNamespaceBadRequest, *UpdateNamespaceUnauthorized, *UpdateNamespaceForbidden, *UpdateNamespaceNotFound, *UpdateNamespaceConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNamespaceParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNamespace",
		Method:             "PATCH",
		PathPattern:        "/v1/admin/namespaces/{namespace}/basic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateNamespaceOK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateNamespaceBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateNamespaceUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateNamespaceForbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateNamespaceNotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateNamespaceConflict:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
