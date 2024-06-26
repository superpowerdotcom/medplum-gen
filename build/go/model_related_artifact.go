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

// checks if the RelatedArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedArtifact{}

// RelatedArtifact Related artifacts such as additional documentation, justification, or bibliographic references.
type RelatedArtifact struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// The type of relationship to the related artifact.
	Type *string `json:"type,omitempty"`
	// A sequence of Unicode characters
	Label *string `json:"label,omitempty"`
	// A sequence of Unicode characters
	Display *string `json:"display,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Citation *string `json:"citation,omitempty"`
	// A URI that is a literal reference
	Url *string `json:"url,omitempty"`
	// The document being referenced, represented as an attachment. This is exclusive with the resource element.
	Document *Attachment `json:"document,omitempty"`
	// A URI that is a reference to a canonical URL on a FHIR resource
	Resource *string `json:"resource,omitempty"`
}

// NewRelatedArtifact instantiates a new RelatedArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedArtifact() *RelatedArtifact {
	this := RelatedArtifact{}
	return &this
}

// NewRelatedArtifactWithDefaults instantiates a new RelatedArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedArtifactWithDefaults() *RelatedArtifact {
	this := RelatedArtifact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelatedArtifact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelatedArtifact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelatedArtifact) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *RelatedArtifact) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *RelatedArtifact) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *RelatedArtifact) SetExtension(v []Extension) {
	o.Extension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelatedArtifact) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelatedArtifact) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RelatedArtifact) SetType(v string) {
	o.Type = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RelatedArtifact) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RelatedArtifact) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RelatedArtifact) SetLabel(v string) {
	o.Label = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *RelatedArtifact) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *RelatedArtifact) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *RelatedArtifact) SetDisplay(v string) {
	o.Display = &v
}

// GetCitation returns the Citation field value if set, zero value otherwise.
func (o *RelatedArtifact) GetCitation() string {
	if o == nil || IsNil(o.Citation) {
		var ret string
		return ret
	}
	return *o.Citation
}

// GetCitationOk returns a tuple with the Citation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetCitationOk() (*string, bool) {
	if o == nil || IsNil(o.Citation) {
		return nil, false
	}
	return o.Citation, true
}

// HasCitation returns a boolean if a field has been set.
func (o *RelatedArtifact) HasCitation() bool {
	if o != nil && !IsNil(o.Citation) {
		return true
	}

	return false
}

// SetCitation gets a reference to the given string and assigns it to the Citation field.
func (o *RelatedArtifact) SetCitation(v string) {
	o.Citation = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RelatedArtifact) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RelatedArtifact) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RelatedArtifact) SetUrl(v string) {
	o.Url = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *RelatedArtifact) GetDocument() Attachment {
	if o == nil || IsNil(o.Document) {
		var ret Attachment
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetDocumentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *RelatedArtifact) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Attachment and assigns it to the Document field.
func (o *RelatedArtifact) SetDocument(v Attachment) {
	o.Document = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *RelatedArtifact) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedArtifact) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *RelatedArtifact) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *RelatedArtifact) SetResource(v string) {
	o.Resource = &v
}

func (o RelatedArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Citation) {
		toSerialize["citation"] = o.Citation
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableRelatedArtifact struct {
	value *RelatedArtifact
	isSet bool
}

func (v NullableRelatedArtifact) Get() *RelatedArtifact {
	return v.value
}

func (v *NullableRelatedArtifact) Set(val *RelatedArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedArtifact(val *RelatedArtifact) *NullableRelatedArtifact {
	return &NullableRelatedArtifact{value: val, isSet: true}
}

func (v NullableRelatedArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


