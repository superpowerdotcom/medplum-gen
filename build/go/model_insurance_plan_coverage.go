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

// checks if the InsurancePlanCoverage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsurancePlanCoverage{}

// InsurancePlanCoverage Details of a Health Insurance product/plan provided by an organization.
type InsurancePlanCoverage struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of coverage  (Medical; Dental; Mental Health; Substance Abuse; Vision; Drug; Short Term; Long Term Care; Hospice; Home Health).
	Type CodeableConcept `json:"type"`
	// Reference to the network that providing the type of coverage.
	Network []Reference `json:"network,omitempty"`
	// Specific benefits under this type of coverage.
	Benefit []InsurancePlanBenefit `json:"benefit"`
}

type _InsurancePlanCoverage InsurancePlanCoverage

// NewInsurancePlanCoverage instantiates a new InsurancePlanCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsurancePlanCoverage(type_ CodeableConcept, benefit []InsurancePlanBenefit) *InsurancePlanCoverage {
	this := InsurancePlanCoverage{}
	this.Type = type_
	this.Benefit = benefit
	return &this
}

// NewInsurancePlanCoverageWithDefaults instantiates a new InsurancePlanCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsurancePlanCoverageWithDefaults() *InsurancePlanCoverage {
	this := InsurancePlanCoverage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InsurancePlanCoverage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InsurancePlanCoverage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InsurancePlanCoverage) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *InsurancePlanCoverage) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *InsurancePlanCoverage) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *InsurancePlanCoverage) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *InsurancePlanCoverage) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *InsurancePlanCoverage) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *InsurancePlanCoverage) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetType returns the Type field value
func (o *InsurancePlanCoverage) GetType() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetTypeOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InsurancePlanCoverage) SetType(v CodeableConcept) {
	o.Type = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InsurancePlanCoverage) GetNetwork() []Reference {
	if o == nil || IsNil(o.Network) {
		var ret []Reference
		return ret
	}
	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetNetworkOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InsurancePlanCoverage) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given []Reference and assigns it to the Network field.
func (o *InsurancePlanCoverage) SetNetwork(v []Reference) {
	o.Network = v
}

// GetBenefit returns the Benefit field value
func (o *InsurancePlanCoverage) GetBenefit() []InsurancePlanBenefit {
	if o == nil {
		var ret []InsurancePlanBenefit
		return ret
	}

	return o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value
// and a boolean to check if the value has been set.
func (o *InsurancePlanCoverage) GetBenefitOk() ([]InsurancePlanBenefit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Benefit, true
}

// SetBenefit sets field value
func (o *InsurancePlanCoverage) SetBenefit(v []InsurancePlanBenefit) {
	o.Benefit = v
}

func (o InsurancePlanCoverage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsurancePlanCoverage) ToMap() (map[string]interface{}, error) {
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
	toSerialize["type"] = o.Type
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	toSerialize["benefit"] = o.Benefit
	return toSerialize, nil
}

func (o *InsurancePlanCoverage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"benefit",
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

	varInsurancePlanCoverage := _InsurancePlanCoverage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInsurancePlanCoverage)

	if err != nil {
		return err
	}

	*o = InsurancePlanCoverage(varInsurancePlanCoverage)

	return err
}

type NullableInsurancePlanCoverage struct {
	value *InsurancePlanCoverage
	isSet bool
}

func (v NullableInsurancePlanCoverage) Get() *InsurancePlanCoverage {
	return v.value
}

func (v *NullableInsurancePlanCoverage) Set(val *InsurancePlanCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableInsurancePlanCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableInsurancePlanCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsurancePlanCoverage(val *InsurancePlanCoverage) *NullableInsurancePlanCoverage {
	return &NullableInsurancePlanCoverage{value: val, isSet: true}
}

func (v NullableInsurancePlanCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsurancePlanCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


