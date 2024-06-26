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

// checks if the AppointmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentResponse{}

// AppointmentResponse A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
type AppointmentResponse struct {
	// This is a AppointmentResponse resource
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
	// This records identifiers associated with this appointment response concern that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Appointment that this response is replying to.
	Appointment Reference `json:"appointment"`
	// An instant in time - known at least to the second
	Start *string `json:"start,omitempty"`
	// An instant in time - known at least to the second
	End *string `json:"end,omitempty"`
	// Role of participant in the appointment.
	ParticipantType []CodeableConcept `json:"participantType,omitempty"`
	// A Person, Location, HealthcareService, or Device that is participating in the appointment.
	Actor *Reference `json:"actor,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	ParticipantStatus *string `json:"participantStatus,omitempty"`
	// A sequence of Unicode characters
	Comment *string `json:"comment,omitempty"`
}

type _AppointmentResponse AppointmentResponse

// NewAppointmentResponse instantiates a new AppointmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentResponse(resourceType string, appointment Reference) *AppointmentResponse {
	this := AppointmentResponse{}
	this.ResourceType = resourceType
	this.Appointment = appointment
	return &this
}

// NewAppointmentResponseWithDefaults instantiates a new AppointmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentResponseWithDefaults() *AppointmentResponse {
	this := AppointmentResponse{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *AppointmentResponse) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *AppointmentResponse) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppointmentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppointmentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppointmentResponse) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppointmentResponse) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppointmentResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *AppointmentResponse) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *AppointmentResponse) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *AppointmentResponse) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *AppointmentResponse) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AppointmentResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AppointmentResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AppointmentResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AppointmentResponse) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AppointmentResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *AppointmentResponse) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *AppointmentResponse) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *AppointmentResponse) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *AppointmentResponse) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AppointmentResponse) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AppointmentResponse) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *AppointmentResponse) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *AppointmentResponse) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *AppointmentResponse) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *AppointmentResponse) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AppointmentResponse) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AppointmentResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *AppointmentResponse) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetAppointment returns the Appointment field value
func (o *AppointmentResponse) GetAppointment() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Appointment
}

// GetAppointmentOk returns a tuple with the Appointment field value
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetAppointmentOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Appointment, true
}

// SetAppointment sets field value
func (o *AppointmentResponse) SetAppointment(v Reference) {
	o.Appointment = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *AppointmentResponse) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *AppointmentResponse) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *AppointmentResponse) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *AppointmentResponse) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *AppointmentResponse) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *AppointmentResponse) SetEnd(v string) {
	o.End = &v
}

// GetParticipantType returns the ParticipantType field value if set, zero value otherwise.
func (o *AppointmentResponse) GetParticipantType() []CodeableConcept {
	if o == nil || IsNil(o.ParticipantType) {
		var ret []CodeableConcept
		return ret
	}
	return o.ParticipantType
}

// GetParticipantTypeOk returns a tuple with the ParticipantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetParticipantTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ParticipantType) {
		return nil, false
	}
	return o.ParticipantType, true
}

// HasParticipantType returns a boolean if a field has been set.
func (o *AppointmentResponse) HasParticipantType() bool {
	if o != nil && !IsNil(o.ParticipantType) {
		return true
	}

	return false
}

// SetParticipantType gets a reference to the given []CodeableConcept and assigns it to the ParticipantType field.
func (o *AppointmentResponse) SetParticipantType(v []CodeableConcept) {
	o.ParticipantType = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *AppointmentResponse) GetActor() Reference {
	if o == nil || IsNil(o.Actor) {
		var ret Reference
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetActorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *AppointmentResponse) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given Reference and assigns it to the Actor field.
func (o *AppointmentResponse) SetActor(v Reference) {
	o.Actor = &v
}

// GetParticipantStatus returns the ParticipantStatus field value if set, zero value otherwise.
func (o *AppointmentResponse) GetParticipantStatus() string {
	if o == nil || IsNil(o.ParticipantStatus) {
		var ret string
		return ret
	}
	return *o.ParticipantStatus
}

// GetParticipantStatusOk returns a tuple with the ParticipantStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetParticipantStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ParticipantStatus) {
		return nil, false
	}
	return o.ParticipantStatus, true
}

// HasParticipantStatus returns a boolean if a field has been set.
func (o *AppointmentResponse) HasParticipantStatus() bool {
	if o != nil && !IsNil(o.ParticipantStatus) {
		return true
	}

	return false
}

// SetParticipantStatus gets a reference to the given string and assigns it to the ParticipantStatus field.
func (o *AppointmentResponse) SetParticipantStatus(v string) {
	o.ParticipantStatus = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AppointmentResponse) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AppointmentResponse) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AppointmentResponse) SetComment(v string) {
	o.Comment = &v
}

func (o AppointmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppointmentResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["appointment"] = o.Appointment
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.ParticipantType) {
		toSerialize["participantType"] = o.ParticipantType
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.ParticipantStatus) {
		toSerialize["participantStatus"] = o.ParticipantStatus
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

func (o *AppointmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"appointment",
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

	varAppointmentResponse := _AppointmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppointmentResponse)

	if err != nil {
		return err
	}

	*o = AppointmentResponse(varAppointmentResponse)

	return err
}

type NullableAppointmentResponse struct {
	value *AppointmentResponse
	isSet bool
}

func (v NullableAppointmentResponse) Get() *AppointmentResponse {
	return v.value
}

func (v *NullableAppointmentResponse) Set(val *AppointmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentResponse(val *AppointmentResponse) *NullableAppointmentResponse {
	return &NullableAppointmentResponse{value: val, isSet: true}
}

func (v NullableAppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


