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

// checks if the SubstanceSpecificationName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstanceSpecificationName{}

// SubstanceSpecificationName The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecificationName struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// Name type.
	Type *CodeableConcept `json:"type,omitempty"`
	// The status of the name.
	Status *CodeableConcept `json:"status,omitempty"`
	// Value of \"true\" or \"false\"
	Preferred *bool `json:"preferred,omitempty"`
	// Language of the name.
	Language []CodeableConcept `json:"language,omitempty"`
	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive.
	Domain []CodeableConcept `json:"domain,omitempty"`
	// The jurisdiction where this name applies.
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// A synonym of this name.
	Synonym []SubstanceSpecificationName `json:"synonym,omitempty"`
	// A translation for this name.
	Translation []SubstanceSpecificationName `json:"translation,omitempty"`
	// Details of the official nature of this name.
	Official []SubstanceSpecificationOfficial `json:"official,omitempty"`
	// Supporting literature.
	Source []Reference `json:"source,omitempty"`
}

// NewSubstanceSpecificationName instantiates a new SubstanceSpecificationName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstanceSpecificationName() *SubstanceSpecificationName {
	this := SubstanceSpecificationName{}
	return &this
}

// NewSubstanceSpecificationNameWithDefaults instantiates a new SubstanceSpecificationName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceSpecificationNameWithDefaults() *SubstanceSpecificationName {
	this := SubstanceSpecificationName{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstanceSpecificationName) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstanceSpecificationName) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstanceSpecificationName) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubstanceSpecificationName) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *SubstanceSpecificationName) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetStatus() CodeableConcept {
	if o == nil || IsNil(o.Status) {
		var ret CodeableConcept
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetStatusOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CodeableConcept and assigns it to the Status field.
func (o *SubstanceSpecificationName) SetStatus(v CodeableConcept) {
	o.Status = &v
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetPreferred() bool {
	if o == nil || IsNil(o.Preferred) {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasPreferred() bool {
	if o != nil && !IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *SubstanceSpecificationName) SetPreferred(v bool) {
	o.Preferred = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetLanguage() []CodeableConcept {
	if o == nil || IsNil(o.Language) {
		var ret []CodeableConcept
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetLanguageOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []CodeableConcept and assigns it to the Language field.
func (o *SubstanceSpecificationName) SetLanguage(v []CodeableConcept) {
	o.Language = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetDomain() []CodeableConcept {
	if o == nil || IsNil(o.Domain) {
		var ret []CodeableConcept
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetDomainOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given []CodeableConcept and assigns it to the Domain field.
func (o *SubstanceSpecificationName) SetDomain(v []CodeableConcept) {
	o.Domain = v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetJurisdiction() []CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret []CodeableConcept
		return ret
	}
	return o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetJurisdictionOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given []CodeableConcept and assigns it to the Jurisdiction field.
func (o *SubstanceSpecificationName) SetJurisdiction(v []CodeableConcept) {
	o.Jurisdiction = v
}

// GetSynonym returns the Synonym field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetSynonym() []SubstanceSpecificationName {
	if o == nil || IsNil(o.Synonym) {
		var ret []SubstanceSpecificationName
		return ret
	}
	return o.Synonym
}

// GetSynonymOk returns a tuple with the Synonym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetSynonymOk() ([]SubstanceSpecificationName, bool) {
	if o == nil || IsNil(o.Synonym) {
		return nil, false
	}
	return o.Synonym, true
}

// HasSynonym returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasSynonym() bool {
	if o != nil && !IsNil(o.Synonym) {
		return true
	}

	return false
}

// SetSynonym gets a reference to the given []SubstanceSpecificationName and assigns it to the Synonym field.
func (o *SubstanceSpecificationName) SetSynonym(v []SubstanceSpecificationName) {
	o.Synonym = v
}

// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetTranslation() []SubstanceSpecificationName {
	if o == nil || IsNil(o.Translation) {
		var ret []SubstanceSpecificationName
		return ret
	}
	return o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetTranslationOk() ([]SubstanceSpecificationName, bool) {
	if o == nil || IsNil(o.Translation) {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasTranslation() bool {
	if o != nil && !IsNil(o.Translation) {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given []SubstanceSpecificationName and assigns it to the Translation field.
func (o *SubstanceSpecificationName) SetTranslation(v []SubstanceSpecificationName) {
	o.Translation = v
}

// GetOfficial returns the Official field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetOfficial() []SubstanceSpecificationOfficial {
	if o == nil || IsNil(o.Official) {
		var ret []SubstanceSpecificationOfficial
		return ret
	}
	return o.Official
}

// GetOfficialOk returns a tuple with the Official field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetOfficialOk() ([]SubstanceSpecificationOfficial, bool) {
	if o == nil || IsNil(o.Official) {
		return nil, false
	}
	return o.Official, true
}

// HasOfficial returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasOfficial() bool {
	if o != nil && !IsNil(o.Official) {
		return true
	}

	return false
}

// SetOfficial gets a reference to the given []SubstanceSpecificationOfficial and assigns it to the Official field.
func (o *SubstanceSpecificationName) SetOfficial(v []SubstanceSpecificationOfficial) {
	o.Official = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SubstanceSpecificationName) GetSource() []Reference {
	if o == nil || IsNil(o.Source) {
		var ret []Reference
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationName) GetSourceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SubstanceSpecificationName) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given []Reference and assigns it to the Source field.
func (o *SubstanceSpecificationName) SetSource(v []Reference) {
	o.Source = v
}

func (o SubstanceSpecificationName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstanceSpecificationName) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	if !IsNil(o.Synonym) {
		toSerialize["synonym"] = o.Synonym
	}
	if !IsNil(o.Translation) {
		toSerialize["translation"] = o.Translation
	}
	if !IsNil(o.Official) {
		toSerialize["official"] = o.Official
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableSubstanceSpecificationName struct {
	value *SubstanceSpecificationName
	isSet bool
}

func (v NullableSubstanceSpecificationName) Get() *SubstanceSpecificationName {
	return v.value
}

func (v *NullableSubstanceSpecificationName) Set(val *SubstanceSpecificationName) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstanceSpecificationName) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstanceSpecificationName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstanceSpecificationName(val *SubstanceSpecificationName) *NullableSubstanceSpecificationName {
	return &NullableSubstanceSpecificationName{value: val, isSet: true}
}

func (v NullableSubstanceSpecificationName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstanceSpecificationName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

