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

// checks if the EncounterHospitalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncounterHospitalization{}

// EncounterHospitalization An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type EncounterHospitalization struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pre-admission identifier.
	PreAdmissionIdentifier *Identifier `json:"preAdmissionIdentifier,omitempty"`
	// The location/organization from which the patient came before admission.
	Origin *Reference `json:"origin,omitempty"`
	// From where patient was admitted (physician referral, transfer).
	AdmitSource *CodeableConcept `json:"admitSource,omitempty"`
	// Whether this hospitalization is a readmission and why if known.
	ReAdmission *CodeableConcept `json:"reAdmission,omitempty"`
	// Diet preferences reported by the patient.
	DietPreference []CodeableConcept `json:"dietPreference,omitempty"`
	// Special courtesies (VIP, board member).
	SpecialCourtesy []CodeableConcept `json:"specialCourtesy,omitempty"`
	// Any special requests that have been made for this hospitalization encounter, such as the provision of specific equipment or other things.
	SpecialArrangement []CodeableConcept `json:"specialArrangement,omitempty"`
	// Location/organization to which the patient is discharged.
	Destination *Reference `json:"destination,omitempty"`
	// Category or kind of location after discharge.
	DischargeDisposition *CodeableConcept `json:"dischargeDisposition,omitempty"`
}

// NewEncounterHospitalization instantiates a new EncounterHospitalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncounterHospitalization() *EncounterHospitalization {
	this := EncounterHospitalization{}
	return &this
}

// NewEncounterHospitalizationWithDefaults instantiates a new EncounterHospitalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncounterHospitalizationWithDefaults() *EncounterHospitalization {
	this := EncounterHospitalization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EncounterHospitalization) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *EncounterHospitalization) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *EncounterHospitalization) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetPreAdmissionIdentifier returns the PreAdmissionIdentifier field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetPreAdmissionIdentifier() Identifier {
	if o == nil || IsNil(o.PreAdmissionIdentifier) {
		var ret Identifier
		return ret
	}
	return *o.PreAdmissionIdentifier
}

// GetPreAdmissionIdentifierOk returns a tuple with the PreAdmissionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetPreAdmissionIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.PreAdmissionIdentifier) {
		return nil, false
	}
	return o.PreAdmissionIdentifier, true
}

// HasPreAdmissionIdentifier returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasPreAdmissionIdentifier() bool {
	if o != nil && !IsNil(o.PreAdmissionIdentifier) {
		return true
	}

	return false
}

// SetPreAdmissionIdentifier gets a reference to the given Identifier and assigns it to the PreAdmissionIdentifier field.
func (o *EncounterHospitalization) SetPreAdmissionIdentifier(v Identifier) {
	o.PreAdmissionIdentifier = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetOrigin() Reference {
	if o == nil || IsNil(o.Origin) {
		var ret Reference
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetOriginOk() (*Reference, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given Reference and assigns it to the Origin field.
func (o *EncounterHospitalization) SetOrigin(v Reference) {
	o.Origin = &v
}

// GetAdmitSource returns the AdmitSource field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetAdmitSource() CodeableConcept {
	if o == nil || IsNil(o.AdmitSource) {
		var ret CodeableConcept
		return ret
	}
	return *o.AdmitSource
}

// GetAdmitSourceOk returns a tuple with the AdmitSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetAdmitSourceOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.AdmitSource) {
		return nil, false
	}
	return o.AdmitSource, true
}

// HasAdmitSource returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasAdmitSource() bool {
	if o != nil && !IsNil(o.AdmitSource) {
		return true
	}

	return false
}

// SetAdmitSource gets a reference to the given CodeableConcept and assigns it to the AdmitSource field.
func (o *EncounterHospitalization) SetAdmitSource(v CodeableConcept) {
	o.AdmitSource = &v
}

// GetReAdmission returns the ReAdmission field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetReAdmission() CodeableConcept {
	if o == nil || IsNil(o.ReAdmission) {
		var ret CodeableConcept
		return ret
	}
	return *o.ReAdmission
}

// GetReAdmissionOk returns a tuple with the ReAdmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetReAdmissionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ReAdmission) {
		return nil, false
	}
	return o.ReAdmission, true
}

// HasReAdmission returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasReAdmission() bool {
	if o != nil && !IsNil(o.ReAdmission) {
		return true
	}

	return false
}

// SetReAdmission gets a reference to the given CodeableConcept and assigns it to the ReAdmission field.
func (o *EncounterHospitalization) SetReAdmission(v CodeableConcept) {
	o.ReAdmission = &v
}

// GetDietPreference returns the DietPreference field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetDietPreference() []CodeableConcept {
	if o == nil || IsNil(o.DietPreference) {
		var ret []CodeableConcept
		return ret
	}
	return o.DietPreference
}

// GetDietPreferenceOk returns a tuple with the DietPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetDietPreferenceOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.DietPreference) {
		return nil, false
	}
	return o.DietPreference, true
}

// HasDietPreference returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasDietPreference() bool {
	if o != nil && !IsNil(o.DietPreference) {
		return true
	}

	return false
}

// SetDietPreference gets a reference to the given []CodeableConcept and assigns it to the DietPreference field.
func (o *EncounterHospitalization) SetDietPreference(v []CodeableConcept) {
	o.DietPreference = v
}

// GetSpecialCourtesy returns the SpecialCourtesy field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetSpecialCourtesy() []CodeableConcept {
	if o == nil || IsNil(o.SpecialCourtesy) {
		var ret []CodeableConcept
		return ret
	}
	return o.SpecialCourtesy
}

// GetSpecialCourtesyOk returns a tuple with the SpecialCourtesy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetSpecialCourtesyOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.SpecialCourtesy) {
		return nil, false
	}
	return o.SpecialCourtesy, true
}

// HasSpecialCourtesy returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasSpecialCourtesy() bool {
	if o != nil && !IsNil(o.SpecialCourtesy) {
		return true
	}

	return false
}

// SetSpecialCourtesy gets a reference to the given []CodeableConcept and assigns it to the SpecialCourtesy field.
func (o *EncounterHospitalization) SetSpecialCourtesy(v []CodeableConcept) {
	o.SpecialCourtesy = v
}

// GetSpecialArrangement returns the SpecialArrangement field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetSpecialArrangement() []CodeableConcept {
	if o == nil || IsNil(o.SpecialArrangement) {
		var ret []CodeableConcept
		return ret
	}
	return o.SpecialArrangement
}

// GetSpecialArrangementOk returns a tuple with the SpecialArrangement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetSpecialArrangementOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.SpecialArrangement) {
		return nil, false
	}
	return o.SpecialArrangement, true
}

// HasSpecialArrangement returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasSpecialArrangement() bool {
	if o != nil && !IsNil(o.SpecialArrangement) {
		return true
	}

	return false
}

// SetSpecialArrangement gets a reference to the given []CodeableConcept and assigns it to the SpecialArrangement field.
func (o *EncounterHospitalization) SetSpecialArrangement(v []CodeableConcept) {
	o.SpecialArrangement = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetDestination() Reference {
	if o == nil || IsNil(o.Destination) {
		var ret Reference
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetDestinationOk() (*Reference, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Reference and assigns it to the Destination field.
func (o *EncounterHospitalization) SetDestination(v Reference) {
	o.Destination = &v
}

// GetDischargeDisposition returns the DischargeDisposition field value if set, zero value otherwise.
func (o *EncounterHospitalization) GetDischargeDisposition() CodeableConcept {
	if o == nil || IsNil(o.DischargeDisposition) {
		var ret CodeableConcept
		return ret
	}
	return *o.DischargeDisposition
}

// GetDischargeDispositionOk returns a tuple with the DischargeDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncounterHospitalization) GetDischargeDispositionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.DischargeDisposition) {
		return nil, false
	}
	return o.DischargeDisposition, true
}

// HasDischargeDisposition returns a boolean if a field has been set.
func (o *EncounterHospitalization) HasDischargeDisposition() bool {
	if o != nil && !IsNil(o.DischargeDisposition) {
		return true
	}

	return false
}

// SetDischargeDisposition gets a reference to the given CodeableConcept and assigns it to the DischargeDisposition field.
func (o *EncounterHospitalization) SetDischargeDisposition(v CodeableConcept) {
	o.DischargeDisposition = &v
}

func (o EncounterHospitalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncounterHospitalization) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PreAdmissionIdentifier) {
		toSerialize["preAdmissionIdentifier"] = o.PreAdmissionIdentifier
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.AdmitSource) {
		toSerialize["admitSource"] = o.AdmitSource
	}
	if !IsNil(o.ReAdmission) {
		toSerialize["reAdmission"] = o.ReAdmission
	}
	if !IsNil(o.DietPreference) {
		toSerialize["dietPreference"] = o.DietPreference
	}
	if !IsNil(o.SpecialCourtesy) {
		toSerialize["specialCourtesy"] = o.SpecialCourtesy
	}
	if !IsNil(o.SpecialArrangement) {
		toSerialize["specialArrangement"] = o.SpecialArrangement
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.DischargeDisposition) {
		toSerialize["dischargeDisposition"] = o.DischargeDisposition
	}
	return toSerialize, nil
}

type NullableEncounterHospitalization struct {
	value *EncounterHospitalization
	isSet bool
}

func (v NullableEncounterHospitalization) Get() *EncounterHospitalization {
	return v.value
}

func (v *NullableEncounterHospitalization) Set(val *EncounterHospitalization) {
	v.value = val
	v.isSet = true
}

func (v NullableEncounterHospitalization) IsSet() bool {
	return v.isSet
}

func (v *NullableEncounterHospitalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncounterHospitalization(val *EncounterHospitalization) *NullableEncounterHospitalization {
	return &NullableEncounterHospitalization{value: val, isSet: true}
}

func (v NullableEncounterHospitalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncounterHospitalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


