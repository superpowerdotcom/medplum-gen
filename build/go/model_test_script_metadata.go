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

// checks if the TestScriptMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestScriptMetadata{}

// TestScriptMetadata A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScriptMetadata struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A link to the FHIR specification that this test is covering.
	Link []TestScriptLink `json:"link,omitempty"`
	// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
	Capability []TestScriptCapability `json:"capability"`
}

type _TestScriptMetadata TestScriptMetadata

// NewTestScriptMetadata instantiates a new TestScriptMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestScriptMetadata(capability []TestScriptCapability) *TestScriptMetadata {
	this := TestScriptMetadata{}
	this.Capability = capability
	return &this
}

// NewTestScriptMetadataWithDefaults instantiates a new TestScriptMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestScriptMetadataWithDefaults() *TestScriptMetadata {
	this := TestScriptMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestScriptMetadata) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptMetadata) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestScriptMetadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestScriptMetadata) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TestScriptMetadata) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptMetadata) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TestScriptMetadata) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TestScriptMetadata) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TestScriptMetadata) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptMetadata) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TestScriptMetadata) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TestScriptMetadata) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *TestScriptMetadata) GetLink() []TestScriptLink {
	if o == nil || IsNil(o.Link) {
		var ret []TestScriptLink
		return ret
	}
	return o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptMetadata) GetLinkOk() ([]TestScriptLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *TestScriptMetadata) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given []TestScriptLink and assigns it to the Link field.
func (o *TestScriptMetadata) SetLink(v []TestScriptLink) {
	o.Link = v
}

// GetCapability returns the Capability field value
func (o *TestScriptMetadata) GetCapability() []TestScriptCapability {
	if o == nil {
		var ret []TestScriptCapability
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *TestScriptMetadata) GetCapabilityOk() ([]TestScriptCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capability, true
}

// SetCapability sets field value
func (o *TestScriptMetadata) SetCapability(v []TestScriptCapability) {
	o.Capability = v
}

func (o TestScriptMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestScriptMetadata) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["capability"] = o.Capability
	return toSerialize, nil
}

func (o *TestScriptMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capability",
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

	varTestScriptMetadata := _TestScriptMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestScriptMetadata)

	if err != nil {
		return err
	}

	*o = TestScriptMetadata(varTestScriptMetadata)

	return err
}

type NullableTestScriptMetadata struct {
	value *TestScriptMetadata
	isSet bool
}

func (v NullableTestScriptMetadata) Get() *TestScriptMetadata {
	return v.value
}

func (v *NullableTestScriptMetadata) Set(val *TestScriptMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTestScriptMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTestScriptMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestScriptMetadata(val *TestScriptMetadata) *NullableTestScriptMetadata {
	return &NullableTestScriptMetadata{value: val, isSet: true}
}

func (v NullableTestScriptMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestScriptMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


