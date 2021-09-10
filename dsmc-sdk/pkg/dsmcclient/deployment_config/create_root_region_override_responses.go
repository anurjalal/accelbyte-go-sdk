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

// CreateRootRegionOverrideReader is a Reader for the CreateRootRegionOverride structure.
type CreateRootRegionOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRootRegionOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRootRegionOverrideCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRootRegionOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRootRegionOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateRootRegionOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateRootRegionOverrideConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateRootRegionOverrideInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreateRootRegionOverrideCreated creates a CreateRootRegionOverrideCreated with default headers values
func NewCreateRootRegionOverrideCreated() *CreateRootRegionOverrideCreated {
	return &CreateRootRegionOverrideCreated{}
}

/*CreateRootRegionOverrideCreated handles this case with default header values.

  region override created
*/
type CreateRootRegionOverrideCreated struct {
	Payload *dsmcclientmodels.ModelsDeploymentWithOverride
}

func (o *CreateRootRegionOverrideCreated) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideCreated  %+v", 201, o.Payload)
}

func (o *CreateRootRegionOverrideCreated) GetPayload() *dsmcclientmodels.ModelsDeploymentWithOverride {
	return o.Payload
}

func (o *CreateRootRegionOverrideCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsDeploymentWithOverride)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootRegionOverrideBadRequest creates a CreateRootRegionOverrideBadRequest with default headers values
func NewCreateRootRegionOverrideBadRequest() *CreateRootRegionOverrideBadRequest {
	return &CreateRootRegionOverrideBadRequest{}
}

/*CreateRootRegionOverrideBadRequest handles this case with default header values.

  malformed request
*/
type CreateRootRegionOverrideBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateRootRegionOverrideBadRequest) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRootRegionOverrideBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateRootRegionOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootRegionOverrideUnauthorized creates a CreateRootRegionOverrideUnauthorized with default headers values
func NewCreateRootRegionOverrideUnauthorized() *CreateRootRegionOverrideUnauthorized {
	return &CreateRootRegionOverrideUnauthorized{}
}

/*CreateRootRegionOverrideUnauthorized handles this case with default header values.

  Unauthorized
*/
type CreateRootRegionOverrideUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateRootRegionOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRootRegionOverrideUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateRootRegionOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootRegionOverrideNotFound creates a CreateRootRegionOverrideNotFound with default headers values
func NewCreateRootRegionOverrideNotFound() *CreateRootRegionOverrideNotFound {
	return &CreateRootRegionOverrideNotFound{}
}

/*CreateRootRegionOverrideNotFound handles this case with default header values.

  deployment not found
*/
type CreateRootRegionOverrideNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateRootRegionOverrideNotFound) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideNotFound  %+v", 404, o.Payload)
}

func (o *CreateRootRegionOverrideNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateRootRegionOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootRegionOverrideConflict creates a CreateRootRegionOverrideConflict with default headers values
func NewCreateRootRegionOverrideConflict() *CreateRootRegionOverrideConflict {
	return &CreateRootRegionOverrideConflict{}
}

/*CreateRootRegionOverrideConflict handles this case with default header values.

  deployment already has region override
*/
type CreateRootRegionOverrideConflict struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateRootRegionOverrideConflict) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideConflict  %+v", 409, o.Payload)
}

func (o *CreateRootRegionOverrideConflict) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateRootRegionOverrideConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootRegionOverrideInternalServerError creates a CreateRootRegionOverrideInternalServerError with default headers values
func NewCreateRootRegionOverrideInternalServerError() *CreateRootRegionOverrideInternalServerError {
	return &CreateRootRegionOverrideInternalServerError{}
}

/*CreateRootRegionOverrideInternalServerError handles this case with default header values.

  Internal Server Error
*/
type CreateRootRegionOverrideInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateRootRegionOverrideInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] createRootRegionOverrideInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRootRegionOverrideInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateRootRegionOverrideInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
