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

// checks if the Substance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Substance{}

// Substance A homogeneous material with a definite composition.
type Substance struct {
	// This is a Substance resource
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
	// Unique identifier for the substance.
	Identifier []Identifier `json:"identifier,omitempty"`
	// A code to indicate if the substance is actively used.
	Status *string `json:"status,omitempty"`
	// A code that classifies the general type of substance.  This is used  for searching, sorting and display purposes.
	Category []CodeableConcept `json:"category,omitempty"`
	// A code (or set of codes) that identify this substance.
	Code CodeableConcept `json:"code"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.
	Instance []SubstanceInstance `json:"instance,omitempty"`
	// A substance can be composed of other substances.
	Ingredient []SubstanceIngredient `json:"ingredient,omitempty"`
}

type _Substance Substance

// NewSubstance instantiates a new Substance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstance(resourceType string, code CodeableConcept) *Substance {
	this := Substance{}
	this.ResourceType = resourceType
	this.Code = code
	return &this
}

// NewSubstanceWithDefaults instantiates a new Substance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceWithDefaults() *Substance {
	this := Substance{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Substance) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Substance) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Substance) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Substance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Substance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Substance) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Substance) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Substance) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Substance) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Substance) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Substance) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Substance) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Substance) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Substance) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Substance) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Substance) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Substance) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Substance) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Substance) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Substance) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Substance) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Substance) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Substance) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Substance) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Substance) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Substance) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Substance) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Substance) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Substance) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Substance) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Substance) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Substance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Substance) SetStatus(v string) {
	o.Status = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Substance) GetCategory() []CodeableConcept {
	if o == nil || IsNil(o.Category) {
		var ret []CodeableConcept
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetCategoryOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Substance) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []CodeableConcept and assigns it to the Category field.
func (o *Substance) SetCategory(v []CodeableConcept) {
	o.Category = v
}

// GetCode returns the Code field value
func (o *Substance) GetCode() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Substance) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Substance) SetCode(v CodeableConcept) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Substance) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Substance) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Substance) SetDescription(v string) {
	o.Description = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *Substance) GetInstance() []SubstanceInstance {
	if o == nil || IsNil(o.Instance) {
		var ret []SubstanceInstance
		return ret
	}
	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetInstanceOk() ([]SubstanceInstance, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *Substance) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given []SubstanceInstance and assigns it to the Instance field.
func (o *Substance) SetInstance(v []SubstanceInstance) {
	o.Instance = v
}

// GetIngredient returns the Ingredient field value if set, zero value otherwise.
func (o *Substance) GetIngredient() []SubstanceIngredient {
	if o == nil || IsNil(o.Ingredient) {
		var ret []SubstanceIngredient
		return ret
	}
	return o.Ingredient
}

// GetIngredientOk returns a tuple with the Ingredient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Substance) GetIngredientOk() ([]SubstanceIngredient, bool) {
	if o == nil || IsNil(o.Ingredient) {
		return nil, false
	}
	return o.Ingredient, true
}

// HasIngredient returns a boolean if a field has been set.
func (o *Substance) HasIngredient() bool {
	if o != nil && !IsNil(o.Ingredient) {
		return true
	}

	return false
}

// SetIngredient gets a reference to the given []SubstanceIngredient and assigns it to the Ingredient field.
func (o *Substance) SetIngredient(v []SubstanceIngredient) {
	o.Ingredient = v
}

func (o Substance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Substance) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Ingredient) {
		toSerialize["ingredient"] = o.Ingredient
	}
	return toSerialize, nil
}

func (o *Substance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"code",
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

	varSubstance := _Substance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubstance)

	if err != nil {
		return err
	}

	*o = Substance(varSubstance)

	return err
}

type NullableSubstance struct {
	value *Substance
	isSet bool
}

func (v NullableSubstance) Get() *Substance {
	return v.value
}

func (v *NullableSubstance) Set(val *Substance) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstance(val *Substance) *NullableSubstance {
	return &NullableSubstance{value: val, isSet: true}
}

func (v NullableSubstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

