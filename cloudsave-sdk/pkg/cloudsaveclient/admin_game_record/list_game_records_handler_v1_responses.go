// Code generated by go-swagger; DO NOT EDIT.

package admin_game_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// ListGameRecordsHandlerV1Reader is a Reader for the ListGameRecordsHandlerV1 structure.
type ListGameRecordsHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGameRecordsHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGameRecordsHandlerV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListGameRecordsHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGameRecordsHandlerV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return result, nil
	}
}

// NewListGameRecordsHandlerV1OK creates a ListGameRecordsHandlerV1OK with default headers values
func NewListGameRecordsHandlerV1OK() *ListGameRecordsHandlerV1OK {
	return &ListGameRecordsHandlerV1OK{}
}

/*ListGameRecordsHandlerV1OK handles this case with default header values.

  Retrieve list of records key by namespace
*/
type ListGameRecordsHandlerV1OK struct {
	Payload *cloudsaveclientmodels.ModelsListGameRecordKeys
}

func (o *ListGameRecordsHandlerV1OK) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/records][%d] listGameRecordsHandlerV1OK  %+v", 200, o.Payload)
}

func (o *ListGameRecordsHandlerV1OK) GetPayload() *cloudsaveclientmodels.ModelsListGameRecordKeys {
	return o.Payload
}

func (o *ListGameRecordsHandlerV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsListGameRecordKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGameRecordsHandlerV1InternalServerError creates a ListGameRecordsHandlerV1InternalServerError with default headers values
func NewListGameRecordsHandlerV1InternalServerError() *ListGameRecordsHandlerV1InternalServerError {
	return &ListGameRecordsHandlerV1InternalServerError{}
}

/*ListGameRecordsHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type ListGameRecordsHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *ListGameRecordsHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/records][%d] listGameRecordsHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ListGameRecordsHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *ListGameRecordsHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGameRecordsHandlerV1Default creates a ListGameRecordsHandlerV1Default with default headers values
func NewListGameRecordsHandlerV1Default(code int) *ListGameRecordsHandlerV1Default {
	return &ListGameRecordsHandlerV1Default{
		_statusCode: code,
	}
}

/*ListGameRecordsHandlerV1Default handles this case with default header values.

  Retrieve list of records key by namespace
*/
type ListGameRecordsHandlerV1Default struct {
	_statusCode int

	Payload *cloudsaveclientmodels.ModelsListGameRecordKeys
}

// Code gets the status code for the list game records handler v1 default response
func (o *ListGameRecordsHandlerV1Default) Code() int {
	return o._statusCode
}

func (o *ListGameRecordsHandlerV1Default) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/records][%d] listGameRecordsHandlerV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ListGameRecordsHandlerV1Default) GetPayload() *cloudsaveclientmodels.ModelsListGameRecordKeys {
	return o.Payload
}

func (o *ListGameRecordsHandlerV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsListGameRecordKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
