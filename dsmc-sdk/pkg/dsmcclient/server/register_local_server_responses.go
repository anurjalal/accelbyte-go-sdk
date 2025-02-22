// Code generated by go-swagger; DO NOT EDIT.

package server

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

// RegisterLocalServerReader is a Reader for the RegisterLocalServer structure.
type RegisterLocalServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterLocalServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterLocalServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterLocalServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegisterLocalServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewRegisterLocalServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRegisterLocalServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/namespaces/{namespace}/servers/local/register returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRegisterLocalServerOK creates a RegisterLocalServerOK with default headers values
func NewRegisterLocalServerOK() *RegisterLocalServerOK {
	return &RegisterLocalServerOK{}
}

/*RegisterLocalServerOK handles this case with default header values.

  server registered
*/
type RegisterLocalServerOK struct {
	Payload *dsmcclientmodels.ModelsServer
}

func (o *RegisterLocalServerOK) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/local/register][%d] registerLocalServerOK  %+v", 200, o.Payload)
}

func (o *RegisterLocalServerOK) GetPayload() *dsmcclientmodels.ModelsServer {
	return o.Payload
}

func (o *RegisterLocalServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterLocalServerBadRequest creates a RegisterLocalServerBadRequest with default headers values
func NewRegisterLocalServerBadRequest() *RegisterLocalServerBadRequest {
	return &RegisterLocalServerBadRequest{}
}

/*RegisterLocalServerBadRequest handles this case with default header values.

  malformed request
*/
type RegisterLocalServerBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *RegisterLocalServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/local/register][%d] registerLocalServerBadRequest  %+v", 400, o.Payload)
}

func (o *RegisterLocalServerBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *RegisterLocalServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterLocalServerUnauthorized creates a RegisterLocalServerUnauthorized with default headers values
func NewRegisterLocalServerUnauthorized() *RegisterLocalServerUnauthorized {
	return &RegisterLocalServerUnauthorized{}
}

/*RegisterLocalServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type RegisterLocalServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *RegisterLocalServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/local/register][%d] registerLocalServerUnauthorized  %+v", 401, o.Payload)
}

func (o *RegisterLocalServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *RegisterLocalServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterLocalServerConflict creates a RegisterLocalServerConflict with default headers values
func NewRegisterLocalServerConflict() *RegisterLocalServerConflict {
	return &RegisterLocalServerConflict{}
}

/*RegisterLocalServerConflict handles this case with default header values.

  server with same name already registered
*/
type RegisterLocalServerConflict struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *RegisterLocalServerConflict) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/local/register][%d] registerLocalServerConflict  %+v", 409, o.Payload)
}

func (o *RegisterLocalServerConflict) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *RegisterLocalServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterLocalServerInternalServerError creates a RegisterLocalServerInternalServerError with default headers values
func NewRegisterLocalServerInternalServerError() *RegisterLocalServerInternalServerError {
	return &RegisterLocalServerInternalServerError{}
}

/*RegisterLocalServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type RegisterLocalServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *RegisterLocalServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/local/register][%d] registerLocalServerInternalServerError  %+v", 500, o.Payload)
}

func (o *RegisterLocalServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *RegisterLocalServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
