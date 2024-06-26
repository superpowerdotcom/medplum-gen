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

// checks if the CoverageCostToBeneficiary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoverageCostToBeneficiary{}

// CoverageCostToBeneficiary Financial instrument which may be used to reimburse or pay for health care products and services. Includes both insurance and self-payment.
type CoverageCostToBeneficiary struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The category of patient centric costs associated with treatment.
	Type *CodeableConcept `json:"type,omitempty"`
	// The amount due from the patient for the cost category.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The amount due from the patient for the cost category.
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// A suite of codes indicating exceptions or reductions to patient costs and their effective periods.
	Exception []CoverageException `json:"exception,omitempty"`
}

// NewCoverageCostToBeneficiary instantiates a new CoverageCostToBeneficiary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoverageCostToBeneficiary() *CoverageCostToBeneficiary {
	this := CoverageCostToBeneficiary{}
	return &this
}

// NewCoverageCostToBeneficiaryWithDefaults instantiates a new CoverageCostToBeneficiary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoverageCostToBeneficiaryWithDefaults() *CoverageCostToBeneficiary {
	this := CoverageCostToBeneficiary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoverageCostToBeneficiary) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CoverageCostToBeneficiary) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CoverageCostToBeneficiary) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *CoverageCostToBeneficiary) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetValueQuantity returns the ValueQuantity field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetValueQuantity() Quantity {
	if o == nil || IsNil(o.ValueQuantity) {
		var ret Quantity
		return ret
	}
	return *o.ValueQuantity
}

// GetValueQuantityOk returns a tuple with the ValueQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetValueQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.ValueQuantity) {
		return nil, false
	}
	return o.ValueQuantity, true
}

// HasValueQuantity returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasValueQuantity() bool {
	if o != nil && !IsNil(o.ValueQuantity) {
		return true
	}

	return false
}

// SetValueQuantity gets a reference to the given Quantity and assigns it to the ValueQuantity field.
func (o *CoverageCostToBeneficiary) SetValueQuantity(v Quantity) {
	o.ValueQuantity = &v
}

// GetValueMoney returns the ValueMoney field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetValueMoney() Money {
	if o == nil || IsNil(o.ValueMoney) {
		var ret Money
		return ret
	}
	return *o.ValueMoney
}

// GetValueMoneyOk returns a tuple with the ValueMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetValueMoneyOk() (*Money, bool) {
	if o == nil || IsNil(o.ValueMoney) {
		return nil, false
	}
	return o.ValueMoney, true
}

// HasValueMoney returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasValueMoney() bool {
	if o != nil && !IsNil(o.ValueMoney) {
		return true
	}

	return false
}

// SetValueMoney gets a reference to the given Money and assigns it to the ValueMoney field.
func (o *CoverageCostToBeneficiary) SetValueMoney(v Money) {
	o.ValueMoney = &v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *CoverageCostToBeneficiary) GetException() []CoverageException {
	if o == nil || IsNil(o.Exception) {
		var ret []CoverageException
		return ret
	}
	return o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageCostToBeneficiary) GetExceptionOk() ([]CoverageException, bool) {
	if o == nil || IsNil(o.Exception) {
		return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *CoverageCostToBeneficiary) HasException() bool {
	if o != nil && !IsNil(o.Exception) {
		return true
	}

	return false
}

// SetException gets a reference to the given []CoverageException and assigns it to the Exception field.
func (o *CoverageCostToBeneficiary) SetException(v []CoverageException) {
	o.Exception = v
}

func (o CoverageCostToBeneficiary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoverageCostToBeneficiary) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ValueQuantity) {
		toSerialize["valueQuantity"] = o.ValueQuantity
	}
	if !IsNil(o.ValueMoney) {
		toSerialize["valueMoney"] = o.ValueMoney
	}
	if !IsNil(o.Exception) {
		toSerialize["exception"] = o.Exception
	}
	return toSerialize, nil
}

type NullableCoverageCostToBeneficiary struct {
	value *CoverageCostToBeneficiary
	isSet bool
}

func (v NullableCoverageCostToBeneficiary) Get() *CoverageCostToBeneficiary {
	return v.value
}

func (v *NullableCoverageCostToBeneficiary) Set(val *CoverageCostToBeneficiary) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverageCostToBeneficiary) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverageCostToBeneficiary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverageCostToBeneficiary(val *CoverageCostToBeneficiary) *NullableCoverageCostToBeneficiary {
	return &NullableCoverageCostToBeneficiary{value: val, isSet: true}
}

func (v NullableCoverageCostToBeneficiary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverageCostToBeneficiary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


