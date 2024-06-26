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

// checks if the Timing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Timing{}

// Timing Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies specific times when the event occurs.
	Event []string `json:"event,omitempty"`
	// A set of rules that describe when the event is scheduled.
	Repeat *TimingRepeat `json:"repeat,omitempty"`
	// A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code *CodeableConcept `json:"code,omitempty"`
}

// NewTiming instantiates a new Timing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTiming() *Timing {
	this := Timing{}
	return &this
}

// NewTimingWithDefaults instantiates a new Timing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimingWithDefaults() *Timing {
	this := Timing{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Timing) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Timing) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Timing) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Timing) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Timing) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Timing) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Timing) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Timing) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Timing) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *Timing) GetEvent() []string {
	if o == nil || IsNil(o.Event) {
		var ret []string
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetEventOk() ([]string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *Timing) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given []string and assigns it to the Event field.
func (o *Timing) SetEvent(v []string) {
	o.Event = v
}

// GetRepeat returns the Repeat field value if set, zero value otherwise.
func (o *Timing) GetRepeat() TimingRepeat {
	if o == nil || IsNil(o.Repeat) {
		var ret TimingRepeat
		return ret
	}
	return *o.Repeat
}

// GetRepeatOk returns a tuple with the Repeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetRepeatOk() (*TimingRepeat, bool) {
	if o == nil || IsNil(o.Repeat) {
		return nil, false
	}
	return o.Repeat, true
}

// HasRepeat returns a boolean if a field has been set.
func (o *Timing) HasRepeat() bool {
	if o != nil && !IsNil(o.Repeat) {
		return true
	}

	return false
}

// SetRepeat gets a reference to the given TimingRepeat and assigns it to the Repeat field.
func (o *Timing) SetRepeat(v TimingRepeat) {
	o.Repeat = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Timing) GetCode() CodeableConcept {
	if o == nil || IsNil(o.Code) {
		var ret CodeableConcept
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timing) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Timing) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given CodeableConcept and assigns it to the Code field.
func (o *Timing) SetCode(v CodeableConcept) {
	o.Code = &v
}

func (o Timing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Timing) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Repeat) {
		toSerialize["repeat"] = o.Repeat
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableTiming struct {
	value *Timing
	isSet bool
}

func (v NullableTiming) Get() *Timing {
	return v.value
}

func (v *NullableTiming) Set(val *Timing) {
	v.value = val
	v.isSet = true
}

func (v NullableTiming) IsSet() bool {
	return v.isSet
}

func (v *NullableTiming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTiming(val *Timing) *NullableTiming {
	return &NullableTiming{value: val, isSet: true}
}

func (v NullableTiming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTiming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

