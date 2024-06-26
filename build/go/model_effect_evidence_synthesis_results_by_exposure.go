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

// checks if the EffectEvidenceSynthesisResultsByExposure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectEvidenceSynthesisResultsByExposure{}

// EffectEvidenceSynthesisResultsByExposure The EffectEvidenceSynthesis resource describes the difference in an outcome between exposures states in a population where the effect estimate is derived from a combination of research studies.
type EffectEvidenceSynthesisResultsByExposure struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Whether these results are for the exposure state or alternative exposure state.
	ExposureState *string `json:"exposureState,omitempty"`
	// Used to define variant exposure states such as low-risk state.
	VariantState *CodeableConcept `json:"variantState,omitempty"`
	// Reference to a RiskEvidenceSynthesis resource.
	RiskEvidenceSynthesis Reference `json:"riskEvidenceSynthesis"`
}

type _EffectEvidenceSynthesisResultsByExposure EffectEvidenceSynthesisResultsByExposure

// NewEffectEvidenceSynthesisResultsByExposure instantiates a new EffectEvidenceSynthesisResultsByExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectEvidenceSynthesisResultsByExposure(riskEvidenceSynthesis Reference) *EffectEvidenceSynthesisResultsByExposure {
	this := EffectEvidenceSynthesisResultsByExposure{}
	this.RiskEvidenceSynthesis = riskEvidenceSynthesis
	return &this
}

// NewEffectEvidenceSynthesisResultsByExposureWithDefaults instantiates a new EffectEvidenceSynthesisResultsByExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectEvidenceSynthesisResultsByExposureWithDefaults() *EffectEvidenceSynthesisResultsByExposure {
	this := EffectEvidenceSynthesisResultsByExposure{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetDescription(v string) {
	o.Description = &v
}

// GetExposureState returns the ExposureState field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetExposureState() string {
	if o == nil || IsNil(o.ExposureState) {
		var ret string
		return ret
	}
	return *o.ExposureState
}

// GetExposureStateOk returns a tuple with the ExposureState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetExposureStateOk() (*string, bool) {
	if o == nil || IsNil(o.ExposureState) {
		return nil, false
	}
	return o.ExposureState, true
}

// HasExposureState returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasExposureState() bool {
	if o != nil && !IsNil(o.ExposureState) {
		return true
	}

	return false
}

// SetExposureState gets a reference to the given string and assigns it to the ExposureState field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetExposureState(v string) {
	o.ExposureState = &v
}

// GetVariantState returns the VariantState field value if set, zero value otherwise.
func (o *EffectEvidenceSynthesisResultsByExposure) GetVariantState() CodeableConcept {
	if o == nil || IsNil(o.VariantState) {
		var ret CodeableConcept
		return ret
	}
	return *o.VariantState
}

// GetVariantStateOk returns a tuple with the VariantState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetVariantStateOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.VariantState) {
		return nil, false
	}
	return o.VariantState, true
}

// HasVariantState returns a boolean if a field has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) HasVariantState() bool {
	if o != nil && !IsNil(o.VariantState) {
		return true
	}

	return false
}

// SetVariantState gets a reference to the given CodeableConcept and assigns it to the VariantState field.
func (o *EffectEvidenceSynthesisResultsByExposure) SetVariantState(v CodeableConcept) {
	o.VariantState = &v
}

// GetRiskEvidenceSynthesis returns the RiskEvidenceSynthesis field value
func (o *EffectEvidenceSynthesisResultsByExposure) GetRiskEvidenceSynthesis() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.RiskEvidenceSynthesis
}

// GetRiskEvidenceSynthesisOk returns a tuple with the RiskEvidenceSynthesis field value
// and a boolean to check if the value has been set.
func (o *EffectEvidenceSynthesisResultsByExposure) GetRiskEvidenceSynthesisOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskEvidenceSynthesis, true
}

// SetRiskEvidenceSynthesis sets field value
func (o *EffectEvidenceSynthesisResultsByExposure) SetRiskEvidenceSynthesis(v Reference) {
	o.RiskEvidenceSynthesis = v
}

func (o EffectEvidenceSynthesisResultsByExposure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectEvidenceSynthesisResultsByExposure) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExposureState) {
		toSerialize["exposureState"] = o.ExposureState
	}
	if !IsNil(o.VariantState) {
		toSerialize["variantState"] = o.VariantState
	}
	toSerialize["riskEvidenceSynthesis"] = o.RiskEvidenceSynthesis
	return toSerialize, nil
}

func (o *EffectEvidenceSynthesisResultsByExposure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"riskEvidenceSynthesis",
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

	varEffectEvidenceSynthesisResultsByExposure := _EffectEvidenceSynthesisResultsByExposure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEffectEvidenceSynthesisResultsByExposure)

	if err != nil {
		return err
	}

	*o = EffectEvidenceSynthesisResultsByExposure(varEffectEvidenceSynthesisResultsByExposure)

	return err
}

type NullableEffectEvidenceSynthesisResultsByExposure struct {
	value *EffectEvidenceSynthesisResultsByExposure
	isSet bool
}

func (v NullableEffectEvidenceSynthesisResultsByExposure) Get() *EffectEvidenceSynthesisResultsByExposure {
	return v.value
}

func (v *NullableEffectEvidenceSynthesisResultsByExposure) Set(val *EffectEvidenceSynthesisResultsByExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectEvidenceSynthesisResultsByExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectEvidenceSynthesisResultsByExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectEvidenceSynthesisResultsByExposure(val *EffectEvidenceSynthesisResultsByExposure) *NullableEffectEvidenceSynthesisResultsByExposure {
	return &NullableEffectEvidenceSynthesisResultsByExposure{value: val, isSet: true}
}

func (v NullableEffectEvidenceSynthesisResultsByExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectEvidenceSynthesisResultsByExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

