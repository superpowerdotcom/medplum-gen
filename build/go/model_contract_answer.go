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

// checks if the ContractAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractAnswer{}

// ContractAnswer Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type ContractAnswer struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueDecimal *float32 `json:"valueDecimal,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueInteger *float32 `json:"valueInteger,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueDate *string `json:"valueDate,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueTime *string `json:"valueTime,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueString *string `json:"valueString,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueUri *string `json:"valueUri,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	ValueReference *Reference `json:"valueReference,omitempty"`
}

// NewContractAnswer instantiates a new ContractAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractAnswer() *ContractAnswer {
	this := ContractAnswer{}
	return &this
}

// NewContractAnswerWithDefaults instantiates a new ContractAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractAnswerWithDefaults() *ContractAnswer {
	this := ContractAnswer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractAnswer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractAnswer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractAnswer) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ContractAnswer) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ContractAnswer) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ContractAnswer) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ContractAnswer) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ContractAnswer) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ContractAnswer) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetValueBoolean returns the ValueBoolean field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueBoolean() bool {
	if o == nil || IsNil(o.ValueBoolean) {
		var ret bool
		return ret
	}
	return *o.ValueBoolean
}

// GetValueBooleanOk returns a tuple with the ValueBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueBoolean) {
		return nil, false
	}
	return o.ValueBoolean, true
}

// HasValueBoolean returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueBoolean() bool {
	if o != nil && !IsNil(o.ValueBoolean) {
		return true
	}

	return false
}

// SetValueBoolean gets a reference to the given bool and assigns it to the ValueBoolean field.
func (o *ContractAnswer) SetValueBoolean(v bool) {
	o.ValueBoolean = &v
}

// GetValueDecimal returns the ValueDecimal field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueDecimal() float32 {
	if o == nil || IsNil(o.ValueDecimal) {
		var ret float32
		return ret
	}
	return *o.ValueDecimal
}

// GetValueDecimalOk returns a tuple with the ValueDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueDecimalOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueDecimal) {
		return nil, false
	}
	return o.ValueDecimal, true
}

// HasValueDecimal returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueDecimal() bool {
	if o != nil && !IsNil(o.ValueDecimal) {
		return true
	}

	return false
}

// SetValueDecimal gets a reference to the given float32 and assigns it to the ValueDecimal field.
func (o *ContractAnswer) SetValueDecimal(v float32) {
	o.ValueDecimal = &v
}

// GetValueInteger returns the ValueInteger field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueInteger() float32 {
	if o == nil || IsNil(o.ValueInteger) {
		var ret float32
		return ret
	}
	return *o.ValueInteger
}

// GetValueIntegerOk returns a tuple with the ValueInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueIntegerOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueInteger) {
		return nil, false
	}
	return o.ValueInteger, true
}

// HasValueInteger returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueInteger() bool {
	if o != nil && !IsNil(o.ValueInteger) {
		return true
	}

	return false
}

// SetValueInteger gets a reference to the given float32 and assigns it to the ValueInteger field.
func (o *ContractAnswer) SetValueInteger(v float32) {
	o.ValueInteger = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueDate() string {
	if o == nil || IsNil(o.ValueDate) {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueDateOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueDate() bool {
	if o != nil && !IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *ContractAnswer) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetValueDateTime returns the ValueDateTime field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueDateTime() string {
	if o == nil || IsNil(o.ValueDateTime) {
		var ret string
		return ret
	}
	return *o.ValueDateTime
}

// GetValueDateTimeOk returns a tuple with the ValueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDateTime) {
		return nil, false
	}
	return o.ValueDateTime, true
}

// HasValueDateTime returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueDateTime() bool {
	if o != nil && !IsNil(o.ValueDateTime) {
		return true
	}

	return false
}

// SetValueDateTime gets a reference to the given string and assigns it to the ValueDateTime field.
func (o *ContractAnswer) SetValueDateTime(v string) {
	o.ValueDateTime = &v
}

// GetValueTime returns the ValueTime field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueTime() string {
	if o == nil || IsNil(o.ValueTime) {
		var ret string
		return ret
	}
	return *o.ValueTime
}

// GetValueTimeOk returns a tuple with the ValueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueTime) {
		return nil, false
	}
	return o.ValueTime, true
}

// HasValueTime returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueTime() bool {
	if o != nil && !IsNil(o.ValueTime) {
		return true
	}

	return false
}

// SetValueTime gets a reference to the given string and assigns it to the ValueTime field.
func (o *ContractAnswer) SetValueTime(v string) {
	o.ValueTime = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueString() string {
	if o == nil || IsNil(o.ValueString) {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.ValueString) {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueString() bool {
	if o != nil && !IsNil(o.ValueString) {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *ContractAnswer) SetValueString(v string) {
	o.ValueString = &v
}

// GetValueUri returns the ValueUri field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueUri() string {
	if o == nil || IsNil(o.ValueUri) {
		var ret string
		return ret
	}
	return *o.ValueUri
}

// GetValueUriOk returns a tuple with the ValueUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueUriOk() (*string, bool) {
	if o == nil || IsNil(o.ValueUri) {
		return nil, false
	}
	return o.ValueUri, true
}

// HasValueUri returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueUri() bool {
	if o != nil && !IsNil(o.ValueUri) {
		return true
	}

	return false
}

// SetValueUri gets a reference to the given string and assigns it to the ValueUri field.
func (o *ContractAnswer) SetValueUri(v string) {
	o.ValueUri = &v
}

// GetValueAttachment returns the ValueAttachment field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueAttachment() Attachment {
	if o == nil || IsNil(o.ValueAttachment) {
		var ret Attachment
		return ret
	}
	return *o.ValueAttachment
}

// GetValueAttachmentOk returns a tuple with the ValueAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueAttachmentOk() (*Attachment, bool) {
	if o == nil || IsNil(o.ValueAttachment) {
		return nil, false
	}
	return o.ValueAttachment, true
}

// HasValueAttachment returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueAttachment() bool {
	if o != nil && !IsNil(o.ValueAttachment) {
		return true
	}

	return false
}

// SetValueAttachment gets a reference to the given Attachment and assigns it to the ValueAttachment field.
func (o *ContractAnswer) SetValueAttachment(v Attachment) {
	o.ValueAttachment = &v
}

// GetValueCoding returns the ValueCoding field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueCoding() Coding {
	if o == nil || IsNil(o.ValueCoding) {
		var ret Coding
		return ret
	}
	return *o.ValueCoding
}

// GetValueCodingOk returns a tuple with the ValueCoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueCodingOk() (*Coding, bool) {
	if o == nil || IsNil(o.ValueCoding) {
		return nil, false
	}
	return o.ValueCoding, true
}

// HasValueCoding returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueCoding() bool {
	if o != nil && !IsNil(o.ValueCoding) {
		return true
	}

	return false
}

// SetValueCoding gets a reference to the given Coding and assigns it to the ValueCoding field.
func (o *ContractAnswer) SetValueCoding(v Coding) {
	o.ValueCoding = &v
}

// GetValueQuantity returns the ValueQuantity field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueQuantity() Quantity {
	if o == nil || IsNil(o.ValueQuantity) {
		var ret Quantity
		return ret
	}
	return *o.ValueQuantity
}

// GetValueQuantityOk returns a tuple with the ValueQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.ValueQuantity) {
		return nil, false
	}
	return o.ValueQuantity, true
}

// HasValueQuantity returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueQuantity() bool {
	if o != nil && !IsNil(o.ValueQuantity) {
		return true
	}

	return false
}

// SetValueQuantity gets a reference to the given Quantity and assigns it to the ValueQuantity field.
func (o *ContractAnswer) SetValueQuantity(v Quantity) {
	o.ValueQuantity = &v
}

// GetValueReference returns the ValueReference field value if set, zero value otherwise.
func (o *ContractAnswer) GetValueReference() Reference {
	if o == nil || IsNil(o.ValueReference) {
		var ret Reference
		return ret
	}
	return *o.ValueReference
}

// GetValueReferenceOk returns a tuple with the ValueReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAnswer) GetValueReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ValueReference) {
		return nil, false
	}
	return o.ValueReference, true
}

// HasValueReference returns a boolean if a field has been set.
func (o *ContractAnswer) HasValueReference() bool {
	if o != nil && !IsNil(o.ValueReference) {
		return true
	}

	return false
}

// SetValueReference gets a reference to the given Reference and assigns it to the ValueReference field.
func (o *ContractAnswer) SetValueReference(v Reference) {
	o.ValueReference = &v
}

func (o ContractAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractAnswer) ToMap() (map[string]interface{}, error) {
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

type NullableContractAnswer struct {
	value *ContractAnswer
	isSet bool
}

func (v NullableContractAnswer) Get() *ContractAnswer {
	return v.value
}

func (v *NullableContractAnswer) Set(val *ContractAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullableContractAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullableContractAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractAnswer(val *ContractAnswer) *NullableContractAnswer {
	return &NullableContractAnswer{value: val, isSet: true}
}

func (v NullableContractAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


