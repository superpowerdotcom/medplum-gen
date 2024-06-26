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

// checks if the ValueSetInclude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueSetInclude{}

// ValueSetInclude A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSetInclude struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// String of characters used to identify a name or a resource
	System *string `json:"system,omitempty"`
	// A sequence of Unicode characters
	Version *string `json:"version,omitempty"`
	// Specifies a concept to be included or excluded.
	Concept []ValueSetConcept `json:"concept,omitempty"`
	// Select concepts by specify a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system. If multiple filters are specified, they SHALL all be true.
	Filter []ValueSetFilter `json:"filter,omitempty"`
	// Selects the concepts found in this value set (based on its value set definition). This is an absolute URI that is a reference to ValueSet.url.  If multiple value sets are specified this includes the union of the contents of all of the referenced value sets.
	ValueSet []string `json:"valueSet,omitempty"`
}

// NewValueSetInclude instantiates a new ValueSetInclude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueSetInclude() *ValueSetInclude {
	this := ValueSetInclude{}
	return &this
}

// NewValueSetIncludeWithDefaults instantiates a new ValueSetInclude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueSetIncludeWithDefaults() *ValueSetInclude {
	this := ValueSetInclude{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ValueSetInclude) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ValueSetInclude) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ValueSetInclude) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ValueSetInclude) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ValueSetInclude) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ValueSetInclude) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ValueSetInclude) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ValueSetInclude) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ValueSetInclude) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ValueSetInclude) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ValueSetInclude) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *ValueSetInclude) SetSystem(v string) {
	o.System = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ValueSetInclude) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ValueSetInclude) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ValueSetInclude) SetVersion(v string) {
	o.Version = &v
}

// GetConcept returns the Concept field value if set, zero value otherwise.
func (o *ValueSetInclude) GetConcept() []ValueSetConcept {
	if o == nil || IsNil(o.Concept) {
		var ret []ValueSetConcept
		return ret
	}
	return o.Concept
}

// GetConceptOk returns a tuple with the Concept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetConceptOk() ([]ValueSetConcept, bool) {
	if o == nil || IsNil(o.Concept) {
		return nil, false
	}
	return o.Concept, true
}

// HasConcept returns a boolean if a field has been set.
func (o *ValueSetInclude) HasConcept() bool {
	if o != nil && !IsNil(o.Concept) {
		return true
	}

	return false
}

// SetConcept gets a reference to the given []ValueSetConcept and assigns it to the Concept field.
func (o *ValueSetInclude) SetConcept(v []ValueSetConcept) {
	o.Concept = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ValueSetInclude) GetFilter() []ValueSetFilter {
	if o == nil || IsNil(o.Filter) {
		var ret []ValueSetFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetFilterOk() ([]ValueSetFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ValueSetInclude) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []ValueSetFilter and assigns it to the Filter field.
func (o *ValueSetInclude) SetFilter(v []ValueSetFilter) {
	o.Filter = v
}

// GetValueSet returns the ValueSet field value if set, zero value otherwise.
func (o *ValueSetInclude) GetValueSet() []string {
	if o == nil || IsNil(o.ValueSet) {
		var ret []string
		return ret
	}
	return o.ValueSet
}

// GetValueSetOk returns a tuple with the ValueSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetInclude) GetValueSetOk() ([]string, bool) {
	if o == nil || IsNil(o.ValueSet) {
		return nil, false
	}
	return o.ValueSet, true
}

// HasValueSet returns a boolean if a field has been set.
func (o *ValueSetInclude) HasValueSet() bool {
	if o != nil && !IsNil(o.ValueSet) {
		return true
	}

	return false
}

// SetValueSet gets a reference to the given []string and assigns it to the ValueSet field.
func (o *ValueSetInclude) SetValueSet(v []string) {
	o.ValueSet = v
}

func (o ValueSetInclude) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueSetInclude) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Concept) {
		toSerialize["concept"] = o.Concept
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.ValueSet) {
		toSerialize["valueSet"] = o.ValueSet
	}
	return toSerialize, nil
}

type NullableValueSetInclude struct {
	value *ValueSetInclude
	isSet bool
}

func (v NullableValueSetInclude) Get() *ValueSetInclude {
	return v.value
}

func (v *NullableValueSetInclude) Set(val *ValueSetInclude) {
	v.value = val
	v.isSet = true
}

func (v NullableValueSetInclude) IsSet() bool {
	return v.isSet
}

func (v *NullableValueSetInclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueSetInclude(val *ValueSetInclude) *NullableValueSetInclude {
	return &NullableValueSetInclude{value: val, isSet: true}
}

func (v NullableValueSetInclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueSetInclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


