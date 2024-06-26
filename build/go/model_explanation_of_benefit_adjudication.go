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

// checks if the ExplanationOfBenefitAdjudication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplanationOfBenefitAdjudication{}

// ExplanationOfBenefitAdjudication This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitAdjudication struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in-aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept `json:"category"`
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *CodeableConcept `json:"reason,omitempty"`
	// Monetary amount associated with the category.
	Amount *Money `json:"amount,omitempty"`
	// A rational number with implicit precision
	Value *float32 `json:"value,omitempty"`
}

type _ExplanationOfBenefitAdjudication ExplanationOfBenefitAdjudication

// NewExplanationOfBenefitAdjudication instantiates a new ExplanationOfBenefitAdjudication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplanationOfBenefitAdjudication(category CodeableConcept) *ExplanationOfBenefitAdjudication {
	this := ExplanationOfBenefitAdjudication{}
	this.Category = category
	return &this
}

// NewExplanationOfBenefitAdjudicationWithDefaults instantiates a new ExplanationOfBenefitAdjudication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplanationOfBenefitAdjudicationWithDefaults() *ExplanationOfBenefitAdjudication {
	this := ExplanationOfBenefitAdjudication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExplanationOfBenefitAdjudication) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExplanationOfBenefitAdjudication) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExplanationOfBenefitAdjudication) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCategory returns the Category field value
func (o *ExplanationOfBenefitAdjudication) GetCategory() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetCategoryOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ExplanationOfBenefitAdjudication) SetCategory(v CodeableConcept) {
	o.Category = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetReason() CodeableConcept {
	if o == nil || IsNil(o.Reason) {
		var ret CodeableConcept
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetReasonOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given CodeableConcept and assigns it to the Reason field.
func (o *ExplanationOfBenefitAdjudication) SetReason(v CodeableConcept) {
	o.Reason = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetAmount() Money {
	if o == nil || IsNil(o.Amount) {
		var ret Money
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetAmountOk() (*Money, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Money and assigns it to the Amount field.
func (o *ExplanationOfBenefitAdjudication) SetAmount(v Money) {
	o.Amount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ExplanationOfBenefitAdjudication) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitAdjudication) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExplanationOfBenefitAdjudication) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ExplanationOfBenefitAdjudication) SetValue(v float32) {
	o.Value = &v
}

func (o ExplanationOfBenefitAdjudication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplanationOfBenefitAdjudication) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *ExplanationOfBenefitAdjudication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
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

	varExplanationOfBenefitAdjudication := _ExplanationOfBenefitAdjudication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExplanationOfBenefitAdjudication)

	if err != nil {
		return err
	}

	*o = ExplanationOfBenefitAdjudication(varExplanationOfBenefitAdjudication)

	return err
}

type NullableExplanationOfBenefitAdjudication struct {
	value *ExplanationOfBenefitAdjudication
	isSet bool
}

func (v NullableExplanationOfBenefitAdjudication) Get() *ExplanationOfBenefitAdjudication {
	return v.value
}

func (v *NullableExplanationOfBenefitAdjudication) Set(val *ExplanationOfBenefitAdjudication) {
	v.value = val
	v.isSet = true
}

func (v NullableExplanationOfBenefitAdjudication) IsSet() bool {
	return v.isSet
}

func (v *NullableExplanationOfBenefitAdjudication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplanationOfBenefitAdjudication(val *ExplanationOfBenefitAdjudication) *NullableExplanationOfBenefitAdjudication {
	return &NullableExplanationOfBenefitAdjudication{value: val, isSet: true}
}

func (v NullableExplanationOfBenefitAdjudication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplanationOfBenefitAdjudication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


