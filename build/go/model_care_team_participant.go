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

// checks if the CareTeamParticipant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CareTeamParticipant{}

// CareTeamParticipant The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeamParticipant struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Indicates specific responsibility of an individual within the care team, such as \"Primary care physician\", \"Trained social worker counselor\", \"Caregiver\", etc.
	Role []CodeableConcept `json:"role,omitempty"`
	// The specific person or organization who is participating/expected to participate in the care team.
	Member *Reference `json:"member,omitempty"`
	// The organization of the practitioner.
	OnBehalfOf *Reference `json:"onBehalfOf,omitempty"`
	// Indicates when the specific member or organization did (or is intended to) come into effect and end.
	Period *Period `json:"period,omitempty"`
}

// NewCareTeamParticipant instantiates a new CareTeamParticipant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCareTeamParticipant() *CareTeamParticipant {
	this := CareTeamParticipant{}
	return &this
}

// NewCareTeamParticipantWithDefaults instantiates a new CareTeamParticipant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCareTeamParticipantWithDefaults() *CareTeamParticipant {
	this := CareTeamParticipant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CareTeamParticipant) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CareTeamParticipant) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CareTeamParticipant) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetRole() []CodeableConcept {
	if o == nil || IsNil(o.Role) {
		var ret []CodeableConcept
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetRoleOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given []CodeableConcept and assigns it to the Role field.
func (o *CareTeamParticipant) SetRole(v []CodeableConcept) {
	o.Role = v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetMember() Reference {
	if o == nil || IsNil(o.Member) {
		var ret Reference
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetMemberOk() (*Reference, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given Reference and assigns it to the Member field.
func (o *CareTeamParticipant) SetMember(v Reference) {
	o.Member = &v
}

// GetOnBehalfOf returns the OnBehalfOf field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetOnBehalfOf() Reference {
	if o == nil || IsNil(o.OnBehalfOf) {
		var ret Reference
		return ret
	}
	return *o.OnBehalfOf
}

// GetOnBehalfOfOk returns a tuple with the OnBehalfOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetOnBehalfOfOk() (*Reference, bool) {
	if o == nil || IsNil(o.OnBehalfOf) {
		return nil, false
	}
	return o.OnBehalfOf, true
}

// HasOnBehalfOf returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasOnBehalfOf() bool {
	if o != nil && !IsNil(o.OnBehalfOf) {
		return true
	}

	return false
}

// SetOnBehalfOf gets a reference to the given Reference and assigns it to the OnBehalfOf field.
func (o *CareTeamParticipant) SetOnBehalfOf(v Reference) {
	o.OnBehalfOf = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CareTeamParticipant) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CareTeamParticipant) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CareTeamParticipant) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *CareTeamParticipant) SetPeriod(v Period) {
	o.Period = &v
}

func (o CareTeamParticipant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CareTeamParticipant) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.OnBehalfOf) {
		toSerialize["onBehalfOf"] = o.OnBehalfOf
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return toSerialize, nil
}

type NullableCareTeamParticipant struct {
	value *CareTeamParticipant
	isSet bool
}

func (v NullableCareTeamParticipant) Get() *CareTeamParticipant {
	return v.value
}

func (v *NullableCareTeamParticipant) Set(val *CareTeamParticipant) {
	v.value = val
	v.isSet = true
}

func (v NullableCareTeamParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableCareTeamParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCareTeamParticipant(val *CareTeamParticipant) *NullableCareTeamParticipant {
	return &NullableCareTeamParticipant{value: val, isSet: true}
}

func (v NullableCareTeamParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCareTeamParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


