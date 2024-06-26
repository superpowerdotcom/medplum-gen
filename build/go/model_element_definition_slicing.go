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

// checks if the ElementDefinitionSlicing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ElementDefinitionSlicing{}

// ElementDefinitionSlicing Captures constraints on each element within the resource, profile, or extension.
type ElementDefinitionSlicing struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
	Discriminator []ElementDefinitionDiscriminator `json:"discriminator,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Value of \"true\" or \"false\"
	Ordered *bool `json:"ordered,omitempty"`
	// Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	Rules *string `json:"rules,omitempty"`
}

// NewElementDefinitionSlicing instantiates a new ElementDefinitionSlicing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElementDefinitionSlicing() *ElementDefinitionSlicing {
	this := ElementDefinitionSlicing{}
	return &this
}

// NewElementDefinitionSlicingWithDefaults instantiates a new ElementDefinitionSlicing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElementDefinitionSlicingWithDefaults() *ElementDefinitionSlicing {
	this := ElementDefinitionSlicing{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ElementDefinitionSlicing) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ElementDefinitionSlicing) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ElementDefinitionSlicing) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDiscriminator returns the Discriminator field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetDiscriminator() []ElementDefinitionDiscriminator {
	if o == nil || IsNil(o.Discriminator) {
		var ret []ElementDefinitionDiscriminator
		return ret
	}
	return o.Discriminator
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetDiscriminatorOk() ([]ElementDefinitionDiscriminator, bool) {
	if o == nil || IsNil(o.Discriminator) {
		return nil, false
	}
	return o.Discriminator, true
}

// HasDiscriminator returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasDiscriminator() bool {
	if o != nil && !IsNil(o.Discriminator) {
		return true
	}

	return false
}

// SetDiscriminator gets a reference to the given []ElementDefinitionDiscriminator and assigns it to the Discriminator field.
func (o *ElementDefinitionSlicing) SetDiscriminator(v []ElementDefinitionDiscriminator) {
	o.Discriminator = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ElementDefinitionSlicing) SetDescription(v string) {
	o.Description = &v
}

// GetOrdered returns the Ordered field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetOrdered() bool {
	if o == nil || IsNil(o.Ordered) {
		var ret bool
		return ret
	}
	return *o.Ordered
}

// GetOrderedOk returns a tuple with the Ordered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetOrderedOk() (*bool, bool) {
	if o == nil || IsNil(o.Ordered) {
		return nil, false
	}
	return o.Ordered, true
}

// HasOrdered returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasOrdered() bool {
	if o != nil && !IsNil(o.Ordered) {
		return true
	}

	return false
}

// SetOrdered gets a reference to the given bool and assigns it to the Ordered field.
func (o *ElementDefinitionSlicing) SetOrdered(v bool) {
	o.Ordered = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ElementDefinitionSlicing) GetRules() string {
	if o == nil || IsNil(o.Rules) {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionSlicing) GetRulesOk() (*string, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ElementDefinitionSlicing) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
func (o *ElementDefinitionSlicing) SetRules(v string) {
	o.Rules = &v
}

func (o ElementDefinitionSlicing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ElementDefinitionSlicing) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Discriminator) {
		toSerialize["discriminator"] = o.Discriminator
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Ordered) {
		toSerialize["ordered"] = o.Ordered
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableElementDefinitionSlicing struct {
	value *ElementDefinitionSlicing
	isSet bool
}

func (v NullableElementDefinitionSlicing) Get() *ElementDefinitionSlicing {
	return v.value
}

func (v *NullableElementDefinitionSlicing) Set(val *ElementDefinitionSlicing) {
	v.value = val
	v.isSet = true
}

func (v NullableElementDefinitionSlicing) IsSet() bool {
	return v.isSet
}

func (v *NullableElementDefinitionSlicing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElementDefinitionSlicing(val *ElementDefinitionSlicing) *NullableElementDefinitionSlicing {
	return &NullableElementDefinitionSlicing{value: val, isSet: true}
}

func (v NullableElementDefinitionSlicing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElementDefinitionSlicing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


