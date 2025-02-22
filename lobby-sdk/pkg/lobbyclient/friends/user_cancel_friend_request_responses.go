// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// UserCancelFriendRequestReader is a Reader for the UserCancelFriendRequest structure.
type UserCancelFriendRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCancelFriendRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCancelFriendRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserCancelFriendRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserCancelFriendRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserCancelFriendRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserCancelFriendRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUserCancelFriendRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /friends/namespaces/{namespace}/me/request/cancel returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUserCancelFriendRequestNoContent creates a UserCancelFriendRequestNoContent with default headers values
func NewUserCancelFriendRequestNoContent() *UserCancelFriendRequestNoContent {
	return &UserCancelFriendRequestNoContent{}
}

/*UserCancelFriendRequestNoContent handles this case with default header values.

  No Content
*/
type UserCancelFriendRequestNoContent struct {
}

func (o *UserCancelFriendRequestNoContent) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestNoContent ", 204)
}

func (o *UserCancelFriendRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCancelFriendRequestBadRequest creates a UserCancelFriendRequestBadRequest with default headers values
func NewUserCancelFriendRequestBadRequest() *UserCancelFriendRequestBadRequest {
	return &UserCancelFriendRequestBadRequest{}
}

/*UserCancelFriendRequestBadRequest handles this case with default header values.

  Bad Request
*/
type UserCancelFriendRequestBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UserCancelFriendRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestBadRequest  %+v", 400, o.Payload)
}

func (o *UserCancelFriendRequestBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UserCancelFriendRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCancelFriendRequestUnauthorized creates a UserCancelFriendRequestUnauthorized with default headers values
func NewUserCancelFriendRequestUnauthorized() *UserCancelFriendRequestUnauthorized {
	return &UserCancelFriendRequestUnauthorized{}
}

/*UserCancelFriendRequestUnauthorized handles this case with default header values.

  Unauthorized
*/
type UserCancelFriendRequestUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UserCancelFriendRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *UserCancelFriendRequestUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UserCancelFriendRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCancelFriendRequestForbidden creates a UserCancelFriendRequestForbidden with default headers values
func NewUserCancelFriendRequestForbidden() *UserCancelFriendRequestForbidden {
	return &UserCancelFriendRequestForbidden{}
}

/*UserCancelFriendRequestForbidden handles this case with default header values.

  Forbidden
*/
type UserCancelFriendRequestForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UserCancelFriendRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestForbidden  %+v", 403, o.Payload)
}

func (o *UserCancelFriendRequestForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UserCancelFriendRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCancelFriendRequestNotFound creates a UserCancelFriendRequestNotFound with default headers values
func NewUserCancelFriendRequestNotFound() *UserCancelFriendRequestNotFound {
	return &UserCancelFriendRequestNotFound{}
}

/*UserCancelFriendRequestNotFound handles this case with default header values.

  Not Found
*/
type UserCancelFriendRequestNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UserCancelFriendRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestNotFound  %+v", 404, o.Payload)
}

func (o *UserCancelFriendRequestNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UserCancelFriendRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCancelFriendRequestInternalServerError creates a UserCancelFriendRequestInternalServerError with default headers values
func NewUserCancelFriendRequestInternalServerError() *UserCancelFriendRequestInternalServerError {
	return &UserCancelFriendRequestInternalServerError{}
}

/*UserCancelFriendRequestInternalServerError handles this case with default header values.

  Internal Server Error
*/
type UserCancelFriendRequestInternalServerError struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UserCancelFriendRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /friends/namespaces/{namespace}/me/request/cancel][%d] userCancelFriendRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *UserCancelFriendRequestInternalServerError) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UserCancelFriendRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
