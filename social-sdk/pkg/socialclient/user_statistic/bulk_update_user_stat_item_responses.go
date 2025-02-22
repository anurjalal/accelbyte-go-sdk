// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// BulkUpdateUserStatItemReader is a Reader for the BulkUpdateUserStatItem structure.
type BulkUpdateUserStatItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkUpdateUserStatItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkUpdateUserStatItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewBulkUpdateUserStatItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /v2/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkUpdateUserStatItemOK creates a BulkUpdateUserStatItemOK with default headers values
func NewBulkUpdateUserStatItemOK() *BulkUpdateUserStatItemOK {
	return &BulkUpdateUserStatItemOK{}
}

/*BulkUpdateUserStatItemOK handles this case with default header values.

  successful operation
*/
type BulkUpdateUserStatItemOK struct {
	Payload []*socialclientmodels.BulkStatItemOperationResult
}

func (o *BulkUpdateUserStatItemOK) Error() string {
	return fmt.Sprintf("[PUT /v2/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] bulkUpdateUserStatItemOK  %+v", 200, o.Payload)
}

func (o *BulkUpdateUserStatItemOK) GetPayload() []*socialclientmodels.BulkStatItemOperationResult {
	return o.Payload
}

func (o *BulkUpdateUserStatItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkUpdateUserStatItemUnprocessableEntity creates a BulkUpdateUserStatItemUnprocessableEntity with default headers values
func NewBulkUpdateUserStatItemUnprocessableEntity() *BulkUpdateUserStatItemUnprocessableEntity {
	return &BulkUpdateUserStatItemUnprocessableEntity{}
}

/*BulkUpdateUserStatItemUnprocessableEntity handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type BulkUpdateUserStatItemUnprocessableEntity struct {
	Payload *socialclientmodels.ValidationErrorEntity
}

func (o *BulkUpdateUserStatItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v2/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] bulkUpdateUserStatItemUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BulkUpdateUserStatItemUnprocessableEntity) GetPayload() *socialclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *BulkUpdateUserStatItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
