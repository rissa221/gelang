/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// BaseCatalogStatus the model 'BaseCatalogStatus'
type BaseCatalogStatus string

// List of BaseCatalog.Status
const (
	BASECATALOGSTATUS_COMING_SOON BaseCatalogStatus = "COMING_SOON"
	BASECATALOGSTATUS_ACTIVE BaseCatalogStatus = "ACTIVE"
	BASECATALOGSTATUS_DEPRECATED BaseCatalogStatus = "DEPRECATED"
	BASECATALOGSTATUS_DISABLED BaseCatalogStatus = "DISABLED"
	BASECATALOGSTATUS_UNKNOWN BaseCatalogStatus = "UNKNOWN"
)

var allowedBaseCatalogStatusEnumValues = []BaseCatalogStatus{
	"COMING_SOON",
	"ACTIVE",
	"DEPRECATED",
	"DISABLED",
	"UNKNOWN",
}

func (v *BaseCatalogStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BaseCatalogStatus(value)
	for _, existing := range allowedBaseCatalogStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BaseCatalogStatus", value)
}

// NewBaseCatalogStatusFromValue returns a pointer to a valid BaseCatalogStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBaseCatalogStatusFromValue(v string) (*BaseCatalogStatus, error) {
	ev := BaseCatalogStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BaseCatalogStatus: valid values are %v", v, allowedBaseCatalogStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BaseCatalogStatus) IsValid() bool {
	for _, existing := range allowedBaseCatalogStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BaseCatalog.Status value
func (v BaseCatalogStatus) Ptr() *BaseCatalogStatus {
	return &v
}

type NullableBaseCatalogStatus struct {
	value *BaseCatalogStatus
	isSet bool
}

func (v NullableBaseCatalogStatus) Get() *BaseCatalogStatus {
	return v.value
}

func (v *NullableBaseCatalogStatus) Set(val *BaseCatalogStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCatalogStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCatalogStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCatalogStatus(val *BaseCatalogStatus) *NullableBaseCatalogStatus {
	return &NullableBaseCatalogStatus{value: val, isSet: true}
}

func (v NullableBaseCatalogStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCatalogStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

