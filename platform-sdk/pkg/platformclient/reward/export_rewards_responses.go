// Code generated by go-swagger; DO NOT EDIT.

package reward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportRewardsReader is a Reader for the ExportRewards structure.
type ExportRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/rewards/export returns an error %d: %s", response.Code(), string(data))
	}
}

// NewExportRewardsOK creates a ExportRewardsOK with default headers values
func NewExportRewardsOK() *ExportRewardsOK {
	return &ExportRewardsOK{}
}

/*ExportRewardsOK handles this case with default header values.

  successful export of reward configs
*/
type ExportRewardsOK struct {
}

func (o *ExportRewardsOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/rewards/export][%d] exportRewardsOK ", 200)
}

func (o *ExportRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
