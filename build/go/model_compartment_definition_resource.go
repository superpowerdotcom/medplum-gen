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

// checks if the CompartmentDefinitionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompartmentDefinitionResource{}

// CompartmentDefinitionResource A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinitionResource struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Code *string `json:"code,omitempty"`
	// The name of a search parameter that represents the link to the compartment. More than one may be listed because a resource may be linked to a compartment in more than one way,.
	Param []string `json:"param,omitempty"`
	// A sequence of Unicode characters
	Documentation *string `json:"documentation,omitempty"`
}

// NewCompartmentDefinitionResource instantiates a new CompartmentDefinitionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompartmentDefinitionResource() *CompartmentDefinitionResource {
	this := CompartmentDefinitionResource{}
	return &this
}

// NewCompartmentDefinitionResourceWithDefaults instantiates a new CompartmentDefinitionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompartmentDefinitionResourceWithDefaults() *CompartmentDefinitionResource {
	this := CompartmentDefinitionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompartmentDefinitionResource) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *CompartmentDefinitionResource) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *CompartmentDefinitionResource) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CompartmentDefinitionResource) SetCode(v string) {
	o.Code = &v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetParam() []string {
	if o == nil || IsNil(o.Param) {
		var ret []string
		return ret
	}
	return o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetParamOk() ([]string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given []string and assigns it to the Param field.
func (o *CompartmentDefinitionResource) SetParam(v []string) {
	o.Param = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *CompartmentDefinitionResource) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentDefinitionResource) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *CompartmentDefinitionResource) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *CompartmentDefinitionResource) SetDocumentation(v string) {
	o.Documentation = &v
}

func (o CompartmentDefinitionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompartmentDefinitionResource) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	return toSerialize, nil
}

type NullableCompartmentDefinitionResource struct {
	value *CompartmentDefinitionResource
	isSet bool
}

func (v NullableCompartmentDefinitionResource) Get() *CompartmentDefinitionResource {
	return v.value
}

func (v *NullableCompartmentDefinitionResource) Set(val *CompartmentDefinitionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCompartmentDefinitionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCompartmentDefinitionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompartmentDefinitionResource(val *CompartmentDefinitionResource) *NullableCompartmentDefinitionResource {
	return &NullableCompartmentDefinitionResource{value: val, isSet: true}
}

func (v NullableCompartmentDefinitionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompartmentDefinitionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


