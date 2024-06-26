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

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project Encapsulation of resources for a specific project or organization.
type Project struct {
	// This is a Project resource
	ResourceType string `json:"resourceType"`
	// A sequence of Unicode characters
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
	Contained []Resource `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An identifier for this project.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Value of \"true\" or \"false\"
	SuperAdmin *bool `json:"superAdmin,omitempty"`
	// Value of \"true\" or \"false\"
	StrictMode *bool `json:"strictMode,omitempty"`
	// Value of \"true\" or \"false\"
	CheckReferencesOnWrite *bool `json:"checkReferencesOnWrite,omitempty"`
	// The user who owns the project.
	Owner *Reference `json:"owner,omitempty"`
	// A list of optional features that are enabled for the project.
	Features []string `json:"features,omitempty"`
	// The default access policy for patients using open registration.
	DefaultPatientAccessPolicy *Reference `json:"defaultPatientAccessPolicy,omitempty"`
	// Option or parameter that can be adjusted within the Medplum Project to customize its behavior.
	Setting []ProjectSetting `json:"setting,omitempty"`
	// Option or parameter that can be adjusted within the Medplum Project to customize its behavior, only visible to project administrators.
	Secret []ProjectSetting `json:"secret,omitempty"`
	// Option or parameter that can be adjusted within the Medplum Project to customize its behavior, only modifiable by system administrators.
	SystemSetting []ProjectSetting `json:"systemSetting,omitempty"`
	// Option or parameter that can be adjusted within the Medplum Project to customize its behavior, only visible to system administrators.
	SystemSecret []ProjectSetting `json:"systemSecret,omitempty"`
	// Web application or web site that is associated with the project.
	Site []ProjectSite `json:"site,omitempty"`
	// Linked Projects whose contents are made available to this one
	Link []ProjectLink `json:"link,omitempty"`
}

type _Project Project

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(resourceType string) *Project {
	this := Project{}
	this.ResourceType = resourceType
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Project) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Project) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Project) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Project) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Project) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Project) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Project) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Project) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Project) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Project) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Project) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Project) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Project) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Project) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Project) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Project) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Project) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Project) GetContained() []Resource {
	if o == nil || IsNil(o.Contained) {
		var ret []Resource
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetContainedOk() ([]Resource, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Project) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []Resource and assigns it to the Contained field.
func (o *Project) SetContained(v []Resource) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Project) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Project) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Project) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Project) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Project) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Project) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Project) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Project) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Project) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Project) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Project) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description = &v
}

// GetSuperAdmin returns the SuperAdmin field value if set, zero value otherwise.
func (o *Project) GetSuperAdmin() bool {
	if o == nil || IsNil(o.SuperAdmin) {
		var ret bool
		return ret
	}
	return *o.SuperAdmin
}

// GetSuperAdminOk returns a tuple with the SuperAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSuperAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.SuperAdmin) {
		return nil, false
	}
	return o.SuperAdmin, true
}

// HasSuperAdmin returns a boolean if a field has been set.
func (o *Project) HasSuperAdmin() bool {
	if o != nil && !IsNil(o.SuperAdmin) {
		return true
	}

	return false
}

// SetSuperAdmin gets a reference to the given bool and assigns it to the SuperAdmin field.
func (o *Project) SetSuperAdmin(v bool) {
	o.SuperAdmin = &v
}

// GetStrictMode returns the StrictMode field value if set, zero value otherwise.
func (o *Project) GetStrictMode() bool {
	if o == nil || IsNil(o.StrictMode) {
		var ret bool
		return ret
	}
	return *o.StrictMode
}

// GetStrictModeOk returns a tuple with the StrictMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetStrictModeOk() (*bool, bool) {
	if o == nil || IsNil(o.StrictMode) {
		return nil, false
	}
	return o.StrictMode, true
}

// HasStrictMode returns a boolean if a field has been set.
func (o *Project) HasStrictMode() bool {
	if o != nil && !IsNil(o.StrictMode) {
		return true
	}

	return false
}

// SetStrictMode gets a reference to the given bool and assigns it to the StrictMode field.
func (o *Project) SetStrictMode(v bool) {
	o.StrictMode = &v
}

// GetCheckReferencesOnWrite returns the CheckReferencesOnWrite field value if set, zero value otherwise.
func (o *Project) GetCheckReferencesOnWrite() bool {
	if o == nil || IsNil(o.CheckReferencesOnWrite) {
		var ret bool
		return ret
	}
	return *o.CheckReferencesOnWrite
}

// GetCheckReferencesOnWriteOk returns a tuple with the CheckReferencesOnWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCheckReferencesOnWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckReferencesOnWrite) {
		return nil, false
	}
	return o.CheckReferencesOnWrite, true
}

// HasCheckReferencesOnWrite returns a boolean if a field has been set.
func (o *Project) HasCheckReferencesOnWrite() bool {
	if o != nil && !IsNil(o.CheckReferencesOnWrite) {
		return true
	}

	return false
}

// SetCheckReferencesOnWrite gets a reference to the given bool and assigns it to the CheckReferencesOnWrite field.
func (o *Project) SetCheckReferencesOnWrite(v bool) {
	o.CheckReferencesOnWrite = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Project) GetOwner() Reference {
	if o == nil || IsNil(o.Owner) {
		var ret Reference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetOwnerOk() (*Reference, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Project) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Reference and assigns it to the Owner field.
func (o *Project) SetOwner(v Reference) {
	o.Owner = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Project) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Project) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *Project) SetFeatures(v []string) {
	o.Features = v
}

// GetDefaultPatientAccessPolicy returns the DefaultPatientAccessPolicy field value if set, zero value otherwise.
func (o *Project) GetDefaultPatientAccessPolicy() Reference {
	if o == nil || IsNil(o.DefaultPatientAccessPolicy) {
		var ret Reference
		return ret
	}
	return *o.DefaultPatientAccessPolicy
}

// GetDefaultPatientAccessPolicyOk returns a tuple with the DefaultPatientAccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDefaultPatientAccessPolicyOk() (*Reference, bool) {
	if o == nil || IsNil(o.DefaultPatientAccessPolicy) {
		return nil, false
	}
	return o.DefaultPatientAccessPolicy, true
}

// HasDefaultPatientAccessPolicy returns a boolean if a field has been set.
func (o *Project) HasDefaultPatientAccessPolicy() bool {
	if o != nil && !IsNil(o.DefaultPatientAccessPolicy) {
		return true
	}

	return false
}

// SetDefaultPatientAccessPolicy gets a reference to the given Reference and assigns it to the DefaultPatientAccessPolicy field.
func (o *Project) SetDefaultPatientAccessPolicy(v Reference) {
	o.DefaultPatientAccessPolicy = &v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *Project) GetSetting() []ProjectSetting {
	if o == nil || IsNil(o.Setting) {
		var ret []ProjectSetting
		return ret
	}
	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSettingOk() ([]ProjectSetting, bool) {
	if o == nil || IsNil(o.Setting) {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *Project) HasSetting() bool {
	if o != nil && !IsNil(o.Setting) {
		return true
	}

	return false
}

// SetSetting gets a reference to the given []ProjectSetting and assigns it to the Setting field.
func (o *Project) SetSetting(v []ProjectSetting) {
	o.Setting = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Project) GetSecret() []ProjectSetting {
	if o == nil || IsNil(o.Secret) {
		var ret []ProjectSetting
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSecretOk() ([]ProjectSetting, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Project) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []ProjectSetting and assigns it to the Secret field.
func (o *Project) SetSecret(v []ProjectSetting) {
	o.Secret = v
}

// GetSystemSetting returns the SystemSetting field value if set, zero value otherwise.
func (o *Project) GetSystemSetting() []ProjectSetting {
	if o == nil || IsNil(o.SystemSetting) {
		var ret []ProjectSetting
		return ret
	}
	return o.SystemSetting
}

// GetSystemSettingOk returns a tuple with the SystemSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSystemSettingOk() ([]ProjectSetting, bool) {
	if o == nil || IsNil(o.SystemSetting) {
		return nil, false
	}
	return o.SystemSetting, true
}

// HasSystemSetting returns a boolean if a field has been set.
func (o *Project) HasSystemSetting() bool {
	if o != nil && !IsNil(o.SystemSetting) {
		return true
	}

	return false
}

// SetSystemSetting gets a reference to the given []ProjectSetting and assigns it to the SystemSetting field.
func (o *Project) SetSystemSetting(v []ProjectSetting) {
	o.SystemSetting = v
}

// GetSystemSecret returns the SystemSecret field value if set, zero value otherwise.
func (o *Project) GetSystemSecret() []ProjectSetting {
	if o == nil || IsNil(o.SystemSecret) {
		var ret []ProjectSetting
		return ret
	}
	return o.SystemSecret
}

// GetSystemSecretOk returns a tuple with the SystemSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSystemSecretOk() ([]ProjectSetting, bool) {
	if o == nil || IsNil(o.SystemSecret) {
		return nil, false
	}
	return o.SystemSecret, true
}

// HasSystemSecret returns a boolean if a field has been set.
func (o *Project) HasSystemSecret() bool {
	if o != nil && !IsNil(o.SystemSecret) {
		return true
	}

	return false
}

// SetSystemSecret gets a reference to the given []ProjectSetting and assigns it to the SystemSecret field.
func (o *Project) SetSystemSecret(v []ProjectSetting) {
	o.SystemSecret = v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *Project) GetSite() []ProjectSite {
	if o == nil || IsNil(o.Site) {
		var ret []ProjectSite
		return ret
	}
	return o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSiteOk() ([]ProjectSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Project) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given []ProjectSite and assigns it to the Site field.
func (o *Project) SetSite(v []ProjectSite) {
	o.Site = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Project) GetLink() []ProjectLink {
	if o == nil || IsNil(o.Link) {
		var ret []ProjectLink
		return ret
	}
	return o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetLinkOk() ([]ProjectLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Project) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given []ProjectLink and assigns it to the Link field.
func (o *Project) SetLink(v []ProjectLink) {
	o.Link = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SuperAdmin) {
		toSerialize["superAdmin"] = o.SuperAdmin
	}
	if !IsNil(o.StrictMode) {
		toSerialize["strictMode"] = o.StrictMode
	}
	if !IsNil(o.CheckReferencesOnWrite) {
		toSerialize["checkReferencesOnWrite"] = o.CheckReferencesOnWrite
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.DefaultPatientAccessPolicy) {
		toSerialize["defaultPatientAccessPolicy"] = o.DefaultPatientAccessPolicy
	}
	if !IsNil(o.Setting) {
		toSerialize["setting"] = o.Setting
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SystemSetting) {
		toSerialize["systemSetting"] = o.SystemSetting
	}
	if !IsNil(o.SystemSecret) {
		toSerialize["systemSecret"] = o.SystemSecret
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

func (o *Project) UnmarshalJSON(data []byte) (err error) {
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

	varProject := _Project{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProject)

	if err != nil {
		return err
	}

	*o = Project(varProject)

	return err
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

