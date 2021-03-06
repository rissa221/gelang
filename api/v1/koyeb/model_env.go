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

// Env struct for Env
type Env struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	ValueFromSecret *string `json:"value_from_secret,omitempty"`
}

// NewEnv instantiates a new Env object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnv() *Env {
	this := Env{}
	return &this
}

// NewEnvWithDefaults instantiates a new Env object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvWithDefaults() *Env {
	this := Env{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Env) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Env) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Env) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Env) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Env) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Env) SetValue(v string) {
	o.Value = &v
}

// GetValueFromSecret returns the ValueFromSecret field value if set, zero value otherwise.
func (o *Env) GetValueFromSecret() string {
	if o == nil || o.ValueFromSecret == nil {
		var ret string
		return ret
	}
	return *o.ValueFromSecret
}

// GetValueFromSecretOk returns a tuple with the ValueFromSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Env) GetValueFromSecretOk() (*string, bool) {
	if o == nil || o.ValueFromSecret == nil {
		return nil, false
	}
	return o.ValueFromSecret, true
}

// HasValueFromSecret returns a boolean if a field has been set.
func (o *Env) HasValueFromSecret() bool {
	if o != nil && o.ValueFromSecret != nil {
		return true
	}

	return false
}

// SetValueFromSecret gets a reference to the given string and assigns it to the ValueFromSecret field.
func (o *Env) SetValueFromSecret(v string) {
	o.ValueFromSecret = &v
}

func (o Env) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueFromSecret != nil {
		toSerialize["value_from_secret"] = o.ValueFromSecret
	}
	return json.Marshal(toSerialize)
}

type NullableEnv struct {
	value *Env
	isSet bool
}

func (v NullableEnv) Get() *Env {
	return v.value
}

func (v *NullableEnv) Set(val *Env) {
	v.value = val
	v.isSet = true
}

func (v NullableEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnv(val *Env) *NullableEnv {
	return &NullableEnv{value: val, isSet: true}
}

func (v NullableEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


