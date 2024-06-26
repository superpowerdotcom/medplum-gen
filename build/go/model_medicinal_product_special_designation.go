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

// checks if the MedicinalProductSpecialDesignation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductSpecialDesignation{}

// MedicinalProductSpecialDesignation Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g. regulatory use).
type MedicinalProductSpecialDesignation struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifier for the designation, or procedure number.
	Identifier []Identifier `json:"identifier,omitempty"`
	// The type of special designation, e.g. orphan drug, minor use.
	Type *CodeableConcept `json:"type,omitempty"`
	// The intended use of the product, e.g. prevention, treatment.
	IntendedUse *CodeableConcept `json:"intendedUse,omitempty"`
	// Condition for which the medicinal use applies.
	IndicationCodeableConcept *CodeableConcept `json:"indicationCodeableConcept,omitempty"`
	// Condition for which the medicinal use applies.
	IndicationReference *Reference `json:"indicationReference,omitempty"`
	// For example granted, pending, expired or withdrawn.
	Status *CodeableConcept `json:"status,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Date *string `json:"date,omitempty"`
	// Animal species for which this applies.
	Species *CodeableConcept `json:"species,omitempty"`
}

// NewMedicinalProductSpecialDesignation instantiates a new MedicinalProductSpecialDesignation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductSpecialDesignation() *MedicinalProductSpecialDesignation {
	this := MedicinalProductSpecialDesignation{}
	return &this
}

// NewMedicinalProductSpecialDesignationWithDefaults instantiates a new MedicinalProductSpecialDesignation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductSpecialDesignationWithDefaults() *MedicinalProductSpecialDesignation {
	this := MedicinalProductSpecialDesignation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductSpecialDesignation) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductSpecialDesignation) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductSpecialDesignation) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *MedicinalProductSpecialDesignation) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *MedicinalProductSpecialDesignation) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetIntendedUse returns the IntendedUse field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetIntendedUse() CodeableConcept {
	if o == nil || IsNil(o.IntendedUse) {
		var ret CodeableConcept
		return ret
	}
	return *o.IntendedUse
}

// GetIntendedUseOk returns a tuple with the IntendedUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetIntendedUseOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.IntendedUse) {
		return nil, false
	}
	return o.IntendedUse, true
}

// HasIntendedUse returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasIntendedUse() bool {
	if o != nil && !IsNil(o.IntendedUse) {
		return true
	}

	return false
}

// SetIntendedUse gets a reference to the given CodeableConcept and assigns it to the IntendedUse field.
func (o *MedicinalProductSpecialDesignation) SetIntendedUse(v CodeableConcept) {
	o.IntendedUse = &v
}

// GetIndicationCodeableConcept returns the IndicationCodeableConcept field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetIndicationCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.IndicationCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.IndicationCodeableConcept
}

// GetIndicationCodeableConceptOk returns a tuple with the IndicationCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetIndicationCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.IndicationCodeableConcept) {
		return nil, false
	}
	return o.IndicationCodeableConcept, true
}

// HasIndicationCodeableConcept returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasIndicationCodeableConcept() bool {
	if o != nil && !IsNil(o.IndicationCodeableConcept) {
		return true
	}

	return false
}

// SetIndicationCodeableConcept gets a reference to the given CodeableConcept and assigns it to the IndicationCodeableConcept field.
func (o *MedicinalProductSpecialDesignation) SetIndicationCodeableConcept(v CodeableConcept) {
	o.IndicationCodeableConcept = &v
}

// GetIndicationReference returns the IndicationReference field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetIndicationReference() Reference {
	if o == nil || IsNil(o.IndicationReference) {
		var ret Reference
		return ret
	}
	return *o.IndicationReference
}

// GetIndicationReferenceOk returns a tuple with the IndicationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetIndicationReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.IndicationReference) {
		return nil, false
	}
	return o.IndicationReference, true
}

// HasIndicationReference returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasIndicationReference() bool {
	if o != nil && !IsNil(o.IndicationReference) {
		return true
	}

	return false
}

// SetIndicationReference gets a reference to the given Reference and assigns it to the IndicationReference field.
func (o *MedicinalProductSpecialDesignation) SetIndicationReference(v Reference) {
	o.IndicationReference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetStatus() CodeableConcept {
	if o == nil || IsNil(o.Status) {
		var ret CodeableConcept
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetStatusOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CodeableConcept and assigns it to the Status field.
func (o *MedicinalProductSpecialDesignation) SetStatus(v CodeableConcept) {
	o.Status = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *MedicinalProductSpecialDesignation) SetDate(v string) {
	o.Date = &v
}

// GetSpecies returns the Species field value if set, zero value otherwise.
func (o *MedicinalProductSpecialDesignation) GetSpecies() CodeableConcept {
	if o == nil || IsNil(o.Species) {
		var ret CodeableConcept
		return ret
	}
	return *o.Species
}

// GetSpeciesOk returns a tuple with the Species field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductSpecialDesignation) GetSpeciesOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Species) {
		return nil, false
	}
	return o.Species, true
}

// HasSpecies returns a boolean if a field has been set.
func (o *MedicinalProductSpecialDesignation) HasSpecies() bool {
	if o != nil && !IsNil(o.Species) {
		return true
	}

	return false
}

// SetSpecies gets a reference to the given CodeableConcept and assigns it to the Species field.
func (o *MedicinalProductSpecialDesignation) SetSpecies(v CodeableConcept) {
	o.Species = &v
}

func (o MedicinalProductSpecialDesignation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductSpecialDesignation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IntendedUse) {
		toSerialize["intendedUse"] = o.IntendedUse
	}
	if !IsNil(o.IndicationCodeableConcept) {
		toSerialize["indicationCodeableConcept"] = o.IndicationCodeableConcept
	}
	if !IsNil(o.IndicationReference) {
		toSerialize["indicationReference"] = o.IndicationReference
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Species) {
		toSerialize["species"] = o.Species
	}
	return toSerialize, nil
}

type NullableMedicinalProductSpecialDesignation struct {
	value *MedicinalProductSpecialDesignation
	isSet bool
}

func (v NullableMedicinalProductSpecialDesignation) Get() *MedicinalProductSpecialDesignation {
	return v.value
}

func (v *NullableMedicinalProductSpecialDesignation) Set(val *MedicinalProductSpecialDesignation) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductSpecialDesignation) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductSpecialDesignation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductSpecialDesignation(val *MedicinalProductSpecialDesignation) *NullableMedicinalProductSpecialDesignation {
	return &NullableMedicinalProductSpecialDesignation{value: val, isSet: true}
}

func (v NullableMedicinalProductSpecialDesignation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductSpecialDesignation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


