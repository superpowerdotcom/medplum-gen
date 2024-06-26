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

// checks if the MedicationKnowledgeAdministrationGuidelines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationKnowledgeAdministrationGuidelines{}

// MedicationKnowledgeAdministrationGuidelines Information about a medication that is used to support knowledge.
type MedicationKnowledgeAdministrationGuidelines struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Dosage for the medication for the specific guidelines.
	Dosage []MedicationKnowledgeDosage `json:"dosage,omitempty"`
	// Indication for use that apply to the specific administration guidelines.
	IndicationCodeableConcept *CodeableConcept `json:"indicationCodeableConcept,omitempty"`
	// Indication for use that apply to the specific administration guidelines.
	IndicationReference *Reference `json:"indicationReference,omitempty"`
	// Characteristics of the patient that are relevant to the administration guidelines (for example, height, weight, gender, etc.).
	PatientCharacteristics []MedicationKnowledgePatientCharacteristics `json:"patientCharacteristics,omitempty"`
}

// NewMedicationKnowledgeAdministrationGuidelines instantiates a new MedicationKnowledgeAdministrationGuidelines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationKnowledgeAdministrationGuidelines() *MedicationKnowledgeAdministrationGuidelines {
	this := MedicationKnowledgeAdministrationGuidelines{}
	return &this
}

// NewMedicationKnowledgeAdministrationGuidelinesWithDefaults instantiates a new MedicationKnowledgeAdministrationGuidelines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationKnowledgeAdministrationGuidelinesWithDefaults() *MedicationKnowledgeAdministrationGuidelines {
	this := MedicationKnowledgeAdministrationGuidelines{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDosage returns the Dosage field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetDosage() []MedicationKnowledgeDosage {
	if o == nil || IsNil(o.Dosage) {
		var ret []MedicationKnowledgeDosage
		return ret
	}
	return o.Dosage
}

// GetDosageOk returns a tuple with the Dosage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetDosageOk() ([]MedicationKnowledgeDosage, bool) {
	if o == nil || IsNil(o.Dosage) {
		return nil, false
	}
	return o.Dosage, true
}

// HasDosage returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasDosage() bool {
	if o != nil && !IsNil(o.Dosage) {
		return true
	}

	return false
}

// SetDosage gets a reference to the given []MedicationKnowledgeDosage and assigns it to the Dosage field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetDosage(v []MedicationKnowledgeDosage) {
	o.Dosage = v
}

// GetIndicationCodeableConcept returns the IndicationCodeableConcept field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetIndicationCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.IndicationCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.IndicationCodeableConcept
}

// GetIndicationCodeableConceptOk returns a tuple with the IndicationCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetIndicationCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.IndicationCodeableConcept) {
		return nil, false
	}
	return o.IndicationCodeableConcept, true
}

// HasIndicationCodeableConcept returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasIndicationCodeableConcept() bool {
	if o != nil && !IsNil(o.IndicationCodeableConcept) {
		return true
	}

	return false
}

// SetIndicationCodeableConcept gets a reference to the given CodeableConcept and assigns it to the IndicationCodeableConcept field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetIndicationCodeableConcept(v CodeableConcept) {
	o.IndicationCodeableConcept = &v
}

// GetIndicationReference returns the IndicationReference field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetIndicationReference() Reference {
	if o == nil || IsNil(o.IndicationReference) {
		var ret Reference
		return ret
	}
	return *o.IndicationReference
}

// GetIndicationReferenceOk returns a tuple with the IndicationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetIndicationReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.IndicationReference) {
		return nil, false
	}
	return o.IndicationReference, true
}

// HasIndicationReference returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasIndicationReference() bool {
	if o != nil && !IsNil(o.IndicationReference) {
		return true
	}

	return false
}

// SetIndicationReference gets a reference to the given Reference and assigns it to the IndicationReference field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetIndicationReference(v Reference) {
	o.IndicationReference = &v
}

// GetPatientCharacteristics returns the PatientCharacteristics field value if set, zero value otherwise.
func (o *MedicationKnowledgeAdministrationGuidelines) GetPatientCharacteristics() []MedicationKnowledgePatientCharacteristics {
	if o == nil || IsNil(o.PatientCharacteristics) {
		var ret []MedicationKnowledgePatientCharacteristics
		return ret
	}
	return o.PatientCharacteristics
}

// GetPatientCharacteristicsOk returns a tuple with the PatientCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) GetPatientCharacteristicsOk() ([]MedicationKnowledgePatientCharacteristics, bool) {
	if o == nil || IsNil(o.PatientCharacteristics) {
		return nil, false
	}
	return o.PatientCharacteristics, true
}

// HasPatientCharacteristics returns a boolean if a field has been set.
func (o *MedicationKnowledgeAdministrationGuidelines) HasPatientCharacteristics() bool {
	if o != nil && !IsNil(o.PatientCharacteristics) {
		return true
	}

	return false
}

// SetPatientCharacteristics gets a reference to the given []MedicationKnowledgePatientCharacteristics and assigns it to the PatientCharacteristics field.
func (o *MedicationKnowledgeAdministrationGuidelines) SetPatientCharacteristics(v []MedicationKnowledgePatientCharacteristics) {
	o.PatientCharacteristics = v
}

func (o MedicationKnowledgeAdministrationGuidelines) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationKnowledgeAdministrationGuidelines) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Dosage) {
		toSerialize["dosage"] = o.Dosage
	}
	if !IsNil(o.IndicationCodeableConcept) {
		toSerialize["indicationCodeableConcept"] = o.IndicationCodeableConcept
	}
	if !IsNil(o.IndicationReference) {
		toSerialize["indicationReference"] = o.IndicationReference
	}
	if !IsNil(o.PatientCharacteristics) {
		toSerialize["patientCharacteristics"] = o.PatientCharacteristics
	}
	return toSerialize, nil
}

type NullableMedicationKnowledgeAdministrationGuidelines struct {
	value *MedicationKnowledgeAdministrationGuidelines
	isSet bool
}

func (v NullableMedicationKnowledgeAdministrationGuidelines) Get() *MedicationKnowledgeAdministrationGuidelines {
	return v.value
}

func (v *NullableMedicationKnowledgeAdministrationGuidelines) Set(val *MedicationKnowledgeAdministrationGuidelines) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationKnowledgeAdministrationGuidelines) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationKnowledgeAdministrationGuidelines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationKnowledgeAdministrationGuidelines(val *MedicationKnowledgeAdministrationGuidelines) *NullableMedicationKnowledgeAdministrationGuidelines {
	return &NullableMedicationKnowledgeAdministrationGuidelines{value: val, isSet: true}
}

func (v NullableMedicationKnowledgeAdministrationGuidelines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationKnowledgeAdministrationGuidelines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

