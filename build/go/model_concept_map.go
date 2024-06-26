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

// checks if the ConceptMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConceptMap{}

// ConceptMap A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMap struct {
	// This is a ConceptMap resource
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
	// String of characters used to identify a name or a resource
	Url *string `json:"url,omitempty"`
	// A formal identifier that is used to identify this concept map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier,omitempty"`
	// A sequence of Unicode characters
	Version *string `json:"version,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A sequence of Unicode characters
	Title *string `json:"title,omitempty"`
	// The status of this concept map. Enables tracking the life-cycle of the content.
	Status *string `json:"status,omitempty"`
	// Value of \"true\" or \"false\"
	Experimental *bool `json:"experimental,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Date *string `json:"date,omitempty"`
	// A sequence of Unicode characters
	Publisher *string `json:"publisher,omitempty"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail `json:"contact,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Description *string `json:"description,omitempty"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate concept map instances.
	UseContext []UsageContext `json:"useContext,omitempty"`
	// A legal or geographic region in which the concept map is intended to be used.
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Purpose *string `json:"purpose,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Copyright *string `json:"copyright,omitempty"`
	// Identifier for the source value set that contains the concepts that are being mapped and provides context for the mappings.
	SourceUri *string `json:"sourceUri,omitempty"`
	// Identifier for the source value set that contains the concepts that are being mapped and provides context for the mappings.
	SourceCanonical *string `json:"sourceCanonical,omitempty"`
	// The target value set provides context for the mappings. Note that the mapping is made between concepts, not between value sets, but the value set provides important context about how the concept mapping choices are made.
	TargetUri *string `json:"targetUri,omitempty"`
	// The target value set provides context for the mappings. Note that the mapping is made between concepts, not between value sets, but the value set provides important context about how the concept mapping choices are made.
	TargetCanonical *string `json:"targetCanonical,omitempty"`
	// A group of mappings that all have the same source and target system.
	Group []ConceptMapGroup `json:"group,omitempty"`
}

type _ConceptMap ConceptMap

// NewConceptMap instantiates a new ConceptMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConceptMap(resourceType string) *ConceptMap {
	this := ConceptMap{}
	this.ResourceType = resourceType
	return &this
}

// NewConceptMapWithDefaults instantiates a new ConceptMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConceptMapWithDefaults() *ConceptMap {
	this := ConceptMap{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *ConceptMap) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ConceptMap) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConceptMap) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConceptMap) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConceptMap) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConceptMap) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConceptMap) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *ConceptMap) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *ConceptMap) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *ConceptMap) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *ConceptMap) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ConceptMap) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ConceptMap) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ConceptMap) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ConceptMap) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ConceptMap) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *ConceptMap) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *ConceptMap) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *ConceptMap) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *ConceptMap) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ConceptMap) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ConceptMap) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ConceptMap) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ConceptMap) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ConceptMap) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ConceptMap) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ConceptMap) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ConceptMap) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ConceptMap) SetUrl(v string) {
	o.Url = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ConceptMap) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ConceptMap) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *ConceptMap) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ConceptMap) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ConceptMap) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ConceptMap) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConceptMap) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConceptMap) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConceptMap) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ConceptMap) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ConceptMap) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ConceptMap) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConceptMap) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConceptMap) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConceptMap) SetStatus(v string) {
	o.Status = &v
}

// GetExperimental returns the Experimental field value if set, zero value otherwise.
func (o *ConceptMap) GetExperimental() bool {
	if o == nil || IsNil(o.Experimental) {
		var ret bool
		return ret
	}
	return *o.Experimental
}

// GetExperimentalOk returns a tuple with the Experimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetExperimentalOk() (*bool, bool) {
	if o == nil || IsNil(o.Experimental) {
		return nil, false
	}
	return o.Experimental, true
}

// HasExperimental returns a boolean if a field has been set.
func (o *ConceptMap) HasExperimental() bool {
	if o != nil && !IsNil(o.Experimental) {
		return true
	}

	return false
}

// SetExperimental gets a reference to the given bool and assigns it to the Experimental field.
func (o *ConceptMap) SetExperimental(v bool) {
	o.Experimental = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ConceptMap) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ConceptMap) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ConceptMap) SetDate(v string) {
	o.Date = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *ConceptMap) GetPublisher() string {
	if o == nil || IsNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetPublisherOk() (*string, bool) {
	if o == nil || IsNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *ConceptMap) HasPublisher() bool {
	if o != nil && !IsNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *ConceptMap) SetPublisher(v string) {
	o.Publisher = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *ConceptMap) GetContact() []ContactDetail {
	if o == nil || IsNil(o.Contact) {
		var ret []ContactDetail
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetContactOk() ([]ContactDetail, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ConceptMap) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given []ContactDetail and assigns it to the Contact field.
func (o *ConceptMap) SetContact(v []ContactDetail) {
	o.Contact = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConceptMap) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConceptMap) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConceptMap) SetDescription(v string) {
	o.Description = &v
}

// GetUseContext returns the UseContext field value if set, zero value otherwise.
func (o *ConceptMap) GetUseContext() []UsageContext {
	if o == nil || IsNil(o.UseContext) {
		var ret []UsageContext
		return ret
	}
	return o.UseContext
}

// GetUseContextOk returns a tuple with the UseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetUseContextOk() ([]UsageContext, bool) {
	if o == nil || IsNil(o.UseContext) {
		return nil, false
	}
	return o.UseContext, true
}

// HasUseContext returns a boolean if a field has been set.
func (o *ConceptMap) HasUseContext() bool {
	if o != nil && !IsNil(o.UseContext) {
		return true
	}

	return false
}

// SetUseContext gets a reference to the given []UsageContext and assigns it to the UseContext field.
func (o *ConceptMap) SetUseContext(v []UsageContext) {
	o.UseContext = v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *ConceptMap) GetJurisdiction() []CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret []CodeableConcept
		return ret
	}
	return o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetJurisdictionOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *ConceptMap) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given []CodeableConcept and assigns it to the Jurisdiction field.
func (o *ConceptMap) SetJurisdiction(v []CodeableConcept) {
	o.Jurisdiction = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *ConceptMap) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *ConceptMap) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *ConceptMap) SetPurpose(v string) {
	o.Purpose = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *ConceptMap) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *ConceptMap) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *ConceptMap) SetCopyright(v string) {
	o.Copyright = &v
}

// GetSourceUri returns the SourceUri field value if set, zero value otherwise.
func (o *ConceptMap) GetSourceUri() string {
	if o == nil || IsNil(o.SourceUri) {
		var ret string
		return ret
	}
	return *o.SourceUri
}

// GetSourceUriOk returns a tuple with the SourceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetSourceUriOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUri) {
		return nil, false
	}
	return o.SourceUri, true
}

// HasSourceUri returns a boolean if a field has been set.
func (o *ConceptMap) HasSourceUri() bool {
	if o != nil && !IsNil(o.SourceUri) {
		return true
	}

	return false
}

// SetSourceUri gets a reference to the given string and assigns it to the SourceUri field.
func (o *ConceptMap) SetSourceUri(v string) {
	o.SourceUri = &v
}

// GetSourceCanonical returns the SourceCanonical field value if set, zero value otherwise.
func (o *ConceptMap) GetSourceCanonical() string {
	if o == nil || IsNil(o.SourceCanonical) {
		var ret string
		return ret
	}
	return *o.SourceCanonical
}

// GetSourceCanonicalOk returns a tuple with the SourceCanonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetSourceCanonicalOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCanonical) {
		return nil, false
	}
	return o.SourceCanonical, true
}

// HasSourceCanonical returns a boolean if a field has been set.
func (o *ConceptMap) HasSourceCanonical() bool {
	if o != nil && !IsNil(o.SourceCanonical) {
		return true
	}

	return false
}

// SetSourceCanonical gets a reference to the given string and assigns it to the SourceCanonical field.
func (o *ConceptMap) SetSourceCanonical(v string) {
	o.SourceCanonical = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise.
func (o *ConceptMap) GetTargetUri() string {
	if o == nil || IsNil(o.TargetUri) {
		var ret string
		return ret
	}
	return *o.TargetUri
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetTargetUriOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUri) {
		return nil, false
	}
	return o.TargetUri, true
}

// HasTargetUri returns a boolean if a field has been set.
func (o *ConceptMap) HasTargetUri() bool {
	if o != nil && !IsNil(o.TargetUri) {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given string and assigns it to the TargetUri field.
func (o *ConceptMap) SetTargetUri(v string) {
	o.TargetUri = &v
}

// GetTargetCanonical returns the TargetCanonical field value if set, zero value otherwise.
func (o *ConceptMap) GetTargetCanonical() string {
	if o == nil || IsNil(o.TargetCanonical) {
		var ret string
		return ret
	}
	return *o.TargetCanonical
}

// GetTargetCanonicalOk returns a tuple with the TargetCanonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetTargetCanonicalOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCanonical) {
		return nil, false
	}
	return o.TargetCanonical, true
}

// HasTargetCanonical returns a boolean if a field has been set.
func (o *ConceptMap) HasTargetCanonical() bool {
	if o != nil && !IsNil(o.TargetCanonical) {
		return true
	}

	return false
}

// SetTargetCanonical gets a reference to the given string and assigns it to the TargetCanonical field.
func (o *ConceptMap) SetTargetCanonical(v string) {
	o.TargetCanonical = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ConceptMap) GetGroup() []ConceptMapGroup {
	if o == nil || IsNil(o.Group) {
		var ret []ConceptMapGroup
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMap) GetGroupOk() ([]ConceptMapGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ConceptMap) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []ConceptMapGroup and assigns it to the Group field.
func (o *ConceptMap) SetGroup(v []ConceptMapGroup) {
	o.Group = v
}

func (o ConceptMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConceptMap) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Experimental) {
		toSerialize["experimental"] = o.Experimental
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Publisher) {
		toSerialize["publisher"] = o.Publisher
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.UseContext) {
		toSerialize["useContext"] = o.UseContext
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.Copyright) {
		toSerialize["copyright"] = o.Copyright
	}
	if !IsNil(o.SourceUri) {
		toSerialize["sourceUri"] = o.SourceUri
	}
	if !IsNil(o.SourceCanonical) {
		toSerialize["sourceCanonical"] = o.SourceCanonical
	}
	if !IsNil(o.TargetUri) {
		toSerialize["targetUri"] = o.TargetUri
	}
	if !IsNil(o.TargetCanonical) {
		toSerialize["targetCanonical"] = o.TargetCanonical
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

func (o *ConceptMap) UnmarshalJSON(data []byte) (err error) {
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

	varConceptMap := _ConceptMap{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConceptMap)

	if err != nil {
		return err
	}

	*o = ConceptMap(varConceptMap)

	return err
}

type NullableConceptMap struct {
	value *ConceptMap
	isSet bool
}

func (v NullableConceptMap) Get() *ConceptMap {
	return v.value
}

func (v *NullableConceptMap) Set(val *ConceptMap) {
	v.value = val
	v.isSet = true
}

func (v NullableConceptMap) IsSet() bool {
	return v.isSet
}

func (v *NullableConceptMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConceptMap(val *ConceptMap) *NullableConceptMap {
	return &NullableConceptMap{value: val, isSet: true}
}

func (v NullableConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConceptMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


