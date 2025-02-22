// Code generated by go-swagger; DO NOT EDIT.

package item

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

// PublicBulkGetItemsReader is a Reader for the PublicBulkGetItems structure.
type PublicBulkGetItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicBulkGetItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicBulkGetItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/items/locale/byIds returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicBulkGetItemsOK creates a PublicBulkGetItemsOK with default headers values
func NewPublicBulkGetItemsOK() *PublicBulkGetItemsOK {
	return &PublicBulkGetItemsOK{}
}

/*PublicBulkGetItemsOK handles this case with default header values.

  successful operation
*/
type PublicBulkGetItemsOK struct {
	Payload []*platformclientmodels.ItemInfo
}

func (o *PublicBulkGetItemsOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/items/locale/byIds][%d] publicBulkGetItemsOK  %+v", 200, o.Payload)
}

func (o *PublicBulkGetItemsOK) GetPayload() []*platformclientmodels.ItemInfo {
	return o.Payload
}

func (o *PublicBulkGetItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
