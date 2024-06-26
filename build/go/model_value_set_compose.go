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

// checks if the ValueSetCompose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueSetCompose{}

// ValueSetCompose A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSetCompose struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A date or partial date (e.g. just year or year + month). There is no time zone. The format is a union of the schema types gYear, gYearMonth and date.  Dates SHALL be valid dates.
	LockedDate *string `json:"lockedDate,omitempty"`
	// Value of \"true\" or \"false\"
	Inactive *bool `json:"inactive,omitempty"`
	// Include one or more codes from a code system or other value set(s).
	Include []ValueSetInclude `json:"include"`
	// Exclude one or more codes from the value set based on code system filters and/or other value sets.
	Exclude []ValueSetInclude `json:"exclude,omitempty"`
}

type _ValueSetCompose ValueSetCompose

// NewValueSetCompose instantiates a new ValueSetCompose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueSetCompose(include []ValueSetInclude) *ValueSetCompose {
	this := ValueSetCompose{}
	this.Include = include
	return &this
}

// NewValueSetComposeWithDefaults instantiates a new ValueSetCompose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueSetComposeWithDefaults() *ValueSetCompose {
	this := ValueSetCompose{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ValueSetCompose) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ValueSetCompose) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ValueSetCompose) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ValueSetCompose) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ValueSetCompose) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ValueSetCompose) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ValueSetCompose) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ValueSetCompose) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ValueSetCompose) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetLockedDate returns the LockedDate field value if set, zero value otherwise.
func (o *ValueSetCompose) GetLockedDate() string {
	if o == nil || IsNil(o.LockedDate) {
		var ret string
		return ret
	}
	return *o.LockedDate
}

// GetLockedDateOk returns a tuple with the LockedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetLockedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LockedDate) {
		return nil, false
	}
	return o.LockedDate, true
}

// HasLockedDate returns a boolean if a field has been set.
func (o *ValueSetCompose) HasLockedDate() bool {
	if o != nil && !IsNil(o.LockedDate) {
		return true
	}

	return false
}

// SetLockedDate gets a reference to the given string and assigns it to the LockedDate field.
func (o *ValueSetCompose) SetLockedDate(v string) {
	o.LockedDate = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *ValueSetCompose) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *ValueSetCompose) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *ValueSetCompose) SetInactive(v bool) {
	o.Inactive = &v
}

// GetInclude returns the Include field value
func (o *ValueSetCompose) GetInclude() []ValueSetInclude {
	if o == nil {
		var ret []ValueSetInclude
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetIncludeOk() ([]ValueSetInclude, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *ValueSetCompose) SetInclude(v []ValueSetInclude) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ValueSetCompose) GetExclude() []ValueSetInclude {
	if o == nil || IsNil(o.Exclude) {
		var ret []ValueSetInclude
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueSetCompose) GetExcludeOk() ([]ValueSetInclude, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ValueSetCompose) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []ValueSetInclude and assigns it to the Exclude field.
func (o *ValueSetCompose) SetExclude(v []ValueSetInclude) {
	o.Exclude = v
}

func (o ValueSetCompose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueSetCompose) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LockedDate) {
		toSerialize["lockedDate"] = o.LockedDate
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	toSerialize["include"] = o.Include
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	return toSerialize, nil
}

func (o *ValueSetCompose) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include",
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

	varValueSetCompose := _ValueSetCompose{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValueSetCompose)

	if err != nil {
		return err
	}

	*o = ValueSetCompose(varValueSetCompose)

	return err
}

type NullableValueSetCompose struct {
	value *ValueSetCompose
	isSet bool
}

func (v NullableValueSetCompose) Get() *ValueSetCompose {
	return v.value
}

func (v *NullableValueSetCompose) Set(val *ValueSetCompose) {
	v.value = val
	v.isSet = true
}

func (v NullableValueSetCompose) IsSet() bool {
	return v.isSet
}

func (v *NullableValueSetCompose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueSetCompose(val *ValueSetCompose) *NullableValueSetCompose {
	return &NullableValueSetCompose{value: val, isSet: true}
}

func (v NullableValueSetCompose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueSetCompose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


