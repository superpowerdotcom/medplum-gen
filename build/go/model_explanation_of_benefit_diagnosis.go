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

// checks if the ExplanationOfBenefitDiagnosis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplanationOfBenefitDiagnosis{}

// ExplanationOfBenefitDiagnosis This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitDiagnosis struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Sequence *float32 `json:"sequence,omitempty"`
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	DiagnosisReference *Reference `json:"diagnosisReference,omitempty"`
	// When the condition was observed or the relative ranking.
	Type []CodeableConcept `json:"type,omitempty"`
	// Indication of whether the diagnosis was present on admission to a facility.
	OnAdmission *CodeableConcept `json:"onAdmission,omitempty"`
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	PackageCode *CodeableConcept `json:"packageCode,omitempty"`
}

// NewExplanationOfBenefitDiagnosis instantiates a new ExplanationOfBenefitDiagnosis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplanationOfBenefitDiagnosis() *ExplanationOfBenefitDiagnosis {
	this := ExplanationOfBenefitDiagnosis{}
	return &this
}

// NewExplanationOfBenefitDiagnosisWithDefaults instantiates a new ExplanationOfBenefitDiagnosis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplanationOfBenefitDiagnosisWithDefaults() *ExplanationOfBenefitDiagnosis {
	this := ExplanationOfBenefitDiagnosis{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExplanationOfBenefitDiagnosis) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExplanationOfBenefitDiagnosis) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExplanationOfBenefitDiagnosis) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetSequence() float32 {
	if o == nil || IsNil(o.Sequence) {
		var ret float32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given float32 and assigns it to the Sequence field.
func (o *ExplanationOfBenefitDiagnosis) SetSequence(v float32) {
	o.Sequence = &v
}

// GetDiagnosisCodeableConcept returns the DiagnosisCodeableConcept field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetDiagnosisCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.DiagnosisCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.DiagnosisCodeableConcept
}

// GetDiagnosisCodeableConceptOk returns a tuple with the DiagnosisCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetDiagnosisCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.DiagnosisCodeableConcept) {
		return nil, false
	}
	return o.DiagnosisCodeableConcept, true
}

// HasDiagnosisCodeableConcept returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasDiagnosisCodeableConcept() bool {
	if o != nil && !IsNil(o.DiagnosisCodeableConcept) {
		return true
	}

	return false
}

// SetDiagnosisCodeableConcept gets a reference to the given CodeableConcept and assigns it to the DiagnosisCodeableConcept field.
func (o *ExplanationOfBenefitDiagnosis) SetDiagnosisCodeableConcept(v CodeableConcept) {
	o.DiagnosisCodeableConcept = &v
}

// GetDiagnosisReference returns the DiagnosisReference field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetDiagnosisReference() Reference {
	if o == nil || IsNil(o.DiagnosisReference) {
		var ret Reference
		return ret
	}
	return *o.DiagnosisReference
}

// GetDiagnosisReferenceOk returns a tuple with the DiagnosisReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetDiagnosisReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.DiagnosisReference) {
		return nil, false
	}
	return o.DiagnosisReference, true
}

// HasDiagnosisReference returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasDiagnosisReference() bool {
	if o != nil && !IsNil(o.DiagnosisReference) {
		return true
	}

	return false
}

// SetDiagnosisReference gets a reference to the given Reference and assigns it to the DiagnosisReference field.
func (o *ExplanationOfBenefitDiagnosis) SetDiagnosisReference(v Reference) {
	o.DiagnosisReference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetType() []CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret []CodeableConcept
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []CodeableConcept and assigns it to the Type field.
func (o *ExplanationOfBenefitDiagnosis) SetType(v []CodeableConcept) {
	o.Type = v
}

// GetOnAdmission returns the OnAdmission field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetOnAdmission() CodeableConcept {
	if o == nil || IsNil(o.OnAdmission) {
		var ret CodeableConcept
		return ret
	}
	return *o.OnAdmission
}

// GetOnAdmissionOk returns a tuple with the OnAdmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetOnAdmissionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.OnAdmission) {
		return nil, false
	}
	return o.OnAdmission, true
}

// HasOnAdmission returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasOnAdmission() bool {
	if o != nil && !IsNil(o.OnAdmission) {
		return true
	}

	return false
}

// SetOnAdmission gets a reference to the given CodeableConcept and assigns it to the OnAdmission field.
func (o *ExplanationOfBenefitDiagnosis) SetOnAdmission(v CodeableConcept) {
	o.OnAdmission = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *ExplanationOfBenefitDiagnosis) GetPackageCode() CodeableConcept {
	if o == nil || IsNil(o.PackageCode) {
		var ret CodeableConcept
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitDiagnosis) GetPackageCodeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *ExplanationOfBenefitDiagnosis) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given CodeableConcept and assigns it to the PackageCode field.
func (o *ExplanationOfBenefitDiagnosis) SetPackageCode(v CodeableConcept) {
	o.PackageCode = &v
}

func (o ExplanationOfBenefitDiagnosis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplanationOfBenefitDiagnosis) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.DiagnosisCodeableConcept) {
		toSerialize["diagnosisCodeableConcept"] = o.DiagnosisCodeableConcept
	}
	if !IsNil(o.DiagnosisReference) {
		toSerialize["diagnosisReference"] = o.DiagnosisReference
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.OnAdmission) {
		toSerialize["onAdmission"] = o.OnAdmission
	}
	if !IsNil(o.PackageCode) {
		toSerialize["packageCode"] = o.PackageCode
	}
	return toSerialize, nil
}

type NullableExplanationOfBenefitDiagnosis struct {
	value *ExplanationOfBenefitDiagnosis
	isSet bool
}

func (v NullableExplanationOfBenefitDiagnosis) Get() *ExplanationOfBenefitDiagnosis {
	return v.value
}

func (v *NullableExplanationOfBenefitDiagnosis) Set(val *ExplanationOfBenefitDiagnosis) {
	v.value = val
	v.isSet = true
}

func (v NullableExplanationOfBenefitDiagnosis) IsSet() bool {
	return v.isSet
}

func (v *NullableExplanationOfBenefitDiagnosis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplanationOfBenefitDiagnosis(val *ExplanationOfBenefitDiagnosis) *NullableExplanationOfBenefitDiagnosis {
	return &NullableExplanationOfBenefitDiagnosis{value: val, isSet: true}
}

func (v NullableExplanationOfBenefitDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplanationOfBenefitDiagnosis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


