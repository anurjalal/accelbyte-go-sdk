// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Slide slide
//
// swagger:model Slide
type Slide struct {

	// alt
	Alt string `json:"alt,omitempty"`

	// preview url
	PreviewURL string `json:"previewUrl,omitempty"`

	// thumbnail url
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`

	// slide type
	// Enum: [image video]
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// video source
	// Enum: [generic youtube vimeo]
	VideoSource string `json:"videoSource,omitempty"`
}

// Validate validates this slide
func (m *Slide) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVideoSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var slideTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image","video"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		slideTypeTypePropEnum = append(slideTypeTypePropEnum, v)
	}
}

const (

	// SlideTypeImage captures enum value "image"
	SlideTypeImage string = "image"

	// SlideTypeVideo captures enum value "video"
	SlideTypeVideo string = "video"
)

// prop value enum
func (m *Slide) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, slideTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Slide) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var slideTypeVideoSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["generic","youtube","vimeo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		slideTypeVideoSourcePropEnum = append(slideTypeVideoSourcePropEnum, v)
	}
}

const (

	// SlideVideoSourceGeneric captures enum value "generic"
	SlideVideoSourceGeneric string = "generic"

	// SlideVideoSourceYoutube captures enum value "youtube"
	SlideVideoSourceYoutube string = "youtube"

	// SlideVideoSourceVimeo captures enum value "vimeo"
	SlideVideoSourceVimeo string = "vimeo"
)

// prop value enum
func (m *Slide) validateVideoSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, slideTypeVideoSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Slide) validateVideoSource(formats strfmt.Registry) error {

	if swag.IsZero(m.VideoSource) { // not required
		return nil
	}

	// value enum
	if err := m.validateVideoSourceEnum("videoSource", "body", m.VideoSource); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Slide) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Slide) UnmarshalBinary(b []byte) error {
	var res Slide
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
