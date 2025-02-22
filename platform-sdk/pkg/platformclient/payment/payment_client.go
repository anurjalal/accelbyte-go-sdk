// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChargePaymentOrder(params *ChargePaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*ChargePaymentOrderOK, *ChargePaymentOrderBadRequest, *ChargePaymentOrderNotFound, *ChargePaymentOrderConflict, error)

	CreateUserPaymentOrder(params *CreateUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserPaymentOrderCreated, *CreateUserPaymentOrderBadRequest, *CreateUserPaymentOrderForbidden, *CreateUserPaymentOrderNotFound, *CreateUserPaymentOrderConflict, *CreateUserPaymentOrderUnprocessableEntity, error)

	GetPaymentOrder(params *GetPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOrderOK, *GetPaymentOrderNotFound, error)

	GetPaymentOrderChargeStatus(params *GetPaymentOrderChargeStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOrderChargeStatusOK, *GetPaymentOrderChargeStatusNotFound, error)

	ListExtOrderNoByExtTxID(params *ListExtOrderNoByExtTxIDParams, authInfo runtime.ClientAuthInfoWriter) (*ListExtOrderNoByExtTxIDOK, error)

	QueryPaymentNotifications(params *QueryPaymentNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPaymentNotificationsOK, error)

	QueryPaymentOrders(params *QueryPaymentOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPaymentOrdersOK, error)

	RefundUserPaymentOrder(params *RefundUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*RefundUserPaymentOrderOK, *RefundUserPaymentOrderNotFound, *RefundUserPaymentOrderConflict, *RefundUserPaymentOrderUnprocessableEntity, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChargePaymentOrder charges payment order without payment flow

  <b>[TEST FACILITY ONLY]</b> Charge payment order without payment flow for unpaid payment order, usually for test usage to simulate real currency payment process.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT", action=4 (UPDATE)</li><li><i>Returns</i>: payment order instance</li></ul>
*/
func (a *Client) ChargePaymentOrder(params *ChargePaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*ChargePaymentOrderOK, *ChargePaymentOrderBadRequest, *ChargePaymentOrderNotFound, *ChargePaymentOrderConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargePaymentOrderParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "chargePaymentOrder",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChargePaymentOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ChargePaymentOrderOK:
		return v, nil, nil, nil, nil
	case *ChargePaymentOrderBadRequest:
		return nil, v, nil, nil, nil
	case *ChargePaymentOrderNotFound:
		return nil, nil, v, nil, nil
	case *ChargePaymentOrderConflict:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateUserPaymentOrder creates payment order

  <b>[SERVICE COMMUNICATION ONLY]</b> This API is used to create payment order from justice service. The result contains the payment station url.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:PAYMENT", action=1 (CREATE)</li><li>It will be forbidden while the user is banned: PAYMENT_INITIATE or ORDER_AND_PAYMENT</li><li><i>Returns</i>: created order</li></ul>
*/
func (a *Client) CreateUserPaymentOrder(params *CreateUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserPaymentOrderCreated, *CreateUserPaymentOrderBadRequest, *CreateUserPaymentOrderForbidden, *CreateUserPaymentOrderNotFound, *CreateUserPaymentOrderConflict, *CreateUserPaymentOrderUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserPaymentOrderParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUserPaymentOrder",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/payment/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserPaymentOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateUserPaymentOrderCreated:
		return v, nil, nil, nil, nil, nil, nil
	case *CreateUserPaymentOrderBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *CreateUserPaymentOrderForbidden:
		return nil, nil, v, nil, nil, nil, nil
	case *CreateUserPaymentOrderNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *CreateUserPaymentOrderConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *CreateUserPaymentOrderUnprocessableEntity:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetPaymentOrder gets payment order

  Get payment order by paymentOrderNo.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT", action=2 (READ)</li><li><i>Returns</i>: payment order instance</li></ul>
*/
func (a *Client) GetPaymentOrder(params *GetPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOrderOK, *GetPaymentOrderNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentOrderParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentOrder",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetPaymentOrderOK:
		return v, nil, nil
	case *GetPaymentOrderNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetPaymentOrderChargeStatus gets payment order charge status

  Get payment order charge status.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT", action=2 (READ)</li><li><i>Returns</i>: payment order charge status</li></ul>
*/
func (a *Client) GetPaymentOrderChargeStatus(params *GetPaymentOrderChargeStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOrderChargeStatusOK, *GetPaymentOrderChargeStatusNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentOrderChargeStatusParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentOrderChargeStatus",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentOrderChargeStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetPaymentOrderChargeStatusOK:
		return v, nil, nil
	case *GetPaymentOrderChargeStatusNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListExtOrderNoByExtTxID lists external order no by external transaction id

  List external order No by external transaction id.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT", action=2 (READ)</li><li><i>Returns</i>: payment orders</li></ul>
*/
func (a *Client) ListExtOrderNoByExtTxID(params *ListExtOrderNoByExtTxIDParams, authInfo runtime.ClientAuthInfoWriter) (*ListExtOrderNoByExtTxIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExtOrderNoByExtTxIDParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listExtOrderNoByExtTxId",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/payment/orders/byExtTxId",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListExtOrderNoByExtTxIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListExtOrderNoByExtTxIDOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QueryPaymentNotifications queries payment notifications

  Query payment notifications.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT:NOTIFICATION", action=2 (READ)</li><li><i>Returns</i>: Payment notifications</li></ul>
*/
func (a *Client) QueryPaymentNotifications(params *QueryPaymentNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPaymentNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryPaymentNotificationsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryPaymentNotifications",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/payment/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryPaymentNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QueryPaymentNotificationsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QueryPaymentOrders queries payment orders

  Query payment orders.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:PAYMENT", action=2 (READ)</li><li><i>Returns</i>: query payment orders</li></ul>
*/
func (a *Client) QueryPaymentOrders(params *QueryPaymentOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPaymentOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryPaymentOrdersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryPaymentOrders",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/payment/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryPaymentOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QueryPaymentOrdersOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RefundUserPaymentOrder refunds payment order

  <b>[SERVICE COMMUNICATION ONLY]</b> This API is used to refund order by paymentOrderNo from justice service.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:PAYMENT", action=4 (UPDATE)</li></ul>
*/
func (a *Client) RefundUserPaymentOrder(params *RefundUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter) (*RefundUserPaymentOrderOK, *RefundUserPaymentOrderNotFound, *RefundUserPaymentOrderConflict, *RefundUserPaymentOrderUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefundUserPaymentOrderParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "refundUserPaymentOrder",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/payment/orders/{paymentOrderNo}/refund",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefundUserPaymentOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RefundUserPaymentOrderOK:
		return v, nil, nil, nil, nil
	case *RefundUserPaymentOrderNotFound:
		return nil, v, nil, nil, nil
	case *RefundUserPaymentOrderConflict:
		return nil, nil, v, nil, nil
	case *RefundUserPaymentOrderUnprocessableEntity:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
