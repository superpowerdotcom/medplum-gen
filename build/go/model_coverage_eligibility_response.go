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

// checks if the CoverageEligibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoverageEligibilityResponse{}

// CoverageEligibilityResponse This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
	// This is a CoverageEligibilityResponse resource
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
	// A unique identifier assigned to this coverage eligiblity request.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Status *string `json:"status,omitempty"`
	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes; benefits for coverages specified or discovered; discovery and return of coverages for the patient; and/or validation that the specified coverage is in-force at the date/period specified or 'now' if not specified.
	Purpose []string `json:"purpose,omitempty"`
	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought.
	Patient Reference `json:"patient"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedDate *string `json:"servicedDate,omitempty"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedPeriod *Period `json:"servicedPeriod,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Created *string `json:"created,omitempty"`
	// The provider which is responsible for the request.
	Requestor *Reference `json:"requestor,omitempty"`
	// Reference to the original request resource.
	Request Reference `json:"request"`
	// The outcome of the request processing.
	Outcome *string `json:"outcome,omitempty"`
	// A sequence of Unicode characters
	Disposition *string `json:"disposition,omitempty"`
	// The Insurer who issued the coverage in question and is the author of the response.
	Insurer Reference `json:"insurer"`
	// Financial instruments for reimbursement for the health care products and services.
	Insurance []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	// A sequence of Unicode characters
	PreAuthRef *string `json:"preAuthRef,omitempty"`
	// A code for the form to be used for printing the content.
	Form *CodeableConcept `json:"form,omitempty"`
	// Errors encountered during the processing of the request.
	Error []CoverageEligibilityResponseError `json:"error,omitempty"`
}

type _CoverageEligibilityResponse CoverageEligibilityResponse

// NewCoverageEligibilityResponse instantiates a new CoverageEligibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoverageEligibilityResponse(resourceType string, patient Reference, request Reference, insurer Reference) *CoverageEligibilityResponse {
	this := CoverageEligibilityResponse{}
	this.ResourceType = resourceType
	this.Patient = patient
	this.Request = request
	this.Insurer = insurer
	return &this
}

// NewCoverageEligibilityResponseWithDefaults instantiates a new CoverageEligibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoverageEligibilityResponseWithDefaults() *CoverageEligibilityResponse {
	this := CoverageEligibilityResponse{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *CoverageEligibilityResponse) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CoverageEligibilityResponse) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoverageEligibilityResponse) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *CoverageEligibilityResponse) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *CoverageEligibilityResponse) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *CoverageEligibilityResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *CoverageEligibilityResponse) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *CoverageEligibilityResponse) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CoverageEligibilityResponse) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CoverageEligibilityResponse) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *CoverageEligibilityResponse) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CoverageEligibilityResponse) SetStatus(v string) {
	o.Status = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetPurpose() []string {
	if o == nil || IsNil(o.Purpose) {
		var ret []string
		return ret
	}
	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetPurposeOk() ([]string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given []string and assigns it to the Purpose field.
func (o *CoverageEligibilityResponse) SetPurpose(v []string) {
	o.Purpose = v
}

// GetPatient returns the Patient field value
func (o *CoverageEligibilityResponse) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *CoverageEligibilityResponse) SetPatient(v Reference) {
	o.Patient = v
}

// GetServicedDate returns the ServicedDate field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetServicedDate() string {
	if o == nil || IsNil(o.ServicedDate) {
		var ret string
		return ret
	}
	return *o.ServicedDate
}

// GetServicedDateOk returns a tuple with the ServicedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetServicedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ServicedDate) {
		return nil, false
	}
	return o.ServicedDate, true
}

// HasServicedDate returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasServicedDate() bool {
	if o != nil && !IsNil(o.ServicedDate) {
		return true
	}

	return false
}

// SetServicedDate gets a reference to the given string and assigns it to the ServicedDate field.
func (o *CoverageEligibilityResponse) SetServicedDate(v string) {
	o.ServicedDate = &v
}

// GetServicedPeriod returns the ServicedPeriod field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetServicedPeriod() Period {
	if o == nil || IsNil(o.ServicedPeriod) {
		var ret Period
		return ret
	}
	return *o.ServicedPeriod
}

// GetServicedPeriodOk returns a tuple with the ServicedPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetServicedPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.ServicedPeriod) {
		return nil, false
	}
	return o.ServicedPeriod, true
}

// HasServicedPeriod returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasServicedPeriod() bool {
	if o != nil && !IsNil(o.ServicedPeriod) {
		return true
	}

	return false
}

// SetServicedPeriod gets a reference to the given Period and assigns it to the ServicedPeriod field.
func (o *CoverageEligibilityResponse) SetServicedPeriod(v Period) {
	o.ServicedPeriod = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CoverageEligibilityResponse) SetCreated(v string) {
	o.Created = &v
}

// GetRequestor returns the Requestor field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetRequestor() Reference {
	if o == nil || IsNil(o.Requestor) {
		var ret Reference
		return ret
	}
	return *o.Requestor
}

// GetRequestorOk returns a tuple with the Requestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetRequestorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Requestor) {
		return nil, false
	}
	return o.Requestor, true
}

// HasRequestor returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasRequestor() bool {
	if o != nil && !IsNil(o.Requestor) {
		return true
	}

	return false
}

// SetRequestor gets a reference to the given Reference and assigns it to the Requestor field.
func (o *CoverageEligibilityResponse) SetRequestor(v Reference) {
	o.Requestor = &v
}

// GetRequest returns the Request field value
func (o *CoverageEligibilityResponse) GetRequest() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetRequestOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *CoverageEligibilityResponse) SetRequest(v Reference) {
	o.Request = v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *CoverageEligibilityResponse) SetOutcome(v string) {
	o.Outcome = &v
}

// GetDisposition returns the Disposition field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetDisposition() string {
	if o == nil || IsNil(o.Disposition) {
		var ret string
		return ret
	}
	return *o.Disposition
}

// GetDispositionOk returns a tuple with the Disposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetDispositionOk() (*string, bool) {
	if o == nil || IsNil(o.Disposition) {
		return nil, false
	}
	return o.Disposition, true
}

// HasDisposition returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasDisposition() bool {
	if o != nil && !IsNil(o.Disposition) {
		return true
	}

	return false
}

// SetDisposition gets a reference to the given string and assigns it to the Disposition field.
func (o *CoverageEligibilityResponse) SetDisposition(v string) {
	o.Disposition = &v
}

// GetInsurer returns the Insurer field value
func (o *CoverageEligibilityResponse) GetInsurer() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Insurer
}

// GetInsurerOk returns a tuple with the Insurer field value
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetInsurerOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Insurer, true
}

// SetInsurer sets field value
func (o *CoverageEligibilityResponse) SetInsurer(v Reference) {
	o.Insurer = v
}

// GetInsurance returns the Insurance field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetInsurance() []CoverageEligibilityResponseInsurance {
	if o == nil || IsNil(o.Insurance) {
		var ret []CoverageEligibilityResponseInsurance
		return ret
	}
	return o.Insurance
}

// GetInsuranceOk returns a tuple with the Insurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetInsuranceOk() ([]CoverageEligibilityResponseInsurance, bool) {
	if o == nil || IsNil(o.Insurance) {
		return nil, false
	}
	return o.Insurance, true
}

// HasInsurance returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasInsurance() bool {
	if o != nil && !IsNil(o.Insurance) {
		return true
	}

	return false
}

// SetInsurance gets a reference to the given []CoverageEligibilityResponseInsurance and assigns it to the Insurance field.
func (o *CoverageEligibilityResponse) SetInsurance(v []CoverageEligibilityResponseInsurance) {
	o.Insurance = v
}

// GetPreAuthRef returns the PreAuthRef field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetPreAuthRef() string {
	if o == nil || IsNil(o.PreAuthRef) {
		var ret string
		return ret
	}
	return *o.PreAuthRef
}

// GetPreAuthRefOk returns a tuple with the PreAuthRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetPreAuthRefOk() (*string, bool) {
	if o == nil || IsNil(o.PreAuthRef) {
		return nil, false
	}
	return o.PreAuthRef, true
}

// HasPreAuthRef returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasPreAuthRef() bool {
	if o != nil && !IsNil(o.PreAuthRef) {
		return true
	}

	return false
}

// SetPreAuthRef gets a reference to the given string and assigns it to the PreAuthRef field.
func (o *CoverageEligibilityResponse) SetPreAuthRef(v string) {
	o.PreAuthRef = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetForm() CodeableConcept {
	if o == nil || IsNil(o.Form) {
		var ret CodeableConcept
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetFormOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasForm() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given CodeableConcept and assigns it to the Form field.
func (o *CoverageEligibilityResponse) SetForm(v CodeableConcept) {
	o.Form = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CoverageEligibilityResponse) GetError() []CoverageEligibilityResponseError {
	if o == nil || IsNil(o.Error) {
		var ret []CoverageEligibilityResponseError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageEligibilityResponse) GetErrorOk() ([]CoverageEligibilityResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CoverageEligibilityResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []CoverageEligibilityResponseError and assigns it to the Error field.
func (o *CoverageEligibilityResponse) SetError(v []CoverageEligibilityResponseError) {
	o.Error = v
}

func (o CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoverageEligibilityResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	toSerialize["patient"] = o.Patient
	if !IsNil(o.ServicedDate) {
		toSerialize["servicedDate"] = o.ServicedDate
	}
	if !IsNil(o.ServicedPeriod) {
		toSerialize["servicedPeriod"] = o.ServicedPeriod
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Requestor) {
		toSerialize["requestor"] = o.Requestor
	}
	toSerialize["request"] = o.Request
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.Disposition) {
		toSerialize["disposition"] = o.Disposition
	}
	toSerialize["insurer"] = o.Insurer
	if !IsNil(o.Insurance) {
		toSerialize["insurance"] = o.Insurance
	}
	if !IsNil(o.PreAuthRef) {
		toSerialize["preAuthRef"] = o.PreAuthRef
	}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *CoverageEligibilityResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"patient",
		"request",
		"insurer",
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

	varCoverageEligibilityResponse := _CoverageEligibilityResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoverageEligibilityResponse)

	if err != nil {
		return err
	}

	*o = CoverageEligibilityResponse(varCoverageEligibilityResponse)

	return err
}

type NullableCoverageEligibilityResponse struct {
	value *CoverageEligibilityResponse
	isSet bool
}

func (v NullableCoverageEligibilityResponse) Get() *CoverageEligibilityResponse {
	return v.value
}

func (v *NullableCoverageEligibilityResponse) Set(val *CoverageEligibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverageEligibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverageEligibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverageEligibilityResponse(val *CoverageEligibilityResponse) *NullableCoverageEligibilityResponse {
	return &NullableCoverageEligibilityResponse{value: val, isSet: true}
}

func (v NullableCoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverageEligibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


