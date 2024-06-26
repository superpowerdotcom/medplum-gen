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

// checks if the ContractAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractAsset{}

// ContractAsset Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type ContractAsset struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Differentiates the kind of the asset .
	Scope *CodeableConcept `json:"scope,omitempty"`
	// Target entity type about which the term may be concerned.
	Type []CodeableConcept `json:"type,omitempty"`
	// Associated entities.
	TypeReference []Reference `json:"typeReference,omitempty"`
	// May be a subtype or part of an offered asset.
	Subtype []CodeableConcept `json:"subtype,omitempty"`
	// Specifies the applicability of the term to an asset resource instance, and instances it refers to orinstances that refer to it, and/or are owned by the offeree.
	Relationship *Coding `json:"relationship,omitempty"`
	// Circumstance of the asset.
	Context []ContractContext `json:"context,omitempty"`
	// A sequence of Unicode characters
	Condition *string `json:"condition,omitempty"`
	// Type of Asset availability for use or ownership.
	PeriodType []CodeableConcept `json:"periodType,omitempty"`
	// Asset relevant contractual time period.
	Period []Period `json:"period,omitempty"`
	// Time period of asset use.
	UsePeriod []Period `json:"usePeriod,omitempty"`
	// A sequence of Unicode characters
	Text *string `json:"text,omitempty"`
	// Id [identifier??] of the clause or question text about the asset in the referenced form or QuestionnaireResponse.
	LinkId []string `json:"linkId,omitempty"`
	// Response to assets.
	Answer []ContractAnswer `json:"answer,omitempty"`
	// Security labels that protects the asset.
	SecurityLabelNumber []float32 `json:"securityLabelNumber,omitempty"`
	// Contract Valued Item List.
	ValuedItem []ContractValuedItem `json:"valuedItem,omitempty"`
}

// NewContractAsset instantiates a new ContractAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractAsset() *ContractAsset {
	this := ContractAsset{}
	return &this
}

// NewContractAssetWithDefaults instantiates a new ContractAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractAssetWithDefaults() *ContractAsset {
	this := ContractAsset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractAsset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractAsset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractAsset) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ContractAsset) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ContractAsset) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ContractAsset) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ContractAsset) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ContractAsset) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ContractAsset) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ContractAsset) GetScope() CodeableConcept {
	if o == nil || IsNil(o.Scope) {
		var ret CodeableConcept
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetScopeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ContractAsset) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given CodeableConcept and assigns it to the Scope field.
func (o *ContractAsset) SetScope(v CodeableConcept) {
	o.Scope = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContractAsset) GetType() []CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret []CodeableConcept
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContractAsset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []CodeableConcept and assigns it to the Type field.
func (o *ContractAsset) SetType(v []CodeableConcept) {
	o.Type = v
}

// GetTypeReference returns the TypeReference field value if set, zero value otherwise.
func (o *ContractAsset) GetTypeReference() []Reference {
	if o == nil || IsNil(o.TypeReference) {
		var ret []Reference
		return ret
	}
	return o.TypeReference
}

// GetTypeReferenceOk returns a tuple with the TypeReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetTypeReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.TypeReference) {
		return nil, false
	}
	return o.TypeReference, true
}

// HasTypeReference returns a boolean if a field has been set.
func (o *ContractAsset) HasTypeReference() bool {
	if o != nil && !IsNil(o.TypeReference) {
		return true
	}

	return false
}

// SetTypeReference gets a reference to the given []Reference and assigns it to the TypeReference field.
func (o *ContractAsset) SetTypeReference(v []Reference) {
	o.TypeReference = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *ContractAsset) GetSubtype() []CodeableConcept {
	if o == nil || IsNil(o.Subtype) {
		var ret []CodeableConcept
		return ret
	}
	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetSubtypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *ContractAsset) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given []CodeableConcept and assigns it to the Subtype field.
func (o *ContractAsset) SetSubtype(v []CodeableConcept) {
	o.Subtype = v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *ContractAsset) GetRelationship() Coding {
	if o == nil || IsNil(o.Relationship) {
		var ret Coding
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetRelationshipOk() (*Coding, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *ContractAsset) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given Coding and assigns it to the Relationship field.
func (o *ContractAsset) SetRelationship(v Coding) {
	o.Relationship = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ContractAsset) GetContext() []ContractContext {
	if o == nil || IsNil(o.Context) {
		var ret []ContractContext
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetContextOk() ([]ContractContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ContractAsset) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []ContractContext and assigns it to the Context field.
func (o *ContractAsset) SetContext(v []ContractContext) {
	o.Context = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ContractAsset) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ContractAsset) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *ContractAsset) SetCondition(v string) {
	o.Condition = &v
}

// GetPeriodType returns the PeriodType field value if set, zero value otherwise.
func (o *ContractAsset) GetPeriodType() []CodeableConcept {
	if o == nil || IsNil(o.PeriodType) {
		var ret []CodeableConcept
		return ret
	}
	return o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetPeriodTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.PeriodType) {
		return nil, false
	}
	return o.PeriodType, true
}

// HasPeriodType returns a boolean if a field has been set.
func (o *ContractAsset) HasPeriodType() bool {
	if o != nil && !IsNil(o.PeriodType) {
		return true
	}

	return false
}

// SetPeriodType gets a reference to the given []CodeableConcept and assigns it to the PeriodType field.
func (o *ContractAsset) SetPeriodType(v []CodeableConcept) {
	o.PeriodType = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ContractAsset) GetPeriod() []Period {
	if o == nil || IsNil(o.Period) {
		var ret []Period
		return ret
	}
	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetPeriodOk() ([]Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ContractAsset) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given []Period and assigns it to the Period field.
func (o *ContractAsset) SetPeriod(v []Period) {
	o.Period = v
}

// GetUsePeriod returns the UsePeriod field value if set, zero value otherwise.
func (o *ContractAsset) GetUsePeriod() []Period {
	if o == nil || IsNil(o.UsePeriod) {
		var ret []Period
		return ret
	}
	return o.UsePeriod
}

// GetUsePeriodOk returns a tuple with the UsePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetUsePeriodOk() ([]Period, bool) {
	if o == nil || IsNil(o.UsePeriod) {
		return nil, false
	}
	return o.UsePeriod, true
}

// HasUsePeriod returns a boolean if a field has been set.
func (o *ContractAsset) HasUsePeriod() bool {
	if o != nil && !IsNil(o.UsePeriod) {
		return true
	}

	return false
}

// SetUsePeriod gets a reference to the given []Period and assigns it to the UsePeriod field.
func (o *ContractAsset) SetUsePeriod(v []Period) {
	o.UsePeriod = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ContractAsset) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ContractAsset) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ContractAsset) SetText(v string) {
	o.Text = &v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *ContractAsset) GetLinkId() []string {
	if o == nil || IsNil(o.LinkId) {
		var ret []string
		return ret
	}
	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetLinkIdOk() ([]string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *ContractAsset) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given []string and assigns it to the LinkId field.
func (o *ContractAsset) SetLinkId(v []string) {
	o.LinkId = v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *ContractAsset) GetAnswer() []ContractAnswer {
	if o == nil || IsNil(o.Answer) {
		var ret []ContractAnswer
		return ret
	}
	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetAnswerOk() ([]ContractAnswer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *ContractAsset) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given []ContractAnswer and assigns it to the Answer field.
func (o *ContractAsset) SetAnswer(v []ContractAnswer) {
	o.Answer = v
}

// GetSecurityLabelNumber returns the SecurityLabelNumber field value if set, zero value otherwise.
func (o *ContractAsset) GetSecurityLabelNumber() []float32 {
	if o == nil || IsNil(o.SecurityLabelNumber) {
		var ret []float32
		return ret
	}
	return o.SecurityLabelNumber
}

// GetSecurityLabelNumberOk returns a tuple with the SecurityLabelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetSecurityLabelNumberOk() ([]float32, bool) {
	if o == nil || IsNil(o.SecurityLabelNumber) {
		return nil, false
	}
	return o.SecurityLabelNumber, true
}

// HasSecurityLabelNumber returns a boolean if a field has been set.
func (o *ContractAsset) HasSecurityLabelNumber() bool {
	if o != nil && !IsNil(o.SecurityLabelNumber) {
		return true
	}

	return false
}

// SetSecurityLabelNumber gets a reference to the given []float32 and assigns it to the SecurityLabelNumber field.
func (o *ContractAsset) SetSecurityLabelNumber(v []float32) {
	o.SecurityLabelNumber = v
}

// GetValuedItem returns the ValuedItem field value if set, zero value otherwise.
func (o *ContractAsset) GetValuedItem() []ContractValuedItem {
	if o == nil || IsNil(o.ValuedItem) {
		var ret []ContractValuedItem
		return ret
	}
	return o.ValuedItem
}

// GetValuedItemOk returns a tuple with the ValuedItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAsset) GetValuedItemOk() ([]ContractValuedItem, bool) {
	if o == nil || IsNil(o.ValuedItem) {
		return nil, false
	}
	return o.ValuedItem, true
}

// HasValuedItem returns a boolean if a field has been set.
func (o *ContractAsset) HasValuedItem() bool {
	if o != nil && !IsNil(o.ValuedItem) {
		return true
	}

	return false
}

// SetValuedItem gets a reference to the given []ContractValuedItem and assigns it to the ValuedItem field.
func (o *ContractAsset) SetValuedItem(v []ContractValuedItem) {
	o.ValuedItem = v
}

func (o ContractAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractAsset) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeReference) {
		toSerialize["typeReference"] = o.TypeReference
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.PeriodType) {
		toSerialize["periodType"] = o.PeriodType
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.UsePeriod) {
		toSerialize["usePeriod"] = o.UsePeriod
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.SecurityLabelNumber) {
		toSerialize["securityLabelNumber"] = o.SecurityLabelNumber
	}
	if !IsNil(o.ValuedItem) {
		toSerialize["valuedItem"] = o.ValuedItem
	}
	return toSerialize, nil
}

type NullableContractAsset struct {
	value *ContractAsset
	isSet bool
}

func (v NullableContractAsset) Get() *ContractAsset {
	return v.value
}

func (v *NullableContractAsset) Set(val *ContractAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableContractAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableContractAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractAsset(val *ContractAsset) *NullableContractAsset {
	return &NullableContractAsset{value: val, isSet: true}
}

func (v NullableContractAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


