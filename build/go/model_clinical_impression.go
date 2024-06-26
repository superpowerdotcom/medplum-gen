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

// checks if the ClinicalImpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClinicalImpression{}

// ClinicalImpression A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called \"ClinicalImpression\" rather than \"ClinicalAssessment\" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression struct {
	// This is a ClinicalImpression resource
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
	// Business identifiers assigned to this clinical impression by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Status *string `json:"status,omitempty"`
	// Captures the reason for the current state of the ClinicalImpression.
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// Categorizes the type of clinical assessment performed.
	Code *CodeableConcept `json:"code,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// The patient or group of individuals assessed as part of this record.
	Subject Reference `json:"subject"`
	// The Encounter during which this ClinicalImpression was created or to which the creation of this record is tightly associated.
	Encounter *Reference `json:"encounter,omitempty"`
	// The point in time or period over which the subject was assessed.
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	// The point in time or period over which the subject was assessed.
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Date *string `json:"date,omitempty"`
	// The clinician performing the assessment.
	Assessor *Reference `json:"assessor,omitempty"`
	// A reference to the last assessment that was conducted on this patient. Assessments are often/usually ongoing in nature; a care provider (practitioner or team) will make new assessments on an ongoing basis as new data arises or the patient's conditions changes.
	Previous *Reference `json:"previous,omitempty"`
	// A list of the relevant problems/conditions for a patient.
	Problem []Reference `json:"problem,omitempty"`
	// One or more sets of investigations (signs, symptoms, etc.). The actual grouping of investigations varies greatly depending on the type and context of the assessment. These investigations may include data generated during the assessment process, or data previously generated and recorded that is pertinent to the outcomes.
	Investigation []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	// Reference to a specific published clinical protocol that was followed during this assessment, and/or that provides evidence in support of the diagnosis.
	Protocol []string `json:"protocol,omitempty"`
	// A sequence of Unicode characters
	Summary *string `json:"summary,omitempty"`
	// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment.
	Finding []ClinicalImpressionFinding `json:"finding,omitempty"`
	// Estimate of likely outcome.
	PrognosisCodeableConcept []CodeableConcept `json:"prognosisCodeableConcept,omitempty"`
	// RiskAssessment expressing likely outcome.
	PrognosisReference []Reference `json:"prognosisReference,omitempty"`
	// Information supporting the clinical impression.
	SupportingInfo []Reference `json:"supportingInfo,omitempty"`
	// Commentary about the impression, typically recorded after the impression itself was made, though supplemental notes by the original author could also appear.
	Note []Annotation `json:"note,omitempty"`
}

type _ClinicalImpression ClinicalImpression

// NewClinicalImpression instantiates a new ClinicalImpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClinicalImpression(resourceType string, subject Reference) *ClinicalImpression {
	this := ClinicalImpression{}
	this.ResourceType = resourceType
	this.Subject = subject
	return &this
}

// NewClinicalImpressionWithDefaults instantiates a new ClinicalImpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClinicalImpressionWithDefaults() *ClinicalImpression {
	this := ClinicalImpression{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *ClinicalImpression) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ClinicalImpression) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClinicalImpression) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClinicalImpression) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClinicalImpression) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClinicalImpression) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClinicalImpression) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *ClinicalImpression) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *ClinicalImpression) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *ClinicalImpression) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *ClinicalImpression) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ClinicalImpression) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ClinicalImpression) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ClinicalImpression) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ClinicalImpression) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ClinicalImpression) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *ClinicalImpression) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *ClinicalImpression) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *ClinicalImpression) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *ClinicalImpression) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ClinicalImpression) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ClinicalImpression) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ClinicalImpression) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ClinicalImpression) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ClinicalImpression) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ClinicalImpression) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ClinicalImpression) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ClinicalImpression) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *ClinicalImpression) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClinicalImpression) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClinicalImpression) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClinicalImpression) SetStatus(v string) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *ClinicalImpression) GetStatusReason() CodeableConcept {
	if o == nil || IsNil(o.StatusReason) {
		var ret CodeableConcept
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetStatusReasonOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *ClinicalImpression) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given CodeableConcept and assigns it to the StatusReason field.
func (o *ClinicalImpression) SetStatusReason(v CodeableConcept) {
	o.StatusReason = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ClinicalImpression) GetCode() CodeableConcept {
	if o == nil || IsNil(o.Code) {
		var ret CodeableConcept
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ClinicalImpression) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given CodeableConcept and assigns it to the Code field.
func (o *ClinicalImpression) SetCode(v CodeableConcept) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClinicalImpression) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClinicalImpression) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClinicalImpression) SetDescription(v string) {
	o.Description = &v
}

// GetSubject returns the Subject field value
func (o *ClinicalImpression) GetSubject() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetSubjectOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *ClinicalImpression) SetSubject(v Reference) {
	o.Subject = v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *ClinicalImpression) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *ClinicalImpression) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *ClinicalImpression) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetEffectiveDateTime returns the EffectiveDateTime field value if set, zero value otherwise.
func (o *ClinicalImpression) GetEffectiveDateTime() string {
	if o == nil || IsNil(o.EffectiveDateTime) {
		var ret string
		return ret
	}
	return *o.EffectiveDateTime
}

// GetEffectiveDateTimeOk returns a tuple with the EffectiveDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetEffectiveDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDateTime) {
		return nil, false
	}
	return o.EffectiveDateTime, true
}

// HasEffectiveDateTime returns a boolean if a field has been set.
func (o *ClinicalImpression) HasEffectiveDateTime() bool {
	if o != nil && !IsNil(o.EffectiveDateTime) {
		return true
	}

	return false
}

// SetEffectiveDateTime gets a reference to the given string and assigns it to the EffectiveDateTime field.
func (o *ClinicalImpression) SetEffectiveDateTime(v string) {
	o.EffectiveDateTime = &v
}

// GetEffectivePeriod returns the EffectivePeriod field value if set, zero value otherwise.
func (o *ClinicalImpression) GetEffectivePeriod() Period {
	if o == nil || IsNil(o.EffectivePeriod) {
		var ret Period
		return ret
	}
	return *o.EffectivePeriod
}

// GetEffectivePeriodOk returns a tuple with the EffectivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetEffectivePeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.EffectivePeriod) {
		return nil, false
	}
	return o.EffectivePeriod, true
}

// HasEffectivePeriod returns a boolean if a field has been set.
func (o *ClinicalImpression) HasEffectivePeriod() bool {
	if o != nil && !IsNil(o.EffectivePeriod) {
		return true
	}

	return false
}

// SetEffectivePeriod gets a reference to the given Period and assigns it to the EffectivePeriod field.
func (o *ClinicalImpression) SetEffectivePeriod(v Period) {
	o.EffectivePeriod = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ClinicalImpression) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ClinicalImpression) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ClinicalImpression) SetDate(v string) {
	o.Date = &v
}

// GetAssessor returns the Assessor field value if set, zero value otherwise.
func (o *ClinicalImpression) GetAssessor() Reference {
	if o == nil || IsNil(o.Assessor) {
		var ret Reference
		return ret
	}
	return *o.Assessor
}

// GetAssessorOk returns a tuple with the Assessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetAssessorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Assessor) {
		return nil, false
	}
	return o.Assessor, true
}

// HasAssessor returns a boolean if a field has been set.
func (o *ClinicalImpression) HasAssessor() bool {
	if o != nil && !IsNil(o.Assessor) {
		return true
	}

	return false
}

// SetAssessor gets a reference to the given Reference and assigns it to the Assessor field.
func (o *ClinicalImpression) SetAssessor(v Reference) {
	o.Assessor = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ClinicalImpression) GetPrevious() Reference {
	if o == nil || IsNil(o.Previous) {
		var ret Reference
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetPreviousOk() (*Reference, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ClinicalImpression) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given Reference and assigns it to the Previous field.
func (o *ClinicalImpression) SetPrevious(v Reference) {
	o.Previous = &v
}

// GetProblem returns the Problem field value if set, zero value otherwise.
func (o *ClinicalImpression) GetProblem() []Reference {
	if o == nil || IsNil(o.Problem) {
		var ret []Reference
		return ret
	}
	return o.Problem
}

// GetProblemOk returns a tuple with the Problem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetProblemOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Problem) {
		return nil, false
	}
	return o.Problem, true
}

// HasProblem returns a boolean if a field has been set.
func (o *ClinicalImpression) HasProblem() bool {
	if o != nil && !IsNil(o.Problem) {
		return true
	}

	return false
}

// SetProblem gets a reference to the given []Reference and assigns it to the Problem field.
func (o *ClinicalImpression) SetProblem(v []Reference) {
	o.Problem = v
}

// GetInvestigation returns the Investigation field value if set, zero value otherwise.
func (o *ClinicalImpression) GetInvestigation() []ClinicalImpressionInvestigation {
	if o == nil || IsNil(o.Investigation) {
		var ret []ClinicalImpressionInvestigation
		return ret
	}
	return o.Investigation
}

// GetInvestigationOk returns a tuple with the Investigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetInvestigationOk() ([]ClinicalImpressionInvestigation, bool) {
	if o == nil || IsNil(o.Investigation) {
		return nil, false
	}
	return o.Investigation, true
}

// HasInvestigation returns a boolean if a field has been set.
func (o *ClinicalImpression) HasInvestigation() bool {
	if o != nil && !IsNil(o.Investigation) {
		return true
	}

	return false
}

// SetInvestigation gets a reference to the given []ClinicalImpressionInvestigation and assigns it to the Investigation field.
func (o *ClinicalImpression) SetInvestigation(v []ClinicalImpressionInvestigation) {
	o.Investigation = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ClinicalImpression) GetProtocol() []string {
	if o == nil || IsNil(o.Protocol) {
		var ret []string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetProtocolOk() ([]string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ClinicalImpression) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given []string and assigns it to the Protocol field.
func (o *ClinicalImpression) SetProtocol(v []string) {
	o.Protocol = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClinicalImpression) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClinicalImpression) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ClinicalImpression) SetSummary(v string) {
	o.Summary = &v
}

// GetFinding returns the Finding field value if set, zero value otherwise.
func (o *ClinicalImpression) GetFinding() []ClinicalImpressionFinding {
	if o == nil || IsNil(o.Finding) {
		var ret []ClinicalImpressionFinding
		return ret
	}
	return o.Finding
}

// GetFindingOk returns a tuple with the Finding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetFindingOk() ([]ClinicalImpressionFinding, bool) {
	if o == nil || IsNil(o.Finding) {
		return nil, false
	}
	return o.Finding, true
}

// HasFinding returns a boolean if a field has been set.
func (o *ClinicalImpression) HasFinding() bool {
	if o != nil && !IsNil(o.Finding) {
		return true
	}

	return false
}

// SetFinding gets a reference to the given []ClinicalImpressionFinding and assigns it to the Finding field.
func (o *ClinicalImpression) SetFinding(v []ClinicalImpressionFinding) {
	o.Finding = v
}

// GetPrognosisCodeableConcept returns the PrognosisCodeableConcept field value if set, zero value otherwise.
func (o *ClinicalImpression) GetPrognosisCodeableConcept() []CodeableConcept {
	if o == nil || IsNil(o.PrognosisCodeableConcept) {
		var ret []CodeableConcept
		return ret
	}
	return o.PrognosisCodeableConcept
}

// GetPrognosisCodeableConceptOk returns a tuple with the PrognosisCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetPrognosisCodeableConceptOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.PrognosisCodeableConcept) {
		return nil, false
	}
	return o.PrognosisCodeableConcept, true
}

// HasPrognosisCodeableConcept returns a boolean if a field has been set.
func (o *ClinicalImpression) HasPrognosisCodeableConcept() bool {
	if o != nil && !IsNil(o.PrognosisCodeableConcept) {
		return true
	}

	return false
}

// SetPrognosisCodeableConcept gets a reference to the given []CodeableConcept and assigns it to the PrognosisCodeableConcept field.
func (o *ClinicalImpression) SetPrognosisCodeableConcept(v []CodeableConcept) {
	o.PrognosisCodeableConcept = v
}

// GetPrognosisReference returns the PrognosisReference field value if set, zero value otherwise.
func (o *ClinicalImpression) GetPrognosisReference() []Reference {
	if o == nil || IsNil(o.PrognosisReference) {
		var ret []Reference
		return ret
	}
	return o.PrognosisReference
}

// GetPrognosisReferenceOk returns a tuple with the PrognosisReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetPrognosisReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.PrognosisReference) {
		return nil, false
	}
	return o.PrognosisReference, true
}

// HasPrognosisReference returns a boolean if a field has been set.
func (o *ClinicalImpression) HasPrognosisReference() bool {
	if o != nil && !IsNil(o.PrognosisReference) {
		return true
	}

	return false
}

// SetPrognosisReference gets a reference to the given []Reference and assigns it to the PrognosisReference field.
func (o *ClinicalImpression) SetPrognosisReference(v []Reference) {
	o.PrognosisReference = v
}

// GetSupportingInfo returns the SupportingInfo field value if set, zero value otherwise.
func (o *ClinicalImpression) GetSupportingInfo() []Reference {
	if o == nil || IsNil(o.SupportingInfo) {
		var ret []Reference
		return ret
	}
	return o.SupportingInfo
}

// GetSupportingInfoOk returns a tuple with the SupportingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetSupportingInfoOk() ([]Reference, bool) {
	if o == nil || IsNil(o.SupportingInfo) {
		return nil, false
	}
	return o.SupportingInfo, true
}

// HasSupportingInfo returns a boolean if a field has been set.
func (o *ClinicalImpression) HasSupportingInfo() bool {
	if o != nil && !IsNil(o.SupportingInfo) {
		return true
	}

	return false
}

// SetSupportingInfo gets a reference to the given []Reference and assigns it to the SupportingInfo field.
func (o *ClinicalImpression) SetSupportingInfo(v []Reference) {
	o.SupportingInfo = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ClinicalImpression) GetNote() []Annotation {
	if o == nil || IsNil(o.Note) {
		var ret []Annotation
		return ret
	}
	return o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpression) GetNoteOk() ([]Annotation, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ClinicalImpression) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given []Annotation and assigns it to the Note field.
func (o *ClinicalImpression) SetNote(v []Annotation) {
	o.Note = v
}

func (o ClinicalImpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClinicalImpression) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.EffectiveDateTime) {
		toSerialize["effectiveDateTime"] = o.EffectiveDateTime
	}
	if !IsNil(o.EffectivePeriod) {
		toSerialize["effectivePeriod"] = o.EffectivePeriod
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Assessor) {
		toSerialize["assessor"] = o.Assessor
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Problem) {
		toSerialize["problem"] = o.Problem
	}
	if !IsNil(o.Investigation) {
		toSerialize["investigation"] = o.Investigation
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Finding) {
		toSerialize["finding"] = o.Finding
	}
	if !IsNil(o.PrognosisCodeableConcept) {
		toSerialize["prognosisCodeableConcept"] = o.PrognosisCodeableConcept
	}
	if !IsNil(o.PrognosisReference) {
		toSerialize["prognosisReference"] = o.PrognosisReference
	}
	if !IsNil(o.SupportingInfo) {
		toSerialize["supportingInfo"] = o.SupportingInfo
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

func (o *ClinicalImpression) UnmarshalJSON(data []byte) (err error) {
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

	varClinicalImpression := _ClinicalImpression{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClinicalImpression)

	if err != nil {
		return err
	}

	*o = ClinicalImpression(varClinicalImpression)

	return err
}

type NullableClinicalImpression struct {
	value *ClinicalImpression
	isSet bool
}

func (v NullableClinicalImpression) Get() *ClinicalImpression {
	return v.value
}

func (v *NullableClinicalImpression) Set(val *ClinicalImpression) {
	v.value = val
	v.isSet = true
}

func (v NullableClinicalImpression) IsSet() bool {
	return v.isSet
}

func (v *NullableClinicalImpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClinicalImpression(val *ClinicalImpression) *NullableClinicalImpression {
	return &NullableClinicalImpression{value: val, isSet: true}
}

func (v NullableClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClinicalImpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


