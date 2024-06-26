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

// checks if the SubstancePolymerStartingMaterial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstancePolymerStartingMaterial{}

// SubstancePolymerStartingMaterial Todo.
type SubstancePolymerStartingMaterial struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Todo.
	Material *CodeableConcept `json:"material,omitempty"`
	// Todo.
	Type *CodeableConcept `json:"type,omitempty"`
	// Value of \"true\" or \"false\"
	IsDefining *bool `json:"isDefining,omitempty"`
	// Todo.
	Amount *SubstanceAmount `json:"amount,omitempty"`
}

// NewSubstancePolymerStartingMaterial instantiates a new SubstancePolymerStartingMaterial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstancePolymerStartingMaterial() *SubstancePolymerStartingMaterial {
	this := SubstancePolymerStartingMaterial{}
	return &this
}

// NewSubstancePolymerStartingMaterialWithDefaults instantiates a new SubstancePolymerStartingMaterial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstancePolymerStartingMaterialWithDefaults() *SubstancePolymerStartingMaterial {
	this := SubstancePolymerStartingMaterial{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstancePolymerStartingMaterial) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstancePolymerStartingMaterial) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstancePolymerStartingMaterial) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetMaterial() CodeableConcept {
	if o == nil || IsNil(o.Material) {
		var ret CodeableConcept
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetMaterialOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Material) {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasMaterial() bool {
	if o != nil && !IsNil(o.Material) {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given CodeableConcept and assigns it to the Material field.
func (o *SubstancePolymerStartingMaterial) SetMaterial(v CodeableConcept) {
	o.Material = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *SubstancePolymerStartingMaterial) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetIsDefining returns the IsDefining field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetIsDefining() bool {
	if o == nil || IsNil(o.IsDefining) {
		var ret bool
		return ret
	}
	return *o.IsDefining
}

// GetIsDefiningOk returns a tuple with the IsDefining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetIsDefiningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefining) {
		return nil, false
	}
	return o.IsDefining, true
}

// HasIsDefining returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasIsDefining() bool {
	if o != nil && !IsNil(o.IsDefining) {
		return true
	}

	return false
}

// SetIsDefining gets a reference to the given bool and assigns it to the IsDefining field.
func (o *SubstancePolymerStartingMaterial) SetIsDefining(v bool) {
	o.IsDefining = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubstancePolymerStartingMaterial) GetAmount() SubstanceAmount {
	if o == nil || IsNil(o.Amount) {
		var ret SubstanceAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerStartingMaterial) GetAmountOk() (*SubstanceAmount, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubstancePolymerStartingMaterial) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given SubstanceAmount and assigns it to the Amount field.
func (o *SubstancePolymerStartingMaterial) SetAmount(v SubstanceAmount) {
	o.Amount = &v
}

func (o SubstancePolymerStartingMaterial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstancePolymerStartingMaterial) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Material) {
		toSerialize["material"] = o.Material
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsDefining) {
		toSerialize["isDefining"] = o.IsDefining
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSubstancePolymerStartingMaterial struct {
	value *SubstancePolymerStartingMaterial
	isSet bool
}

func (v NullableSubstancePolymerStartingMaterial) Get() *SubstancePolymerStartingMaterial {
	return v.value
}

func (v *NullableSubstancePolymerStartingMaterial) Set(val *SubstancePolymerStartingMaterial) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstancePolymerStartingMaterial) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstancePolymerStartingMaterial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstancePolymerStartingMaterial(val *SubstancePolymerStartingMaterial) *NullableSubstancePolymerStartingMaterial {
	return &NullableSubstancePolymerStartingMaterial{value: val, isSet: true}
}

func (v NullableSubstancePolymerStartingMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstancePolymerStartingMaterial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

