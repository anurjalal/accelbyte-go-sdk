// Code generated by go-swagger; DO NOT EDIT.

package payment_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// PublicGetPaymentMethodsReader is a Reader for the PublicGetPaymentMethods structure.
type PublicGetPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetPaymentMethodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetPaymentMethodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/payment/methods returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetPaymentMethodsOK creates a PublicGetPaymentMethodsOK with default headers values
func NewPublicGetPaymentMethodsOK() *PublicGetPaymentMethodsOK {
	return &PublicGetPaymentMethodsOK{}
}

/*PublicGetPaymentMethodsOK handles this case with default header values.

  successful operation
*/
type PublicGetPaymentMethodsOK struct {
	Payload []*platformclientmodels.PaymentMethod
}

func (o *PublicGetPaymentMethodsOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/payment/methods][%d] publicGetPaymentMethodsOK  %+v", 200, o.Payload)
}

func (o *PublicGetPaymentMethodsOK) GetPayload() []*platformclientmodels.PaymentMethod {
	return o.Payload
}

func (o *PublicGetPaymentMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetPaymentMethodsNotFound creates a PublicGetPaymentMethodsNotFound with default headers values
func NewPublicGetPaymentMethodsNotFound() *PublicGetPaymentMethodsNotFound {
	return &PublicGetPaymentMethodsNotFound{}
}

/*PublicGetPaymentMethodsNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>33141</td><td>Payment Order [{paymentOrderNo}] does not exist</td></tr></table>
*/
type PublicGetPaymentMethodsNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PublicGetPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/payment/methods][%d] publicGetPaymentMethodsNotFound  %+v", 404, o.Payload)
}

func (o *PublicGetPaymentMethodsNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicGetPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
