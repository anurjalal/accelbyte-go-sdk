// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// UpdatePrivateCustomAttributesPartiallyReader is a Reader for the UpdatePrivateCustomAttributesPartially structure.
type UpdatePrivateCustomAttributesPartiallyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrivateCustomAttributesPartiallyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePrivateCustomAttributesPartiallyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePrivateCustomAttributesPartiallyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePrivateCustomAttributesPartiallyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdatePrivateCustomAttributesPartiallyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdatePrivateCustomAttributesPartiallyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdatePrivateCustomAttributesPartiallyOK creates a UpdatePrivateCustomAttributesPartiallyOK with default headers values
func NewUpdatePrivateCustomAttributesPartiallyOK() *UpdatePrivateCustomAttributesPartiallyOK {
	return &UpdatePrivateCustomAttributesPartiallyOK{}
}

/*UpdatePrivateCustomAttributesPartiallyOK handles this case with default header values.

  successful operation
*/
type UpdatePrivateCustomAttributesPartiallyOK struct {
	Payload map[string]interface{}
}

func (o *UpdatePrivateCustomAttributesPartiallyOK) Error() string {
	return fmt.Sprintf("[PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes][%d] updatePrivateCustomAttributesPartiallyOK  %+v", 200, o.Payload)
}

func (o *UpdatePrivateCustomAttributesPartiallyOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *UpdatePrivateCustomAttributesPartiallyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePrivateCustomAttributesPartiallyBadRequest creates a UpdatePrivateCustomAttributesPartiallyBadRequest with default headers values
func NewUpdatePrivateCustomAttributesPartiallyBadRequest() *UpdatePrivateCustomAttributesPartiallyBadRequest {
	return &UpdatePrivateCustomAttributesPartiallyBadRequest{}
}

/*UpdatePrivateCustomAttributesPartiallyBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type UpdatePrivateCustomAttributesPartiallyBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *UpdatePrivateCustomAttributesPartiallyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes][%d] updatePrivateCustomAttributesPartiallyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePrivateCustomAttributesPartiallyBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *UpdatePrivateCustomAttributesPartiallyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePrivateCustomAttributesPartiallyUnauthorized creates a UpdatePrivateCustomAttributesPartiallyUnauthorized with default headers values
func NewUpdatePrivateCustomAttributesPartiallyUnauthorized() *UpdatePrivateCustomAttributesPartiallyUnauthorized {
	return &UpdatePrivateCustomAttributesPartiallyUnauthorized{}
}

/*UpdatePrivateCustomAttributesPartiallyUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type UpdatePrivateCustomAttributesPartiallyUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdatePrivateCustomAttributesPartiallyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes][%d] updatePrivateCustomAttributesPartiallyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdatePrivateCustomAttributesPartiallyUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdatePrivateCustomAttributesPartiallyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePrivateCustomAttributesPartiallyForbidden creates a UpdatePrivateCustomAttributesPartiallyForbidden with default headers values
func NewUpdatePrivateCustomAttributesPartiallyForbidden() *UpdatePrivateCustomAttributesPartiallyForbidden {
	return &UpdatePrivateCustomAttributesPartiallyForbidden{}
}

/*UpdatePrivateCustomAttributesPartiallyForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type UpdatePrivateCustomAttributesPartiallyForbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdatePrivateCustomAttributesPartiallyForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes][%d] updatePrivateCustomAttributesPartiallyForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePrivateCustomAttributesPartiallyForbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdatePrivateCustomAttributesPartiallyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePrivateCustomAttributesPartiallyNotFound creates a UpdatePrivateCustomAttributesPartiallyNotFound with default headers values
func NewUpdatePrivateCustomAttributesPartiallyNotFound() *UpdatePrivateCustomAttributesPartiallyNotFound {
	return &UpdatePrivateCustomAttributesPartiallyNotFound{}
}

/*UpdatePrivateCustomAttributesPartiallyNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>11440</td><td>user profile not found</td></tr></table>
*/
type UpdatePrivateCustomAttributesPartiallyNotFound struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdatePrivateCustomAttributesPartiallyNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/admin/namespaces/{namespace}/users/{userId}/profiles/privateCustomAttributes][%d] updatePrivateCustomAttributesPartiallyNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePrivateCustomAttributesPartiallyNotFound) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdatePrivateCustomAttributesPartiallyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
