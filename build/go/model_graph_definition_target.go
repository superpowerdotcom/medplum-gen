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

// checks if the GraphDefinitionTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphDefinitionTarget{}

// GraphDefinitionTarget A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinitionTarget struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Type *string `json:"type,omitempty"`
	// A sequence of Unicode characters
	Params *string `json:"params,omitempty"`
	// A URI that is a reference to a canonical URL on a FHIR resource
	Profile *string `json:"profile,omitempty"`
	// Compartment Consistency Rules.
	Compartment []GraphDefinitionCompartment `json:"compartment,omitempty"`
	// Additional links from target resource.
	Link []GraphDefinitionLink `json:"link,omitempty"`
}

// NewGraphDefinitionTarget instantiates a new GraphDefinitionTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphDefinitionTarget() *GraphDefinitionTarget {
	this := GraphDefinitionTarget{}
	return &this
}

// NewGraphDefinitionTargetWithDefaults instantiates a new GraphDefinitionTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphDefinitionTargetWithDefaults() *GraphDefinitionTarget {
	this := GraphDefinitionTarget{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GraphDefinitionTarget) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *GraphDefinitionTarget) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *GraphDefinitionTarget) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GraphDefinitionTarget) SetType(v string) {
	o.Type = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetParams() string {
	if o == nil || IsNil(o.Params) {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetParamsOk() (*string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *GraphDefinitionTarget) SetParams(v string) {
	o.Params = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *GraphDefinitionTarget) SetProfile(v string) {
	o.Profile = &v
}

// GetCompartment returns the Compartment field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetCompartment() []GraphDefinitionCompartment {
	if o == nil || IsNil(o.Compartment) {
		var ret []GraphDefinitionCompartment
		return ret
	}
	return o.Compartment
}

// GetCompartmentOk returns a tuple with the Compartment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetCompartmentOk() ([]GraphDefinitionCompartment, bool) {
	if o == nil || IsNil(o.Compartment) {
		return nil, false
	}
	return o.Compartment, true
}

// HasCompartment returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasCompartment() bool {
	if o != nil && !IsNil(o.Compartment) {
		return true
	}

	return false
}

// SetCompartment gets a reference to the given []GraphDefinitionCompartment and assigns it to the Compartment field.
func (o *GraphDefinitionTarget) SetCompartment(v []GraphDefinitionCompartment) {
	o.Compartment = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *GraphDefinitionTarget) GetLink() []GraphDefinitionLink {
	if o == nil || IsNil(o.Link) {
		var ret []GraphDefinitionLink
		return ret
	}
	return o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDefinitionTarget) GetLinkOk() ([]GraphDefinitionLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *GraphDefinitionTarget) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given []GraphDefinitionLink and assigns it to the Link field.
func (o *GraphDefinitionTarget) SetLink(v []GraphDefinitionLink) {
	o.Link = v
}

func (o GraphDefinitionTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphDefinitionTarget) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Compartment) {
		toSerialize["compartment"] = o.Compartment
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableGraphDefinitionTarget struct {
	value *GraphDefinitionTarget
	isSet bool
}

func (v NullableGraphDefinitionTarget) Get() *GraphDefinitionTarget {
	return v.value
}

func (v *NullableGraphDefinitionTarget) Set(val *GraphDefinitionTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphDefinitionTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphDefinitionTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphDefinitionTarget(val *GraphDefinitionTarget) *NullableGraphDefinitionTarget {
	return &NullableGraphDefinitionTarget{value: val, isSet: true}
}

func (v NullableGraphDefinitionTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphDefinitionTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


