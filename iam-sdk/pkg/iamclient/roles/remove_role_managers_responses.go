// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveRoleManagersReader is a Reader for the RemoveRoleManagers structure.
type RemoveRoleManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveRoleManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveRoleManagersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveRoleManagersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveRoleManagersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRemoveRoleManagersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRemoveRoleManagersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/roles/{roleId}/managers returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRemoveRoleManagersNoContent creates a RemoveRoleManagersNoContent with default headers values
func NewRemoveRoleManagersNoContent() *RemoveRoleManagersNoContent {
	return &RemoveRoleManagersNoContent{}
}

/*RemoveRoleManagersNoContent handles this case with default header values.

  Operation succeeded
*/
type RemoveRoleManagersNoContent struct {
}

func (o *RemoveRoleManagersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/managers][%d] removeRoleManagersNoContent ", 204)
}

func (o *RemoveRoleManagersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleManagersBadRequest creates a RemoveRoleManagersBadRequest with default headers values
func NewRemoveRoleManagersBadRequest() *RemoveRoleManagersBadRequest {
	return &RemoveRoleManagersBadRequest{}
}

/*RemoveRoleManagersBadRequest handles this case with default header values.

  Invalid request
*/
type RemoveRoleManagersBadRequest struct {
}

func (o *RemoveRoleManagersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/managers][%d] removeRoleManagersBadRequest ", 400)
}

func (o *RemoveRoleManagersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleManagersUnauthorized creates a RemoveRoleManagersUnauthorized with default headers values
func NewRemoveRoleManagersUnauthorized() *RemoveRoleManagersUnauthorized {
	return &RemoveRoleManagersUnauthorized{}
}

/*RemoveRoleManagersUnauthorized handles this case with default header values.

  Unauthorized access
*/
type RemoveRoleManagersUnauthorized struct {
}

func (o *RemoveRoleManagersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/managers][%d] removeRoleManagersUnauthorized ", 401)
}

func (o *RemoveRoleManagersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleManagersForbidden creates a RemoveRoleManagersForbidden with default headers values
func NewRemoveRoleManagersForbidden() *RemoveRoleManagersForbidden {
	return &RemoveRoleManagersForbidden{}
}

/*RemoveRoleManagersForbidden handles this case with default header values.

  Forbidden
*/
type RemoveRoleManagersForbidden struct {
}

func (o *RemoveRoleManagersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/managers][%d] removeRoleManagersForbidden ", 403)
}

func (o *RemoveRoleManagersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleManagersNotFound creates a RemoveRoleManagersNotFound with default headers values
func NewRemoveRoleManagersNotFound() *RemoveRoleManagersNotFound {
	return &RemoveRoleManagersNotFound{}
}

/*RemoveRoleManagersNotFound handles this case with default header values.

  Data not found
*/
type RemoveRoleManagersNotFound struct {
}

func (o *RemoveRoleManagersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/managers][%d] removeRoleManagersNotFound ", 404)
}

func (o *RemoveRoleManagersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
