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

// checks if the CommunicationRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationRequestPayload{}

// CommunicationRequestPayload A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequestPayload struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentString *string `json:"contentString,omitempty"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentReference *Reference `json:"contentReference,omitempty"`
}

// NewCommunicationRequestPayload instantiates a new CommunicationRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationRequestPayload() *CommunicationRequestPayload {
	this := CommunicationRequestPayload{}
	return &this
}

// NewCommunicationRequestPayloadWithDefaults instantiates a new CommunicationRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationRequestPayloadWithDefaults() *CommunicationRequestPayload {
	this := CommunicationRequestPayload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommunicationRequestPayload) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CommunicationRequestPayload) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CommunicationRequestPayload) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetContentString returns the ContentString field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetContentString() string {
	if o == nil || IsNil(o.ContentString) {
		var ret string
		return ret
	}
	return *o.ContentString
}

// GetContentStringOk returns a tuple with the ContentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetContentStringOk() (*string, bool) {
	if o == nil || IsNil(o.ContentString) {
		return nil, false
	}
	return o.ContentString, true
}

// HasContentString returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasContentString() bool {
	if o != nil && !IsNil(o.ContentString) {
		return true
	}

	return false
}

// SetContentString gets a reference to the given string and assigns it to the ContentString field.
func (o *CommunicationRequestPayload) SetContentString(v string) {
	o.ContentString = &v
}

// GetContentAttachment returns the ContentAttachment field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetContentAttachment() Attachment {
	if o == nil || IsNil(o.ContentAttachment) {
		var ret Attachment
		return ret
	}
	return *o.ContentAttachment
}

// GetContentAttachmentOk returns a tuple with the ContentAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetContentAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.ContentAttachment) {
		return nil, false
	}
	return o.ContentAttachment, true
}

// HasContentAttachment returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasContentAttachment() bool {
	if o != nil && !IsNil(o.ContentAttachment) {
		return true
	}

	return false
}

// SetContentAttachment gets a reference to the given Attachment and assigns it to the ContentAttachment field.
func (o *CommunicationRequestPayload) SetContentAttachment(v Attachment) {
	o.ContentAttachment = &v
}

// GetContentReference returns the ContentReference field value if set, zero value otherwise.
func (o *CommunicationRequestPayload) GetContentReference() Reference {
	if o == nil || IsNil(o.ContentReference) {
		var ret Reference
		return ret
	}
	return *o.ContentReference
}

// GetContentReferenceOk returns a tuple with the ContentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequestPayload) GetContentReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ContentReference) {
		return nil, false
	}
	return o.ContentReference, true
}

// HasContentReference returns a boolean if a field has been set.
func (o *CommunicationRequestPayload) HasContentReference() bool {
	if o != nil && !IsNil(o.ContentReference) {
		return true
	}

	return false
}

// SetContentReference gets a reference to the given Reference and assigns it to the ContentReference field.
func (o *CommunicationRequestPayload) SetContentReference(v Reference) {
	o.ContentReference = &v
}

func (o CommunicationRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationRequestPayload) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ContentString) {
		toSerialize["contentString"] = o.ContentString
	}
	if !IsNil(o.ContentAttachment) {
		toSerialize["contentAttachment"] = o.ContentAttachment
	}
	if !IsNil(o.ContentReference) {
		toSerialize["contentReference"] = o.ContentReference
	}
	return toSerialize, nil
}

type NullableCommunicationRequestPayload struct {
	value *CommunicationRequestPayload
	isSet bool
}

func (v NullableCommunicationRequestPayload) Get() *CommunicationRequestPayload {
	return v.value
}

func (v *NullableCommunicationRequestPayload) Set(val *CommunicationRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationRequestPayload(val *CommunicationRequestPayload) *NullableCommunicationRequestPayload {
	return &NullableCommunicationRequestPayload{value: val, isSet: true}
}

func (v NullableCommunicationRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

