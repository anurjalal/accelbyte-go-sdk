// Code generated by go-swagger; DO NOT EDIT.

package public_game_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// PutGameRecordHandlerV1Reader is a Reader for the PutGameRecordHandlerV1 structure.
type PutGameRecordHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGameRecordHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGameRecordHandlerV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutGameRecordHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutGameRecordHandlerV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return result, nil
	}
}

// NewPutGameRecordHandlerV1OK creates a PutGameRecordHandlerV1OK with default headers values
func NewPutGameRecordHandlerV1OK() *PutGameRecordHandlerV1OK {
	return &PutGameRecordHandlerV1OK{}
}

/*PutGameRecordHandlerV1OK handles this case with default header values.

  Record saved
*/
type PutGameRecordHandlerV1OK struct {
}

func (o *PutGameRecordHandlerV1OK) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/records/{key}][%d] putGameRecordHandlerV1OK ", 200)
}

func (o *PutGameRecordHandlerV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGameRecordHandlerV1InternalServerError creates a PutGameRecordHandlerV1InternalServerError with default headers values
func NewPutGameRecordHandlerV1InternalServerError() *PutGameRecordHandlerV1InternalServerError {
	return &PutGameRecordHandlerV1InternalServerError{}
}

/*PutGameRecordHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type PutGameRecordHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *PutGameRecordHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/records/{key}][%d] putGameRecordHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PutGameRecordHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *PutGameRecordHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGameRecordHandlerV1Default creates a PutGameRecordHandlerV1Default with default headers values
func NewPutGameRecordHandlerV1Default(code int) *PutGameRecordHandlerV1Default {
	return &PutGameRecordHandlerV1Default{
		_statusCode: code,
	}
}

/*PutGameRecordHandlerV1Default handles this case with default header values.

  Record saved
*/
type PutGameRecordHandlerV1Default struct {
	_statusCode int
}

// Code gets the status code for the put game record handler v1 default response
func (o *PutGameRecordHandlerV1Default) Code() int {
	return o._statusCode
}

func (o *PutGameRecordHandlerV1Default) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/records/{key}][%d] putGameRecordHandlerV1 default ", o._statusCode)
}

func (o *PutGameRecordHandlerV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
