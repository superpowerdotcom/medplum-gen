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

// checks if the MedicationKnowledgeIngredient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationKnowledgeIngredient{}

// MedicationKnowledgeIngredient Information about a medication that is used to support knowledge.
type MedicationKnowledgeIngredient struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual ingredient - either a substance (simple ingredient) or another medication.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	// The actual ingredient - either a substance (simple ingredient) or another medication.
	ItemReference *Reference `json:"itemReference,omitempty"`
	// Value of \"true\" or \"false\"
	IsActive *bool `json:"isActive,omitempty"`
	// Specifies how many (or how much) of the items there are in this Medication.  For example, 250 mg per tablet.  This is expressed as a ratio where the numerator is 250mg and the denominator is 1 tablet.
	Strength *Ratio `json:"strength,omitempty"`
}

// NewMedicationKnowledgeIngredient instantiates a new MedicationKnowledgeIngredient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationKnowledgeIngredient() *MedicationKnowledgeIngredient {
	this := MedicationKnowledgeIngredient{}
	return &this
}

// NewMedicationKnowledgeIngredientWithDefaults instantiates a new MedicationKnowledgeIngredient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationKnowledgeIngredientWithDefaults() *MedicationKnowledgeIngredient {
	this := MedicationKnowledgeIngredient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationKnowledgeIngredient) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationKnowledgeIngredient) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationKnowledgeIngredient) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetItemCodeableConcept returns the ItemCodeableConcept field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetItemCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ItemCodeableConcept
}

// GetItemCodeableConceptOk returns a tuple with the ItemCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetItemCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		return nil, false
	}
	return o.ItemCodeableConcept, true
}

// HasItemCodeableConcept returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasItemCodeableConcept() bool {
	if o != nil && !IsNil(o.ItemCodeableConcept) {
		return true
	}

	return false
}

// SetItemCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ItemCodeableConcept field.
func (o *MedicationKnowledgeIngredient) SetItemCodeableConcept(v CodeableConcept) {
	o.ItemCodeableConcept = &v
}

// GetItemReference returns the ItemReference field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetItemReference() Reference {
	if o == nil || IsNil(o.ItemReference) {
		var ret Reference
		return ret
	}
	return *o.ItemReference
}

// GetItemReferenceOk returns a tuple with the ItemReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetItemReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ItemReference) {
		return nil, false
	}
	return o.ItemReference, true
}

// HasItemReference returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasItemReference() bool {
	if o != nil && !IsNil(o.ItemReference) {
		return true
	}

	return false
}

// SetItemReference gets a reference to the given Reference and assigns it to the ItemReference field.
func (o *MedicationKnowledgeIngredient) SetItemReference(v Reference) {
	o.ItemReference = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *MedicationKnowledgeIngredient) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *MedicationKnowledgeIngredient) GetStrength() Ratio {
	if o == nil || IsNil(o.Strength) {
		var ret Ratio
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationKnowledgeIngredient) GetStrengthOk() (*Ratio, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *MedicationKnowledgeIngredient) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given Ratio and assigns it to the Strength field.
func (o *MedicationKnowledgeIngredient) SetStrength(v Ratio) {
	o.Strength = &v
}

func (o MedicationKnowledgeIngredient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationKnowledgeIngredient) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ItemCodeableConcept) {
		toSerialize["itemCodeableConcept"] = o.ItemCodeableConcept
	}
	if !IsNil(o.ItemReference) {
		toSerialize["itemReference"] = o.ItemReference
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	return toSerialize, nil
}

type NullableMedicationKnowledgeIngredient struct {
	value *MedicationKnowledgeIngredient
	isSet bool
}

func (v NullableMedicationKnowledgeIngredient) Get() *MedicationKnowledgeIngredient {
	return v.value
}

func (v *NullableMedicationKnowledgeIngredient) Set(val *MedicationKnowledgeIngredient) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationKnowledgeIngredient) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationKnowledgeIngredient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationKnowledgeIngredient(val *MedicationKnowledgeIngredient) *NullableMedicationKnowledgeIngredient {
	return &NullableMedicationKnowledgeIngredient{value: val, isSet: true}
}

func (v NullableMedicationKnowledgeIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationKnowledgeIngredient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


