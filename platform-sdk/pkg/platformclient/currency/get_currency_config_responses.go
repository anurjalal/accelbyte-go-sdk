// Code generated by go-swagger; DO NOT EDIT.

package currency

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

// GetCurrencyConfigReader is a Reader for the GetCurrencyConfig structure.
type GetCurrencyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrencyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrencyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCurrencyConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/currencies/{currencyCode}/config returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetCurrencyConfigOK creates a GetCurrencyConfigOK with default headers values
func NewGetCurrencyConfigOK() *GetCurrencyConfigOK {
	return &GetCurrencyConfigOK{}
}

/*GetCurrencyConfigOK handles this case with default header values.

  successful operation
*/
type GetCurrencyConfigOK struct {
	Payload *platformclientmodels.CurrencyConfig
}

func (o *GetCurrencyConfigOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/currencies/{currencyCode}/config][%d] getCurrencyConfigOK  %+v", 200, o.Payload)
}

func (o *GetCurrencyConfigOK) GetPayload() *platformclientmodels.CurrencyConfig {
	return o.Payload
}

func (o *GetCurrencyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.CurrencyConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrencyConfigNotFound creates a GetCurrencyConfigNotFound with default headers values
func NewGetCurrencyConfigNotFound() *GetCurrencyConfigNotFound {
	return &GetCurrencyConfigNotFound{}
}

/*GetCurrencyConfigNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>36141</td><td>Currency [{currencyCode}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type GetCurrencyConfigNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetCurrencyConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/currencies/{currencyCode}/config][%d] getCurrencyConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrencyConfigNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetCurrencyConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
