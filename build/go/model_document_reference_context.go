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

// checks if the DocumentReferenceContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentReferenceContext{}

// DocumentReferenceContext A reference to a document of any kind for any purpose. Provides metadata about the document so that the document can be discovered and managed. The scope of a document is any seralized object with a mime-type, so includes formal patient centric documents (CDA), cliical notes, scanned paper, and non-patient specific documents like policy text.
type DocumentReferenceContext struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Describes the clinical encounter or type of care that the document content is associated with.
	Encounter []Reference `json:"encounter,omitempty"`
	// This list of codes represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented. In some cases, the event is inherent in the type Code, such as a \"History and Physical Report\" in which the procedure being documented is necessarily a \"History and Physical\" act.
	Event []CodeableConcept `json:"event,omitempty"`
	// The time period over which the service that is described by the document was provided.
	Period *Period `json:"period,omitempty"`
	// The kind of facility where the patient was seen.
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	// This property may convey specifics about the practice setting where the content was created, often reflecting the clinical specialty.
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	// The Patient Information as known when the document was published. May be a reference to a version specific, or contained.
	SourcePatientInfo *Reference `json:"sourcePatientInfo,omitempty"`
	// Related identifiers or resources associated with the DocumentReference.
	Related []Reference `json:"related,omitempty"`
}

// NewDocumentReferenceContext instantiates a new DocumentReferenceContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentReferenceContext() *DocumentReferenceContext {
	this := DocumentReferenceContext{}
	return &this
}

// NewDocumentReferenceContextWithDefaults instantiates a new DocumentReferenceContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentReferenceContextWithDefaults() *DocumentReferenceContext {
	this := DocumentReferenceContext{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentReferenceContext) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DocumentReferenceContext) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DocumentReferenceContext) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetEncounter() []Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret []Reference
		return ret
	}
	return o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetEncounterOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given []Reference and assigns it to the Encounter field.
func (o *DocumentReferenceContext) SetEncounter(v []Reference) {
	o.Encounter = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetEvent() []CodeableConcept {
	if o == nil || IsNil(o.Event) {
		var ret []CodeableConcept
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetEventOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given []CodeableConcept and assigns it to the Event field.
func (o *DocumentReferenceContext) SetEvent(v []CodeableConcept) {
	o.Event = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *DocumentReferenceContext) SetPeriod(v Period) {
	o.Period = &v
}

// GetFacilityType returns the FacilityType field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetFacilityType() CodeableConcept {
	if o == nil || IsNil(o.FacilityType) {
		var ret CodeableConcept
		return ret
	}
	return *o.FacilityType
}

// GetFacilityTypeOk returns a tuple with the FacilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetFacilityTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.FacilityType) {
		return nil, false
	}
	return o.FacilityType, true
}

// HasFacilityType returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasFacilityType() bool {
	if o != nil && !IsNil(o.FacilityType) {
		return true
	}

	return false
}

// SetFacilityType gets a reference to the given CodeableConcept and assigns it to the FacilityType field.
func (o *DocumentReferenceContext) SetFacilityType(v CodeableConcept) {
	o.FacilityType = &v
}

// GetPracticeSetting returns the PracticeSetting field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetPracticeSetting() CodeableConcept {
	if o == nil || IsNil(o.PracticeSetting) {
		var ret CodeableConcept
		return ret
	}
	return *o.PracticeSetting
}

// GetPracticeSettingOk returns a tuple with the PracticeSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetPracticeSettingOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.PracticeSetting) {
		return nil, false
	}
	return o.PracticeSetting, true
}

// HasPracticeSetting returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasPracticeSetting() bool {
	if o != nil && !IsNil(o.PracticeSetting) {
		return true
	}

	return false
}

// SetPracticeSetting gets a reference to the given CodeableConcept and assigns it to the PracticeSetting field.
func (o *DocumentReferenceContext) SetPracticeSetting(v CodeableConcept) {
	o.PracticeSetting = &v
}

// GetSourcePatientInfo returns the SourcePatientInfo field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetSourcePatientInfo() Reference {
	if o == nil || IsNil(o.SourcePatientInfo) {
		var ret Reference
		return ret
	}
	return *o.SourcePatientInfo
}

// GetSourcePatientInfoOk returns a tuple with the SourcePatientInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetSourcePatientInfoOk() (*Reference, bool) {
	if o == nil || IsNil(o.SourcePatientInfo) {
		return nil, false
	}
	return o.SourcePatientInfo, true
}

// HasSourcePatientInfo returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasSourcePatientInfo() bool {
	if o != nil && !IsNil(o.SourcePatientInfo) {
		return true
	}

	return false
}

// SetSourcePatientInfo gets a reference to the given Reference and assigns it to the SourcePatientInfo field.
func (o *DocumentReferenceContext) SetSourcePatientInfo(v Reference) {
	o.SourcePatientInfo = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *DocumentReferenceContext) GetRelated() []Reference {
	if o == nil || IsNil(o.Related) {
		var ret []Reference
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContext) GetRelatedOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *DocumentReferenceContext) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given []Reference and assigns it to the Related field.
func (o *DocumentReferenceContext) SetRelated(v []Reference) {
	o.Related = v
}

func (o DocumentReferenceContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentReferenceContext) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.FacilityType) {
		toSerialize["facilityType"] = o.FacilityType
	}
	if !IsNil(o.PracticeSetting) {
		toSerialize["practiceSetting"] = o.PracticeSetting
	}
	if !IsNil(o.SourcePatientInfo) {
		toSerialize["sourcePatientInfo"] = o.SourcePatientInfo
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	return toSerialize, nil
}

type NullableDocumentReferenceContext struct {
	value *DocumentReferenceContext
	isSet bool
}

func (v NullableDocumentReferenceContext) Get() *DocumentReferenceContext {
	return v.value
}

func (v *NullableDocumentReferenceContext) Set(val *DocumentReferenceContext) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentReferenceContext) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentReferenceContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentReferenceContext(val *DocumentReferenceContext) *NullableDocumentReferenceContext {
	return &NullableDocumentReferenceContext{value: val, isSet: true}
}

func (v NullableDocumentReferenceContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentReferenceContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


