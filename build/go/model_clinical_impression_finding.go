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

// checks if the ClinicalImpressionFinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClinicalImpressionFinding{}

// ClinicalImpressionFinding A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called \"ClinicalImpression\" rather than \"ClinicalAssessment\" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpressionFinding struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Specific text or code for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	// Specific reference for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemReference *Reference `json:"itemReference,omitempty"`
	// A sequence of Unicode characters
	Basis *string `json:"basis,omitempty"`
}

// NewClinicalImpressionFinding instantiates a new ClinicalImpressionFinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClinicalImpressionFinding() *ClinicalImpressionFinding {
	this := ClinicalImpressionFinding{}
	return &this
}

// NewClinicalImpressionFindingWithDefaults instantiates a new ClinicalImpressionFinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClinicalImpressionFindingWithDefaults() *ClinicalImpressionFinding {
	this := ClinicalImpressionFinding{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClinicalImpressionFinding) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ClinicalImpressionFinding) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ClinicalImpressionFinding) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetItemCodeableConcept returns the ItemCodeableConcept field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetItemCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ItemCodeableConcept
}

// GetItemCodeableConceptOk returns a tuple with the ItemCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetItemCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		return nil, false
	}
	return o.ItemCodeableConcept, true
}

// HasItemCodeableConcept returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasItemCodeableConcept() bool {
	if o != nil && !IsNil(o.ItemCodeableConcept) {
		return true
	}

	return false
}

// SetItemCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ItemCodeableConcept field.
func (o *ClinicalImpressionFinding) SetItemCodeableConcept(v CodeableConcept) {
	o.ItemCodeableConcept = &v
}

// GetItemReference returns the ItemReference field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetItemReference() Reference {
	if o == nil || IsNil(o.ItemReference) {
		var ret Reference
		return ret
	}
	return *o.ItemReference
}

// GetItemReferenceOk returns a tuple with the ItemReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetItemReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ItemReference) {
		return nil, false
	}
	return o.ItemReference, true
}

// HasItemReference returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasItemReference() bool {
	if o != nil && !IsNil(o.ItemReference) {
		return true
	}

	return false
}

// SetItemReference gets a reference to the given Reference and assigns it to the ItemReference field.
func (o *ClinicalImpressionFinding) SetItemReference(v Reference) {
	o.ItemReference = &v
}

// GetBasis returns the Basis field value if set, zero value otherwise.
func (o *ClinicalImpressionFinding) GetBasis() string {
	if o == nil || IsNil(o.Basis) {
		var ret string
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinicalImpressionFinding) GetBasisOk() (*string, bool) {
	if o == nil || IsNil(o.Basis) {
		return nil, false
	}
	return o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *ClinicalImpressionFinding) HasBasis() bool {
	if o != nil && !IsNil(o.Basis) {
		return true
	}

	return false
}

// SetBasis gets a reference to the given string and assigns it to the Basis field.
func (o *ClinicalImpressionFinding) SetBasis(v string) {
	o.Basis = &v
}

func (o ClinicalImpressionFinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClinicalImpressionFinding) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Basis) {
		toSerialize["basis"] = o.Basis
	}
	return toSerialize, nil
}

type NullableClinicalImpressionFinding struct {
	value *ClinicalImpressionFinding
	isSet bool
}

func (v NullableClinicalImpressionFinding) Get() *ClinicalImpressionFinding {
	return v.value
}

func (v *NullableClinicalImpressionFinding) Set(val *ClinicalImpressionFinding) {
	v.value = val
	v.isSet = true
}

func (v NullableClinicalImpressionFinding) IsSet() bool {
	return v.isSet
}

func (v *NullableClinicalImpressionFinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClinicalImpressionFinding(val *ClinicalImpressionFinding) *NullableClinicalImpressionFinding {
	return &NullableClinicalImpressionFinding{value: val, isSet: true}
}

func (v NullableClinicalImpressionFinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClinicalImpressionFinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


