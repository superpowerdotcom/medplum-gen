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

// checks if the VerificationResultValidator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationResultValidator{}

// VerificationResultValidator Describes validation requirements, source(s), status and dates for one or more elements.
type VerificationResultValidator struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the organization validating information.
	Organization Reference `json:"organization"`
	// A sequence of Unicode characters
	IdentityCertificate *string `json:"identityCertificate,omitempty"`
	// Signed assertion by the validator that they have validated the information.
	AttestationSignature *Signature `json:"attestationSignature,omitempty"`
}

type _VerificationResultValidator VerificationResultValidator

// NewVerificationResultValidator instantiates a new VerificationResultValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationResultValidator(organization Reference) *VerificationResultValidator {
	this := VerificationResultValidator{}
	this.Organization = organization
	return &this
}

// NewVerificationResultValidatorWithDefaults instantiates a new VerificationResultValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationResultValidatorWithDefaults() *VerificationResultValidator {
	this := VerificationResultValidator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VerificationResultValidator) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VerificationResultValidator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VerificationResultValidator) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *VerificationResultValidator) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *VerificationResultValidator) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *VerificationResultValidator) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *VerificationResultValidator) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *VerificationResultValidator) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *VerificationResultValidator) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetOrganization returns the Organization field value
func (o *VerificationResultValidator) GetOrganization() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetOrganizationOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *VerificationResultValidator) SetOrganization(v Reference) {
	o.Organization = v
}

// GetIdentityCertificate returns the IdentityCertificate field value if set, zero value otherwise.
func (o *VerificationResultValidator) GetIdentityCertificate() string {
	if o == nil || IsNil(o.IdentityCertificate) {
		var ret string
		return ret
	}
	return *o.IdentityCertificate
}

// GetIdentityCertificateOk returns a tuple with the IdentityCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetIdentityCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityCertificate) {
		return nil, false
	}
	return o.IdentityCertificate, true
}

// HasIdentityCertificate returns a boolean if a field has been set.
func (o *VerificationResultValidator) HasIdentityCertificate() bool {
	if o != nil && !IsNil(o.IdentityCertificate) {
		return true
	}

	return false
}

// SetIdentityCertificate gets a reference to the given string and assigns it to the IdentityCertificate field.
func (o *VerificationResultValidator) SetIdentityCertificate(v string) {
	o.IdentityCertificate = &v
}

// GetAttestationSignature returns the AttestationSignature field value if set, zero value otherwise.
func (o *VerificationResultValidator) GetAttestationSignature() Signature {
	if o == nil || IsNil(o.AttestationSignature) {
		var ret Signature
		return ret
	}
	return *o.AttestationSignature
}

// GetAttestationSignatureOk returns a tuple with the AttestationSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResultValidator) GetAttestationSignatureOk() (*Signature, bool) {
	if o == nil || IsNil(o.AttestationSignature) {
		return nil, false
	}
	return o.AttestationSignature, true
}

// HasAttestationSignature returns a boolean if a field has been set.
func (o *VerificationResultValidator) HasAttestationSignature() bool {
	if o != nil && !IsNil(o.AttestationSignature) {
		return true
	}

	return false
}

// SetAttestationSignature gets a reference to the given Signature and assigns it to the AttestationSignature field.
func (o *VerificationResultValidator) SetAttestationSignature(v Signature) {
	o.AttestationSignature = &v
}

func (o VerificationResultValidator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationResultValidator) ToMap() (map[string]interface{}, error) {
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
	toSerialize["organization"] = o.Organization
	if !IsNil(o.IdentityCertificate) {
		toSerialize["identityCertificate"] = o.IdentityCertificate
	}
	if !IsNil(o.AttestationSignature) {
		toSerialize["attestationSignature"] = o.AttestationSignature
	}
	return toSerialize, nil
}

func (o *VerificationResultValidator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organization",
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

	varVerificationResultValidator := _VerificationResultValidator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerificationResultValidator)

	if err != nil {
		return err
	}

	*o = VerificationResultValidator(varVerificationResultValidator)

	return err
}

type NullableVerificationResultValidator struct {
	value *VerificationResultValidator
	isSet bool
}

func (v NullableVerificationResultValidator) Get() *VerificationResultValidator {
	return v.value
}

func (v *NullableVerificationResultValidator) Set(val *VerificationResultValidator) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationResultValidator) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationResultValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationResultValidator(val *VerificationResultValidator) *NullableVerificationResultValidator {
	return &NullableVerificationResultValidator{value: val, isSet: true}
}

func (v NullableVerificationResultValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationResultValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


