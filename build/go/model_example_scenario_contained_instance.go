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

// checks if the ExampleScenarioContainedInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExampleScenarioContainedInstance{}

// ExampleScenarioContainedInstance Example of workflow instance.
type ExampleScenarioContainedInstance struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	ResourceId *string `json:"resourceId,omitempty"`
	// A sequence of Unicode characters
	VersionId *string `json:"versionId,omitempty"`
}

// NewExampleScenarioContainedInstance instantiates a new ExampleScenarioContainedInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleScenarioContainedInstance() *ExampleScenarioContainedInstance {
	this := ExampleScenarioContainedInstance{}
	return &this
}

// NewExampleScenarioContainedInstanceWithDefaults instantiates a new ExampleScenarioContainedInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleScenarioContainedInstanceWithDefaults() *ExampleScenarioContainedInstance {
	this := ExampleScenarioContainedInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExampleScenarioContainedInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioContainedInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExampleScenarioContainedInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExampleScenarioContainedInstance) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExampleScenarioContainedInstance) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioContainedInstance) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExampleScenarioContainedInstance) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExampleScenarioContainedInstance) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExampleScenarioContainedInstance) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioContainedInstance) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExampleScenarioContainedInstance) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExampleScenarioContainedInstance) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *ExampleScenarioContainedInstance) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioContainedInstance) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *ExampleScenarioContainedInstance) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *ExampleScenarioContainedInstance) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *ExampleScenarioContainedInstance) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioContainedInstance) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *ExampleScenarioContainedInstance) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *ExampleScenarioContainedInstance) SetVersionId(v string) {
	o.VersionId = &v
}

func (o ExampleScenarioContainedInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExampleScenarioContainedInstance) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableExampleScenarioContainedInstance struct {
	value *ExampleScenarioContainedInstance
	isSet bool
}

func (v NullableExampleScenarioContainedInstance) Get() *ExampleScenarioContainedInstance {
	return v.value
}

func (v *NullableExampleScenarioContainedInstance) Set(val *ExampleScenarioContainedInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleScenarioContainedInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleScenarioContainedInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleScenarioContainedInstance(val *ExampleScenarioContainedInstance) *NullableExampleScenarioContainedInstance {
	return &NullableExampleScenarioContainedInstance{value: val, isSet: true}
}

func (v NullableExampleScenarioContainedInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleScenarioContainedInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


