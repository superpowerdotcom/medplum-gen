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

// checks if the QuestionnaireEnableWhen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireEnableWhen{}

// QuestionnaireEnableWhen A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type QuestionnaireEnableWhen struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Question *string `json:"question,omitempty"`
	// Specifies the criteria by which the question is enabled.
	Operator *string `json:"operator,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerBoolean *bool `json:"answerBoolean,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerDecimal *float32 `json:"answerDecimal,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerInteger *float32 `json:"answerInteger,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerDate *string `json:"answerDate,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerDateTime *string `json:"answerDateTime,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerTime *string `json:"answerTime,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerString *string `json:"answerString,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerCoding *Coding `json:"answerCoding,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerQuantity *Quantity `json:"answerQuantity,omitempty"`
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	AnswerReference *Reference `json:"answerReference,omitempty"`
}

// NewQuestionnaireEnableWhen instantiates a new QuestionnaireEnableWhen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireEnableWhen() *QuestionnaireEnableWhen {
	this := QuestionnaireEnableWhen{}
	return &this
}

// NewQuestionnaireEnableWhenWithDefaults instantiates a new QuestionnaireEnableWhen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireEnableWhenWithDefaults() *QuestionnaireEnableWhen {
	this := QuestionnaireEnableWhen{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuestionnaireEnableWhen) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *QuestionnaireEnableWhen) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *QuestionnaireEnableWhen) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetQuestion() string {
	if o == nil || IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetQuestionOk() (*string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *QuestionnaireEnableWhen) SetQuestion(v string) {
	o.Question = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *QuestionnaireEnableWhen) SetOperator(v string) {
	o.Operator = &v
}

// GetAnswerBoolean returns the AnswerBoolean field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerBoolean() bool {
	if o == nil || IsNil(o.AnswerBoolean) {
		var ret bool
		return ret
	}
	return *o.AnswerBoolean
}

// GetAnswerBooleanOk returns a tuple with the AnswerBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.AnswerBoolean) {
		return nil, false
	}
	return o.AnswerBoolean, true
}

// HasAnswerBoolean returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerBoolean() bool {
	if o != nil && !IsNil(o.AnswerBoolean) {
		return true
	}

	return false
}

// SetAnswerBoolean gets a reference to the given bool and assigns it to the AnswerBoolean field.
func (o *QuestionnaireEnableWhen) SetAnswerBoolean(v bool) {
	o.AnswerBoolean = &v
}

// GetAnswerDecimal returns the AnswerDecimal field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerDecimal() float32 {
	if o == nil || IsNil(o.AnswerDecimal) {
		var ret float32
		return ret
	}
	return *o.AnswerDecimal
}

// GetAnswerDecimalOk returns a tuple with the AnswerDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerDecimalOk() (*float32, bool) {
	if o == nil || IsNil(o.AnswerDecimal) {
		return nil, false
	}
	return o.AnswerDecimal, true
}

// HasAnswerDecimal returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerDecimal() bool {
	if o != nil && !IsNil(o.AnswerDecimal) {
		return true
	}

	return false
}

// SetAnswerDecimal gets a reference to the given float32 and assigns it to the AnswerDecimal field.
func (o *QuestionnaireEnableWhen) SetAnswerDecimal(v float32) {
	o.AnswerDecimal = &v
}

// GetAnswerInteger returns the AnswerInteger field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerInteger() float32 {
	if o == nil || IsNil(o.AnswerInteger) {
		var ret float32
		return ret
	}
	return *o.AnswerInteger
}

// GetAnswerIntegerOk returns a tuple with the AnswerInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerIntegerOk() (*float32, bool) {
	if o == nil || IsNil(o.AnswerInteger) {
		return nil, false
	}
	return o.AnswerInteger, true
}

// HasAnswerInteger returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerInteger() bool {
	if o != nil && !IsNil(o.AnswerInteger) {
		return true
	}

	return false
}

// SetAnswerInteger gets a reference to the given float32 and assigns it to the AnswerInteger field.
func (o *QuestionnaireEnableWhen) SetAnswerInteger(v float32) {
	o.AnswerInteger = &v
}

// GetAnswerDate returns the AnswerDate field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerDate() string {
	if o == nil || IsNil(o.AnswerDate) {
		var ret string
		return ret
	}
	return *o.AnswerDate
}

// GetAnswerDateOk returns a tuple with the AnswerDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerDateOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerDate) {
		return nil, false
	}
	return o.AnswerDate, true
}

// HasAnswerDate returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerDate() bool {
	if o != nil && !IsNil(o.AnswerDate) {
		return true
	}

	return false
}

// SetAnswerDate gets a reference to the given string and assigns it to the AnswerDate field.
func (o *QuestionnaireEnableWhen) SetAnswerDate(v string) {
	o.AnswerDate = &v
}

// GetAnswerDateTime returns the AnswerDateTime field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerDateTime() string {
	if o == nil || IsNil(o.AnswerDateTime) {
		var ret string
		return ret
	}
	return *o.AnswerDateTime
}

// GetAnswerDateTimeOk returns a tuple with the AnswerDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerDateTime) {
		return nil, false
	}
	return o.AnswerDateTime, true
}

// HasAnswerDateTime returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerDateTime() bool {
	if o != nil && !IsNil(o.AnswerDateTime) {
		return true
	}

	return false
}

// SetAnswerDateTime gets a reference to the given string and assigns it to the AnswerDateTime field.
func (o *QuestionnaireEnableWhen) SetAnswerDateTime(v string) {
	o.AnswerDateTime = &v
}

// GetAnswerTime returns the AnswerTime field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerTime() string {
	if o == nil || IsNil(o.AnswerTime) {
		var ret string
		return ret
	}
	return *o.AnswerTime
}

// GetAnswerTimeOk returns a tuple with the AnswerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerTime) {
		return nil, false
	}
	return o.AnswerTime, true
}

// HasAnswerTime returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerTime() bool {
	if o != nil && !IsNil(o.AnswerTime) {
		return true
	}

	return false
}

// SetAnswerTime gets a reference to the given string and assigns it to the AnswerTime field.
func (o *QuestionnaireEnableWhen) SetAnswerTime(v string) {
	o.AnswerTime = &v
}

// GetAnswerString returns the AnswerString field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerString() string {
	if o == nil || IsNil(o.AnswerString) {
		var ret string
		return ret
	}
	return *o.AnswerString
}

// GetAnswerStringOk returns a tuple with the AnswerString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerStringOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerString) {
		return nil, false
	}
	return o.AnswerString, true
}

// HasAnswerString returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerString() bool {
	if o != nil && !IsNil(o.AnswerString) {
		return true
	}

	return false
}

// SetAnswerString gets a reference to the given string and assigns it to the AnswerString field.
func (o *QuestionnaireEnableWhen) SetAnswerString(v string) {
	o.AnswerString = &v
}

// GetAnswerCoding returns the AnswerCoding field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerCoding() Coding {
	if o == nil || IsNil(o.AnswerCoding) {
		var ret Coding
		return ret
	}
	return *o.AnswerCoding
}

// GetAnswerCodingOk returns a tuple with the AnswerCoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerCodingOk() (*Coding, bool) {
	if o == nil || IsNil(o.AnswerCoding) {
		return nil, false
	}
	return o.AnswerCoding, true
}

// HasAnswerCoding returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerCoding() bool {
	if o != nil && !IsNil(o.AnswerCoding) {
		return true
	}

	return false
}

// SetAnswerCoding gets a reference to the given Coding and assigns it to the AnswerCoding field.
func (o *QuestionnaireEnableWhen) SetAnswerCoding(v Coding) {
	o.AnswerCoding = &v
}

// GetAnswerQuantity returns the AnswerQuantity field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerQuantity() Quantity {
	if o == nil || IsNil(o.AnswerQuantity) {
		var ret Quantity
		return ret
	}
	return *o.AnswerQuantity
}

// GetAnswerQuantityOk returns a tuple with the AnswerQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.AnswerQuantity) {
		return nil, false
	}
	return o.AnswerQuantity, true
}

// HasAnswerQuantity returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerQuantity() bool {
	if o != nil && !IsNil(o.AnswerQuantity) {
		return true
	}

	return false
}

// SetAnswerQuantity gets a reference to the given Quantity and assigns it to the AnswerQuantity field.
func (o *QuestionnaireEnableWhen) SetAnswerQuantity(v Quantity) {
	o.AnswerQuantity = &v
}

// GetAnswerReference returns the AnswerReference field value if set, zero value otherwise.
func (o *QuestionnaireEnableWhen) GetAnswerReference() Reference {
	if o == nil || IsNil(o.AnswerReference) {
		var ret Reference
		return ret
	}
	return *o.AnswerReference
}

// GetAnswerReferenceOk returns a tuple with the AnswerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireEnableWhen) GetAnswerReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.AnswerReference) {
		return nil, false
	}
	return o.AnswerReference, true
}

// HasAnswerReference returns a boolean if a field has been set.
func (o *QuestionnaireEnableWhen) HasAnswerReference() bool {
	if o != nil && !IsNil(o.AnswerReference) {
		return true
	}

	return false
}

// SetAnswerReference gets a reference to the given Reference and assigns it to the AnswerReference field.
func (o *QuestionnaireEnableWhen) SetAnswerReference(v Reference) {
	o.AnswerReference = &v
}

func (o QuestionnaireEnableWhen) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireEnableWhen) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.AnswerBoolean) {
		toSerialize["answerBoolean"] = o.AnswerBoolean
	}
	if !IsNil(o.AnswerDecimal) {
		toSerialize["answerDecimal"] = o.AnswerDecimal
	}
	if !IsNil(o.AnswerInteger) {
		toSerialize["answerInteger"] = o.AnswerInteger
	}
	if !IsNil(o.AnswerDate) {
		toSerialize["answerDate"] = o.AnswerDate
	}
	if !IsNil(o.AnswerDateTime) {
		toSerialize["answerDateTime"] = o.AnswerDateTime
	}
	if !IsNil(o.AnswerTime) {
		toSerialize["answerTime"] = o.AnswerTime
	}
	if !IsNil(o.AnswerString) {
		toSerialize["answerString"] = o.AnswerString
	}
	if !IsNil(o.AnswerCoding) {
		toSerialize["answerCoding"] = o.AnswerCoding
	}
	if !IsNil(o.AnswerQuantity) {
		toSerialize["answerQuantity"] = o.AnswerQuantity
	}
	if !IsNil(o.AnswerReference) {
		toSerialize["answerReference"] = o.AnswerReference
	}
	return toSerialize, nil
}

type NullableQuestionnaireEnableWhen struct {
	value *QuestionnaireEnableWhen
	isSet bool
}

func (v NullableQuestionnaireEnableWhen) Get() *QuestionnaireEnableWhen {
	return v.value
}

func (v *NullableQuestionnaireEnableWhen) Set(val *QuestionnaireEnableWhen) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireEnableWhen) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireEnableWhen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireEnableWhen(val *QuestionnaireEnableWhen) *NullableQuestionnaireEnableWhen {
	return &NullableQuestionnaireEnableWhen{value: val, isSet: true}
}

func (v NullableQuestionnaireEnableWhen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireEnableWhen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


