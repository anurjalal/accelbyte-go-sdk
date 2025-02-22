// Code generated by go-swagger; DO NOT EDIT.

package slot_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNamespaceSlotConfigReader is a Reader for the DeleteNamespaceSlotConfig structure.
type DeleteNamespaceSlotConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNamespaceSlotConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNamespaceSlotConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /admin/namespaces/{namespace}/config returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteNamespaceSlotConfigNoContent creates a DeleteNamespaceSlotConfigNoContent with default headers values
func NewDeleteNamespaceSlotConfigNoContent() *DeleteNamespaceSlotConfigNoContent {
	return &DeleteNamespaceSlotConfigNoContent{}
}

/*DeleteNamespaceSlotConfigNoContent handles this case with default header values.

  Successful delete of namespace slot config
*/
type DeleteNamespaceSlotConfigNoContent struct {
}

func (o *DeleteNamespaceSlotConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/config][%d] deleteNamespaceSlotConfigNoContent ", 204)
}

func (o *DeleteNamespaceSlotConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
