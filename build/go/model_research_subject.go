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

// checks if the ResearchSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResearchSubject{}

// ResearchSubject A physical entity which is the primary unit of operational and/or administrative interest in a study.
type ResearchSubject struct {
	// This is a ResearchSubject resource
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
	// Identifiers assigned to this research subject for a study.
	Identifier []Identifier `json:"identifier,omitempty"`
	// The current state of the subject.
	Status *string `json:"status,omitempty"`
	// The dates the subject began and ended their participation in the study.
	Period *Period `json:"period,omitempty"`
	// Reference to the study the subject is participating in.
	Study Reference `json:"study"`
	// The record of the person or animal who is involved in the study.
	Individual Reference `json:"individual"`
	// A sequence of Unicode characters
	AssignedArm *string `json:"assignedArm,omitempty"`
	// A sequence of Unicode characters
	ActualArm *string `json:"actualArm,omitempty"`
	// A record of the patient's informed agreement to participate in the study.
	Consent *Reference `json:"consent,omitempty"`
}

type _ResearchSubject ResearchSubject

// NewResearchSubject instantiates a new ResearchSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResearchSubject(resourceType string, study Reference, individual Reference) *ResearchSubject {
	this := ResearchSubject{}
	this.ResourceType = resourceType
	this.Study = study
	this.Individual = individual
	return &this
}

// NewResearchSubjectWithDefaults instantiates a new ResearchSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResearchSubjectWithDefaults() *ResearchSubject {
	this := ResearchSubject{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *ResearchSubject) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ResearchSubject) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResearchSubject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResearchSubject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResearchSubject) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ResearchSubject) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ResearchSubject) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *ResearchSubject) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *ResearchSubject) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *ResearchSubject) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *ResearchSubject) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ResearchSubject) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ResearchSubject) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ResearchSubject) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ResearchSubject) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ResearchSubject) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *ResearchSubject) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *ResearchSubject) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *ResearchSubject) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *ResearchSubject) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ResearchSubject) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ResearchSubject) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ResearchSubject) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ResearchSubject) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ResearchSubject) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ResearchSubject) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ResearchSubject) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ResearchSubject) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *ResearchSubject) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResearchSubject) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResearchSubject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResearchSubject) SetStatus(v string) {
	o.Status = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ResearchSubject) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ResearchSubject) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *ResearchSubject) SetPeriod(v Period) {
	o.Period = &v
}

// GetStudy returns the Study field value
func (o *ResearchSubject) GetStudy() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Study
}

// GetStudyOk returns a tuple with the Study field value
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetStudyOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Study, true
}

// SetStudy sets field value
func (o *ResearchSubject) SetStudy(v Reference) {
	o.Study = v
}

// GetIndividual returns the Individual field value
func (o *ResearchSubject) GetIndividual() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetIndividualOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Individual, true
}

// SetIndividual sets field value
func (o *ResearchSubject) SetIndividual(v Reference) {
	o.Individual = v
}

// GetAssignedArm returns the AssignedArm field value if set, zero value otherwise.
func (o *ResearchSubject) GetAssignedArm() string {
	if o == nil || IsNil(o.AssignedArm) {
		var ret string
		return ret
	}
	return *o.AssignedArm
}

// GetAssignedArmOk returns a tuple with the AssignedArm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetAssignedArmOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedArm) {
		return nil, false
	}
	return o.AssignedArm, true
}

// HasAssignedArm returns a boolean if a field has been set.
func (o *ResearchSubject) HasAssignedArm() bool {
	if o != nil && !IsNil(o.AssignedArm) {
		return true
	}

	return false
}

// SetAssignedArm gets a reference to the given string and assigns it to the AssignedArm field.
func (o *ResearchSubject) SetAssignedArm(v string) {
	o.AssignedArm = &v
}

// GetActualArm returns the ActualArm field value if set, zero value otherwise.
func (o *ResearchSubject) GetActualArm() string {
	if o == nil || IsNil(o.ActualArm) {
		var ret string
		return ret
	}
	return *o.ActualArm
}

// GetActualArmOk returns a tuple with the ActualArm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetActualArmOk() (*string, bool) {
	if o == nil || IsNil(o.ActualArm) {
		return nil, false
	}
	return o.ActualArm, true
}

// HasActualArm returns a boolean if a field has been set.
func (o *ResearchSubject) HasActualArm() bool {
	if o != nil && !IsNil(o.ActualArm) {
		return true
	}

	return false
}

// SetActualArm gets a reference to the given string and assigns it to the ActualArm field.
func (o *ResearchSubject) SetActualArm(v string) {
	o.ActualArm = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *ResearchSubject) GetConsent() Reference {
	if o == nil || IsNil(o.Consent) {
		var ret Reference
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResearchSubject) GetConsentOk() (*Reference, bool) {
	if o == nil || IsNil(o.Consent) {
		return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *ResearchSubject) HasConsent() bool {
	if o != nil && !IsNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given Reference and assigns it to the Consent field.
func (o *ResearchSubject) SetConsent(v Reference) {
	o.Consent = &v
}

func (o ResearchSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResearchSubject) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["study"] = o.Study
	toSerialize["individual"] = o.Individual
	if !IsNil(o.AssignedArm) {
		toSerialize["assignedArm"] = o.AssignedArm
	}
	if !IsNil(o.ActualArm) {
		toSerialize["actualArm"] = o.ActualArm
	}
	if !IsNil(o.Consent) {
		toSerialize["consent"] = o.Consent
	}
	return toSerialize, nil
}

func (o *ResearchSubject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"study",
		"individual",
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

	varResearchSubject := _ResearchSubject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResearchSubject)

	if err != nil {
		return err
	}

	*o = ResearchSubject(varResearchSubject)

	return err
}

type NullableResearchSubject struct {
	value *ResearchSubject
	isSet bool
}

func (v NullableResearchSubject) Get() *ResearchSubject {
	return v.value
}

func (v *NullableResearchSubject) Set(val *ResearchSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableResearchSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableResearchSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResearchSubject(val *ResearchSubject) *NullableResearchSubject {
	return &NullableResearchSubject{value: val, isSet: true}
}

func (v NullableResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResearchSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


