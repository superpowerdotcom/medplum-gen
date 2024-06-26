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

// checks if the ImmunizationRecommendationDateCriterion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImmunizationRecommendationDateCriterion{}

// ImmunizationRecommendationDateCriterion A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendationDateCriterion struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Date classification of recommendation.  For example, earliest date to give, latest date to give, etc.
	Code CodeableConcept `json:"code"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Value *string `json:"value,omitempty"`
}

type _ImmunizationRecommendationDateCriterion ImmunizationRecommendationDateCriterion

// NewImmunizationRecommendationDateCriterion instantiates a new ImmunizationRecommendationDateCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImmunizationRecommendationDateCriterion(code CodeableConcept) *ImmunizationRecommendationDateCriterion {
	this := ImmunizationRecommendationDateCriterion{}
	this.Code = code
	return &this
}

// NewImmunizationRecommendationDateCriterionWithDefaults instantiates a new ImmunizationRecommendationDateCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImmunizationRecommendationDateCriterionWithDefaults() *ImmunizationRecommendationDateCriterion {
	this := ImmunizationRecommendationDateCriterion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImmunizationRecommendationDateCriterion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationRecommendationDateCriterion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImmunizationRecommendationDateCriterion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImmunizationRecommendationDateCriterion) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ImmunizationRecommendationDateCriterion) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationRecommendationDateCriterion) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ImmunizationRecommendationDateCriterion) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ImmunizationRecommendationDateCriterion) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ImmunizationRecommendationDateCriterion) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationRecommendationDateCriterion) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ImmunizationRecommendationDateCriterion) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ImmunizationRecommendationDateCriterion) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value
func (o *ImmunizationRecommendationDateCriterion) GetCode() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ImmunizationRecommendationDateCriterion) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ImmunizationRecommendationDateCriterion) SetCode(v CodeableConcept) {
	o.Code = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ImmunizationRecommendationDateCriterion) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationRecommendationDateCriterion) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ImmunizationRecommendationDateCriterion) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ImmunizationRecommendationDateCriterion) SetValue(v string) {
	o.Value = &v
}

func (o ImmunizationRecommendationDateCriterion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImmunizationRecommendationDateCriterion) ToMap() (map[string]interface{}, error) {
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
	toSerialize["code"] = o.Code
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *ImmunizationRecommendationDateCriterion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varImmunizationRecommendationDateCriterion := _ImmunizationRecommendationDateCriterion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImmunizationRecommendationDateCriterion)

	if err != nil {
		return err
	}

	*o = ImmunizationRecommendationDateCriterion(varImmunizationRecommendationDateCriterion)

	return err
}

type NullableImmunizationRecommendationDateCriterion struct {
	value *ImmunizationRecommendationDateCriterion
	isSet bool
}

func (v NullableImmunizationRecommendationDateCriterion) Get() *ImmunizationRecommendationDateCriterion {
	return v.value
}

func (v *NullableImmunizationRecommendationDateCriterion) Set(val *ImmunizationRecommendationDateCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableImmunizationRecommendationDateCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableImmunizationRecommendationDateCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImmunizationRecommendationDateCriterion(val *ImmunizationRecommendationDateCriterion) *NullableImmunizationRecommendationDateCriterion {
	return &NullableImmunizationRecommendationDateCriterion{value: val, isSet: true}
}

func (v NullableImmunizationRecommendationDateCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImmunizationRecommendationDateCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


