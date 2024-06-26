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

// checks if the ContractOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractOffer{}

// ContractOffer Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type ContractOffer struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique identifier for this particular Contract Provision.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Offer Recipient.
	Party []ContractParty `json:"party,omitempty"`
	// The owner of an asset has the residual control rights over the asset: the right to decide all usages of the asset in any way not inconsistent with a prior contract, custom, or law (Hart, 1995, p. 30).
	Topic *Reference `json:"topic,omitempty"`
	// Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions, e.g. life time maximum benefit.
	Type *CodeableConcept `json:"type,omitempty"`
	// Type of choice made by accepting party with respect to an offer made by an offeror/ grantee.
	Decision *CodeableConcept `json:"decision,omitempty"`
	// How the decision about a Contract was conveyed.
	DecisionMode []CodeableConcept `json:"decisionMode,omitempty"`
	// Response to offer text.
	Answer []ContractAnswer `json:"answer,omitempty"`
	// A sequence of Unicode characters
	Text *string `json:"text,omitempty"`
	// The id of the clause or question text of the offer in the referenced questionnaire/response.
	LinkId []string `json:"linkId,omitempty"`
	// Security labels that protects the offer.
	SecurityLabelNumber []float32 `json:"securityLabelNumber,omitempty"`
}

// NewContractOffer instantiates a new ContractOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractOffer() *ContractOffer {
	this := ContractOffer{}
	return &this
}

// NewContractOfferWithDefaults instantiates a new ContractOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractOfferWithDefaults() *ContractOffer {
	this := ContractOffer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractOffer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractOffer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractOffer) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ContractOffer) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ContractOffer) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ContractOffer) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ContractOffer) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ContractOffer) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ContractOffer) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ContractOffer) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ContractOffer) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *ContractOffer) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetParty returns the Party field value if set, zero value otherwise.
func (o *ContractOffer) GetParty() []ContractParty {
	if o == nil || IsNil(o.Party) {
		var ret []ContractParty
		return ret
	}
	return o.Party
}

// GetPartyOk returns a tuple with the Party field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetPartyOk() ([]ContractParty, bool) {
	if o == nil || IsNil(o.Party) {
		return nil, false
	}
	return o.Party, true
}

// HasParty returns a boolean if a field has been set.
func (o *ContractOffer) HasParty() bool {
	if o != nil && !IsNil(o.Party) {
		return true
	}

	return false
}

// SetParty gets a reference to the given []ContractParty and assigns it to the Party field.
func (o *ContractOffer) SetParty(v []ContractParty) {
	o.Party = v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ContractOffer) GetTopic() Reference {
	if o == nil || IsNil(o.Topic) {
		var ret Reference
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetTopicOk() (*Reference, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ContractOffer) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given Reference and assigns it to the Topic field.
func (o *ContractOffer) SetTopic(v Reference) {
	o.Topic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContractOffer) GetType() CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret CodeableConcept
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContractOffer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodeableConcept and assigns it to the Type field.
func (o *ContractOffer) SetType(v CodeableConcept) {
	o.Type = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *ContractOffer) GetDecision() CodeableConcept {
	if o == nil || IsNil(o.Decision) {
		var ret CodeableConcept
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetDecisionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Decision) {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *ContractOffer) HasDecision() bool {
	if o != nil && !IsNil(o.Decision) {
		return true
	}

	return false
}

// SetDecision gets a reference to the given CodeableConcept and assigns it to the Decision field.
func (o *ContractOffer) SetDecision(v CodeableConcept) {
	o.Decision = &v
}

// GetDecisionMode returns the DecisionMode field value if set, zero value otherwise.
func (o *ContractOffer) GetDecisionMode() []CodeableConcept {
	if o == nil || IsNil(o.DecisionMode) {
		var ret []CodeableConcept
		return ret
	}
	return o.DecisionMode
}

// GetDecisionModeOk returns a tuple with the DecisionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetDecisionModeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.DecisionMode) {
		return nil, false
	}
	return o.DecisionMode, true
}

// HasDecisionMode returns a boolean if a field has been set.
func (o *ContractOffer) HasDecisionMode() bool {
	if o != nil && !IsNil(o.DecisionMode) {
		return true
	}

	return false
}

// SetDecisionMode gets a reference to the given []CodeableConcept and assigns it to the DecisionMode field.
func (o *ContractOffer) SetDecisionMode(v []CodeableConcept) {
	o.DecisionMode = v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *ContractOffer) GetAnswer() []ContractAnswer {
	if o == nil || IsNil(o.Answer) {
		var ret []ContractAnswer
		return ret
	}
	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetAnswerOk() ([]ContractAnswer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *ContractOffer) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given []ContractAnswer and assigns it to the Answer field.
func (o *ContractOffer) SetAnswer(v []ContractAnswer) {
	o.Answer = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ContractOffer) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ContractOffer) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ContractOffer) SetText(v string) {
	o.Text = &v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *ContractOffer) GetLinkId() []string {
	if o == nil || IsNil(o.LinkId) {
		var ret []string
		return ret
	}
	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetLinkIdOk() ([]string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *ContractOffer) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given []string and assigns it to the LinkId field.
func (o *ContractOffer) SetLinkId(v []string) {
	o.LinkId = v
}

// GetSecurityLabelNumber returns the SecurityLabelNumber field value if set, zero value otherwise.
func (o *ContractOffer) GetSecurityLabelNumber() []float32 {
	if o == nil || IsNil(o.SecurityLabelNumber) {
		var ret []float32
		return ret
	}
	return o.SecurityLabelNumber
}

// GetSecurityLabelNumberOk returns a tuple with the SecurityLabelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOffer) GetSecurityLabelNumberOk() ([]float32, bool) {
	if o == nil || IsNil(o.SecurityLabelNumber) {
		return nil, false
	}
	return o.SecurityLabelNumber, true
}

// HasSecurityLabelNumber returns a boolean if a field has been set.
func (o *ContractOffer) HasSecurityLabelNumber() bool {
	if o != nil && !IsNil(o.SecurityLabelNumber) {
		return true
	}

	return false
}

// SetSecurityLabelNumber gets a reference to the given []float32 and assigns it to the SecurityLabelNumber field.
func (o *ContractOffer) SetSecurityLabelNumber(v []float32) {
	o.SecurityLabelNumber = v
}

func (o ContractOffer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractOffer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Party) {
		toSerialize["party"] = o.Party
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Decision) {
		toSerialize["decision"] = o.Decision
	}
	if !IsNil(o.DecisionMode) {
		toSerialize["decisionMode"] = o.DecisionMode
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.SecurityLabelNumber) {
		toSerialize["securityLabelNumber"] = o.SecurityLabelNumber
	}
	return toSerialize, nil
}

type NullableContractOffer struct {
	value *ContractOffer
	isSet bool
}

func (v NullableContractOffer) Get() *ContractOffer {
	return v.value
}

func (v *NullableContractOffer) Set(val *ContractOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableContractOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableContractOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractOffer(val *ContractOffer) *NullableContractOffer {
	return &NullableContractOffer{value: val, isSet: true}
}

func (v NullableContractOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


