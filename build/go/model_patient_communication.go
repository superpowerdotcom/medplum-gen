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

// checks if the PatientCommunication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatientCommunication{}

// PatientCommunication Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type PatientCommunication struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. \"en\" for English, or \"en-US\" for American English versus \"en-EN\" for England English.
	Language CodeableConcept `json:"language"`
	// Value of \"true\" or \"false\"
	Preferred *bool `json:"preferred,omitempty"`
}

type _PatientCommunication PatientCommunication

// NewPatientCommunication instantiates a new PatientCommunication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatientCommunication(language CodeableConcept) *PatientCommunication {
	this := PatientCommunication{}
	this.Language = language
	return &this
}

// NewPatientCommunicationWithDefaults instantiates a new PatientCommunication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatientCommunicationWithDefaults() *PatientCommunication {
	this := PatientCommunication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatientCommunication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientCommunication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatientCommunication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatientCommunication) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *PatientCommunication) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientCommunication) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *PatientCommunication) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *PatientCommunication) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *PatientCommunication) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientCommunication) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *PatientCommunication) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *PatientCommunication) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetLanguage returns the Language field value
func (o *PatientCommunication) GetLanguage() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *PatientCommunication) GetLanguageOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *PatientCommunication) SetLanguage(v CodeableConcept) {
	o.Language = v
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *PatientCommunication) GetPreferred() bool {
	if o == nil || IsNil(o.Preferred) {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatientCommunication) GetPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *PatientCommunication) HasPreferred() bool {
	if o != nil && !IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *PatientCommunication) SetPreferred(v bool) {
	o.Preferred = &v
}

func (o PatientCommunication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatientCommunication) ToMap() (map[string]interface{}, error) {
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
	toSerialize["language"] = o.Language
	if !IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	return toSerialize, nil
}

func (o *PatientCommunication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
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

	varPatientCommunication := _PatientCommunication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatientCommunication)

	if err != nil {
		return err
	}

	*o = PatientCommunication(varPatientCommunication)

	return err
}

type NullablePatientCommunication struct {
	value *PatientCommunication
	isSet bool
}

func (v NullablePatientCommunication) Get() *PatientCommunication {
	return v.value
}

func (v *NullablePatientCommunication) Set(val *PatientCommunication) {
	v.value = val
	v.isSet = true
}

func (v NullablePatientCommunication) IsSet() bool {
	return v.isSet
}

func (v *NullablePatientCommunication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatientCommunication(val *PatientCommunication) *NullablePatientCommunication {
	return &NullablePatientCommunication{value: val, isSet: true}
}

func (v NullablePatientCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatientCommunication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


