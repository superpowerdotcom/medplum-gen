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

// checks if the ElementDefinitionDiscriminator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ElementDefinitionDiscriminator{}

// ElementDefinitionDiscriminator Captures constraints on each element within the resource, profile, or extension.
type ElementDefinitionDiscriminator struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How the element value is interpreted when discrimination is evaluated.
	Type *string `json:"type,omitempty"`
	// A sequence of Unicode characters
	Path *string `json:"path,omitempty"`
}

// NewElementDefinitionDiscriminator instantiates a new ElementDefinitionDiscriminator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElementDefinitionDiscriminator() *ElementDefinitionDiscriminator {
	this := ElementDefinitionDiscriminator{}
	return &this
}

// NewElementDefinitionDiscriminatorWithDefaults instantiates a new ElementDefinitionDiscriminator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElementDefinitionDiscriminatorWithDefaults() *ElementDefinitionDiscriminator {
	this := ElementDefinitionDiscriminator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ElementDefinitionDiscriminator) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionDiscriminator) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ElementDefinitionDiscriminator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ElementDefinitionDiscriminator) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ElementDefinitionDiscriminator) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionDiscriminator) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ElementDefinitionDiscriminator) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ElementDefinitionDiscriminator) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ElementDefinitionDiscriminator) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionDiscriminator) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ElementDefinitionDiscriminator) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ElementDefinitionDiscriminator) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ElementDefinitionDiscriminator) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionDiscriminator) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ElementDefinitionDiscriminator) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ElementDefinitionDiscriminator) SetType(v string) {
	o.Type = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ElementDefinitionDiscriminator) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionDiscriminator) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ElementDefinitionDiscriminator) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ElementDefinitionDiscriminator) SetPath(v string) {
	o.Path = &v
}

func (o ElementDefinitionDiscriminator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ElementDefinitionDiscriminator) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableElementDefinitionDiscriminator struct {
	value *ElementDefinitionDiscriminator
	isSet bool
}

func (v NullableElementDefinitionDiscriminator) Get() *ElementDefinitionDiscriminator {
	return v.value
}

func (v *NullableElementDefinitionDiscriminator) Set(val *ElementDefinitionDiscriminator) {
	v.value = val
	v.isSet = true
}

func (v NullableElementDefinitionDiscriminator) IsSet() bool {
	return v.isSet
}

func (v *NullableElementDefinitionDiscriminator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElementDefinitionDiscriminator(val *ElementDefinitionDiscriminator) *NullableElementDefinitionDiscriminator {
	return &NullableElementDefinitionDiscriminator{value: val, isSet: true}
}

func (v NullableElementDefinitionDiscriminator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElementDefinitionDiscriminator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


