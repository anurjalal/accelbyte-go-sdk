// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// GrantUserEntitlementReader is a Reader for the GrantUserEntitlement structure.
type GrantUserEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantUserEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGrantUserEntitlementCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGrantUserEntitlementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGrantUserEntitlementUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /admin/namespaces/{namespace}/users/{userId}/entitlements returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGrantUserEntitlementCreated creates a GrantUserEntitlementCreated with default headers values
func NewGrantUserEntitlementCreated() *GrantUserEntitlementCreated {
	return &GrantUserEntitlementCreated{}
}

/*GrantUserEntitlementCreated handles this case with default header values.

  successful operation
*/
type GrantUserEntitlementCreated struct {
	Payload []*platformclientmodels.StackableEntitlementInfo
}

func (o *GrantUserEntitlementCreated) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/entitlements][%d] grantUserEntitlementCreated  %+v", 201, o.Payload)
}

func (o *GrantUserEntitlementCreated) GetPayload() []*platformclientmodels.StackableEntitlementInfo {
	return o.Payload
}

func (o *GrantUserEntitlementCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserEntitlementNotFound creates a GrantUserEntitlementNotFound with default headers values
func NewGrantUserEntitlementNotFound() *GrantUserEntitlementNotFound {
	return &GrantUserEntitlementNotFound{}
}

/*GrantUserEntitlementNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30341</td><td>Item [{itemId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type GrantUserEntitlementNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GrantUserEntitlementNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/entitlements][%d] grantUserEntitlementNotFound  %+v", 404, o.Payload)
}

func (o *GrantUserEntitlementNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GrantUserEntitlementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserEntitlementUnprocessableEntity creates a GrantUserEntitlementUnprocessableEntity with default headers values
func NewGrantUserEntitlementUnprocessableEntity() *GrantUserEntitlementUnprocessableEntity {
	return &GrantUserEntitlementUnprocessableEntity{}
}

/*GrantUserEntitlementUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type GrantUserEntitlementUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *GrantUserEntitlementUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/entitlements][%d] grantUserEntitlementUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GrantUserEntitlementUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *GrantUserEntitlementUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
