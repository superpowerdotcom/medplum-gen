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

// checks if the Practitioner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Practitioner{}

// Practitioner A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner struct {
	// This is a Practitioner resource
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
	// An identifier that applies to this person in this role.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Value of \"true\" or \"false\"
	Active *bool `json:"active,omitempty"`
	// The name(s) associated with the practitioner.
	Name []HumanName `json:"name,omitempty"`
	// A contact detail for the practitioner, e.g. a telephone number or an email address.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Address(es) of the practitioner that are not role specific (typically home address).  Work addresses are not typically entered in this property as they are usually role dependent.
	Address []Address `json:"address,omitempty"`
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender *string `json:"gender,omitempty"`
	// A date or partial date (e.g. just year or year + month). There is no time zone. The format is a union of the schema types gYear, gYearMonth and date.  Dates SHALL be valid dates.
	BirthDate *string `json:"birthDate,omitempty"`
	// Image of the person.
	Photo []Attachment `json:"photo,omitempty"`
	// The official certifications, training, and licenses that authorize or otherwise pertain to the provision of care by the practitioner.  For example, a medical license issued by a medical board authorizing the practitioner to practice medicine within a certian locality.
	Qualification []PractitionerQualification `json:"qualification,omitempty"`
	// A language the practitioner can use in patient communication.
	Communication []CodeableConcept `json:"communication,omitempty"`
}

type _Practitioner Practitioner

// NewPractitioner instantiates a new Practitioner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPractitioner(resourceType string) *Practitioner {
	this := Practitioner{}
	this.ResourceType = resourceType
	return &this
}

// NewPractitionerWithDefaults instantiates a new Practitioner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPractitionerWithDefaults() *Practitioner {
	this := Practitioner{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Practitioner) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Practitioner) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Practitioner) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Practitioner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Practitioner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Practitioner) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Practitioner) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Practitioner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Practitioner) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Practitioner) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Practitioner) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Practitioner) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Practitioner) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Practitioner) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Practitioner) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Practitioner) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Practitioner) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Practitioner) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Practitioner) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Practitioner) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Practitioner) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Practitioner) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Practitioner) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Practitioner) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Practitioner) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Practitioner) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Practitioner) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Practitioner) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Practitioner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Practitioner) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Practitioner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Practitioner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Practitioner) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Practitioner) GetName() []HumanName {
	if o == nil || IsNil(o.Name) {
		var ret []HumanName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetNameOk() ([]HumanName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Practitioner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []HumanName and assigns it to the Name field.
func (o *Practitioner) SetName(v []HumanName) {
	o.Name = v
}

// GetTelecom returns the Telecom field value if set, zero value otherwise.
func (o *Practitioner) GetTelecom() []ContactPoint {
	if o == nil || IsNil(o.Telecom) {
		var ret []ContactPoint
		return ret
	}
	return o.Telecom
}

// GetTelecomOk returns a tuple with the Telecom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetTelecomOk() ([]ContactPoint, bool) {
	if o == nil || IsNil(o.Telecom) {
		return nil, false
	}
	return o.Telecom, true
}

// HasTelecom returns a boolean if a field has been set.
func (o *Practitioner) HasTelecom() bool {
	if o != nil && !IsNil(o.Telecom) {
		return true
	}

	return false
}

// SetTelecom gets a reference to the given []ContactPoint and assigns it to the Telecom field.
func (o *Practitioner) SetTelecom(v []ContactPoint) {
	o.Telecom = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Practitioner) GetAddress() []Address {
	if o == nil || IsNil(o.Address) {
		var ret []Address
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetAddressOk() ([]Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Practitioner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []Address and assigns it to the Address field.
func (o *Practitioner) SetAddress(v []Address) {
	o.Address = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Practitioner) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Practitioner) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *Practitioner) SetGender(v string) {
	o.Gender = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *Practitioner) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *Practitioner) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *Practitioner) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *Practitioner) GetPhoto() []Attachment {
	if o == nil || IsNil(o.Photo) {
		var ret []Attachment
		return ret
	}
	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetPhotoOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *Practitioner) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given []Attachment and assigns it to the Photo field.
func (o *Practitioner) SetPhoto(v []Attachment) {
	o.Photo = v
}

// GetQualification returns the Qualification field value if set, zero value otherwise.
func (o *Practitioner) GetQualification() []PractitionerQualification {
	if o == nil || IsNil(o.Qualification) {
		var ret []PractitionerQualification
		return ret
	}
	return o.Qualification
}

// GetQualificationOk returns a tuple with the Qualification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetQualificationOk() ([]PractitionerQualification, bool) {
	if o == nil || IsNil(o.Qualification) {
		return nil, false
	}
	return o.Qualification, true
}

// HasQualification returns a boolean if a field has been set.
func (o *Practitioner) HasQualification() bool {
	if o != nil && !IsNil(o.Qualification) {
		return true
	}

	return false
}

// SetQualification gets a reference to the given []PractitionerQualification and assigns it to the Qualification field.
func (o *Practitioner) SetQualification(v []PractitionerQualification) {
	o.Qualification = v
}

// GetCommunication returns the Communication field value if set, zero value otherwise.
func (o *Practitioner) GetCommunication() []CodeableConcept {
	if o == nil || IsNil(o.Communication) {
		var ret []CodeableConcept
		return ret
	}
	return o.Communication
}

// GetCommunicationOk returns a tuple with the Communication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Practitioner) GetCommunicationOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Communication) {
		return nil, false
	}
	return o.Communication, true
}

// HasCommunication returns a boolean if a field has been set.
func (o *Practitioner) HasCommunication() bool {
	if o != nil && !IsNil(o.Communication) {
		return true
	}

	return false
}

// SetCommunication gets a reference to the given []CodeableConcept and assigns it to the Communication field.
func (o *Practitioner) SetCommunication(v []CodeableConcept) {
	o.Communication = v
}

func (o Practitioner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Practitioner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Telecom) {
		toSerialize["telecom"] = o.Telecom
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.Qualification) {
		toSerialize["qualification"] = o.Qualification
	}
	if !IsNil(o.Communication) {
		toSerialize["communication"] = o.Communication
	}
	return toSerialize, nil
}

func (o *Practitioner) UnmarshalJSON(data []byte) (err error) {
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

	varPractitioner := _Practitioner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPractitioner)

	if err != nil {
		return err
	}

	*o = Practitioner(varPractitioner)

	return err
}

type NullablePractitioner struct {
	value *Practitioner
	isSet bool
}

func (v NullablePractitioner) Get() *Practitioner {
	return v.value
}

func (v *NullablePractitioner) Set(val *Practitioner) {
	v.value = val
	v.isSet = true
}

func (v NullablePractitioner) IsSet() bool {
	return v.isSet
}

func (v *NullablePractitioner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePractitioner(val *Practitioner) *NullablePractitioner {
	return &NullablePractitioner{value: val, isSet: true}
}

func (v NullablePractitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePractitioner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

