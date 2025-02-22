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

// PublicCreateJusticeUserReader is a Reader for the PublicCreateJusticeUser structure.
type PublicCreateJusticeUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicCreateJusticeUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPublicCreateJusticeUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicCreateJusticeUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicCreateJusticeUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicCreateJusticeUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicCreateJusticeUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPublicCreateJusticeUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicCreateJusticeUserCreated creates a PublicCreateJusticeUserCreated with default headers values
func NewPublicCreateJusticeUserCreated() *PublicCreateJusticeUserCreated {
	return &PublicCreateJusticeUserCreated{}
}

/*PublicCreateJusticeUserCreated handles this case with default header values.

  Created
*/
type PublicCreateJusticeUserCreated struct {
	Payload *iamclientmodels.ModelCreateJusticeUserResponse
}

func (o *PublicCreateJusticeUserCreated) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserCreated  %+v", 201, o.Payload)
}

func (o *PublicCreateJusticeUserCreated) GetPayload() *iamclientmodels.ModelCreateJusticeUserResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelCreateJusticeUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicCreateJusticeUserBadRequest creates a PublicCreateJusticeUserBadRequest with default headers values
func NewPublicCreateJusticeUserBadRequest() *PublicCreateJusticeUserBadRequest {
	return &PublicCreateJusticeUserBadRequest{}
}

/*PublicCreateJusticeUserBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PublicCreateJusticeUserBadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicCreateJusticeUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserBadRequest  %+v", 400, o.Payload)
}

func (o *PublicCreateJusticeUserBadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicCreateJusticeUserUnauthorized creates a PublicCreateJusticeUserUnauthorized with default headers values
func NewPublicCreateJusticeUserUnauthorized() *PublicCreateJusticeUserUnauthorized {
	return &PublicCreateJusticeUserUnauthorized{}
}

/*PublicCreateJusticeUserUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr><tr><td>20022</td><td>token is not user token</td></tr></table>
*/
type PublicCreateJusticeUserUnauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicCreateJusticeUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserUnauthorized  %+v", 401, o.Payload)
}

func (o *PublicCreateJusticeUserUnauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicCreateJusticeUserForbidden creates a PublicCreateJusticeUserForbidden with default headers values
func NewPublicCreateJusticeUserForbidden() *PublicCreateJusticeUserForbidden {
	return &PublicCreateJusticeUserForbidden{}
}

/*PublicCreateJusticeUserForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type PublicCreateJusticeUserForbidden struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicCreateJusticeUserForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserForbidden  %+v", 403, o.Payload)
}

func (o *PublicCreateJusticeUserForbidden) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicCreateJusticeUserNotFound creates a PublicCreateJusticeUserNotFound with default headers values
func NewPublicCreateJusticeUserNotFound() *PublicCreateJusticeUserNotFound {
	return &PublicCreateJusticeUserNotFound{}
}

/*PublicCreateJusticeUserNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type PublicCreateJusticeUserNotFound struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicCreateJusticeUserNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserNotFound  %+v", 404, o.Payload)
}

func (o *PublicCreateJusticeUserNotFound) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicCreateJusticeUserInternalServerError creates a PublicCreateJusticeUserInternalServerError with default headers values
func NewPublicCreateJusticeUserInternalServerError() *PublicCreateJusticeUserInternalServerError {
	return &PublicCreateJusticeUserInternalServerError{}
}

/*PublicCreateJusticeUserInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type PublicCreateJusticeUserInternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicCreateJusticeUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/platforms/justice/{targetNamespace}][%d] publicCreateJusticeUserInternalServerError  %+v", 500, o.Payload)
}

func (o *PublicCreateJusticeUserInternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicCreateJusticeUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
