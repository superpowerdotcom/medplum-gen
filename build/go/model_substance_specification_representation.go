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

// checks if the SubstanceSpecificationRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstanceSpecificationRepresentation{}

// SubstanceSpecificationRepresentation The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecificationRepresentation struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of structure (e.g. Full, Partial, Representative).
	Type *CodeableConcept `json:"type,omitempty"`
	// A sequence of Unicode characters
	Representation *string `json:"representation,omitempty"`
	// An attached file with the structural representation.
	Attachment *Attachment `json:"attachment,omitempty"`
}

// NewSubstanceSpecificationRepresentation instantiates a new SubstanceSpecificationRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstanceSpecificationRepresentation() *SubstanceSpecificationRepresentation {
	this := SubstanceSpecificationRepresentation{}
	return &this
}

// NewSubstanceSpecificationRepresentationWithDefaults instantiates a new SubstanceSpecificationRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceSpecificationRepresentationWithDefaults() *SubstanceSpecificationRepresentation {
	this := SubstanceSpecificationRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstanceSpecificationRepresentation) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstanceSpecificationRepresentation) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstanceSpecificationRepresentation) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *SubstanceSpecificationRepresentation) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetRepresentation returns the Representation field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetRepresentation() string {
	if o == nil || IsNil(o.Representation) {
		var ret string
		return ret
	}
	return *o.Representation
}

// GetRepresentationOk returns a tuple with the Representation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetRepresentationOk() (*string, bool) {
	if o == nil || IsNil(o.Representation) {
		return nil, false
	}
	return o.Representation, true
}

// HasRepresentation returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasRepresentation() bool {
	if o != nil && !IsNil(o.Representation) {
		return true
	}

	return false
}

// SetRepresentation gets a reference to the given string and assigns it to the Representation field.
func (o *SubstanceSpecificationRepresentation) SetRepresentation(v string) {
	o.Representation = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *SubstanceSpecificationRepresentation) GetAttachment() Attachment {
	if o == nil || IsNil(o.Attachment) {
		var ret Attachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationRepresentation) GetAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *SubstanceSpecificationRepresentation) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given Attachment and assigns it to the Attachment field.
func (o *SubstanceSpecificationRepresentation) SetAttachment(v Attachment) {
	o.Attachment = &v
}

func (o SubstanceSpecificationRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstanceSpecificationRepresentation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Representation) {
		toSerialize["representation"] = o.Representation
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	return toSerialize, nil
}

type NullableSubstanceSpecificationRepresentation struct {
	value *SubstanceSpecificationRepresentation
	isSet bool
}

func (v NullableSubstanceSpecificationRepresentation) Get() *SubstanceSpecificationRepresentation {
	return v.value
}

func (v *NullableSubstanceSpecificationRepresentation) Set(val *SubstanceSpecificationRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstanceSpecificationRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstanceSpecificationRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstanceSpecificationRepresentation(val *SubstanceSpecificationRepresentation) *NullableSubstanceSpecificationRepresentation {
	return &NullableSubstanceSpecificationRepresentation{value: val, isSet: true}
}

func (v NullableSubstanceSpecificationRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstanceSpecificationRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


