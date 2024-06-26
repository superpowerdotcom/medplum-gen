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

// checks if the MedicinalProductAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductAuthorization{}

// MedicinalProductAuthorization The regulatory authorization of a medicinal product.
type MedicinalProductAuthorization struct {
	// This is a MedicinalProductAuthorization resource
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
	// Business identifier for the marketing authorization, as assigned by a regulator.
	Identifier []Identifier `json:"identifier,omitempty"`
	// The medicinal product that is being authorized.
	Subject *Reference `json:"subject,omitempty"`
	// The country in which the marketing authorization has been granted.
	Country []CodeableConcept `json:"country,omitempty"`
	// Jurisdiction within a country.
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// The status of the marketing authorization.
	Status *CodeableConcept `json:"status,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	StatusDate *string `json:"statusDate,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	RestoreDate *string `json:"restoreDate,omitempty"`
	// The beginning of the time period in which the marketing authorization is in the specific status shall be specified A complete date consisting of day, month and year shall be specified using the ISO 8601 date format.
	ValidityPeriod *Period `json:"validityPeriod,omitempty"`
	// A period of time after authorization before generic product applicatiosn can be submitted.
	DataExclusivityPeriod *Period `json:"dataExclusivityPeriod,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	DateOfFirstAuthorization *string `json:"dateOfFirstAuthorization,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	InternationalBirthDate *string `json:"internationalBirthDate,omitempty"`
	// The legal framework against which this authorization is granted.
	LegalBasis *CodeableConcept `json:"legalBasis,omitempty"`
	// Authorization in areas within a country.
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization `json:"jurisdictionalAuthorization,omitempty"`
	// Marketing Authorization Holder.
	Holder *Reference `json:"holder,omitempty"`
	// Medicines Regulatory Agency.
	Regulator *Reference `json:"regulator,omitempty"`
	// The regulatory procedure for granting or amending a marketing authorization.
	Procedure *MedicinalProductAuthorizationProcedure `json:"procedure,omitempty"`
}

type _MedicinalProductAuthorization MedicinalProductAuthorization

// NewMedicinalProductAuthorization instantiates a new MedicinalProductAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductAuthorization(resourceType string) *MedicinalProductAuthorization {
	this := MedicinalProductAuthorization{}
	this.ResourceType = resourceType
	return &this
}

// NewMedicinalProductAuthorizationWithDefaults instantiates a new MedicinalProductAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductAuthorizationWithDefaults() *MedicinalProductAuthorization {
	this := MedicinalProductAuthorization{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *MedicinalProductAuthorization) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *MedicinalProductAuthorization) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductAuthorization) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *MedicinalProductAuthorization) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *MedicinalProductAuthorization) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MedicinalProductAuthorization) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *MedicinalProductAuthorization) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *MedicinalProductAuthorization) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductAuthorization) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductAuthorization) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *MedicinalProductAuthorization) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *MedicinalProductAuthorization) SetSubject(v Reference) {
	o.Subject = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetCountry() []CodeableConcept {
	if o == nil || IsNil(o.Country) {
		var ret []CodeableConcept
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetCountryOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []CodeableConcept and assigns it to the Country field.
func (o *MedicinalProductAuthorization) SetCountry(v []CodeableConcept) {
	o.Country = v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetJurisdiction() []CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret []CodeableConcept
		return ret
	}
	return o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetJurisdictionOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given []CodeableConcept and assigns it to the Jurisdiction field.
func (o *MedicinalProductAuthorization) SetJurisdiction(v []CodeableConcept) {
	o.Jurisdiction = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetStatus() CodeableConcept {
	if o == nil || IsNil(o.Status) {
		var ret CodeableConcept
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetStatusOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CodeableConcept and assigns it to the Status field.
func (o *MedicinalProductAuthorization) SetStatus(v CodeableConcept) {
	o.Status = &v
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetStatusDate() string {
	if o == nil || IsNil(o.StatusDate) {
		var ret string
		return ret
	}
	return *o.StatusDate
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetStatusDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDate) {
		return nil, false
	}
	return o.StatusDate, true
}

// HasStatusDate returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasStatusDate() bool {
	if o != nil && !IsNil(o.StatusDate) {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given string and assigns it to the StatusDate field.
func (o *MedicinalProductAuthorization) SetStatusDate(v string) {
	o.StatusDate = &v
}

// GetRestoreDate returns the RestoreDate field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetRestoreDate() string {
	if o == nil || IsNil(o.RestoreDate) {
		var ret string
		return ret
	}
	return *o.RestoreDate
}

// GetRestoreDateOk returns a tuple with the RestoreDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetRestoreDateOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreDate) {
		return nil, false
	}
	return o.RestoreDate, true
}

// HasRestoreDate returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasRestoreDate() bool {
	if o != nil && !IsNil(o.RestoreDate) {
		return true
	}

	return false
}

// SetRestoreDate gets a reference to the given string and assigns it to the RestoreDate field.
func (o *MedicinalProductAuthorization) SetRestoreDate(v string) {
	o.RestoreDate = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetValidityPeriod() Period {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret Period
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetValidityPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given Period and assigns it to the ValidityPeriod field.
func (o *MedicinalProductAuthorization) SetValidityPeriod(v Period) {
	o.ValidityPeriod = &v
}

// GetDataExclusivityPeriod returns the DataExclusivityPeriod field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetDataExclusivityPeriod() Period {
	if o == nil || IsNil(o.DataExclusivityPeriod) {
		var ret Period
		return ret
	}
	return *o.DataExclusivityPeriod
}

// GetDataExclusivityPeriodOk returns a tuple with the DataExclusivityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetDataExclusivityPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.DataExclusivityPeriod) {
		return nil, false
	}
	return o.DataExclusivityPeriod, true
}

// HasDataExclusivityPeriod returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasDataExclusivityPeriod() bool {
	if o != nil && !IsNil(o.DataExclusivityPeriod) {
		return true
	}

	return false
}

// SetDataExclusivityPeriod gets a reference to the given Period and assigns it to the DataExclusivityPeriod field.
func (o *MedicinalProductAuthorization) SetDataExclusivityPeriod(v Period) {
	o.DataExclusivityPeriod = &v
}

// GetDateOfFirstAuthorization returns the DateOfFirstAuthorization field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetDateOfFirstAuthorization() string {
	if o == nil || IsNil(o.DateOfFirstAuthorization) {
		var ret string
		return ret
	}
	return *o.DateOfFirstAuthorization
}

// GetDateOfFirstAuthorizationOk returns a tuple with the DateOfFirstAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetDateOfFirstAuthorizationOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfFirstAuthorization) {
		return nil, false
	}
	return o.DateOfFirstAuthorization, true
}

// HasDateOfFirstAuthorization returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasDateOfFirstAuthorization() bool {
	if o != nil && !IsNil(o.DateOfFirstAuthorization) {
		return true
	}

	return false
}

// SetDateOfFirstAuthorization gets a reference to the given string and assigns it to the DateOfFirstAuthorization field.
func (o *MedicinalProductAuthorization) SetDateOfFirstAuthorization(v string) {
	o.DateOfFirstAuthorization = &v
}

// GetInternationalBirthDate returns the InternationalBirthDate field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetInternationalBirthDate() string {
	if o == nil || IsNil(o.InternationalBirthDate) {
		var ret string
		return ret
	}
	return *o.InternationalBirthDate
}

// GetInternationalBirthDateOk returns a tuple with the InternationalBirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetInternationalBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.InternationalBirthDate) {
		return nil, false
	}
	return o.InternationalBirthDate, true
}

// HasInternationalBirthDate returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasInternationalBirthDate() bool {
	if o != nil && !IsNil(o.InternationalBirthDate) {
		return true
	}

	return false
}

// SetInternationalBirthDate gets a reference to the given string and assigns it to the InternationalBirthDate field.
func (o *MedicinalProductAuthorization) SetInternationalBirthDate(v string) {
	o.InternationalBirthDate = &v
}

// GetLegalBasis returns the LegalBasis field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetLegalBasis() CodeableConcept {
	if o == nil || IsNil(o.LegalBasis) {
		var ret CodeableConcept
		return ret
	}
	return *o.LegalBasis
}

// GetLegalBasisOk returns a tuple with the LegalBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetLegalBasisOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.LegalBasis) {
		return nil, false
	}
	return o.LegalBasis, true
}

// HasLegalBasis returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasLegalBasis() bool {
	if o != nil && !IsNil(o.LegalBasis) {
		return true
	}

	return false
}

// SetLegalBasis gets a reference to the given CodeableConcept and assigns it to the LegalBasis field.
func (o *MedicinalProductAuthorization) SetLegalBasis(v CodeableConcept) {
	o.LegalBasis = &v
}

// GetJurisdictionalAuthorization returns the JurisdictionalAuthorization field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetJurisdictionalAuthorization() []MedicinalProductAuthorizationJurisdictionalAuthorization {
	if o == nil || IsNil(o.JurisdictionalAuthorization) {
		var ret []MedicinalProductAuthorizationJurisdictionalAuthorization
		return ret
	}
	return o.JurisdictionalAuthorization
}

// GetJurisdictionalAuthorizationOk returns a tuple with the JurisdictionalAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetJurisdictionalAuthorizationOk() ([]MedicinalProductAuthorizationJurisdictionalAuthorization, bool) {
	if o == nil || IsNil(o.JurisdictionalAuthorization) {
		return nil, false
	}
	return o.JurisdictionalAuthorization, true
}

// HasJurisdictionalAuthorization returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasJurisdictionalAuthorization() bool {
	if o != nil && !IsNil(o.JurisdictionalAuthorization) {
		return true
	}

	return false
}

// SetJurisdictionalAuthorization gets a reference to the given []MedicinalProductAuthorizationJurisdictionalAuthorization and assigns it to the JurisdictionalAuthorization field.
func (o *MedicinalProductAuthorization) SetJurisdictionalAuthorization(v []MedicinalProductAuthorizationJurisdictionalAuthorization) {
	o.JurisdictionalAuthorization = v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetHolder() Reference {
	if o == nil || IsNil(o.Holder) {
		var ret Reference
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetHolderOk() (*Reference, bool) {
	if o == nil || IsNil(o.Holder) {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasHolder() bool {
	if o != nil && !IsNil(o.Holder) {
		return true
	}

	return false
}

// SetHolder gets a reference to the given Reference and assigns it to the Holder field.
func (o *MedicinalProductAuthorization) SetHolder(v Reference) {
	o.Holder = &v
}

// GetRegulator returns the Regulator field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetRegulator() Reference {
	if o == nil || IsNil(o.Regulator) {
		var ret Reference
		return ret
	}
	return *o.Regulator
}

// GetRegulatorOk returns a tuple with the Regulator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetRegulatorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Regulator) {
		return nil, false
	}
	return o.Regulator, true
}

// HasRegulator returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasRegulator() bool {
	if o != nil && !IsNil(o.Regulator) {
		return true
	}

	return false
}

// SetRegulator gets a reference to the given Reference and assigns it to the Regulator field.
func (o *MedicinalProductAuthorization) SetRegulator(v Reference) {
	o.Regulator = &v
}

// GetProcedure returns the Procedure field value if set, zero value otherwise.
func (o *MedicinalProductAuthorization) GetProcedure() MedicinalProductAuthorizationProcedure {
	if o == nil || IsNil(o.Procedure) {
		var ret MedicinalProductAuthorizationProcedure
		return ret
	}
	return *o.Procedure
}

// GetProcedureOk returns a tuple with the Procedure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorization) GetProcedureOk() (*MedicinalProductAuthorizationProcedure, bool) {
	if o == nil || IsNil(o.Procedure) {
		return nil, false
	}
	return o.Procedure, true
}

// HasProcedure returns a boolean if a field has been set.
func (o *MedicinalProductAuthorization) HasProcedure() bool {
	if o != nil && !IsNil(o.Procedure) {
		return true
	}

	return false
}

// SetProcedure gets a reference to the given MedicinalProductAuthorizationProcedure and assigns it to the Procedure field.
func (o *MedicinalProductAuthorization) SetProcedure(v MedicinalProductAuthorizationProcedure) {
	o.Procedure = &v
}

func (o MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductAuthorization) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDate) {
		toSerialize["statusDate"] = o.StatusDate
	}
	if !IsNil(o.RestoreDate) {
		toSerialize["restoreDate"] = o.RestoreDate
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	if !IsNil(o.DataExclusivityPeriod) {
		toSerialize["dataExclusivityPeriod"] = o.DataExclusivityPeriod
	}
	if !IsNil(o.DateOfFirstAuthorization) {
		toSerialize["dateOfFirstAuthorization"] = o.DateOfFirstAuthorization
	}
	if !IsNil(o.InternationalBirthDate) {
		toSerialize["internationalBirthDate"] = o.InternationalBirthDate
	}
	if !IsNil(o.LegalBasis) {
		toSerialize["legalBasis"] = o.LegalBasis
	}
	if !IsNil(o.JurisdictionalAuthorization) {
		toSerialize["jurisdictionalAuthorization"] = o.JurisdictionalAuthorization
	}
	if !IsNil(o.Holder) {
		toSerialize["holder"] = o.Holder
	}
	if !IsNil(o.Regulator) {
		toSerialize["regulator"] = o.Regulator
	}
	if !IsNil(o.Procedure) {
		toSerialize["procedure"] = o.Procedure
	}
	return toSerialize, nil
}

func (o *MedicinalProductAuthorization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
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

	varMedicinalProductAuthorization := _MedicinalProductAuthorization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicinalProductAuthorization)

	if err != nil {
		return err
	}

	*o = MedicinalProductAuthorization(varMedicinalProductAuthorization)

	return err
}

type NullableMedicinalProductAuthorization struct {
	value *MedicinalProductAuthorization
	isSet bool
}

func (v NullableMedicinalProductAuthorization) Get() *MedicinalProductAuthorization {
	return v.value
}

func (v *NullableMedicinalProductAuthorization) Set(val *MedicinalProductAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductAuthorization(val *MedicinalProductAuthorization) *NullableMedicinalProductAuthorization {
	return &NullableMedicinalProductAuthorization{value: val, isSet: true}
}

func (v NullableMedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

