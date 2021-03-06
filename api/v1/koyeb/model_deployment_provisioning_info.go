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

// DeploymentProvisioningInfo struct for DeploymentProvisioningInfo
type DeploymentProvisioningInfo struct {
	// The git sha for this build (we resolve the reference at the start of the build).
	Sha *string `json:"sha,omitempty"`
	// The docker image built as a result of this build.
	Image *string `json:"image,omitempty"`
	// The id of the job that ran the build.
	BuildJobId *string `json:"build_job_id,omitempty"`
	// Some info about the build.
	Stages *[]DeploymentProvisioningInfoStage `json:"stages,omitempty"`
}

// NewDeploymentProvisioningInfo instantiates a new DeploymentProvisioningInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfo() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// NewDeploymentProvisioningInfoWithDefaults instantiates a new DeploymentProvisioningInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoWithDefaults() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasSha() bool {
	if o != nil && o.Sha != nil {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *DeploymentProvisioningInfo) SetSha(v string) {
	o.Sha = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DeploymentProvisioningInfo) SetImage(v string) {
	o.Image = &v
}

// GetBuildJobId returns the BuildJobId field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetBuildJobId() string {
	if o == nil || o.BuildJobId == nil {
		var ret string
		return ret
	}
	return *o.BuildJobId
}

// GetBuildJobIdOk returns a tuple with the BuildJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetBuildJobIdOk() (*string, bool) {
	if o == nil || o.BuildJobId == nil {
		return nil, false
	}
	return o.BuildJobId, true
}

// HasBuildJobId returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasBuildJobId() bool {
	if o != nil && o.BuildJobId != nil {
		return true
	}

	return false
}

// SetBuildJobId gets a reference to the given string and assigns it to the BuildJobId field.
func (o *DeploymentProvisioningInfo) SetBuildJobId(v string) {
	o.BuildJobId = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetStages() []DeploymentProvisioningInfoStage {
	if o == nil || o.Stages == nil {
		var ret []DeploymentProvisioningInfoStage
		return ret
	}
	return *o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetStagesOk() (*[]DeploymentProvisioningInfoStage, bool) {
	if o == nil || o.Stages == nil {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasStages() bool {
	if o != nil && o.Stages != nil {
		return true
	}

	return false
}

// SetStages gets a reference to the given []DeploymentProvisioningInfoStage and assigns it to the Stages field.
func (o *DeploymentProvisioningInfo) SetStages(v []DeploymentProvisioningInfoStage) {
	o.Stages = &v
}

func (o DeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.BuildJobId != nil {
		toSerialize["build_job_id"] = o.BuildJobId
	}
	if o.Stages != nil {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfo struct {
	value *DeploymentProvisioningInfo
	isSet bool
}

func (v NullableDeploymentProvisioningInfo) Get() *DeploymentProvisioningInfo {
	return v.value
}

func (v *NullableDeploymentProvisioningInfo) Set(val *DeploymentProvisioningInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfo(val *DeploymentProvisioningInfo) *NullableDeploymentProvisioningInfo {
	return &NullableDeploymentProvisioningInfo{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


