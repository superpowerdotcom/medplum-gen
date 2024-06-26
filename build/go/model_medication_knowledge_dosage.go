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

// checks if the MedicationKnowledgeDosage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationKnowledgeDosage{}

// MedicationKnowledgeDosage Information about a medication that is used to support knowledge.
type MedicationKnowledgeDosage struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of dosage (for example, prophylaxis, maintenance, therapeutic, etc.).
	Type CodeableConcept `json:"type"`
	// Dosage for the medication for the specific guidelines.
	Dosage []Dosage `json:"dosage"`
}

type _MedicationKnowledgeDosage MedicationKnowledgeDosage

// NewMedicationKnowledgeDosage instantiates a new MedicationKnowledgeDosage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationKnowledgeDosage(type_ CodeableConcept, dosage []Dosage) *MedicationKnowledgeDosage {
	this := MedicationKnowledgeDosage{}
	this.Type = type_
	this.Dosage = dosage
	return &this
}

// NewMedicationKnowledgeDosageWithDefaults instantiates a new MedicationKnowledgeDosage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationKnowledgeDosageWithDefaults() *MedicationKnowledgeDosage {
	this := MedicationKnowledgeDosage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationKnowledgeDosage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeDosage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationKnowledgeDosage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationKnowledgeDosage) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationKnowledgeDosage) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeDosage) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeDosage) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationKnowledgeDosage) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationKnowledgeDosage) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeDosage) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeDosage) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationKnowledgeDosage) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value
func (o *MedicationKnowledgeDosage) GetType() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeDosage) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MedicationKnowledgeDosage) SetType(v CodeableConcept) {
	o.Type = v
}

// GetDosage returns the Dosage field value
func (o *MedicationKnowledgeDosage) GetDosage() []Dosage {
	if o == nil {
		var ret []Dosage
		return ret
	}

	return o.Dosage
}

// GetDosageOk returns a tuple with the Dosage field value
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeDosage) GetDosageOk() ([]Dosage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dosage, true
}

// SetDosage sets field value
func (o *MedicationKnowledgeDosage) SetDosage(v []Dosage) {
	o.Dosage = v
}

func (o MedicationKnowledgeDosage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationKnowledgeDosage) ToMap() (map[string]interface{}, error) {
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
	toSerialize["dosage"] = o.Dosage
	return toSerialize, nil
}

func (o *MedicationKnowledgeDosage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"dosage",
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

	varMedicationKnowledgeDosage := _MedicationKnowledgeDosage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicationKnowledgeDosage)

	if err != nil {
		return err
	}

	*o = MedicationKnowledgeDosage(varMedicationKnowledgeDosage)

	return err
}

type NullableMedicationKnowledgeDosage struct {
	value *MedicationKnowledgeDosage
	isSet bool
}

func (v NullableMedicationKnowledgeDosage) Get() *MedicationKnowledgeDosage {
	return v.value
}

func (v *NullableMedicationKnowledgeDosage) Set(val *MedicationKnowledgeDosage) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationKnowledgeDosage) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationKnowledgeDosage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationKnowledgeDosage(val *MedicationKnowledgeDosage) *NullableMedicationKnowledgeDosage {
	return &NullableMedicationKnowledgeDosage{value: val, isSet: true}
}

func (v NullableMedicationKnowledgeDosage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationKnowledgeDosage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


