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
	"bytes"
	"fmt"
)

// checks if the TestScriptSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestScriptSetup{}

// TestScriptSetup A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScriptSetup struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Action would contain either an operation or an assertion.
	Action []TestScriptAction `json:"action"`
}

type _TestScriptSetup TestScriptSetup

// NewTestScriptSetup instantiates a new TestScriptSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestScriptSetup(action []TestScriptAction) *TestScriptSetup {
	this := TestScriptSetup{}
	this.Action = action
	return &this
}

// NewTestScriptSetupWithDefaults instantiates a new TestScriptSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestScriptSetupWithDefaults() *TestScriptSetup {
	this := TestScriptSetup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestScriptSetup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptSetup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestScriptSetup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestScriptSetup) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TestScriptSetup) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptSetup) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TestScriptSetup) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TestScriptSetup) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TestScriptSetup) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptSetup) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TestScriptSetup) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TestScriptSetup) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetAction returns the Action field value
func (o *TestScriptSetup) GetAction() []TestScriptAction {
	if o == nil {
		var ret []TestScriptAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TestScriptSetup) GetActionOk() ([]TestScriptAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action, true
}

// SetAction sets field value
func (o *TestScriptSetup) SetAction(v []TestScriptAction) {
	o.Action = v
}

func (o TestScriptSetup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestScriptSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.ModifierExtension) {
		toSerialize["modifierExtension"] = o.ModifierExtension
	}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *TestScriptSetup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varTestScriptSetup := _TestScriptSetup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestScriptSetup)

	if err != nil {
		return err
	}

	*o = TestScriptSetup(varTestScriptSetup)

	return err
}

type NullableTestScriptSetup struct {
	value *TestScriptSetup
	isSet bool
}

func (v NullableTestScriptSetup) Get() *TestScriptSetup {
	return v.value
}

func (v *NullableTestScriptSetup) Set(val *TestScriptSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableTestScriptSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableTestScriptSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestScriptSetup(val *TestScriptSetup) *NullableTestScriptSetup {
	return &NullableTestScriptSetup{value: val, isSet: true}
}

func (v NullableTestScriptSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestScriptSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

