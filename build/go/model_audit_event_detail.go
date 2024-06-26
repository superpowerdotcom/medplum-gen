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

// checks if the AuditEventDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEventDetail{}

// AuditEventDetail A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEventDetail struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Type *string `json:"type,omitempty"`
	// The  value of the extra detail.
	ValueString *string `json:"valueString,omitempty"`
	// The  value of the extra detail.
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`
}

// NewAuditEventDetail instantiates a new AuditEventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEventDetail() *AuditEventDetail {
	this := AuditEventDetail{}
	return &this
}

// NewAuditEventDetailWithDefaults instantiates a new AuditEventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEventDetailWithDefaults() *AuditEventDetail {
	this := AuditEventDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditEventDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditEventDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditEventDetail) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AuditEventDetail) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AuditEventDetail) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *AuditEventDetail) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *AuditEventDetail) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *AuditEventDetail) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *AuditEventDetail) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditEventDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditEventDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditEventDetail) SetType(v string) {
	o.Type = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *AuditEventDetail) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *AuditEventDetail) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *AuditEventDetail) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueBase64Binary returns the ValueBase64Binary field value if set, zero value otherwise.
func (o *AuditEventDetail) GetValueBase64Binary() string {
	if o == nil || IsNil(o.ValueBase64Binary) {
		var ret string
		return ret
	}
	return *o.ValueBase64Binary
}

// GetValueBase64BinaryOk returns a tuple with the ValueBase64Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDetail) GetValueBase64BinaryOk() (*string, bool) {
	if o == nil || IsNil(o.ValueBase64Binary) {
		return nil, false
	}
	return o.ValueBase64Binary, true
}

// HasValueBase64Binary returns a boolean if a field has been set.
func (o *AuditEventDetail) HasValueBase64Binary() bool {
	if o != nil && !IsNil(o.ValueBase64Binary) {
		return true
	}

	return false
}

// SetValueBase64Binary gets a reference to the given string and assigns it to the ValueBase64Binary field.
func (o *AuditEventDetail) SetValueBase64Binary(v string) {
	o.ValueBase64Binary = &v
}

func (o AuditEventDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEventDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ValueString) {
		toSerialize["valueString"] = o.ValueString
	}
	if !IsNil(o.ValueBase64Binary) {
		toSerialize["valueBase64Binary"] = o.ValueBase64Binary
	}
	return toSerialize, nil
}

type NullableAuditEventDetail struct {
	value *AuditEventDetail
	isSet bool
}

func (v NullableAuditEventDetail) Get() *AuditEventDetail {
	return v.value
}

func (v *NullableAuditEventDetail) Set(val *AuditEventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEventDetail(val *AuditEventDetail) *NullableAuditEventDetail {
	return &NullableAuditEventDetail{value: val, isSet: true}
}

func (v NullableAuditEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


