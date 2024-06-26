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

// checks if the Media type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Media{}

// Media A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {
	// This is a Media resource
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
	// Identifiers associated with the image - these may include identifiers for the image itself, identifiers for the context of its collection (e.g. series ids) and context ids such as accession numbers or other workflow identifiers.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A procedure that is fulfilled in whole or in part by the creation of this media.
	BasedOn []Reference `json:"basedOn,omitempty"`
	// A larger event of which this particular event is a component or step.
	PartOf []Reference `json:"partOf,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Status *string `json:"status,omitempty"`
	// A code that classifies whether the media is an image, video or audio recording or some other media category.
	Type *CodeableConcept `json:"type,omitempty"`
	// Details of the type of the media - usually, how it was acquired (what type of device). If images sourced from a DICOM system, are wrapped in a Media resource, then this is the modality.
	Modality *CodeableConcept `json:"modality,omitempty"`
	// The name of the imaging view e.g. Lateral or Antero-posterior (AP).
	View *CodeableConcept `json:"view,omitempty"`
	// Who/What this Media is a record of.
	Subject *Reference `json:"subject,omitempty"`
	// The encounter that establishes the context for this media.
	Encounter *Reference `json:"encounter,omitempty"`
	// The date and time(s) at which the media was collected.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	// The date and time(s) at which the media was collected.
	CreatedPeriod *Period `json:"createdPeriod,omitempty"`
	// An instant in time - known at least to the second
	Issued *string `json:"issued,omitempty"`
	// The person who administered the collection of the image.
	Operator *Reference `json:"operator,omitempty"`
	// Describes why the event occurred in coded or textual form.
	ReasonCode []CodeableConcept `json:"reasonCode,omitempty"`
	// Indicates the site on the subject's body where the observation was made (i.e. the target site).
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// A sequence of Unicode characters
	DeviceName *string `json:"deviceName,omitempty"`
	// The device used to collect the media.
	Device *Reference `json:"device,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Height *float32 `json:"height,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Width *float32 `json:"width,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Frames *float32 `json:"frames,omitempty"`
	// A rational number with implicit precision
	Duration *float32 `json:"duration,omitempty"`
	// The actual content of the media - inline or by direct reference to the media source file.
	Content Attachment `json:"content"`
	// Comments made about the media by the performer, subject or other participants.
	Note []Annotation `json:"note,omitempty"`
}

type _Media Media

// NewMedia instantiates a new Media object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedia(resourceType string, content Attachment) *Media {
	this := Media{}
	this.ResourceType = resourceType
	this.Content = content
	return &this
}

// NewMediaWithDefaults instantiates a new Media object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaWithDefaults() *Media {
	this := Media{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Media) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Media) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Media) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Media) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Media) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Media) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Media) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Media) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Media) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Media) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Media) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Media) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Media) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Media) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Media) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Media) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Media) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Media) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Media) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Media) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Media) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Media) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Media) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Media) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Media) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Media) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Media) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Media) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Media) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Media) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *Media) GetBasedOn() []Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret []Reference
		return ret
	}
	return o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetBasedOnOk() ([]Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *Media) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given []Reference and assigns it to the BasedOn field.
func (o *Media) SetBasedOn(v []Reference) {
	o.BasedOn = v
}

// GetPartOf returns the PartOf field value if set, zero value otherwise.
func (o *Media) GetPartOf() []Reference {
	if o == nil || IsNil(o.PartOf) {
		var ret []Reference
		return ret
	}
	return o.PartOf
}

// GetPartOfOk returns a tuple with the PartOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetPartOfOk() ([]Reference, bool) {
	if o == nil || IsNil(o.PartOf) {
		return nil, false
	}
	return o.PartOf, true
}

// HasPartOf returns a boolean if a field has been set.
func (o *Media) HasPartOf() bool {
	if o != nil && !IsNil(o.PartOf) {
		return true
	}

	return false
}

// SetPartOf gets a reference to the given []Reference and assigns it to the PartOf field.
func (o *Media) SetPartOf(v []Reference) {
	o.PartOf = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Media) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Media) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Media) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Media) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Media) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *Media) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetModality returns the Modality field value if set, zero value otherwise.
func (o *Media) GetModality() CodeableConcept {
	if o == nil || IsNil(o.Modality) {
		var ret CodeableConcept
		return ret
	}
	return *o.Modality
}

// GetModalityOk returns a tuple with the Modality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetModalityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Modality) {
		return nil, false
	}
	return o.Modality, true
}

// HasModality returns a boolean if a field has been set.
func (o *Media) HasModality() bool {
	if o != nil && !IsNil(o.Modality) {
		return true
	}

	return false
}

// SetModality gets a reference to the given CodeableConcept and assigns it to the Modality field.
func (o *Media) SetModality(v CodeableConcept) {
	o.Modality = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *Media) GetView() CodeableConcept {
	if o == nil || IsNil(o.View) {
		var ret CodeableConcept
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetViewOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *Media) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given CodeableConcept and assigns it to the View field.
func (o *Media) SetView(v CodeableConcept) {
	o.View = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Media) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Media) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *Media) SetSubject(v Reference) {
	o.Subject = &v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *Media) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *Media) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *Media) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *Media) GetCreatedDateTime() string {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Media) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *Media) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetCreatedPeriod returns the CreatedPeriod field value if set, zero value otherwise.
func (o *Media) GetCreatedPeriod() Period {
	if o == nil || IsNil(o.CreatedPeriod) {
		var ret Period
		return ret
	}
	return *o.CreatedPeriod
}

// GetCreatedPeriodOk returns a tuple with the CreatedPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetCreatedPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.CreatedPeriod) {
		return nil, false
	}
	return o.CreatedPeriod, true
}

// HasCreatedPeriod returns a boolean if a field has been set.
func (o *Media) HasCreatedPeriod() bool {
	if o != nil && !IsNil(o.CreatedPeriod) {
		return true
	}

	return false
}

// SetCreatedPeriod gets a reference to the given Period and assigns it to the CreatedPeriod field.
func (o *Media) SetCreatedPeriod(v Period) {
	o.CreatedPeriod = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *Media) GetIssued() string {
	if o == nil || IsNil(o.Issued) {
		var ret string
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetIssuedOk() (*string, bool) {
	if o == nil || IsNil(o.Issued) {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *Media) HasIssued() bool {
	if o != nil && !IsNil(o.Issued) {
		return true
	}

	return false
}

// SetIssued gets a reference to the given string and assigns it to the Issued field.
func (o *Media) SetIssued(v string) {
	o.Issued = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Media) GetOperator() Reference {
	if o == nil || IsNil(o.Operator) {
		var ret Reference
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetOperatorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Media) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given Reference and assigns it to the Operator field.
func (o *Media) SetOperator(v Reference) {
	o.Operator = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *Media) GetReasonCode() []CodeableConcept {
	if o == nil || IsNil(o.ReasonCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetReasonCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *Media) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given []CodeableConcept and assigns it to the ReasonCode field.
func (o *Media) SetReasonCode(v []CodeableConcept) {
	o.ReasonCode = v
}

// GetBodySite returns the BodySite field value if set, zero value otherwise.
func (o *Media) GetBodySite() CodeableConcept {
	if o == nil || IsNil(o.BodySite) {
		var ret CodeableConcept
		return ret
	}
	return *o.BodySite
}

// GetBodySiteOk returns a tuple with the BodySite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetBodySiteOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.BodySite) {
		return nil, false
	}
	return o.BodySite, true
}

// HasBodySite returns a boolean if a field has been set.
func (o *Media) HasBodySite() bool {
	if o != nil && !IsNil(o.BodySite) {
		return true
	}

	return false
}

// SetBodySite gets a reference to the given CodeableConcept and assigns it to the BodySite field.
func (o *Media) SetBodySite(v CodeableConcept) {
	o.BodySite = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *Media) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *Media) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *Media) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Media) GetDevice() Reference {
	if o == nil || IsNil(o.Device) {
		var ret Reference
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetDeviceOk() (*Reference, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Media) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Reference and assigns it to the Device field.
func (o *Media) SetDevice(v Reference) {
	o.Device = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Media) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Media) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *Media) SetHeight(v float32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Media) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Media) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *Media) SetWidth(v float32) {
	o.Width = &v
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *Media) GetFrames() float32 {
	if o == nil || IsNil(o.Frames) {
		var ret float32
		return ret
	}
	return *o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetFramesOk() (*float32, bool) {
	if o == nil || IsNil(o.Frames) {
		return nil, false
	}
	return o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *Media) HasFrames() bool {
	if o != nil && !IsNil(o.Frames) {
		return true
	}

	return false
}

// SetFrames gets a reference to the given float32 and assigns it to the Frames field.
func (o *Media) SetFrames(v float32) {
	o.Frames = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Media) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Media) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *Media) SetDuration(v float32) {
	o.Duration = &v
}

// GetContent returns the Content field value
func (o *Media) GetContent() Attachment {
	if o == nil {
		var ret Attachment
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Media) GetContentOk() (*Attachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Media) SetContent(v Attachment) {
	o.Content = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *Media) GetNote() []Annotation {
	if o == nil || IsNil(o.Note) {
		var ret []Annotation
		return ret
	}
	return o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetNoteOk() ([]Annotation, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *Media) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given []Annotation and assigns it to the Note field.
func (o *Media) SetNote(v []Annotation) {
	o.Note = v
}

func (o Media) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Media) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Modality) {
		toSerialize["modality"] = o.Modality
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.CreatedPeriod) {
		toSerialize["createdPeriod"] = o.CreatedPeriod
	}
	if !IsNil(o.Issued) {
		toSerialize["issued"] = o.Issued
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if !IsNil(o.BodySite) {
		toSerialize["bodySite"] = o.BodySite
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Frames) {
		toSerialize["frames"] = o.Frames
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["content"] = o.Content
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

func (o *Media) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"content",
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

	varMedia := _Media{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedia)

	if err != nil {
		return err
	}

	*o = Media(varMedia)

	return err
}

type NullableMedia struct {
	value *Media
	isSet bool
}

func (v NullableMedia) Get() *Media {
	return v.value
}

func (v *NullableMedia) Set(val *Media) {
	v.value = val
	v.isSet = true
}

func (v NullableMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedia(val *Media) *NullableMedia {
	return &NullableMedia{value: val, isSet: true}
}

func (v NullableMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


