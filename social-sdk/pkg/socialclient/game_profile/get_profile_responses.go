// Code generated by go-swagger; DO NOT EDIT.

package game_profile

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

// GetProfileReader is a Reader for the GetProfile structure.
type GetProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/users/{userId}/profiles/{profileId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetProfileOK creates a GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {
	return &GetProfileOK{}
}

/*GetProfileOK handles this case with default header values.

  successful operation
*/
type GetProfileOK struct {
	Payload *socialclientmodels.GameProfileInfo
}

func (o *GetProfileOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/profiles/{profileId}][%d] getProfileOK  %+v", 200, o.Payload)
}

func (o *GetProfileOK) GetPayload() *socialclientmodels.GameProfileInfo {
	return o.Payload
}

func (o *GetProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.GameProfileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileNotFound creates a GetProfileNotFound with default headers values
func NewGetProfileNotFound() *GetProfileNotFound {
	return &GetProfileNotFound{}
}

/*GetProfileNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12041</td><td>Game profile with id [{profileId}] is not found</td></tr></table>
*/
type GetProfileNotFound struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *GetProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/profiles/{profileId}][%d] getProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetProfileNotFound) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
