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

// GetCustomAttributesInfo1Reader is a Reader for the GetCustomAttributesInfo1 structure.
type GetCustomAttributesInfo1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomAttributesInfo1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomAttributesInfo1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCustomAttributesInfo1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCustomAttributesInfo1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetCustomAttributesInfo1OK creates a GetCustomAttributesInfo1OK with default headers values
func NewGetCustomAttributesInfo1OK() *GetCustomAttributesInfo1OK {
	return &GetCustomAttributesInfo1OK{}
}

/*GetCustomAttributesInfo1OK handles this case with default header values.

  Successful operation
*/
type GetCustomAttributesInfo1OK struct {
	Payload map[string]interface{}
}

func (o *GetCustomAttributesInfo1OK) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] getCustomAttributesInfo1OK  %+v", 200, o.Payload)
}

func (o *GetCustomAttributesInfo1OK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetCustomAttributesInfo1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomAttributesInfo1Unauthorized creates a GetCustomAttributesInfo1Unauthorized with default headers values
func NewGetCustomAttributesInfo1Unauthorized() *GetCustomAttributesInfo1Unauthorized {
	return &GetCustomAttributesInfo1Unauthorized{}
}

/*GetCustomAttributesInfo1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type GetCustomAttributesInfo1Unauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetCustomAttributesInfo1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] getCustomAttributesInfo1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetCustomAttributesInfo1Unauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetCustomAttributesInfo1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomAttributesInfo1NotFound creates a GetCustomAttributesInfo1NotFound with default headers values
func NewGetCustomAttributesInfo1NotFound() *GetCustomAttributesInfo1NotFound {
	return &GetCustomAttributesInfo1NotFound{}
}

/*GetCustomAttributesInfo1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>11440</td><td>user profile not found</td></tr></table>
*/
type GetCustomAttributesInfo1NotFound struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetCustomAttributesInfo1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] getCustomAttributesInfo1NotFound  %+v", 404, o.Payload)
}

func (o *GetCustomAttributesInfo1NotFound) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetCustomAttributesInfo1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
