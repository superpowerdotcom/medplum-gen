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

// checks if the ExplanationOfBenefitProcessNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplanationOfBenefitProcessNote{}

// ExplanationOfBenefitProcessNote This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitProcessNote struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An integer with a value that is positive (e.g. >0)
	Number *float32 `json:"number,omitempty"`
	// The business purpose of the note text.
	Type *string `json:"type,omitempty"`
	// A sequence of Unicode characters
	Text *string `json:"text,omitempty"`
	// A code to define the language used in the text of the note.
	Language *CodeableConcept `json:"language,omitempty"`
}

// NewExplanationOfBenefitProcessNote instantiates a new ExplanationOfBenefitProcessNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplanationOfBenefitProcessNote() *ExplanationOfBenefitProcessNote {
	this := ExplanationOfBenefitProcessNote{}
	return &this
}

// NewExplanationOfBenefitProcessNoteWithDefaults instantiates a new ExplanationOfBenefitProcessNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplanationOfBenefitProcessNoteWithDefaults() *ExplanationOfBenefitProcessNote {
	this := ExplanationOfBenefitProcessNote{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExplanationOfBenefitProcessNote) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExplanationOfBenefitProcessNote) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExplanationOfBenefitProcessNote) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetNumber() float32 {
	if o == nil || IsNil(o.Number) {
		var ret float32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float32 and assigns it to the Number field.
func (o *ExplanationOfBenefitProcessNote) SetNumber(v float32) {
	o.Number = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExplanationOfBenefitProcessNote) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ExplanationOfBenefitProcessNote) SetText(v string) {
	o.Text = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ExplanationOfBenefitProcessNote) GetLanguage() CodeableConcept {
	if o == nil || IsNil(o.Language) {
		var ret CodeableConcept
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplanationOfBenefitProcessNote) GetLanguageOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ExplanationOfBenefitProcessNote) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given CodeableConcept and assigns it to the Language field.
func (o *ExplanationOfBenefitProcessNote) SetLanguage(v CodeableConcept) {
	o.Language = &v
}

func (o ExplanationOfBenefitProcessNote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplanationOfBenefitProcessNote) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableExplanationOfBenefitProcessNote struct {
	value *ExplanationOfBenefitProcessNote
	isSet bool
}

func (v NullableExplanationOfBenefitProcessNote) Get() *ExplanationOfBenefitProcessNote {
	return v.value
}

func (v *NullableExplanationOfBenefitProcessNote) Set(val *ExplanationOfBenefitProcessNote) {
	v.value = val
	v.isSet = true
}

func (v NullableExplanationOfBenefitProcessNote) IsSet() bool {
	return v.isSet
}

func (v *NullableExplanationOfBenefitProcessNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplanationOfBenefitProcessNote(val *ExplanationOfBenefitProcessNote) *NullableExplanationOfBenefitProcessNote {
	return &NullableExplanationOfBenefitProcessNote{value: val, isSet: true}
}

func (v NullableExplanationOfBenefitProcessNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplanationOfBenefitProcessNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


