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
	"time"
)

// Deployment struct for Deployment
type Deployment struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AllocatedAt *time.Time `json:"allocated_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	ChildId *string `json:"child_id,omitempty"`
	Status *DeploymentStatus `json:"status,omitempty"`
	Metadata *DeploymentMetadata `json:"metadata,omitempty"`
	Messages *[]string `json:"messages,omitempty"`
	Version *string `json:"version,omitempty"`
	DeploymentGroup *string `json:"deployment_group,omitempty"`
	Definition *ServiceDefinition `json:"definition,omitempty"`
	State *ServiceRevisionState `json:"state,omitempty"`
}

// NewDeployment instantiates a new Deployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployment() *Deployment {
	this := Deployment{}
	var status DeploymentStatus = DEPLOYMENTSTATUS_PENDING
	this.Status = &status
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	var status DeploymentStatus = DEPLOYMENTSTATUS_PENDING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Deployment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Deployment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Deployment) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Deployment) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Deployment) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Deployment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Deployment) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Deployment) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Deployment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAllocatedAt returns the AllocatedAt field value if set, zero value otherwise.
func (o *Deployment) GetAllocatedAt() time.Time {
	if o == nil || o.AllocatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AllocatedAt
}

// GetAllocatedAtOk returns a tuple with the AllocatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetAllocatedAtOk() (*time.Time, bool) {
	if o == nil || o.AllocatedAt == nil {
		return nil, false
	}
	return o.AllocatedAt, true
}

// HasAllocatedAt returns a boolean if a field has been set.
func (o *Deployment) HasAllocatedAt() bool {
	if o != nil && o.AllocatedAt != nil {
		return true
	}

	return false
}

// SetAllocatedAt gets a reference to the given time.Time and assigns it to the AllocatedAt field.
func (o *Deployment) SetAllocatedAt(v time.Time) {
	o.AllocatedAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Deployment) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Deployment) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Deployment) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Deployment) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Deployment) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Deployment) SetAppId(v string) {
	o.AppId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *Deployment) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *Deployment) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *Deployment) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Deployment) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Deployment) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Deployment) SetParentId(v string) {
	o.ParentId = &v
}

// GetChildId returns the ChildId field value if set, zero value otherwise.
func (o *Deployment) GetChildId() string {
	if o == nil || o.ChildId == nil {
		var ret string
		return ret
	}
	return *o.ChildId
}

// GetChildIdOk returns a tuple with the ChildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetChildIdOk() (*string, bool) {
	if o == nil || o.ChildId == nil {
		return nil, false
	}
	return o.ChildId, true
}

// HasChildId returns a boolean if a field has been set.
func (o *Deployment) HasChildId() bool {
	if o != nil && o.ChildId != nil {
		return true
	}

	return false
}

// SetChildId gets a reference to the given string and assigns it to the ChildId field.
func (o *Deployment) SetChildId(v string) {
	o.ChildId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Deployment) GetStatus() DeploymentStatus {
	if o == nil || o.Status == nil {
		var ret DeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetStatusOk() (*DeploymentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Deployment) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentStatus and assigns it to the Status field.
func (o *Deployment) SetStatus(v DeploymentStatus) {
	o.Status = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Deployment) GetMetadata() DeploymentMetadata {
	if o == nil || o.Metadata == nil {
		var ret DeploymentMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetMetadataOk() (*DeploymentMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Deployment) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DeploymentMetadata and assigns it to the Metadata field.
func (o *Deployment) SetMetadata(v DeploymentMetadata) {
	o.Metadata = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Deployment) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Deployment) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *Deployment) SetMessages(v []string) {
	o.Messages = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Deployment) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Deployment) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Deployment) SetVersion(v string) {
	o.Version = &v
}

// GetDeploymentGroup returns the DeploymentGroup field value if set, zero value otherwise.
func (o *Deployment) GetDeploymentGroup() string {
	if o == nil || o.DeploymentGroup == nil {
		var ret string
		return ret
	}
	return *o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDeploymentGroupOk() (*string, bool) {
	if o == nil || o.DeploymentGroup == nil {
		return nil, false
	}
	return o.DeploymentGroup, true
}

// HasDeploymentGroup returns a boolean if a field has been set.
func (o *Deployment) HasDeploymentGroup() bool {
	if o != nil && o.DeploymentGroup != nil {
		return true
	}

	return false
}

// SetDeploymentGroup gets a reference to the given string and assigns it to the DeploymentGroup field.
func (o *Deployment) SetDeploymentGroup(v string) {
	o.DeploymentGroup = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *Deployment) GetDefinition() ServiceDefinition {
	if o == nil || o.Definition == nil {
		var ret ServiceDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDefinitionOk() (*ServiceDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *Deployment) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given ServiceDefinition and assigns it to the Definition field.
func (o *Deployment) SetDefinition(v ServiceDefinition) {
	o.Definition = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Deployment) GetState() ServiceRevisionState {
	if o == nil || o.State == nil {
		var ret ServiceRevisionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetStateOk() (*ServiceRevisionState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Deployment) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ServiceRevisionState and assigns it to the State field.
func (o *Deployment) SetState(v ServiceRevisionState) {
	o.State = &v
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.AllocatedAt != nil {
		toSerialize["allocated_at"] = o.AllocatedAt
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.ChildId != nil {
		toSerialize["child_id"] = o.ChildId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.DeploymentGroup != nil {
		toSerialize["deployment_group"] = o.DeploymentGroup
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableDeployment struct {
	value *Deployment
	isSet bool
}

func (v NullableDeployment) Get() *Deployment {
	return v.value
}

func (v *NullableDeployment) Set(val *Deployment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployment(val *Deployment) *NullableDeployment {
	return &NullableDeployment{value: val, isSet: true}
}

func (v NullableDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


