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

// checks if the NamingSystemUniqueId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamingSystemUniqueId{}

// NamingSystemUniqueId A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a \"System\" used within the Identifier and Coding data types.
type NamingSystemUniqueId struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the unique identifier scheme used for this particular identifier.
	Type *string `json:"type,omitempty"`
	// A sequence of Unicode characters
	Value *string `json:"value,omitempty"`
	// Value of \"true\" or \"false\"
	Preferred *bool `json:"preferred,omitempty"`
	// A sequence of Unicode characters
	Comment *string `json:"comment,omitempty"`
	// Identifies the period of time over which this identifier is considered appropriate to refer to the naming system.  Outside of this window, the identifier might be non-deterministic.
	Period *Period `json:"period,omitempty"`
}

// NewNamingSystemUniqueId instantiates a new NamingSystemUniqueId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingSystemUniqueId() *NamingSystemUniqueId {
	this := NamingSystemUniqueId{}
	return &this
}

// NewNamingSystemUniqueIdWithDefaults instantiates a new NamingSystemUniqueId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingSystemUniqueIdWithDefaults() *NamingSystemUniqueId {
	this := NamingSystemUniqueId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NamingSystemUniqueId) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *NamingSystemUniqueId) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *NamingSystemUniqueId) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NamingSystemUniqueId) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NamingSystemUniqueId) SetValue(v string) {
	o.Value = &v
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetPreferred() bool {
	if o == nil || IsNil(o.Preferred) {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasPreferred() bool {
	if o != nil && !IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *NamingSystemUniqueId) SetPreferred(v bool) {
	o.Preferred = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NamingSystemUniqueId) SetComment(v string) {
	o.Comment = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *NamingSystemUniqueId) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingSystemUniqueId) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *NamingSystemUniqueId) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *NamingSystemUniqueId) SetPeriod(v Period) {
	o.Period = &v
}

func (o NamingSystemUniqueId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamingSystemUniqueId) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return toSerialize, nil
}

type NullableNamingSystemUniqueId struct {
	value *NamingSystemUniqueId
	isSet bool
}

func (v NullableNamingSystemUniqueId) Get() *NamingSystemUniqueId {
	return v.value
}

func (v *NullableNamingSystemUniqueId) Set(val *NamingSystemUniqueId) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingSystemUniqueId) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingSystemUniqueId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingSystemUniqueId(val *NamingSystemUniqueId) *NullableNamingSystemUniqueId {
	return &NullableNamingSystemUniqueId{value: val, isSet: true}
}

func (v NullableNamingSystemUniqueId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingSystemUniqueId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


