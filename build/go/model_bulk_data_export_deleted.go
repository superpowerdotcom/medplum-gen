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
	"fmt"
)

// checks if the BulkDataExportDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkDataExportDeleted{}

// BulkDataExportDeleted An array of deleted file items following the same structure as the output array.
type BulkDataExportDeleted struct {
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Type string `json:"type"`
	// String of characters used to identify a name or a resource
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _BulkDataExportDeleted BulkDataExportDeleted

// NewBulkDataExportDeleted instantiates a new BulkDataExportDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkDataExportDeleted(type_ string, url string) *BulkDataExportDeleted {
	this := BulkDataExportDeleted{}
	this.Type = type_
	this.Url = url
	return &this
}

// NewBulkDataExportDeletedWithDefaults instantiates a new BulkDataExportDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkDataExportDeletedWithDefaults() *BulkDataExportDeleted {
	this := BulkDataExportDeleted{}
	return &this
}

// GetType returns the Type field value
func (o *BulkDataExportDeleted) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkDataExportDeleted) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BulkDataExportDeleted) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *BulkDataExportDeleted) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BulkDataExportDeleted) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BulkDataExportDeleted) SetUrl(v string) {
	o.Url = v
}

func (o BulkDataExportDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkDataExportDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkDataExportDeleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"url",
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

	varBulkDataExportDeleted := _BulkDataExportDeleted{}

	err = json.Unmarshal(data, &varBulkDataExportDeleted)

	if err != nil {
		return err
	}

	*o = BulkDataExportDeleted(varBulkDataExportDeleted)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkDataExportDeleted struct {
	value *BulkDataExportDeleted
	isSet bool
}

func (v NullableBulkDataExportDeleted) Get() *BulkDataExportDeleted {
	return v.value
}

func (v *NullableBulkDataExportDeleted) Set(val *BulkDataExportDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkDataExportDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkDataExportDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkDataExportDeleted(val *BulkDataExportDeleted) *NullableBulkDataExportDeleted {
	return &NullableBulkDataExportDeleted{value: val, isSet: true}
}

func (v NullableBulkDataExportDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkDataExportDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

