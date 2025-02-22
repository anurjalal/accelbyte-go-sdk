// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncSteamInventoryReader is a Reader for the SyncSteamInventory structure.
type SyncSteamInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncSteamInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSyncSteamInventoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /public/namespaces/{namespace}/users/{userId}/iap/steam/sync returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncSteamInventoryNoContent creates a SyncSteamInventoryNoContent with default headers values
func NewSyncSteamInventoryNoContent() *SyncSteamInventoryNoContent {
	return &SyncSteamInventoryNoContent{}
}

/*SyncSteamInventoryNoContent handles this case with default header values.

  successful operation
*/
type SyncSteamInventoryNoContent struct {
}

func (o *SyncSteamInventoryNoContent) Error() string {
	return fmt.Sprintf("[PUT /public/namespaces/{namespace}/users/{userId}/iap/steam/sync][%d] syncSteamInventoryNoContent ", 204)
}

func (o *SyncSteamInventoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
