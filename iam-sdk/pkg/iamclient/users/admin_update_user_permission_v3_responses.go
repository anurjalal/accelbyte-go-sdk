// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminUpdateUserPermissionV3Reader is a Reader for the AdminUpdateUserPermissionV3 structure.
type AdminUpdateUserPermissionV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserPermissionV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminUpdateUserPermissionV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserPermissionV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminUpdateUserPermissionV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminUpdateUserPermissionV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminUpdateUserPermissionV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminUpdateUserPermissionV3NoContent creates a AdminUpdateUserPermissionV3NoContent with default headers values
func NewAdminUpdateUserPermissionV3NoContent() *AdminUpdateUserPermissionV3NoContent {
	return &AdminUpdateUserPermissionV3NoContent{}
}

/*AdminUpdateUserPermissionV3NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminUpdateUserPermissionV3NoContent struct {
}

func (o *AdminUpdateUserPermissionV3NoContent) Error() string {
	return fmt.Sprintf("[PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions][%d] adminUpdateUserPermissionV3NoContent ", 204)
}

func (o *AdminUpdateUserPermissionV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserPermissionV3BadRequest creates a AdminUpdateUserPermissionV3BadRequest with default headers values
func NewAdminUpdateUserPermissionV3BadRequest() *AdminUpdateUserPermissionV3BadRequest {
	return &AdminUpdateUserPermissionV3BadRequest{}
}

/*AdminUpdateUserPermissionV3BadRequest handles this case with default header values.

  Invalid request
*/
type AdminUpdateUserPermissionV3BadRequest struct {
}

func (o *AdminUpdateUserPermissionV3BadRequest) Error() string {
	return fmt.Sprintf("[PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions][%d] adminUpdateUserPermissionV3BadRequest ", 400)
}

func (o *AdminUpdateUserPermissionV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserPermissionV3Unauthorized creates a AdminUpdateUserPermissionV3Unauthorized with default headers values
func NewAdminUpdateUserPermissionV3Unauthorized() *AdminUpdateUserPermissionV3Unauthorized {
	return &AdminUpdateUserPermissionV3Unauthorized{}
}

/*AdminUpdateUserPermissionV3Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminUpdateUserPermissionV3Unauthorized struct {
}

func (o *AdminUpdateUserPermissionV3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions][%d] adminUpdateUserPermissionV3Unauthorized ", 401)
}

func (o *AdminUpdateUserPermissionV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserPermissionV3Forbidden creates a AdminUpdateUserPermissionV3Forbidden with default headers values
func NewAdminUpdateUserPermissionV3Forbidden() *AdminUpdateUserPermissionV3Forbidden {
	return &AdminUpdateUserPermissionV3Forbidden{}
}

/*AdminUpdateUserPermissionV3Forbidden handles this case with default header values.

  Forbidden
*/
type AdminUpdateUserPermissionV3Forbidden struct {
}

func (o *AdminUpdateUserPermissionV3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions][%d] adminUpdateUserPermissionV3Forbidden ", 403)
}

func (o *AdminUpdateUserPermissionV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminUpdateUserPermissionV3NotFound creates a AdminUpdateUserPermissionV3NotFound with default headers values
func NewAdminUpdateUserPermissionV3NotFound() *AdminUpdateUserPermissionV3NotFound {
	return &AdminUpdateUserPermissionV3NotFound{}
}

/*AdminUpdateUserPermissionV3NotFound handles this case with default header values.

  Data not found
*/
type AdminUpdateUserPermissionV3NotFound struct {
}

func (o *AdminUpdateUserPermissionV3NotFound) Error() string {
	return fmt.Sprintf("[PUT /iam/v3/admin/namespaces/{namespace}/users/{userId}/permissions][%d] adminUpdateUserPermissionV3NotFound ", 404)
}

func (o *AdminUpdateUserPermissionV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
