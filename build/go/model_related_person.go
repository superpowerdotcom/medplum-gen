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

// checks if the RelatedPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedPerson{}

// RelatedPerson Information about a person that is involved in the care for a patient, but who is not the target of healthcare, nor has a formal responsibility in the care process.
type RelatedPerson struct {
	// This is a RelatedPerson resource
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
	// Identifier for a person within a particular scope.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Value of \"true\" or \"false\"
	Active *bool `json:"active,omitempty"`
	// The patient this person is related to.
	Patient Reference `json:"patient"`
	// The nature of the relationship between a patient and the related person.
	Relationship []CodeableConcept `json:"relationship,omitempty"`
	// A name associated with the person.
	Name []HumanName `json:"name,omitempty"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender *string `json:"gender,omitempty"`
	// A date or partial date (e.g. just year or year + month). There is no time zone. The format is a union of the schema types gYear, gYearMonth and date.  Dates SHALL be valid dates.
	BirthDate *string `json:"birthDate,omitempty"`
	// Address where the related person can be contacted or visited.
	Address []Address `json:"address,omitempty"`
	// Image of the person.
	Photo []Attachment `json:"photo,omitempty"`
	// The period of time during which this relationship is or was active. If there are no dates defined, then the interval is unknown.
	Period *Period `json:"period,omitempty"`
	// A language which may be used to communicate with about the patient's health.
	Communication []RelatedPersonCommunication `json:"communication,omitempty"`
}

type _RelatedPerson RelatedPerson

// NewRelatedPerson instantiates a new RelatedPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedPerson(resourceType string, patient Reference) *RelatedPerson {
	this := RelatedPerson{}
	this.ResourceType = resourceType
	this.Patient = patient
	return &this
}

// NewRelatedPersonWithDefaults instantiates a new RelatedPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedPersonWithDefaults() *RelatedPerson {
	this := RelatedPerson{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *RelatedPerson) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *RelatedPerson) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelatedPerson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelatedPerson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelatedPerson) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RelatedPerson) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RelatedPerson) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *RelatedPerson) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *RelatedPerson) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *RelatedPerson) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *RelatedPerson) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *RelatedPerson) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *RelatedPerson) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *RelatedPerson) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RelatedPerson) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RelatedPerson) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *RelatedPerson) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *RelatedPerson) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *RelatedPerson) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *RelatedPerson) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *RelatedPerson) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *RelatedPerson) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *RelatedPerson) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *RelatedPerson) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *RelatedPerson) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *RelatedPerson) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *RelatedPerson) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *RelatedPerson) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *RelatedPerson) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RelatedPerson) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RelatedPerson) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RelatedPerson) SetActive(v bool) {
	o.Active = &v
}

// GetPatient returns the Patient field value
func (o *RelatedPerson) GetPatient() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Patient
}

// GetPatientOk returns a tuple with the Patient field value
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetPatientOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patient, true
}

// SetPatient sets field value
func (o *RelatedPerson) SetPatient(v Reference) {
	o.Patient = v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *RelatedPerson) GetRelationship() []CodeableConcept {
	if o == nil || IsNil(o.Relationship) {
		var ret []CodeableConcept
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetRelationshipOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *RelatedPerson) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given []CodeableConcept and assigns it to the Relationship field.
func (o *RelatedPerson) SetRelationship(v []CodeableConcept) {
	o.Relationship = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RelatedPerson) GetName() []HumanName {
	if o == nil || IsNil(o.Name) {
		var ret []HumanName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetNameOk() ([]HumanName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RelatedPerson) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []HumanName and assigns it to the Name field.
func (o *RelatedPerson) SetName(v []HumanName) {
	o.Name = v
}

// GetTelecom returns the Telecom field value if set, zero value otherwise.
func (o *RelatedPerson) GetTelecom() []ContactPoint {
	if o == nil || IsNil(o.Telecom) {
		var ret []ContactPoint
		return ret
	}
	return o.Telecom
}

// GetTelecomOk returns a tuple with the Telecom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetTelecomOk() ([]ContactPoint, bool) {
	if o == nil || IsNil(o.Telecom) {
		return nil, false
	}
	return o.Telecom, true
}

// HasTelecom returns a boolean if a field has been set.
func (o *RelatedPerson) HasTelecom() bool {
	if o != nil && !IsNil(o.Telecom) {
		return true
	}

	return false
}

// SetTelecom gets a reference to the given []ContactPoint and assigns it to the Telecom field.
func (o *RelatedPerson) SetTelecom(v []ContactPoint) {
	o.Telecom = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *RelatedPerson) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *RelatedPerson) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *RelatedPerson) SetGender(v string) {
	o.Gender = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *RelatedPerson) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *RelatedPerson) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *RelatedPerson) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RelatedPerson) GetAddress() []Address {
	if o == nil || IsNil(o.Address) {
		var ret []Address
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetAddressOk() ([]Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RelatedPerson) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []Address and assigns it to the Address field.
func (o *RelatedPerson) SetAddress(v []Address) {
	o.Address = v
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *RelatedPerson) GetPhoto() []Attachment {
	if o == nil || IsNil(o.Photo) {
		var ret []Attachment
		return ret
	}
	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetPhotoOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *RelatedPerson) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given []Attachment and assigns it to the Photo field.
func (o *RelatedPerson) SetPhoto(v []Attachment) {
	o.Photo = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *RelatedPerson) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *RelatedPerson) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *RelatedPerson) SetPeriod(v Period) {
	o.Period = &v
}

// GetCommunication returns the Communication field value if set, zero value otherwise.
func (o *RelatedPerson) GetCommunication() []RelatedPersonCommunication {
	if o == nil || IsNil(o.Communication) {
		var ret []RelatedPersonCommunication
		return ret
	}
	return o.Communication
}

// GetCommunicationOk returns a tuple with the Communication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPerson) GetCommunicationOk() ([]RelatedPersonCommunication, bool) {
	if o == nil || IsNil(o.Communication) {
		return nil, false
	}
	return o.Communication, true
}

// HasCommunication returns a boolean if a field has been set.
func (o *RelatedPerson) HasCommunication() bool {
	if o != nil && !IsNil(o.Communication) {
		return true
	}

	return false
}

// SetCommunication gets a reference to the given []RelatedPersonCommunication and assigns it to the Communication field.
func (o *RelatedPerson) SetCommunication(v []RelatedPersonCommunication) {
	o.Communication = v
}

func (o RelatedPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedPerson) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["patient"] = o.Patient
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Telecom) {
		toSerialize["telecom"] = o.Telecom
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Communication) {
		toSerialize["communication"] = o.Communication
	}
	return toSerialize, nil
}

func (o *RelatedPerson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"patient",
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

	varRelatedPerson := _RelatedPerson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelatedPerson)

	if err != nil {
		return err
	}

	*o = RelatedPerson(varRelatedPerson)

	return err
}

type NullableRelatedPerson struct {
	value *RelatedPerson
	isSet bool
}

func (v NullableRelatedPerson) Get() *RelatedPerson {
	return v.value
}

func (v *NullableRelatedPerson) Set(val *RelatedPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedPerson(val *RelatedPerson) *NullableRelatedPerson {
	return &NullableRelatedPerson{value: val, isSet: true}
}

func (v NullableRelatedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


