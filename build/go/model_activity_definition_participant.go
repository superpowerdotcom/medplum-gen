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

// checks if the ActivityDefinitionParticipant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityDefinitionParticipant{}

// ActivityDefinitionParticipant This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinitionParticipant struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Type *string `json:"type,omitempty"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `json:"role,omitempty"`
}

// NewActivityDefinitionParticipant instantiates a new ActivityDefinitionParticipant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityDefinitionParticipant() *ActivityDefinitionParticipant {
	this := ActivityDefinitionParticipant{}
	return &this
}

// NewActivityDefinitionParticipantWithDefaults instantiates a new ActivityDefinitionParticipant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityDefinitionParticipantWithDefaults() *ActivityDefinitionParticipant {
	this := ActivityDefinitionParticipant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivityDefinitionParticipant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDefinitionParticipant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivityDefinitionParticipant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActivityDefinitionParticipant) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ActivityDefinitionParticipant) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDefinitionParticipant) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ActivityDefinitionParticipant) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ActivityDefinitionParticipant) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ActivityDefinitionParticipant) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDefinitionParticipant) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ActivityDefinitionParticipant) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ActivityDefinitionParticipant) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActivityDefinitionParticipant) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDefinitionParticipant) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActivityDefinitionParticipant) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActivityDefinitionParticipant) SetType(v string) {
	o.Type = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ActivityDefinitionParticipant) GetRole() CodeableConcept {
	if o == nil || IsNil(o.Role) {
		var ret CodeableConcept
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDefinitionParticipant) GetRoleOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ActivityDefinitionParticipant) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given CodeableConcept and assigns it to the Role field.
func (o *ActivityDefinitionParticipant) SetRole(v CodeableConcept) {
	o.Role = &v
}

func (o ActivityDefinitionParticipant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityDefinitionParticipant) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableActivityDefinitionParticipant struct {
	value *ActivityDefinitionParticipant
	isSet bool
}

func (v NullableActivityDefinitionParticipant) Get() *ActivityDefinitionParticipant {
	return v.value
}

func (v *NullableActivityDefinitionParticipant) Set(val *ActivityDefinitionParticipant) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityDefinitionParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityDefinitionParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityDefinitionParticipant(val *ActivityDefinitionParticipant) *NullableActivityDefinitionParticipant {
	return &NullableActivityDefinitionParticipant{value: val, isSet: true}
}

func (v NullableActivityDefinitionParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityDefinitionParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


