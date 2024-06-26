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

// checks if the MessageDefinitionAllowedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDefinitionAllowedResponse{}

// MessageDefinitionAllowedResponse Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinitionAllowedResponse struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A URI that is a reference to a canonical URL on a FHIR resource
	Message string `json:"message"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Situation *string `json:"situation,omitempty"`
}

type _MessageDefinitionAllowedResponse MessageDefinitionAllowedResponse

// NewMessageDefinitionAllowedResponse instantiates a new MessageDefinitionAllowedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDefinitionAllowedResponse(message string) *MessageDefinitionAllowedResponse {
	this := MessageDefinitionAllowedResponse{}
	this.Message = message
	return &this
}

// NewMessageDefinitionAllowedResponseWithDefaults instantiates a new MessageDefinitionAllowedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDefinitionAllowedResponseWithDefaults() *MessageDefinitionAllowedResponse {
	this := MessageDefinitionAllowedResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageDefinitionAllowedResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDefinitionAllowedResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageDefinitionAllowedResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MessageDefinitionAllowedResponse) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MessageDefinitionAllowedResponse) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDefinitionAllowedResponse) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MessageDefinitionAllowedResponse) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MessageDefinitionAllowedResponse) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MessageDefinitionAllowedResponse) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDefinitionAllowedResponse) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MessageDefinitionAllowedResponse) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MessageDefinitionAllowedResponse) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetMessage returns the Message field value
func (o *MessageDefinitionAllowedResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MessageDefinitionAllowedResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MessageDefinitionAllowedResponse) SetMessage(v string) {
	o.Message = v
}

// GetSituation returns the Situation field value if set, zero value otherwise.
func (o *MessageDefinitionAllowedResponse) GetSituation() string {
	if o == nil || IsNil(o.Situation) {
		var ret string
		return ret
	}
	return *o.Situation
}

// GetSituationOk returns a tuple with the Situation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDefinitionAllowedResponse) GetSituationOk() (*string, bool) {
	if o == nil || IsNil(o.Situation) {
		return nil, false
	}
	return o.Situation, true
}

// HasSituation returns a boolean if a field has been set.
func (o *MessageDefinitionAllowedResponse) HasSituation() bool {
	if o != nil && !IsNil(o.Situation) {
		return true
	}

	return false
}

// SetSituation gets a reference to the given string and assigns it to the Situation field.
func (o *MessageDefinitionAllowedResponse) SetSituation(v string) {
	o.Situation = &v
}

func (o MessageDefinitionAllowedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDefinitionAllowedResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["message"] = o.Message
	if !IsNil(o.Situation) {
		toSerialize["situation"] = o.Situation
	}
	return toSerialize, nil
}

func (o *MessageDefinitionAllowedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varMessageDefinitionAllowedResponse := _MessageDefinitionAllowedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDefinitionAllowedResponse)

	if err != nil {
		return err
	}

	*o = MessageDefinitionAllowedResponse(varMessageDefinitionAllowedResponse)

	return err
}

type NullableMessageDefinitionAllowedResponse struct {
	value *MessageDefinitionAllowedResponse
	isSet bool
}

func (v NullableMessageDefinitionAllowedResponse) Get() *MessageDefinitionAllowedResponse {
	return v.value
}

func (v *NullableMessageDefinitionAllowedResponse) Set(val *MessageDefinitionAllowedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDefinitionAllowedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDefinitionAllowedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDefinitionAllowedResponse(val *MessageDefinitionAllowedResponse) *NullableMessageDefinitionAllowedResponse {
	return &NullableMessageDefinitionAllowedResponse{value: val, isSet: true}
}

func (v NullableMessageDefinitionAllowedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDefinitionAllowedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

