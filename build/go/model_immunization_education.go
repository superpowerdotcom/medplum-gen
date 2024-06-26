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

// checks if the ImmunizationEducation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImmunizationEducation{}

// ImmunizationEducation Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
type ImmunizationEducation struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	DocumentType *string `json:"documentType,omitempty"`
	// String of characters used to identify a name or a resource
	Reference *string `json:"reference,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	PublicationDate *string `json:"publicationDate,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	PresentationDate *string `json:"presentationDate,omitempty"`
}

// NewImmunizationEducation instantiates a new ImmunizationEducation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImmunizationEducation() *ImmunizationEducation {
	this := ImmunizationEducation{}
	return &this
}

// NewImmunizationEducationWithDefaults instantiates a new ImmunizationEducation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImmunizationEducationWithDefaults() *ImmunizationEducation {
	this := ImmunizationEducation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImmunizationEducation) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ImmunizationEducation) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ImmunizationEducation) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *ImmunizationEducation) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ImmunizationEducation) SetReference(v string) {
	o.Reference = &v
}

// GetPublicationDate returns the PublicationDate field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetPublicationDate() string {
	if o == nil || IsNil(o.PublicationDate) {
		var ret string
		return ret
	}
	return *o.PublicationDate
}

// GetPublicationDateOk returns a tuple with the PublicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetPublicationDateOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationDate) {
		return nil, false
	}
	return o.PublicationDate, true
}

// HasPublicationDate returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasPublicationDate() bool {
	if o != nil && !IsNil(o.PublicationDate) {
		return true
	}

	return false
}

// SetPublicationDate gets a reference to the given string and assigns it to the PublicationDate field.
func (o *ImmunizationEducation) SetPublicationDate(v string) {
	o.PublicationDate = &v
}

// GetPresentationDate returns the PresentationDate field value if set, zero value otherwise.
func (o *ImmunizationEducation) GetPresentationDate() string {
	if o == nil || IsNil(o.PresentationDate) {
		var ret string
		return ret
	}
	return *o.PresentationDate
}

// GetPresentationDateOk returns a tuple with the PresentationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImmunizationEducation) GetPresentationDateOk() (*string, bool) {
	if o == nil || IsNil(o.PresentationDate) {
		return nil, false
	}
	return o.PresentationDate, true
}

// HasPresentationDate returns a boolean if a field has been set.
func (o *ImmunizationEducation) HasPresentationDate() bool {
	if o != nil && !IsNil(o.PresentationDate) {
		return true
	}

	return false
}

// SetPresentationDate gets a reference to the given string and assigns it to the PresentationDate field.
func (o *ImmunizationEducation) SetPresentationDate(v string) {
	o.PresentationDate = &v
}

func (o ImmunizationEducation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImmunizationEducation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.PublicationDate) {
		toSerialize["publicationDate"] = o.PublicationDate
	}
	if !IsNil(o.PresentationDate) {
		toSerialize["presentationDate"] = o.PresentationDate
	}
	return toSerialize, nil
}

type NullableImmunizationEducation struct {
	value *ImmunizationEducation
	isSet bool
}

func (v NullableImmunizationEducation) Get() *ImmunizationEducation {
	return v.value
}

func (v *NullableImmunizationEducation) Set(val *ImmunizationEducation) {
	v.value = val
	v.isSet = true
}

func (v NullableImmunizationEducation) IsSet() bool {
	return v.isSet
}

func (v *NullableImmunizationEducation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImmunizationEducation(val *ImmunizationEducation) *NullableImmunizationEducation {
	return &NullableImmunizationEducation{value: val, isSet: true}
}

func (v NullableImmunizationEducation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImmunizationEducation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


