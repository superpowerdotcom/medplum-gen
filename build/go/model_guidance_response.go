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

// checks if the GuidanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuidanceResponse{}

// GuidanceResponse A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
type GuidanceResponse struct {
	// This is a GuidanceResponse resource
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
	// The identifier of the request associated with this response. If an identifier was given as part of the request, it will be reproduced here to enable the requester to more easily identify the response in a multi-request scenario.
	RequestIdentifier *Identifier `json:"requestIdentifier,omitempty"`
	// Allows a service to provide  unique, business identifiers for the response.
	Identifier []Identifier `json:"identifier,omitempty"`
	// An identifier, CodeableConcept or canonical reference to the guidance that was requested.
	ModuleUri *string `json:"moduleUri,omitempty"`
	// An identifier, CodeableConcept or canonical reference to the guidance that was requested.
	ModuleCanonical *string `json:"moduleCanonical,omitempty"`
	// An identifier, CodeableConcept or canonical reference to the guidance that was requested.
	ModuleCodeableConcept *CodeableConcept `json:"moduleCodeableConcept,omitempty"`
	// The status of the response. If the evaluation is completed successfully, the status will indicate success. However, in order to complete the evaluation, the engine may require more information. In this case, the status will be data-required, and the response will contain a description of the additional required information. If the evaluation completed successfully, but the engine determines that a potentially more accurate response could be provided if more data was available, the status will be data-requested, and the response will contain a description of the additional requested information.
	Status *string `json:"status,omitempty"`
	// The patient for which the request was processed.
	Subject *Reference `json:"subject,omitempty"`
	// The encounter during which this response was created or to which the creation of this record is tightly associated.
	Encounter *Reference `json:"encounter,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	// Provides a reference to the device that performed the guidance.
	Performer *Reference `json:"performer,omitempty"`
	// Describes the reason for the guidance response in coded or textual form.
	ReasonCode []CodeableConcept `json:"reasonCode,omitempty"`
	// Indicates the reason the request was initiated. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonReference []Reference `json:"reasonReference,omitempty"`
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation `json:"note,omitempty"`
	// Messages resulting from the evaluation of the artifact or artifacts. As part of evaluating the request, the engine may produce informational or warning messages. These messages will be provided by this element.
	EvaluationMessage []Reference `json:"evaluationMessage,omitempty"`
	// The output parameters of the evaluation, if any. Many modules will result in the return of specific resources such as procedure or communication requests that are returned as part of the operation result. However, modules may define specific outputs that would be returned as the result of the evaluation, and these would be returned in this element.
	OutputParameters *Reference `json:"outputParameters,omitempty"`
	// The actions, if any, produced by the evaluation of the artifact.
	Result *Reference `json:"result,omitempty"`
	// If the evaluation could not be completed due to lack of information, or additional information would potentially result in a more accurate response, this element will a description of the data required in order to proceed with the evaluation. A subsequent request to the service should include this data.
	DataRequirement []DataRequirement `json:"dataRequirement,omitempty"`
}

type _GuidanceResponse GuidanceResponse

// NewGuidanceResponse instantiates a new GuidanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceResponse(resourceType string) *GuidanceResponse {
	this := GuidanceResponse{}
	this.ResourceType = resourceType
	return &this
}

// NewGuidanceResponseWithDefaults instantiates a new GuidanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceResponseWithDefaults() *GuidanceResponse {
	this := GuidanceResponse{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *GuidanceResponse) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *GuidanceResponse) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuidanceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuidanceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GuidanceResponse) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GuidanceResponse) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GuidanceResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GuidanceResponse) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *GuidanceResponse) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *GuidanceResponse) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *GuidanceResponse) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GuidanceResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GuidanceResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GuidanceResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *GuidanceResponse) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GuidanceResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *GuidanceResponse) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *GuidanceResponse) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *GuidanceResponse) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *GuidanceResponse) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *GuidanceResponse) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *GuidanceResponse) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *GuidanceResponse) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *GuidanceResponse) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *GuidanceResponse) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *GuidanceResponse) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetRequestIdentifier returns the RequestIdentifier field value if set, zero value otherwise.
func (o *GuidanceResponse) GetRequestIdentifier() Identifier {
	if o == nil || IsNil(o.RequestIdentifier) {
		var ret Identifier
		return ret
	}
	return *o.RequestIdentifier
}

// GetRequestIdentifierOk returns a tuple with the RequestIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetRequestIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.RequestIdentifier) {
		return nil, false
	}
	return o.RequestIdentifier, true
}

// HasRequestIdentifier returns a boolean if a field has been set.
func (o *GuidanceResponse) HasRequestIdentifier() bool {
	if o != nil && !IsNil(o.RequestIdentifier) {
		return true
	}

	return false
}

// SetRequestIdentifier gets a reference to the given Identifier and assigns it to the RequestIdentifier field.
func (o *GuidanceResponse) SetRequestIdentifier(v Identifier) {
	o.RequestIdentifier = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *GuidanceResponse) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *GuidanceResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *GuidanceResponse) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetModuleUri returns the ModuleUri field value if set, zero value otherwise.
func (o *GuidanceResponse) GetModuleUri() string {
	if o == nil || IsNil(o.ModuleUri) {
		var ret string
		return ret
	}
	return *o.ModuleUri
}

// GetModuleUriOk returns a tuple with the ModuleUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetModuleUriOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleUri) {
		return nil, false
	}
	return o.ModuleUri, true
}

// HasModuleUri returns a boolean if a field has been set.
func (o *GuidanceResponse) HasModuleUri() bool {
	if o != nil && !IsNil(o.ModuleUri) {
		return true
	}

	return false
}

// SetModuleUri gets a reference to the given string and assigns it to the ModuleUri field.
func (o *GuidanceResponse) SetModuleUri(v string) {
	o.ModuleUri = &v
}

// GetModuleCanonical returns the ModuleCanonical field value if set, zero value otherwise.
func (o *GuidanceResponse) GetModuleCanonical() string {
	if o == nil || IsNil(o.ModuleCanonical) {
		var ret string
		return ret
	}
	return *o.ModuleCanonical
}

// GetModuleCanonicalOk returns a tuple with the ModuleCanonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetModuleCanonicalOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleCanonical) {
		return nil, false
	}
	return o.ModuleCanonical, true
}

// HasModuleCanonical returns a boolean if a field has been set.
func (o *GuidanceResponse) HasModuleCanonical() bool {
	if o != nil && !IsNil(o.ModuleCanonical) {
		return true
	}

	return false
}

// SetModuleCanonical gets a reference to the given string and assigns it to the ModuleCanonical field.
func (o *GuidanceResponse) SetModuleCanonical(v string) {
	o.ModuleCanonical = &v
}

// GetModuleCodeableConcept returns the ModuleCodeableConcept field value if set, zero value otherwise.
func (o *GuidanceResponse) GetModuleCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ModuleCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ModuleCodeableConcept
}

// GetModuleCodeableConceptOk returns a tuple with the ModuleCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetModuleCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ModuleCodeableConcept) {
		return nil, false
	}
	return o.ModuleCodeableConcept, true
}

// HasModuleCodeableConcept returns a boolean if a field has been set.
func (o *GuidanceResponse) HasModuleCodeableConcept() bool {
	if o != nil && !IsNil(o.ModuleCodeableConcept) {
		return true
	}

	return false
}

// SetModuleCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ModuleCodeableConcept field.
func (o *GuidanceResponse) SetModuleCodeableConcept(v CodeableConcept) {
	o.ModuleCodeableConcept = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GuidanceResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GuidanceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GuidanceResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *GuidanceResponse) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *GuidanceResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *GuidanceResponse) SetSubject(v Reference) {
	o.Subject = &v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *GuidanceResponse) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *GuidanceResponse) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *GuidanceResponse) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetOccurrenceDateTime returns the OccurrenceDateTime field value if set, zero value otherwise.
func (o *GuidanceResponse) GetOccurrenceDateTime() string {
	if o == nil || IsNil(o.OccurrenceDateTime) {
		var ret string
		return ret
	}
	return *o.OccurrenceDateTime
}

// GetOccurrenceDateTimeOk returns a tuple with the OccurrenceDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetOccurrenceDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OccurrenceDateTime) {
		return nil, false
	}
	return o.OccurrenceDateTime, true
}

// HasOccurrenceDateTime returns a boolean if a field has been set.
func (o *GuidanceResponse) HasOccurrenceDateTime() bool {
	if o != nil && !IsNil(o.OccurrenceDateTime) {
		return true
	}

	return false
}

// SetOccurrenceDateTime gets a reference to the given string and assigns it to the OccurrenceDateTime field.
func (o *GuidanceResponse) SetOccurrenceDateTime(v string) {
	o.OccurrenceDateTime = &v
}

// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *GuidanceResponse) GetPerformer() Reference {
	if o == nil || IsNil(o.Performer) {
		var ret Reference
		return ret
	}
	return *o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetPerformerOk() (*Reference, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *GuidanceResponse) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given Reference and assigns it to the Performer field.
func (o *GuidanceResponse) SetPerformer(v Reference) {
	o.Performer = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *GuidanceResponse) GetReasonCode() []CodeableConcept {
	if o == nil || IsNil(o.ReasonCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetReasonCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *GuidanceResponse) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given []CodeableConcept and assigns it to the ReasonCode field.
func (o *GuidanceResponse) SetReasonCode(v []CodeableConcept) {
	o.ReasonCode = v
}

// GetReasonReference returns the ReasonReference field value if set, zero value otherwise.
func (o *GuidanceResponse) GetReasonReference() []Reference {
	if o == nil || IsNil(o.ReasonReference) {
		var ret []Reference
		return ret
	}
	return o.ReasonReference
}

// GetReasonReferenceOk returns a tuple with the ReasonReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetReasonReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ReasonReference) {
		return nil, false
	}
	return o.ReasonReference, true
}

// HasReasonReference returns a boolean if a field has been set.
func (o *GuidanceResponse) HasReasonReference() bool {
	if o != nil && !IsNil(o.ReasonReference) {
		return true
	}

	return false
}

// SetReasonReference gets a reference to the given []Reference and assigns it to the ReasonReference field.
func (o *GuidanceResponse) SetReasonReference(v []Reference) {
	o.ReasonReference = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *GuidanceResponse) GetNote() []Annotation {
	if o == nil || IsNil(o.Note) {
		var ret []Annotation
		return ret
	}
	return o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetNoteOk() ([]Annotation, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *GuidanceResponse) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given []Annotation and assigns it to the Note field.
func (o *GuidanceResponse) SetNote(v []Annotation) {
	o.Note = v
}

// GetEvaluationMessage returns the EvaluationMessage field value if set, zero value otherwise.
func (o *GuidanceResponse) GetEvaluationMessage() []Reference {
	if o == nil || IsNil(o.EvaluationMessage) {
		var ret []Reference
		return ret
	}
	return o.EvaluationMessage
}

// GetEvaluationMessageOk returns a tuple with the EvaluationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetEvaluationMessageOk() ([]Reference, bool) {
	if o == nil || IsNil(o.EvaluationMessage) {
		return nil, false
	}
	return o.EvaluationMessage, true
}

// HasEvaluationMessage returns a boolean if a field has been set.
func (o *GuidanceResponse) HasEvaluationMessage() bool {
	if o != nil && !IsNil(o.EvaluationMessage) {
		return true
	}

	return false
}

// SetEvaluationMessage gets a reference to the given []Reference and assigns it to the EvaluationMessage field.
func (o *GuidanceResponse) SetEvaluationMessage(v []Reference) {
	o.EvaluationMessage = v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise.
func (o *GuidanceResponse) GetOutputParameters() Reference {
	if o == nil || IsNil(o.OutputParameters) {
		var ret Reference
		return ret
	}
	return *o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetOutputParametersOk() (*Reference, bool) {
	if o == nil || IsNil(o.OutputParameters) {
		return nil, false
	}
	return o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *GuidanceResponse) HasOutputParameters() bool {
	if o != nil && !IsNil(o.OutputParameters) {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given Reference and assigns it to the OutputParameters field.
func (o *GuidanceResponse) SetOutputParameters(v Reference) {
	o.OutputParameters = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GuidanceResponse) GetResult() Reference {
	if o == nil || IsNil(o.Result) {
		var ret Reference
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetResultOk() (*Reference, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GuidanceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Reference and assigns it to the Result field.
func (o *GuidanceResponse) SetResult(v Reference) {
	o.Result = &v
}

// GetDataRequirement returns the DataRequirement field value if set, zero value otherwise.
func (o *GuidanceResponse) GetDataRequirement() []DataRequirement {
	if o == nil || IsNil(o.DataRequirement) {
		var ret []DataRequirement
		return ret
	}
	return o.DataRequirement
}

// GetDataRequirementOk returns a tuple with the DataRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceResponse) GetDataRequirementOk() ([]DataRequirement, bool) {
	if o == nil || IsNil(o.DataRequirement) {
		return nil, false
	}
	return o.DataRequirement, true
}

// HasDataRequirement returns a boolean if a field has been set.
func (o *GuidanceResponse) HasDataRequirement() bool {
	if o != nil && !IsNil(o.DataRequirement) {
		return true
	}

	return false
}

// SetDataRequirement gets a reference to the given []DataRequirement and assigns it to the DataRequirement field.
func (o *GuidanceResponse) SetDataRequirement(v []DataRequirement) {
	o.DataRequirement = v
}

func (o GuidanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuidanceResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RequestIdentifier) {
		toSerialize["requestIdentifier"] = o.RequestIdentifier
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.ModuleUri) {
		toSerialize["moduleUri"] = o.ModuleUri
	}
	if !IsNil(o.ModuleCanonical) {
		toSerialize["moduleCanonical"] = o.ModuleCanonical
	}
	if !IsNil(o.ModuleCodeableConcept) {
		toSerialize["moduleCodeableConcept"] = o.ModuleCodeableConcept
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.OccurrenceDateTime) {
		toSerialize["occurrenceDateTime"] = o.OccurrenceDateTime
	}
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
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
	if !IsNil(o.EvaluationMessage) {
		toSerialize["evaluationMessage"] = o.EvaluationMessage
	}
	if !IsNil(o.OutputParameters) {
		toSerialize["outputParameters"] = o.OutputParameters
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.DataRequirement) {
		toSerialize["dataRequirement"] = o.DataRequirement
	}
	return toSerialize, nil
}

func (o *GuidanceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGuidanceResponse := _GuidanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuidanceResponse)

	if err != nil {
		return err
	}

	*o = GuidanceResponse(varGuidanceResponse)

	return err
}

type NullableGuidanceResponse struct {
	value *GuidanceResponse
	isSet bool
}

func (v NullableGuidanceResponse) Get() *GuidanceResponse {
	return v.value
}

func (v *NullableGuidanceResponse) Set(val *GuidanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceResponse(val *GuidanceResponse) *NullableGuidanceResponse {
	return &NullableGuidanceResponse{value: val, isSet: true}
}

func (v NullableGuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

