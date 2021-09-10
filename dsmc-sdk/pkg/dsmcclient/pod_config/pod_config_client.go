// Code generated by go-swagger; DO NOT EDIT.

package pod_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pod config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pod config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePodConfig(params *CreatePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePodConfigCreated, *CreatePodConfigBadRequest, *CreatePodConfigUnauthorized, *CreatePodConfigConflict, *CreatePodConfigInternalServerError, error)

	DeletePodConfig(params *DeletePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePodConfigNoContent, *DeletePodConfigBadRequest, *DeletePodConfigUnauthorized, *DeletePodConfigNotFound, *DeletePodConfigConflict, *DeletePodConfigInternalServerError, error)

	GetAllPodConfig(params *GetAllPodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPodConfigOK, *GetAllPodConfigBadRequest, *GetAllPodConfigUnauthorized, *GetAllPodConfigInternalServerError, error)

	GetPodConfig(params *GetPodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetPodConfigOK, *GetPodConfigBadRequest, *GetPodConfigUnauthorized, *GetPodConfigNotFound, *GetPodConfigInternalServerError, error)

	UpdatePodConfig(params *UpdatePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePodConfigOK, *UpdatePodConfigBadRequest, *UpdatePodConfigUnauthorized, *UpdatePodConfigNotFound, *UpdatePodConfigConflict, *UpdatePodConfigInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePodConfig creates pod config

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint create a dedicated servers pod config in a namespace.
*/
func (a *Client) CreatePodConfig(params *CreatePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePodConfigCreated, *CreatePodConfigBadRequest, *CreatePodConfigUnauthorized, *CreatePodConfigConflict, *CreatePodConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePodConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePodConfig",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/pods/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePodConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreatePodConfigCreated:
		return v, nil, nil, nil, nil, nil
	case *CreatePodConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *CreatePodConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *CreatePodConfigConflict:
		return nil, nil, nil, v, nil, nil
	case *CreatePodConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeletePodConfig deletes pod config

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [DELETE]

Required scope: social

This endpoint delete a dedicated server pod config in a namespace
*/
func (a *Client) DeletePodConfig(params *DeletePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePodConfigNoContent, *DeletePodConfigBadRequest, *DeletePodConfigUnauthorized, *DeletePodConfigNotFound, *DeletePodConfigConflict, *DeletePodConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePodConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePodConfig",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/pods/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePodConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeletePodConfigNoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *DeletePodConfigBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeletePodConfigUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeletePodConfigNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *DeletePodConfigConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *DeletePodConfigInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetAllPodConfig gets all pod configs

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get a all pod configs in a namespace
*/
func (a *Client) GetAllPodConfig(params *GetAllPodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPodConfigOK, *GetAllPodConfigBadRequest, *GetAllPodConfigUnauthorized, *GetAllPodConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPodConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAllPodConfig",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/pods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllPodConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetAllPodConfigOK:
		return v, nil, nil, nil, nil
	case *GetAllPodConfigBadRequest:
		return nil, v, nil, nil, nil
	case *GetAllPodConfigUnauthorized:
		return nil, nil, v, nil, nil
	case *GetAllPodConfigInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetPodConfig gets pod config

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get a dedicated server pod config in a namespace
*/
func (a *Client) GetPodConfig(params *GetPodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetPodConfigOK, *GetPodConfigBadRequest, *GetPodConfigUnauthorized, *GetPodConfigNotFound, *GetPodConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPodConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPodConfig",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/pods/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPodConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetPodConfigOK:
		return v, nil, nil, nil, nil, nil
	case *GetPodConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetPodConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetPodConfigNotFound:
		return nil, nil, nil, v, nil, nil
	case *GetPodConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdatePodConfig updates pod config

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint update a dedicated servers pod config in a namespace.
*/
func (a *Client) UpdatePodConfig(params *UpdatePodConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePodConfigOK, *UpdatePodConfigBadRequest, *UpdatePodConfigUnauthorized, *UpdatePodConfigNotFound, *UpdatePodConfigConflict, *UpdatePodConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePodConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdatePodConfig",
		Method:             "PATCH",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/configs/pods/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePodConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdatePodConfigOK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdatePodConfigBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdatePodConfigUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdatePodConfigNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdatePodConfigConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdatePodConfigInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
