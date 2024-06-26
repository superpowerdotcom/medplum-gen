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

// checks if the InsurancePlanGeneralCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsurancePlanGeneralCost{}

// InsurancePlanGeneralCost Details of a Health Insurance product/plan provided by an organization.
type InsurancePlanGeneralCost struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of cost.
	Type *CodeableConcept `json:"type,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	GroupSize *float32 `json:"groupSize,omitempty"`
	// Value of the cost.
	Cost *Money `json:"cost,omitempty"`
	// A sequence of Unicode characters
	Comment *string `json:"comment,omitempty"`
}

// NewInsurancePlanGeneralCost instantiates a new InsurancePlanGeneralCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsurancePlanGeneralCost() *InsurancePlanGeneralCost {
	this := InsurancePlanGeneralCost{}
	return &this
}

// NewInsurancePlanGeneralCostWithDefaults instantiates a new InsurancePlanGeneralCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsurancePlanGeneralCostWithDefaults() *InsurancePlanGeneralCost {
	this := InsurancePlanGeneralCost{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InsurancePlanGeneralCost) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *InsurancePlanGeneralCost) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *InsurancePlanGeneralCost) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *InsurancePlanGeneralCost) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetGroupSize returns the GroupSize field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetGroupSize() float32 {
	if o == nil || IsNil(o.GroupSize) {
		var ret float32
		return ret
	}
	return *o.GroupSize
}

// GetGroupSizeOk returns a tuple with the GroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetGroupSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.GroupSize) {
		return nil, false
	}
	return o.GroupSize, true
}

// HasGroupSize returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasGroupSize() bool {
	if o != nil && !IsNil(o.GroupSize) {
		return true
	}

	return false
}

// SetGroupSize gets a reference to the given float32 and assigns it to the GroupSize field.
func (o *InsurancePlanGeneralCost) SetGroupSize(v float32) {
	o.GroupSize = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetCost() Money {
	if o == nil || IsNil(o.Cost) {
		var ret Money
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetCostOk() (*Money, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given Money and assigns it to the Cost field.
func (o *InsurancePlanGeneralCost) SetCost(v Money) {
	o.Cost = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InsurancePlanGeneralCost) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanGeneralCost) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InsurancePlanGeneralCost) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InsurancePlanGeneralCost) SetComment(v string) {
	o.Comment = &v
}

func (o InsurancePlanGeneralCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsurancePlanGeneralCost) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.GroupSize) {
		toSerialize["groupSize"] = o.GroupSize
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableInsurancePlanGeneralCost struct {
	value *InsurancePlanGeneralCost
	isSet bool
}

func (v NullableInsurancePlanGeneralCost) Get() *InsurancePlanGeneralCost {
	return v.value
}

func (v *NullableInsurancePlanGeneralCost) Set(val *InsurancePlanGeneralCost) {
	v.value = val
	v.isSet = true
}

func (v NullableInsurancePlanGeneralCost) IsSet() bool {
	return v.isSet
}

func (v *NullableInsurancePlanGeneralCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsurancePlanGeneralCost(val *InsurancePlanGeneralCost) *NullableInsurancePlanGeneralCost {
	return &NullableInsurancePlanGeneralCost{value: val, isSet: true}
}

func (v NullableInsurancePlanGeneralCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsurancePlanGeneralCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


