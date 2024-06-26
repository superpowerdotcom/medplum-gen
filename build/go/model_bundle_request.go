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

// checks if the BundleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleRequest{}

// BundleRequest A container for a collection of resources.
type BundleRequest struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// In a transaction or batch, this is the HTTP action to be executed for this entry. In a history bundle, this indicates the HTTP action that occurred.
	Method *string `json:"method,omitempty"`
	// String of characters used to identify a name or a resource
	Url *string `json:"url,omitempty"`
	// A sequence of Unicode characters
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`
	// An instant in time - known at least to the second
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`
	// A sequence of Unicode characters
	IfMatch *string `json:"ifMatch,omitempty"`
	// A sequence of Unicode characters
	IfNoneExist *string `json:"ifNoneExist,omitempty"`
}

// NewBundleRequest instantiates a new BundleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleRequest() *BundleRequest {
	this := BundleRequest{}
	return &this
}

// NewBundleRequestWithDefaults instantiates a new BundleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleRequestWithDefaults() *BundleRequest {
	this := BundleRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BundleRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BundleRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BundleRequest) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *BundleRequest) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *BundleRequest) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *BundleRequest) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *BundleRequest) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *BundleRequest) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *BundleRequest) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *BundleRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *BundleRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *BundleRequest) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BundleRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BundleRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BundleRequest) SetUrl(v string) {
	o.Url = &v
}

// GetIfNoneMatch returns the IfNoneMatch field value if set, zero value otherwise.
func (o *BundleRequest) GetIfNoneMatch() string {
	if o == nil || IsNil(o.IfNoneMatch) {
		var ret string
		return ret
	}
	return *o.IfNoneMatch
}

// GetIfNoneMatchOk returns a tuple with the IfNoneMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetIfNoneMatchOk() (*string, bool) {
	if o == nil || IsNil(o.IfNoneMatch) {
		return nil, false
	}
	return o.IfNoneMatch, true
}

// HasIfNoneMatch returns a boolean if a field has been set.
func (o *BundleRequest) HasIfNoneMatch() bool {
	if o != nil && !IsNil(o.IfNoneMatch) {
		return true
	}

	return false
}

// SetIfNoneMatch gets a reference to the given string and assigns it to the IfNoneMatch field.
func (o *BundleRequest) SetIfNoneMatch(v string) {
	o.IfNoneMatch = &v
}

// GetIfModifiedSince returns the IfModifiedSince field value if set, zero value otherwise.
func (o *BundleRequest) GetIfModifiedSince() string {
	if o == nil || IsNil(o.IfModifiedSince) {
		var ret string
		return ret
	}
	return *o.IfModifiedSince
}

// GetIfModifiedSinceOk returns a tuple with the IfModifiedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetIfModifiedSinceOk() (*string, bool) {
	if o == nil || IsNil(o.IfModifiedSince) {
		return nil, false
	}
	return o.IfModifiedSince, true
}

// HasIfModifiedSince returns a boolean if a field has been set.
func (o *BundleRequest) HasIfModifiedSince() bool {
	if o != nil && !IsNil(o.IfModifiedSince) {
		return true
	}

	return false
}

// SetIfModifiedSince gets a reference to the given string and assigns it to the IfModifiedSince field.
func (o *BundleRequest) SetIfModifiedSince(v string) {
	o.IfModifiedSince = &v
}

// GetIfMatch returns the IfMatch field value if set, zero value otherwise.
func (o *BundleRequest) GetIfMatch() string {
	if o == nil || IsNil(o.IfMatch) {
		var ret string
		return ret
	}
	return *o.IfMatch
}

// GetIfMatchOk returns a tuple with the IfMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetIfMatchOk() (*string, bool) {
	if o == nil || IsNil(o.IfMatch) {
		return nil, false
	}
	return o.IfMatch, true
}

// HasIfMatch returns a boolean if a field has been set.
func (o *BundleRequest) HasIfMatch() bool {
	if o != nil && !IsNil(o.IfMatch) {
		return true
	}

	return false
}

// SetIfMatch gets a reference to the given string and assigns it to the IfMatch field.
func (o *BundleRequest) SetIfMatch(v string) {
	o.IfMatch = &v
}

// GetIfNoneExist returns the IfNoneExist field value if set, zero value otherwise.
func (o *BundleRequest) GetIfNoneExist() string {
	if o == nil || IsNil(o.IfNoneExist) {
		var ret string
		return ret
	}
	return *o.IfNoneExist
}

// GetIfNoneExistOk returns a tuple with the IfNoneExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetIfNoneExistOk() (*string, bool) {
	if o == nil || IsNil(o.IfNoneExist) {
		return nil, false
	}
	return o.IfNoneExist, true
}

// HasIfNoneExist returns a boolean if a field has been set.
func (o *BundleRequest) HasIfNoneExist() bool {
	if o != nil && !IsNil(o.IfNoneExist) {
		return true
	}

	return false
}

// SetIfNoneExist gets a reference to the given string and assigns it to the IfNoneExist field.
func (o *BundleRequest) SetIfNoneExist(v string) {
	o.IfNoneExist = &v
}

func (o BundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.IfNoneMatch) {
		toSerialize["ifNoneMatch"] = o.IfNoneMatch
	}
	if !IsNil(o.IfModifiedSince) {
		toSerialize["ifModifiedSince"] = o.IfModifiedSince
	}
	if !IsNil(o.IfMatch) {
		toSerialize["ifMatch"] = o.IfMatch
	}
	if !IsNil(o.IfNoneExist) {
		toSerialize["ifNoneExist"] = o.IfNoneExist
	}
	return toSerialize, nil
}

type NullableBundleRequest struct {
	value *BundleRequest
	isSet bool
}

func (v NullableBundleRequest) Get() *BundleRequest {
	return v.value
}

func (v *NullableBundleRequest) Set(val *BundleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleRequest(val *BundleRequest) *NullableBundleRequest {
	return &NullableBundleRequest{value: val, isSet: true}
}

func (v NullableBundleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


