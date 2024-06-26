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

// checks if the QuestionnaireInitial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireInitial{}

// QuestionnaireInitial A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type QuestionnaireInitial struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual value to for an initial answer.
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The actual value to for an initial answer.
	ValueDecimal *float32 `json:"valueDecimal,omitempty"`
	// The actual value to for an initial answer.
	ValueInteger *float32 `json:"valueInteger,omitempty"`
	// The actual value to for an initial answer.
	ValueDate *string `json:"valueDate,omitempty"`
	// The actual value to for an initial answer.
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	// The actual value to for an initial answer.
	ValueTime *string `json:"valueTime,omitempty"`
	// The actual value to for an initial answer.
	ValueString *string `json:"valueString,omitempty"`
	// The actual value to for an initial answer.
	ValueUri *string `json:"valueUri,omitempty"`
	// The actual value to for an initial answer.
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// The actual value to for an initial answer.
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// The actual value to for an initial answer.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The actual value to for an initial answer.
	ValueReference *Reference `json:"valueReference,omitempty"`
}

// NewQuestionnaireInitial instantiates a new QuestionnaireInitial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireInitial() *QuestionnaireInitial {
	this := QuestionnaireInitial{}
	return &this
}

// NewQuestionnaireInitialWithDefaults instantiates a new QuestionnaireInitial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireInitialWithDefaults() *QuestionnaireInitial {
	this := QuestionnaireInitial{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuestionnaireInitial) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *QuestionnaireInitial) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *QuestionnaireInitial) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetValueBoolean returns the ValueBoolean field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueBoolean() bool {
	if o == nil || IsNil(o.ValueBoolean) {
		var ret bool
		return ret
	}
	return *o.ValueBoolean
}

// GetValueBooleanOk returns a tuple with the ValueBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueBoolean) {
		return nil, false
	}
	return o.ValueBoolean, true
}

// HasValueBoolean returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueBoolean() bool {
	if o != nil && !IsNil(o.ValueBoolean) {
		return true
	}

	return false
}

// SetValueBoolean gets a reference to the given bool and assigns it to the ValueBoolean field.
func (o *QuestionnaireInitial) SetValueBoolean(v bool) {
	o.ValueBoolean = &v
}

// GetValueDecimal returns the ValueDecimal field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueDecimal() float32 {
	if o == nil || IsNil(o.ValueDecimal) {
		var ret float32
		return ret
	}
	return *o.ValueDecimal
}

// GetValueDecimalOk returns a tuple with the ValueDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueDecimalOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueDecimal) {
		return nil, false
	}
	return o.ValueDecimal, true
}

// HasValueDecimal returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueDecimal() bool {
	if o != nil && !IsNil(o.ValueDecimal) {
		return true
	}

	return false
}

// SetValueDecimal gets a reference to the given float32 and assigns it to the ValueDecimal field.
func (o *QuestionnaireInitial) SetValueDecimal(v float32) {
	o.ValueDecimal = &v
}

// GetValueInteger returns the ValueInteger field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueInteger() float32 {
	if o == nil || IsNil(o.ValueInteger) {
		var ret float32
		return ret
	}
	return *o.ValueInteger
}

// GetValueIntegerOk returns a tuple with the ValueInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueIntegerOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueInteger) {
		return nil, false
	}
	return o.ValueInteger, true
}

// HasValueInteger returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueInteger() bool {
	if o != nil && !IsNil(o.ValueInteger) {
		return true
	}

	return false
}

// SetValueInteger gets a reference to the given float32 and assigns it to the ValueInteger field.
func (o *QuestionnaireInitial) SetValueInteger(v float32) {
	o.ValueInteger = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueDate() string {
	if o == nil || IsNil(o.ValueDate) {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueDateOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueDate() bool {
	if o != nil && !IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *QuestionnaireInitial) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetValueDateTime returns the ValueDateTime field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueDateTime() string {
	if o == nil || IsNil(o.ValueDateTime) {
		var ret string
		return ret
	}
	return *o.ValueDateTime
}

// GetValueDateTimeOk returns a tuple with the ValueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDateTime) {
		return nil, false
	}
	return o.ValueDateTime, true
}

// HasValueDateTime returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueDateTime() bool {
	if o != nil && !IsNil(o.ValueDateTime) {
		return true
	}

	return false
}

// SetValueDateTime gets a reference to the given string and assigns it to the ValueDateTime field.
func (o *QuestionnaireInitial) SetValueDateTime(v string) {
	o.ValueDateTime = &v
}

// GetValueTime returns the ValueTime field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueTime() string {
	if o == nil || IsNil(o.ValueTime) {
		var ret string
		return ret
	}
	return *o.ValueTime
}

// GetValueTimeOk returns a tuple with the ValueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueTime) {
		return nil, false
	}
	return o.ValueTime, true
}

// HasValueTime returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueTime() bool {
	if o != nil && !IsNil(o.ValueTime) {
		return true
	}

	return false
}

// SetValueTime gets a reference to the given string and assigns it to the ValueTime field.
func (o *QuestionnaireInitial) SetValueTime(v string) {
	o.ValueTime = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *QuestionnaireInitial) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueUri returns the ValueUri field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueUri() string {
	if o == nil || IsNil(o.ValueUri) {
		var ret string
		return ret
	}
	return *o.ValueUri
}

// GetValueUriOk returns a tuple with the ValueUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueUriOk() (*string, bool) {
	if o == nil || IsNil(o.ValueUri) {
		return nil, false
	}
	return o.ValueUri, true
}

// HasValueUri returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueUri() bool {
	if o != nil && !IsNil(o.ValueUri) {
		return true
	}

	return false
}

// SetValueUri gets a reference to the given string and assigns it to the ValueUri field.
func (o *QuestionnaireInitial) SetValueUri(v string) {
	o.ValueUri = &v
}

// GetValueAttachment returns the ValueAttachment field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueAttachment() Attachment {
	if o == nil || IsNil(o.ValueAttachment) {
		var ret Attachment
		return ret
	}
	return *o.ValueAttachment
}

// GetValueAttachmentOk returns a tuple with the ValueAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.ValueAttachment) {
		return nil, false
	}
	return o.ValueAttachment, true
}

// HasValueAttachment returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueAttachment() bool {
	if o != nil && !IsNil(o.ValueAttachment) {
		return true
	}

	return false
}

// SetValueAttachment gets a reference to the given Attachment and assigns it to the ValueAttachment field.
func (o *QuestionnaireInitial) SetValueAttachment(v Attachment) {
	o.ValueAttachment = &v
}

// GetValueCoding returns the ValueCoding field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueCoding() Coding {
	if o == nil || IsNil(o.ValueCoding) {
		var ret Coding
		return ret
	}
	return *o.ValueCoding
}

// GetValueCodingOk returns a tuple with the ValueCoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueCodingOk() (*Coding, bool) {
	if o == nil || IsNil(o.ValueCoding) {
		return nil, false
	}
	return o.ValueCoding, true
}

// HasValueCoding returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueCoding() bool {
	if o != nil && !IsNil(o.ValueCoding) {
		return true
	}

	return false
}

// SetValueCoding gets a reference to the given Coding and assigns it to the ValueCoding field.
func (o *QuestionnaireInitial) SetValueCoding(v Coding) {
	o.ValueCoding = &v
}

// GetValueQuantity returns the ValueQuantity field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueQuantity() Quantity {
	if o == nil || IsNil(o.ValueQuantity) {
		var ret Quantity
		return ret
	}
	return *o.ValueQuantity
}

// GetValueQuantityOk returns a tuple with the ValueQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.ValueQuantity) {
		return nil, false
	}
	return o.ValueQuantity, true
}

// HasValueQuantity returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueQuantity() bool {
	if o != nil && !IsNil(o.ValueQuantity) {
		return true
	}

	return false
}

// SetValueQuantity gets a reference to the given Quantity and assigns it to the ValueQuantity field.
func (o *QuestionnaireInitial) SetValueQuantity(v Quantity) {
	o.ValueQuantity = &v
}

// GetValueReference returns the ValueReference field value if set, zero value otherwise.
func (o *QuestionnaireInitial) GetValueReference() Reference {
	if o == nil || IsNil(o.ValueReference) {
		var ret Reference
		return ret
	}
	return *o.ValueReference
}

// GetValueReferenceOk returns a tuple with the ValueReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireInitial) GetValueReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ValueReference) {
		return nil, false
	}
	return o.ValueReference, true
}

// HasValueReference returns a boolean if a field has been set.
func (o *QuestionnaireInitial) HasValueReference() bool {
	if o != nil && !IsNil(o.ValueReference) {
		return true
	}

	return false
}

// SetValueReference gets a reference to the given Reference and assigns it to the ValueReference field.
func (o *QuestionnaireInitial) SetValueReference(v Reference) {
	o.ValueReference = &v
}

func (o QuestionnaireInitial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireInitial) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ValueBoolean) {
		toSerialize["valueBoolean"] = o.ValueBoolean
	}
	if !IsNil(o.ValueDecimal) {
		toSerialize["valueDecimal"] = o.ValueDecimal
	}
	if !IsNil(o.ValueInteger) {
		toSerialize["valueInteger"] = o.ValueInteger
	}
	if !IsNil(o.ValueDate) {
		toSerialize["valueDate"] = o.ValueDate
	}
	if !IsNil(o.ValueDateTime) {
		toSerialize["valueDateTime"] = o.ValueDateTime
	}
	if !IsNil(o.ValueTime) {
		toSerialize["valueTime"] = o.ValueTime
	}
	if !IsNil(o.ValueString) {
		toSerialize["valueString"] = o.ValueString
	}
	if !IsNil(o.ValueUri) {
		toSerialize["valueUri"] = o.ValueUri
	}
	if !IsNil(o.ValueAttachment) {
		toSerialize["valueAttachment"] = o.ValueAttachment
	}
	if !IsNil(o.ValueCoding) {
		toSerialize["valueCoding"] = o.ValueCoding
	}
	if !IsNil(o.ValueQuantity) {
		toSerialize["valueQuantity"] = o.ValueQuantity
	}
	if !IsNil(o.ValueReference) {
		toSerialize["valueReference"] = o.ValueReference
	}
	return toSerialize, nil
}

type NullableQuestionnaireInitial struct {
	value *QuestionnaireInitial
	isSet bool
}

func (v NullableQuestionnaireInitial) Get() *QuestionnaireInitial {
	return v.value
}

func (v *NullableQuestionnaireInitial) Set(val *QuestionnaireInitial) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireInitial) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireInitial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireInitial(val *QuestionnaireInitial) *NullableQuestionnaireInitial {
	return &NullableQuestionnaireInitial{value: val, isSet: true}
}

func (v NullableQuestionnaireInitial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireInitial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


