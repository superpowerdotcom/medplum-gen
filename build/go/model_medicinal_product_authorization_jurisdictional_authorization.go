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

// checks if the MedicinalProductAuthorizationJurisdictionalAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductAuthorizationJurisdictionalAuthorization{}

// MedicinalProductAuthorizationJurisdictionalAuthorization The regulatory authorization of a medicinal product.
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The assigned number for the marketing authorization.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Country of authorization.
	Country *CodeableConcept `json:"country,omitempty"`
	// Jurisdiction within a country.
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// The legal status of supply in a jurisdiction or region.
	LegalStatusOfSupply *CodeableConcept `json:"legalStatusOfSupply,omitempty"`
	// The start and expected end date of the authorization.
	ValidityPeriod *Period `json:"validityPeriod,omitempty"`
}

// NewMedicinalProductAuthorizationJurisdictionalAuthorization instantiates a new MedicinalProductAuthorizationJurisdictionalAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductAuthorizationJurisdictionalAuthorization() *MedicinalProductAuthorizationJurisdictionalAuthorization {
	this := MedicinalProductAuthorizationJurisdictionalAuthorization{}
	return &this
}

// NewMedicinalProductAuthorizationJurisdictionalAuthorizationWithDefaults instantiates a new MedicinalProductAuthorizationJurisdictionalAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductAuthorizationJurisdictionalAuthorizationWithDefaults() *MedicinalProductAuthorizationJurisdictionalAuthorization {
	this := MedicinalProductAuthorizationJurisdictionalAuthorization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetCountry() CodeableConcept {
	if o == nil || IsNil(o.Country) {
		var ret CodeableConcept
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetCountryOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given CodeableConcept and assigns it to the Country field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetCountry(v CodeableConcept) {
	o.Country = &v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetJurisdiction() []CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret []CodeableConcept
		return ret
	}
	return o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetJurisdictionOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given []CodeableConcept and assigns it to the Jurisdiction field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetJurisdiction(v []CodeableConcept) {
	o.Jurisdiction = v
}

// GetLegalStatusOfSupply returns the LegalStatusOfSupply field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetLegalStatusOfSupply() CodeableConcept {
	if o == nil || IsNil(o.LegalStatusOfSupply) {
		var ret CodeableConcept
		return ret
	}
	return *o.LegalStatusOfSupply
}

// GetLegalStatusOfSupplyOk returns a tuple with the LegalStatusOfSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetLegalStatusOfSupplyOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.LegalStatusOfSupply) {
		return nil, false
	}
	return o.LegalStatusOfSupply, true
}

// HasLegalStatusOfSupply returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasLegalStatusOfSupply() bool {
	if o != nil && !IsNil(o.LegalStatusOfSupply) {
		return true
	}

	return false
}

// SetLegalStatusOfSupply gets a reference to the given CodeableConcept and assigns it to the LegalStatusOfSupply field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetLegalStatusOfSupply(v CodeableConcept) {
	o.LegalStatusOfSupply = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetValidityPeriod() Period {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret Period
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) GetValidityPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given Period and assigns it to the ValidityPeriod field.
func (o *MedicinalProductAuthorizationJurisdictionalAuthorization) SetValidityPeriod(v Period) {
	o.ValidityPeriod = &v
}

func (o MedicinalProductAuthorizationJurisdictionalAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductAuthorizationJurisdictionalAuthorization) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	if !IsNil(o.LegalStatusOfSupply) {
		toSerialize["legalStatusOfSupply"] = o.LegalStatusOfSupply
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	return toSerialize, nil
}

type NullableMedicinalProductAuthorizationJurisdictionalAuthorization struct {
	value *MedicinalProductAuthorizationJurisdictionalAuthorization
	isSet bool
}

func (v NullableMedicinalProductAuthorizationJurisdictionalAuthorization) Get() *MedicinalProductAuthorizationJurisdictionalAuthorization {
	return v.value
}

func (v *NullableMedicinalProductAuthorizationJurisdictionalAuthorization) Set(val *MedicinalProductAuthorizationJurisdictionalAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductAuthorizationJurisdictionalAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductAuthorizationJurisdictionalAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductAuthorizationJurisdictionalAuthorization(val *MedicinalProductAuthorizationJurisdictionalAuthorization) *NullableMedicinalProductAuthorizationJurisdictionalAuthorization {
	return &NullableMedicinalProductAuthorizationJurisdictionalAuthorization{value: val, isSet: true}
}

func (v NullableMedicinalProductAuthorizationJurisdictionalAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductAuthorizationJurisdictionalAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


