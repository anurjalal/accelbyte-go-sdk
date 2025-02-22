// Code generated by go-swagger; DO NOT EDIT.

package entitlement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserDistributionReceiverReader is a Reader for the UpdateUserDistributionReceiver structure.
type UpdateUserDistributionReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserDistributionReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateUserDistributionReceiverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/users/{userId}/entitlements/receivers/{extUserId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateUserDistributionReceiverNoContent creates a UpdateUserDistributionReceiverNoContent with default headers values
func NewUpdateUserDistributionReceiverNoContent() *UpdateUserDistributionReceiverNoContent {
	return &UpdateUserDistributionReceiverNoContent{}
}

/*UpdateUserDistributionReceiverNoContent handles this case with default header values.

  create distribution receiver successfully
*/
type UpdateUserDistributionReceiverNoContent struct {
}

func (o *UpdateUserDistributionReceiverNoContent) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/users/{userId}/entitlements/receivers/{extUserId}][%d] updateUserDistributionReceiverNoContent ", 204)
}

func (o *UpdateUserDistributionReceiverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
