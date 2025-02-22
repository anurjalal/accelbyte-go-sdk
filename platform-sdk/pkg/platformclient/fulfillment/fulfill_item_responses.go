// Code generated by go-swagger; DO NOT EDIT.

package fulfillment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// FulfillItemReader is a Reader for the FulfillItem structure.
type FulfillItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FulfillItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFulfillItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFulfillItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFulfillItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewFulfillItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /admin/namespaces/{namespace}/users/{userId}/fulfillment returns an error %d: %s", response.Code(), string(data))
	}
}

// NewFulfillItemOK creates a FulfillItemOK with default headers values
func NewFulfillItemOK() *FulfillItemOK {
	return &FulfillItemOK{}
}

/*FulfillItemOK handles this case with default header values.

  successful operation
*/
type FulfillItemOK struct {
	Payload *platformclientmodels.FulfillmentResult
}

func (o *FulfillItemOK) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/fulfillment][%d] fulfillItemOK  %+v", 200, o.Payload)
}

func (o *FulfillItemOK) GetPayload() *platformclientmodels.FulfillmentResult {
	return o.Payload
}

func (o *FulfillItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FulfillmentResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFulfillItemBadRequest creates a FulfillItemBadRequest with default headers values
func NewFulfillItemBadRequest() *FulfillItemBadRequest {
	return &FulfillItemBadRequest{}
}

/*FulfillItemBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>35121</td><td>Transaction amount [{actualAmount}] exceed max amount [{maxAmount}] per day</td></tr><tr><td>35122</td><td>Transaction amount [{actualAmount}] exceed max amount [{maxAmount}] per transaction</td></tr><tr><td>35123</td><td>Wallet [{walletId}] is inactive</td></tr><tr><td>35125</td><td>Balance exceed max balance [{maxAmount}]</td></tr><tr><td>38121</td><td>Duplicate permanent item exists</td></tr><tr><td>38122</td><td>Subscription endDate required</td></tr></table>
*/
type FulfillItemBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *FulfillItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/fulfillment][%d] fulfillItemBadRequest  %+v", 400, o.Payload)
}

func (o *FulfillItemBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *FulfillItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFulfillItemNotFound creates a FulfillItemNotFound with default headers values
func NewFulfillItemNotFound() *FulfillItemNotFound {
	return &FulfillItemNotFound{}
}

/*FulfillItemNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30341</td><td>Item [{itemId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type FulfillItemNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *FulfillItemNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/fulfillment][%d] fulfillItemNotFound  %+v", 404, o.Payload)
}

func (o *FulfillItemNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *FulfillItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFulfillItemConflict creates a FulfillItemConflict with default headers values
func NewFulfillItemConflict() *FulfillItemConflict {
	return &FulfillItemConflict{}
}

/*FulfillItemConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20006</td><td>optimistic lock</td></tr></table>
*/
type FulfillItemConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *FulfillItemConflict) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/fulfillment][%d] fulfillItemConflict  %+v", 409, o.Payload)
}

func (o *FulfillItemConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *FulfillItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
