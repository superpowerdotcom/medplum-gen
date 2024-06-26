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

// checks if the DeviceDefinitionUdiDeviceIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceDefinitionUdiDeviceIdentifier{}

// DeviceDefinitionUdiDeviceIdentifier The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceDefinitionUdiDeviceIdentifier struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	DeviceIdentifier *string `json:"deviceIdentifier,omitempty"`
	// String of characters used to identify a name or a resource
	Issuer *string `json:"issuer,omitempty"`
	// String of characters used to identify a name or a resource
	Jurisdiction *string `json:"jurisdiction,omitempty"`
}

// NewDeviceDefinitionUdiDeviceIdentifier instantiates a new DeviceDefinitionUdiDeviceIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceDefinitionUdiDeviceIdentifier() *DeviceDefinitionUdiDeviceIdentifier {
	this := DeviceDefinitionUdiDeviceIdentifier{}
	return &this
}

// NewDeviceDefinitionUdiDeviceIdentifierWithDefaults instantiates a new DeviceDefinitionUdiDeviceIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceDefinitionUdiDeviceIdentifierWithDefaults() *DeviceDefinitionUdiDeviceIdentifier {
	this := DeviceDefinitionUdiDeviceIdentifier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDeviceIdentifier returns the DeviceIdentifier field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetDeviceIdentifier() string {
	if o == nil || IsNil(o.DeviceIdentifier) {
		var ret string
		return ret
	}
	return *o.DeviceIdentifier
}

// GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetDeviceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceIdentifier) {
		return nil, false
	}
	return o.DeviceIdentifier, true
}

// HasDeviceIdentifier returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasDeviceIdentifier() bool {
	if o != nil && !IsNil(o.DeviceIdentifier) {
		return true
	}

	return false
}

// SetDeviceIdentifier gets a reference to the given string and assigns it to the DeviceIdentifier field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetDeviceIdentifier(v string) {
	o.DeviceIdentifier = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetJurisdiction() string {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret string
		return ret
	}
	return *o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) GetJurisdictionOk() (*string, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *DeviceDefinitionUdiDeviceIdentifier) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given string and assigns it to the Jurisdiction field.
func (o *DeviceDefinitionUdiDeviceIdentifier) SetJurisdiction(v string) {
	o.Jurisdiction = &v
}

func (o DeviceDefinitionUdiDeviceIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceDefinitionUdiDeviceIdentifier) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceIdentifier) {
		toSerialize["deviceIdentifier"] = o.DeviceIdentifier
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	return toSerialize, nil
}

type NullableDeviceDefinitionUdiDeviceIdentifier struct {
	value *DeviceDefinitionUdiDeviceIdentifier
	isSet bool
}

func (v NullableDeviceDefinitionUdiDeviceIdentifier) Get() *DeviceDefinitionUdiDeviceIdentifier {
	return v.value
}

func (v *NullableDeviceDefinitionUdiDeviceIdentifier) Set(val *DeviceDefinitionUdiDeviceIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceDefinitionUdiDeviceIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceDefinitionUdiDeviceIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceDefinitionUdiDeviceIdentifier(val *DeviceDefinitionUdiDeviceIdentifier) *NullableDeviceDefinitionUdiDeviceIdentifier {
	return &NullableDeviceDefinitionUdiDeviceIdentifier{value: val, isSet: true}
}

func (v NullableDeviceDefinitionUdiDeviceIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceDefinitionUdiDeviceIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


