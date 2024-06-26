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

// checks if the TerminologyCapabilitiesVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminologyCapabilitiesVersion{}

// TerminologyCapabilitiesVersion A TerminologyCapabilities resource documents a set of capabilities (behaviors) of a FHIR Terminology Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type TerminologyCapabilitiesVersion struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Code *string `json:"code,omitempty"`
	// Value of \"true\" or \"false\"
	IsDefault *bool `json:"isDefault,omitempty"`
	// Value of \"true\" or \"false\"
	Compositional *bool `json:"compositional,omitempty"`
	// Language Displays supported.
	Language []string `json:"language,omitempty"`
	// Filter Properties supported.
	Filter []TerminologyCapabilitiesFilter `json:"filter,omitempty"`
	// Properties supported for $lookup.
	Property []string `json:"property,omitempty"`
}

// NewTerminologyCapabilitiesVersion instantiates a new TerminologyCapabilitiesVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminologyCapabilitiesVersion() *TerminologyCapabilitiesVersion {
	this := TerminologyCapabilitiesVersion{}
	return &this
}

// NewTerminologyCapabilitiesVersionWithDefaults instantiates a new TerminologyCapabilitiesVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminologyCapabilitiesVersionWithDefaults() *TerminologyCapabilitiesVersion {
	this := TerminologyCapabilitiesVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminologyCapabilitiesVersion) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TerminologyCapabilitiesVersion) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TerminologyCapabilitiesVersion) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TerminologyCapabilitiesVersion) SetCode(v string) {
	o.Code = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *TerminologyCapabilitiesVersion) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetCompositional returns the Compositional field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetCompositional() bool {
	if o == nil || IsNil(o.Compositional) {
		var ret bool
		return ret
	}
	return *o.Compositional
}

// GetCompositionalOk returns a tuple with the Compositional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetCompositionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Compositional) {
		return nil, false
	}
	return o.Compositional, true
}

// HasCompositional returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasCompositional() bool {
	if o != nil && !IsNil(o.Compositional) {
		return true
	}

	return false
}

// SetCompositional gets a reference to the given bool and assigns it to the Compositional field.
func (o *TerminologyCapabilitiesVersion) SetCompositional(v bool) {
	o.Compositional = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetLanguage() []string {
	if o == nil || IsNil(o.Language) {
		var ret []string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetLanguageOk() ([]string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []string and assigns it to the Language field.
func (o *TerminologyCapabilitiesVersion) SetLanguage(v []string) {
	o.Language = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetFilter() []TerminologyCapabilitiesFilter {
	if o == nil || IsNil(o.Filter) {
		var ret []TerminologyCapabilitiesFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetFilterOk() ([]TerminologyCapabilitiesFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []TerminologyCapabilitiesFilter and assigns it to the Filter field.
func (o *TerminologyCapabilitiesVersion) SetFilter(v []TerminologyCapabilitiesFilter) {
	o.Filter = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesVersion) GetProperty() []string {
	if o == nil || IsNil(o.Property) {
		var ret []string
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesVersion) GetPropertyOk() ([]string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesVersion) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given []string and assigns it to the Property field.
func (o *TerminologyCapabilitiesVersion) SetProperty(v []string) {
	o.Property = v
}

func (o TerminologyCapabilitiesVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminologyCapabilitiesVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.ModifierExtension) {
		toSerialize["modifierExtension"] = o.ModifierExtension
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.Compositional) {
		toSerialize["compositional"] = o.Compositional
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	return toSerialize, nil
}

type NullableTerminologyCapabilitiesVersion struct {
	value *TerminologyCapabilitiesVersion
	isSet bool
}

func (v NullableTerminologyCapabilitiesVersion) Get() *TerminologyCapabilitiesVersion {
	return v.value
}

func (v *NullableTerminologyCapabilitiesVersion) Set(val *TerminologyCapabilitiesVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminologyCapabilitiesVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminologyCapabilitiesVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminologyCapabilitiesVersion(val *TerminologyCapabilitiesVersion) *NullableTerminologyCapabilitiesVersion {
	return &NullableTerminologyCapabilitiesVersion{value: val, isSet: true}
}

func (v NullableTerminologyCapabilitiesVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminologyCapabilitiesVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

