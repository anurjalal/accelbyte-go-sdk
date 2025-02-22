// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateClientPermissionReader is a Reader for the UpdateClientPermission structure.
type UpdateClientPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClientPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateClientPermissionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClientPermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateClientPermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateClientPermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateClientPermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/clients/{clientId}/clientpermissions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateClientPermissionNoContent creates a UpdateClientPermissionNoContent with default headers values
func NewUpdateClientPermissionNoContent() *UpdateClientPermissionNoContent {
	return &UpdateClientPermissionNoContent{}
}

/*UpdateClientPermissionNoContent handles this case with default header values.

  Operation succeeded
*/
type UpdateClientPermissionNoContent struct {
}

func (o *UpdateClientPermissionNoContent) Error() string {
	return fmt.Sprintf("[POST /iam/clients/{clientId}/clientpermissions][%d] updateClientPermissionNoContent ", 204)
}

func (o *UpdateClientPermissionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClientPermissionBadRequest creates a UpdateClientPermissionBadRequest with default headers values
func NewUpdateClientPermissionBadRequest() *UpdateClientPermissionBadRequest {
	return &UpdateClientPermissionBadRequest{}
}

/*UpdateClientPermissionBadRequest handles this case with default header values.

  Invalid request
*/
type UpdateClientPermissionBadRequest struct {
}

func (o *UpdateClientPermissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/clients/{clientId}/clientpermissions][%d] updateClientPermissionBadRequest ", 400)
}

func (o *UpdateClientPermissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClientPermissionUnauthorized creates a UpdateClientPermissionUnauthorized with default headers values
func NewUpdateClientPermissionUnauthorized() *UpdateClientPermissionUnauthorized {
	return &UpdateClientPermissionUnauthorized{}
}

/*UpdateClientPermissionUnauthorized handles this case with default header values.

  Unauthorized access
*/
type UpdateClientPermissionUnauthorized struct {
}

func (o *UpdateClientPermissionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/clients/{clientId}/clientpermissions][%d] updateClientPermissionUnauthorized ", 401)
}

func (o *UpdateClientPermissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClientPermissionForbidden creates a UpdateClientPermissionForbidden with default headers values
func NewUpdateClientPermissionForbidden() *UpdateClientPermissionForbidden {
	return &UpdateClientPermissionForbidden{}
}

/*UpdateClientPermissionForbidden handles this case with default header values.

  Forbidden
*/
type UpdateClientPermissionForbidden struct {
}

func (o *UpdateClientPermissionForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/clients/{clientId}/clientpermissions][%d] updateClientPermissionForbidden ", 403)
}

func (o *UpdateClientPermissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClientPermissionNotFound creates a UpdateClientPermissionNotFound with default headers values
func NewUpdateClientPermissionNotFound() *UpdateClientPermissionNotFound {
	return &UpdateClientPermissionNotFound{}
}

/*UpdateClientPermissionNotFound handles this case with default header values.

  Data not found
*/
type UpdateClientPermissionNotFound struct {
}

func (o *UpdateClientPermissionNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/clients/{clientId}/clientpermissions][%d] updateClientPermissionNotFound ", 404)
}

func (o *UpdateClientPermissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
