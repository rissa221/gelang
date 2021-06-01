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
)

// UpdateOrganizationReply struct for UpdateOrganizationReply
type UpdateOrganizationReply struct {
	Organization *Organization `json:"organization,omitempty"`
}

// NewUpdateOrganizationReply instantiates a new UpdateOrganizationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationReply() *UpdateOrganizationReply {
	this := UpdateOrganizationReply{}
	return &this
}

// NewUpdateOrganizationReplyWithDefaults instantiates a new UpdateOrganizationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationReplyWithDefaults() *UpdateOrganizationReply {
	this := UpdateOrganizationReply{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UpdateOrganizationReply) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationReply) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UpdateOrganizationReply) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *UpdateOrganizationReply) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o UpdateOrganizationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationReply struct {
	value *UpdateOrganizationReply
	isSet bool
}

func (v NullableUpdateOrganizationReply) Get() *UpdateOrganizationReply {
	return v.value
}

func (v *NullableUpdateOrganizationReply) Set(val *UpdateOrganizationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationReply(val *UpdateOrganizationReply) *NullableUpdateOrganizationReply {
	return &NullableUpdateOrganizationReply{value: val, isSet: true}
}

func (v NullableUpdateOrganizationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

