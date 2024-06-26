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

// checks if the Provenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Provenance{}

// Provenance Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance struct {
	// This is a Provenance resource
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
	// The Reference(s) that were generated or updated by  the activity described in this resource. A provenance can point to more than one target if multiple resources were created/updated by the same activity.
	Target []Reference `json:"target"`
	// The period during which the activity occurred.
	OccurredPeriod *Period `json:"occurredPeriod,omitempty"`
	// The period during which the activity occurred.
	OccurredDateTime *string `json:"occurredDateTime,omitempty"`
	// An instant in time - known at least to the second
	Recorded *string `json:"recorded,omitempty"`
	// Policy or plan the activity was defined by. Typically, a single activity may have multiple applicable policy documents, such as patient consent, guarantor funding, etc.
	Policy []string `json:"policy,omitempty"`
	// Where the activity occurred, if relevant.
	Location *Reference `json:"location,omitempty"`
	// The reason that the activity was taking place.
	Reason []CodeableConcept `json:"reason,omitempty"`
	// An activity is something that occurs over a period of time and acts upon or with entities; it may include consuming, processing, transforming, modifying, relocating, using, or generating entities.
	Activity *CodeableConcept `json:"activity,omitempty"`
	// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
	Agent []ProvenanceAgent `json:"agent"`
	// An entity used in this activity.
	Entity []ProvenanceEntity `json:"entity,omitempty"`
	// A digital signature on the target Reference(s). The signer should match a Provenance.agent. The purpose of the signature is indicated.
	Signature []Signature `json:"signature,omitempty"`
}

type _Provenance Provenance

// NewProvenance instantiates a new Provenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvenance(resourceType string, target []Reference, agent []ProvenanceAgent) *Provenance {
	this := Provenance{}
	this.ResourceType = resourceType
	this.Target = target
	this.Agent = agent
	return &this
}

// NewProvenanceWithDefaults instantiates a new Provenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvenanceWithDefaults() *Provenance {
	this := Provenance{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Provenance) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Provenance) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Provenance) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Provenance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Provenance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Provenance) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Provenance) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Provenance) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Provenance) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Provenance) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Provenance) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Provenance) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Provenance) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Provenance) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Provenance) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Provenance) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Provenance) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Provenance) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Provenance) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Provenance) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Provenance) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Provenance) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Provenance) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Provenance) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Provenance) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Provenance) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Provenance) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetTarget returns the Target field value
func (o *Provenance) GetTarget() []Reference {
	if o == nil {
		var ret []Reference
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Provenance) GetTargetOk() ([]Reference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target, true
}

// SetTarget sets field value
func (o *Provenance) SetTarget(v []Reference) {
	o.Target = v
}

// GetOccurredPeriod returns the OccurredPeriod field value if set, zero value otherwise.
func (o *Provenance) GetOccurredPeriod() Period {
	if o == nil || IsNil(o.OccurredPeriod) {
		var ret Period
		return ret
	}
	return *o.OccurredPeriod
}

// GetOccurredPeriodOk returns a tuple with the OccurredPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetOccurredPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.OccurredPeriod) {
		return nil, false
	}
	return o.OccurredPeriod, true
}

// HasOccurredPeriod returns a boolean if a field has been set.
func (o *Provenance) HasOccurredPeriod() bool {
	if o != nil && !IsNil(o.OccurredPeriod) {
		return true
	}

	return false
}

// SetOccurredPeriod gets a reference to the given Period and assigns it to the OccurredPeriod field.
func (o *Provenance) SetOccurredPeriod(v Period) {
	o.OccurredPeriod = &v
}

// GetOccurredDateTime returns the OccurredDateTime field value if set, zero value otherwise.
func (o *Provenance) GetOccurredDateTime() string {
	if o == nil || IsNil(o.OccurredDateTime) {
		var ret string
		return ret
	}
	return *o.OccurredDateTime
}

// GetOccurredDateTimeOk returns a tuple with the OccurredDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetOccurredDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OccurredDateTime) {
		return nil, false
	}
	return o.OccurredDateTime, true
}

// HasOccurredDateTime returns a boolean if a field has been set.
func (o *Provenance) HasOccurredDateTime() bool {
	if o != nil && !IsNil(o.OccurredDateTime) {
		return true
	}

	return false
}

// SetOccurredDateTime gets a reference to the given string and assigns it to the OccurredDateTime field.
func (o *Provenance) SetOccurredDateTime(v string) {
	o.OccurredDateTime = &v
}

// GetRecorded returns the Recorded field value if set, zero value otherwise.
func (o *Provenance) GetRecorded() string {
	if o == nil || IsNil(o.Recorded) {
		var ret string
		return ret
	}
	return *o.Recorded
}

// GetRecordedOk returns a tuple with the Recorded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetRecordedOk() (*string, bool) {
	if o == nil || IsNil(o.Recorded) {
		return nil, false
	}
	return o.Recorded, true
}

// HasRecorded returns a boolean if a field has been set.
func (o *Provenance) HasRecorded() bool {
	if o != nil && !IsNil(o.Recorded) {
		return true
	}

	return false
}

// SetRecorded gets a reference to the given string and assigns it to the Recorded field.
func (o *Provenance) SetRecorded(v string) {
	o.Recorded = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *Provenance) GetPolicy() []string {
	if o == nil || IsNil(o.Policy) {
		var ret []string
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *Provenance) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given []string and assigns it to the Policy field.
func (o *Provenance) SetPolicy(v []string) {
	o.Policy = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Provenance) GetLocation() Reference {
	if o == nil || IsNil(o.Location) {
		var ret Reference
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetLocationOk() (*Reference, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Provenance) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Reference and assigns it to the Location field.
func (o *Provenance) SetLocation(v Reference) {
	o.Location = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Provenance) GetReason() []CodeableConcept {
	if o == nil || IsNil(o.Reason) {
		var ret []CodeableConcept
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetReasonOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Provenance) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given []CodeableConcept and assigns it to the Reason field.
func (o *Provenance) SetReason(v []CodeableConcept) {
	o.Reason = v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *Provenance) GetActivity() CodeableConcept {
	if o == nil || IsNil(o.Activity) {
		var ret CodeableConcept
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetActivityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *Provenance) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given CodeableConcept and assigns it to the Activity field.
func (o *Provenance) SetActivity(v CodeableConcept) {
	o.Activity = &v
}

// GetAgent returns the Agent field value
func (o *Provenance) GetAgent() []ProvenanceAgent {
	if o == nil {
		var ret []ProvenanceAgent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *Provenance) GetAgentOk() ([]ProvenanceAgent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agent, true
}

// SetAgent sets field value
func (o *Provenance) SetAgent(v []ProvenanceAgent) {
	o.Agent = v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *Provenance) GetEntity() []ProvenanceEntity {
	if o == nil || IsNil(o.Entity) {
		var ret []ProvenanceEntity
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetEntityOk() ([]ProvenanceEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *Provenance) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given []ProvenanceEntity and assigns it to the Entity field.
func (o *Provenance) SetEntity(v []ProvenanceEntity) {
	o.Entity = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *Provenance) GetSignature() []Signature {
	if o == nil || IsNil(o.Signature) {
		var ret []Signature
		return ret
	}
	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provenance) GetSignatureOk() ([]Signature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *Provenance) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given []Signature and assigns it to the Signature field.
func (o *Provenance) SetSignature(v []Signature) {
	o.Signature = v
}

func (o Provenance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Provenance) ToMap() (map[string]interface{}, error) {
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
	toSerialize["target"] = o.Target
	if !IsNil(o.OccurredPeriod) {
		toSerialize["occurredPeriod"] = o.OccurredPeriod
	}
	if !IsNil(o.OccurredDateTime) {
		toSerialize["occurredDateTime"] = o.OccurredDateTime
	}
	if !IsNil(o.Recorded) {
		toSerialize["recorded"] = o.Recorded
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	toSerialize["agent"] = o.Agent
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	return toSerialize, nil
}

func (o *Provenance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"target",
		"agent",
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

	varProvenance := _Provenance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvenance)

	if err != nil {
		return err
	}

	*o = Provenance(varProvenance)

	return err
}

type NullableProvenance struct {
	value *Provenance
	isSet bool
}

func (v NullableProvenance) Get() *Provenance {
	return v.value
}

func (v *NullableProvenance) Set(val *Provenance) {
	v.value = val
	v.isSet = true
}

func (v NullableProvenance) IsSet() bool {
	return v.isSet
}

func (v *NullableProvenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvenance(val *Provenance) *NullableProvenance {
	return &NullableProvenance{value: val, isSet: true}
}

func (v NullableProvenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


