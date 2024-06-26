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

// checks if the CompositionRelatesTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompositionRelatesTo{}

// CompositionRelatesTo A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
type CompositionRelatesTo struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Code *string `json:"code,omitempty"`
	// The target composition/document of this relationship.
	TargetIdentifier *Identifier `json:"targetIdentifier,omitempty"`
	// The target composition/document of this relationship.
	TargetReference *Reference `json:"targetReference,omitempty"`
}

// NewCompositionRelatesTo instantiates a new CompositionRelatesTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositionRelatesTo() *CompositionRelatesTo {
	this := CompositionRelatesTo{}
	return &this
}

// NewCompositionRelatesToWithDefaults instantiates a new CompositionRelatesTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositionRelatesToWithDefaults() *CompositionRelatesTo {
	this := CompositionRelatesTo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompositionRelatesTo) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CompositionRelatesTo) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CompositionRelatesTo) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CompositionRelatesTo) SetCode(v string) {
	o.Code = &v
}

// GetTargetIdentifier returns the TargetIdentifier field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetTargetIdentifier() Identifier {
	if o == nil || IsNil(o.TargetIdentifier) {
		var ret Identifier
		return ret
	}
	return *o.TargetIdentifier
}

// GetTargetIdentifierOk returns a tuple with the TargetIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetTargetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.TargetIdentifier) {
		return nil, false
	}
	return o.TargetIdentifier, true
}

// HasTargetIdentifier returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasTargetIdentifier() bool {
	if o != nil && !IsNil(o.TargetIdentifier) {
		return true
	}

	return false
}

// SetTargetIdentifier gets a reference to the given Identifier and assigns it to the TargetIdentifier field.
func (o *CompositionRelatesTo) SetTargetIdentifier(v Identifier) {
	o.TargetIdentifier = &v
}

// GetTargetReference returns the TargetReference field value if set, zero value otherwise.
func (o *CompositionRelatesTo) GetTargetReference() Reference {
	if o == nil || IsNil(o.TargetReference) {
		var ret Reference
		return ret
	}
	return *o.TargetReference
}

// GetTargetReferenceOk returns a tuple with the TargetReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionRelatesTo) GetTargetReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.TargetReference) {
		return nil, false
	}
	return o.TargetReference, true
}

// HasTargetReference returns a boolean if a field has been set.
func (o *CompositionRelatesTo) HasTargetReference() bool {
	if o != nil && !IsNil(o.TargetReference) {
		return true
	}

	return false
}

// SetTargetReference gets a reference to the given Reference and assigns it to the TargetReference field.
func (o *CompositionRelatesTo) SetTargetReference(v Reference) {
	o.TargetReference = &v
}

func (o CompositionRelatesTo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositionRelatesTo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.TargetIdentifier) {
		toSerialize["targetIdentifier"] = o.TargetIdentifier
	}
	if !IsNil(o.TargetReference) {
		toSerialize["targetReference"] = o.TargetReference
	}
	return toSerialize, nil
}

type NullableCompositionRelatesTo struct {
	value *CompositionRelatesTo
	isSet bool
}

func (v NullableCompositionRelatesTo) Get() *CompositionRelatesTo {
	return v.value
}

func (v *NullableCompositionRelatesTo) Set(val *CompositionRelatesTo) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositionRelatesTo) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositionRelatesTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositionRelatesTo(val *CompositionRelatesTo) *NullableCompositionRelatesTo {
	return &NullableCompositionRelatesTo{value: val, isSet: true}
}

func (v NullableCompositionRelatesTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositionRelatesTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


