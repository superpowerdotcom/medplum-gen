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

// checks if the Encounter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Encounter{}

// Encounter An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter struct {
	// This is a Encounter resource
	ResourceType string `json:"resourceType"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Id *string `json:"id,omitempty"`
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `json:"meta,omitempty"`
	// String of characters used to identify a name or a resource
	ImplicitRules *string `json:"implicitRules,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Language *string `json:"language,omitempty"`
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it \"clinically safe\" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative `json:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []ResourceList `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifier(s) by which this encounter is known.
	Identifier []Identifier `json:"identifier,omitempty"`
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status *string `json:"status,omitempty"`
	// The status history permits the encounter resource to contain the status history without needing to read through the historical versions of the resource, or even have the server store them.
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class Coding `json:"class"`
	// The class history permits the tracking of the encounters transitions without needing to go  through the resource history.  This would be used for a case where an admission starts of as an emergency encounter, then transitions into an inpatient scenario. Doing this and not restarting a new encounter ensures that any lab/diagnostic results can more easily follow the patient and not require re-processing and not get lost or cancelled during a kind of discharge from emergency to inpatient.
	ClassHistory []EncounterClassHistory `json:"classHistory,omitempty"`
	// Specific type of encounter (e.g. e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	Type []CodeableConcept `json:"type,omitempty"`
	// Broad categorization of the service that is to be provided (e.g. cardiology).
	ServiceType *CodeableConcept `json:"serviceType,omitempty"`
	// Indicates the urgency of the encounter.
	Priority *CodeableConcept `json:"priority,omitempty"`
	// The patient or group present at the encounter.
	Subject *Reference `json:"subject,omitempty"`
	// Where a specific encounter should be classified as a part of a specific episode(s) of care this field should be used. This association can facilitate grouping of related encounters together for a specific purpose, such as government reporting, issue tracking, association via a common problem.  The association is recorded on the encounter as these are typically created after the episode of care and grouped on entry rather than editing the episode of care to append another encounter to it (the episode of care could span years).
	EpisodeOfCare []Reference `json:"episodeOfCare,omitempty"`
	// The request this encounter satisfies (e.g. incoming referral or procedure request).
	BasedOn []Reference `json:"basedOn,omitempty"`
	// The list of people responsible for providing the service.
	Participant []EncounterParticipant `json:"participant,omitempty"`
	// The appointment that scheduled this encounter.
	Appointment []Reference `json:"appointment,omitempty"`
	// The start and end time of the encounter.
	Period *Period `json:"period,omitempty"`
	// Quantity of time the encounter lasted. This excludes the time during leaves of absence.
	Length *Duration `json:"length,omitempty"`
	// Reason the encounter takes place, expressed as a code. For admissions, this can be used for a coded admission diagnosis.
	ReasonCode []CodeableConcept `json:"reasonCode,omitempty"`
	// Reason the encounter takes place, expressed as a code. For admissions, this can be used for a coded admission diagnosis.
	ReasonReference []Reference `json:"reasonReference,omitempty"`
	// The list of diagnosis relevant to this encounter.
	Diagnosis []EncounterDiagnosis `json:"diagnosis,omitempty"`
	// The set of accounts that may be used for billing for this Encounter.
	Account []Reference `json:"account,omitempty"`
	// Details about the admission to a healthcare service.
	Hospitalization *EncounterHospitalization `json:"hospitalization,omitempty"`
	// List of locations where  the patient has been during this encounter.
	Location []EncounterLocation `json:"location,omitempty"`
	// The organization that is primarily responsible for this Encounter's services. This MAY be the same as the organization on the Patient record, however it could be different, such as if the actor performing the services was from an external organization (which may be billed seperately) for an external consultation.  Refer to the example bundle showing an abbreviated set of Encounters for a colonoscopy.
	ServiceProvider *Reference `json:"serviceProvider,omitempty"`
	// Another Encounter of which this encounter is a part of (administratively or in time).
	PartOf *Reference `json:"partOf,omitempty"`
}

type _Encounter Encounter

// NewEncounter instantiates a new Encounter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncounter(resourceType string, class Coding) *Encounter {
	this := Encounter{}
	this.ResourceType = resourceType
	this.Class = class
	return &this
}

// NewEncounterWithDefaults instantiates a new Encounter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncounterWithDefaults() *Encounter {
	this := Encounter{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Encounter) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Encounter) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Encounter) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Encounter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Encounter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Encounter) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Encounter) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Encounter) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Encounter) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Encounter) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Encounter) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Encounter) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Encounter) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Encounter) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Encounter) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Encounter) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Encounter) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Encounter) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Encounter) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Encounter) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Encounter) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Encounter) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Encounter) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Encounter) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Encounter) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Encounter) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Encounter) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Encounter) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Encounter) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Encounter) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Encounter) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Encounter) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Encounter) SetStatus(v string) {
	o.Status = &v
}

// GetStatusHistory returns the StatusHistory field value if set, zero value otherwise.
func (o *Encounter) GetStatusHistory() []EncounterStatusHistory {
	if o == nil || IsNil(o.StatusHistory) {
		var ret []EncounterStatusHistory
		return ret
	}
	return o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetStatusHistoryOk() ([]EncounterStatusHistory, bool) {
	if o == nil || IsNil(o.StatusHistory) {
		return nil, false
	}
	return o.StatusHistory, true
}

// HasStatusHistory returns a boolean if a field has been set.
func (o *Encounter) HasStatusHistory() bool {
	if o != nil && !IsNil(o.StatusHistory) {
		return true
	}

	return false
}

// SetStatusHistory gets a reference to the given []EncounterStatusHistory and assigns it to the StatusHistory field.
func (o *Encounter) SetStatusHistory(v []EncounterStatusHistory) {
	o.StatusHistory = v
}

// GetClass returns the Class field value
func (o *Encounter) GetClass() Coding {
	if o == nil {
		var ret Coding
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *Encounter) GetClassOk() (*Coding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *Encounter) SetClass(v Coding) {
	o.Class = v
}

// GetClassHistory returns the ClassHistory field value if set, zero value otherwise.
func (o *Encounter) GetClassHistory() []EncounterClassHistory {
	if o == nil || IsNil(o.ClassHistory) {
		var ret []EncounterClassHistory
		return ret
	}
	return o.ClassHistory
}

// GetClassHistoryOk returns a tuple with the ClassHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetClassHistoryOk() ([]EncounterClassHistory, bool) {
	if o == nil || IsNil(o.ClassHistory) {
		return nil, false
	}
	return o.ClassHistory, true
}

// HasClassHistory returns a boolean if a field has been set.
func (o *Encounter) HasClassHistory() bool {
	if o != nil && !IsNil(o.ClassHistory) {
		return true
	}

	return false
}

// SetClassHistory gets a reference to the given []EncounterClassHistory and assigns it to the ClassHistory field.
func (o *Encounter) SetClassHistory(v []EncounterClassHistory) {
	o.ClassHistory = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Encounter) GetType() []CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret []CodeableConcept
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Encounter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []CodeableConcept and assigns it to the Type field.
func (o *Encounter) SetType(v []CodeableConcept) {
	o.Type = v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *Encounter) GetServiceType() CodeableConcept {
	if o == nil || IsNil(o.ServiceType) {
		var ret CodeableConcept
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetServiceTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *Encounter) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given CodeableConcept and assigns it to the ServiceType field.
func (o *Encounter) SetServiceType(v CodeableConcept) {
	o.ServiceType = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Encounter) GetPriority() CodeableConcept {
	if o == nil || IsNil(o.Priority) {
		var ret CodeableConcept
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetPriorityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Encounter) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given CodeableConcept and assigns it to the Priority field.
func (o *Encounter) SetPriority(v CodeableConcept) {
	o.Priority = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Encounter) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Encounter) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *Encounter) SetSubject(v Reference) {
	o.Subject = &v
}

// GetEpisodeOfCare returns the EpisodeOfCare field value if set, zero value otherwise.
func (o *Encounter) GetEpisodeOfCare() []Reference {
	if o == nil || IsNil(o.EpisodeOfCare) {
		var ret []Reference
		return ret
	}
	return o.EpisodeOfCare
}

// GetEpisodeOfCareOk returns a tuple with the EpisodeOfCare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetEpisodeOfCareOk() ([]Reference, bool) {
	if o == nil || IsNil(o.EpisodeOfCare) {
		return nil, false
	}
	return o.EpisodeOfCare, true
}

// HasEpisodeOfCare returns a boolean if a field has been set.
func (o *Encounter) HasEpisodeOfCare() bool {
	if o != nil && !IsNil(o.EpisodeOfCare) {
		return true
	}

	return false
}

// SetEpisodeOfCare gets a reference to the given []Reference and assigns it to the EpisodeOfCare field.
func (o *Encounter) SetEpisodeOfCare(v []Reference) {
	o.EpisodeOfCare = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *Encounter) GetBasedOn() []Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret []Reference
		return ret
	}
	return o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetBasedOnOk() ([]Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *Encounter) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given []Reference and assigns it to the BasedOn field.
func (o *Encounter) SetBasedOn(v []Reference) {
	o.BasedOn = v
}

// GetParticipant returns the Participant field value if set, zero value otherwise.
func (o *Encounter) GetParticipant() []EncounterParticipant {
	if o == nil || IsNil(o.Participant) {
		var ret []EncounterParticipant
		return ret
	}
	return o.Participant
}

// GetParticipantOk returns a tuple with the Participant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetParticipantOk() ([]EncounterParticipant, bool) {
	if o == nil || IsNil(o.Participant) {
		return nil, false
	}
	return o.Participant, true
}

// HasParticipant returns a boolean if a field has been set.
func (o *Encounter) HasParticipant() bool {
	if o != nil && !IsNil(o.Participant) {
		return true
	}

	return false
}

// SetParticipant gets a reference to the given []EncounterParticipant and assigns it to the Participant field.
func (o *Encounter) SetParticipant(v []EncounterParticipant) {
	o.Participant = v
}

// GetAppointment returns the Appointment field value if set, zero value otherwise.
func (o *Encounter) GetAppointment() []Reference {
	if o == nil || IsNil(o.Appointment) {
		var ret []Reference
		return ret
	}
	return o.Appointment
}

// GetAppointmentOk returns a tuple with the Appointment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetAppointmentOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Appointment) {
		return nil, false
	}
	return o.Appointment, true
}

// HasAppointment returns a boolean if a field has been set.
func (o *Encounter) HasAppointment() bool {
	if o != nil && !IsNil(o.Appointment) {
		return true
	}

	return false
}

// SetAppointment gets a reference to the given []Reference and assigns it to the Appointment field.
func (o *Encounter) SetAppointment(v []Reference) {
	o.Appointment = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *Encounter) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *Encounter) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *Encounter) SetPeriod(v Period) {
	o.Period = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Encounter) GetLength() Duration {
	if o == nil || IsNil(o.Length) {
		var ret Duration
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetLengthOk() (*Duration, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Encounter) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given Duration and assigns it to the Length field.
func (o *Encounter) SetLength(v Duration) {
	o.Length = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *Encounter) GetReasonCode() []CodeableConcept {
	if o == nil || IsNil(o.ReasonCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetReasonCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *Encounter) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given []CodeableConcept and assigns it to the ReasonCode field.
func (o *Encounter) SetReasonCode(v []CodeableConcept) {
	o.ReasonCode = v
}

// GetReasonReference returns the ReasonReference field value if set, zero value otherwise.
func (o *Encounter) GetReasonReference() []Reference {
	if o == nil || IsNil(o.ReasonReference) {
		var ret []Reference
		return ret
	}
	return o.ReasonReference
}

// GetReasonReferenceOk returns a tuple with the ReasonReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetReasonReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ReasonReference) {
		return nil, false
	}
	return o.ReasonReference, true
}

// HasReasonReference returns a boolean if a field has been set.
func (o *Encounter) HasReasonReference() bool {
	if o != nil && !IsNil(o.ReasonReference) {
		return true
	}

	return false
}

// SetReasonReference gets a reference to the given []Reference and assigns it to the ReasonReference field.
func (o *Encounter) SetReasonReference(v []Reference) {
	o.ReasonReference = v
}

// GetDiagnosis returns the Diagnosis field value if set, zero value otherwise.
func (o *Encounter) GetDiagnosis() []EncounterDiagnosis {
	if o == nil || IsNil(o.Diagnosis) {
		var ret []EncounterDiagnosis
		return ret
	}
	return o.Diagnosis
}

// GetDiagnosisOk returns a tuple with the Diagnosis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetDiagnosisOk() ([]EncounterDiagnosis, bool) {
	if o == nil || IsNil(o.Diagnosis) {
		return nil, false
	}
	return o.Diagnosis, true
}

// HasDiagnosis returns a boolean if a field has been set.
func (o *Encounter) HasDiagnosis() bool {
	if o != nil && !IsNil(o.Diagnosis) {
		return true
	}

	return false
}

// SetDiagnosis gets a reference to the given []EncounterDiagnosis and assigns it to the Diagnosis field.
func (o *Encounter) SetDiagnosis(v []EncounterDiagnosis) {
	o.Diagnosis = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Encounter) GetAccount() []Reference {
	if o == nil || IsNil(o.Account) {
		var ret []Reference
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetAccountOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Encounter) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []Reference and assigns it to the Account field.
func (o *Encounter) SetAccount(v []Reference) {
	o.Account = v
}

// GetHospitalization returns the Hospitalization field value if set, zero value otherwise.
func (o *Encounter) GetHospitalization() EncounterHospitalization {
	if o == nil || IsNil(o.Hospitalization) {
		var ret EncounterHospitalization
		return ret
	}
	return *o.Hospitalization
}

// GetHospitalizationOk returns a tuple with the Hospitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetHospitalizationOk() (*EncounterHospitalization, bool) {
	if o == nil || IsNil(o.Hospitalization) {
		return nil, false
	}
	return o.Hospitalization, true
}

// HasHospitalization returns a boolean if a field has been set.
func (o *Encounter) HasHospitalization() bool {
	if o != nil && !IsNil(o.Hospitalization) {
		return true
	}

	return false
}

// SetHospitalization gets a reference to the given EncounterHospitalization and assigns it to the Hospitalization field.
func (o *Encounter) SetHospitalization(v EncounterHospitalization) {
	o.Hospitalization = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Encounter) GetLocation() []EncounterLocation {
	if o == nil || IsNil(o.Location) {
		var ret []EncounterLocation
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetLocationOk() ([]EncounterLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Encounter) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given []EncounterLocation and assigns it to the Location field.
func (o *Encounter) SetLocation(v []EncounterLocation) {
	o.Location = v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *Encounter) GetServiceProvider() Reference {
	if o == nil || IsNil(o.ServiceProvider) {
		var ret Reference
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetServiceProviderOk() (*Reference, bool) {
	if o == nil || IsNil(o.ServiceProvider) {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *Encounter) HasServiceProvider() bool {
	if o != nil && !IsNil(o.ServiceProvider) {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given Reference and assigns it to the ServiceProvider field.
func (o *Encounter) SetServiceProvider(v Reference) {
	o.ServiceProvider = &v
}

// GetPartOf returns the PartOf field value if set, zero value otherwise.
func (o *Encounter) GetPartOf() Reference {
	if o == nil || IsNil(o.PartOf) {
		var ret Reference
		return ret
	}
	return *o.PartOf
}

// GetPartOfOk returns a tuple with the PartOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encounter) GetPartOfOk() (*Reference, bool) {
	if o == nil || IsNil(o.PartOf) {
		return nil, false
	}
	return o.PartOf, true
}

// HasPartOf returns a boolean if a field has been set.
func (o *Encounter) HasPartOf() bool {
	if o != nil && !IsNil(o.PartOf) {
		return true
	}

	return false
}

// SetPartOf gets a reference to the given Reference and assigns it to the PartOf field.
func (o *Encounter) SetPartOf(v Reference) {
	o.PartOf = &v
}

func (o Encounter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Encounter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ImplicitRules) {
		toSerialize["implicitRules"] = o.ImplicitRules
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Contained) {
		toSerialize["contained"] = o.Contained
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.ModifierExtension) {
		toSerialize["modifierExtension"] = o.ModifierExtension
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusHistory) {
		toSerialize["statusHistory"] = o.StatusHistory
	}
	toSerialize["class"] = o.Class
	if !IsNil(o.ClassHistory) {
		toSerialize["classHistory"] = o.ClassHistory
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.EpisodeOfCare) {
		toSerialize["episodeOfCare"] = o.EpisodeOfCare
	}
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	if !IsNil(o.Participant) {
		toSerialize["participant"] = o.Participant
	}
	if !IsNil(o.Appointment) {
		toSerialize["appointment"] = o.Appointment
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if !IsNil(o.ReasonReference) {
		toSerialize["reasonReference"] = o.ReasonReference
	}
	if !IsNil(o.Diagnosis) {
		toSerialize["diagnosis"] = o.Diagnosis
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Hospitalization) {
		toSerialize["hospitalization"] = o.Hospitalization
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.ServiceProvider) {
		toSerialize["serviceProvider"] = o.ServiceProvider
	}
	if !IsNil(o.PartOf) {
		toSerialize["partOf"] = o.PartOf
	}
	return toSerialize, nil
}

func (o *Encounter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"class",
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

	varEncounter := _Encounter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEncounter)

	if err != nil {
		return err
	}

	*o = Encounter(varEncounter)

	return err
}

type NullableEncounter struct {
	value *Encounter
	isSet bool
}

func (v NullableEncounter) Get() *Encounter {
	return v.value
}

func (v *NullableEncounter) Set(val *Encounter) {
	v.value = val
	v.isSet = true
}

func (v NullableEncounter) IsSet() bool {
	return v.isSet
}

func (v *NullableEncounter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncounter(val *Encounter) *NullableEncounter {
	return &NullableEncounter{value: val, isSet: true}
}

func (v NullableEncounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncounter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


