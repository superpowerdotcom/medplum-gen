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

// checks if the ImplementationGuideManifest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImplementationGuideManifest{}

// ImplementationGuideManifest A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuideManifest struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A URI that is a literal reference
	Rendering *string `json:"rendering,omitempty"`
	// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
	Resource []ImplementationGuideResource1 `json:"resource"`
	// Information about a page within the IG.
	Page []ImplementationGuidePage1 `json:"page,omitempty"`
	// Indicates a relative path to an image that exists within the IG.
	Image []string `json:"image,omitempty"`
	// Indicates the relative path of an additional non-page, non-image file that is part of the IG - e.g. zip, jar and similar files that could be the target of a hyperlink in a derived IG.
	Other []string `json:"other,omitempty"`
}

type _ImplementationGuideManifest ImplementationGuideManifest

// NewImplementationGuideManifest instantiates a new ImplementationGuideManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImplementationGuideManifest(resource []ImplementationGuideResource1) *ImplementationGuideManifest {
	this := ImplementationGuideManifest{}
	this.Resource = resource
	return &this
}

// NewImplementationGuideManifestWithDefaults instantiates a new ImplementationGuideManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImplementationGuideManifestWithDefaults() *ImplementationGuideManifest {
	this := ImplementationGuideManifest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImplementationGuideManifest) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ImplementationGuideManifest) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ImplementationGuideManifest) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetRendering returns the Rendering field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetRendering() string {
	if o == nil || IsNil(o.Rendering) {
		var ret string
		return ret
	}
	return *o.Rendering
}

// GetRenderingOk returns a tuple with the Rendering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetRenderingOk() (*string, bool) {
	if o == nil || IsNil(o.Rendering) {
		return nil, false
	}
	return o.Rendering, true
}

// HasRendering returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasRendering() bool {
	if o != nil && !IsNil(o.Rendering) {
		return true
	}

	return false
}

// SetRendering gets a reference to the given string and assigns it to the Rendering field.
func (o *ImplementationGuideManifest) SetRendering(v string) {
	o.Rendering = &v
}

// GetResource returns the Resource field value
func (o *ImplementationGuideManifest) GetResource() []ImplementationGuideResource1 {
	if o == nil {
		var ret []ImplementationGuideResource1
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetResourceOk() ([]ImplementationGuideResource1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource, true
}

// SetResource sets field value
func (o *ImplementationGuideManifest) SetResource(v []ImplementationGuideResource1) {
	o.Resource = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetPage() []ImplementationGuidePage1 {
	if o == nil || IsNil(o.Page) {
		var ret []ImplementationGuidePage1
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetPageOk() ([]ImplementationGuidePage1, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given []ImplementationGuidePage1 and assigns it to the Page field.
func (o *ImplementationGuideManifest) SetPage(v []ImplementationGuidePage1) {
	o.Page = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetImage() []string {
	if o == nil || IsNil(o.Image) {
		var ret []string
		return ret
	}
	return o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetImageOk() ([]string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given []string and assigns it to the Image field.
func (o *ImplementationGuideManifest) SetImage(v []string) {
	o.Image = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *ImplementationGuideManifest) GetOther() []string {
	if o == nil || IsNil(o.Other) {
		var ret []string
		return ret
	}
	return o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplementationGuideManifest) GetOtherOk() ([]string, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *ImplementationGuideManifest) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given []string and assigns it to the Other field.
func (o *ImplementationGuideManifest) SetOther(v []string) {
	o.Other = v
}

func (o ImplementationGuideManifest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImplementationGuideManifest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Rendering) {
		toSerialize["rendering"] = o.Rendering
	}
	toSerialize["resource"] = o.Resource
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	return toSerialize, nil
}

func (o *ImplementationGuideManifest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource",
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

	varImplementationGuideManifest := _ImplementationGuideManifest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImplementationGuideManifest)

	if err != nil {
		return err
	}

	*o = ImplementationGuideManifest(varImplementationGuideManifest)

	return err
}

type NullableImplementationGuideManifest struct {
	value *ImplementationGuideManifest
	isSet bool
}

func (v NullableImplementationGuideManifest) Get() *ImplementationGuideManifest {
	return v.value
}

func (v *NullableImplementationGuideManifest) Set(val *ImplementationGuideManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableImplementationGuideManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableImplementationGuideManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImplementationGuideManifest(val *ImplementationGuideManifest) *NullableImplementationGuideManifest {
	return &NullableImplementationGuideManifest{value: val, isSet: true}
}

func (v NullableImplementationGuideManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImplementationGuideManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


