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

// checks if the ProjectMembership type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMembership{}

// ProjectMembership Medplum project membership. A project membership grants a user access to a project.
type ProjectMembership struct {
	// This is a ProjectMembership resource
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
	// Project where the memberships are available.
	Project Reference `json:"project"`
	// The project administrator who invited the user to the project.
	InvitedBy *Reference `json:"invitedBy,omitempty"`
	// User that is granted access to the project.
	User Reference `json:"user"`
	// Reference to the resource that represents the user profile within the project.
	Profile Reference `json:"profile"`
	// A sequence of Unicode characters
	UserName *string `json:"userName,omitempty"`
	// A sequence of Unicode characters
	ExternalId *string `json:"externalId,omitempty"`
	// The access policy for the user within the project memebership.
	AccessPolicy *Reference `json:"accessPolicy,omitempty"`
	// Extended access configuration using parameterized access policies.
	Access []ProjectMembershipAccess `json:"access,omitempty"`
	// The user configuration for the user within the project memebership such as menu links, saved searches, and features.
	UserConfiguration *Reference `json:"userConfiguration,omitempty"`
	// Value of \"true\" or \"false\"
	Admin *bool `json:"admin,omitempty"`
}

type _ProjectMembership ProjectMembership

// NewProjectMembership instantiates a new ProjectMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMembership(resourceType string, project Reference, user Reference, profile Reference) *ProjectMembership {
	this := ProjectMembership{}
	this.ResourceType = resourceType
	this.Project = project
	this.User = user
	this.Profile = profile
	return &this
}

// NewProjectMembershipWithDefaults instantiates a new ProjectMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMembershipWithDefaults() *ProjectMembership {
	this := ProjectMembership{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *ProjectMembership) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ProjectMembership) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectMembership) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectMembership) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectMembership) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProjectMembership) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProjectMembership) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *ProjectMembership) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *ProjectMembership) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *ProjectMembership) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *ProjectMembership) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ProjectMembership) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ProjectMembership) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ProjectMembership) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ProjectMembership) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ProjectMembership) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *ProjectMembership) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *ProjectMembership) GetContained() []Resource {
	if o == nil || IsNil(o.Contained) {
		var ret []Resource
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetContainedOk() ([]Resource, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *ProjectMembership) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []Resource and assigns it to the Contained field.
func (o *ProjectMembership) SetContained(v []Resource) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ProjectMembership) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ProjectMembership) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ProjectMembership) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ProjectMembership) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ProjectMembership) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ProjectMembership) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetProject returns the Project field value
func (o *ProjectMembership) GetProject() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetProjectOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectMembership) SetProject(v Reference) {
	o.Project = v
}

// GetInvitedBy returns the InvitedBy field value if set, zero value otherwise.
func (o *ProjectMembership) GetInvitedBy() Reference {
	if o == nil || IsNil(o.InvitedBy) {
		var ret Reference
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetInvitedByOk() (*Reference, bool) {
	if o == nil || IsNil(o.InvitedBy) {
		return nil, false
	}
	return o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *ProjectMembership) HasInvitedBy() bool {
	if o != nil && !IsNil(o.InvitedBy) {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given Reference and assigns it to the InvitedBy field.
func (o *ProjectMembership) SetInvitedBy(v Reference) {
	o.InvitedBy = &v
}

// GetUser returns the User field value
func (o *ProjectMembership) GetUser() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetUserOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ProjectMembership) SetUser(v Reference) {
	o.User = v
}

// GetProfile returns the Profile field value
func (o *ProjectMembership) GetProfile() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetProfileOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ProjectMembership) SetProfile(v Reference) {
	o.Profile = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ProjectMembership) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ProjectMembership) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ProjectMembership) SetUserName(v string) {
	o.UserName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ProjectMembership) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ProjectMembership) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ProjectMembership) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *ProjectMembership) GetAccessPolicy() Reference {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret Reference
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetAccessPolicyOk() (*Reference, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *ProjectMembership) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given Reference and assigns it to the AccessPolicy field.
func (o *ProjectMembership) SetAccessPolicy(v Reference) {
	o.AccessPolicy = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ProjectMembership) GetAccess() []ProjectMembershipAccess {
	if o == nil || IsNil(o.Access) {
		var ret []ProjectMembershipAccess
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetAccessOk() ([]ProjectMembershipAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ProjectMembership) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []ProjectMembershipAccess and assigns it to the Access field.
func (o *ProjectMembership) SetAccess(v []ProjectMembershipAccess) {
	o.Access = v
}

// GetUserConfiguration returns the UserConfiguration field value if set, zero value otherwise.
func (o *ProjectMembership) GetUserConfiguration() Reference {
	if o == nil || IsNil(o.UserConfiguration) {
		var ret Reference
		return ret
	}
	return *o.UserConfiguration
}

// GetUserConfigurationOk returns a tuple with the UserConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetUserConfigurationOk() (*Reference, bool) {
	if o == nil || IsNil(o.UserConfiguration) {
		return nil, false
	}
	return o.UserConfiguration, true
}

// HasUserConfiguration returns a boolean if a field has been set.
func (o *ProjectMembership) HasUserConfiguration() bool {
	if o != nil && !IsNil(o.UserConfiguration) {
		return true
	}

	return false
}

// SetUserConfiguration gets a reference to the given Reference and assigns it to the UserConfiguration field.
func (o *ProjectMembership) SetUserConfiguration(v Reference) {
	o.UserConfiguration = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ProjectMembership) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMembership) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ProjectMembership) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *ProjectMembership) SetAdmin(v bool) {
	o.Admin = &v
}

func (o ProjectMembership) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMembership) ToMap() (map[string]interface{}, error) {
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
	toSerialize["project"] = o.Project
	if !IsNil(o.InvitedBy) {
		toSerialize["invitedBy"] = o.InvitedBy
	}
	toSerialize["user"] = o.User
	toSerialize["profile"] = o.Profile
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.AccessPolicy) {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.UserConfiguration) {
		toSerialize["userConfiguration"] = o.UserConfiguration
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

func (o *ProjectMembership) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"project",
		"user",
		"profile",
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

	varProjectMembership := _ProjectMembership{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMembership)

	if err != nil {
		return err
	}

	*o = ProjectMembership(varProjectMembership)

	return err
}

type NullableProjectMembership struct {
	value *ProjectMembership
	isSet bool
}

func (v NullableProjectMembership) Get() *ProjectMembership {
	return v.value
}

func (v *NullableProjectMembership) Set(val *ProjectMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMembership(val *ProjectMembership) *NullableProjectMembership {
	return &NullableProjectMembership{value: val, isSet: true}
}

func (v NullableProjectMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

