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

// GetUserByLoginIDReader is a Reader for the GetUserByLoginID structure.
type GetUserByLoginIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByLoginIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByLoginIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserByLoginIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserByLoginIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUserByLoginIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/namespaces/{namespace}/users/byLoginId returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserByLoginIDOK creates a GetUserByLoginIDOK with default headers values
func NewGetUserByLoginIDOK() *GetUserByLoginIDOK {
	return &GetUserByLoginIDOK{}
}

/*GetUserByLoginIDOK handles this case with default header values.

  OK
*/
type GetUserByLoginIDOK struct {
	Payload *iamclientmodels.ModelPublicUserResponse
}

func (o *GetUserByLoginIDOK) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/byLoginId][%d] getUserByLoginIdOK  %+v", 200, o.Payload)
}

func (o *GetUserByLoginIDOK) GetPayload() *iamclientmodels.ModelPublicUserResponse {
	return o.Payload
}

func (o *GetUserByLoginIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelPublicUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByLoginIDBadRequest creates a GetUserByLoginIDBadRequest with default headers values
func NewGetUserByLoginIDBadRequest() *GetUserByLoginIDBadRequest {
	return &GetUserByLoginIDBadRequest{}
}

/*GetUserByLoginIDBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type GetUserByLoginIDBadRequest struct {
}

func (o *GetUserByLoginIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/byLoginId][%d] getUserByLoginIdBadRequest ", 400)
}

func (o *GetUserByLoginIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserByLoginIDNotFound creates a GetUserByLoginIDNotFound with default headers values
func NewGetUserByLoginIDNotFound() *GetUserByLoginIDNotFound {
	return &GetUserByLoginIDNotFound{}
}

/*GetUserByLoginIDNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type GetUserByLoginIDNotFound struct {
}

func (o *GetUserByLoginIDNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/byLoginId][%d] getUserByLoginIdNotFound ", 404)
}

func (o *GetUserByLoginIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserByLoginIDInternalServerError creates a GetUserByLoginIDInternalServerError with default headers values
func NewGetUserByLoginIDInternalServerError() *GetUserByLoginIDInternalServerError {
	return &GetUserByLoginIDInternalServerError{}
}

/*GetUserByLoginIDInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type GetUserByLoginIDInternalServerError struct {
}

func (o *GetUserByLoginIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/byLoginId][%d] getUserByLoginIdInternalServerError ", 500)
}

func (o *GetUserByLoginIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
