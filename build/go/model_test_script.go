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

// checks if the TestScript type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestScript{}

// TestScript A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScript struct {
	// This is a TestScript resource
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
	// A formal identifier that is used to identify this test script when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier,omitempty"`
	// A sequence of Unicode characters
	Version *string `json:"version,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A sequence of Unicode characters
	Title *string `json:"title,omitempty"`
	// The status of this test script. Enables tracking the life-cycle of the content.
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
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate test script instances.
	UseContext []UsageContext `json:"useContext,omitempty"`
	// A legal or geographic region in which the test script is intended to be used.
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Purpose *string `json:"purpose,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Copyright *string `json:"copyright,omitempty"`
	// An abstract server used in operations within this test script in the origin element.
	Origin []TestScriptOrigin `json:"origin,omitempty"`
	// An abstract server used in operations within this test script in the destination element.
	Destination []TestScriptDestination `json:"destination,omitempty"`
	// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
	Metadata *TestScriptMetadata `json:"metadata,omitempty"`
	// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
	Fixture []TestScriptFixture `json:"fixture,omitempty"`
	// Reference to the profile to be used for validation.
	Profile []Reference `json:"profile,omitempty"`
	// Variable is set based either on element value in response body or on header field value in the response headers.
	Variable []TestScriptVariable `json:"variable,omitempty"`
	// A series of required setup operations before tests are executed.
	Setup *TestScriptSetup `json:"setup,omitempty"`
	// A test in this script.
	Test []TestScriptTest `json:"test,omitempty"`
	// A series of operations required to clean up after all the tests are executed (successfully or otherwise).
	Teardown *TestScriptTeardown `json:"teardown,omitempty"`
}

type _TestScript TestScript

// NewTestScript instantiates a new TestScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestScript(resourceType string) *TestScript {
	this := TestScript{}
	this.ResourceType = resourceType
	return &this
}

// NewTestScriptWithDefaults instantiates a new TestScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestScriptWithDefaults() *TestScript {
	this := TestScript{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *TestScript) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *TestScript) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *TestScript) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestScript) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestScript) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestScript) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TestScript) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TestScript) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *TestScript) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *TestScript) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *TestScript) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *TestScript) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TestScript) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TestScript) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *TestScript) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TestScript) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TestScript) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *TestScript) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *TestScript) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *TestScript) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *TestScript) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TestScript) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TestScript) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TestScript) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TestScript) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TestScript) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TestScript) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TestScript) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TestScript) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TestScript) SetUrl(v string) {
	o.Url = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *TestScript) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *TestScript) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *TestScript) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TestScript) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TestScript) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TestScript) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestScript) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestScript) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestScript) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TestScript) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TestScript) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TestScript) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TestScript) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TestScript) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TestScript) SetStatus(v string) {
	o.Status = &v
}

// GetExperimental returns the Experimental field value if set, zero value otherwise.
func (o *TestScript) GetExperimental() bool {
	if o == nil || IsNil(o.Experimental) {
		var ret bool
		return ret
	}
	return *o.Experimental
}

// GetExperimentalOk returns a tuple with the Experimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetExperimentalOk() (*bool, bool) {
	if o == nil || IsNil(o.Experimental) {
		return nil, false
	}
	return o.Experimental, true
}

// HasExperimental returns a boolean if a field has been set.
func (o *TestScript) HasExperimental() bool {
	if o != nil && !IsNil(o.Experimental) {
		return true
	}

	return false
}

// SetExperimental gets a reference to the given bool and assigns it to the Experimental field.
func (o *TestScript) SetExperimental(v bool) {
	o.Experimental = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TestScript) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TestScript) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *TestScript) SetDate(v string) {
	o.Date = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *TestScript) GetPublisher() string {
	if o == nil || IsNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetPublisherOk() (*string, bool) {
	if o == nil || IsNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *TestScript) HasPublisher() bool {
	if o != nil && !IsNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *TestScript) SetPublisher(v string) {
	o.Publisher = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *TestScript) GetContact() []ContactDetail {
	if o == nil || IsNil(o.Contact) {
		var ret []ContactDetail
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetContactOk() ([]ContactDetail, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *TestScript) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given []ContactDetail and assigns it to the Contact field.
func (o *TestScript) SetContact(v []ContactDetail) {
	o.Contact = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TestScript) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TestScript) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TestScript) SetDescription(v string) {
	o.Description = &v
}

// GetUseContext returns the UseContext field value if set, zero value otherwise.
func (o *TestScript) GetUseContext() []UsageContext {
	if o == nil || IsNil(o.UseContext) {
		var ret []UsageContext
		return ret
	}
	return o.UseContext
}

// GetUseContextOk returns a tuple with the UseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetUseContextOk() ([]UsageContext, bool) {
	if o == nil || IsNil(o.UseContext) {
		return nil, false
	}
	return o.UseContext, true
}

// HasUseContext returns a boolean if a field has been set.
func (o *TestScript) HasUseContext() bool {
	if o != nil && !IsNil(o.UseContext) {
		return true
	}

	return false
}

// SetUseContext gets a reference to the given []UsageContext and assigns it to the UseContext field.
func (o *TestScript) SetUseContext(v []UsageContext) {
	o.UseContext = v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *TestScript) GetJurisdiction() []CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret []CodeableConcept
		return ret
	}
	return o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetJurisdictionOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *TestScript) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given []CodeableConcept and assigns it to the Jurisdiction field.
func (o *TestScript) SetJurisdiction(v []CodeableConcept) {
	o.Jurisdiction = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *TestScript) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *TestScript) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *TestScript) SetPurpose(v string) {
	o.Purpose = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *TestScript) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *TestScript) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *TestScript) SetCopyright(v string) {
	o.Copyright = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *TestScript) GetOrigin() []TestScriptOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret []TestScriptOrigin
		return ret
	}
	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetOriginOk() ([]TestScriptOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *TestScript) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given []TestScriptOrigin and assigns it to the Origin field.
func (o *TestScript) SetOrigin(v []TestScriptOrigin) {
	o.Origin = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TestScript) GetDestination() []TestScriptDestination {
	if o == nil || IsNil(o.Destination) {
		var ret []TestScriptDestination
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetDestinationOk() ([]TestScriptDestination, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TestScript) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given []TestScriptDestination and assigns it to the Destination field.
func (o *TestScript) SetDestination(v []TestScriptDestination) {
	o.Destination = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TestScript) GetMetadata() TestScriptMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret TestScriptMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetMetadataOk() (*TestScriptMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TestScript) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given TestScriptMetadata and assigns it to the Metadata field.
func (o *TestScript) SetMetadata(v TestScriptMetadata) {
	o.Metadata = &v
}

// GetFixture returns the Fixture field value if set, zero value otherwise.
func (o *TestScript) GetFixture() []TestScriptFixture {
	if o == nil || IsNil(o.Fixture) {
		var ret []TestScriptFixture
		return ret
	}
	return o.Fixture
}

// GetFixtureOk returns a tuple with the Fixture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetFixtureOk() ([]TestScriptFixture, bool) {
	if o == nil || IsNil(o.Fixture) {
		return nil, false
	}
	return o.Fixture, true
}

// HasFixture returns a boolean if a field has been set.
func (o *TestScript) HasFixture() bool {
	if o != nil && !IsNil(o.Fixture) {
		return true
	}

	return false
}

// SetFixture gets a reference to the given []TestScriptFixture and assigns it to the Fixture field.
func (o *TestScript) SetFixture(v []TestScriptFixture) {
	o.Fixture = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *TestScript) GetProfile() []Reference {
	if o == nil || IsNil(o.Profile) {
		var ret []Reference
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetProfileOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *TestScript) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given []Reference and assigns it to the Profile field.
func (o *TestScript) SetProfile(v []Reference) {
	o.Profile = v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *TestScript) GetVariable() []TestScriptVariable {
	if o == nil || IsNil(o.Variable) {
		var ret []TestScriptVariable
		return ret
	}
	return o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetVariableOk() ([]TestScriptVariable, bool) {
	if o == nil || IsNil(o.Variable) {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *TestScript) HasVariable() bool {
	if o != nil && !IsNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given []TestScriptVariable and assigns it to the Variable field.
func (o *TestScript) SetVariable(v []TestScriptVariable) {
	o.Variable = v
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *TestScript) GetSetup() TestScriptSetup {
	if o == nil || IsNil(o.Setup) {
		var ret TestScriptSetup
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetSetupOk() (*TestScriptSetup, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *TestScript) HasSetup() bool {
	if o != nil && !IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given TestScriptSetup and assigns it to the Setup field.
func (o *TestScript) SetSetup(v TestScriptSetup) {
	o.Setup = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *TestScript) GetTest() []TestScriptTest {
	if o == nil || IsNil(o.Test) {
		var ret []TestScriptTest
		return ret
	}
	return o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetTestOk() ([]TestScriptTest, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *TestScript) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given []TestScriptTest and assigns it to the Test field.
func (o *TestScript) SetTest(v []TestScriptTest) {
	o.Test = v
}

// GetTeardown returns the Teardown field value if set, zero value otherwise.
func (o *TestScript) GetTeardown() TestScriptTeardown {
	if o == nil || IsNil(o.Teardown) {
		var ret TestScriptTeardown
		return ret
	}
	return *o.Teardown
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScript) GetTeardownOk() (*TestScriptTeardown, bool) {
	if o == nil || IsNil(o.Teardown) {
		return nil, false
	}
	return o.Teardown, true
}

// HasTeardown returns a boolean if a field has been set.
func (o *TestScript) HasTeardown() bool {
	if o != nil && !IsNil(o.Teardown) {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given TestScriptTeardown and assigns it to the Teardown field.
func (o *TestScript) SetTeardown(v TestScriptTeardown) {
	o.Teardown = &v
}

func (o TestScript) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestScript) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Fixture) {
		toSerialize["fixture"] = o.Fixture
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	if !IsNil(o.Setup) {
		toSerialize["setup"] = o.Setup
	}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	if !IsNil(o.Teardown) {
		toSerialize["teardown"] = o.Teardown
	}
	return toSerialize, nil
}

func (o *TestScript) UnmarshalJSON(data []byte) (err error) {
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

	varTestScript := _TestScript{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestScript)

	if err != nil {
		return err
	}

	*o = TestScript(varTestScript)

	return err
}

type NullableTestScript struct {
	value *TestScript
	isSet bool
}

func (v NullableTestScript) Get() *TestScript {
	return v.value
}

func (v *NullableTestScript) Set(val *TestScript) {
	v.value = val
	v.isSet = true
}

func (v NullableTestScript) IsSet() bool {
	return v.isSet
}

func (v *NullableTestScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestScript(val *TestScript) *NullableTestScript {
	return &NullableTestScript{value: val, isSet: true}
}

func (v NullableTestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


