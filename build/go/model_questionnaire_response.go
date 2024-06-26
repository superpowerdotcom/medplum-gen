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

// checks if the QuestionnaireResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireResponse{}

// QuestionnaireResponse A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse struct {
	// This is a QuestionnaireResponse resource
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
	// A business identifier assigned to a particular completed (or partially completed) questionnaire.
	Identifier *Identifier `json:"identifier,omitempty"`
	// The order, proposal or plan that is fulfilled in whole or in part by this QuestionnaireResponse.  For example, a ServiceRequest seeking an intake assessment or a decision support recommendation to assess for post-partum depression.
	BasedOn []Reference `json:"basedOn,omitempty"`
	// A procedure or observation that this questionnaire was performed as part of the execution of.  For example, the surgery a checklist was executed as part of.
	PartOf []Reference `json:"partOf,omitempty"`
	// A URI that is a reference to a canonical URL on a FHIR resource
	Questionnaire *string `json:"questionnaire,omitempty"`
	// The position of the questionnaire response within its overall lifecycle.
	Status *string `json:"status,omitempty"`
	// The subject of the questionnaire response.  This could be a patient, organization, practitioner, device, etc.  This is who/what the answers apply to, but is not necessarily the source of information.
	Subject *Reference `json:"subject,omitempty"`
	// The Encounter during which this questionnaire response was created or to which the creation of this record is tightly associated.
	Encounter *Reference `json:"encounter,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Authored *string `json:"authored,omitempty"`
	// Person who received the answers to the questions in the QuestionnaireResponse and recorded them in the system.
	Author *Reference `json:"author,omitempty"`
	// The person who answered the questions about the subject.
	Source *Reference `json:"source,omitempty"`
	// A group or question item from the original questionnaire for which answers are provided.
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

type _QuestionnaireResponse QuestionnaireResponse

// NewQuestionnaireResponse instantiates a new QuestionnaireResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireResponse(resourceType string) *QuestionnaireResponse {
	this := QuestionnaireResponse{}
	this.ResourceType = resourceType
	return &this
}

// NewQuestionnaireResponseWithDefaults instantiates a new QuestionnaireResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireResponseWithDefaults() *QuestionnaireResponse {
	this := QuestionnaireResponse{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *QuestionnaireResponse) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *QuestionnaireResponse) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuestionnaireResponse) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *QuestionnaireResponse) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *QuestionnaireResponse) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *QuestionnaireResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *QuestionnaireResponse) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *QuestionnaireResponse) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *QuestionnaireResponse) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *QuestionnaireResponse) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetIdentifier() Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret Identifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetIdentifierOk() (*Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given Identifier and assigns it to the Identifier field.
func (o *QuestionnaireResponse) SetIdentifier(v Identifier) {
	o.Identifier = &v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetBasedOn() []Reference {
	if o == nil || IsNil(o.BasedOn) {
		var ret []Reference
		return ret
	}
	return o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetBasedOnOk() ([]Reference, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given []Reference and assigns it to the BasedOn field.
func (o *QuestionnaireResponse) SetBasedOn(v []Reference) {
	o.BasedOn = v
}

// GetPartOf returns the PartOf field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetPartOf() []Reference {
	if o == nil || IsNil(o.PartOf) {
		var ret []Reference
		return ret
	}
	return o.PartOf
}

// GetPartOfOk returns a tuple with the PartOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetPartOfOk() ([]Reference, bool) {
	if o == nil || IsNil(o.PartOf) {
		return nil, false
	}
	return o.PartOf, true
}

// HasPartOf returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasPartOf() bool {
	if o != nil && !IsNil(o.PartOf) {
		return true
	}

	return false
}

// SetPartOf gets a reference to the given []Reference and assigns it to the PartOf field.
func (o *QuestionnaireResponse) SetPartOf(v []Reference) {
	o.PartOf = v
}

// GetQuestionnaire returns the Questionnaire field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetQuestionnaire() string {
	if o == nil || IsNil(o.Questionnaire) {
		var ret string
		return ret
	}
	return *o.Questionnaire
}

// GetQuestionnaireOk returns a tuple with the Questionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetQuestionnaireOk() (*string, bool) {
	if o == nil || IsNil(o.Questionnaire) {
		return nil, false
	}
	return o.Questionnaire, true
}

// HasQuestionnaire returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasQuestionnaire() bool {
	if o != nil && !IsNil(o.Questionnaire) {
		return true
	}

	return false
}

// SetQuestionnaire gets a reference to the given string and assigns it to the Questionnaire field.
func (o *QuestionnaireResponse) SetQuestionnaire(v string) {
	o.Questionnaire = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QuestionnaireResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetSubject() Reference {
	if o == nil || IsNil(o.Subject) {
		var ret Reference
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetSubjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Reference and assigns it to the Subject field.
func (o *QuestionnaireResponse) SetSubject(v Reference) {
	o.Subject = &v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *QuestionnaireResponse) SetEncounter(v Reference) {
	o.Encounter = &v
}

// GetAuthored returns the Authored field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetAuthored() string {
	if o == nil || IsNil(o.Authored) {
		var ret string
		return ret
	}
	return *o.Authored
}

// GetAuthoredOk returns a tuple with the Authored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetAuthoredOk() (*string, bool) {
	if o == nil || IsNil(o.Authored) {
		return nil, false
	}
	return o.Authored, true
}

// HasAuthored returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasAuthored() bool {
	if o != nil && !IsNil(o.Authored) {
		return true
	}

	return false
}

// SetAuthored gets a reference to the given string and assigns it to the Authored field.
func (o *QuestionnaireResponse) SetAuthored(v string) {
	o.Authored = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetAuthor() Reference {
	if o == nil || IsNil(o.Author) {
		var ret Reference
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetAuthorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Reference and assigns it to the Author field.
func (o *QuestionnaireResponse) SetAuthor(v Reference) {
	o.Author = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetSource() Reference {
	if o == nil || IsNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetSourceOk() (*Reference, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *QuestionnaireResponse) SetSource(v Reference) {
	o.Source = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *QuestionnaireResponse) GetItem() []QuestionnaireResponseItem {
	if o == nil || IsNil(o.Item) {
		var ret []QuestionnaireResponseItem
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireResponse) GetItemOk() ([]QuestionnaireResponseItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *QuestionnaireResponse) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []QuestionnaireResponseItem and assigns it to the Item field.
func (o *QuestionnaireResponse) SetItem(v []QuestionnaireResponseItem) {
	o.Item = v
}

func (o QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	if !IsNil(o.PartOf) {
		toSerialize["partOf"] = o.PartOf
	}
	if !IsNil(o.Questionnaire) {
		toSerialize["questionnaire"] = o.Questionnaire
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	if !IsNil(o.Authored) {
		toSerialize["authored"] = o.Authored
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

func (o *QuestionnaireResponse) UnmarshalJSON(data []byte) (err error) {
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

	varQuestionnaireResponse := _QuestionnaireResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestionnaireResponse)

	if err != nil {
		return err
	}

	*o = QuestionnaireResponse(varQuestionnaireResponse)

	return err
}

type NullableQuestionnaireResponse struct {
	value *QuestionnaireResponse
	isSet bool
}

func (v NullableQuestionnaireResponse) Get() *QuestionnaireResponse {
	return v.value
}

func (v *NullableQuestionnaireResponse) Set(val *QuestionnaireResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireResponse(val *QuestionnaireResponse) *NullableQuestionnaireResponse {
	return &NullableQuestionnaireResponse{value: val, isSet: true}
}

func (v NullableQuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


