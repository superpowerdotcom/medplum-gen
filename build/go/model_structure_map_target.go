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

// checks if the StructureMapTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructureMapTarget{}

// StructureMapTarget A Map of relationships between 2 structures that can be used to transform data.
type StructureMapTarget struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Context *string `json:"context,omitempty"`
	// How to interpret the context.
	ContextType *string `json:"contextType,omitempty"`
	// A sequence of Unicode characters
	Element *string `json:"element,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Variable *string `json:"variable,omitempty"`
	// If field is a list, how to manage the list.
	ListMode []string `json:"listMode,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	ListRuleId *string `json:"listRuleId,omitempty"`
	// How the data is copied / created.
	Transform *string `json:"transform,omitempty"`
	// Parameters to the transform.
	Parameter []StructureMapParameter `json:"parameter,omitempty"`
}

// NewStructureMapTarget instantiates a new StructureMapTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructureMapTarget() *StructureMapTarget {
	this := StructureMapTarget{}
	return &this
}

// NewStructureMapTargetWithDefaults instantiates a new StructureMapTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructureMapTargetWithDefaults() *StructureMapTarget {
	this := StructureMapTarget{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructureMapTarget) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructureMapTarget) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructureMapTarget) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *StructureMapTarget) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *StructureMapTarget) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *StructureMapTarget) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *StructureMapTarget) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *StructureMapTarget) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *StructureMapTarget) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *StructureMapTarget) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *StructureMapTarget) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *StructureMapTarget) SetContext(v string) {
	o.Context = &v
}

// GetContextType returns the ContextType field value if set, zero value otherwise.
func (o *StructureMapTarget) GetContextType() string {
	if o == nil || IsNil(o.ContextType) {
		var ret string
		return ret
	}
	return *o.ContextType
}

// GetContextTypeOk returns a tuple with the ContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetContextTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContextType) {
		return nil, false
	}
	return o.ContextType, true
}

// HasContextType returns a boolean if a field has been set.
func (o *StructureMapTarget) HasContextType() bool {
	if o != nil && !IsNil(o.ContextType) {
		return true
	}

	return false
}

// SetContextType gets a reference to the given string and assigns it to the ContextType field.
func (o *StructureMapTarget) SetContextType(v string) {
	o.ContextType = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *StructureMapTarget) GetElement() string {
	if o == nil || IsNil(o.Element) {
		var ret string
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetElementOk() (*string, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *StructureMapTarget) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given string and assigns it to the Element field.
func (o *StructureMapTarget) SetElement(v string) {
	o.Element = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *StructureMapTarget) GetVariable() string {
	if o == nil || IsNil(o.Variable) {
		var ret string
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetVariableOk() (*string, bool) {
	if o == nil || IsNil(o.Variable) {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *StructureMapTarget) HasVariable() bool {
	if o != nil && !IsNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given string and assigns it to the Variable field.
func (o *StructureMapTarget) SetVariable(v string) {
	o.Variable = &v
}

// GetListMode returns the ListMode field value if set, zero value otherwise.
func (o *StructureMapTarget) GetListMode() []string {
	if o == nil || IsNil(o.ListMode) {
		var ret []string
		return ret
	}
	return o.ListMode
}

// GetListModeOk returns a tuple with the ListMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetListModeOk() ([]string, bool) {
	if o == nil || IsNil(o.ListMode) {
		return nil, false
	}
	return o.ListMode, true
}

// HasListMode returns a boolean if a field has been set.
func (o *StructureMapTarget) HasListMode() bool {
	if o != nil && !IsNil(o.ListMode) {
		return true
	}

	return false
}

// SetListMode gets a reference to the given []string and assigns it to the ListMode field.
func (o *StructureMapTarget) SetListMode(v []string) {
	o.ListMode = v
}

// GetListRuleId returns the ListRuleId field value if set, zero value otherwise.
func (o *StructureMapTarget) GetListRuleId() string {
	if o == nil || IsNil(o.ListRuleId) {
		var ret string
		return ret
	}
	return *o.ListRuleId
}

// GetListRuleIdOk returns a tuple with the ListRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetListRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListRuleId) {
		return nil, false
	}
	return o.ListRuleId, true
}

// HasListRuleId returns a boolean if a field has been set.
func (o *StructureMapTarget) HasListRuleId() bool {
	if o != nil && !IsNil(o.ListRuleId) {
		return true
	}

	return false
}

// SetListRuleId gets a reference to the given string and assigns it to the ListRuleId field.
func (o *StructureMapTarget) SetListRuleId(v string) {
	o.ListRuleId = &v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *StructureMapTarget) GetTransform() string {
	if o == nil || IsNil(o.Transform) {
		var ret string
		return ret
	}
	return *o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetTransformOk() (*string, bool) {
	if o == nil || IsNil(o.Transform) {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *StructureMapTarget) HasTransform() bool {
	if o != nil && !IsNil(o.Transform) {
		return true
	}

	return false
}

// SetTransform gets a reference to the given string and assigns it to the Transform field.
func (o *StructureMapTarget) SetTransform(v string) {
	o.Transform = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *StructureMapTarget) GetParameter() []StructureMapParameter {
	if o == nil || IsNil(o.Parameter) {
		var ret []StructureMapParameter
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructureMapTarget) GetParameterOk() ([]StructureMapParameter, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *StructureMapTarget) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []StructureMapParameter and assigns it to the Parameter field.
func (o *StructureMapTarget) SetParameter(v []StructureMapParameter) {
	o.Parameter = v
}

func (o StructureMapTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructureMapTarget) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.ContextType) {
		toSerialize["contextType"] = o.ContextType
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	if !IsNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	if !IsNil(o.ListMode) {
		toSerialize["listMode"] = o.ListMode
	}
	if !IsNil(o.ListRuleId) {
		toSerialize["listRuleId"] = o.ListRuleId
	}
	if !IsNil(o.Transform) {
		toSerialize["transform"] = o.Transform
	}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	return toSerialize, nil
}

type NullableStructureMapTarget struct {
	value *StructureMapTarget
	isSet bool
}

func (v NullableStructureMapTarget) Get() *StructureMapTarget {
	return v.value
}

func (v *NullableStructureMapTarget) Set(val *StructureMapTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableStructureMapTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableStructureMapTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructureMapTarget(val *StructureMapTarget) *NullableStructureMapTarget {
	return &NullableStructureMapTarget{value: val, isSet: true}
}

func (v NullableStructureMapTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructureMapTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


