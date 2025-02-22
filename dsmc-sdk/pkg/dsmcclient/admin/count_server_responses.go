// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// CountServerReader is a Reader for the CountServer structure.
type CountServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCountServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCountServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/admin/namespaces/{namespace}/servers/count returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCountServerOK creates a CountServerOK with default headers values
func NewCountServerOK() *CountServerOK {
	return &CountServerOK{}
}

/*CountServerOK handles this case with default header values.

  servers listed
*/
type CountServerOK struct {
	Payload *dsmcclientmodels.ModelsCountServerResponse
}

func (o *CountServerOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count][%d] countServerOK  %+v", 200, o.Payload)
}

func (o *CountServerOK) GetPayload() *dsmcclientmodels.ModelsCountServerResponse {
	return o.Payload
}

func (o *CountServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsCountServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountServerUnauthorized creates a CountServerUnauthorized with default headers values
func NewCountServerUnauthorized() *CountServerUnauthorized {
	return &CountServerUnauthorized{}
}

/*CountServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type CountServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CountServerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count][%d] countServerUnauthorized  %+v", 401, o.Payload)
}

func (o *CountServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CountServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountServerInternalServerError creates a CountServerInternalServerError with default headers values
func NewCountServerInternalServerError() *CountServerInternalServerError {
	return &CountServerInternalServerError{}
}

/*CountServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type CountServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CountServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count][%d] countServerInternalServerError  %+v", 500, o.Payload)
}

func (o *CountServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CountServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
