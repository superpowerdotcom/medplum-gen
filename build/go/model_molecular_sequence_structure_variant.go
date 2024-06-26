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

// checks if the MolecularSequenceStructureVariant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularSequenceStructureVariant{}

// MolecularSequenceStructureVariant Raw data describing a biological sequence.
type MolecularSequenceStructureVariant struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Information about chromosome structure variation DNA change type.
	VariantType *CodeableConcept `json:"variantType,omitempty"`
	// Value of \"true\" or \"false\"
	Exact *bool `json:"exact,omitempty"`
	// A whole number
	Length *float32 `json:"length,omitempty"`
	// Structural variant outer.
	Outer *MolecularSequenceOuter `json:"outer,omitempty"`
	// Structural variant inner.
	Inner *MolecularSequenceInner `json:"inner,omitempty"`
}

// NewMolecularSequenceStructureVariant instantiates a new MolecularSequenceStructureVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularSequenceStructureVariant() *MolecularSequenceStructureVariant {
	this := MolecularSequenceStructureVariant{}
	return &this
}

// NewMolecularSequenceStructureVariantWithDefaults instantiates a new MolecularSequenceStructureVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularSequenceStructureVariantWithDefaults() *MolecularSequenceStructureVariant {
	this := MolecularSequenceStructureVariant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MolecularSequenceStructureVariant) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MolecularSequenceStructureVariant) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MolecularSequenceStructureVariant) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetVariantType returns the VariantType field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetVariantType() CodeableConcept {
	if o == nil || IsNil(o.VariantType) {
		var ret CodeableConcept
		return ret
	}
	return *o.VariantType
}

// GetVariantTypeOk returns a tuple with the VariantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetVariantTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.VariantType) {
		return nil, false
	}
	return o.VariantType, true
}

// HasVariantType returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasVariantType() bool {
	if o != nil && !IsNil(o.VariantType) {
		return true
	}

	return false
}

// SetVariantType gets a reference to the given CodeableConcept and assigns it to the VariantType field.
func (o *MolecularSequenceStructureVariant) SetVariantType(v CodeableConcept) {
	o.VariantType = &v
}

// GetExact returns the Exact field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetExact() bool {
	if o == nil || IsNil(o.Exact) {
		var ret bool
		return ret
	}
	return *o.Exact
}

// GetExactOk returns a tuple with the Exact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetExactOk() (*bool, bool) {
	if o == nil || IsNil(o.Exact) {
		return nil, false
	}
	return o.Exact, true
}

// HasExact returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasExact() bool {
	if o != nil && !IsNil(o.Exact) {
		return true
	}

	return false
}

// SetExact gets a reference to the given bool and assigns it to the Exact field.
func (o *MolecularSequenceStructureVariant) SetExact(v bool) {
	o.Exact = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *MolecularSequenceStructureVariant) SetLength(v float32) {
	o.Length = &v
}

// GetOuter returns the Outer field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetOuter() MolecularSequenceOuter {
	if o == nil || IsNil(o.Outer) {
		var ret MolecularSequenceOuter
		return ret
	}
	return *o.Outer
}

// GetOuterOk returns a tuple with the Outer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetOuterOk() (*MolecularSequenceOuter, bool) {
	if o == nil || IsNil(o.Outer) {
		return nil, false
	}
	return o.Outer, true
}

// HasOuter returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasOuter() bool {
	if o != nil && !IsNil(o.Outer) {
		return true
	}

	return false
}

// SetOuter gets a reference to the given MolecularSequenceOuter and assigns it to the Outer field.
func (o *MolecularSequenceStructureVariant) SetOuter(v MolecularSequenceOuter) {
	o.Outer = &v
}

// GetInner returns the Inner field value if set, zero value otherwise.
func (o *MolecularSequenceStructureVariant) GetInner() MolecularSequenceInner {
	if o == nil || IsNil(o.Inner) {
		var ret MolecularSequenceInner
		return ret
	}
	return *o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceStructureVariant) GetInnerOk() (*MolecularSequenceInner, bool) {
	if o == nil || IsNil(o.Inner) {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *MolecularSequenceStructureVariant) HasInner() bool {
	if o != nil && !IsNil(o.Inner) {
		return true
	}

	return false
}

// SetInner gets a reference to the given MolecularSequenceInner and assigns it to the Inner field.
func (o *MolecularSequenceStructureVariant) SetInner(v MolecularSequenceInner) {
	o.Inner = &v
}

func (o MolecularSequenceStructureVariant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularSequenceStructureVariant) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.VariantType) {
		toSerialize["variantType"] = o.VariantType
	}
	if !IsNil(o.Exact) {
		toSerialize["exact"] = o.Exact
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Outer) {
		toSerialize["outer"] = o.Outer
	}
	if !IsNil(o.Inner) {
		toSerialize["inner"] = o.Inner
	}
	return toSerialize, nil
}

type NullableMolecularSequenceStructureVariant struct {
	value *MolecularSequenceStructureVariant
	isSet bool
}

func (v NullableMolecularSequenceStructureVariant) Get() *MolecularSequenceStructureVariant {
	return v.value
}

func (v *NullableMolecularSequenceStructureVariant) Set(val *MolecularSequenceStructureVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularSequenceStructureVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularSequenceStructureVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularSequenceStructureVariant(val *MolecularSequenceStructureVariant) *NullableMolecularSequenceStructureVariant {
	return &NullableMolecularSequenceStructureVariant{value: val, isSet: true}
}

func (v NullableMolecularSequenceStructureVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularSequenceStructureVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

