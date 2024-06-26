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

// checks if the SmartAppLaunch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartAppLaunch{}

// SmartAppLaunch This resource contains context details for a SMART App Launch.
type SmartAppLaunch struct {
	// This is a SmartAppLaunch resource
	ResourceType string `json:"resourceType"`
	// A sequence of Unicode characters
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
	Contained []string `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Optional patient indicating that the app was launched in the patient context.
	Patient *Reference `json:"patient,omitempty"`
	// Optional encounter indicating that the app was launched in the encounter context.
	Encounter *Reference `json:"encounter,omitempty"`
}

type _SmartAppLaunch SmartAppLaunch

// NewSmartAppLaunch instantiates a new SmartAppLaunch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartAppLaunch(resourceType string) *SmartAppLaunch {
	this := SmartAppLaunch{}
	this.ResourceType = resourceType
	return &this
}

// NewSmartAppLaunchWithDefaults instantiates a new SmartAppLaunch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartAppLaunchWithDefaults() *SmartAppLaunch {
	this := SmartAppLaunch{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *SmartAppLaunch) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *SmartAppLaunch) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SmartAppLaunch) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *SmartAppLaunch) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *SmartAppLaunch) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *SmartAppLaunch) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *SmartAppLaunch) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetContained() []string {
	if o == nil || IsNil(o.Contained) {
		var ret []string
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetContainedOk() ([]string, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []string and assigns it to the Contained field.
func (o *SmartAppLaunch) SetContained(v []string) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SmartAppLaunch) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SmartAppLaunch) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetPatient returns the Patient field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetPatient() Reference {
	if o == nil || IsNil(o.Patient) {
		var ret Reference
		return ret
	}
	return *o.Patient
}

// GetPatientOk returns a tuple with the Patient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetPatientOk() (*Reference, bool) {
	if o == nil || IsNil(o.Patient) {
		return nil, false
	}
	return o.Patient, true
}

// HasPatient returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasPatient() bool {
	if o != nil && !IsNil(o.Patient) {
		return true
	}

	return false
}

// SetPatient gets a reference to the given Reference and assigns it to the Patient field.
func (o *SmartAppLaunch) SetPatient(v Reference) {
	o.Patient = &v
}

// GetEncounter returns the Encounter field value if set, zero value otherwise.
func (o *SmartAppLaunch) GetEncounter() Reference {
	if o == nil || IsNil(o.Encounter) {
		var ret Reference
		return ret
	}
	return *o.Encounter
}

// GetEncounterOk returns a tuple with the Encounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAppLaunch) GetEncounterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Encounter) {
		return nil, false
	}
	return o.Encounter, true
}

// HasEncounter returns a boolean if a field has been set.
func (o *SmartAppLaunch) HasEncounter() bool {
	if o != nil && !IsNil(o.Encounter) {
		return true
	}

	return false
}

// SetEncounter gets a reference to the given Reference and assigns it to the Encounter field.
func (o *SmartAppLaunch) SetEncounter(v Reference) {
	o.Encounter = &v
}

func (o SmartAppLaunch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartAppLaunch) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Patient) {
		toSerialize["patient"] = o.Patient
	}
	if !IsNil(o.Encounter) {
		toSerialize["encounter"] = o.Encounter
	}
	return toSerialize, nil
}

func (o *SmartAppLaunch) UnmarshalJSON(data []byte) (err error) {
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

	varSmartAppLaunch := _SmartAppLaunch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartAppLaunch)

	if err != nil {
		return err
	}

	*o = SmartAppLaunch(varSmartAppLaunch)

	return err
}

type NullableSmartAppLaunch struct {
	value *SmartAppLaunch
	isSet bool
}

func (v NullableSmartAppLaunch) Get() *SmartAppLaunch {
	return v.value
}

func (v *NullableSmartAppLaunch) Set(val *SmartAppLaunch) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartAppLaunch) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartAppLaunch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartAppLaunch(val *SmartAppLaunch) *NullableSmartAppLaunch {
	return &NullableSmartAppLaunch{value: val, isSet: true}
}

func (v NullableSmartAppLaunch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartAppLaunch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


