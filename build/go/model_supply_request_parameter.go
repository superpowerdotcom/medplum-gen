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

// checks if the SupplyRequestParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyRequestParameter{}

// SupplyRequestParameter A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequestParameter struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code or string that identifies the device detail being asserted.
	Code *CodeableConcept `json:"code,omitempty"`
	// The value of the device detail.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// The value of the device detail.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The value of the device detail.
	ValueRange *Range `json:"valueRange,omitempty"`
	// The value of the device detail.
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
}

// NewSupplyRequestParameter instantiates a new SupplyRequestParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyRequestParameter() *SupplyRequestParameter {
	this := SupplyRequestParameter{}
	return &this
}

// NewSupplyRequestParameterWithDefaults instantiates a new SupplyRequestParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyRequestParameterWithDefaults() *SupplyRequestParameter {
	this := SupplyRequestParameter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupplyRequestParameter) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SupplyRequestParameter) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SupplyRequestParameter) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetCode() CodeableConcept {
	if o == nil || IsNil(o.Code) {
		var ret CodeableConcept
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given CodeableConcept and assigns it to the Code field.
func (o *SupplyRequestParameter) SetCode(v CodeableConcept) {
	o.Code = &v
}

// GetValueCodeableConcept returns the ValueCodeableConcept field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetValueCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ValueCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ValueCodeableConcept
}

// GetValueCodeableConceptOk returns a tuple with the ValueCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetValueCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ValueCodeableConcept) {
		return nil, false
	}
	return o.ValueCodeableConcept, true
}

// HasValueCodeableConcept returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasValueCodeableConcept() bool {
	if o != nil && !IsNil(o.ValueCodeableConcept) {
		return true
	}

	return false
}

// SetValueCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ValueCodeableConcept field.
func (o *SupplyRequestParameter) SetValueCodeableConcept(v CodeableConcept) {
	o.ValueCodeableConcept = &v
}

// GetValueQuantity returns the ValueQuantity field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetValueQuantity() Quantity {
	if o == nil || IsNil(o.ValueQuantity) {
		var ret Quantity
		return ret
	}
	return *o.ValueQuantity
}

// GetValueQuantityOk returns a tuple with the ValueQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetValueQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.ValueQuantity) {
		return nil, false
	}
	return o.ValueQuantity, true
}

// HasValueQuantity returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasValueQuantity() bool {
	if o != nil && !IsNil(o.ValueQuantity) {
		return true
	}

	return false
}

// SetValueQuantity gets a reference to the given Quantity and assigns it to the ValueQuantity field.
func (o *SupplyRequestParameter) SetValueQuantity(v Quantity) {
	o.ValueQuantity = &v
}

// GetValueRange returns the ValueRange field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetValueRange() Range {
	if o == nil || IsNil(o.ValueRange) {
		var ret Range
		return ret
	}
	return *o.ValueRange
}

// GetValueRangeOk returns a tuple with the ValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetValueRangeOk() (*Range, bool) {
	if o == nil || IsNil(o.ValueRange) {
		return nil, false
	}
	return o.ValueRange, true
}

// HasValueRange returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasValueRange() bool {
	if o != nil && !IsNil(o.ValueRange) {
		return true
	}

	return false
}

// SetValueRange gets a reference to the given Range and assigns it to the ValueRange field.
func (o *SupplyRequestParameter) SetValueRange(v Range) {
	o.ValueRange = &v
}

// GetValueBoolean returns the ValueBoolean field value if set, zero value otherwise.
func (o *SupplyRequestParameter) GetValueBoolean() bool {
	if o == nil || IsNil(o.ValueBoolean) {
		var ret bool
		return ret
	}
	return *o.ValueBoolean
}

// GetValueBooleanOk returns a tuple with the ValueBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequestParameter) GetValueBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueBoolean) {
		return nil, false
	}
	return o.ValueBoolean, true
}

// HasValueBoolean returns a boolean if a field has been set.
func (o *SupplyRequestParameter) HasValueBoolean() bool {
	if o != nil && !IsNil(o.ValueBoolean) {
		return true
	}

	return false
}

// SetValueBoolean gets a reference to the given bool and assigns it to the ValueBoolean field.
func (o *SupplyRequestParameter) SetValueBoolean(v bool) {
	o.ValueBoolean = &v
}

func (o SupplyRequestParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyRequestParameter) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ValueCodeableConcept) {
		toSerialize["valueCodeableConcept"] = o.ValueCodeableConcept
	}
	if !IsNil(o.ValueQuantity) {
		toSerialize["valueQuantity"] = o.ValueQuantity
	}
	if !IsNil(o.ValueRange) {
		toSerialize["valueRange"] = o.ValueRange
	}
	if !IsNil(o.ValueBoolean) {
		toSerialize["valueBoolean"] = o.ValueBoolean
	}
	return toSerialize, nil
}

type NullableSupplyRequestParameter struct {
	value *SupplyRequestParameter
	isSet bool
}

func (v NullableSupplyRequestParameter) Get() *SupplyRequestParameter {
	return v.value
}

func (v *NullableSupplyRequestParameter) Set(val *SupplyRequestParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyRequestParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyRequestParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyRequestParameter(val *SupplyRequestParameter) *NullableSupplyRequestParameter {
	return &NullableSupplyRequestParameter{value: val, isSet: true}
}

func (v NullableSupplyRequestParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyRequestParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

