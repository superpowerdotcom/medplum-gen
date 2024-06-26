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

// checks if the SubstancePolymerDegreeOfPolymerisation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstancePolymerDegreeOfPolymerisation{}

// SubstancePolymerDegreeOfPolymerisation Todo.
type SubstancePolymerDegreeOfPolymerisation struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Todo.
	Degree *CodeableConcept `json:"degree,omitempty"`
	// Todo.
	Amount *SubstanceAmount `json:"amount,omitempty"`
}

// NewSubstancePolymerDegreeOfPolymerisation instantiates a new SubstancePolymerDegreeOfPolymerisation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstancePolymerDegreeOfPolymerisation() *SubstancePolymerDegreeOfPolymerisation {
	this := SubstancePolymerDegreeOfPolymerisation{}
	return &this
}

// NewSubstancePolymerDegreeOfPolymerisationWithDefaults instantiates a new SubstancePolymerDegreeOfPolymerisation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstancePolymerDegreeOfPolymerisationWithDefaults() *SubstancePolymerDegreeOfPolymerisation {
	this := SubstancePolymerDegreeOfPolymerisation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstancePolymerDegreeOfPolymerisation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstancePolymerDegreeOfPolymerisation) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstancePolymerDegreeOfPolymerisation) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstancePolymerDegreeOfPolymerisation) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstancePolymerDegreeOfPolymerisation) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstancePolymerDegreeOfPolymerisation) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *SubstancePolymerDegreeOfPolymerisation) GetDegree() CodeableConcept {
	if o == nil || IsNil(o.Degree) {
		var ret CodeableConcept
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) GetDegreeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Degree) {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) HasDegree() bool {
	if o != nil && !IsNil(o.Degree) {
		return true
	}

	return false
}

// SetDegree gets a reference to the given CodeableConcept and assigns it to the Degree field.
func (o *SubstancePolymerDegreeOfPolymerisation) SetDegree(v CodeableConcept) {
	o.Degree = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubstancePolymerDegreeOfPolymerisation) GetAmount() SubstanceAmount {
	if o == nil || IsNil(o.Amount) {
		var ret SubstanceAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) GetAmountOk() (*SubstanceAmount, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubstancePolymerDegreeOfPolymerisation) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given SubstanceAmount and assigns it to the Amount field.
func (o *SubstancePolymerDegreeOfPolymerisation) SetAmount(v SubstanceAmount) {
	o.Amount = &v
}

func (o SubstancePolymerDegreeOfPolymerisation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstancePolymerDegreeOfPolymerisation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Degree) {
		toSerialize["degree"] = o.Degree
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSubstancePolymerDegreeOfPolymerisation struct {
	value *SubstancePolymerDegreeOfPolymerisation
	isSet bool
}

func (v NullableSubstancePolymerDegreeOfPolymerisation) Get() *SubstancePolymerDegreeOfPolymerisation {
	return v.value
}

func (v *NullableSubstancePolymerDegreeOfPolymerisation) Set(val *SubstancePolymerDegreeOfPolymerisation) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstancePolymerDegreeOfPolymerisation) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstancePolymerDegreeOfPolymerisation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstancePolymerDegreeOfPolymerisation(val *SubstancePolymerDegreeOfPolymerisation) *NullableSubstancePolymerDegreeOfPolymerisation {
	return &NullableSubstancePolymerDegreeOfPolymerisation{value: val, isSet: true}
}

func (v NullableSubstancePolymerDegreeOfPolymerisation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstancePolymerDegreeOfPolymerisation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


