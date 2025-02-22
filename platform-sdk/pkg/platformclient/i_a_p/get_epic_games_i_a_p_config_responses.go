// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// GetEpicGamesIAPConfigReader is a Reader for the GetEpicGamesIAPConfig structure.
type GetEpicGamesIAPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEpicGamesIAPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEpicGamesIAPConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEpicGamesIAPConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/iap/config/epicgames returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetEpicGamesIAPConfigOK creates a GetEpicGamesIAPConfigOK with default headers values
func NewGetEpicGamesIAPConfigOK() *GetEpicGamesIAPConfigOK {
	return &GetEpicGamesIAPConfigOK{}
}

/*GetEpicGamesIAPConfigOK handles this case with default header values.

  successful operation
*/
type GetEpicGamesIAPConfigOK struct {
	Payload *platformclientmodels.EpicGamesIAPConfigInfo
}

func (o *GetEpicGamesIAPConfigOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/iap/config/epicgames][%d] getEpicGamesIAPConfigOK  %+v", 200, o.Payload)
}

func (o *GetEpicGamesIAPConfigOK) GetPayload() *platformclientmodels.EpicGamesIAPConfigInfo {
	return o.Payload
}

func (o *GetEpicGamesIAPConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.EpicGamesIAPConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEpicGamesIAPConfigNotFound creates a GetEpicGamesIAPConfigNotFound with default headers values
func NewGetEpicGamesIAPConfigNotFound() *GetEpicGamesIAPConfigNotFound {
	return &GetEpicGamesIAPConfigNotFound{}
}

/*GetEpicGamesIAPConfigNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>39243</td><td>EpicGames config does not exist</td></tr></table>
*/
type GetEpicGamesIAPConfigNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetEpicGamesIAPConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/iap/config/epicgames][%d] getEpicGamesIAPConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetEpicGamesIAPConfigNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetEpicGamesIAPConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
