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

// checks if the TaskRestriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskRestriction{}

// TaskRestriction A task to be performed.
type TaskRestriction struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Repetitions *float32 `json:"repetitions,omitempty"`
	// Over what time-period is fulfillment sought.
	Period *Period `json:"period,omitempty"`
	// For requests that are targeted to more than on potential recipient/target, for whom is fulfillment sought?
	Recipient []Reference `json:"recipient,omitempty"`
}

// NewTaskRestriction instantiates a new TaskRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskRestriction() *TaskRestriction {
	this := TaskRestriction{}
	return &this
}

// NewTaskRestrictionWithDefaults instantiates a new TaskRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskRestrictionWithDefaults() *TaskRestriction {
	this := TaskRestriction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskRestriction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskRestriction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskRestriction) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TaskRestriction) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TaskRestriction) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TaskRestriction) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TaskRestriction) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TaskRestriction) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TaskRestriction) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetRepetitions returns the Repetitions field value if set, zero value otherwise.
func (o *TaskRestriction) GetRepetitions() float32 {
	if o == nil || IsNil(o.Repetitions) {
		var ret float32
		return ret
	}
	return *o.Repetitions
}

// GetRepetitionsOk returns a tuple with the Repetitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetRepetitionsOk() (*float32, bool) {
	if o == nil || IsNil(o.Repetitions) {
		return nil, false
	}
	return o.Repetitions, true
}

// HasRepetitions returns a boolean if a field has been set.
func (o *TaskRestriction) HasRepetitions() bool {
	if o != nil && !IsNil(o.Repetitions) {
		return true
	}

	return false
}

// SetRepetitions gets a reference to the given float32 and assigns it to the Repetitions field.
func (o *TaskRestriction) SetRepetitions(v float32) {
	o.Repetitions = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *TaskRestriction) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *TaskRestriction) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *TaskRestriction) SetPeriod(v Period) {
	o.Period = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *TaskRestriction) GetRecipient() []Reference {
	if o == nil || IsNil(o.Recipient) {
		var ret []Reference
		return ret
	}
	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskRestriction) GetRecipientOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *TaskRestriction) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given []Reference and assigns it to the Recipient field.
func (o *TaskRestriction) SetRecipient(v []Reference) {
	o.Recipient = v
}

func (o TaskRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskRestriction) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Repetitions) {
		toSerialize["repetitions"] = o.Repetitions
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	return toSerialize, nil
}

type NullableTaskRestriction struct {
	value *TaskRestriction
	isSet bool
}

func (v NullableTaskRestriction) Get() *TaskRestriction {
	return v.value
}

func (v *NullableTaskRestriction) Set(val *TaskRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskRestriction(val *TaskRestriction) *NullableTaskRestriction {
	return &NullableTaskRestriction{value: val, isSet: true}
}

func (v NullableTaskRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


