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

// GetCountriesReader is a Reader for the GetCountries structure.
type GetCountriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCountriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCountriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCountriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCountriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/admin/namespaces/{namespace}/misc/countries returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetCountriesOK creates a GetCountriesOK with default headers values
func NewGetCountriesOK() *GetCountriesOK {
	return &GetCountriesOK{}
}

/*GetCountriesOK handles this case with default header values.

  successful operation
*/
type GetCountriesOK struct {
	Payload []*basicclientmodels.CountryObject
}

func (o *GetCountriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/misc/countries][%d] getCountriesOK  %+v", 200, o.Payload)
}

func (o *GetCountriesOK) GetPayload() []*basicclientmodels.CountryObject {
	return o.Payload
}

func (o *GetCountriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCountriesBadRequest creates a GetCountriesBadRequest with default headers values
func NewGetCountriesBadRequest() *GetCountriesBadRequest {
	return &GetCountriesBadRequest{}
}

/*GetCountriesBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type GetCountriesBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *GetCountriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/misc/countries][%d] getCountriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCountriesBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *GetCountriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCountriesUnauthorized creates a GetCountriesUnauthorized with default headers values
func NewGetCountriesUnauthorized() *GetCountriesUnauthorized {
	return &GetCountriesUnauthorized{}
}

/*GetCountriesUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type GetCountriesUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetCountriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/misc/countries][%d] getCountriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCountriesUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetCountriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
