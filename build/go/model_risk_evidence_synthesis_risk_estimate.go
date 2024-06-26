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

// checks if the RiskEvidenceSynthesisRiskEstimate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvidenceSynthesisRiskEstimate{}

// RiskEvidenceSynthesisRiskEstimate The RiskEvidenceSynthesis resource describes the likelihood of an outcome in a population plus exposure state where the risk estimate is derived from a combination of research studies.
type RiskEvidenceSynthesisRiskEstimate struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Examples include proportion and mean.
	Type *CodeableConcept `json:"type,omitempty"`
	// A rational number with implicit precision
	Value *float32 `json:"value,omitempty"`
	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *CodeableConcept `json:"unitOfMeasure,omitempty"`
	// A whole number
	DenominatorCount *float32 `json:"denominatorCount,omitempty"`
	// A whole number
	NumeratorCount *float32 `json:"numeratorCount,omitempty"`
	// A description of the precision of the estimate for the effect.
	PrecisionEstimate []RiskEvidenceSynthesisPrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// NewRiskEvidenceSynthesisRiskEstimate instantiates a new RiskEvidenceSynthesisRiskEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvidenceSynthesisRiskEstimate() *RiskEvidenceSynthesisRiskEstimate {
	this := RiskEvidenceSynthesisRiskEstimate{}
	return &this
}

// NewRiskEvidenceSynthesisRiskEstimateWithDefaults instantiates a new RiskEvidenceSynthesisRiskEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvidenceSynthesisRiskEstimateWithDefaults() *RiskEvidenceSynthesisRiskEstimate {
	this := RiskEvidenceSynthesisRiskEstimate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetValue(v float32) {
	o.Value = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetUnitOfMeasure() CodeableConcept {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret CodeableConcept
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetUnitOfMeasureOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given CodeableConcept and assigns it to the UnitOfMeasure field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetUnitOfMeasure(v CodeableConcept) {
	o.UnitOfMeasure = &v
}

// GetDenominatorCount returns the DenominatorCount field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetDenominatorCount() float32 {
	if o == nil || IsNil(o.DenominatorCount) {
		var ret float32
		return ret
	}
	return *o.DenominatorCount
}

// GetDenominatorCountOk returns a tuple with the DenominatorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetDenominatorCountOk() (*float32, bool) {
	if o == nil || IsNil(o.DenominatorCount) {
		return nil, false
	}
	return o.DenominatorCount, true
}

// HasDenominatorCount returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasDenominatorCount() bool {
	if o != nil && !IsNil(o.DenominatorCount) {
		return true
	}

	return false
}

// SetDenominatorCount gets a reference to the given float32 and assigns it to the DenominatorCount field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetDenominatorCount(v float32) {
	o.DenominatorCount = &v
}

// GetNumeratorCount returns the NumeratorCount field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetNumeratorCount() float32 {
	if o == nil || IsNil(o.NumeratorCount) {
		var ret float32
		return ret
	}
	return *o.NumeratorCount
}

// GetNumeratorCountOk returns a tuple with the NumeratorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetNumeratorCountOk() (*float32, bool) {
	if o == nil || IsNil(o.NumeratorCount) {
		return nil, false
	}
	return o.NumeratorCount, true
}

// HasNumeratorCount returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasNumeratorCount() bool {
	if o != nil && !IsNil(o.NumeratorCount) {
		return true
	}

	return false
}

// SetNumeratorCount gets a reference to the given float32 and assigns it to the NumeratorCount field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetNumeratorCount(v float32) {
	o.NumeratorCount = &v
}

// GetPrecisionEstimate returns the PrecisionEstimate field value if set, zero value otherwise.
func (o *RiskEvidenceSynthesisRiskEstimate) GetPrecisionEstimate() []RiskEvidenceSynthesisPrecisionEstimate {
	if o == nil || IsNil(o.PrecisionEstimate) {
		var ret []RiskEvidenceSynthesisPrecisionEstimate
		return ret
	}
	return o.PrecisionEstimate
}

// GetPrecisionEstimateOk returns a tuple with the PrecisionEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) GetPrecisionEstimateOk() ([]RiskEvidenceSynthesisPrecisionEstimate, bool) {
	if o == nil || IsNil(o.PrecisionEstimate) {
		return nil, false
	}
	return o.PrecisionEstimate, true
}

// HasPrecisionEstimate returns a boolean if a field has been set.
func (o *RiskEvidenceSynthesisRiskEstimate) HasPrecisionEstimate() bool {
	if o != nil && !IsNil(o.PrecisionEstimate) {
		return true
	}

	return false
}

// SetPrecisionEstimate gets a reference to the given []RiskEvidenceSynthesisPrecisionEstimate and assigns it to the PrecisionEstimate field.
func (o *RiskEvidenceSynthesisRiskEstimate) SetPrecisionEstimate(v []RiskEvidenceSynthesisPrecisionEstimate) {
	o.PrecisionEstimate = v
}

func (o RiskEvidenceSynthesisRiskEstimate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvidenceSynthesisRiskEstimate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	if !IsNil(o.DenominatorCount) {
		toSerialize["denominatorCount"] = o.DenominatorCount
	}
	if !IsNil(o.NumeratorCount) {
		toSerialize["numeratorCount"] = o.NumeratorCount
	}
	if !IsNil(o.PrecisionEstimate) {
		toSerialize["precisionEstimate"] = o.PrecisionEstimate
	}
	return toSerialize, nil
}

type NullableRiskEvidenceSynthesisRiskEstimate struct {
	value *RiskEvidenceSynthesisRiskEstimate
	isSet bool
}

func (v NullableRiskEvidenceSynthesisRiskEstimate) Get() *RiskEvidenceSynthesisRiskEstimate {
	return v.value
}

func (v *NullableRiskEvidenceSynthesisRiskEstimate) Set(val *RiskEvidenceSynthesisRiskEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvidenceSynthesisRiskEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvidenceSynthesisRiskEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvidenceSynthesisRiskEstimate(val *RiskEvidenceSynthesisRiskEstimate) *NullableRiskEvidenceSynthesisRiskEstimate {
	return &NullableRiskEvidenceSynthesisRiskEstimate{value: val, isSet: true}
}

func (v NullableRiskEvidenceSynthesisRiskEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvidenceSynthesisRiskEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


