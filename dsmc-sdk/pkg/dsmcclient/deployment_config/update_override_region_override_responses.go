// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// UpdateOverrideRegionOverrideReader is a Reader for the UpdateOverrideRegionOverride structure.
type UpdateOverrideRegionOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOverrideRegionOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOverrideRegionOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOverrideRegionOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateOverrideRegionOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateOverrideRegionOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateOverrideRegionOverrideInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateOverrideRegionOverrideOK creates a UpdateOverrideRegionOverrideOK with default headers values
func NewUpdateOverrideRegionOverrideOK() *UpdateOverrideRegionOverrideOK {
	return &UpdateOverrideRegionOverrideOK{}
}

/*UpdateOverrideRegionOverrideOK handles this case with default header values.

  deployment region override updated
*/
type UpdateOverrideRegionOverrideOK struct {
	Payload *dsmcclientmodels.ModelsDeploymentWithOverride
}

func (o *UpdateOverrideRegionOverrideOK) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] updateOverrideRegionOverrideOK  %+v", 200, o.Payload)
}

func (o *UpdateOverrideRegionOverrideOK) GetPayload() *dsmcclientmodels.ModelsDeploymentWithOverride {
	return o.Payload
}

func (o *UpdateOverrideRegionOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsDeploymentWithOverride)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOverrideRegionOverrideBadRequest creates a UpdateOverrideRegionOverrideBadRequest with default headers values
func NewUpdateOverrideRegionOverrideBadRequest() *UpdateOverrideRegionOverrideBadRequest {
	return &UpdateOverrideRegionOverrideBadRequest{}
}

/*UpdateOverrideRegionOverrideBadRequest handles this case with default header values.

  malformed request
*/
type UpdateOverrideRegionOverrideBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateOverrideRegionOverrideBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] updateOverrideRegionOverrideBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOverrideRegionOverrideBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateOverrideRegionOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOverrideRegionOverrideUnauthorized creates a UpdateOverrideRegionOverrideUnauthorized with default headers values
func NewUpdateOverrideRegionOverrideUnauthorized() *UpdateOverrideRegionOverrideUnauthorized {
	return &UpdateOverrideRegionOverrideUnauthorized{}
}

/*UpdateOverrideRegionOverrideUnauthorized handles this case with default header values.

  Unauthorized
*/
type UpdateOverrideRegionOverrideUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateOverrideRegionOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] updateOverrideRegionOverrideUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOverrideRegionOverrideUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateOverrideRegionOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOverrideRegionOverrideNotFound creates a UpdateOverrideRegionOverrideNotFound with default headers values
func NewUpdateOverrideRegionOverrideNotFound() *UpdateOverrideRegionOverrideNotFound {
	return &UpdateOverrideRegionOverrideNotFound{}
}

/*UpdateOverrideRegionOverrideNotFound handles this case with default header values.

  deployment not found
*/
type UpdateOverrideRegionOverrideNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateOverrideRegionOverrideNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] updateOverrideRegionOverrideNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOverrideRegionOverrideNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateOverrideRegionOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOverrideRegionOverrideInternalServerError creates a UpdateOverrideRegionOverrideInternalServerError with default headers values
func NewUpdateOverrideRegionOverrideInternalServerError() *UpdateOverrideRegionOverrideInternalServerError {
	return &UpdateOverrideRegionOverrideInternalServerError{}
}

/*UpdateOverrideRegionOverrideInternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateOverrideRegionOverrideInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateOverrideRegionOverrideInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] updateOverrideRegionOverrideInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOverrideRegionOverrideInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateOverrideRegionOverrideInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
