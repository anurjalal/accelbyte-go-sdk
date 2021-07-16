// Code generated by go-swagger; DO NOT EDIT.

package order_dedicated

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

// SyncOrdersReader is a Reader for the SyncOrders structure.
type SyncOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/orders returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncOrdersOK creates a SyncOrdersOK with default headers values
func NewSyncOrdersOK() *SyncOrdersOK {
	return &SyncOrdersOK{}
}

/*SyncOrdersOK handles this case with default header values.

  successful operation
*/
type SyncOrdersOK struct {
	Payload *platformclientmodels.OrderSyncResult
}

func (o *SyncOrdersOK) Error() string {
	return fmt.Sprintf("[GET /admin/orders][%d] syncOrdersOK  %+v", 200, o.Payload)
}

func (o *SyncOrdersOK) GetPayload() *platformclientmodels.OrderSyncResult {
	return o.Payload
}

func (o *SyncOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.OrderSyncResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
