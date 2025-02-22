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

// AdminDeleteRolePermissionsV4Reader is a Reader for the AdminDeleteRolePermissionsV4 structure.
type AdminDeleteRolePermissionsV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteRolePermissionsV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeleteRolePermissionsV4NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteRolePermissionsV4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminDeleteRolePermissionsV4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminDeleteRolePermissionsV4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/v4/admin/roles/{roleId}/permissions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeleteRolePermissionsV4NoContent creates a AdminDeleteRolePermissionsV4NoContent with default headers values
func NewAdminDeleteRolePermissionsV4NoContent() *AdminDeleteRolePermissionsV4NoContent {
	return &AdminDeleteRolePermissionsV4NoContent{}
}

/*AdminDeleteRolePermissionsV4NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminDeleteRolePermissionsV4NoContent struct {
}

func (o *AdminDeleteRolePermissionsV4NoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/v4/admin/roles/{roleId}/permissions][%d] adminDeleteRolePermissionsV4NoContent ", 204)
}

func (o *AdminDeleteRolePermissionsV4NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteRolePermissionsV4Unauthorized creates a AdminDeleteRolePermissionsV4Unauthorized with default headers values
func NewAdminDeleteRolePermissionsV4Unauthorized() *AdminDeleteRolePermissionsV4Unauthorized {
	return &AdminDeleteRolePermissionsV4Unauthorized{}
}

/*AdminDeleteRolePermissionsV4Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminDeleteRolePermissionsV4Unauthorized struct {
}

func (o *AdminDeleteRolePermissionsV4Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/v4/admin/roles/{roleId}/permissions][%d] adminDeleteRolePermissionsV4Unauthorized ", 401)
}

func (o *AdminDeleteRolePermissionsV4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteRolePermissionsV4Forbidden creates a AdminDeleteRolePermissionsV4Forbidden with default headers values
func NewAdminDeleteRolePermissionsV4Forbidden() *AdminDeleteRolePermissionsV4Forbidden {
	return &AdminDeleteRolePermissionsV4Forbidden{}
}

/*AdminDeleteRolePermissionsV4Forbidden handles this case with default header values.

  Forbidden
*/
type AdminDeleteRolePermissionsV4Forbidden struct {
}

func (o *AdminDeleteRolePermissionsV4Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/v4/admin/roles/{roleId}/permissions][%d] adminDeleteRolePermissionsV4Forbidden ", 403)
}

func (o *AdminDeleteRolePermissionsV4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteRolePermissionsV4NotFound creates a AdminDeleteRolePermissionsV4NotFound with default headers values
func NewAdminDeleteRolePermissionsV4NotFound() *AdminDeleteRolePermissionsV4NotFound {
	return &AdminDeleteRolePermissionsV4NotFound{}
}

/*AdminDeleteRolePermissionsV4NotFound handles this case with default header values.

  Data not found
*/
type AdminDeleteRolePermissionsV4NotFound struct {
}

func (o *AdminDeleteRolePermissionsV4NotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/v4/admin/roles/{roleId}/permissions][%d] adminDeleteRolePermissionsV4NotFound ", 404)
}

func (o *AdminDeleteRolePermissionsV4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
