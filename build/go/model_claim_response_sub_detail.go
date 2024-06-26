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

// checks if the ClaimResponseSubDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimResponseSubDetail{}

// ClaimResponseSubDetail This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponseSubDetail struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	SubDetailSequence *float32 `json:"subDetailSequence,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []float32 `json:"noteNumber,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseAdjudication `json:"adjudication,omitempty"`
}

// NewClaimResponseSubDetail instantiates a new ClaimResponseSubDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimResponseSubDetail() *ClaimResponseSubDetail {
	this := ClaimResponseSubDetail{}
	return &this
}

// NewClaimResponseSubDetailWithDefaults instantiates a new ClaimResponseSubDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimResponseSubDetailWithDefaults() *ClaimResponseSubDetail {
	this := ClaimResponseSubDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClaimResponseSubDetail) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ClaimResponseSubDetail) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ClaimResponseSubDetail) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSubDetailSequence returns the SubDetailSequence field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetSubDetailSequence() float32 {
	if o == nil || IsNil(o.SubDetailSequence) {
		var ret float32
		return ret
	}
	return *o.SubDetailSequence
}

// GetSubDetailSequenceOk returns a tuple with the SubDetailSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetSubDetailSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.SubDetailSequence) {
		return nil, false
	}
	return o.SubDetailSequence, true
}

// HasSubDetailSequence returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasSubDetailSequence() bool {
	if o != nil && !IsNil(o.SubDetailSequence) {
		return true
	}

	return false
}

// SetSubDetailSequence gets a reference to the given float32 and assigns it to the SubDetailSequence field.
func (o *ClaimResponseSubDetail) SetSubDetailSequence(v float32) {
	o.SubDetailSequence = &v
}

// GetNoteNumber returns the NoteNumber field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetNoteNumber() []float32 {
	if o == nil || IsNil(o.NoteNumber) {
		var ret []float32
		return ret
	}
	return o.NoteNumber
}

// GetNoteNumberOk returns a tuple with the NoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetNoteNumberOk() ([]float32, bool) {
	if o == nil || IsNil(o.NoteNumber) {
		return nil, false
	}
	return o.NoteNumber, true
}

// HasNoteNumber returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasNoteNumber() bool {
	if o != nil && !IsNil(o.NoteNumber) {
		return true
	}

	return false
}

// SetNoteNumber gets a reference to the given []float32 and assigns it to the NoteNumber field.
func (o *ClaimResponseSubDetail) SetNoteNumber(v []float32) {
	o.NoteNumber = v
}

// GetAdjudication returns the Adjudication field value if set, zero value otherwise.
func (o *ClaimResponseSubDetail) GetAdjudication() []ClaimResponseAdjudication {
	if o == nil || IsNil(o.Adjudication) {
		var ret []ClaimResponseAdjudication
		return ret
	}
	return o.Adjudication
}

// GetAdjudicationOk returns a tuple with the Adjudication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseSubDetail) GetAdjudicationOk() ([]ClaimResponseAdjudication, bool) {
	if o == nil || IsNil(o.Adjudication) {
		return nil, false
	}
	return o.Adjudication, true
}

// HasAdjudication returns a boolean if a field has been set.
func (o *ClaimResponseSubDetail) HasAdjudication() bool {
	if o != nil && !IsNil(o.Adjudication) {
		return true
	}

	return false
}

// SetAdjudication gets a reference to the given []ClaimResponseAdjudication and assigns it to the Adjudication field.
func (o *ClaimResponseSubDetail) SetAdjudication(v []ClaimResponseAdjudication) {
	o.Adjudication = v
}

func (o ClaimResponseSubDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimResponseSubDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SubDetailSequence) {
		toSerialize["subDetailSequence"] = o.SubDetailSequence
	}
	if !IsNil(o.NoteNumber) {
		toSerialize["noteNumber"] = o.NoteNumber
	}
	if !IsNil(o.Adjudication) {
		toSerialize["adjudication"] = o.Adjudication
	}
	return toSerialize, nil
}

type NullableClaimResponseSubDetail struct {
	value *ClaimResponseSubDetail
	isSet bool
}

func (v NullableClaimResponseSubDetail) Get() *ClaimResponseSubDetail {
	return v.value
}

func (v *NullableClaimResponseSubDetail) Set(val *ClaimResponseSubDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimResponseSubDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimResponseSubDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimResponseSubDetail(val *ClaimResponseSubDetail) *NullableClaimResponseSubDetail {
	return &NullableClaimResponseSubDetail{value: val, isSet: true}
}

func (v NullableClaimResponseSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimResponseSubDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

