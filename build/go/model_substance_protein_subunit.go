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

// checks if the SubstanceProteinSubunit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstanceProteinSubunit{}

// SubstanceProteinSubunit A SubstanceProtein is defined as a single unit of a linear amino acid sequence, or a combination of subunits that are either covalently linked or have a defined invariant stoichiometric relationship. This includes all synthetic, recombinant and purified SubstanceProteins of defined sequence, whether the use is therapeutic or prophylactic. This set of elements will be used to describe albumins, coagulation factors, cytokines, growth factors, peptide/SubstanceProtein hormones, enzymes, toxins, toxoids, recombinant vaccines, and immunomodulators.
type SubstanceProteinSubunit struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A whole number
	Subunit *float32 `json:"subunit,omitempty"`
	// A sequence of Unicode characters
	Sequence *string `json:"sequence,omitempty"`
	// A whole number
	Length *float32 `json:"length,omitempty"`
	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end using standard single-letter amino acid codes. Uppercase shall be used for L-amino acids and lowercase for D-amino acids. Transcribed SubstanceProteins will always be described using the translated sequence; for synthetic peptide containing amino acids that are not represented with a single letter code an X should be used within the sequence. The modified amino acids will be distinguished by their position in the sequence.
	SequenceAttachment *Attachment `json:"sequenceAttachment,omitempty"`
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID.
	NTerminalModificationId *Identifier `json:"nTerminalModificationId,omitempty"`
	// A sequence of Unicode characters
	NTerminalModification *string `json:"nTerminalModification,omitempty"`
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID.
	CTerminalModificationId *Identifier `json:"cTerminalModificationId,omitempty"`
	// A sequence of Unicode characters
	CTerminalModification *string `json:"cTerminalModification,omitempty"`
}

// NewSubstanceProteinSubunit instantiates a new SubstanceProteinSubunit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstanceProteinSubunit() *SubstanceProteinSubunit {
	this := SubstanceProteinSubunit{}
	return &this
}

// NewSubstanceProteinSubunitWithDefaults instantiates a new SubstanceProteinSubunit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceProteinSubunitWithDefaults() *SubstanceProteinSubunit {
	this := SubstanceProteinSubunit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstanceProteinSubunit) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstanceProteinSubunit) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstanceProteinSubunit) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSubunit returns the Subunit field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetSubunit() float32 {
	if o == nil || IsNil(o.Subunit) {
		var ret float32
		return ret
	}
	return *o.Subunit
}

// GetSubunitOk returns a tuple with the Subunit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetSubunitOk() (*float32, bool) {
	if o == nil || IsNil(o.Subunit) {
		return nil, false
	}
	return o.Subunit, true
}

// HasSubunit returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasSubunit() bool {
	if o != nil && !IsNil(o.Subunit) {
		return true
	}

	return false
}

// SetSubunit gets a reference to the given float32 and assigns it to the Subunit field.
func (o *SubstanceProteinSubunit) SetSubunit(v float32) {
	o.Subunit = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetSequence() string {
	if o == nil || IsNil(o.Sequence) {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetSequenceOk() (*string, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *SubstanceProteinSubunit) SetSequence(v string) {
	o.Sequence = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *SubstanceProteinSubunit) SetLength(v float32) {
	o.Length = &v
}

// GetSequenceAttachment returns the SequenceAttachment field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetSequenceAttachment() Attachment {
	if o == nil || IsNil(o.SequenceAttachment) {
		var ret Attachment
		return ret
	}
	return *o.SequenceAttachment
}

// GetSequenceAttachmentOk returns a tuple with the SequenceAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetSequenceAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.SequenceAttachment) {
		return nil, false
	}
	return o.SequenceAttachment, true
}

// HasSequenceAttachment returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasSequenceAttachment() bool {
	if o != nil && !IsNil(o.SequenceAttachment) {
		return true
	}

	return false
}

// SetSequenceAttachment gets a reference to the given Attachment and assigns it to the SequenceAttachment field.
func (o *SubstanceProteinSubunit) SetSequenceAttachment(v Attachment) {
	o.SequenceAttachment = &v
}

// GetNTerminalModificationId returns the NTerminalModificationId field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetNTerminalModificationId() Identifier {
	if o == nil || IsNil(o.NTerminalModificationId) {
		var ret Identifier
		return ret
	}
	return *o.NTerminalModificationId
}

// GetNTerminalModificationIdOk returns a tuple with the NTerminalModificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetNTerminalModificationIdOk() (*Identifier, bool) {
	if o == nil || IsNil(o.NTerminalModificationId) {
		return nil, false
	}
	return o.NTerminalModificationId, true
}

// HasNTerminalModificationId returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasNTerminalModificationId() bool {
	if o != nil && !IsNil(o.NTerminalModificationId) {
		return true
	}

	return false
}

// SetNTerminalModificationId gets a reference to the given Identifier and assigns it to the NTerminalModificationId field.
func (o *SubstanceProteinSubunit) SetNTerminalModificationId(v Identifier) {
	o.NTerminalModificationId = &v
}

// GetNTerminalModification returns the NTerminalModification field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetNTerminalModification() string {
	if o == nil || IsNil(o.NTerminalModification) {
		var ret string
		return ret
	}
	return *o.NTerminalModification
}

// GetNTerminalModificationOk returns a tuple with the NTerminalModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetNTerminalModificationOk() (*string, bool) {
	if o == nil || IsNil(o.NTerminalModification) {
		return nil, false
	}
	return o.NTerminalModification, true
}

// HasNTerminalModification returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasNTerminalModification() bool {
	if o != nil && !IsNil(o.NTerminalModification) {
		return true
	}

	return false
}

// SetNTerminalModification gets a reference to the given string and assigns it to the NTerminalModification field.
func (o *SubstanceProteinSubunit) SetNTerminalModification(v string) {
	o.NTerminalModification = &v
}

// GetCTerminalModificationId returns the CTerminalModificationId field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetCTerminalModificationId() Identifier {
	if o == nil || IsNil(o.CTerminalModificationId) {
		var ret Identifier
		return ret
	}
	return *o.CTerminalModificationId
}

// GetCTerminalModificationIdOk returns a tuple with the CTerminalModificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetCTerminalModificationIdOk() (*Identifier, bool) {
	if o == nil || IsNil(o.CTerminalModificationId) {
		return nil, false
	}
	return o.CTerminalModificationId, true
}

// HasCTerminalModificationId returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasCTerminalModificationId() bool {
	if o != nil && !IsNil(o.CTerminalModificationId) {
		return true
	}

	return false
}

// SetCTerminalModificationId gets a reference to the given Identifier and assigns it to the CTerminalModificationId field.
func (o *SubstanceProteinSubunit) SetCTerminalModificationId(v Identifier) {
	o.CTerminalModificationId = &v
}

// GetCTerminalModification returns the CTerminalModification field value if set, zero value otherwise.
func (o *SubstanceProteinSubunit) GetCTerminalModification() string {
	if o == nil || IsNil(o.CTerminalModification) {
		var ret string
		return ret
	}
	return *o.CTerminalModification
}

// GetCTerminalModificationOk returns a tuple with the CTerminalModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceProteinSubunit) GetCTerminalModificationOk() (*string, bool) {
	if o == nil || IsNil(o.CTerminalModification) {
		return nil, false
	}
	return o.CTerminalModification, true
}

// HasCTerminalModification returns a boolean if a field has been set.
func (o *SubstanceProteinSubunit) HasCTerminalModification() bool {
	if o != nil && !IsNil(o.CTerminalModification) {
		return true
	}

	return false
}

// SetCTerminalModification gets a reference to the given string and assigns it to the CTerminalModification field.
func (o *SubstanceProteinSubunit) SetCTerminalModification(v string) {
	o.CTerminalModification = &v
}

func (o SubstanceProteinSubunit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstanceProteinSubunit) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Subunit) {
		toSerialize["subunit"] = o.Subunit
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.SequenceAttachment) {
		toSerialize["sequenceAttachment"] = o.SequenceAttachment
	}
	if !IsNil(o.NTerminalModificationId) {
		toSerialize["nTerminalModificationId"] = o.NTerminalModificationId
	}
	if !IsNil(o.NTerminalModification) {
		toSerialize["nTerminalModification"] = o.NTerminalModification
	}
	if !IsNil(o.CTerminalModificationId) {
		toSerialize["cTerminalModificationId"] = o.CTerminalModificationId
	}
	if !IsNil(o.CTerminalModification) {
		toSerialize["cTerminalModification"] = o.CTerminalModification
	}
	return toSerialize, nil
}

type NullableSubstanceProteinSubunit struct {
	value *SubstanceProteinSubunit
	isSet bool
}

func (v NullableSubstanceProteinSubunit) Get() *SubstanceProteinSubunit {
	return v.value
}

func (v *NullableSubstanceProteinSubunit) Set(val *SubstanceProteinSubunit) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstanceProteinSubunit) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstanceProteinSubunit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstanceProteinSubunit(val *SubstanceProteinSubunit) *NullableSubstanceProteinSubunit {
	return &NullableSubstanceProteinSubunit{value: val, isSet: true}
}

func (v NullableSubstanceProteinSubunit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstanceProteinSubunit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


