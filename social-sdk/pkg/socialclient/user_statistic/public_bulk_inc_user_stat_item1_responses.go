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

// PublicBulkIncUserStatItem1Reader is a Reader for the PublicBulkIncUserStatItem1 structure.
type PublicBulkIncUserStatItem1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicBulkIncUserStatItem1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicBulkIncUserStatItem1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPublicBulkIncUserStatItem1UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicBulkIncUserStatItem1OK creates a PublicBulkIncUserStatItem1OK with default headers values
func NewPublicBulkIncUserStatItem1OK() *PublicBulkIncUserStatItem1OK {
	return &PublicBulkIncUserStatItem1OK{}
}

/*PublicBulkIncUserStatItem1OK handles this case with default header values.

  successful operation
*/
type PublicBulkIncUserStatItem1OK struct {
	Payload []*socialclientmodels.BulkStatItemOperationResult
}

func (o *PublicBulkIncUserStatItem1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] publicBulkIncUserStatItem1OK  %+v", 200, o.Payload)
}

func (o *PublicBulkIncUserStatItem1OK) GetPayload() []*socialclientmodels.BulkStatItemOperationResult {
	return o.Payload
}

func (o *PublicBulkIncUserStatItem1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicBulkIncUserStatItem1UnprocessableEntity creates a PublicBulkIncUserStatItem1UnprocessableEntity with default headers values
func NewPublicBulkIncUserStatItem1UnprocessableEntity() *PublicBulkIncUserStatItem1UnprocessableEntity {
	return &PublicBulkIncUserStatItem1UnprocessableEntity{}
}

/*PublicBulkIncUserStatItem1UnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PublicBulkIncUserStatItem1UnprocessableEntity struct {
	Payload *socialclientmodels.ValidationErrorEntity
}

func (o *PublicBulkIncUserStatItem1UnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] publicBulkIncUserStatItem1UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PublicBulkIncUserStatItem1UnprocessableEntity) GetPayload() *socialclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *PublicBulkIncUserStatItem1UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
