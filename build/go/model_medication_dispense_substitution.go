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

// checks if the MedicationDispenseSubstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicationDispenseSubstitution{}

// MedicationDispenseSubstitution Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispenseSubstitution struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Value of \"true\" or \"false\"
	WasSubstituted *bool `json:"wasSubstituted,omitempty"`
	// A code signifying whether a different drug was dispensed from what was prescribed.
	Type *CodeableConcept `json:"type,omitempty"`
	// Indicates the reason for the substitution (or lack of substitution) from what was prescribed.
	Reason []CodeableConcept `json:"reason,omitempty"`
	// The person or organization that has primary responsibility for the substitution.
	ResponsibleParty []Reference `json:"responsibleParty,omitempty"`
}

// NewMedicationDispenseSubstitution instantiates a new MedicationDispenseSubstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicationDispenseSubstitution() *MedicationDispenseSubstitution {
	this := MedicationDispenseSubstitution{}
	return &this
}

// NewMedicationDispenseSubstitutionWithDefaults instantiates a new MedicationDispenseSubstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicationDispenseSubstitutionWithDefaults() *MedicationDispenseSubstitution {
	this := MedicationDispenseSubstitution{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicationDispenseSubstitution) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicationDispenseSubstitution) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicationDispenseSubstitution) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetWasSubstituted returns the WasSubstituted field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetWasSubstituted() bool {
	if o == nil || IsNil(o.WasSubstituted) {
		var ret bool
		return ret
	}
	return *o.WasSubstituted
}

// GetWasSubstitutedOk returns a tuple with the WasSubstituted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetWasSubstitutedOk() (*bool, bool) {
	if o == nil || IsNil(o.WasSubstituted) {
		return nil, false
	}
	return o.WasSubstituted, true
}

// HasWasSubstituted returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasWasSubstituted() bool {
	if o != nil && !IsNil(o.WasSubstituted) {
		return true
	}

	return false
}

// SetWasSubstituted gets a reference to the given bool and assigns it to the WasSubstituted field.
func (o *MedicationDispenseSubstitution) SetWasSubstituted(v bool) {
	o.WasSubstituted = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *MedicationDispenseSubstitution) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetReason() []CodeableConcept {
	if o == nil || IsNil(o.Reason) {
		var ret []CodeableConcept
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetReasonOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given []CodeableConcept and assigns it to the Reason field.
func (o *MedicationDispenseSubstitution) SetReason(v []CodeableConcept) {
	o.Reason = v
}

// GetResponsibleParty returns the ResponsibleParty field value if set, zero value otherwise.
func (o *MedicationDispenseSubstitution) GetResponsibleParty() []Reference {
	if o == nil || IsNil(o.ResponsibleParty) {
		var ret []Reference
		return ret
	}
	return o.ResponsibleParty
}

// GetResponsiblePartyOk returns a tuple with the ResponsibleParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicationDispenseSubstitution) GetResponsiblePartyOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ResponsibleParty) {
		return nil, false
	}
	return o.ResponsibleParty, true
}

// HasResponsibleParty returns a boolean if a field has been set.
func (o *MedicationDispenseSubstitution) HasResponsibleParty() bool {
	if o != nil && !IsNil(o.ResponsibleParty) {
		return true
	}

	return false
}

// SetResponsibleParty gets a reference to the given []Reference and assigns it to the ResponsibleParty field.
func (o *MedicationDispenseSubstitution) SetResponsibleParty(v []Reference) {
	o.ResponsibleParty = v
}

func (o MedicationDispenseSubstitution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicationDispenseSubstitution) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.WasSubstituted) {
		toSerialize["wasSubstituted"] = o.WasSubstituted
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ResponsibleParty) {
		toSerialize["responsibleParty"] = o.ResponsibleParty
	}
	return toSerialize, nil
}

type NullableMedicationDispenseSubstitution struct {
	value *MedicationDispenseSubstitution
	isSet bool
}

func (v NullableMedicationDispenseSubstitution) Get() *MedicationDispenseSubstitution {
	return v.value
}

func (v *NullableMedicationDispenseSubstitution) Set(val *MedicationDispenseSubstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicationDispenseSubstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicationDispenseSubstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicationDispenseSubstitution(val *MedicationDispenseSubstitution) *NullableMedicationDispenseSubstitution {
	return &NullableMedicationDispenseSubstitution{value: val, isSet: true}
}

func (v NullableMedicationDispenseSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicationDispenseSubstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


