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

// checks if the QuestionnaireAnswerOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireAnswerOption{}

// QuestionnaireAnswerOption A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type QuestionnaireAnswerOption struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueInteger *float32 `json:"valueInteger,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueDate *string `json:"valueDate,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueTime *string `json:"valueTime,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueString *string `json:"valueString,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// A potential answer that's allowed as the answer to this question.
	ValueReference *Reference `json:"valueReference,omitempty"`
	// Value of \"true\" or \"false\"
	InitialSelected *bool `json:"initialSelected,omitempty"`
}

// NewQuestionnaireAnswerOption instantiates a new QuestionnaireAnswerOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireAnswerOption() *QuestionnaireAnswerOption {
	this := QuestionnaireAnswerOption{}
	return &this
}

// NewQuestionnaireAnswerOptionWithDefaults instantiates a new QuestionnaireAnswerOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireAnswerOptionWithDefaults() *QuestionnaireAnswerOption {
	this := QuestionnaireAnswerOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuestionnaireAnswerOption) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *QuestionnaireAnswerOption) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *QuestionnaireAnswerOption) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetValueInteger returns the ValueInteger field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueInteger() float32 {
	if o == nil || IsNil(o.ValueInteger) {
		var ret float32
		return ret
	}
	return *o.ValueInteger
}

// GetValueIntegerOk returns a tuple with the ValueInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueIntegerOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueInteger) {
		return nil, false
	}
	return o.ValueInteger, true
}

// HasValueInteger returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueInteger() bool {
	if o != nil && !IsNil(o.ValueInteger) {
		return true
	}

	return false
}

// SetValueInteger gets a reference to the given float32 and assigns it to the ValueInteger field.
func (o *QuestionnaireAnswerOption) SetValueInteger(v float32) {
	o.ValueInteger = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueDate() string {
	if o == nil || IsNil(o.ValueDate) {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueDateOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueDate() bool {
	if o != nil && !IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *QuestionnaireAnswerOption) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetValueTime returns the ValueTime field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueTime() string {
	if o == nil || IsNil(o.ValueTime) {
		var ret string
		return ret
	}
	return *o.ValueTime
}

// GetValueTimeOk returns a tuple with the ValueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueTime) {
		return nil, false
	}
	return o.ValueTime, true
}

// HasValueTime returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueTime() bool {
	if o != nil && !IsNil(o.ValueTime) {
		return true
	}

	return false
}

// SetValueTime gets a reference to the given string and assigns it to the ValueTime field.
func (o *QuestionnaireAnswerOption) SetValueTime(v string) {
	o.ValueTime = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *QuestionnaireAnswerOption) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueCoding returns the ValueCoding field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueCoding() Coding {
	if o == nil || IsNil(o.ValueCoding) {
		var ret Coding
		return ret
	}
	return *o.ValueCoding
}

// GetValueCodingOk returns a tuple with the ValueCoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueCodingOk() (*Coding, bool) {
	if o == nil || IsNil(o.ValueCoding) {
		return nil, false
	}
	return o.ValueCoding, true
}

// HasValueCoding returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueCoding() bool {
	if o != nil && !IsNil(o.ValueCoding) {
		return true
	}

	return false
}

// SetValueCoding gets a reference to the given Coding and assigns it to the ValueCoding field.
func (o *QuestionnaireAnswerOption) SetValueCoding(v Coding) {
	o.ValueCoding = &v
}

// GetValueReference returns the ValueReference field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetValueReference() Reference {
	if o == nil || IsNil(o.ValueReference) {
		var ret Reference
		return ret
	}
	return *o.ValueReference
}

// GetValueReferenceOk returns a tuple with the ValueReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetValueReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ValueReference) {
		return nil, false
	}
	return o.ValueReference, true
}

// HasValueReference returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasValueReference() bool {
	if o != nil && !IsNil(o.ValueReference) {
		return true
	}

	return false
}

// SetValueReference gets a reference to the given Reference and assigns it to the ValueReference field.
func (o *QuestionnaireAnswerOption) SetValueReference(v Reference) {
	o.ValueReference = &v
}

// GetInitialSelected returns the InitialSelected field value if set, zero value otherwise.
func (o *QuestionnaireAnswerOption) GetInitialSelected() bool {
	if o == nil || IsNil(o.InitialSelected) {
		var ret bool
		return ret
	}
	return *o.InitialSelected
}

// GetInitialSelectedOk returns a tuple with the InitialSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerOption) GetInitialSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.InitialSelected) {
		return nil, false
	}
	return o.InitialSelected, true
}

// HasInitialSelected returns a boolean if a field has been set.
func (o *QuestionnaireAnswerOption) HasInitialSelected() bool {
	if o != nil && !IsNil(o.InitialSelected) {
		return true
	}

	return false
}

// SetInitialSelected gets a reference to the given bool and assigns it to the InitialSelected field.
func (o *QuestionnaireAnswerOption) SetInitialSelected(v bool) {
	o.InitialSelected = &v
}

func (o QuestionnaireAnswerOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireAnswerOption) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ValueInteger) {
		toSerialize["valueInteger"] = o.ValueInteger
	}
	if !IsNil(o.ValueDate) {
		toSerialize["valueDate"] = o.ValueDate
	}
	if !IsNil(o.ValueTime) {
		toSerialize["valueTime"] = o.ValueTime
	}
	if !IsNil(o.ValueString) {
		toSerialize["valueString"] = o.ValueString
	}
	if !IsNil(o.ValueCoding) {
		toSerialize["valueCoding"] = o.ValueCoding
	}
	if !IsNil(o.ValueReference) {
		toSerialize["valueReference"] = o.ValueReference
	}
	if !IsNil(o.InitialSelected) {
		toSerialize["initialSelected"] = o.InitialSelected
	}
	return toSerialize, nil
}

type NullableQuestionnaireAnswerOption struct {
	value *QuestionnaireAnswerOption
	isSet bool
}

func (v NullableQuestionnaireAnswerOption) Get() *QuestionnaireAnswerOption {
	return v.value
}

func (v *NullableQuestionnaireAnswerOption) Set(val *QuestionnaireAnswerOption) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireAnswerOption) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireAnswerOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireAnswerOption(val *QuestionnaireAnswerOption) *NullableQuestionnaireAnswerOption {
	return &NullableQuestionnaireAnswerOption{value: val, isSet: true}
}

func (v NullableQuestionnaireAnswerOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireAnswerOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

