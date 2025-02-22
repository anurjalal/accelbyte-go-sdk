// Code generated by go-swagger; DO NOT EDIT.

package admin_player_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// AdminDeletePlayerPublicRecordHandlerV1Reader is a Reader for the AdminDeletePlayerPublicRecordHandlerV1 structure.
type AdminDeletePlayerPublicRecordHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeletePlayerPublicRecordHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeletePlayerPublicRecordHandlerV1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminDeletePlayerPublicRecordHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /cloudsave/v1/admin/namespaces/{namespace}/users/{userID}/records/{key}/public returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeletePlayerPublicRecordHandlerV1NoContent creates a AdminDeletePlayerPublicRecordHandlerV1NoContent with default headers values
func NewAdminDeletePlayerPublicRecordHandlerV1NoContent() *AdminDeletePlayerPublicRecordHandlerV1NoContent {
	return &AdminDeletePlayerPublicRecordHandlerV1NoContent{}
}

/*AdminDeletePlayerPublicRecordHandlerV1NoContent handles this case with default header values.

  Player public record deleted
*/
type AdminDeletePlayerPublicRecordHandlerV1NoContent struct {
}

func (o *AdminDeletePlayerPublicRecordHandlerV1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /cloudsave/v1/admin/namespaces/{namespace}/users/{userID}/records/{key}/public][%d] adminDeletePlayerPublicRecordHandlerV1NoContent ", 204)
}

func (o *AdminDeletePlayerPublicRecordHandlerV1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlayerPublicRecordHandlerV1InternalServerError creates a AdminDeletePlayerPublicRecordHandlerV1InternalServerError with default headers values
func NewAdminDeletePlayerPublicRecordHandlerV1InternalServerError() *AdminDeletePlayerPublicRecordHandlerV1InternalServerError {
	return &AdminDeletePlayerPublicRecordHandlerV1InternalServerError{}
}

/*AdminDeletePlayerPublicRecordHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminDeletePlayerPublicRecordHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *AdminDeletePlayerPublicRecordHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloudsave/v1/admin/namespaces/{namespace}/users/{userID}/records/{key}/public][%d] adminDeletePlayerPublicRecordHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDeletePlayerPublicRecordHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminDeletePlayerPublicRecordHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
