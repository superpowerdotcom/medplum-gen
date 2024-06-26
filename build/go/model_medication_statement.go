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

// checks if the MedicationStatement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationStatement{}

// MedicationStatement A record of a medication that is being consumed by a patient.   A MedicationStatement may indicate that the patient may be taking the medication now or has taken the medication in the past or will be taking the medication in the future.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay.   The medication information may come from sources such as the patient's memory, from a prescription bottle,  or from a list of medications the patient, clinician or other party maintains.   The primary difference between a medication statement and a medication administration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication.  A medication statement is often, if not always, less specific.  There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise.  As stated earlier, the medication statement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party maintains.  Medication administration is more formal and is not missing detailed information.
type MedicationStatement struct {
	// This is a MedicationStatement resource
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
	// Identifiers associated with this Medication Statement that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []Reference `json:"basedOn,omitempty"`
	// A larger event of which this particular event is a component or step.
	PartOf []Reference `json:"partOf,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Status *string `json:"status,omitempty"`
	// Captures the reason for the current state of the MedicationStatement.
	StatusReason []CodeableConcept `json:"statusReason,omitempty"`
	// Indicates where the medication is expected to be consumed or administered.
	Category *CodeableConcept `json:"category,omitempty"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationReference *Reference `json:"medicationReference,omitempty"`
	// The person, animal or group who is/was taking the medication.
	Subject Reference `json:"subject"`
	// The encounter or episode of care that establishes the context for this MedicationStatement.
	Context *Reference `json:"context,omitempty"`
	// The interval of time during which it is being asserted that the patient is/was/will be taking the medication (or was not taking, when the MedicationStatement.taken element is No).
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	// The interval of time during which it is being asserted that the patient is/was/will be taking the medication (or was not taking, when the MedicationStatement.taken element is No).
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	DateAsserted *string `json:"dateAsserted,omitempty"`
	// The person or organization that provided the information about the taking of this medication. Note: Use derivedFrom when a MedicationStatement is derived from other resources, e.g. Claim or MedicationRequest.
	InformationSource *Reference `json:"informationSource,omitempty"`
	// Allows linking the MedicationStatement to the underlying MedicationRequest, or to other information that supports or is used to derive the MedicationStatement.
	DerivedFrom []Reference `json:"derivedFrom,omitempty"`
	// A reason for why the medication is being/was taken.
	ReasonCode []CodeableConcept `json:"reasonCode,omitempty"`
	// Condition or observation that supports why the medication is being/was taken.
	ReasonReference []Reference `json:"reasonReference,omitempty"`
	// Provides extra information about the medication statement that is not conveyed by the other attributes.
	Note []Annotation `json:"note,omitempty"`
	// Indicates how the medication is/was or should be taken by the patient.
	Dosage []Dosage `json:"dosage,omitempty"`
}

type _MedicationStatement MedicationStatement

// NewMedicationStatement instantiates a new MedicationStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationStatement(resourceType string, subject Reference) *MedicationStatement {
	this := MedicationStatement{}
	this.ResourceType = resourceType
	this.Subject = subject
	return &this
}

// NewMedicationStatementWithDefaults instantiates a new MedicationStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationStatementWithDefaults() *MedicationStatement {
	this := MedicationStatement{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *MedicationStatement) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *MedicationStatement) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationStatement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationStatement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationStatement) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MedicationStatement) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MedicationStatement) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *MedicationStatement) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *MedicationStatement) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *MedicationStatement) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *MedicationStatement) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MedicationStatement) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MedicationStatement) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MedicationStatement) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MedicationStatement) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MedicationStatement) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *MedicationStatement) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *MedicationStatement) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *MedicationStatement) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *MedicationStatement) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationStatement) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationStatement) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationStatement) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationStatement) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationStatement) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationStatement) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MedicationStatement) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MedicationStatement) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *MedicationStatement) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *MedicationStatement) GetBasedOn() []Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret []Reference
		return ret
	}
	return o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetBasedOnOk() ([]Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *MedicationStatement) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given []Reference and assigns it to the BasedOn field.
func (o *MedicationStatement) SetBasedOn(v []Reference) {
	o.BasedOn = v
}

// GetPartOf returns the PartOf field value if set, zero value otherwise.
func (o *MedicationStatement) GetPartOf() []Reference {
	if o == nil || IsNil(o.PartOf) {
		var ret []Reference
		return ret
	}
	return o.PartOf
}

// GetPartOfOk returns a tuple with the PartOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetPartOfOk() ([]Reference, bool) {
	if o == nil || IsNil(o.PartOf) {
		return nil, false
	}
	return o.PartOf, true
}

// HasPartOf returns a boolean if a field has been set.
func (o *MedicationStatement) HasPartOf() bool {
	if o != nil && !IsNil(o.PartOf) {
		return true
	}

	return false
}

// SetPartOf gets a reference to the given []Reference and assigns it to the PartOf field.
func (o *MedicationStatement) SetPartOf(v []Reference) {
	o.PartOf = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MedicationStatement) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MedicationStatement) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MedicationStatement) SetStatus(v string) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *MedicationStatement) GetStatusReason() []CodeableConcept {
	if o == nil || IsNil(o.StatusReason) {
		var ret []CodeableConcept
		return ret
	}
	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetStatusReasonOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *MedicationStatement) HasStatusReason() bool {
	if o != nil && !IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given []CodeableConcept and assigns it to the StatusReason field.
func (o *MedicationStatement) SetStatusReason(v []CodeableConcept) {
	o.StatusReason = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MedicationStatement) GetCategory() CodeableConcept {
	if o == nil || IsNil(o.Category) {
		var ret CodeableConcept
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetCategoryOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MedicationStatement) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CodeableConcept and assigns it to the Category field.
func (o *MedicationStatement) SetCategory(v CodeableConcept) {
	o.Category = &v
}

// GetMedicationCodeableConcept returns the MedicationCodeableConcept field value if set, zero value otherwise.
func (o *MedicationStatement) GetMedicationCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.MedicationCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.MedicationCodeableConcept
}

// GetMedicationCodeableConceptOk returns a tuple with the MedicationCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetMedicationCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.MedicationCodeableConcept) {
		return nil, false
	}
	return o.MedicationCodeableConcept, true
}

// HasMedicationCodeableConcept returns a boolean if a field has been set.
func (o *MedicationStatement) HasMedicationCodeableConcept() bool {
	if o != nil && !IsNil(o.MedicationCodeableConcept) {
		return true
	}

	return false
}

// SetMedicationCodeableConcept gets a reference to the given CodeableConcept and assigns it to the MedicationCodeableConcept field.
func (o *MedicationStatement) SetMedicationCodeableConcept(v CodeableConcept) {
	o.MedicationCodeableConcept = &v
}

// GetMedicationReference returns the MedicationReference field value if set, zero value otherwise.
func (o *MedicationStatement) GetMedicationReference() Reference {
	if o == nil || IsNil(o.MedicationReference) {
		var ret Reference
		return ret
	}
	return *o.MedicationReference
}

// GetMedicationReferenceOk returns a tuple with the MedicationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetMedicationReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.MedicationReference) {
		return nil, false
	}
	return o.MedicationReference, true
}

// HasMedicationReference returns a boolean if a field has been set.
func (o *MedicationStatement) HasMedicationReference() bool {
	if o != nil && !IsNil(o.MedicationReference) {
		return true
	}

	return false
}

// SetMedicationReference gets a reference to the given Reference and assigns it to the MedicationReference field.
func (o *MedicationStatement) SetMedicationReference(v Reference) {
	o.MedicationReference = &v
}

// GetSubject returns the Subject field value
func (o *MedicationStatement) GetSubject() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetSubjectOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *MedicationStatement) SetSubject(v Reference) {
	o.Subject = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MedicationStatement) GetContext() Reference {
	if o == nil || IsNil(o.Context) {
		var ret Reference
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetContextOk() (*Reference, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MedicationStatement) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Reference and assigns it to the Context field.
func (o *MedicationStatement) SetContext(v Reference) {
	o.Context = &v
}

// GetEffectiveDateTime returns the EffectiveDateTime field value if set, zero value otherwise.
func (o *MedicationStatement) GetEffectiveDateTime() string {
	if o == nil || IsNil(o.EffectiveDateTime) {
		var ret string
		return ret
	}
	return *o.EffectiveDateTime
}

// GetEffectiveDateTimeOk returns a tuple with the EffectiveDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetEffectiveDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDateTime) {
		return nil, false
	}
	return o.EffectiveDateTime, true
}

// HasEffectiveDateTime returns a boolean if a field has been set.
func (o *MedicationStatement) HasEffectiveDateTime() bool {
	if o != nil && !IsNil(o.EffectiveDateTime) {
		return true
	}

	return false
}

// SetEffectiveDateTime gets a reference to the given string and assigns it to the EffectiveDateTime field.
func (o *MedicationStatement) SetEffectiveDateTime(v string) {
	o.EffectiveDateTime = &v
}

// GetEffectivePeriod returns the EffectivePeriod field value if set, zero value otherwise.
func (o *MedicationStatement) GetEffectivePeriod() Period {
	if o == nil || IsNil(o.EffectivePeriod) {
		var ret Period
		return ret
	}
	return *o.EffectivePeriod
}

// GetEffectivePeriodOk returns a tuple with the EffectivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetEffectivePeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.EffectivePeriod) {
		return nil, false
	}
	return o.EffectivePeriod, true
}

// HasEffectivePeriod returns a boolean if a field has been set.
func (o *MedicationStatement) HasEffectivePeriod() bool {
	if o != nil && !IsNil(o.EffectivePeriod) {
		return true
	}

	return false
}

// SetEffectivePeriod gets a reference to the given Period and assigns it to the EffectivePeriod field.
func (o *MedicationStatement) SetEffectivePeriod(v Period) {
	o.EffectivePeriod = &v
}

// GetDateAsserted returns the DateAsserted field value if set, zero value otherwise.
func (o *MedicationStatement) GetDateAsserted() string {
	if o == nil || IsNil(o.DateAsserted) {
		var ret string
		return ret
	}
	return *o.DateAsserted
}

// GetDateAssertedOk returns a tuple with the DateAsserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetDateAssertedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAsserted) {
		return nil, false
	}
	return o.DateAsserted, true
}

// HasDateAsserted returns a boolean if a field has been set.
func (o *MedicationStatement) HasDateAsserted() bool {
	if o != nil && !IsNil(o.DateAsserted) {
		return true
	}

	return false
}

// SetDateAsserted gets a reference to the given string and assigns it to the DateAsserted field.
func (o *MedicationStatement) SetDateAsserted(v string) {
	o.DateAsserted = &v
}

// GetInformationSource returns the InformationSource field value if set, zero value otherwise.
func (o *MedicationStatement) GetInformationSource() Reference {
	if o == nil || IsNil(o.InformationSource) {
		var ret Reference
		return ret
	}
	return *o.InformationSource
}

// GetInformationSourceOk returns a tuple with the InformationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetInformationSourceOk() (*Reference, bool) {
	if o == nil || IsNil(o.InformationSource) {
		return nil, false
	}
	return o.InformationSource, true
}

// HasInformationSource returns a boolean if a field has been set.
func (o *MedicationStatement) HasInformationSource() bool {
	if o != nil && !IsNil(o.InformationSource) {
		return true
	}

	return false
}

// SetInformationSource gets a reference to the given Reference and assigns it to the InformationSource field.
func (o *MedicationStatement) SetInformationSource(v Reference) {
	o.InformationSource = &v
}

// GetDerivedFrom returns the DerivedFrom field value if set, zero value otherwise.
func (o *MedicationStatement) GetDerivedFrom() []Reference {
	if o == nil || IsNil(o.DerivedFrom) {
		var ret []Reference
		return ret
	}
	return o.DerivedFrom
}

// GetDerivedFromOk returns a tuple with the DerivedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetDerivedFromOk() ([]Reference, bool) {
	if o == nil || IsNil(o.DerivedFrom) {
		return nil, false
	}
	return o.DerivedFrom, true
}

// HasDerivedFrom returns a boolean if a field has been set.
func (o *MedicationStatement) HasDerivedFrom() bool {
	if o != nil && !IsNil(o.DerivedFrom) {
		return true
	}

	return false
}

// SetDerivedFrom gets a reference to the given []Reference and assigns it to the DerivedFrom field.
func (o *MedicationStatement) SetDerivedFrom(v []Reference) {
	o.DerivedFrom = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *MedicationStatement) GetReasonCode() []CodeableConcept {
	if o == nil || IsNil(o.ReasonCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetReasonCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *MedicationStatement) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given []CodeableConcept and assigns it to the ReasonCode field.
func (o *MedicationStatement) SetReasonCode(v []CodeableConcept) {
	o.ReasonCode = v
}

// GetReasonReference returns the ReasonReference field value if set, zero value otherwise.
func (o *MedicationStatement) GetReasonReference() []Reference {
	if o == nil || IsNil(o.ReasonReference) {
		var ret []Reference
		return ret
	}
	return o.ReasonReference
}

// GetReasonReferenceOk returns a tuple with the ReasonReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetReasonReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ReasonReference) {
		return nil, false
	}
	return o.ReasonReference, true
}

// HasReasonReference returns a boolean if a field has been set.
func (o *MedicationStatement) HasReasonReference() bool {
	if o != nil && !IsNil(o.ReasonReference) {
		return true
	}

	return false
}

// SetReasonReference gets a reference to the given []Reference and assigns it to the ReasonReference field.
func (o *MedicationStatement) SetReasonReference(v []Reference) {
	o.ReasonReference = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *MedicationStatement) GetNote() []Annotation {
	if o == nil || IsNil(o.Note) {
		var ret []Annotation
		return ret
	}
	return o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetNoteOk() ([]Annotation, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *MedicationStatement) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given []Annotation and assigns it to the Note field.
func (o *MedicationStatement) SetNote(v []Annotation) {
	o.Note = v
}

// GetDosage returns the Dosage field value if set, zero value otherwise.
func (o *MedicationStatement) GetDosage() []Dosage {
	if o == nil || IsNil(o.Dosage) {
		var ret []Dosage
		return ret
	}
	return o.Dosage
}

// GetDosageOk returns a tuple with the Dosage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationStatement) GetDosageOk() ([]Dosage, bool) {
	if o == nil || IsNil(o.Dosage) {
		return nil, false
	}
	return o.Dosage, true
}

// HasDosage returns a boolean if a field has been set.
func (o *MedicationStatement) HasDosage() bool {
	if o != nil && !IsNil(o.Dosage) {
		return true
	}

	return false
}

// SetDosage gets a reference to the given []Dosage and assigns it to the Dosage field.
func (o *MedicationStatement) SetDosage(v []Dosage) {
	o.Dosage = v
}

func (o MedicationStatement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationStatement) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PartOf) {
		toSerialize["partOf"] = o.PartOf
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.MedicationCodeableConcept) {
		toSerialize["medicationCodeableConcept"] = o.MedicationCodeableConcept
	}
	if !IsNil(o.MedicationReference) {
		toSerialize["medicationReference"] = o.MedicationReference
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.EffectiveDateTime) {
		toSerialize["effectiveDateTime"] = o.EffectiveDateTime
	}
	if !IsNil(o.EffectivePeriod) {
		toSerialize["effectivePeriod"] = o.EffectivePeriod
	}
	if !IsNil(o.DateAsserted) {
		toSerialize["dateAsserted"] = o.DateAsserted
	}
	if !IsNil(o.InformationSource) {
		toSerialize["informationSource"] = o.InformationSource
	}
	if !IsNil(o.DerivedFrom) {
		toSerialize["derivedFrom"] = o.DerivedFrom
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if !IsNil(o.ReasonReference) {
		toSerialize["reasonReference"] = o.ReasonReference
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.Dosage) {
		toSerialize["dosage"] = o.Dosage
	}
	return toSerialize, nil
}

func (o *MedicationStatement) UnmarshalJSON(data []byte) (err error) {
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

	varMedicationStatement := _MedicationStatement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicationStatement)

	if err != nil {
		return err
	}

	*o = MedicationStatement(varMedicationStatement)

	return err
}

type NullableMedicationStatement struct {
	value *MedicationStatement
	isSet bool
}

func (v NullableMedicationStatement) Get() *MedicationStatement {
	return v.value
}

func (v *NullableMedicationStatement) Set(val *MedicationStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationStatement(val *MedicationStatement) *NullableMedicationStatement {
	return &NullableMedicationStatement{value: val, isSet: true}
}

func (v NullableMedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


