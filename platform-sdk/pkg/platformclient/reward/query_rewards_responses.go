// Code generated by go-swagger; DO NOT EDIT.

package reward

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

// QueryRewardsReader is a Reader for the QueryRewards structure.
type QueryRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewQueryRewardsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/rewards/byCriteria returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryRewardsOK creates a QueryRewardsOK with default headers values
func NewQueryRewardsOK() *QueryRewardsOK {
	return &QueryRewardsOK{}
}

/*QueryRewardsOK handles this case with default header values.

  successful operation
*/
type QueryRewardsOK struct {
	Payload *platformclientmodels.RewardPagingSlicedResult
}

func (o *QueryRewardsOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/rewards/byCriteria][%d] queryRewardsOK  %+v", 200, o.Payload)
}

func (o *QueryRewardsOK) GetPayload() *platformclientmodels.RewardPagingSlicedResult {
	return o.Payload
}

func (o *QueryRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.RewardPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRewardsUnprocessableEntity creates a QueryRewardsUnprocessableEntity with default headers values
func NewQueryRewardsUnprocessableEntity() *QueryRewardsUnprocessableEntity {
	return &QueryRewardsUnprocessableEntity{}
}

/*QueryRewardsUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type QueryRewardsUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *QueryRewardsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/rewards/byCriteria][%d] queryRewardsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *QueryRewardsUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *QueryRewardsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
