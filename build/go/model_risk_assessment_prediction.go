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

// checks if the RiskAssessmentPrediction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskAssessmentPrediction{}

// RiskAssessmentPrediction An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessmentPrediction struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// One of the potential outcomes for the patient (e.g. remission, death,  a particular condition).
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// Indicates how likely the outcome is (in the specified timeframe).
	ProbabilityDecimal *float32 `json:"probabilityDecimal,omitempty"`
	// Indicates how likely the outcome is (in the specified timeframe).
	ProbabilityRange *Range `json:"probabilityRange,omitempty"`
	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value (e.g. low, medium, or high).
	QualitativeRisk *CodeableConcept `json:"qualitativeRisk,omitempty"`
	// A rational number with implicit precision
	RelativeRisk *float32 `json:"relativeRisk,omitempty"`
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	WhenPeriod *Period `json:"whenPeriod,omitempty"`
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	WhenRange *Range `json:"whenRange,omitempty"`
	// A sequence of Unicode characters
	Rationale *string `json:"rationale,omitempty"`
}

// NewRiskAssessmentPrediction instantiates a new RiskAssessmentPrediction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAssessmentPrediction() *RiskAssessmentPrediction {
	this := RiskAssessmentPrediction{}
	return &this
}

// NewRiskAssessmentPredictionWithDefaults instantiates a new RiskAssessmentPrediction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAssessmentPredictionWithDefaults() *RiskAssessmentPrediction {
	this := RiskAssessmentPrediction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskAssessmentPrediction) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *RiskAssessmentPrediction) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *RiskAssessmentPrediction) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetOutcome() CodeableConcept {
	if o == nil || IsNil(o.Outcome) {
		var ret CodeableConcept
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetOutcomeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given CodeableConcept and assigns it to the Outcome field.
func (o *RiskAssessmentPrediction) SetOutcome(v CodeableConcept) {
	o.Outcome = &v
}

// GetProbabilityDecimal returns the ProbabilityDecimal field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetProbabilityDecimal() float32 {
	if o == nil || IsNil(o.ProbabilityDecimal) {
		var ret float32
		return ret
	}
	return *o.ProbabilityDecimal
}

// GetProbabilityDecimalOk returns a tuple with the ProbabilityDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetProbabilityDecimalOk() (*float32, bool) {
	if o == nil || IsNil(o.ProbabilityDecimal) {
		return nil, false
	}
	return o.ProbabilityDecimal, true
}

// HasProbabilityDecimal returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasProbabilityDecimal() bool {
	if o != nil && !IsNil(o.ProbabilityDecimal) {
		return true
	}

	return false
}

// SetProbabilityDecimal gets a reference to the given float32 and assigns it to the ProbabilityDecimal field.
func (o *RiskAssessmentPrediction) SetProbabilityDecimal(v float32) {
	o.ProbabilityDecimal = &v
}

// GetProbabilityRange returns the ProbabilityRange field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetProbabilityRange() Range {
	if o == nil || IsNil(o.ProbabilityRange) {
		var ret Range
		return ret
	}
	return *o.ProbabilityRange
}

// GetProbabilityRangeOk returns a tuple with the ProbabilityRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetProbabilityRangeOk() (*Range, bool) {
	if o == nil || IsNil(o.ProbabilityRange) {
		return nil, false
	}
	return o.ProbabilityRange, true
}

// HasProbabilityRange returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasProbabilityRange() bool {
	if o != nil && !IsNil(o.ProbabilityRange) {
		return true
	}

	return false
}

// SetProbabilityRange gets a reference to the given Range and assigns it to the ProbabilityRange field.
func (o *RiskAssessmentPrediction) SetProbabilityRange(v Range) {
	o.ProbabilityRange = &v
}

// GetQualitativeRisk returns the QualitativeRisk field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetQualitativeRisk() CodeableConcept {
	if o == nil || IsNil(o.QualitativeRisk) {
		var ret CodeableConcept
		return ret
	}
	return *o.QualitativeRisk
}

// GetQualitativeRiskOk returns a tuple with the QualitativeRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetQualitativeRiskOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.QualitativeRisk) {
		return nil, false
	}
	return o.QualitativeRisk, true
}

// HasQualitativeRisk returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasQualitativeRisk() bool {
	if o != nil && !IsNil(o.QualitativeRisk) {
		return true
	}

	return false
}

// SetQualitativeRisk gets a reference to the given CodeableConcept and assigns it to the QualitativeRisk field.
func (o *RiskAssessmentPrediction) SetQualitativeRisk(v CodeableConcept) {
	o.QualitativeRisk = &v
}

// GetRelativeRisk returns the RelativeRisk field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetRelativeRisk() float32 {
	if o == nil || IsNil(o.RelativeRisk) {
		var ret float32
		return ret
	}
	return *o.RelativeRisk
}

// GetRelativeRiskOk returns a tuple with the RelativeRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetRelativeRiskOk() (*float32, bool) {
	if o == nil || IsNil(o.RelativeRisk) {
		return nil, false
	}
	return o.RelativeRisk, true
}

// HasRelativeRisk returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasRelativeRisk() bool {
	if o != nil && !IsNil(o.RelativeRisk) {
		return true
	}

	return false
}

// SetRelativeRisk gets a reference to the given float32 and assigns it to the RelativeRisk field.
func (o *RiskAssessmentPrediction) SetRelativeRisk(v float32) {
	o.RelativeRisk = &v
}

// GetWhenPeriod returns the WhenPeriod field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetWhenPeriod() Period {
	if o == nil || IsNil(o.WhenPeriod) {
		var ret Period
		return ret
	}
	return *o.WhenPeriod
}

// GetWhenPeriodOk returns a tuple with the WhenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetWhenPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.WhenPeriod) {
		return nil, false
	}
	return o.WhenPeriod, true
}

// HasWhenPeriod returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasWhenPeriod() bool {
	if o != nil && !IsNil(o.WhenPeriod) {
		return true
	}

	return false
}

// SetWhenPeriod gets a reference to the given Period and assigns it to the WhenPeriod field.
func (o *RiskAssessmentPrediction) SetWhenPeriod(v Period) {
	o.WhenPeriod = &v
}

// GetWhenRange returns the WhenRange field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetWhenRange() Range {
	if o == nil || IsNil(o.WhenRange) {
		var ret Range
		return ret
	}
	return *o.WhenRange
}

// GetWhenRangeOk returns a tuple with the WhenRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetWhenRangeOk() (*Range, bool) {
	if o == nil || IsNil(o.WhenRange) {
		return nil, false
	}
	return o.WhenRange, true
}

// HasWhenRange returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasWhenRange() bool {
	if o != nil && !IsNil(o.WhenRange) {
		return true
	}

	return false
}

// SetWhenRange gets a reference to the given Range and assigns it to the WhenRange field.
func (o *RiskAssessmentPrediction) SetWhenRange(v Range) {
	o.WhenRange = &v
}

// GetRationale returns the Rationale field value if set, zero value otherwise.
func (o *RiskAssessmentPrediction) GetRationale() string {
	if o == nil || IsNil(o.Rationale) {
		var ret string
		return ret
	}
	return *o.Rationale
}

// GetRationaleOk returns a tuple with the Rationale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessmentPrediction) GetRationaleOk() (*string, bool) {
	if o == nil || IsNil(o.Rationale) {
		return nil, false
	}
	return o.Rationale, true
}

// HasRationale returns a boolean if a field has been set.
func (o *RiskAssessmentPrediction) HasRationale() bool {
	if o != nil && !IsNil(o.Rationale) {
		return true
	}

	return false
}

// SetRationale gets a reference to the given string and assigns it to the Rationale field.
func (o *RiskAssessmentPrediction) SetRationale(v string) {
	o.Rationale = &v
}

func (o RiskAssessmentPrediction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskAssessmentPrediction) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.ProbabilityDecimal) {
		toSerialize["probabilityDecimal"] = o.ProbabilityDecimal
	}
	if !IsNil(o.ProbabilityRange) {
		toSerialize["probabilityRange"] = o.ProbabilityRange
	}
	if !IsNil(o.QualitativeRisk) {
		toSerialize["qualitativeRisk"] = o.QualitativeRisk
	}
	if !IsNil(o.RelativeRisk) {
		toSerialize["relativeRisk"] = o.RelativeRisk
	}
	if !IsNil(o.WhenPeriod) {
		toSerialize["whenPeriod"] = o.WhenPeriod
	}
	if !IsNil(o.WhenRange) {
		toSerialize["whenRange"] = o.WhenRange
	}
	if !IsNil(o.Rationale) {
		toSerialize["rationale"] = o.Rationale
	}
	return toSerialize, nil
}

type NullableRiskAssessmentPrediction struct {
	value *RiskAssessmentPrediction
	isSet bool
}

func (v NullableRiskAssessmentPrediction) Get() *RiskAssessmentPrediction {
	return v.value
}

func (v *NullableRiskAssessmentPrediction) Set(val *RiskAssessmentPrediction) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAssessmentPrediction) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAssessmentPrediction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAssessmentPrediction(val *RiskAssessmentPrediction) *NullableRiskAssessmentPrediction {
	return &NullableRiskAssessmentPrediction{value: val, isSet: true}
}

func (v NullableRiskAssessmentPrediction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAssessmentPrediction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


