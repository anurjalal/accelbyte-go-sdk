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

// SearchItemsReader is a Reader for the SearchItems structure.
type SearchItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/items/search returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSearchItemsOK creates a SearchItemsOK with default headers values
func NewSearchItemsOK() *SearchItemsOK {
	return &SearchItemsOK{}
}

/*SearchItemsOK handles this case with default header values.

  successful operation
*/
type SearchItemsOK struct {
	Payload *platformclientmodels.FullItemPagingSlicedResult
}

func (o *SearchItemsOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/items/search][%d] searchItemsOK  %+v", 200, o.Payload)
}

func (o *SearchItemsOK) GetPayload() *platformclientmodels.FullItemPagingSlicedResult {
	return o.Payload
}

func (o *SearchItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FullItemPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
