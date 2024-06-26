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
	"fmt"
)

// checks if the ProjectMembershipAccessParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMembershipAccessParameter{}

// ProjectMembershipAccessParameter User options that control the display of the application.
type ProjectMembershipAccessParameter struct {
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Name string `json:"name"`
	// A sequence of Unicode characters
	ValueString *string `json:"valueString,omitempty"`
	// Value of the parameter - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueReference *Reference `json:"valueReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectMembershipAccessParameter ProjectMembershipAccessParameter

// NewProjectMembershipAccessParameter instantiates a new ProjectMembershipAccessParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMembershipAccessParameter(name string) *ProjectMembershipAccessParameter {
	this := ProjectMembershipAccessParameter{}
	this.Name = name
	return &this
}

// NewProjectMembershipAccessParameterWithDefaults instantiates a new ProjectMembershipAccessParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMembershipAccessParameterWithDefaults() *ProjectMembershipAccessParameter {
	this := ProjectMembershipAccessParameter{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectMembershipAccessParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMembershipAccessParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMembershipAccessParameter) SetName(v string) {
	o.Name = v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *ProjectMembershipAccessParameter) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembershipAccessParameter) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *ProjectMembershipAccessParameter) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *ProjectMembershipAccessParameter) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueReference returns the ValueReference field value if set, zero value otherwise.
func (o *ProjectMembershipAccessParameter) GetValueReference() Reference {
	if o == nil || IsNil(o.ValueReference) {
		var ret Reference
		return ret
	}
	return *o.ValueReference
}

// GetValueReferenceOk returns a tuple with the ValueReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembershipAccessParameter) GetValueReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ValueReference) {
		return nil, false
	}
	return o.ValueReference, true
}

// HasValueReference returns a boolean if a field has been set.
func (o *ProjectMembershipAccessParameter) HasValueReference() bool {
	if o != nil && !IsNil(o.ValueReference) {
		return true
	}

	return false
}

// SetValueReference gets a reference to the given Reference and assigns it to the ValueReference field.
func (o *ProjectMembershipAccessParameter) SetValueReference(v Reference) {
	o.ValueReference = &v
}

func (o ProjectMembershipAccessParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMembershipAccessParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ValueString) {
		toSerialize["valueString"] = o.ValueString
	}
	if !IsNil(o.ValueReference) {
		toSerialize["valueReference"] = o.ValueReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectMembershipAccessParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectMembershipAccessParameter := _ProjectMembershipAccessParameter{}

	err = json.Unmarshal(data, &varProjectMembershipAccessParameter)

	if err != nil {
		return err
	}

	*o = ProjectMembershipAccessParameter(varProjectMembershipAccessParameter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "valueString")
		delete(additionalProperties, "valueReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectMembershipAccessParameter struct {
	value *ProjectMembershipAccessParameter
	isSet bool
}

func (v NullableProjectMembershipAccessParameter) Get() *ProjectMembershipAccessParameter {
	return v.value
}

func (v *NullableProjectMembershipAccessParameter) Set(val *ProjectMembershipAccessParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMembershipAccessParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMembershipAccessParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMembershipAccessParameter(val *ProjectMembershipAccessParameter) *NullableProjectMembershipAccessParameter {
	return &NullableProjectMembershipAccessParameter{value: val, isSet: true}
}

func (v NullableProjectMembershipAccessParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMembershipAccessParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


