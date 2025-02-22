// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployment config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeployment(params *CreateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentCreated, *CreateDeploymentBadRequest, *CreateDeploymentUnauthorized, *CreateDeploymentConflict, *CreateDeploymentInternalServerError, error)

	CreateDeploymentOverride(params *CreateDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentOverrideCreated, *CreateDeploymentOverrideBadRequest, *CreateDeploymentOverrideUnauthorized, *CreateDeploymentOverrideNotFound, *CreateDeploymentOverrideConflict, *CreateDeploymentOverrideInternalServerError, error)

	CreateOverrideRegionOverride(params *CreateOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOverrideRegionOverrideCreated, *CreateOverrideRegionOverrideBadRequest, *CreateOverrideRegionOverrideUnauthorized, *CreateOverrideRegionOverrideNotFound, *CreateOverrideRegionOverrideConflict, *CreateOverrideRegionOverrideInternalServerError, error)

	CreateRootRegionOverride(params *CreateRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRootRegionOverrideCreated, *CreateRootRegionOverrideBadRequest, *CreateRootRegionOverrideUnauthorized, *CreateRootRegionOverrideNotFound, *CreateRootRegionOverrideConflict, *CreateRootRegionOverrideInternalServerError, error)

	DeleteDeployment(params *DeleteDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentNoContent, *DeleteDeploymentBadRequest, *DeleteDeploymentUnauthorized, *DeleteDeploymentNotFound, *DeleteDeploymentInternalServerError, error)

	DeleteDeploymentOverride(params *DeleteDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentOverrideOK, *DeleteDeploymentOverrideBadRequest, *DeleteDeploymentOverrideUnauthorized, *DeleteDeploymentOverrideNotFound, *DeleteDeploymentOverrideInternalServerError, error)

	DeleteOverrideRegionOverride(params *DeleteOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOverrideRegionOverrideOK, *DeleteOverrideRegionOverrideBadRequest, *DeleteOverrideRegionOverrideUnauthorized, *DeleteOverrideRegionOverrideNotFound, *DeleteOverrideRegionOverrideInternalServerError, error)

	DeleteRootRegionOverride(params *DeleteRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRootRegionOverrideOK, *DeleteRootRegionOverrideBadRequest, *DeleteRootRegionOverrideUnauthorized, *DeleteRootRegionOverrideNotFound, *DeleteRootRegionOverrideInternalServerError, error)

	GetAllDeployment(params *GetAllDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllDeploymentOK, *GetAllDeploymentBadRequest, *GetAllDeploymentUnauthorized, *GetAllDeploymentInternalServerError, error)

	GetDeployment(params *GetDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentOK, *GetDeploymentBadRequest, *GetDeploymentUnauthorized, *GetDeploymentNotFound, *GetDeploymentInternalServerError, error)

	UpdateDeployment(params *UpdateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentOK, *UpdateDeploymentBadRequest, *UpdateDeploymentUnauthorized, *UpdateDeploymentNotFound, *UpdateDeploymentInternalServerError, error)

	UpdateDeploymentOverride(params *UpdateDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentOverrideOK, *UpdateDeploymentOverrideBadRequest, *UpdateDeploymentOverrideUnauthorized, *UpdateDeploymentOverrideNotFound, *UpdateDeploymentOverrideInternalServerError, error)

	UpdateOverrideRegionOverride(params *UpdateOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOverrideRegionOverrideOK, *UpdateOverrideRegionOverrideBadRequest, *UpdateOverrideRegionOverrideUnauthorized, *UpdateOverrideRegionOverrideNotFound, *UpdateOverrideRegionOverrideInternalServerError, error)

	UpdateRootRegionOverride(params *UpdateRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRootRegionOverrideOK, *UpdateRootRegionOverrideBadRequest, *UpdateRootRegionOverrideUnauthorized, *UpdateRootRegionOverrideNotFound, *UpdateRootRegionOverrideInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDeployment creates deployment

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint create a dedicated servers deployment in a namespace.
*/
func (a *Client) CreateDeployment(params *CreateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentCreated, *CreateDeploymentBadRequest, *CreateDeploymentUnauthorized, *CreateDeploymentConflict, *CreateDeploymentInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDeployment",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateDeploymentCreated:
		return v, nil, nil, nil, nil, nil
	case *CreateDeploymentBadRequest:
		return nil, v, nil, nil, nil, nil
	case *CreateDeploymentUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *CreateDeploymentConflict:
		return nil, nil, nil, v, nil, nil
	case *CreateDeploymentInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateDeploymentOverride creates deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint create a dedicated servers deployment override in a namespace.
*/
func (a *Client) CreateDeploymentOverride(params *CreateDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentOverrideCreated, *CreateDeploymentOverrideBadRequest, *CreateDeploymentOverrideUnauthorized, *CreateDeploymentOverrideNotFound, *CreateDeploymentOverrideConflict, *CreateDeploymentOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDeploymentOverride",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/version/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateDeploymentOverrideCreated:
		return v, nil, nil, nil, nil, nil, nil
	case *CreateDeploymentOverrideBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *CreateDeploymentOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *CreateDeploymentOverrideNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *CreateDeploymentOverrideConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *CreateDeploymentOverrideInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateOverrideRegionOverride creates region override for deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint creates a dedicated servers deployment override in a namespace in a region for deployment overrides.
*/
func (a *Client) CreateOverrideRegionOverride(params *CreateOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOverrideRegionOverrideCreated, *CreateOverrideRegionOverrideBadRequest, *CreateOverrideRegionOverrideUnauthorized, *CreateOverrideRegionOverrideNotFound, *CreateOverrideRegionOverrideConflict, *CreateOverrideRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOverrideRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateOverrideRegionOverride",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateOverrideRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateOverrideRegionOverrideCreated:
		return v, nil, nil, nil, nil, nil, nil
	case *CreateOverrideRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *CreateOverrideRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *CreateOverrideRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *CreateOverrideRegionOverrideConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *CreateOverrideRegionOverrideInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateRootRegionOverride creates region override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint creates a dedicated servers deployment override in a namespace in a region for root deployment.
*/
func (a *Client) CreateRootRegionOverride(params *CreateRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRootRegionOverrideCreated, *CreateRootRegionOverrideBadRequest, *CreateRootRegionOverrideUnauthorized, *CreateRootRegionOverrideNotFound, *CreateRootRegionOverrideConflict, *CreateRootRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRootRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRootRegionOverride",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRootRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateRootRegionOverrideCreated:
		return v, nil, nil, nil, nil, nil, nil
	case *CreateRootRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *CreateRootRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *CreateRootRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *CreateRootRegionOverrideConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *CreateRootRegionOverrideInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteDeployment deletes deployment

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint delete a dedicated server deployment in a namespace
*/
func (a *Client) DeleteDeployment(params *DeleteDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentNoContent, *DeleteDeploymentBadRequest, *DeleteDeploymentUnauthorized, *DeleteDeploymentNotFound, *DeleteDeploymentInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDeployment",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteDeploymentNoContent:
		return v, nil, nil, nil, nil, nil
	case *DeleteDeploymentBadRequest:
		return nil, v, nil, nil, nil, nil
	case *DeleteDeploymentUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *DeleteDeploymentNotFound:
		return nil, nil, nil, v, nil, nil
	case *DeleteDeploymentInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteDeploymentOverride deletes deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [DELETE]

Required scope: social

This endpoint delete a dedicated server deployment override in a namespace
*/
func (a *Client) DeleteDeploymentOverride(params *DeleteDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentOverrideOK, *DeleteDeploymentOverrideBadRequest, *DeleteDeploymentOverrideUnauthorized, *DeleteDeploymentOverrideNotFound, *DeleteDeploymentOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDeploymentOverride",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeploymentOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteDeploymentOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *DeleteDeploymentOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *DeleteDeploymentOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *DeleteDeploymentOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *DeleteDeploymentOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteOverrideRegionOverride deletes region override for deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [DELETE]

Required scope: social

This endpoint delete a dedicated server deployment override in a namespace in a region for deployment overrides
*/
func (a *Client) DeleteOverrideRegionOverride(params *DeleteOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOverrideRegionOverrideOK, *DeleteOverrideRegionOverrideBadRequest, *DeleteOverrideRegionOverrideUnauthorized, *DeleteOverrideRegionOverrideNotFound, *DeleteOverrideRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOverrideRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOverrideRegionOverride",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOverrideRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteOverrideRegionOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *DeleteOverrideRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *DeleteOverrideRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *DeleteOverrideRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *DeleteOverrideRegionOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteRootRegionOverride deletes region override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [DELETE]

Required scope: social

This endpoint delete a dedicated server deployment override in a namespace in a region for root deployment
*/
func (a *Client) DeleteRootRegionOverride(params *DeleteRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRootRegionOverrideOK, *DeleteRootRegionOverrideBadRequest, *DeleteRootRegionOverrideUnauthorized, *DeleteRootRegionOverrideNotFound, *DeleteRootRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRootRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRootRegionOverride",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRootRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteRootRegionOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *DeleteRootRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *DeleteRootRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *DeleteRootRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *DeleteRootRegionOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetAllDeployment gets all deployments

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get a all deployments in a namespace
*/
func (a *Client) GetAllDeployment(params *GetAllDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllDeploymentOK, *GetAllDeploymentBadRequest, *GetAllDeploymentUnauthorized, *GetAllDeploymentInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllDeploymentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAllDeployment",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetAllDeploymentOK:
		return v, nil, nil, nil, nil
	case *GetAllDeploymentBadRequest:
		return nil, v, nil, nil, nil
	case *GetAllDeploymentUnauthorized:
		return nil, nil, v, nil, nil
	case *GetAllDeploymentInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetDeployment gets deployment

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get a dedicated server deployment in a namespace
*/
func (a *Client) GetDeployment(params *GetDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentOK, *GetDeploymentBadRequest, *GetDeploymentUnauthorized, *GetDeploymentNotFound, *GetDeploymentInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeployment",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetDeploymentOK:
		return v, nil, nil, nil, nil, nil
	case *GetDeploymentBadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetDeploymentUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetDeploymentNotFound:
		return nil, nil, nil, v, nil, nil
	case *GetDeploymentInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateDeployment updates deployment

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint update a dedicated servers deployment in a namespace.
*/
func (a *Client) UpdateDeployment(params *UpdateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentOK, *UpdateDeploymentBadRequest, *UpdateDeploymentUnauthorized, *UpdateDeploymentNotFound, *UpdateDeploymentInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateDeployment",
		Method:             "PATCH",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateDeploymentOK:
		return v, nil, nil, nil, nil, nil
	case *UpdateDeploymentBadRequest:
		return nil, v, nil, nil, nil, nil
	case *UpdateDeploymentUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *UpdateDeploymentNotFound:
		return nil, nil, nil, v, nil, nil
	case *UpdateDeploymentInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateDeploymentOverride updates deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint update a dedicated servers deployment override in a namespace.
*/
func (a *Client) UpdateDeploymentOverride(params *UpdateDeploymentOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentOverrideOK, *UpdateDeploymentOverrideBadRequest, *UpdateDeploymentOverrideUnauthorized, *UpdateDeploymentOverrideNotFound, *UpdateDeploymentOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateDeploymentOverride",
		Method:             "PATCH",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDeploymentOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateDeploymentOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *UpdateDeploymentOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *UpdateDeploymentOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *UpdateDeploymentOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *UpdateDeploymentOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateOverrideRegionOverride updates region override for deployment override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint update a dedicated servers deployment override in a namespace in a region for deployment overrides.
*/
func (a *Client) UpdateOverrideRegionOverride(params *UpdateOverrideRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOverrideRegionOverrideOK, *UpdateOverrideRegionOverrideBadRequest, *UpdateOverrideRegionOverrideUnauthorized, *UpdateOverrideRegionOverrideNotFound, *UpdateOverrideRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOverrideRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateOverrideRegionOverride",
		Method:             "PATCH",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOverrideRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateOverrideRegionOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *UpdateOverrideRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *UpdateOverrideRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *UpdateOverrideRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *UpdateOverrideRegionOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateRootRegionOverride updates region override

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint update a dedicated servers deployment override in a namespace in a region for root deployment.
*/
func (a *Client) UpdateRootRegionOverride(params *UpdateRootRegionOverrideParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRootRegionOverrideOK, *UpdateRootRegionOverrideBadRequest, *UpdateRootRegionOverrideUnauthorized, *UpdateRootRegionOverrideNotFound, *UpdateRootRegionOverrideInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRootRegionOverrideParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateRootRegionOverride",
		Method:             "PATCH",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRootRegionOverrideReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateRootRegionOverrideOK:
		return v, nil, nil, nil, nil, nil
	case *UpdateRootRegionOverrideBadRequest:
		return nil, v, nil, nil, nil, nil
	case *UpdateRootRegionOverrideUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *UpdateRootRegionOverrideNotFound:
		return nil, nil, nil, v, nil, nil
	case *UpdateRootRegionOverrideInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
