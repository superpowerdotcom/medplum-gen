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

// checks if the MedicinalProductInteraction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MedicinalProductInteraction{}

// MedicinalProductInteraction The interactions of the medicinal product with other medicinal products, or other forms of interactions.
type MedicinalProductInteraction struct {
	// This is a MedicinalProductInteraction resource
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
	// The medication for which this is a described interaction.
	Subject []Reference `json:"subject,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// The specific medication, food or laboratory test that interacts.
	Interactant []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`
	// The type of the interaction e.g. drug-drug interaction, drug-food interaction, drug-lab test interaction.
	Type *CodeableConcept `json:"type,omitempty"`
	// The effect of the interaction, for example \"reduced gastric absorption of primary medication\".
	Effect *CodeableConcept `json:"effect,omitempty"`
	// The incidence of the interaction, e.g. theoretical, observed.
	Incidence *CodeableConcept `json:"incidence,omitempty"`
	// Actions for managing the interaction.
	Management *CodeableConcept `json:"management,omitempty"`
}

type _MedicinalProductInteraction MedicinalProductInteraction

// NewMedicinalProductInteraction instantiates a new MedicinalProductInteraction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedicinalProductInteraction(resourceType string) *MedicinalProductInteraction {
	this := MedicinalProductInteraction{}
	this.ResourceType = resourceType
	return &this
}

// NewMedicinalProductInteractionWithDefaults instantiates a new MedicinalProductInteraction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMedicinalProductInteractionWithDefaults() *MedicinalProductInteraction {
	this := MedicinalProductInteraction{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *MedicinalProductInteraction) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *MedicinalProductInteraction) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MedicinalProductInteraction) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *MedicinalProductInteraction) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *MedicinalProductInteraction) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MedicinalProductInteraction) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *MedicinalProductInteraction) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *MedicinalProductInteraction) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MedicinalProductInteraction) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MedicinalProductInteraction) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetSubject() []Reference {
	if o == nil || IsNil(o.Subject) {
		var ret []Reference
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetSubjectOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []Reference and assigns it to the Subject field.
func (o *MedicinalProductInteraction) SetSubject(v []Reference) {
	o.Subject = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MedicinalProductInteraction) SetDescription(v string) {
	o.Description = &v
}

// GetInteractant returns the Interactant field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetInteractant() []MedicinalProductInteractionInteractant {
	if o == nil || IsNil(o.Interactant) {
		var ret []MedicinalProductInteractionInteractant
		return ret
	}
	return o.Interactant
}

// GetInteractantOk returns a tuple with the Interactant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetInteractantOk() ([]MedicinalProductInteractionInteractant, bool) {
	if o == nil || IsNil(o.Interactant) {
		return nil, false
	}
	return o.Interactant, true
}

// HasInteractant returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasInteractant() bool {
	if o != nil && !IsNil(o.Interactant) {
		return true
	}

	return false
}

// SetInteractant gets a reference to the given []MedicinalProductInteractionInteractant and assigns it to the Interactant field.
func (o *MedicinalProductInteraction) SetInteractant(v []MedicinalProductInteractionInteractant) {
	o.Interactant = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *MedicinalProductInteraction) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetEffect() CodeableConcept {
	if o == nil || IsNil(o.Effect) {
		var ret CodeableConcept
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetEffectOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Effect) {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasEffect() bool {
	if o != nil && !IsNil(o.Effect) {
		return true
	}

	return false
}

// SetEffect gets a reference to the given CodeableConcept and assigns it to the Effect field.
func (o *MedicinalProductInteraction) SetEffect(v CodeableConcept) {
	o.Effect = &v
}

// GetIncidence returns the Incidence field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetIncidence() CodeableConcept {
	if o == nil || IsNil(o.Incidence) {
		var ret CodeableConcept
		return ret
	}
	return *o.Incidence
}

// GetIncidenceOk returns a tuple with the Incidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetIncidenceOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Incidence) {
		return nil, false
	}
	return o.Incidence, true
}

// HasIncidence returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasIncidence() bool {
	if o != nil && !IsNil(o.Incidence) {
		return true
	}

	return false
}

// SetIncidence gets a reference to the given CodeableConcept and assigns it to the Incidence field.
func (o *MedicinalProductInteraction) SetIncidence(v CodeableConcept) {
	o.Incidence = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *MedicinalProductInteraction) GetManagement() CodeableConcept {
	if o == nil || IsNil(o.Management) {
		var ret CodeableConcept
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MedicinalProductInteraction) GetManagementOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Management) {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *MedicinalProductInteraction) HasManagement() bool {
	if o != nil && !IsNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given CodeableConcept and assigns it to the Management field.
func (o *MedicinalProductInteraction) SetManagement(v CodeableConcept) {
	o.Management = &v
}

func (o MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MedicinalProductInteraction) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Interactant) {
		toSerialize["interactant"] = o.Interactant
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Effect) {
		toSerialize["effect"] = o.Effect
	}
	if !IsNil(o.Incidence) {
		toSerialize["incidence"] = o.Incidence
	}
	if !IsNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	return toSerialize, nil
}

func (o *MedicinalProductInteraction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
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

	varMedicinalProductInteraction := _MedicinalProductInteraction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedicinalProductInteraction)

	if err != nil {
		return err
	}

	*o = MedicinalProductInteraction(varMedicinalProductInteraction)

	return err
}

type NullableMedicinalProductInteraction struct {
	value *MedicinalProductInteraction
	isSet bool
}

func (v NullableMedicinalProductInteraction) Get() *MedicinalProductInteraction {
	return v.value
}

func (v *NullableMedicinalProductInteraction) Set(val *MedicinalProductInteraction) {
	v.value = val
	v.isSet = true
}

func (v NullableMedicinalProductInteraction) IsSet() bool {
	return v.isSet
}

func (v *NullableMedicinalProductInteraction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedicinalProductInteraction(val *MedicinalProductInteraction) *NullableMedicinalProductInteraction {
	return &NullableMedicinalProductInteraction{value: val, isSet: true}
}

func (v NullableMedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedicinalProductInteraction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


