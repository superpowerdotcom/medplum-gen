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

// checks if the ConceptMapGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConceptMapGroup{}

// ConceptMapGroup A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMapGroup struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// String of characters used to identify a name or a resource
	Source *string `json:"source,omitempty"`
	// A sequence of Unicode characters
	SourceVersion *string `json:"sourceVersion,omitempty"`
	// String of characters used to identify a name or a resource
	Target *string `json:"target,omitempty"`
	// A sequence of Unicode characters
	TargetVersion *string `json:"targetVersion,omitempty"`
	// Mappings for an individual concept in the source to one or more concepts in the target.
	Element []ConceptMapElement `json:"element"`
	// What to do when there is no mapping for the source concept. \"Unmapped\" does not include codes that are unmatched, and the unmapped element is ignored in a code is specified to have equivalence = unmatched.
	Unmapped *ConceptMapUnmapped `json:"unmapped,omitempty"`
}

type _ConceptMapGroup ConceptMapGroup

// NewConceptMapGroup instantiates a new ConceptMapGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConceptMapGroup(element []ConceptMapElement) *ConceptMapGroup {
	this := ConceptMapGroup{}
	this.Element = element
	return &this
}

// NewConceptMapGroupWithDefaults instantiates a new ConceptMapGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConceptMapGroupWithDefaults() *ConceptMapGroup {
	this := ConceptMapGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConceptMapGroup) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ConceptMapGroup) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ConceptMapGroup) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConceptMapGroup) SetSource(v string) {
	o.Source = &v
}

// GetSourceVersion returns the SourceVersion field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetSourceVersion() string {
	if o == nil || IsNil(o.SourceVersion) {
		var ret string
		return ret
	}
	return *o.SourceVersion
}

// GetSourceVersionOk returns a tuple with the SourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetSourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SourceVersion) {
		return nil, false
	}
	return o.SourceVersion, true
}

// HasSourceVersion returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasSourceVersion() bool {
	if o != nil && !IsNil(o.SourceVersion) {
		return true
	}

	return false
}

// SetSourceVersion gets a reference to the given string and assigns it to the SourceVersion field.
func (o *ConceptMapGroup) SetSourceVersion(v string) {
	o.SourceVersion = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ConceptMapGroup) SetTarget(v string) {
	o.Target = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetTargetVersion() string {
	if o == nil || IsNil(o.TargetVersion) {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetTargetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetVersion) {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasTargetVersion() bool {
	if o != nil && !IsNil(o.TargetVersion) {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *ConceptMapGroup) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetElement returns the Element field value
func (o *ConceptMapGroup) GetElement() []ConceptMapElement {
	if o == nil {
		var ret []ConceptMapElement
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetElementOk() ([]ConceptMapElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Element, true
}

// SetElement sets field value
func (o *ConceptMapGroup) SetElement(v []ConceptMapElement) {
	o.Element = v
}

// GetUnmapped returns the Unmapped field value if set, zero value otherwise.
func (o *ConceptMapGroup) GetUnmapped() ConceptMapUnmapped {
	if o == nil || IsNil(o.Unmapped) {
		var ret ConceptMapUnmapped
		return ret
	}
	return *o.Unmapped
}

// GetUnmappedOk returns a tuple with the Unmapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptMapGroup) GetUnmappedOk() (*ConceptMapUnmapped, bool) {
	if o == nil || IsNil(o.Unmapped) {
		return nil, false
	}
	return o.Unmapped, true
}

// HasUnmapped returns a boolean if a field has been set.
func (o *ConceptMapGroup) HasUnmapped() bool {
	if o != nil && !IsNil(o.Unmapped) {
		return true
	}

	return false
}

// SetUnmapped gets a reference to the given ConceptMapUnmapped and assigns it to the Unmapped field.
func (o *ConceptMapGroup) SetUnmapped(v ConceptMapUnmapped) {
	o.Unmapped = &v
}

func (o ConceptMapGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConceptMapGroup) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.SourceVersion) {
		toSerialize["sourceVersion"] = o.SourceVersion
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.TargetVersion) {
		toSerialize["targetVersion"] = o.TargetVersion
	}
	toSerialize["element"] = o.Element
	if !IsNil(o.Unmapped) {
		toSerialize["unmapped"] = o.Unmapped
	}
	return toSerialize, nil
}

func (o *ConceptMapGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"element",
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

	varConceptMapGroup := _ConceptMapGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConceptMapGroup)

	if err != nil {
		return err
	}

	*o = ConceptMapGroup(varConceptMapGroup)

	return err
}

type NullableConceptMapGroup struct {
	value *ConceptMapGroup
	isSet bool
}

func (v NullableConceptMapGroup) Get() *ConceptMapGroup {
	return v.value
}

func (v *NullableConceptMapGroup) Set(val *ConceptMapGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableConceptMapGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableConceptMapGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConceptMapGroup(val *ConceptMapGroup) *NullableConceptMapGroup {
	return &NullableConceptMapGroup{value: val, isSet: true}
}

func (v NullableConceptMapGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConceptMapGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


