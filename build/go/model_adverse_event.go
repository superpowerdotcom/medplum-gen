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

// checks if the AdverseEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdverseEvent{}

// AdverseEvent Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent struct {
	// This is a AdverseEvent resource
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
	// Business identifiers assigned to this adverse event by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier *Identifier `json:"identifier,omitempty"`
	// Whether the event actually happened, or just had the potential to. Note that this is independent of whether anyone was affected or harmed or how severely.
	Actuality *string `json:"actuality,omitempty"`
	// The overall type of event, intended for search and filtering purposes.
	Category []CodeableConcept `json:"category,omitempty"`
	// This element defines the specific type of event that occurred or that was prevented from occurring.
	Event *CodeableConcept `json:"event,omitempty"`
	// This subject or group impacted by the event.
	Subject Reference `json:"subject"`
	// The Encounter during which AdverseEvent was created or to which the creation of this record is tightly associated.
	Encounter *Reference `json:"encounter,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Date *string `json:"date,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Detected *string `json:"detected,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	RecordedDate *string `json:"recordedDate,omitempty"`
	// Includes information about the reaction that occurred as a result of exposure to a substance (for example, a drug or a chemical).
	ResultingCondition []Reference `json:"resultingCondition,omitempty"`
	// The information about where the adverse event occurred.
	Location *Reference `json:"location,omitempty"`
	// Assessment whether this event was of real importance.
	Seriousness *CodeableConcept `json:"seriousness,omitempty"`
	// Describes the severity of the adverse event, in relation to the subject. Contrast to AdverseEvent.seriousness - a severe rash might not be serious, but a mild heart problem is.
	Severity *CodeableConcept `json:"severity,omitempty"`
	// Describes the type of outcome from the adverse event.
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// Information on who recorded the adverse event.  May be the patient or a practitioner.
	Recorder *Reference `json:"recorder,omitempty"`
	// Parties that may or should contribute or have contributed information to the adverse event, which can consist of one or more activities.  Such information includes information leading to the decision to perform the activity and how to perform the activity (e.g. consultant), information that the activity itself seeks to reveal (e.g. informant of clinical history), or information about what activity was performed (e.g. informant witness).
	Contributor []Reference `json:"contributor,omitempty"`
	// Describes the entity that is suspected to have caused the adverse event.
	SuspectEntity []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
	// AdverseEvent.subjectMedicalHistory.
	SubjectMedicalHistory []Reference `json:"subjectMedicalHistory,omitempty"`
	// AdverseEvent.referenceDocument.
	ReferenceDocument []Reference `json:"referenceDocument,omitempty"`
	// AdverseEvent.study.
	Study []Reference `json:"study,omitempty"`
}

type _AdverseEvent AdverseEvent

// NewAdverseEvent instantiates a new AdverseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdverseEvent(resourceType string, subject Reference) *AdverseEvent {
	this := AdverseEvent{}
	this.ResourceType = resourceType
	this.Subject = subject
	return &this
}

// NewAdverseEventWithDefaults instantiates a new AdverseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdverseEventWithDefaults() *AdverseEvent {
	this := AdverseEvent{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *AdverseEvent) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *AdverseEvent) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdverseEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdverseEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdverseEvent) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AdverseEvent) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AdverseEvent) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *AdverseEvent) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *AdverseEvent) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *AdverseEvent) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *AdverseEvent) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AdverseEvent) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AdverseEvent) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AdverseEvent) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AdverseEvent) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AdverseEvent) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *AdverseEvent) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *AdverseEvent) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *AdverseEvent) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *AdverseEvent) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AdverseEvent) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AdverseEvent) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *AdverseEvent) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *AdverseEvent) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *AdverseEvent) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *AdverseEvent) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AdverseEvent) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AdverseEvent) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *AdverseEvent) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetActuality returns the Actuality field value if set, zero value otherwise.
func (o *AdverseEvent) GetActuality() string {
	if o == nil || IsNil(o.Actuality) {
		var ret string
		return ret
	}
	return *o.Actuality
}

// GetActualityOk returns a tuple with the Actuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetActualityOk() (*string, bool) {
	if o == nil || IsNil(o.Actuality) {
		return nil, false
	}
	return o.Actuality, true
}

// HasActuality returns a boolean if a field has been set.
func (o *AdverseEvent) HasActuality() bool {
	if o != nil && !IsNil(o.Actuality) {
		return true
	}

	return false
}

// SetActuality gets a reference to the given string and assigns it to the Actuality field.
func (o *AdverseEvent) SetActuality(v string) {
	o.Actuality = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AdverseEvent) GetCategory() []CodeableConcept {
	if o == nil || IsNil(o.Category) {
		var ret []CodeableConcept
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetCategoryOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AdverseEvent) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []CodeableConcept and assigns it to the Category field.
func (o *AdverseEvent) SetCategory(v []CodeableConcept) {
	o.Category = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AdverseEvent) GetEvent() CodeableConcept {
	if o == nil || IsNil(o.Event) {
		var ret CodeableConcept
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetEventOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AdverseEvent) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given CodeableConcept and assigns it to the Event field.
func (o *AdverseEvent) SetEvent(v CodeableConcept) {
	o.Event = &v
}

// GetSubject returns the Subject field value
func (o *AdverseEvent) GetSubject() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetSubjectOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *AdverseEvent) SetSubject(v Reference) {
	o.Subject = v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *AdverseEvent) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *AdverseEvent) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *AdverseEvent) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AdverseEvent) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AdverseEvent) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AdverseEvent) SetDate(v string) {
	o.Date = &v
}

// GetDetected returns the Detected field value if set, zero value otherwise.
func (o *AdverseEvent) GetDetected() string {
	if o == nil || IsNil(o.Detected) {
		var ret string
		return ret
	}
	return *o.Detected
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetDetectedOk() (*string, bool) {
	if o == nil || IsNil(o.Detected) {
		return nil, false
	}
	return o.Detected, true
}

// HasDetected returns a boolean if a field has been set.
func (o *AdverseEvent) HasDetected() bool {
	if o != nil && !IsNil(o.Detected) {
		return true
	}

	return false
}

// SetDetected gets a reference to the given string and assigns it to the Detected field.
func (o *AdverseEvent) SetDetected(v string) {
	o.Detected = &v
}

// GetRecordedDate returns the RecordedDate field value if set, zero value otherwise.
func (o *AdverseEvent) GetRecordedDate() string {
	if o == nil || IsNil(o.RecordedDate) {
		var ret string
		return ret
	}
	return *o.RecordedDate
}

// GetRecordedDateOk returns a tuple with the RecordedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetRecordedDateOk() (*string, bool) {
	if o == nil || IsNil(o.RecordedDate) {
		return nil, false
	}
	return o.RecordedDate, true
}

// HasRecordedDate returns a boolean if a field has been set.
func (o *AdverseEvent) HasRecordedDate() bool {
	if o != nil && !IsNil(o.RecordedDate) {
		return true
	}

	return false
}

// SetRecordedDate gets a reference to the given string and assigns it to the RecordedDate field.
func (o *AdverseEvent) SetRecordedDate(v string) {
	o.RecordedDate = &v
}

// GetResultingCondition returns the ResultingCondition field value if set, zero value otherwise.
func (o *AdverseEvent) GetResultingCondition() []Reference {
	if o == nil || IsNil(o.ResultingCondition) {
		var ret []Reference
		return ret
	}
	return o.ResultingCondition
}

// GetResultingConditionOk returns a tuple with the ResultingCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetResultingConditionOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ResultingCondition) {
		return nil, false
	}
	return o.ResultingCondition, true
}

// HasResultingCondition returns a boolean if a field has been set.
func (o *AdverseEvent) HasResultingCondition() bool {
	if o != nil && !IsNil(o.ResultingCondition) {
		return true
	}

	return false
}

// SetResultingCondition gets a reference to the given []Reference and assigns it to the ResultingCondition field.
func (o *AdverseEvent) SetResultingCondition(v []Reference) {
	o.ResultingCondition = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AdverseEvent) GetLocation() Reference {
	if o == nil || IsNil(o.Location) {
		var ret Reference
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetLocationOk() (*Reference, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AdverseEvent) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Reference and assigns it to the Location field.
func (o *AdverseEvent) SetLocation(v Reference) {
	o.Location = &v
}

// GetSeriousness returns the Seriousness field value if set, zero value otherwise.
func (o *AdverseEvent) GetSeriousness() CodeableConcept {
	if o == nil || IsNil(o.Seriousness) {
		var ret CodeableConcept
		return ret
	}
	return *o.Seriousness
}

// GetSeriousnessOk returns a tuple with the Seriousness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetSeriousnessOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Seriousness) {
		return nil, false
	}
	return o.Seriousness, true
}

// HasSeriousness returns a boolean if a field has been set.
func (o *AdverseEvent) HasSeriousness() bool {
	if o != nil && !IsNil(o.Seriousness) {
		return true
	}

	return false
}

// SetSeriousness gets a reference to the given CodeableConcept and assigns it to the Seriousness field.
func (o *AdverseEvent) SetSeriousness(v CodeableConcept) {
	o.Seriousness = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdverseEvent) GetSeverity() CodeableConcept {
	if o == nil || IsNil(o.Severity) {
		var ret CodeableConcept
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetSeverityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdverseEvent) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given CodeableConcept and assigns it to the Severity field.
func (o *AdverseEvent) SetSeverity(v CodeableConcept) {
	o.Severity = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *AdverseEvent) GetOutcome() CodeableConcept {
	if o == nil || IsNil(o.Outcome) {
		var ret CodeableConcept
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetOutcomeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *AdverseEvent) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given CodeableConcept and assigns it to the Outcome field.
func (o *AdverseEvent) SetOutcome(v CodeableConcept) {
	o.Outcome = &v
}

// GetRecorder returns the Recorder field value if set, zero value otherwise.
func (o *AdverseEvent) GetRecorder() Reference {
	if o == nil || IsNil(o.Recorder) {
		var ret Reference
		return ret
	}
	return *o.Recorder
}

// GetRecorderOk returns a tuple with the Recorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetRecorderOk() (*Reference, bool) {
	if o == nil || IsNil(o.Recorder) {
		return nil, false
	}
	return o.Recorder, true
}

// HasRecorder returns a boolean if a field has been set.
func (o *AdverseEvent) HasRecorder() bool {
	if o != nil && !IsNil(o.Recorder) {
		return true
	}

	return false
}

// SetRecorder gets a reference to the given Reference and assigns it to the Recorder field.
func (o *AdverseEvent) SetRecorder(v Reference) {
	o.Recorder = &v
}

// GetContributor returns the Contributor field value if set, zero value otherwise.
func (o *AdverseEvent) GetContributor() []Reference {
	if o == nil || IsNil(o.Contributor) {
		var ret []Reference
		return ret
	}
	return o.Contributor
}

// GetContributorOk returns a tuple with the Contributor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetContributorOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Contributor) {
		return nil, false
	}
	return o.Contributor, true
}

// HasContributor returns a boolean if a field has been set.
func (o *AdverseEvent) HasContributor() bool {
	if o != nil && !IsNil(o.Contributor) {
		return true
	}

	return false
}

// SetContributor gets a reference to the given []Reference and assigns it to the Contributor field.
func (o *AdverseEvent) SetContributor(v []Reference) {
	o.Contributor = v
}

// GetSuspectEntity returns the SuspectEntity field value if set, zero value otherwise.
func (o *AdverseEvent) GetSuspectEntity() []AdverseEventSuspectEntity {
	if o == nil || IsNil(o.SuspectEntity) {
		var ret []AdverseEventSuspectEntity
		return ret
	}
	return o.SuspectEntity
}

// GetSuspectEntityOk returns a tuple with the SuspectEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetSuspectEntityOk() ([]AdverseEventSuspectEntity, bool) {
	if o == nil || IsNil(o.SuspectEntity) {
		return nil, false
	}
	return o.SuspectEntity, true
}

// HasSuspectEntity returns a boolean if a field has been set.
func (o *AdverseEvent) HasSuspectEntity() bool {
	if o != nil && !IsNil(o.SuspectEntity) {
		return true
	}

	return false
}

// SetSuspectEntity gets a reference to the given []AdverseEventSuspectEntity and assigns it to the SuspectEntity field.
func (o *AdverseEvent) SetSuspectEntity(v []AdverseEventSuspectEntity) {
	o.SuspectEntity = v
}

// GetSubjectMedicalHistory returns the SubjectMedicalHistory field value if set, zero value otherwise.
func (o *AdverseEvent) GetSubjectMedicalHistory() []Reference {
	if o == nil || IsNil(o.SubjectMedicalHistory) {
		var ret []Reference
		return ret
	}
	return o.SubjectMedicalHistory
}

// GetSubjectMedicalHistoryOk returns a tuple with the SubjectMedicalHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetSubjectMedicalHistoryOk() ([]Reference, bool) {
	if o == nil || IsNil(o.SubjectMedicalHistory) {
		return nil, false
	}
	return o.SubjectMedicalHistory, true
}

// HasSubjectMedicalHistory returns a boolean if a field has been set.
func (o *AdverseEvent) HasSubjectMedicalHistory() bool {
	if o != nil && !IsNil(o.SubjectMedicalHistory) {
		return true
	}

	return false
}

// SetSubjectMedicalHistory gets a reference to the given []Reference and assigns it to the SubjectMedicalHistory field.
func (o *AdverseEvent) SetSubjectMedicalHistory(v []Reference) {
	o.SubjectMedicalHistory = v
}

// GetReferenceDocument returns the ReferenceDocument field value if set, zero value otherwise.
func (o *AdverseEvent) GetReferenceDocument() []Reference {
	if o == nil || IsNil(o.ReferenceDocument) {
		var ret []Reference
		return ret
	}
	return o.ReferenceDocument
}

// GetReferenceDocumentOk returns a tuple with the ReferenceDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetReferenceDocumentOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ReferenceDocument) {
		return nil, false
	}
	return o.ReferenceDocument, true
}

// HasReferenceDocument returns a boolean if a field has been set.
func (o *AdverseEvent) HasReferenceDocument() bool {
	if o != nil && !IsNil(o.ReferenceDocument) {
		return true
	}

	return false
}

// SetReferenceDocument gets a reference to the given []Reference and assigns it to the ReferenceDocument field.
func (o *AdverseEvent) SetReferenceDocument(v []Reference) {
	o.ReferenceDocument = v
}

// GetStudy returns the Study field value if set, zero value otherwise.
func (o *AdverseEvent) GetStudy() []Reference {
	if o == nil || IsNil(o.Study) {
		var ret []Reference
		return ret
	}
	return o.Study
}

// GetStudyOk returns a tuple with the Study field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEvent) GetStudyOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Study) {
		return nil, false
	}
	return o.Study, true
}

// HasStudy returns a boolean if a field has been set.
func (o *AdverseEvent) HasStudy() bool {
	if o != nil && !IsNil(o.Study) {
		return true
	}

	return false
}

// SetStudy gets a reference to the given []Reference and assigns it to the Study field.
func (o *AdverseEvent) SetStudy(v []Reference) {
	o.Study = v
}

func (o AdverseEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdverseEvent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Actuality) {
		toSerialize["actuality"] = o.Actuality
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Detected) {
		toSerialize["detected"] = o.Detected
	}
	if !IsNil(o.RecordedDate) {
		toSerialize["recordedDate"] = o.RecordedDate
	}
	if !IsNil(o.ResultingCondition) {
		toSerialize["resultingCondition"] = o.ResultingCondition
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Seriousness) {
		toSerialize["seriousness"] = o.Seriousness
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.Recorder) {
		toSerialize["recorder"] = o.Recorder
	}
	if !IsNil(o.Contributor) {
		toSerialize["contributor"] = o.Contributor
	}
	if !IsNil(o.SuspectEntity) {
		toSerialize["suspectEntity"] = o.SuspectEntity
	}
	if !IsNil(o.SubjectMedicalHistory) {
		toSerialize["subjectMedicalHistory"] = o.SubjectMedicalHistory
	}
	if !IsNil(o.ReferenceDocument) {
		toSerialize["referenceDocument"] = o.ReferenceDocument
	}
	if !IsNil(o.Study) {
		toSerialize["study"] = o.Study
	}
	return toSerialize, nil
}

func (o *AdverseEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"subject",
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

	varAdverseEvent := _AdverseEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdverseEvent)

	if err != nil {
		return err
	}

	*o = AdverseEvent(varAdverseEvent)

	return err
}

type NullableAdverseEvent struct {
	value *AdverseEvent
	isSet bool
}

func (v NullableAdverseEvent) Get() *AdverseEvent {
	return v.value
}

func (v *NullableAdverseEvent) Set(val *AdverseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAdverseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAdverseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdverseEvent(val *AdverseEvent) *NullableAdverseEvent {
	return &NullableAdverseEvent{value: val, isSet: true}
}

func (v NullableAdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdverseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

