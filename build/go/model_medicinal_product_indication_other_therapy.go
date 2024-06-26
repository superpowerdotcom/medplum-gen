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

// checks if the MedicinalProductIndicationOtherTherapy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductIndicationOtherTherapy{}

// MedicinalProductIndicationOtherTherapy Indication for the Medicinal Product.
type MedicinalProductIndicationOtherTherapy struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of relationship between the medicinal product indication or contraindication and another therapy.
	TherapyRelationshipType CodeableConcept `json:"therapyRelationshipType"`
	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication.
	MedicationReference *Reference `json:"medicationReference,omitempty"`
}

type _MedicinalProductIndicationOtherTherapy MedicinalProductIndicationOtherTherapy

// NewMedicinalProductIndicationOtherTherapy instantiates a new MedicinalProductIndicationOtherTherapy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductIndicationOtherTherapy(therapyRelationshipType CodeableConcept) *MedicinalProductIndicationOtherTherapy {
	this := MedicinalProductIndicationOtherTherapy{}
	this.TherapyRelationshipType = therapyRelationshipType
	return &this
}

// NewMedicinalProductIndicationOtherTherapyWithDefaults instantiates a new MedicinalProductIndicationOtherTherapy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductIndicationOtherTherapyWithDefaults() *MedicinalProductIndicationOtherTherapy {
	this := MedicinalProductIndicationOtherTherapy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductIndicationOtherTherapy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductIndicationOtherTherapy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductIndicationOtherTherapy) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductIndicationOtherTherapy) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductIndicationOtherTherapy) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductIndicationOtherTherapy) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductIndicationOtherTherapy) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductIndicationOtherTherapy) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductIndicationOtherTherapy) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetTherapyRelationshipType returns the TherapyRelationshipType field value
func (o *MedicinalProductIndicationOtherTherapy) GetTherapyRelationshipType() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.TherapyRelationshipType
}

// GetTherapyRelationshipTypeOk returns a tuple with the TherapyRelationshipType field value
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetTherapyRelationshipTypeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TherapyRelationshipType, true
}

// SetTherapyRelationshipType sets field value
func (o *MedicinalProductIndicationOtherTherapy) SetTherapyRelationshipType(v CodeableConcept) {
	o.TherapyRelationshipType = v
}

// GetMedicationCodeableConcept returns the MedicationCodeableConcept field value if set, zero value otherwise.
func (o *MedicinalProductIndicationOtherTherapy) GetMedicationCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.MedicationCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.MedicationCodeableConcept
}

// GetMedicationCodeableConceptOk returns a tuple with the MedicationCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetMedicationCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.MedicationCodeableConcept) {
		return nil, false
	}
	return o.MedicationCodeableConcept, true
}

// HasMedicationCodeableConcept returns a boolean if a field has been set.
func (o *MedicinalProductIndicationOtherTherapy) HasMedicationCodeableConcept() bool {
	if o != nil && !IsNil(o.MedicationCodeableConcept) {
		return true
	}

	return false
}

// SetMedicationCodeableConcept gets a reference to the given CodeableConcept and assigns it to the MedicationCodeableConcept field.
func (o *MedicinalProductIndicationOtherTherapy) SetMedicationCodeableConcept(v CodeableConcept) {
	o.MedicationCodeableConcept = &v
}

// GetMedicationReference returns the MedicationReference field value if set, zero value otherwise.
func (o *MedicinalProductIndicationOtherTherapy) GetMedicationReference() Reference {
	if o == nil || IsNil(o.MedicationReference) {
		var ret Reference
		return ret
	}
	return *o.MedicationReference
}

// GetMedicationReferenceOk returns a tuple with the MedicationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIndicationOtherTherapy) GetMedicationReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.MedicationReference) {
		return nil, false
	}
	return o.MedicationReference, true
}

// HasMedicationReference returns a boolean if a field has been set.
func (o *MedicinalProductIndicationOtherTherapy) HasMedicationReference() bool {
	if o != nil && !IsNil(o.MedicationReference) {
		return true
	}

	return false
}

// SetMedicationReference gets a reference to the given Reference and assigns it to the MedicationReference field.
func (o *MedicinalProductIndicationOtherTherapy) SetMedicationReference(v Reference) {
	o.MedicationReference = &v
}

func (o MedicinalProductIndicationOtherTherapy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductIndicationOtherTherapy) ToMap() (map[string]interface{}, error) {
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
	toSerialize["therapyRelationshipType"] = o.TherapyRelationshipType
	if !IsNil(o.MedicationCodeableConcept) {
		toSerialize["medicationCodeableConcept"] = o.MedicationCodeableConcept
	}
	if !IsNil(o.MedicationReference) {
		toSerialize["medicationReference"] = o.MedicationReference
	}
	return toSerialize, nil
}

func (o *MedicinalProductIndicationOtherTherapy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"therapyRelationshipType",
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

	varMedicinalProductIndicationOtherTherapy := _MedicinalProductIndicationOtherTherapy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicinalProductIndicationOtherTherapy)

	if err != nil {
		return err
	}

	*o = MedicinalProductIndicationOtherTherapy(varMedicinalProductIndicationOtherTherapy)

	return err
}

type NullableMedicinalProductIndicationOtherTherapy struct {
	value *MedicinalProductIndicationOtherTherapy
	isSet bool
}

func (v NullableMedicinalProductIndicationOtherTherapy) Get() *MedicinalProductIndicationOtherTherapy {
	return v.value
}

func (v *NullableMedicinalProductIndicationOtherTherapy) Set(val *MedicinalProductIndicationOtherTherapy) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductIndicationOtherTherapy) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductIndicationOtherTherapy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductIndicationOtherTherapy(val *MedicinalProductIndicationOtherTherapy) *NullableMedicinalProductIndicationOtherTherapy {
	return &NullableMedicinalProductIndicationOtherTherapy{value: val, isSet: true}
}

func (v NullableMedicinalProductIndicationOtherTherapy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductIndicationOtherTherapy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

