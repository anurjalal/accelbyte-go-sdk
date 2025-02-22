// Code generated by go-swagger; DO NOT EDIT.

package bans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// AdminGetBansTypeWithNamespaceV3Reader is a Reader for the AdminGetBansTypeWithNamespaceV3 structure.
type AdminGetBansTypeWithNamespaceV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetBansTypeWithNamespaceV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetBansTypeWithNamespaceV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetBansTypeWithNamespaceV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminGetBansTypeWithNamespaceV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/admin/namespaces/{namespace}/bantypes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminGetBansTypeWithNamespaceV3OK creates a AdminGetBansTypeWithNamespaceV3OK with default headers values
func NewAdminGetBansTypeWithNamespaceV3OK() *AdminGetBansTypeWithNamespaceV3OK {
	return &AdminGetBansTypeWithNamespaceV3OK{}
}

/*AdminGetBansTypeWithNamespaceV3OK handles this case with default header values.

  OK
*/
type AdminGetBansTypeWithNamespaceV3OK struct {
	Payload *iamclientmodels.AccountcommonBansV3
}

func (o *AdminGetBansTypeWithNamespaceV3OK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/bantypes][%d] adminGetBansTypeWithNamespaceV3OK  %+v", 200, o.Payload)
}

func (o *AdminGetBansTypeWithNamespaceV3OK) GetPayload() *iamclientmodels.AccountcommonBansV3 {
	return o.Payload
}

func (o *AdminGetBansTypeWithNamespaceV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.AccountcommonBansV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBansTypeWithNamespaceV3Unauthorized creates a AdminGetBansTypeWithNamespaceV3Unauthorized with default headers values
func NewAdminGetBansTypeWithNamespaceV3Unauthorized() *AdminGetBansTypeWithNamespaceV3Unauthorized {
	return &AdminGetBansTypeWithNamespaceV3Unauthorized{}
}

/*AdminGetBansTypeWithNamespaceV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type AdminGetBansTypeWithNamespaceV3Unauthorized struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminGetBansTypeWithNamespaceV3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/bantypes][%d] adminGetBansTypeWithNamespaceV3Unauthorized  %+v", 401, o.Payload)
}

func (o *AdminGetBansTypeWithNamespaceV3Unauthorized) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminGetBansTypeWithNamespaceV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBansTypeWithNamespaceV3Forbidden creates a AdminGetBansTypeWithNamespaceV3Forbidden with default headers values
func NewAdminGetBansTypeWithNamespaceV3Forbidden() *AdminGetBansTypeWithNamespaceV3Forbidden {
	return &AdminGetBansTypeWithNamespaceV3Forbidden{}
}

/*AdminGetBansTypeWithNamespaceV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type AdminGetBansTypeWithNamespaceV3Forbidden struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminGetBansTypeWithNamespaceV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/bantypes][%d] adminGetBansTypeWithNamespaceV3Forbidden  %+v", 403, o.Payload)
}

func (o *AdminGetBansTypeWithNamespaceV3Forbidden) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminGetBansTypeWithNamespaceV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
