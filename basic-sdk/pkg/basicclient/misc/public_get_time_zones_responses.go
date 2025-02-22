// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// PublicGetTimeZonesReader is a Reader for the PublicGetTimeZones structure.
type PublicGetTimeZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetTimeZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetTimeZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicGetTimeZonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/public/namespaces/{namespace}/misc/timezones returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetTimeZonesOK creates a PublicGetTimeZonesOK with default headers values
func NewPublicGetTimeZonesOK() *PublicGetTimeZonesOK {
	return &PublicGetTimeZonesOK{}
}

/*PublicGetTimeZonesOK handles this case with default header values.

  successful operation
*/
type PublicGetTimeZonesOK struct {
	Payload []string
}

func (o *PublicGetTimeZonesOK) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/misc/timezones][%d] publicGetTimeZonesOK  %+v", 200, o.Payload)
}

func (o *PublicGetTimeZonesOK) GetPayload() []string {
	return o.Payload
}

func (o *PublicGetTimeZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetTimeZonesBadRequest creates a PublicGetTimeZonesBadRequest with default headers values
func NewPublicGetTimeZonesBadRequest() *PublicGetTimeZonesBadRequest {
	return &PublicGetTimeZonesBadRequest{}
}

/*PublicGetTimeZonesBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PublicGetTimeZonesBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *PublicGetTimeZonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/misc/timezones][%d] publicGetTimeZonesBadRequest  %+v", 400, o.Payload)
}

func (o *PublicGetTimeZonesBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *PublicGetTimeZonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
