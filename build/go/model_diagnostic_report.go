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

// checks if the DiagnosticReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticReport{}

// DiagnosticReport The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport struct {
	// This is a DiagnosticReport resource
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
	// Identifiers assigned to this report by the performer or other systems.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Details concerning a service requested.
	BasedOn []Reference `json:"basedOn,omitempty"`
	// The status of the diagnostic report.
	Status *string `json:"status,omitempty"`
	// A code that classifies the clinical discipline, department or diagnostic service that created the report (e.g. cardiology, biochemistry, hematology, MRI). This is used for searching, sorting and display purposes.
	Category []CodeableConcept `json:"category,omitempty"`
	// A code or name that describes this diagnostic report.
	Code CodeableConcept `json:"code"`
	// The subject of the report. Usually, but not always, this is a patient. However, diagnostic services also perform analyses on specimens collected from a variety of other sources.
	Subject *Reference `json:"subject,omitempty"`
	// The healthcare event  (e.g. a patient and healthcare provider interaction) which this DiagnosticReport is about.
	Encounter *Reference `json:"encounter,omitempty"`
	// The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	// The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// An instant in time - known at least to the second
	Issued *string `json:"issued,omitempty"`
	// The diagnostic service that is responsible for issuing the report.
	Performer []Reference `json:"performer,omitempty"`
	// The practitioner or organization that is responsible for the report's conclusions and interpretations.
	ResultsInterpreter []Reference `json:"resultsInterpreter,omitempty"`
	// Details about the specimens on which this diagnostic report is based.
	Specimen []Reference `json:"specimen,omitempty"`
	// [Observations](observation.html)  that are part of this diagnostic report.
	Result []Reference `json:"result,omitempty"`
	// One or more links to full details of any imaging performed during the diagnostic investigation. Typically, this is imaging performed by DICOM enabled modalities, but this is not required. A fully enabled PACS viewer can use this information to provide views of the source images.
	ImagingStudy []Reference `json:"imagingStudy,omitempty"`
	// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
	Media []DiagnosticReportMedia `json:"media,omitempty"`
	// A sequence of Unicode characters
	Conclusion *string `json:"conclusion,omitempty"`
	// One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report.
	ConclusionCode []CodeableConcept `json:"conclusionCode,omitempty"`
	// Rich text representation of the entire result as issued by the diagnostic service. Multiple formats are allowed but they SHALL be semantically equivalent.
	PresentedForm []Attachment `json:"presentedForm,omitempty"`
}

type _DiagnosticReport DiagnosticReport

// NewDiagnosticReport instantiates a new DiagnosticReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticReport(resourceType string, code CodeableConcept) *DiagnosticReport {
	this := DiagnosticReport{}
	this.ResourceType = resourceType
	this.Code = code
	return &this
}

// NewDiagnosticReportWithDefaults instantiates a new DiagnosticReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticReportWithDefaults() *DiagnosticReport {
	this := DiagnosticReport{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *DiagnosticReport) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *DiagnosticReport) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiagnosticReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiagnosticReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiagnosticReport) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DiagnosticReport) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DiagnosticReport) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *DiagnosticReport) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *DiagnosticReport) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *DiagnosticReport) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *DiagnosticReport) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *DiagnosticReport) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *DiagnosticReport) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *DiagnosticReport) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *DiagnosticReport) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *DiagnosticReport) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *DiagnosticReport) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *DiagnosticReport) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *DiagnosticReport) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *DiagnosticReport) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DiagnosticReport) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DiagnosticReport) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DiagnosticReport) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DiagnosticReport) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DiagnosticReport) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DiagnosticReport) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *DiagnosticReport) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *DiagnosticReport) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *DiagnosticReport) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *DiagnosticReport) GetBasedOn() []Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret []Reference
		return ret
	}
	return o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetBasedOnOk() ([]Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *DiagnosticReport) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given []Reference and assigns it to the BasedOn field.
func (o *DiagnosticReport) SetBasedOn(v []Reference) {
	o.BasedOn = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiagnosticReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiagnosticReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiagnosticReport) SetStatus(v string) {
	o.Status = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *DiagnosticReport) GetCategory() []CodeableConcept {
	if o == nil || IsNil(o.Category) {
		var ret []CodeableConcept
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetCategoryOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *DiagnosticReport) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []CodeableConcept and assigns it to the Category field.
func (o *DiagnosticReport) SetCategory(v []CodeableConcept) {
	o.Category = v
}

// GetCode returns the Code field value
func (o *DiagnosticReport) GetCode() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DiagnosticReport) SetCode(v CodeableConcept) {
	o.Code = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *DiagnosticReport) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *DiagnosticReport) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *DiagnosticReport) SetSubject(v Reference) {
	o.Subject = &v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *DiagnosticReport) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *DiagnosticReport) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *DiagnosticReport) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetEffectiveDateTime returns the EffectiveDateTime field value if set, zero value otherwise.
func (o *DiagnosticReport) GetEffectiveDateTime() string {
	if o == nil || IsNil(o.EffectiveDateTime) {
		var ret string
		return ret
	}
	return *o.EffectiveDateTime
}

// GetEffectiveDateTimeOk returns a tuple with the EffectiveDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetEffectiveDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDateTime) {
		return nil, false
	}
	return o.EffectiveDateTime, true
}

// HasEffectiveDateTime returns a boolean if a field has been set.
func (o *DiagnosticReport) HasEffectiveDateTime() bool {
	if o != nil && !IsNil(o.EffectiveDateTime) {
		return true
	}

	return false
}

// SetEffectiveDateTime gets a reference to the given string and assigns it to the EffectiveDateTime field.
func (o *DiagnosticReport) SetEffectiveDateTime(v string) {
	o.EffectiveDateTime = &v
}

// GetEffectivePeriod returns the EffectivePeriod field value if set, zero value otherwise.
func (o *DiagnosticReport) GetEffectivePeriod() Period {
	if o == nil || IsNil(o.EffectivePeriod) {
		var ret Period
		return ret
	}
	return *o.EffectivePeriod
}

// GetEffectivePeriodOk returns a tuple with the EffectivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetEffectivePeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.EffectivePeriod) {
		return nil, false
	}
	return o.EffectivePeriod, true
}

// HasEffectivePeriod returns a boolean if a field has been set.
func (o *DiagnosticReport) HasEffectivePeriod() bool {
	if o != nil && !IsNil(o.EffectivePeriod) {
		return true
	}

	return false
}

// SetEffectivePeriod gets a reference to the given Period and assigns it to the EffectivePeriod field.
func (o *DiagnosticReport) SetEffectivePeriod(v Period) {
	o.EffectivePeriod = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *DiagnosticReport) GetIssued() string {
	if o == nil || IsNil(o.Issued) {
		var ret string
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetIssuedOk() (*string, bool) {
	if o == nil || IsNil(o.Issued) {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *DiagnosticReport) HasIssued() bool {
	if o != nil && !IsNil(o.Issued) {
		return true
	}

	return false
}

// SetIssued gets a reference to the given string and assigns it to the Issued field.
func (o *DiagnosticReport) SetIssued(v string) {
	o.Issued = &v
}

// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *DiagnosticReport) GetPerformer() []Reference {
	if o == nil || IsNil(o.Performer) {
		var ret []Reference
		return ret
	}
	return o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetPerformerOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *DiagnosticReport) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given []Reference and assigns it to the Performer field.
func (o *DiagnosticReport) SetPerformer(v []Reference) {
	o.Performer = v
}

// GetResultsInterpreter returns the ResultsInterpreter field value if set, zero value otherwise.
func (o *DiagnosticReport) GetResultsInterpreter() []Reference {
	if o == nil || IsNil(o.ResultsInterpreter) {
		var ret []Reference
		return ret
	}
	return o.ResultsInterpreter
}

// GetResultsInterpreterOk returns a tuple with the ResultsInterpreter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetResultsInterpreterOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ResultsInterpreter) {
		return nil, false
	}
	return o.ResultsInterpreter, true
}

// HasResultsInterpreter returns a boolean if a field has been set.
func (o *DiagnosticReport) HasResultsInterpreter() bool {
	if o != nil && !IsNil(o.ResultsInterpreter) {
		return true
	}

	return false
}

// SetResultsInterpreter gets a reference to the given []Reference and assigns it to the ResultsInterpreter field.
func (o *DiagnosticReport) SetResultsInterpreter(v []Reference) {
	o.ResultsInterpreter = v
}

// GetSpecimen returns the Specimen field value if set, zero value otherwise.
func (o *DiagnosticReport) GetSpecimen() []Reference {
	if o == nil || IsNil(o.Specimen) {
		var ret []Reference
		return ret
	}
	return o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetSpecimenOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Specimen) {
		return nil, false
	}
	return o.Specimen, true
}

// HasSpecimen returns a boolean if a field has been set.
func (o *DiagnosticReport) HasSpecimen() bool {
	if o != nil && !IsNil(o.Specimen) {
		return true
	}

	return false
}

// SetSpecimen gets a reference to the given []Reference and assigns it to the Specimen field.
func (o *DiagnosticReport) SetSpecimen(v []Reference) {
	o.Specimen = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DiagnosticReport) GetResult() []Reference {
	if o == nil || IsNil(o.Result) {
		var ret []Reference
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetResultOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DiagnosticReport) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Reference and assigns it to the Result field.
func (o *DiagnosticReport) SetResult(v []Reference) {
	o.Result = v
}

// GetImagingStudy returns the ImagingStudy field value if set, zero value otherwise.
func (o *DiagnosticReport) GetImagingStudy() []Reference {
	if o == nil || IsNil(o.ImagingStudy) {
		var ret []Reference
		return ret
	}
	return o.ImagingStudy
}

// GetImagingStudyOk returns a tuple with the ImagingStudy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetImagingStudyOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ImagingStudy) {
		return nil, false
	}
	return o.ImagingStudy, true
}

// HasImagingStudy returns a boolean if a field has been set.
func (o *DiagnosticReport) HasImagingStudy() bool {
	if o != nil && !IsNil(o.ImagingStudy) {
		return true
	}

	return false
}

// SetImagingStudy gets a reference to the given []Reference and assigns it to the ImagingStudy field.
func (o *DiagnosticReport) SetImagingStudy(v []Reference) {
	o.ImagingStudy = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *DiagnosticReport) GetMedia() []DiagnosticReportMedia {
	if o == nil || IsNil(o.Media) {
		var ret []DiagnosticReportMedia
		return ret
	}
	return o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetMediaOk() ([]DiagnosticReportMedia, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *DiagnosticReport) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given []DiagnosticReportMedia and assigns it to the Media field.
func (o *DiagnosticReport) SetMedia(v []DiagnosticReportMedia) {
	o.Media = v
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise.
func (o *DiagnosticReport) GetConclusion() string {
	if o == nil || IsNil(o.Conclusion) {
		var ret string
		return ret
	}
	return *o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetConclusionOk() (*string, bool) {
	if o == nil || IsNil(o.Conclusion) {
		return nil, false
	}
	return o.Conclusion, true
}

// HasConclusion returns a boolean if a field has been set.
func (o *DiagnosticReport) HasConclusion() bool {
	if o != nil && !IsNil(o.Conclusion) {
		return true
	}

	return false
}

// SetConclusion gets a reference to the given string and assigns it to the Conclusion field.
func (o *DiagnosticReport) SetConclusion(v string) {
	o.Conclusion = &v
}

// GetConclusionCode returns the ConclusionCode field value if set, zero value otherwise.
func (o *DiagnosticReport) GetConclusionCode() []CodeableConcept {
	if o == nil || IsNil(o.ConclusionCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ConclusionCode
}

// GetConclusionCodeOk returns a tuple with the ConclusionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetConclusionCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ConclusionCode) {
		return nil, false
	}
	return o.ConclusionCode, true
}

// HasConclusionCode returns a boolean if a field has been set.
func (o *DiagnosticReport) HasConclusionCode() bool {
	if o != nil && !IsNil(o.ConclusionCode) {
		return true
	}

	return false
}

// SetConclusionCode gets a reference to the given []CodeableConcept and assigns it to the ConclusionCode field.
func (o *DiagnosticReport) SetConclusionCode(v []CodeableConcept) {
	o.ConclusionCode = v
}

// GetPresentedForm returns the PresentedForm field value if set, zero value otherwise.
func (o *DiagnosticReport) GetPresentedForm() []Attachment {
	if o == nil || IsNil(o.PresentedForm) {
		var ret []Attachment
		return ret
	}
	return o.PresentedForm
}

// GetPresentedFormOk returns a tuple with the PresentedForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetPresentedFormOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.PresentedForm) {
		return nil, false
	}
	return o.PresentedForm, true
}

// HasPresentedForm returns a boolean if a field has been set.
func (o *DiagnosticReport) HasPresentedForm() bool {
	if o != nil && !IsNil(o.PresentedForm) {
		return true
	}

	return false
}

// SetPresentedForm gets a reference to the given []Attachment and assigns it to the PresentedForm field.
func (o *DiagnosticReport) SetPresentedForm(v []Attachment) {
	o.PresentedForm = v
}

func (o DiagnosticReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticReport) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.EffectiveDateTime) {
		toSerialize["effectiveDateTime"] = o.EffectiveDateTime
	}
	if !IsNil(o.EffectivePeriod) {
		toSerialize["effectivePeriod"] = o.EffectivePeriod
	}
	if !IsNil(o.Issued) {
		toSerialize["issued"] = o.Issued
	}
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
	}
	if !IsNil(o.ResultsInterpreter) {
		toSerialize["resultsInterpreter"] = o.ResultsInterpreter
	}
	if !IsNil(o.Specimen) {
		toSerialize["specimen"] = o.Specimen
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.ImagingStudy) {
		toSerialize["imagingStudy"] = o.ImagingStudy
	}
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !IsNil(o.Conclusion) {
		toSerialize["conclusion"] = o.Conclusion
	}
	if !IsNil(o.ConclusionCode) {
		toSerialize["conclusionCode"] = o.ConclusionCode
	}
	if !IsNil(o.PresentedForm) {
		toSerialize["presentedForm"] = o.PresentedForm
	}
	return toSerialize, nil
}

func (o *DiagnosticReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"code",
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

	varDiagnosticReport := _DiagnosticReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiagnosticReport)

	if err != nil {
		return err
	}

	*o = DiagnosticReport(varDiagnosticReport)

	return err
}

type NullableDiagnosticReport struct {
	value *DiagnosticReport
	isSet bool
}

func (v NullableDiagnosticReport) Get() *DiagnosticReport {
	return v.value
}

func (v *NullableDiagnosticReport) Set(val *DiagnosticReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticReport(val *DiagnosticReport) *NullableDiagnosticReport {
	return &NullableDiagnosticReport{value: val, isSet: true}
}

func (v NullableDiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


