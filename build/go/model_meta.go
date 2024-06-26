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
)

// checks if the Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Meta{}

// Meta The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	VersionId *string `json:"versionId,omitempty"`
	// An instant in time - known at least to the second
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// String of characters used to identify a name or a resource
	Source *string `json:"source,omitempty"`
	// A list of profiles (references to [[[StructureDefinition]]] resources) that this resource claims to conform to. The URL is a reference to [[[StructureDefinition.url]]].
	Profile []string `json:"profile,omitempty"`
	// Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Security []Coding `json:"security,omitempty"`
	// Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	Tag []Coding `json:"tag,omitempty"`
	// String of characters used to identify a name or a resource
	Project *string `json:"project,omitempty"`
	// The individual, device or organization who initiated the last change.
	Author *Reference `json:"author,omitempty"`
	// Optional account reference that can be used for sub-project compartments.
	Account *Reference `json:"account,omitempty"`
	// The list of compartments containing this resource. This is readonly and is set by the server.
	Compartment []Reference `json:"compartment,omitempty"`
}

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta() *Meta {
	this := Meta{}
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Meta) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Meta) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Meta) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Meta) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Meta) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Meta) SetExtension(v []Extension) {
	o.Extension = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *Meta) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *Meta) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *Meta) SetVersionId(v string) {
	o.VersionId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Meta) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Meta) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *Meta) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Meta) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Meta) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Meta) SetSource(v string) {
	o.Source = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Meta) GetProfile() []string {
	if o == nil || IsNil(o.Profile) {
		var ret []string
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetProfileOk() ([]string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Meta) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given []string and assigns it to the Profile field.
func (o *Meta) SetProfile(v []string) {
	o.Profile = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *Meta) GetSecurity() []Coding {
	if o == nil || IsNil(o.Security) {
		var ret []Coding
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetSecurityOk() ([]Coding, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *Meta) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []Coding and assigns it to the Security field.
func (o *Meta) SetSecurity(v []Coding) {
	o.Security = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Meta) GetTag() []Coding {
	if o == nil || IsNil(o.Tag) {
		var ret []Coding
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetTagOk() ([]Coding, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Meta) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Coding and assigns it to the Tag field.
func (o *Meta) SetTag(v []Coding) {
	o.Tag = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Meta) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Meta) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Meta) SetProject(v string) {
	o.Project = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Meta) GetAuthor() Reference {
	if o == nil || IsNil(o.Author) {
		var ret Reference
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetAuthorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Meta) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Reference and assigns it to the Author field.
func (o *Meta) SetAuthor(v Reference) {
	o.Author = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Meta) GetAccount() Reference {
	if o == nil || IsNil(o.Account) {
		var ret Reference
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetAccountOk() (*Reference, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Meta) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Reference and assigns it to the Account field.
func (o *Meta) SetAccount(v Reference) {
	o.Account = &v
}

// GetCompartment returns the Compartment field value if set, zero value otherwise.
func (o *Meta) GetCompartment() []Reference {
	if o == nil || IsNil(o.Compartment) {
		var ret []Reference
		return ret
	}
	return o.Compartment
}

// GetCompartmentOk returns a tuple with the Compartment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetCompartmentOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Compartment) {
		return nil, false
	}
	return o.Compartment, true
}

// HasCompartment returns a boolean if a field has been set.
func (o *Meta) HasCompartment() bool {
	if o != nil && !IsNil(o.Compartment) {
		return true
	}

	return false
}

// SetCompartment gets a reference to the given []Reference and assigns it to the Compartment field.
func (o *Meta) SetCompartment(v []Reference) {
	o.Compartment = v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Compartment) {
		toSerialize["compartment"] = o.Compartment
	}
	return toSerialize, nil
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


