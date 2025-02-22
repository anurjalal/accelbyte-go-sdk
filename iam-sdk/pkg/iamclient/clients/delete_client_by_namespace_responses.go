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

// DeleteClientByNamespaceReader is a Reader for the DeleteClientByNamespace structure.
type DeleteClientByNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientByNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClientByNamespaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClientByNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteClientByNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteClientByNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/namespaces/{namespace}/clients/{clientId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteClientByNamespaceNoContent creates a DeleteClientByNamespaceNoContent with default headers values
func NewDeleteClientByNamespaceNoContent() *DeleteClientByNamespaceNoContent {
	return &DeleteClientByNamespaceNoContent{}
}

/*DeleteClientByNamespaceNoContent handles this case with default header values.

  Operation succeeded
*/
type DeleteClientByNamespaceNoContent struct {
}

func (o *DeleteClientByNamespaceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/clients/{clientId}][%d] deleteClientByNamespaceNoContent ", 204)
}

func (o *DeleteClientByNamespaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClientByNamespaceUnauthorized creates a DeleteClientByNamespaceUnauthorized with default headers values
func NewDeleteClientByNamespaceUnauthorized() *DeleteClientByNamespaceUnauthorized {
	return &DeleteClientByNamespaceUnauthorized{}
}

/*DeleteClientByNamespaceUnauthorized handles this case with default header values.

  Unauthorized access
*/
type DeleteClientByNamespaceUnauthorized struct {
}

func (o *DeleteClientByNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/clients/{clientId}][%d] deleteClientByNamespaceUnauthorized ", 401)
}

func (o *DeleteClientByNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClientByNamespaceForbidden creates a DeleteClientByNamespaceForbidden with default headers values
func NewDeleteClientByNamespaceForbidden() *DeleteClientByNamespaceForbidden {
	return &DeleteClientByNamespaceForbidden{}
}

/*DeleteClientByNamespaceForbidden handles this case with default header values.

  Forbidden
*/
type DeleteClientByNamespaceForbidden struct {
}

func (o *DeleteClientByNamespaceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/clients/{clientId}][%d] deleteClientByNamespaceForbidden ", 403)
}

func (o *DeleteClientByNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClientByNamespaceNotFound creates a DeleteClientByNamespaceNotFound with default headers values
func NewDeleteClientByNamespaceNotFound() *DeleteClientByNamespaceNotFound {
	return &DeleteClientByNamespaceNotFound{}
}

/*DeleteClientByNamespaceNotFound handles this case with default header values.

  Data not found
*/
type DeleteClientByNamespaceNotFound struct {
}

func (o *DeleteClientByNamespaceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/clients/{clientId}][%d] deleteClientByNamespaceNotFound ", 404)
}

func (o *DeleteClientByNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
