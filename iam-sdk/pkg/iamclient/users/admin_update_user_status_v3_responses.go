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

// AdminUpdateUserStatusV3Reader is a Reader for the AdminUpdateUserStatusV3 structure.
type AdminUpdateUserStatusV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserStatusV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminUpdateUserStatusV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserStatusV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminUpdateUserStatusV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminUpdateUserStatusV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminUpdateUserStatusV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminUpdateUserStatusV3NoContent creates a AdminUpdateUserStatusV3NoContent with default headers values
func NewAdminUpdateUserStatusV3NoContent() *AdminUpdateUserStatusV3NoContent {
	return &AdminUpdateUserStatusV3NoContent{}
}

/*AdminUpdateUserStatusV3NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminUpdateUserStatusV3NoContent struct {
}

func (o *AdminUpdateUserStatusV3NoContent) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status][%d] adminUpdateUserStatusV3NoContent ", 204)
}

func (o *AdminUpdateUserStatusV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserStatusV3BadRequest creates a AdminUpdateUserStatusV3BadRequest with default headers values
func NewAdminUpdateUserStatusV3BadRequest() *AdminUpdateUserStatusV3BadRequest {
	return &AdminUpdateUserStatusV3BadRequest{}
}

/*AdminUpdateUserStatusV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type AdminUpdateUserStatusV3BadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminUpdateUserStatusV3BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status][%d] adminUpdateUserStatusV3BadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateUserStatusV3BadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminUpdateUserStatusV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserStatusV3Unauthorized creates a AdminUpdateUserStatusV3Unauthorized with default headers values
func NewAdminUpdateUserStatusV3Unauthorized() *AdminUpdateUserStatusV3Unauthorized {
	return &AdminUpdateUserStatusV3Unauthorized{}
}

/*AdminUpdateUserStatusV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type AdminUpdateUserStatusV3Unauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminUpdateUserStatusV3Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status][%d] adminUpdateUserStatusV3Unauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateUserStatusV3Unauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminUpdateUserStatusV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserStatusV3Forbidden creates a AdminUpdateUserStatusV3Forbidden with default headers values
func NewAdminUpdateUserStatusV3Forbidden() *AdminUpdateUserStatusV3Forbidden {
	return &AdminUpdateUserStatusV3Forbidden{}
}

/*AdminUpdateUserStatusV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type AdminUpdateUserStatusV3Forbidden struct {
}

func (o *AdminUpdateUserStatusV3Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status][%d] adminUpdateUserStatusV3Forbidden ", 403)
}

func (o *AdminUpdateUserStatusV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserStatusV3NotFound creates a AdminUpdateUserStatusV3NotFound with default headers values
func NewAdminUpdateUserStatusV3NotFound() *AdminUpdateUserStatusV3NotFound {
	return &AdminUpdateUserStatusV3NotFound{}
}

/*AdminUpdateUserStatusV3NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type AdminUpdateUserStatusV3NotFound struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminUpdateUserStatusV3NotFound) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/users/{userId}/status][%d] adminUpdateUserStatusV3NotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateUserStatusV3NotFound) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminUpdateUserStatusV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
