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

// DeprecatedRedeployRequestInfo struct for DeprecatedRedeployRequestInfo
type DeprecatedRedeployRequestInfo struct {
	DeploymentGroup *string `json:"deployment_group,omitempty"`
	Sha *string `json:"sha,omitempty"`
}

// NewDeprecatedRedeployRequestInfo instantiates a new DeprecatedRedeployRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedRedeployRequestInfo() *DeprecatedRedeployRequestInfo {
	this := DeprecatedRedeployRequestInfo{}
	return &this
}

// NewDeprecatedRedeployRequestInfoWithDefaults instantiates a new DeprecatedRedeployRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedRedeployRequestInfoWithDefaults() *DeprecatedRedeployRequestInfo {
	this := DeprecatedRedeployRequestInfo{}
	return &this
}

// GetDeploymentGroup returns the DeploymentGroup field value if set, zero value otherwise.
func (o *DeprecatedRedeployRequestInfo) GetDeploymentGroup() string {
	if o == nil || o.DeploymentGroup == nil {
		var ret string
		return ret
	}
	return *o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedRedeployRequestInfo) GetDeploymentGroupOk() (*string, bool) {
	if o == nil || o.DeploymentGroup == nil {
		return nil, false
	}
	return o.DeploymentGroup, true
}

// HasDeploymentGroup returns a boolean if a field has been set.
func (o *DeprecatedRedeployRequestInfo) HasDeploymentGroup() bool {
	if o != nil && o.DeploymentGroup != nil {
		return true
	}

	return false
}

// SetDeploymentGroup gets a reference to the given string and assigns it to the DeploymentGroup field.
func (o *DeprecatedRedeployRequestInfo) SetDeploymentGroup(v string) {
	o.DeploymentGroup = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *DeprecatedRedeployRequestInfo) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedRedeployRequestInfo) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *DeprecatedRedeployRequestInfo) HasSha() bool {
	if o != nil && o.Sha != nil {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *DeprecatedRedeployRequestInfo) SetSha(v string) {
	o.Sha = &v
}

func (o DeprecatedRedeployRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentGroup != nil {
		toSerialize["deployment_group"] = o.DeploymentGroup
	}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedRedeployRequestInfo struct {
	value *DeprecatedRedeployRequestInfo
	isSet bool
}

func (v NullableDeprecatedRedeployRequestInfo) Get() *DeprecatedRedeployRequestInfo {
	return v.value
}

func (v *NullableDeprecatedRedeployRequestInfo) Set(val *DeprecatedRedeployRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedRedeployRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedRedeployRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedRedeployRequestInfo(val *DeprecatedRedeployRequestInfo) *NullableDeprecatedRedeployRequestInfo {
	return &NullableDeprecatedRedeployRequestInfo{value: val, isSet: true}
}

func (v NullableDeprecatedRedeployRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedRedeployRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


