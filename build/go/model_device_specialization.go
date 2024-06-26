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

// checks if the DeviceSpecialization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSpecialization{}

// DeviceSpecialization A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
type DeviceSpecialization struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The standard that is used to operate and communicate.
	SystemType CodeableConcept `json:"systemType"`
	// A sequence of Unicode characters
	Version *string `json:"version,omitempty"`
}

type _DeviceSpecialization DeviceSpecialization

// NewDeviceSpecialization instantiates a new DeviceSpecialization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSpecialization(systemType CodeableConcept) *DeviceSpecialization {
	this := DeviceSpecialization{}
	this.SystemType = systemType
	return &this
}

// NewDeviceSpecializationWithDefaults instantiates a new DeviceSpecialization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSpecializationWithDefaults() *DeviceSpecialization {
	this := DeviceSpecialization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceSpecialization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSpecialization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceSpecialization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceSpecialization) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DeviceSpecialization) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSpecialization) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DeviceSpecialization) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DeviceSpecialization) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DeviceSpecialization) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSpecialization) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DeviceSpecialization) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DeviceSpecialization) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSystemType returns the SystemType field value
func (o *DeviceSpecialization) GetSystemType() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value
// and a boolean to check if the value has been set.
func (o *DeviceSpecialization) GetSystemTypeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemType, true
}

// SetSystemType sets field value
func (o *DeviceSpecialization) SetSystemType(v CodeableConcept) {
	o.SystemType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceSpecialization) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSpecialization) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceSpecialization) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeviceSpecialization) SetVersion(v string) {
	o.Version = &v
}

func (o DeviceSpecialization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSpecialization) ToMap() (map[string]interface{}, error) {
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
	toSerialize["systemType"] = o.SystemType
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *DeviceSpecialization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"systemType",
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

	varDeviceSpecialization := _DeviceSpecialization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceSpecialization)

	if err != nil {
		return err
	}

	*o = DeviceSpecialization(varDeviceSpecialization)

	return err
}

type NullableDeviceSpecialization struct {
	value *DeviceSpecialization
	isSet bool
}

func (v NullableDeviceSpecialization) Get() *DeviceSpecialization {
	return v.value
}

func (v *NullableDeviceSpecialization) Set(val *DeviceSpecialization) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSpecialization) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSpecialization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSpecialization(val *DeviceSpecialization) *NullableDeviceSpecialization {
	return &NullableDeviceSpecialization{value: val, isSet: true}
}

func (v NullableDeviceSpecialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSpecialization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


