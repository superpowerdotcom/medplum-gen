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

// checks if the MedicinalProductIngredient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductIngredient{}

// MedicinalProductIngredient An ingredient of a manufactured item or pharmaceutical product.
type MedicinalProductIngredient struct {
	// This is a MedicinalProductIngredient resource
	ResourceType string `json:"resourceType"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Id *string `json:"id,omitempty"`
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `json:"meta,omitempty"`
	// String of characters used to identify a name or a resource
	ImplicitRules *string `json:"implicitRules,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Language *string `json:"language,omitempty"`
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it \"clinically safe\" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative `json:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []ResourceList `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The identifier(s) of this Ingredient that are assigned by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier *Identifier `json:"identifier,omitempty"`
	// Ingredient role e.g. Active ingredient, excipient.
	Role CodeableConcept `json:"role"`
	// Value of \"true\" or \"false\"
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`
	// Manufacturer of this Ingredient.
	Manufacturer []Reference `json:"manufacturer,omitempty"`
	// A specified substance that comprises this ingredient.
	SpecifiedSubstance []MedicinalProductIngredientSpecifiedSubstance `json:"specifiedSubstance,omitempty"`
	// The ingredient substance.
	Substance *MedicinalProductIngredientSubstance `json:"substance,omitempty"`
}

type _MedicinalProductIngredient MedicinalProductIngredient

// NewMedicinalProductIngredient instantiates a new MedicinalProductIngredient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductIngredient(resourceType string, role CodeableConcept) *MedicinalProductIngredient {
	this := MedicinalProductIngredient{}
	this.ResourceType = resourceType
	this.Role = role
	return &this
}

// NewMedicinalProductIngredientWithDefaults instantiates a new MedicinalProductIngredient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductIngredientWithDefaults() *MedicinalProductIngredient {
	this := MedicinalProductIngredient{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *MedicinalProductIngredient) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *MedicinalProductIngredient) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductIngredient) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *MedicinalProductIngredient) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *MedicinalProductIngredient) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MedicinalProductIngredient) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *MedicinalProductIngredient) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *MedicinalProductIngredient) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductIngredient) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductIngredient) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *MedicinalProductIngredient) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetRole returns the Role field value
func (o *MedicinalProductIngredient) GetRole() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetRoleOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *MedicinalProductIngredient) SetRole(v CodeableConcept) {
	o.Role = v
}

// GetAllergenicIndicator returns the AllergenicIndicator field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetAllergenicIndicator() bool {
	if o == nil || IsNil(o.AllergenicIndicator) {
		var ret bool
		return ret
	}
	return *o.AllergenicIndicator
}

// GetAllergenicIndicatorOk returns a tuple with the AllergenicIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetAllergenicIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.AllergenicIndicator) {
		return nil, false
	}
	return o.AllergenicIndicator, true
}

// HasAllergenicIndicator returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasAllergenicIndicator() bool {
	if o != nil && !IsNil(o.AllergenicIndicator) {
		return true
	}

	return false
}

// SetAllergenicIndicator gets a reference to the given bool and assigns it to the AllergenicIndicator field.
func (o *MedicinalProductIngredient) SetAllergenicIndicator(v bool) {
	o.AllergenicIndicator = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetManufacturer() []Reference {
	if o == nil || IsNil(o.Manufacturer) {
		var ret []Reference
		return ret
	}
	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetManufacturerOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given []Reference and assigns it to the Manufacturer field.
func (o *MedicinalProductIngredient) SetManufacturer(v []Reference) {
	o.Manufacturer = v
}

// GetSpecifiedSubstance returns the SpecifiedSubstance field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetSpecifiedSubstance() []MedicinalProductIngredientSpecifiedSubstance {
	if o == nil || IsNil(o.SpecifiedSubstance) {
		var ret []MedicinalProductIngredientSpecifiedSubstance
		return ret
	}
	return o.SpecifiedSubstance
}

// GetSpecifiedSubstanceOk returns a tuple with the SpecifiedSubstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetSpecifiedSubstanceOk() ([]MedicinalProductIngredientSpecifiedSubstance, bool) {
	if o == nil || IsNil(o.SpecifiedSubstance) {
		return nil, false
	}
	return o.SpecifiedSubstance, true
}

// HasSpecifiedSubstance returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasSpecifiedSubstance() bool {
	if o != nil && !IsNil(o.SpecifiedSubstance) {
		return true
	}

	return false
}

// SetSpecifiedSubstance gets a reference to the given []MedicinalProductIngredientSpecifiedSubstance and assigns it to the SpecifiedSubstance field.
func (o *MedicinalProductIngredient) SetSpecifiedSubstance(v []MedicinalProductIngredientSpecifiedSubstance) {
	o.SpecifiedSubstance = v
}

// GetSubstance returns the Substance field value if set, zero value otherwise.
func (o *MedicinalProductIngredient) GetSubstance() MedicinalProductIngredientSubstance {
	if o == nil || IsNil(o.Substance) {
		var ret MedicinalProductIngredientSubstance
		return ret
	}
	return *o.Substance
}

// GetSubstanceOk returns a tuple with the Substance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductIngredient) GetSubstanceOk() (*MedicinalProductIngredientSubstance, bool) {
	if o == nil || IsNil(o.Substance) {
		return nil, false
	}
	return o.Substance, true
}

// HasSubstance returns a boolean if a field has been set.
func (o *MedicinalProductIngredient) HasSubstance() bool {
	if o != nil && !IsNil(o.Substance) {
		return true
	}

	return false
}

// SetSubstance gets a reference to the given MedicinalProductIngredientSubstance and assigns it to the Substance field.
func (o *MedicinalProductIngredient) SetSubstance(v MedicinalProductIngredientSubstance) {
	o.Substance = &v
}

func (o MedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductIngredient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ImplicitRules) {
		toSerialize["implicitRules"] = o.ImplicitRules
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Contained) {
		toSerialize["contained"] = o.Contained
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
	toSerialize["role"] = o.Role
	if !IsNil(o.AllergenicIndicator) {
		toSerialize["allergenicIndicator"] = o.AllergenicIndicator
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.SpecifiedSubstance) {
		toSerialize["specifiedSubstance"] = o.SpecifiedSubstance
	}
	if !IsNil(o.Substance) {
		toSerialize["substance"] = o.Substance
	}
	return toSerialize, nil
}

func (o *MedicinalProductIngredient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"role",
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

	varMedicinalProductIngredient := _MedicinalProductIngredient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicinalProductIngredient)

	if err != nil {
		return err
	}

	*o = MedicinalProductIngredient(varMedicinalProductIngredient)

	return err
}

type NullableMedicinalProductIngredient struct {
	value *MedicinalProductIngredient
	isSet bool
}

func (v NullableMedicinalProductIngredient) Get() *MedicinalProductIngredient {
	return v.value
}

func (v *NullableMedicinalProductIngredient) Set(val *MedicinalProductIngredient) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductIngredient) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductIngredient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductIngredient(val *MedicinalProductIngredient) *NullableMedicinalProductIngredient {
	return &NullableMedicinalProductIngredient{value: val, isSet: true}
}

func (v NullableMedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductIngredient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


