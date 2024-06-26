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

// checks if the ExampleScenarioProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExampleScenarioProcess{}

// ExampleScenarioProcess Example of workflow instance.
type ExampleScenarioProcess struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Title *string `json:"title,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	Description *string `json:"description,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	PreConditions *string `json:"preConditions,omitempty"`
	// A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
	PostConditions *string `json:"postConditions,omitempty"`
	// Each step of the process.
	Step []ExampleScenarioStep `json:"step,omitempty"`
}

// NewExampleScenarioProcess instantiates a new ExampleScenarioProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleScenarioProcess() *ExampleScenarioProcess {
	this := ExampleScenarioProcess{}
	return &this
}

// NewExampleScenarioProcessWithDefaults instantiates a new ExampleScenarioProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleScenarioProcessWithDefaults() *ExampleScenarioProcess {
	this := ExampleScenarioProcess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExampleScenarioProcess) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ExampleScenarioProcess) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ExampleScenarioProcess) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ExampleScenarioProcess) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExampleScenarioProcess) SetDescription(v string) {
	o.Description = &v
}

// GetPreConditions returns the PreConditions field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetPreConditions() string {
	if o == nil || IsNil(o.PreConditions) {
		var ret string
		return ret
	}
	return *o.PreConditions
}

// GetPreConditionsOk returns a tuple with the PreConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetPreConditionsOk() (*string, bool) {
	if o == nil || IsNil(o.PreConditions) {
		return nil, false
	}
	return o.PreConditions, true
}

// HasPreConditions returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasPreConditions() bool {
	if o != nil && !IsNil(o.PreConditions) {
		return true
	}

	return false
}

// SetPreConditions gets a reference to the given string and assigns it to the PreConditions field.
func (o *ExampleScenarioProcess) SetPreConditions(v string) {
	o.PreConditions = &v
}

// GetPostConditions returns the PostConditions field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetPostConditions() string {
	if o == nil || IsNil(o.PostConditions) {
		var ret string
		return ret
	}
	return *o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetPostConditionsOk() (*string, bool) {
	if o == nil || IsNil(o.PostConditions) {
		return nil, false
	}
	return o.PostConditions, true
}

// HasPostConditions returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasPostConditions() bool {
	if o != nil && !IsNil(o.PostConditions) {
		return true
	}

	return false
}

// SetPostConditions gets a reference to the given string and assigns it to the PostConditions field.
func (o *ExampleScenarioProcess) SetPostConditions(v string) {
	o.PostConditions = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *ExampleScenarioProcess) GetStep() []ExampleScenarioStep {
	if o == nil || IsNil(o.Step) {
		var ret []ExampleScenarioStep
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleScenarioProcess) GetStepOk() ([]ExampleScenarioStep, bool) {
	if o == nil || IsNil(o.Step) {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *ExampleScenarioProcess) HasStep() bool {
	if o != nil && !IsNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given []ExampleScenarioStep and assigns it to the Step field.
func (o *ExampleScenarioProcess) SetStep(v []ExampleScenarioStep) {
	o.Step = v
}

func (o ExampleScenarioProcess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExampleScenarioProcess) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PreConditions) {
		toSerialize["preConditions"] = o.PreConditions
	}
	if !IsNil(o.PostConditions) {
		toSerialize["postConditions"] = o.PostConditions
	}
	if !IsNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	return toSerialize, nil
}

type NullableExampleScenarioProcess struct {
	value *ExampleScenarioProcess
	isSet bool
}

func (v NullableExampleScenarioProcess) Get() *ExampleScenarioProcess {
	return v.value
}

func (v *NullableExampleScenarioProcess) Set(val *ExampleScenarioProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleScenarioProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleScenarioProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleScenarioProcess(val *ExampleScenarioProcess) *NullableExampleScenarioProcess {
	return &NullableExampleScenarioProcess{value: val, isSet: true}
}

func (v NullableExampleScenarioProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleScenarioProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


