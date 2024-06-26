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

// checks if the EncounterClassHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncounterClassHistory{}

// EncounterClassHistory An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type EncounterClassHistory struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// inpatient | outpatient | ambulatory | emergency +.
	Class Coding `json:"class"`
	// The time that the episode was in the specified class.
	Period Period `json:"period"`
}

type _EncounterClassHistory EncounterClassHistory

// NewEncounterClassHistory instantiates a new EncounterClassHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncounterClassHistory(class Coding, period Period) *EncounterClassHistory {
	this := EncounterClassHistory{}
	this.Class = class
	this.Period = period
	return &this
}

// NewEncounterClassHistoryWithDefaults instantiates a new EncounterClassHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncounterClassHistoryWithDefaults() *EncounterClassHistory {
	this := EncounterClassHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EncounterClassHistory) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterClassHistory) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EncounterClassHistory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EncounterClassHistory) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *EncounterClassHistory) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterClassHistory) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *EncounterClassHistory) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *EncounterClassHistory) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *EncounterClassHistory) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterClassHistory) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *EncounterClassHistory) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *EncounterClassHistory) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetClass returns the Class field value
func (o *EncounterClassHistory) GetClass() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *EncounterClassHistory) GetClassOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *EncounterClassHistory) SetClass(v Coding) {
	o.Class = v
}

// GetPeriod returns the Period field value
func (o *EncounterClassHistory) GetPeriod() Period {
	if o == nil {
		var ret Period
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *EncounterClassHistory) GetPeriodOk() (*Period, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *EncounterClassHistory) SetPeriod(v Period) {
	o.Period = v
}

func (o EncounterClassHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncounterClassHistory) ToMap() (map[string]interface{}, error) {
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
	toSerialize["class"] = o.Class
	toSerialize["period"] = o.Period
	return toSerialize, nil
}

func (o *EncounterClassHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"class",
		"period",
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

	varEncounterClassHistory := _EncounterClassHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEncounterClassHistory)

	if err != nil {
		return err
	}

	*o = EncounterClassHistory(varEncounterClassHistory)

	return err
}

type NullableEncounterClassHistory struct {
	value *EncounterClassHistory
	isSet bool
}

func (v NullableEncounterClassHistory) Get() *EncounterClassHistory {
	return v.value
}

func (v *NullableEncounterClassHistory) Set(val *EncounterClassHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableEncounterClassHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableEncounterClassHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncounterClassHistory(val *EncounterClassHistory) *NullableEncounterClassHistory {
	return &NullableEncounterClassHistory{value: val, isSet: true}
}

func (v NullableEncounterClassHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncounterClassHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


