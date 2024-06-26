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

// checks if the ClaimResponseDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimResponseDetail{}

// ClaimResponseDetail This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponseDetail struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	DetailSequence *float32 `json:"detailSequence,omitempty"`
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []float32 `json:"noteNumber,omitempty"`
	// The adjudication results.
	Adjudication []ClaimResponseAdjudication `json:"adjudication"`
	// A sub-detail adjudication of a simple product or service.
	SubDetail []ClaimResponseSubDetail `json:"subDetail,omitempty"`
}

type _ClaimResponseDetail ClaimResponseDetail

// NewClaimResponseDetail instantiates a new ClaimResponseDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimResponseDetail(adjudication []ClaimResponseAdjudication) *ClaimResponseDetail {
	this := ClaimResponseDetail{}
	this.Adjudication = adjudication
	return &this
}

// NewClaimResponseDetailWithDefaults instantiates a new ClaimResponseDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimResponseDetailWithDefaults() *ClaimResponseDetail {
	this := ClaimResponseDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClaimResponseDetail) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ClaimResponseDetail) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ClaimResponseDetail) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDetailSequence returns the DetailSequence field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetDetailSequence() float32 {
	if o == nil || IsNil(o.DetailSequence) {
		var ret float32
		return ret
	}
	return *o.DetailSequence
}

// GetDetailSequenceOk returns a tuple with the DetailSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetDetailSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.DetailSequence) {
		return nil, false
	}
	return o.DetailSequence, true
}

// HasDetailSequence returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasDetailSequence() bool {
	if o != nil && !IsNil(o.DetailSequence) {
		return true
	}

	return false
}

// SetDetailSequence gets a reference to the given float32 and assigns it to the DetailSequence field.
func (o *ClaimResponseDetail) SetDetailSequence(v float32) {
	o.DetailSequence = &v
}

// GetNoteNumber returns the NoteNumber field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetNoteNumber() []float32 {
	if o == nil || IsNil(o.NoteNumber) {
		var ret []float32
		return ret
	}
	return o.NoteNumber
}

// GetNoteNumberOk returns a tuple with the NoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetNoteNumberOk() ([]float32, bool) {
	if o == nil || IsNil(o.NoteNumber) {
		return nil, false
	}
	return o.NoteNumber, true
}

// HasNoteNumber returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasNoteNumber() bool {
	if o != nil && !IsNil(o.NoteNumber) {
		return true
	}

	return false
}

// SetNoteNumber gets a reference to the given []float32 and assigns it to the NoteNumber field.
func (o *ClaimResponseDetail) SetNoteNumber(v []float32) {
	o.NoteNumber = v
}

// GetAdjudication returns the Adjudication field value
func (o *ClaimResponseDetail) GetAdjudication() []ClaimResponseAdjudication {
	if o == nil {
		var ret []ClaimResponseAdjudication
		return ret
	}

	return o.Adjudication
}

// GetAdjudicationOk returns a tuple with the Adjudication field value
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetAdjudicationOk() ([]ClaimResponseAdjudication, bool) {
	if o == nil {
		return nil, false
	}
	return o.Adjudication, true
}

// SetAdjudication sets field value
func (o *ClaimResponseDetail) SetAdjudication(v []ClaimResponseAdjudication) {
	o.Adjudication = v
}

// GetSubDetail returns the SubDetail field value if set, zero value otherwise.
func (o *ClaimResponseDetail) GetSubDetail() []ClaimResponseSubDetail {
	if o == nil || IsNil(o.SubDetail) {
		var ret []ClaimResponseSubDetail
		return ret
	}
	return o.SubDetail
}

// GetSubDetailOk returns a tuple with the SubDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimResponseDetail) GetSubDetailOk() ([]ClaimResponseSubDetail, bool) {
	if o == nil || IsNil(o.SubDetail) {
		return nil, false
	}
	return o.SubDetail, true
}

// HasSubDetail returns a boolean if a field has been set.
func (o *ClaimResponseDetail) HasSubDetail() bool {
	if o != nil && !IsNil(o.SubDetail) {
		return true
	}

	return false
}

// SetSubDetail gets a reference to the given []ClaimResponseSubDetail and assigns it to the SubDetail field.
func (o *ClaimResponseDetail) SetSubDetail(v []ClaimResponseSubDetail) {
	o.SubDetail = v
}

func (o ClaimResponseDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimResponseDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DetailSequence) {
		toSerialize["detailSequence"] = o.DetailSequence
	}
	if !IsNil(o.NoteNumber) {
		toSerialize["noteNumber"] = o.NoteNumber
	}
	toSerialize["adjudication"] = o.Adjudication
	if !IsNil(o.SubDetail) {
		toSerialize["subDetail"] = o.SubDetail
	}
	return toSerialize, nil
}

func (o *ClaimResponseDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"adjudication",
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

	varClaimResponseDetail := _ClaimResponseDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClaimResponseDetail)

	if err != nil {
		return err
	}

	*o = ClaimResponseDetail(varClaimResponseDetail)

	return err
}

type NullableClaimResponseDetail struct {
	value *ClaimResponseDetail
	isSet bool
}

func (v NullableClaimResponseDetail) Get() *ClaimResponseDetail {
	return v.value
}

func (v *NullableClaimResponseDetail) Set(val *ClaimResponseDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimResponseDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimResponseDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimResponseDetail(val *ClaimResponseDetail) *NullableClaimResponseDetail {
	return &NullableClaimResponseDetail{value: val, isSet: true}
}

func (v NullableClaimResponseDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimResponseDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


