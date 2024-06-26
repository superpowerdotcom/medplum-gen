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

// checks if the MedicationKnowledgeMedicineClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationKnowledgeMedicineClassification{}

// MedicationKnowledgeMedicineClassification Information about a medication that is used to support knowledge.
type MedicationKnowledgeMedicineClassification struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of category for the medication (for example, therapeutic classification, therapeutic sub-classification).
	Type CodeableConcept `json:"type"`
	// Specific category assigned to the medication (e.g. anti-infective, anti-hypertensive, antibiotic, etc.).
	Classification []CodeableConcept `json:"classification,omitempty"`
}

type _MedicationKnowledgeMedicineClassification MedicationKnowledgeMedicineClassification

// NewMedicationKnowledgeMedicineClassification instantiates a new MedicationKnowledgeMedicineClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationKnowledgeMedicineClassification(type_ CodeableConcept) *MedicationKnowledgeMedicineClassification {
	this := MedicationKnowledgeMedicineClassification{}
	this.Type = type_
	return &this
}

// NewMedicationKnowledgeMedicineClassificationWithDefaults instantiates a new MedicationKnowledgeMedicineClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationKnowledgeMedicineClassificationWithDefaults() *MedicationKnowledgeMedicineClassification {
	this := MedicationKnowledgeMedicineClassification{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationKnowledgeMedicineClassification) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeMedicineClassification) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationKnowledgeMedicineClassification) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationKnowledgeMedicineClassification) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationKnowledgeMedicineClassification) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeMedicineClassification) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeMedicineClassification) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationKnowledgeMedicineClassification) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationKnowledgeMedicineClassification) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeMedicineClassification) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeMedicineClassification) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationKnowledgeMedicineClassification) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value
func (o *MedicationKnowledgeMedicineClassification) GetType() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeMedicineClassification) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MedicationKnowledgeMedicineClassification) SetType(v CodeableConcept) {
	o.Type = v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *MedicationKnowledgeMedicineClassification) GetClassification() []CodeableConcept {
	if o == nil || IsNil(o.Classification) {
		var ret []CodeableConcept
		return ret
	}
	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeMedicineClassification) GetClassificationOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *MedicationKnowledgeMedicineClassification) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given []CodeableConcept and assigns it to the Classification field.
func (o *MedicationKnowledgeMedicineClassification) SetClassification(v []CodeableConcept) {
	o.Classification = v
}

func (o MedicationKnowledgeMedicineClassification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationKnowledgeMedicineClassification) ToMap() (map[string]interface{}, error) {
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
	toSerialize["type"] = o.Type
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	return toSerialize, nil
}

func (o *MedicationKnowledgeMedicineClassification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varMedicationKnowledgeMedicineClassification := _MedicationKnowledgeMedicineClassification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicationKnowledgeMedicineClassification)

	if err != nil {
		return err
	}

	*o = MedicationKnowledgeMedicineClassification(varMedicationKnowledgeMedicineClassification)

	return err
}

type NullableMedicationKnowledgeMedicineClassification struct {
	value *MedicationKnowledgeMedicineClassification
	isSet bool
}

func (v NullableMedicationKnowledgeMedicineClassification) Get() *MedicationKnowledgeMedicineClassification {
	return v.value
}

func (v *NullableMedicationKnowledgeMedicineClassification) Set(val *MedicationKnowledgeMedicineClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationKnowledgeMedicineClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationKnowledgeMedicineClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationKnowledgeMedicineClassification(val *MedicationKnowledgeMedicineClassification) *NullableMedicationKnowledgeMedicineClassification {
	return &NullableMedicationKnowledgeMedicineClassification{value: val, isSet: true}
}

func (v NullableMedicationKnowledgeMedicineClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationKnowledgeMedicineClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


