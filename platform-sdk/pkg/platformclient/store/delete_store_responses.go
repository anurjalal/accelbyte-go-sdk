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

// DeleteStoreReader is a Reader for the DeleteStore structure.
type DeleteStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewDeleteStoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /admin/namespaces/{namespace}/stores/{storeId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteStoreOK creates a DeleteStoreOK with default headers values
func NewDeleteStoreOK() *DeleteStoreOK {
	return &DeleteStoreOK{}
}

/*DeleteStoreOK handles this case with default header values.

  successful operation
*/
type DeleteStoreOK struct {
	Payload *platformclientmodels.StoreInfo
}

func (o *DeleteStoreOK) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/stores/{storeId}][%d] deleteStoreOK  %+v", 200, o.Payload)
}

func (o *DeleteStoreOK) GetPayload() *platformclientmodels.StoreInfo {
	return o.Payload
}

func (o *DeleteStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.StoreInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStoreNotFound creates a DeleteStoreNotFound with default headers values
func NewDeleteStoreNotFound() *DeleteStoreNotFound {
	return &DeleteStoreNotFound{}
}

/*DeleteStoreNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30141</td><td>Store [{storeId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type DeleteStoreNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DeleteStoreNotFound) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/stores/{storeId}][%d] deleteStoreNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStoreNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStoreConflict creates a DeleteStoreConflict with default headers values
func NewDeleteStoreConflict() *DeleteStoreConflict {
	return &DeleteStoreConflict{}
}

/*DeleteStoreConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30173</td><td>Published store can't modify content</td></tr></table>
*/
type DeleteStoreConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DeleteStoreConflict) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/stores/{storeId}][%d] deleteStoreConflict  %+v", 409, o.Payload)
}

func (o *DeleteStoreConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteStoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
