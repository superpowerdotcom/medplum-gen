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

// checks if the SubstanceSpecificationMoiety type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstanceSpecificationMoiety{}

// SubstanceSpecificationMoiety The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecificationMoiety struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Role that the moiety is playing.
	Role *CodeableConcept `json:"role,omitempty"`
	// Identifier by which this moiety substance is known.
	Identifier *Identifier `json:"identifier,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// Stereochemistry type.
	Stereochemistry *CodeableConcept `json:"stereochemistry,omitempty"`
	// Optical activity type.
	OpticalActivity *CodeableConcept `json:"opticalActivity,omitempty"`
	// A sequence of Unicode characters
	MolecularFormula *string `json:"molecularFormula,omitempty"`
	// Quantitative value for this moiety.
	AmountQuantity *Quantity `json:"amountQuantity,omitempty"`
	// Quantitative value for this moiety.
	AmountString *string `json:"amountString,omitempty"`
}

// NewSubstanceSpecificationMoiety instantiates a new SubstanceSpecificationMoiety object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstanceSpecificationMoiety() *SubstanceSpecificationMoiety {
	this := SubstanceSpecificationMoiety{}
	return &this
}

// NewSubstanceSpecificationMoietyWithDefaults instantiates a new SubstanceSpecificationMoiety object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceSpecificationMoietyWithDefaults() *SubstanceSpecificationMoiety {
	this := SubstanceSpecificationMoiety{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstanceSpecificationMoiety) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstanceSpecificationMoiety) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstanceSpecificationMoiety) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetRole() CodeableConcept {
	if o == nil || IsNil(o.Role) {
		var ret CodeableConcept
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetRoleOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given CodeableConcept and assigns it to the Role field.
func (o *SubstanceSpecificationMoiety) SetRole(v CodeableConcept) {
	o.Role = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *SubstanceSpecificationMoiety) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubstanceSpecificationMoiety) SetName(v string) {
	o.Name = &v
}

// GetStereochemistry returns the Stereochemistry field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetStereochemistry() CodeableConcept {
	if o == nil || IsNil(o.Stereochemistry) {
		var ret CodeableConcept
		return ret
	}
	return *o.Stereochemistry
}

// GetStereochemistryOk returns a tuple with the Stereochemistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetStereochemistryOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Stereochemistry) {
		return nil, false
	}
	return o.Stereochemistry, true
}

// HasStereochemistry returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasStereochemistry() bool {
	if o != nil && !IsNil(o.Stereochemistry) {
		return true
	}

	return false
}

// SetStereochemistry gets a reference to the given CodeableConcept and assigns it to the Stereochemistry field.
func (o *SubstanceSpecificationMoiety) SetStereochemistry(v CodeableConcept) {
	o.Stereochemistry = &v
}

// GetOpticalActivity returns the OpticalActivity field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetOpticalActivity() CodeableConcept {
	if o == nil || IsNil(o.OpticalActivity) {
		var ret CodeableConcept
		return ret
	}
	return *o.OpticalActivity
}

// GetOpticalActivityOk returns a tuple with the OpticalActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetOpticalActivityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.OpticalActivity) {
		return nil, false
	}
	return o.OpticalActivity, true
}

// HasOpticalActivity returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasOpticalActivity() bool {
	if o != nil && !IsNil(o.OpticalActivity) {
		return true
	}

	return false
}

// SetOpticalActivity gets a reference to the given CodeableConcept and assigns it to the OpticalActivity field.
func (o *SubstanceSpecificationMoiety) SetOpticalActivity(v CodeableConcept) {
	o.OpticalActivity = &v
}

// GetMolecularFormula returns the MolecularFormula field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetMolecularFormula() string {
	if o == nil || IsNil(o.MolecularFormula) {
		var ret string
		return ret
	}
	return *o.MolecularFormula
}

// GetMolecularFormulaOk returns a tuple with the MolecularFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetMolecularFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.MolecularFormula) {
		return nil, false
	}
	return o.MolecularFormula, true
}

// HasMolecularFormula returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasMolecularFormula() bool {
	if o != nil && !IsNil(o.MolecularFormula) {
		return true
	}

	return false
}

// SetMolecularFormula gets a reference to the given string and assigns it to the MolecularFormula field.
func (o *SubstanceSpecificationMoiety) SetMolecularFormula(v string) {
	o.MolecularFormula = &v
}

// GetAmountQuantity returns the AmountQuantity field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetAmountQuantity() Quantity {
	if o == nil || IsNil(o.AmountQuantity) {
		var ret Quantity
		return ret
	}
	return *o.AmountQuantity
}

// GetAmountQuantityOk returns a tuple with the AmountQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetAmountQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.AmountQuantity) {
		return nil, false
	}
	return o.AmountQuantity, true
}

// HasAmountQuantity returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasAmountQuantity() bool {
	if o != nil && !IsNil(o.AmountQuantity) {
		return true
	}

	return false
}

// SetAmountQuantity gets a reference to the given Quantity and assigns it to the AmountQuantity field.
func (o *SubstanceSpecificationMoiety) SetAmountQuantity(v Quantity) {
	o.AmountQuantity = &v
}

// GetAmountString returns the AmountString field value if set, zero value otherwise.
func (o *SubstanceSpecificationMoiety) GetAmountString() string {
	if o == nil || IsNil(o.AmountString) {
		var ret string
		return ret
	}
	return *o.AmountString
}

// GetAmountStringOk returns a tuple with the AmountString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationMoiety) GetAmountStringOk() (*string, bool) {
	if o == nil || IsNil(o.AmountString) {
		return nil, false
	}
	return o.AmountString, true
}

// HasAmountString returns a boolean if a field has been set.
func (o *SubstanceSpecificationMoiety) HasAmountString() bool {
	if o != nil && !IsNil(o.AmountString) {
		return true
	}

	return false
}

// SetAmountString gets a reference to the given string and assigns it to the AmountString field.
func (o *SubstanceSpecificationMoiety) SetAmountString(v string) {
	o.AmountString = &v
}

func (o SubstanceSpecificationMoiety) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstanceSpecificationMoiety) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Stereochemistry) {
		toSerialize["stereochemistry"] = o.Stereochemistry
	}
	if !IsNil(o.OpticalActivity) {
		toSerialize["opticalActivity"] = o.OpticalActivity
	}
	if !IsNil(o.MolecularFormula) {
		toSerialize["molecularFormula"] = o.MolecularFormula
	}
	if !IsNil(o.AmountQuantity) {
		toSerialize["amountQuantity"] = o.AmountQuantity
	}
	if !IsNil(o.AmountString) {
		toSerialize["amountString"] = o.AmountString
	}
	return toSerialize, nil
}

type NullableSubstanceSpecificationMoiety struct {
	value *SubstanceSpecificationMoiety
	isSet bool
}

func (v NullableSubstanceSpecificationMoiety) Get() *SubstanceSpecificationMoiety {
	return v.value
}

func (v *NullableSubstanceSpecificationMoiety) Set(val *SubstanceSpecificationMoiety) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstanceSpecificationMoiety) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstanceSpecificationMoiety) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstanceSpecificationMoiety(val *SubstanceSpecificationMoiety) *NullableSubstanceSpecificationMoiety {
	return &NullableSubstanceSpecificationMoiety{value: val, isSet: true}
}

func (v NullableSubstanceSpecificationMoiety) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstanceSpecificationMoiety) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


