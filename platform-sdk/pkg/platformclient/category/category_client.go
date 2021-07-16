// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new category API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for category API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCategory(params *CreateCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCategoryCreated, *CreateCategoryBadRequest, *CreateCategoryNotFound, *CreateCategoryConflict, *CreateCategoryUnprocessableEntity, error)

	DeleteCategory(params *DeleteCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCategoryOK, *DeleteCategoryNotFound, *DeleteCategoryConflict, error)

	DownloadCategories(params *DownloadCategoriesParams) (*DownloadCategoriesOK, *DownloadCategoriesNotFound, error)

	GetCategory(params *GetCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoryOK, *GetCategoryNotFound, error)

	GetChildCategories(params *GetChildCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetChildCategoriesOK, error)

	GetDescendantCategories(params *GetDescendantCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDescendantCategoriesOK, error)

	GetRootCategories(params *GetRootCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetRootCategoriesOK, error)

	ListCategoriesBasic(params *ListCategoriesBasicParams, authInfo runtime.ClientAuthInfoWriter) (*ListCategoriesBasicOK, error)

	PublicGetCategory(params *PublicGetCategoryParams) (*PublicGetCategoryOK, *PublicGetCategoryNotFound, error)

	PublicGetChildCategories(params *PublicGetChildCategoriesParams) (*PublicGetChildCategoriesOK, error)

	PublicGetDescendantCategories(params *PublicGetDescendantCategoriesParams) (*PublicGetDescendantCategoriesOK, error)

	PublicGetRootCategories(params *PublicGetRootCategoriesParams) (*PublicGetRootCategoriesOK, error)

	UpdateCategory(params *UpdateCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCategoryOK, *UpdateCategoryBadRequest, *UpdateCategoryNotFound, *UpdateCategoryConflict, *UpdateCategoryUnprocessableEntity, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCategory creates category

  This API is used to create category.<p>A category is a path separated by "/". A category also has localized display names. Example:<p><pre><code>{
	"categoryPath": "/games",
	"localizationDisplayNames": \{"en" : "Games"}
}</code></pre>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=1 (CREATE)</li><li><i>Returns</i>: created category data</li></ul>
*/
func (a *Client) CreateCategory(params *CreateCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCategoryCreated, *CreateCategoryBadRequest, *CreateCategoryNotFound, *CreateCategoryConflict, *CreateCategoryUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCategoryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCategory",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateCategoryCreated:
		return v, nil, nil, nil, nil, nil
	case *CreateCategoryBadRequest:
		return nil, v, nil, nil, nil, nil
	case *CreateCategoryNotFound:
		return nil, nil, v, nil, nil, nil
	case *CreateCategoryConflict:
		return nil, nil, nil, v, nil, nil
	case *CreateCategoryUnprocessableEntity:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteCategory deletes category

  This API is used to delete category by category path. <p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=8 (DELETE)</li><li><i>Returns</i>: the deleted category data</li></ul>
*/
func (a *Client) DeleteCategory(params *DeleteCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCategoryOK, *DeleteCategoryNotFound, *DeleteCategoryConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCategoryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCategory",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/categories/{categoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteCategoryOK:
		return v, nil, nil, nil
	case *DeleteCategoryNotFound:
		return nil, v, nil, nil
	case *DeleteCategoryConflict:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DownloadCategories downloads store s structured categories

  This API is used to download store's structured categories.<p>Other detail info: <ul><li><i>Optional permission</i>: resource="SANDBOX", action=1(CREATE) (user with this permission can view draft store content)</li><li><i>Returns</i>: structured categories</li></ul>
*/
func (a *Client) DownloadCategories(params *DownloadCategoriesParams) (*DownloadCategoriesOK, *DownloadCategoriesNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadCategories",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/categories/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *DownloadCategoriesOK:
		return v, nil, nil
	case *DownloadCategoriesNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetCategory gets category

  This API is used to get category by category path.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=2 (READ)</li><li><i>Returns</i>: category data</li></ul>
*/
func (a *Client) GetCategory(params *GetCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoryOK, *GetCategoryNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCategory",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/categories/{categoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetCategoryOK:
		return v, nil, nil
	case *GetCategoryNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetChildCategories gets child categories

  This API is used to get child categories by category path.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=2 (READ)</li><li><i>Returns</i>: list of child categories data</li></ul>
*/
func (a *Client) GetChildCategories(params *GetChildCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetChildCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChildCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChildCategories",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/categories/{categoryPath}/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChildCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetChildCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetDescendantCategories gets descendant categories

  This API is used to get descendant categories by category path.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=2 (READ)</li><li><i>Returns</i>: list of descendant categories data</li></ul>
*/
func (a *Client) GetDescendantCategories(params *GetDescendantCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDescendantCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDescendantCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDescendantCategories",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/categories/{categoryPath}/descendants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDescendantCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetDescendantCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetRootCategories gets root categories

  This API is used to get root categories.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=2 (READ)</li><li><i>Returns</i>: root category data</li></ul>
*/
func (a *Client) GetRootCategories(params *GetRootCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetRootCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRootCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRootCategories",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRootCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetRootCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListCategoriesBasic lists categories basic info

  This API is used to list all categories' basic info of a store ordered by category path.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=2 (READ)</li><li><i>Returns</i>: list of categories' paths</li></ul>
*/
func (a *Client) ListCategoriesBasic(params *ListCategoriesBasicParams, authInfo runtime.ClientAuthInfoWriter) (*ListCategoriesBasicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCategoriesBasicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCategoriesBasic",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/categories/basic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCategoriesBasicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListCategoriesBasicOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetCategory gets category

  This API is used to get category by category path.<p>Other detail info: <ul><li><i>Optional permission</i>: resource="SANDBOX", action=1 (CREATE)(user with this permission can view draft store category)</li><li><i>Returns</i>: category data</li></ul>
*/
func (a *Client) PublicGetCategory(params *PublicGetCategoryParams) (*PublicGetCategoryOK, *PublicGetCategoryNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetCategoryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetCategory",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/categories/{categoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetCategoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetCategoryOK:
		return v, nil, nil
	case *PublicGetCategoryNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetChildCategories gets child categories

  This API is used to get child categories by category path.<p>Other detail info: <ul><li><i>Optional permission</i>: resource="SANDBOX", action=1(CREATE) (user with this permission can view draft store category)</li><li><i>Returns</i>: list of child categories data</li></ul>
*/
func (a *Client) PublicGetChildCategories(params *PublicGetChildCategoriesParams) (*PublicGetChildCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetChildCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetChildCategories",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/categories/{categoryPath}/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetChildCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetChildCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetDescendantCategories gets descendant categories

  This API is used to get descendant categories by category path.<p>Other detail info: <ul><li><i>Optional permission</i>: resource="SANDBOX", action=1(CREATE) (user with this permission can view draft store category)</li><li><i>Returns</i>: list of descendant categories data</li></ul>
*/
func (a *Client) PublicGetDescendantCategories(params *PublicGetDescendantCategoriesParams) (*PublicGetDescendantCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetDescendantCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetDescendantCategories",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/categories/{categoryPath}/descendants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetDescendantCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetDescendantCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetRootCategories gets root categories

  This API is used to get root categories.<p>Other detail info: <ul><li><i>Optional permission</i>: resource="SANDBOX", action=1(CREATE) (user with this permission can view draft store category)</li><li><i>Returns</i>: root category data</li></ul>
*/
func (a *Client) PublicGetRootCategories(params *PublicGetRootCategoriesParams) (*PublicGetRootCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetRootCategoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetRootCategories",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetRootCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetRootCategoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateCategory updates category

  This API is used to update category. <p> The category update data is a category object, example as:<pre><code>{
	"storeId": "store-id",
	"localizationDisplayNames": {"en" : "Games"}
}</code></pre>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CATEGORY", action=4 (UPDATE)</li><li><i>Returns</i>: the updated category data</li></ul>
*/
func (a *Client) UpdateCategory(params *UpdateCategoryParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCategoryOK, *UpdateCategoryBadRequest, *UpdateCategoryNotFound, *UpdateCategoryConflict, *UpdateCategoryUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCategoryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCategory",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/categories/{categoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCategoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateCategoryOK:
		return v, nil, nil, nil, nil, nil
	case *UpdateCategoryBadRequest:
		return nil, v, nil, nil, nil, nil
	case *UpdateCategoryNotFound:
		return nil, nil, v, nil, nil, nil
	case *UpdateCategoryConflict:
		return nil, nil, nil, v, nil, nil
	case *UpdateCategoryUnprocessableEntity:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
