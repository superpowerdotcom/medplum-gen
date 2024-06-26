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

// checks if the ImagingStudyPerformer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagingStudyPerformer{}

// ImagingStudyPerformer Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudyPerformer struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Distinguishes the type of involvement of the performer in the series.
	Function *CodeableConcept `json:"function,omitempty"`
	// Indicates who or what performed the series.
	Actor Reference `json:"actor"`
}

type _ImagingStudyPerformer ImagingStudyPerformer

// NewImagingStudyPerformer instantiates a new ImagingStudyPerformer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagingStudyPerformer(actor Reference) *ImagingStudyPerformer {
	this := ImagingStudyPerformer{}
	this.Actor = actor
	return &this
}

// NewImagingStudyPerformerWithDefaults instantiates a new ImagingStudyPerformer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagingStudyPerformerWithDefaults() *ImagingStudyPerformer {
	this := ImagingStudyPerformer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImagingStudyPerformer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagingStudyPerformer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImagingStudyPerformer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImagingStudyPerformer) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ImagingStudyPerformer) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagingStudyPerformer) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ImagingStudyPerformer) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ImagingStudyPerformer) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ImagingStudyPerformer) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagingStudyPerformer) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ImagingStudyPerformer) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ImagingStudyPerformer) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *ImagingStudyPerformer) GetFunction() CodeableConcept {
	if o == nil || IsNil(o.Function) {
		var ret CodeableConcept
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagingStudyPerformer) GetFunctionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *ImagingStudyPerformer) HasFunction() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given CodeableConcept and assigns it to the Function field.
func (o *ImagingStudyPerformer) SetFunction(v CodeableConcept) {
	o.Function = &v
}

// GetActor returns the Actor field value
func (o *ImagingStudyPerformer) GetActor() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *ImagingStudyPerformer) GetActorOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *ImagingStudyPerformer) SetActor(v Reference) {
	o.Actor = v
}

func (o ImagingStudyPerformer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImagingStudyPerformer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	toSerialize["actor"] = o.Actor
	return toSerialize, nil
}

func (o *ImagingStudyPerformer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actor",
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

	varImagingStudyPerformer := _ImagingStudyPerformer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImagingStudyPerformer)

	if err != nil {
		return err
	}

	*o = ImagingStudyPerformer(varImagingStudyPerformer)

	return err
}

type NullableImagingStudyPerformer struct {
	value *ImagingStudyPerformer
	isSet bool
}

func (v NullableImagingStudyPerformer) Get() *ImagingStudyPerformer {
	return v.value
}

func (v *NullableImagingStudyPerformer) Set(val *ImagingStudyPerformer) {
	v.value = val
	v.isSet = true
}

func (v NullableImagingStudyPerformer) IsSet() bool {
	return v.isSet
}

func (v *NullableImagingStudyPerformer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagingStudyPerformer(val *ImagingStudyPerformer) *NullableImagingStudyPerformer {
	return &NullableImagingStudyPerformer{value: val, isSet: true}
}

func (v NullableImagingStudyPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImagingStudyPerformer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


