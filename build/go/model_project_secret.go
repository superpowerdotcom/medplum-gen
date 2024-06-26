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

// checks if the ProjectSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSecret{}

// ProjectSecret Secure environment variable that can be used to store secrets for bots.
type ProjectSecret struct {
	// A sequence of Unicode characters
	Name string `json:"name"`
	// A sequence of Unicode characters
	ValueString *string `json:"valueString,omitempty"`
	// Value of \"true\" or \"false\"
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// A rational number with implicit precision
	ValueDecimal *float32 `json:"valueDecimal,omitempty"`
	// A whole number
	ValueInteger *float32 `json:"valueInteger,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectSecret ProjectSecret

// NewProjectSecret instantiates a new ProjectSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSecret(name string) *ProjectSecret {
	this := ProjectSecret{}
	this.Name = name
	return &this
}

// NewProjectSecretWithDefaults instantiates a new ProjectSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSecretWithDefaults() *ProjectSecret {
	this := ProjectSecret{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectSecret) SetName(v string) {
	o.Name = v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *ProjectSecret) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecret) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *ProjectSecret) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *ProjectSecret) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueBoolean returns the ValueBoolean field value if set, zero value otherwise.
func (o *ProjectSecret) GetValueBoolean() bool {
	if o == nil || IsNil(o.ValueBoolean) {
		var ret bool
		return ret
	}
	return *o.ValueBoolean
}

// GetValueBooleanOk returns a tuple with the ValueBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecret) GetValueBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueBoolean) {
		return nil, false
	}
	return o.ValueBoolean, true
}

// HasValueBoolean returns a boolean if a field has been set.
func (o *ProjectSecret) HasValueBoolean() bool {
	if o != nil && !IsNil(o.ValueBoolean) {
		return true
	}

	return false
}

// SetValueBoolean gets a reference to the given bool and assigns it to the ValueBoolean field.
func (o *ProjectSecret) SetValueBoolean(v bool) {
	o.ValueBoolean = &v
}

// GetValueDecimal returns the ValueDecimal field value if set, zero value otherwise.
func (o *ProjectSecret) GetValueDecimal() float32 {
	if o == nil || IsNil(o.ValueDecimal) {
		var ret float32
		return ret
	}
	return *o.ValueDecimal
}

// GetValueDecimalOk returns a tuple with the ValueDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecret) GetValueDecimalOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueDecimal) {
		return nil, false
	}
	return o.ValueDecimal, true
}

// HasValueDecimal returns a boolean if a field has been set.
func (o *ProjectSecret) HasValueDecimal() bool {
	if o != nil && !IsNil(o.ValueDecimal) {
		return true
	}

	return false
}

// SetValueDecimal gets a reference to the given float32 and assigns it to the ValueDecimal field.
func (o *ProjectSecret) SetValueDecimal(v float32) {
	o.ValueDecimal = &v
}

// GetValueInteger returns the ValueInteger field value if set, zero value otherwise.
func (o *ProjectSecret) GetValueInteger() float32 {
	if o == nil || IsNil(o.ValueInteger) {
		var ret float32
		return ret
	}
	return *o.ValueInteger
}

// GetValueIntegerOk returns a tuple with the ValueInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecret) GetValueIntegerOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueInteger) {
		return nil, false
	}
	return o.ValueInteger, true
}

// HasValueInteger returns a boolean if a field has been set.
func (o *ProjectSecret) HasValueInteger() bool {
	if o != nil && !IsNil(o.ValueInteger) {
		return true
	}

	return false
}

// SetValueInteger gets a reference to the given float32 and assigns it to the ValueInteger field.
func (o *ProjectSecret) SetValueInteger(v float32) {
	o.ValueInteger = &v
}

func (o ProjectSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ValueString) {
		toSerialize["valueString"] = o.ValueString
	}
	if !IsNil(o.ValueBoolean) {
		toSerialize["valueBoolean"] = o.ValueBoolean
	}
	if !IsNil(o.ValueDecimal) {
		toSerialize["valueDecimal"] = o.ValueDecimal
	}
	if !IsNil(o.ValueInteger) {
		toSerialize["valueInteger"] = o.ValueInteger
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectSecret) UnmarshalJSON(data []byte) (err error) {
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

	varProjectSecret := _ProjectSecret{}

	err = json.Unmarshal(data, &varProjectSecret)

	if err != nil {
		return err
	}

	*o = ProjectSecret(varProjectSecret)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "valueString")
		delete(additionalProperties, "valueBoolean")
		delete(additionalProperties, "valueDecimal")
		delete(additionalProperties, "valueInteger")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectSecret struct {
	value *ProjectSecret
	isSet bool
}

func (v NullableProjectSecret) Get() *ProjectSecret {
	return v.value
}

func (v *NullableProjectSecret) Set(val *ProjectSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSecret(val *ProjectSecret) *NullableProjectSecret {
	return &NullableProjectSecret{value: val, isSet: true}
}

func (v NullableProjectSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


