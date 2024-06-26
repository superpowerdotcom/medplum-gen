/*
Medplum - OpenAPI 3.0

Medplum OpenAPI 3.0 specification.  Learn more about Medplum at [https://www.medplum.com](https://www.medplum.com).

API version: 1.0.5
Contact: hello@medplum.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package medplum

import (
	"encoding/json"
)

// checks if the AccessPolicyResourceWriteCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyResourceWriteCriteria{}

// AccessPolicyResourceWriteCriteria Invariants that must be satisfied for the resource to be written.
type AccessPolicyResourceWriteCriteria struct {
	// A sequence of Unicode characters
	Pre *string `json:"pre,omitempty"`
	// A sequence of Unicode characters
	Post *string `json:"post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyResourceWriteCriteria AccessPolicyResourceWriteCriteria

// NewAccessPolicyResourceWriteCriteria instantiates a new AccessPolicyResourceWriteCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyResourceWriteCriteria() *AccessPolicyResourceWriteCriteria {
	this := AccessPolicyResourceWriteCriteria{}
	return &this
}

// NewAccessPolicyResourceWriteCriteriaWithDefaults instantiates a new AccessPolicyResourceWriteCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyResourceWriteCriteriaWithDefaults() *AccessPolicyResourceWriteCriteria {
	this := AccessPolicyResourceWriteCriteria{}
	return &this
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *AccessPolicyResourceWriteCriteria) GetPre() string {
	if o == nil || IsNil(o.Pre) {
		var ret string
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyResourceWriteCriteria) GetPreOk() (*string, bool) {
	if o == nil || IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *AccessPolicyResourceWriteCriteria) HasPre() bool {
	if o != nil && !IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given string and assigns it to the Pre field.
func (o *AccessPolicyResourceWriteCriteria) SetPre(v string) {
	o.Pre = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *AccessPolicyResourceWriteCriteria) GetPost() string {
	if o == nil || IsNil(o.Post) {
		var ret string
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyResourceWriteCriteria) GetPostOk() (*string, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *AccessPolicyResourceWriteCriteria) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given string and assigns it to the Post field.
func (o *AccessPolicyResourceWriteCriteria) SetPost(v string) {
	o.Post = &v
}

func (o AccessPolicyResourceWriteCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyResourceWriteCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pre) {
		toSerialize["pre"] = o.Pre
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyResourceWriteCriteria) UnmarshalJSON(data []byte) (err error) {
	varAccessPolicyResourceWriteCriteria := _AccessPolicyResourceWriteCriteria{}

	err = json.Unmarshal(data, &varAccessPolicyResourceWriteCriteria)

	if err != nil {
		return err
	}

	*o = AccessPolicyResourceWriteCriteria(varAccessPolicyResourceWriteCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pre")
		delete(additionalProperties, "post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyResourceWriteCriteria struct {
	value *AccessPolicyResourceWriteCriteria
	isSet bool
}

func (v NullableAccessPolicyResourceWriteCriteria) Get() *AccessPolicyResourceWriteCriteria {
	return v.value
}

func (v *NullableAccessPolicyResourceWriteCriteria) Set(val *AccessPolicyResourceWriteCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyResourceWriteCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyResourceWriteCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyResourceWriteCriteria(val *AccessPolicyResourceWriteCriteria) *NullableAccessPolicyResourceWriteCriteria {
	return &NullableAccessPolicyResourceWriteCriteria{value: val, isSet: true}
}

func (v NullableAccessPolicyResourceWriteCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyResourceWriteCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


