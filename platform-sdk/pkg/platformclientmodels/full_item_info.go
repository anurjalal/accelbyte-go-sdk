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

// FullItemInfo full item info
//
// swagger:model FullItemInfo
type FullItemInfo struct {

	// App id, required when itemType is APP
	AppID string `json:"appId,omitempty"`

	// App type, required when itemType is APP
	// Enum: [GAME SOFTWARE DLC DEMO]
	AppType string `json:"appType,omitempty"`

	// Base app id
	BaseAppID string `json:"baseAppId,omitempty"`

	// booth name to get tickets while it's item type is CODE
	BoothName string `json:"boothName,omitempty"`

	// Item category path
	// Required: true
	CategoryPath *string `json:"categoryPath"`

	// customized item clazz
	Clazz string `json:"clazz,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// Entitlement type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	EntitlementType *string `json:"entitlementType"`

	// customized item properties
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Features
	// Unique: true
	Features []string `json:"features"`

	// images
	Images []*Image `json:"images"`

	// Item id
	// Required: true
	ItemID *string `json:"itemId"`

	// Bundle item's item ids
	ItemIds []string `json:"itemIds"`

	// Item type
	// Required: true
	// Enum: [APP COINS INGAMEITEM BUNDLE CODE SUBSCRIPTION]
	ItemType *string `json:"itemType"`

	// Item localizations
	// Required: true
	Localizations map[string]Localization `json:"localizations"`

	// Max count, -1 means UNLIMITED, unset when itemType is CODE
	MaxCount int32 `json:"maxCount,omitempty"`

	// Max count per user, -1 means UNLIMITED
	MaxCountPerUser int32 `json:"maxCountPerUser,omitempty"`

	// Name
	// Required: true
	Name *string `json:"name"`

	// Item namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// recurring for subscription
	Recurring *Recurring `json:"recurring,omitempty"`

	// Item region data
	// Required: true
	RegionData map[string][]RegionDataItem `json:"regionData"`

	// Sku
	Sku string `json:"sku,omitempty"`

	// Whether stack the CONSUMABLE entitlement
	Stackable bool `json:"stackable,omitempty"`

	// Item status
	// Required: true
	// Enum: [ACTIVE INACTIVE]
	Status *string `json:"status"`

	// Tags
	// Unique: true
	Tags []string `json:"tags"`

	// The target currency code of coin Item
	TargetCurrencyCode string `json:"targetCurrencyCode,omitempty"`

	// Target item id if this item is mapping from game namespace
	TargetItemID string `json:"targetItemId,omitempty"`

	// The target namespace of a cross namespace item
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// thumbnail url
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// Item use count, required when entitlement type is consumable or itemType is COINS
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this full item info
func (m *FullItemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
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

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
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

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fullItemInfoTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GAME","SOFTWARE","DLC","DEMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullItemInfoTypeAppTypePropEnum = append(fullItemInfoTypeAppTypePropEnum, v)
	}
}

const (

	// FullItemInfoAppTypeGAME captures enum value "GAME"
	FullItemInfoAppTypeGAME string = "GAME"

	// FullItemInfoAppTypeSOFTWARE captures enum value "SOFTWARE"
	FullItemInfoAppTypeSOFTWARE string = "SOFTWARE"

	// FullItemInfoAppTypeDLC captures enum value "DLC"
	FullItemInfoAppTypeDLC string = "DLC"

	// FullItemInfoAppTypeDEMO captures enum value "DEMO"
	FullItemInfoAppTypeDEMO string = "DEMO"
)

// prop value enum
func (m *FullItemInfo) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullItemInfoTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullItemInfo) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateCategoryPath(formats strfmt.Registry) error {

	if err := validate.Required("categoryPath", "body", m.CategoryPath); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var fullItemInfoTypeEntitlementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullItemInfoTypeEntitlementTypePropEnum = append(fullItemInfoTypeEntitlementTypePropEnum, v)
	}
}

const (

	// FullItemInfoEntitlementTypeDURABLE captures enum value "DURABLE"
	FullItemInfoEntitlementTypeDURABLE string = "DURABLE"

	// FullItemInfoEntitlementTypeCONSUMABLE captures enum value "CONSUMABLE"
	FullItemInfoEntitlementTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *FullItemInfo) validateEntitlementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullItemInfoTypeEntitlementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullItemInfo) validateEntitlementType(formats strfmt.Registry) error {

	if err := validate.Required("entitlementType", "body", m.EntitlementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntitlementTypeEnum("entitlementType", "body", *m.EntitlementType); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateImages(formats strfmt.Registry) error {

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

func (m *FullItemInfo) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("itemId", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var fullItemInfoTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","COINS","INGAMEITEM","BUNDLE","CODE","SUBSCRIPTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullItemInfoTypeItemTypePropEnum = append(fullItemInfoTypeItemTypePropEnum, v)
	}
}

const (

	// FullItemInfoItemTypeAPP captures enum value "APP"
	FullItemInfoItemTypeAPP string = "APP"

	// FullItemInfoItemTypeCOINS captures enum value "COINS"
	FullItemInfoItemTypeCOINS string = "COINS"

	// FullItemInfoItemTypeINGAMEITEM captures enum value "INGAMEITEM"
	FullItemInfoItemTypeINGAMEITEM string = "INGAMEITEM"

	// FullItemInfoItemTypeBUNDLE captures enum value "BUNDLE"
	FullItemInfoItemTypeBUNDLE string = "BUNDLE"

	// FullItemInfoItemTypeCODE captures enum value "CODE"
	FullItemInfoItemTypeCODE string = "CODE"

	// FullItemInfoItemTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	FullItemInfoItemTypeSUBSCRIPTION string = "SUBSCRIPTION"
)

// prop value enum
func (m *FullItemInfo) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullItemInfoTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullItemInfo) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateLocalizations(formats strfmt.Registry) error {

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

func (m *FullItemInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateRecurring(formats strfmt.Registry) error {

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

func (m *FullItemInfo) validateRegionData(formats strfmt.Registry) error {

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

var fullItemInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullItemInfoTypeStatusPropEnum = append(fullItemInfoTypeStatusPropEnum, v)
	}
}

const (

	// FullItemInfoStatusACTIVE captures enum value "ACTIVE"
	FullItemInfoStatusACTIVE string = "ACTIVE"

	// FullItemInfoStatusINACTIVE captures enum value "INACTIVE"
	FullItemInfoStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *FullItemInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullItemInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullItemInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *FullItemInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullItemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullItemInfo) UnmarshalBinary(b []byte) error {
	var res FullItemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
