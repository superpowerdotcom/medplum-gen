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

// checks if the QuestionnaireResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireResponseItem{}

// QuestionnaireResponseItem A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponseItem struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	LinkId *string `json:"linkId,omitempty"`
	// String of characters used to identify a name or a resource
	Definition *string `json:"definition,omitempty"`
	// A sequence of Unicode characters
	Text *string `json:"text,omitempty"`
	// The respondent's answer(s) to the question.
	Answer []QuestionnaireResponseAnswer `json:"answer,omitempty"`
	// Questions or sub-groups nested beneath a question or group.
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// NewQuestionnaireResponseItem instantiates a new QuestionnaireResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireResponseItem() *QuestionnaireResponseItem {
	this := QuestionnaireResponseItem{}
	return &this
}

// NewQuestionnaireResponseItemWithDefaults instantiates a new QuestionnaireResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireResponseItemWithDefaults() *QuestionnaireResponseItem {
	this := QuestionnaireResponseItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuestionnaireResponseItem) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *QuestionnaireResponseItem) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *QuestionnaireResponseItem) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *QuestionnaireResponseItem) SetLinkId(v string) {
	o.LinkId = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetDefinition() string {
	if o == nil || IsNil(o.Definition) {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetDefinitionOk() (*string, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *QuestionnaireResponseItem) SetDefinition(v string) {
	o.Definition = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *QuestionnaireResponseItem) SetText(v string) {
	o.Text = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetAnswer() []QuestionnaireResponseAnswer {
	if o == nil || IsNil(o.Answer) {
		var ret []QuestionnaireResponseAnswer
		return ret
	}
	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetAnswerOk() ([]QuestionnaireResponseAnswer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given []QuestionnaireResponseAnswer and assigns it to the Answer field.
func (o *QuestionnaireResponseItem) SetAnswer(v []QuestionnaireResponseAnswer) {
	o.Answer = v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *QuestionnaireResponseItem) GetItem() []QuestionnaireResponseItem {
	if o == nil || IsNil(o.Item) {
		var ret []QuestionnaireResponseItem
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponseItem) GetItemOk() ([]QuestionnaireResponseItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *QuestionnaireResponseItem) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []QuestionnaireResponseItem and assigns it to the Item field.
func (o *QuestionnaireResponseItem) SetItem(v []QuestionnaireResponseItem) {
	o.Item = v
}

func (o QuestionnaireResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireResponseItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableQuestionnaireResponseItem struct {
	value *QuestionnaireResponseItem
	isSet bool
}

func (v NullableQuestionnaireResponseItem) Get() *QuestionnaireResponseItem {
	return v.value
}

func (v *NullableQuestionnaireResponseItem) Set(val *QuestionnaireResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireResponseItem(val *QuestionnaireResponseItem) *NullableQuestionnaireResponseItem {
	return &NullableQuestionnaireResponseItem{value: val, isSet: true}
}

func (v NullableQuestionnaireResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


