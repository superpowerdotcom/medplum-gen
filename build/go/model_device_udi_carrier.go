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

// checks if the DeviceUdiCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceUdiCarrier{}

// DeviceUdiCarrier A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
type DeviceUdiCarrier struct {
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
	// A stream of bytes
	CarrierAIDC *string `json:"carrierAIDC,omitempty"`
	// A sequence of Unicode characters
	CarrierHRF *string `json:"carrierHRF,omitempty"`
	// A coded entry to indicate how the data was entered.
	EntryType *string `json:"entryType,omitempty"`
}

// NewDeviceUdiCarrier instantiates a new DeviceUdiCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUdiCarrier() *DeviceUdiCarrier {
	this := DeviceUdiCarrier{}
	return &this
}

// NewDeviceUdiCarrierWithDefaults instantiates a new DeviceUdiCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUdiCarrierWithDefaults() *DeviceUdiCarrier {
	this := DeviceUdiCarrier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceUdiCarrier) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DeviceUdiCarrier) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DeviceUdiCarrier) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDeviceIdentifier returns the DeviceIdentifier field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetDeviceIdentifier() string {
	if o == nil || IsNil(o.DeviceIdentifier) {
		var ret string
		return ret
	}
	return *o.DeviceIdentifier
}

// GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetDeviceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceIdentifier) {
		return nil, false
	}
	return o.DeviceIdentifier, true
}

// HasDeviceIdentifier returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasDeviceIdentifier() bool {
	if o != nil && !IsNil(o.DeviceIdentifier) {
		return true
	}

	return false
}

// SetDeviceIdentifier gets a reference to the given string and assigns it to the DeviceIdentifier field.
func (o *DeviceUdiCarrier) SetDeviceIdentifier(v string) {
	o.DeviceIdentifier = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *DeviceUdiCarrier) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetJurisdiction() string {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret string
		return ret
	}
	return *o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetJurisdictionOk() (*string, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given string and assigns it to the Jurisdiction field.
func (o *DeviceUdiCarrier) SetJurisdiction(v string) {
	o.Jurisdiction = &v
}

// GetCarrierAIDC returns the CarrierAIDC field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetCarrierAIDC() string {
	if o == nil || IsNil(o.CarrierAIDC) {
		var ret string
		return ret
	}
	return *o.CarrierAIDC
}

// GetCarrierAIDCOk returns a tuple with the CarrierAIDC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetCarrierAIDCOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAIDC) {
		return nil, false
	}
	return o.CarrierAIDC, true
}

// HasCarrierAIDC returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasCarrierAIDC() bool {
	if o != nil && !IsNil(o.CarrierAIDC) {
		return true
	}

	return false
}

// SetCarrierAIDC gets a reference to the given string and assigns it to the CarrierAIDC field.
func (o *DeviceUdiCarrier) SetCarrierAIDC(v string) {
	o.CarrierAIDC = &v
}

// GetCarrierHRF returns the CarrierHRF field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetCarrierHRF() string {
	if o == nil || IsNil(o.CarrierHRF) {
		var ret string
		return ret
	}
	return *o.CarrierHRF
}

// GetCarrierHRFOk returns a tuple with the CarrierHRF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetCarrierHRFOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierHRF) {
		return nil, false
	}
	return o.CarrierHRF, true
}

// HasCarrierHRF returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasCarrierHRF() bool {
	if o != nil && !IsNil(o.CarrierHRF) {
		return true
	}

	return false
}

// SetCarrierHRF gets a reference to the given string and assigns it to the CarrierHRF field.
func (o *DeviceUdiCarrier) SetCarrierHRF(v string) {
	o.CarrierHRF = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *DeviceUdiCarrier) GetEntryType() string {
	if o == nil || IsNil(o.EntryType) {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUdiCarrier) GetEntryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *DeviceUdiCarrier) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *DeviceUdiCarrier) SetEntryType(v string) {
	o.EntryType = &v
}

func (o DeviceUdiCarrier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceUdiCarrier) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CarrierAIDC) {
		toSerialize["carrierAIDC"] = o.CarrierAIDC
	}
	if !IsNil(o.CarrierHRF) {
		toSerialize["carrierHRF"] = o.CarrierHRF
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	return toSerialize, nil
}

type NullableDeviceUdiCarrier struct {
	value *DeviceUdiCarrier
	isSet bool
}

func (v NullableDeviceUdiCarrier) Get() *DeviceUdiCarrier {
	return v.value
}

func (v *NullableDeviceUdiCarrier) Set(val *DeviceUdiCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUdiCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUdiCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUdiCarrier(val *DeviceUdiCarrier) *NullableDeviceUdiCarrier {
	return &NullableDeviceUdiCarrier{value: val, isSet: true}
}

func (v NullableDeviceUdiCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUdiCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

