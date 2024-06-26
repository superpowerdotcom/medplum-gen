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

// checks if the MolecularSequenceVariant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularSequenceVariant{}

// MolecularSequenceVariant Raw data describing a biological sequence.
type MolecularSequenceVariant struct {
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
	// A sequence of Unicode characters
	ObservedAllele *string `json:"observedAllele,omitempty"`
	// A sequence of Unicode characters
	ReferenceAllele *string `json:"referenceAllele,omitempty"`
	// A sequence of Unicode characters
	Cigar *string `json:"cigar,omitempty"`
	// A pointer to an Observation containing variant information.
	VariantPointer *Reference `json:"variantPointer,omitempty"`
}

// NewMolecularSequenceVariant instantiates a new MolecularSequenceVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularSequenceVariant() *MolecularSequenceVariant {
	this := MolecularSequenceVariant{}
	return &this
}

// NewMolecularSequenceVariantWithDefaults instantiates a new MolecularSequenceVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularSequenceVariantWithDefaults() *MolecularSequenceVariant {
	this := MolecularSequenceVariant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MolecularSequenceVariant) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MolecularSequenceVariant) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MolecularSequenceVariant) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *MolecularSequenceVariant) SetStart(v float32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *MolecularSequenceVariant) SetEnd(v float32) {
	o.End = &v
}

// GetObservedAllele returns the ObservedAllele field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetObservedAllele() string {
	if o == nil || IsNil(o.ObservedAllele) {
		var ret string
		return ret
	}
	return *o.ObservedAllele
}

// GetObservedAlleleOk returns a tuple with the ObservedAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetObservedAlleleOk() (*string, bool) {
	if o == nil || IsNil(o.ObservedAllele) {
		return nil, false
	}
	return o.ObservedAllele, true
}

// HasObservedAllele returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasObservedAllele() bool {
	if o != nil && !IsNil(o.ObservedAllele) {
		return true
	}

	return false
}

// SetObservedAllele gets a reference to the given string and assigns it to the ObservedAllele field.
func (o *MolecularSequenceVariant) SetObservedAllele(v string) {
	o.ObservedAllele = &v
}

// GetReferenceAllele returns the ReferenceAllele field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetReferenceAllele() string {
	if o == nil || IsNil(o.ReferenceAllele) {
		var ret string
		return ret
	}
	return *o.ReferenceAllele
}

// GetReferenceAlleleOk returns a tuple with the ReferenceAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetReferenceAlleleOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceAllele) {
		return nil, false
	}
	return o.ReferenceAllele, true
}

// HasReferenceAllele returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasReferenceAllele() bool {
	if o != nil && !IsNil(o.ReferenceAllele) {
		return true
	}

	return false
}

// SetReferenceAllele gets a reference to the given string and assigns it to the ReferenceAllele field.
func (o *MolecularSequenceVariant) SetReferenceAllele(v string) {
	o.ReferenceAllele = &v
}

// GetCigar returns the Cigar field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetCigar() string {
	if o == nil || IsNil(o.Cigar) {
		var ret string
		return ret
	}
	return *o.Cigar
}

// GetCigarOk returns a tuple with the Cigar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetCigarOk() (*string, bool) {
	if o == nil || IsNil(o.Cigar) {
		return nil, false
	}
	return o.Cigar, true
}

// HasCigar returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasCigar() bool {
	if o != nil && !IsNil(o.Cigar) {
		return true
	}

	return false
}

// SetCigar gets a reference to the given string and assigns it to the Cigar field.
func (o *MolecularSequenceVariant) SetCigar(v string) {
	o.Cigar = &v
}

// GetVariantPointer returns the VariantPointer field value if set, zero value otherwise.
func (o *MolecularSequenceVariant) GetVariantPointer() Reference {
	if o == nil || IsNil(o.VariantPointer) {
		var ret Reference
		return ret
	}
	return *o.VariantPointer
}

// GetVariantPointerOk returns a tuple with the VariantPointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceVariant) GetVariantPointerOk() (*Reference, bool) {
	if o == nil || IsNil(o.VariantPointer) {
		return nil, false
	}
	return o.VariantPointer, true
}

// HasVariantPointer returns a boolean if a field has been set.
func (o *MolecularSequenceVariant) HasVariantPointer() bool {
	if o != nil && !IsNil(o.VariantPointer) {
		return true
	}

	return false
}

// SetVariantPointer gets a reference to the given Reference and assigns it to the VariantPointer field.
func (o *MolecularSequenceVariant) SetVariantPointer(v Reference) {
	o.VariantPointer = &v
}

func (o MolecularSequenceVariant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularSequenceVariant) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ObservedAllele) {
		toSerialize["observedAllele"] = o.ObservedAllele
	}
	if !IsNil(o.ReferenceAllele) {
		toSerialize["referenceAllele"] = o.ReferenceAllele
	}
	if !IsNil(o.Cigar) {
		toSerialize["cigar"] = o.Cigar
	}
	if !IsNil(o.VariantPointer) {
		toSerialize["variantPointer"] = o.VariantPointer
	}
	return toSerialize, nil
}

type NullableMolecularSequenceVariant struct {
	value *MolecularSequenceVariant
	isSet bool
}

func (v NullableMolecularSequenceVariant) Get() *MolecularSequenceVariant {
	return v.value
}

func (v *NullableMolecularSequenceVariant) Set(val *MolecularSequenceVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularSequenceVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularSequenceVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularSequenceVariant(val *MolecularSequenceVariant) *NullableMolecularSequenceVariant {
	return &NullableMolecularSequenceVariant{value: val, isSet: true}
}

func (v NullableMolecularSequenceVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularSequenceVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

