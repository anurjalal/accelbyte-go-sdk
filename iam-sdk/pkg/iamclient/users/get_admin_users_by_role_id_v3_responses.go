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

// GetAdminUsersByRoleIDV3Reader is a Reader for the GetAdminUsersByRoleIDV3 structure.
type GetAdminUsersByRoleIDV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminUsersByRoleIDV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminUsersByRoleIDV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdminUsersByRoleIDV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAdminUsersByRoleIDV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAdminUsersByRoleIDV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAdminUsersByRoleIDV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAdminUsersByRoleIDV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetAdminUsersByRoleIDV3OK creates a GetAdminUsersByRoleIDV3OK with default headers values
func NewGetAdminUsersByRoleIDV3OK() *GetAdminUsersByRoleIDV3OK {
	return &GetAdminUsersByRoleIDV3OK{}
}

/*GetAdminUsersByRoleIDV3OK handles this case with default header values.

  OK
*/
type GetAdminUsersByRoleIDV3OK struct {
	Payload *iamclientmodels.ModelGetUsersResponseWithPaginationV3
}

func (o *GetAdminUsersByRoleIDV3OK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3OK  %+v", 200, o.Payload)
}

func (o *GetAdminUsersByRoleIDV3OK) GetPayload() *iamclientmodels.ModelGetUsersResponseWithPaginationV3 {
	return o.Payload
}

func (o *GetAdminUsersByRoleIDV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelGetUsersResponseWithPaginationV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersByRoleIDV3BadRequest creates a GetAdminUsersByRoleIDV3BadRequest with default headers values
func NewGetAdminUsersByRoleIDV3BadRequest() *GetAdminUsersByRoleIDV3BadRequest {
	return &GetAdminUsersByRoleIDV3BadRequest{}
}

/*GetAdminUsersByRoleIDV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>10157</td><td>specified role is not admin role</td></tr></table>
*/
type GetAdminUsersByRoleIDV3BadRequest struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *GetAdminUsersByRoleIDV3BadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3BadRequest  %+v", 400, o.Payload)
}

func (o *GetAdminUsersByRoleIDV3BadRequest) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetAdminUsersByRoleIDV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersByRoleIDV3Unauthorized creates a GetAdminUsersByRoleIDV3Unauthorized with default headers values
func NewGetAdminUsersByRoleIDV3Unauthorized() *GetAdminUsersByRoleIDV3Unauthorized {
	return &GetAdminUsersByRoleIDV3Unauthorized{}
}

/*GetAdminUsersByRoleIDV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type GetAdminUsersByRoleIDV3Unauthorized struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *GetAdminUsersByRoleIDV3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3Unauthorized  %+v", 401, o.Payload)
}

func (o *GetAdminUsersByRoleIDV3Unauthorized) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetAdminUsersByRoleIDV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersByRoleIDV3Forbidden creates a GetAdminUsersByRoleIDV3Forbidden with default headers values
func NewGetAdminUsersByRoleIDV3Forbidden() *GetAdminUsersByRoleIDV3Forbidden {
	return &GetAdminUsersByRoleIDV3Forbidden{}
}

/*GetAdminUsersByRoleIDV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type GetAdminUsersByRoleIDV3Forbidden struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *GetAdminUsersByRoleIDV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3Forbidden  %+v", 403, o.Payload)
}

func (o *GetAdminUsersByRoleIDV3Forbidden) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetAdminUsersByRoleIDV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersByRoleIDV3NotFound creates a GetAdminUsersByRoleIDV3NotFound with default headers values
func NewGetAdminUsersByRoleIDV3NotFound() *GetAdminUsersByRoleIDV3NotFound {
	return &GetAdminUsersByRoleIDV3NotFound{}
}

/*GetAdminUsersByRoleIDV3NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10156</td><td>role not found</td></tr></table>
*/
type GetAdminUsersByRoleIDV3NotFound struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *GetAdminUsersByRoleIDV3NotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3NotFound  %+v", 404, o.Payload)
}

func (o *GetAdminUsersByRoleIDV3NotFound) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetAdminUsersByRoleIDV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminUsersByRoleIDV3InternalServerError creates a GetAdminUsersByRoleIDV3InternalServerError with default headers values
func NewGetAdminUsersByRoleIDV3InternalServerError() *GetAdminUsersByRoleIDV3InternalServerError {
	return &GetAdminUsersByRoleIDV3InternalServerError{}
}

/*GetAdminUsersByRoleIDV3InternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetAdminUsersByRoleIDV3InternalServerError struct {
}

func (o *GetAdminUsersByRoleIDV3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/roles/{roleId}/users][%d] getAdminUsersByRoleIdV3InternalServerError ", 500)
}

func (o *GetAdminUsersByRoleIDV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
