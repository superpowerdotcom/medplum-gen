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

// checks if the DocumentReferenceContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentReferenceContent{}

// DocumentReferenceContent A reference to a document of any kind for any purpose. Provides metadata about the document so that the document can be discovered and managed. The scope of a document is any seralized object with a mime-type, so includes formal patient centric documents (CDA), cliical notes, scanned paper, and non-patient specific documents like policy text.
type DocumentReferenceContent struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The document or URL of the document along with critical metadata to prove content has integrity.
	Attachment Attachment `json:"attachment"`
	// An identifier of the document encoding, structure, and template that the document conforms to beyond the base format indicated in the mimeType.
	Format *Coding `json:"format,omitempty"`
}

type _DocumentReferenceContent DocumentReferenceContent

// NewDocumentReferenceContent instantiates a new DocumentReferenceContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentReferenceContent(attachment Attachment) *DocumentReferenceContent {
	this := DocumentReferenceContent{}
	this.Attachment = attachment
	return &this
}

// NewDocumentReferenceContentWithDefaults instantiates a new DocumentReferenceContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentReferenceContentWithDefaults() *DocumentReferenceContent {
	this := DocumentReferenceContent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentReferenceContent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentReferenceContent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentReferenceContent) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DocumentReferenceContent) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContent) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DocumentReferenceContent) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DocumentReferenceContent) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DocumentReferenceContent) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContent) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DocumentReferenceContent) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DocumentReferenceContent) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetAttachment returns the Attachment field value
func (o *DocumentReferenceContent) GetAttachment() Attachment {
	if o == nil {
		var ret Attachment
		return ret
	}

	return o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContent) GetAttachmentOk() (*Attachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachment, true
}

// SetAttachment sets field value
func (o *DocumentReferenceContent) SetAttachment(v Attachment) {
	o.Attachment = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DocumentReferenceContent) GetFormat() Coding {
	if o == nil || IsNil(o.Format) {
		var ret Coding
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferenceContent) GetFormatOk() (*Coding, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DocumentReferenceContent) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given Coding and assigns it to the Format field.
func (o *DocumentReferenceContent) SetFormat(v Coding) {
	o.Format = &v
}

func (o DocumentReferenceContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentReferenceContent) ToMap() (map[string]interface{}, error) {
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
	toSerialize["attachment"] = o.Attachment
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	return toSerialize, nil
}

func (o *DocumentReferenceContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachment",
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

	varDocumentReferenceContent := _DocumentReferenceContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentReferenceContent)

	if err != nil {
		return err
	}

	*o = DocumentReferenceContent(varDocumentReferenceContent)

	return err
}

type NullableDocumentReferenceContent struct {
	value *DocumentReferenceContent
	isSet bool
}

func (v NullableDocumentReferenceContent) Get() *DocumentReferenceContent {
	return v.value
}

func (v *NullableDocumentReferenceContent) Set(val *DocumentReferenceContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentReferenceContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentReferenceContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentReferenceContent(val *DocumentReferenceContent) *NullableDocumentReferenceContent {
	return &NullableDocumentReferenceContent{value: val, isSet: true}
}

func (v NullableDocumentReferenceContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentReferenceContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


