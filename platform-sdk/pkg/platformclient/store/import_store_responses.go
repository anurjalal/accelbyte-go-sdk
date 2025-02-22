// Code generated by go-swagger; DO NOT EDIT.

package store

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

// ImportStoreReader is a Reader for the ImportStore structure.
type ImportStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImportStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/stores/import returns an error %d: %s", response.Code(), string(data))
	}
}

// NewImportStoreOK creates a ImportStoreOK with default headers values
func NewImportStoreOK() *ImportStoreOK {
	return &ImportStoreOK{}
}

/*ImportStoreOK handles this case with default header values.

  successful operation
*/
type ImportStoreOK struct {
	Payload *platformclientmodels.StoreInfo
}

func (o *ImportStoreOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/stores/import][%d] importStoreOK  %+v", 200, o.Payload)
}

func (o *ImportStoreOK) GetPayload() *platformclientmodels.StoreInfo {
	return o.Payload
}

func (o *ImportStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.StoreInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportStoreBadRequest creates a ImportStoreBadRequest with default headers values
func NewImportStoreBadRequest() *ImportStoreBadRequest {
	return &ImportStoreBadRequest{}
}

/*ImportStoreBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30121</td><td>Store data is invalid</td></tr><tr><td>30122</td><td>Store's meta mismatch</td></tr></table>
*/
type ImportStoreBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *ImportStoreBadRequest) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/stores/import][%d] importStoreBadRequest  %+v", 400, o.Payload)
}

func (o *ImportStoreBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *ImportStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportStoreNotFound creates a ImportStoreNotFound with default headers values
func NewImportStoreNotFound() *ImportStoreNotFound {
	return &ImportStoreNotFound{}
}

/*ImportStoreNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30141</td><td>Store [{storeId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type ImportStoreNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *ImportStoreNotFound) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/stores/import][%d] importStoreNotFound  %+v", 404, o.Payload)
}

func (o *ImportStoreNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *ImportStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
