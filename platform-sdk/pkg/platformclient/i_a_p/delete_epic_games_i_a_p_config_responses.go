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

// DeleteEpicGamesIAPConfigReader is a Reader for the DeleteEpicGamesIAPConfig structure.
type DeleteEpicGamesIAPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEpicGamesIAPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEpicGamesIAPConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /admin/namespaces/{namespace}/iap/config/epicgames returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteEpicGamesIAPConfigNoContent creates a DeleteEpicGamesIAPConfigNoContent with default headers values
func NewDeleteEpicGamesIAPConfigNoContent() *DeleteEpicGamesIAPConfigNoContent {
	return &DeleteEpicGamesIAPConfigNoContent{}
}

/*DeleteEpicGamesIAPConfigNoContent handles this case with default header values.

  Delete successfully
*/
type DeleteEpicGamesIAPConfigNoContent struct {
}

func (o *DeleteEpicGamesIAPConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/iap/config/epicgames][%d] deleteEpicGamesIAPConfigNoContent ", 204)
}

func (o *DeleteEpicGamesIAPConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
