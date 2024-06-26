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

// checks if the NutritionOrderNutrient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NutritionOrderNutrient{}

// NutritionOrderNutrient A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrderNutrient struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The nutrient that is being modified such as carbohydrate or sodium.
	Modifier *CodeableConcept `json:"modifier,omitempty"`
	// The quantity of the specified nutrient to include in diet.
	Amount *Quantity `json:"amount,omitempty"`
}

// NewNutritionOrderNutrient instantiates a new NutritionOrderNutrient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNutritionOrderNutrient() *NutritionOrderNutrient {
	this := NutritionOrderNutrient{}
	return &this
}

// NewNutritionOrderNutrientWithDefaults instantiates a new NutritionOrderNutrient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNutritionOrderNutrientWithDefaults() *NutritionOrderNutrient {
	this := NutritionOrderNutrient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NutritionOrderNutrient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderNutrient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NutritionOrderNutrient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NutritionOrderNutrient) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *NutritionOrderNutrient) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderNutrient) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *NutritionOrderNutrient) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *NutritionOrderNutrient) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *NutritionOrderNutrient) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderNutrient) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *NutritionOrderNutrient) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *NutritionOrderNutrient) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetModifier returns the Modifier field value if set, zero value otherwise.
func (o *NutritionOrderNutrient) GetModifier() CodeableConcept {
	if o == nil || IsNil(o.Modifier) {
		var ret CodeableConcept
		return ret
	}
	return *o.Modifier
}

// GetModifierOk returns a tuple with the Modifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderNutrient) GetModifierOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Modifier) {
		return nil, false
	}
	return o.Modifier, true
}

// HasModifier returns a boolean if a field has been set.
func (o *NutritionOrderNutrient) HasModifier() bool {
	if o != nil && !IsNil(o.Modifier) {
		return true
	}

	return false
}

// SetModifier gets a reference to the given CodeableConcept and assigns it to the Modifier field.
func (o *NutritionOrderNutrient) SetModifier(v CodeableConcept) {
	o.Modifier = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NutritionOrderNutrient) GetAmount() Quantity {
	if o == nil || IsNil(o.Amount) {
		var ret Quantity
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderNutrient) GetAmountOk() (*Quantity, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NutritionOrderNutrient) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Quantity and assigns it to the Amount field.
func (o *NutritionOrderNutrient) SetAmount(v Quantity) {
	o.Amount = &v
}

func (o NutritionOrderNutrient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NutritionOrderNutrient) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Modifier) {
		toSerialize["modifier"] = o.Modifier
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableNutritionOrderNutrient struct {
	value *NutritionOrderNutrient
	isSet bool
}

func (v NullableNutritionOrderNutrient) Get() *NutritionOrderNutrient {
	return v.value
}

func (v *NullableNutritionOrderNutrient) Set(val *NutritionOrderNutrient) {
	v.value = val
	v.isSet = true
}

func (v NullableNutritionOrderNutrient) IsSet() bool {
	return v.isSet
}

func (v *NullableNutritionOrderNutrient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNutritionOrderNutrient(val *NutritionOrderNutrient) *NullableNutritionOrderNutrient {
	return &NullableNutritionOrderNutrient{value: val, isSet: true}
}

func (v NullableNutritionOrderNutrient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNutritionOrderNutrient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


