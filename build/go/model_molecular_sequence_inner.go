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

// checks if the MolecularSequenceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularSequenceInner{}

// MolecularSequenceInner Raw data describing a biological sequence.
type MolecularSequenceInner struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A whole number
	Start *float32 `json:"start,omitempty"`
	// A whole number
	End *float32 `json:"end,omitempty"`
}

// NewMolecularSequenceInner instantiates a new MolecularSequenceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularSequenceInner() *MolecularSequenceInner {
	this := MolecularSequenceInner{}
	return &this
}

// NewMolecularSequenceInnerWithDefaults instantiates a new MolecularSequenceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularSequenceInnerWithDefaults() *MolecularSequenceInner {
	this := MolecularSequenceInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MolecularSequenceInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MolecularSequenceInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MolecularSequenceInner) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MolecularSequenceInner) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceInner) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MolecularSequenceInner) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MolecularSequenceInner) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MolecularSequenceInner) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceInner) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MolecularSequenceInner) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MolecularSequenceInner) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MolecularSequenceInner) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceInner) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MolecularSequenceInner) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *MolecularSequenceInner) SetStart(v float32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MolecularSequenceInner) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceInner) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MolecularSequenceInner) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *MolecularSequenceInner) SetEnd(v float32) {
	o.End = &v
}

func (o MolecularSequenceInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularSequenceInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableMolecularSequenceInner struct {
	value *MolecularSequenceInner
	isSet bool
}

func (v NullableMolecularSequenceInner) Get() *MolecularSequenceInner {
	return v.value
}

func (v *NullableMolecularSequenceInner) Set(val *MolecularSequenceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularSequenceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularSequenceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularSequenceInner(val *MolecularSequenceInner) *NullableMolecularSequenceInner {
	return &NullableMolecularSequenceInner{value: val, isSet: true}
}

func (v NullableMolecularSequenceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularSequenceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


