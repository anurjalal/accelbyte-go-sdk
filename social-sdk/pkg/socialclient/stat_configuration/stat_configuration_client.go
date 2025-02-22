// Code generated by go-swagger; DO NOT EDIT.

package stat_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stat configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stat configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStat(params *CreateStatParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStatCreated, *CreateStatConflict, error)

	CreateStat1(params *CreateStat1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateStat1Created, *CreateStat1Conflict, error)

	DeleteStat(params *DeleteStatParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStatNoContent, *DeleteStatNotFound, error)

	ExportStats(params *ExportStatsParams, authInfo runtime.ClientAuthInfoWriter) (*ExportStatsOK, error)

	GetStat(params *GetStatParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatOK, *GetStatNotFound, error)

	GetStats(params *GetStatsParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatsOK, error)

	ImportStats(params *ImportStatsParams, authInfo runtime.ClientAuthInfoWriter) (*ImportStatsOK, *ImportStatsBadRequest, error)

	QueryStats(params *QueryStatsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryStatsOK, error)

	UpdateStat(params *UpdateStatParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateStatOK, *UpdateStatNotFound, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateStat creates stat

  Create stat.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=1 (CREATE)</li><li><i>Returns</i>: created stat template</li></ul>
*/
func (a *Client) CreateStat(params *CreateStatParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStatCreated, *CreateStatConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStatParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStat",
		Method:             "POST",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStatReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *CreateStatCreated:
		return v, nil, nil
	case *CreateStatConflict:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateStat1 creates stat

  Create stat.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:STAT", action=1 (CREATE)</li><li><i>Returns</i>: created stat template</li></ul>
*/
func (a *Client) CreateStat1(params *CreateStat1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateStat1Created, *CreateStat1Conflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStat1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStat_1",
		Method:             "POST",
		PathPattern:        "/v1/public/namespaces/{namespace}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStat1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *CreateStat1Created:
		return v, nil, nil
	case *CreateStat1Conflict:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteStat deletes stat

  Deletes stat template.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=8 (DELETE)</li></ul>
*/
func (a *Client) DeleteStat(params *DeleteStatParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStatNoContent, *DeleteStatNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStatParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStat",
		Method:             "DELETE",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/{statCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStatReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteStatNoContent:
		return v, nil, nil
	case *DeleteStatNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ExportStats exports all stat configurations

  Export all stat configurations for a given namespace into file  At current, only JSON file is supported.<p>Other detail info:<ul><li><i>*Required permission*: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=2 (READ)</li></ul>
*/
func (a *Client) ExportStats(params *ExportStatsParams, authInfo runtime.ClientAuthInfoWriter) (*ExportStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportStatsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportStats",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ExportStatsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetStat gets stat by stat code

  Get stat by statCode.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=2 (READ)</li><li><i>Returns</i>: stat info</ul>
*/
func (a *Client) GetStat(params *GetStatParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatOK, *GetStatNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStat",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/{statCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetStatOK:
		return v, nil, nil
	case *GetStatNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetStats lists stats

  List stats by pagination.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=2 (READ)</li><li><i>Returns</i>: stats</li></ul>
*/
func (a *Client) GetStats(params *GetStatsParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStats",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetStatsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ImportStats imports stat configurations

  Import stat configurations for a given namespace from file. At current, only JSON file is supported.<p>Other detail info:<ul><li><i>*Required permission*: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=1 (CREATE)</li></ul>
*/
func (a *Client) ImportStats(params *ImportStatsParams, authInfo runtime.ClientAuthInfoWriter) (*ImportStatsOK, *ImportStatsBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportStatsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importStats",
		Method:             "POST",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *ImportStatsOK:
		return v, nil, nil
	case *ImportStatsBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QueryStats queries stats by keyword

  Query stats stats by keyword.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=2 (READ)</li><li><i>Returns<i>: stats</li></ul>
*/
func (a *Client) QueryStats(params *QueryStatsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryStatsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryStats",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QueryStatsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateStat updates stat

  Update stat.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STAT", action=4 (UPDATE)</li><li><i>Returns</i>: updated stat</li></ul>
*/
func (a *Client) UpdateStat(params *UpdateStatParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateStatOK, *UpdateStatNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateStat",
		Method:             "PATCH",
		PathPattern:        "/v1/admin/namespaces/{namespace}/stats/{statCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateStatReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateStatOK:
		return v, nil, nil
	case *UpdateStatNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
