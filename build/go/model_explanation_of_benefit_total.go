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

// checks if the ExplanationOfBenefitTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplanationOfBenefitTotal{}

// ExplanationOfBenefitTotal This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitTotal struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept `json:"category"`
	// Monetary total amount associated with the category.
	Amount Money `json:"amount"`
}

type _ExplanationOfBenefitTotal ExplanationOfBenefitTotal

// NewExplanationOfBenefitTotal instantiates a new ExplanationOfBenefitTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplanationOfBenefitTotal(category CodeableConcept, amount Money) *ExplanationOfBenefitTotal {
	this := ExplanationOfBenefitTotal{}
	this.Category = category
	this.Amount = amount
	return &this
}

// NewExplanationOfBenefitTotalWithDefaults instantiates a new ExplanationOfBenefitTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplanationOfBenefitTotalWithDefaults() *ExplanationOfBenefitTotal {
	this := ExplanationOfBenefitTotal{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExplanationOfBenefitTotal) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitTotal) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExplanationOfBenefitTotal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExplanationOfBenefitTotal) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitTotal) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitTotal) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitTotal) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExplanationOfBenefitTotal) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitTotal) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitTotal) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitTotal) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExplanationOfBenefitTotal) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCategory returns the Category field value
func (o *ExplanationOfBenefitTotal) GetCategory() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitTotal) GetCategoryOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ExplanationOfBenefitTotal) SetCategory(v CodeableConcept) {
	o.Category = v
}

// GetAmount returns the Amount field value
func (o *ExplanationOfBenefitTotal) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitTotal) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ExplanationOfBenefitTotal) SetAmount(v Money) {
	o.Amount = v
}

func (o ExplanationOfBenefitTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplanationOfBenefitTotal) ToMap() (map[string]interface{}, error) {
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
	toSerialize["category"] = o.Category
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *ExplanationOfBenefitTotal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
		"amount",
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

	varExplanationOfBenefitTotal := _ExplanationOfBenefitTotal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExplanationOfBenefitTotal)

	if err != nil {
		return err
	}

	*o = ExplanationOfBenefitTotal(varExplanationOfBenefitTotal)

	return err
}

type NullableExplanationOfBenefitTotal struct {
	value *ExplanationOfBenefitTotal
	isSet bool
}

func (v NullableExplanationOfBenefitTotal) Get() *ExplanationOfBenefitTotal {
	return v.value
}

func (v *NullableExplanationOfBenefitTotal) Set(val *ExplanationOfBenefitTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableExplanationOfBenefitTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableExplanationOfBenefitTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplanationOfBenefitTotal(val *ExplanationOfBenefitTotal) *NullableExplanationOfBenefitTotal {
	return &NullableExplanationOfBenefitTotal{value: val, isSet: true}
}

func (v NullableExplanationOfBenefitTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplanationOfBenefitTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


