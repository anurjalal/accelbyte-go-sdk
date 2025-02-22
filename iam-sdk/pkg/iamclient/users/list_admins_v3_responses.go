// Code generated by go-swagger; DO NOT EDIT.

package users

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

// ListAdminsV3Reader is a Reader for the ListAdminsV3 structure.
type ListAdminsV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAdminsV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAdminsV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAdminsV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListAdminsV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListAdminsV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/admin/namespaces/{namespace}/admins returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListAdminsV3OK creates a ListAdminsV3OK with default headers values
func NewListAdminsV3OK() *ListAdminsV3OK {
	return &ListAdminsV3OK{}
}

/*ListAdminsV3OK handles this case with default header values.

  Operation succeeded
*/
type ListAdminsV3OK struct {
	Payload *iamclientmodels.ModelGetUsersResponseWithPaginationV3
}

func (o *ListAdminsV3OK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/admins][%d] listAdminsV3OK  %+v", 200, o.Payload)
}

func (o *ListAdminsV3OK) GetPayload() *iamclientmodels.ModelGetUsersResponseWithPaginationV3 {
	return o.Payload
}

func (o *ListAdminsV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelGetUsersResponseWithPaginationV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAdminsV3Unauthorized creates a ListAdminsV3Unauthorized with default headers values
func NewListAdminsV3Unauthorized() *ListAdminsV3Unauthorized {
	return &ListAdminsV3Unauthorized{}
}

/*ListAdminsV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type ListAdminsV3Unauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *ListAdminsV3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/admins][%d] listAdminsV3Unauthorized  %+v", 401, o.Payload)
}

func (o *ListAdminsV3Unauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *ListAdminsV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAdminsV3Forbidden creates a ListAdminsV3Forbidden with default headers values
func NewListAdminsV3Forbidden() *ListAdminsV3Forbidden {
	return &ListAdminsV3Forbidden{}
}

/*ListAdminsV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20003</td><td>forbidden access</td></tr></table>
*/
type ListAdminsV3Forbidden struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *ListAdminsV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/admins][%d] listAdminsV3Forbidden  %+v", 403, o.Payload)
}

func (o *ListAdminsV3Forbidden) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *ListAdminsV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAdminsV3InternalServerError creates a ListAdminsV3InternalServerError with default headers values
func NewListAdminsV3InternalServerError() *ListAdminsV3InternalServerError {
	return &ListAdminsV3InternalServerError{}
}

/*ListAdminsV3InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type ListAdminsV3InternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *ListAdminsV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/admins][%d] listAdminsV3InternalServerError  %+v", 500, o.Payload)
}

func (o *ListAdminsV3InternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *ListAdminsV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
