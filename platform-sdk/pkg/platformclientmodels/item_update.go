// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemUpdate A DTO object for updating item API call.
//
// swagger:model ItemUpdate
type ItemUpdate struct {

	// appId, required if appType is present, alpha numeric, max length is 255
	AppID string `json:"appId,omitempty"`

	// appType
	// Enum: [GAME SOFTWARE DLC DEMO]
	AppType string `json:"appType,omitempty"`

	// baseAppId, can set value of game appId if want to link to a game, only publisher namespace item will take effect
	BaseAppID string `json:"baseAppId,omitempty"`

	// booth name to get tickets while ItemType is CODE, Campaign or KeyGroup should located in targetNamespace if targetNamespace not null
	BoothName string `json:"boothName,omitempty"`

	// Category Path, A path separated by "/", will not show in store if it set to blank, max length is 255
	CategoryPath string `json:"categoryPath,omitempty"`

	// customized item clazz
	Clazz string `json:"clazz,omitempty"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// Entitlement Type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	EntitlementType *string `json:"entitlementType"`

	// customized item properties
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Features, allowed characters from a-zA-Z0-9_:- and start/end in alphanumeric with length range from 1 to 127
	// Unique: true
	Features []string `json:"features"`

	// images
	Images []*Image `json:"images"`

	// itemIds, should be empty if item type is not "BUNDLE"
	ItemIds []string `json:"itemIds"`

	// Item Type
	// Required: true
	// Enum: [APP COINS INGAMEITEM BUNDLE CODE SUBSCRIPTION]
	ItemType *string `json:"itemType"`

	// Localization, key language, value localization content
	Localizations map[string]Localization `json:"localizations,omitempty"`

	// Max count, -1 means UNLIMITED, new value should >= old value if both old value and new value is limited, unset when itemType is CODE
	MaxCount int32 `json:"maxCount,omitempty"`

	// Max count per user, -1 means UNLIMITED
	MaxCountPerUser int32 `json:"maxCountPerUser,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// recurring for subscription
	Recurring *Recurring `json:"recurring,omitempty"`

	// region data map, key is region, value is region data list
	RegionData map[string][]RegionDataItem `json:"regionData,omitempty"`

	// sku, allowed characters from a-zA-Z0-9_:- and start/end in alphanumeric, max length is 127
	Sku string `json:"sku,omitempty"`

	// Whether to stack the entitlement when entitlement type is "CONSUMABLE"
	Stackable bool `json:"stackable,omitempty"`

	// status
	// Enum: [ACTIVE INACTIVE]
	Status string `json:"status,omitempty"`

	// Tags, allowed characters from a-zA-Z:_- with length range from 1 to 30, should start and end in upper/lowercase, an item has max 5 tags.
	// Unique: true
	Tags []string `json:"tags"`

	// target currency code, required if item type is "COINS"
	TargetCurrencyCode string `json:"targetCurrencyCode,omitempty"`

	// target namespace, require when sell a game's item on the publisher namespace
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// thumbnail Url
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`

	// Represent entitlement count when entitlement type is "CONSUMABLE", and represent coin count when itemType is "COINS"
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this item update
func (m *ItemUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var itemUpdateTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GAME","SOFTWARE","DLC","DEMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemUpdateTypeAppTypePropEnum = append(itemUpdateTypeAppTypePropEnum, v)
	}
}

const (

	// ItemUpdateAppTypeGAME captures enum value "GAME"
	ItemUpdateAppTypeGAME string = "GAME"

	// ItemUpdateAppTypeSOFTWARE captures enum value "SOFTWARE"
	ItemUpdateAppTypeSOFTWARE string = "SOFTWARE"

	// ItemUpdateAppTypeDLC captures enum value "DLC"
	ItemUpdateAppTypeDLC string = "DLC"

	// ItemUpdateAppTypeDEMO captures enum value "DEMO"
	ItemUpdateAppTypeDEMO string = "DEMO"
)

// prop value enum
func (m *ItemUpdate) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemUpdateTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemUpdate) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

var itemUpdateTypeEntitlementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemUpdateTypeEntitlementTypePropEnum = append(itemUpdateTypeEntitlementTypePropEnum, v)
	}
}

const (

	// ItemUpdateEntitlementTypeDURABLE captures enum value "DURABLE"
	ItemUpdateEntitlementTypeDURABLE string = "DURABLE"

	// ItemUpdateEntitlementTypeCONSUMABLE captures enum value "CONSUMABLE"
	ItemUpdateEntitlementTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *ItemUpdate) validateEntitlementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemUpdateTypeEntitlementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemUpdate) validateEntitlementType(formats strfmt.Registry) error {

	if err := validate.Required("entitlementType", "body", m.EntitlementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntitlementTypeEnum("entitlementType", "body", *m.EntitlementType); err != nil {
		return err
	}

	return nil
}

func (m *ItemUpdate) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *ItemUpdate) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var itemUpdateTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","COINS","INGAMEITEM","BUNDLE","CODE","SUBSCRIPTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemUpdateTypeItemTypePropEnum = append(itemUpdateTypeItemTypePropEnum, v)
	}
}

const (

	// ItemUpdateItemTypeAPP captures enum value "APP"
	ItemUpdateItemTypeAPP string = "APP"

	// ItemUpdateItemTypeCOINS captures enum value "COINS"
	ItemUpdateItemTypeCOINS string = "COINS"

	// ItemUpdateItemTypeINGAMEITEM captures enum value "INGAMEITEM"
	ItemUpdateItemTypeINGAMEITEM string = "INGAMEITEM"

	// ItemUpdateItemTypeBUNDLE captures enum value "BUNDLE"
	ItemUpdateItemTypeBUNDLE string = "BUNDLE"

	// ItemUpdateItemTypeCODE captures enum value "CODE"
	ItemUpdateItemTypeCODE string = "CODE"

	// ItemUpdateItemTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	ItemUpdateItemTypeSUBSCRIPTION string = "SUBSCRIPTION"
)

// prop value enum
func (m *ItemUpdate) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemUpdateTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemUpdate) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *ItemUpdate) validateLocalizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Localizations) { // not required
		return nil
	}

	for k := range m.Localizations {

		if err := validate.Required("localizations"+"."+k, "body", m.Localizations[k]); err != nil {
			return err
		}
		if val, ok := m.Localizations[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ItemUpdate) validateRecurring(formats strfmt.Registry) error {

	if swag.IsZero(m.Recurring) { // not required
		return nil
	}

	if m.Recurring != nil {
		if err := m.Recurring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recurring")
			}
			return err
		}
	}

	return nil
}

func (m *ItemUpdate) validateRegionData(formats strfmt.Registry) error {

	if swag.IsZero(m.RegionData) { // not required
		return nil
	}

	for k := range m.RegionData {

		if err := validate.Required("regionData"+"."+k, "body", m.RegionData[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.RegionData[k]); i++ {

			if err := m.RegionData[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regionData" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var itemUpdateTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemUpdateTypeStatusPropEnum = append(itemUpdateTypeStatusPropEnum, v)
	}
}

const (

	// ItemUpdateStatusACTIVE captures enum value "ACTIVE"
	ItemUpdateStatusACTIVE string = "ACTIVE"

	// ItemUpdateStatusINACTIVE captures enum value "INACTIVE"
	ItemUpdateStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *ItemUpdate) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemUpdateTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemUpdate) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ItemUpdate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemUpdate) UnmarshalBinary(b []byte) error {
	var res ItemUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
