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

// checks if the TerminologyCapabilitiesSoftware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminologyCapabilitiesSoftware{}

// TerminologyCapabilitiesSoftware A TerminologyCapabilities resource documents a set of capabilities (behaviors) of a FHIR Terminology Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type TerminologyCapabilitiesSoftware struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A sequence of Unicode characters
	Version *string `json:"version,omitempty"`
}

// NewTerminologyCapabilitiesSoftware instantiates a new TerminologyCapabilitiesSoftware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminologyCapabilitiesSoftware() *TerminologyCapabilitiesSoftware {
	this := TerminologyCapabilitiesSoftware{}
	return &this
}

// NewTerminologyCapabilitiesSoftwareWithDefaults instantiates a new TerminologyCapabilitiesSoftware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminologyCapabilitiesSoftwareWithDefaults() *TerminologyCapabilitiesSoftware {
	this := TerminologyCapabilitiesSoftware{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesSoftware) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesSoftware) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesSoftware) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminologyCapabilitiesSoftware) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesSoftware) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesSoftware) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesSoftware) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TerminologyCapabilitiesSoftware) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesSoftware) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesSoftware) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesSoftware) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TerminologyCapabilitiesSoftware) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesSoftware) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesSoftware) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesSoftware) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TerminologyCapabilitiesSoftware) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TerminologyCapabilitiesSoftware) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminologyCapabilitiesSoftware) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TerminologyCapabilitiesSoftware) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TerminologyCapabilitiesSoftware) SetVersion(v string) {
	o.Version = &v
}

func (o TerminologyCapabilitiesSoftware) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminologyCapabilitiesSoftware) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableTerminologyCapabilitiesSoftware struct {
	value *TerminologyCapabilitiesSoftware
	isSet bool
}

func (v NullableTerminologyCapabilitiesSoftware) Get() *TerminologyCapabilitiesSoftware {
	return v.value
}

func (v *NullableTerminologyCapabilitiesSoftware) Set(val *TerminologyCapabilitiesSoftware) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminologyCapabilitiesSoftware) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminologyCapabilitiesSoftware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminologyCapabilitiesSoftware(val *TerminologyCapabilitiesSoftware) *NullableTerminologyCapabilitiesSoftware {
	return &NullableTerminologyCapabilitiesSoftware{value: val, isSet: true}
}

func (v NullableTerminologyCapabilitiesSoftware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminologyCapabilitiesSoftware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


