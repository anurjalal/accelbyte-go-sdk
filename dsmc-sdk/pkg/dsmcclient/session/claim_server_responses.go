// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// ClaimServerReader is a Reader for the ClaimServer structure.
type ClaimServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewClaimServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewClaimServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewClaimServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewClaimServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 425:
		result := NewClaimServerStatus425()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewClaimServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewClaimServerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/namespaces/{namespace}/sessions/claim returns an error %d: %s", response.Code(), string(data))
	}
}

// NewClaimServerNoContent creates a ClaimServerNoContent with default headers values
func NewClaimServerNoContent() *ClaimServerNoContent {
	return &ClaimServerNoContent{}
}

/*ClaimServerNoContent handles this case with default header values.

  DS claimed for session
*/
type ClaimServerNoContent struct {
}

func (o *ClaimServerNoContent) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerNoContent ", 204)
}

func (o *ClaimServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimServerUnauthorized creates a ClaimServerUnauthorized with default headers values
func NewClaimServerUnauthorized() *ClaimServerUnauthorized {
	return &ClaimServerUnauthorized{}
}

/*ClaimServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type ClaimServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerUnauthorized  %+v", 401, o.Payload)
}

func (o *ClaimServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimServerNotFound creates a ClaimServerNotFound with default headers values
func NewClaimServerNotFound() *ClaimServerNotFound {
	return &ClaimServerNotFound{}
}

/*ClaimServerNotFound handles this case with default header values.

  session not found
*/
type ClaimServerNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerNotFound) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerNotFound  %+v", 404, o.Payload)
}

func (o *ClaimServerNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimServerConflict creates a ClaimServerConflict with default headers values
func NewClaimServerConflict() *ClaimServerConflict {
	return &ClaimServerConflict{}
}

/*ClaimServerConflict handles this case with default header values.

  DS is already claimed
*/
type ClaimServerConflict struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerConflict) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerConflict  %+v", 409, o.Payload)
}

func (o *ClaimServerConflict) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimServerStatus425 creates a ClaimServerStatus425 with default headers values
func NewClaimServerStatus425() *ClaimServerStatus425 {
	return &ClaimServerStatus425{}
}

/*ClaimServerStatus425 handles this case with default header values.

  DS is not ready to be claimed
*/
type ClaimServerStatus425 struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerStatus425) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerStatus425  %+v", 425, o.Payload)
}

func (o *ClaimServerStatus425) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerStatus425) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimServerInternalServerError creates a ClaimServerInternalServerError with default headers values
func NewClaimServerInternalServerError() *ClaimServerInternalServerError {
	return &ClaimServerInternalServerError{}
}

/*ClaimServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type ClaimServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerInternalServerError  %+v", 500, o.Payload)
}

func (o *ClaimServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimServerServiceUnavailable creates a ClaimServerServiceUnavailable with default headers values
func NewClaimServerServiceUnavailable() *ClaimServerServiceUnavailable {
	return &ClaimServerServiceUnavailable{}
}

/*ClaimServerServiceUnavailable handles this case with default header values.

  DS is unreachable
*/
type ClaimServerServiceUnavailable struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ClaimServerServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/sessions/claim][%d] claimServerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ClaimServerServiceUnavailable) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ClaimServerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
