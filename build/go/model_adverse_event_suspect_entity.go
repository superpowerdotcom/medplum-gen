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

// checks if the AdverseEventSuspectEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdverseEventSuspectEntity{}

// AdverseEventSuspectEntity Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEventSuspectEntity struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the actual instance of what caused the adverse event.  May be a substance, medication, medication administration, medication statement or a device.
	Instance Reference `json:"instance"`
	// Information on the possible cause of the event.
	Causality []AdverseEventCausality `json:"causality,omitempty"`
}

type _AdverseEventSuspectEntity AdverseEventSuspectEntity

// NewAdverseEventSuspectEntity instantiates a new AdverseEventSuspectEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdverseEventSuspectEntity(instance Reference) *AdverseEventSuspectEntity {
	this := AdverseEventSuspectEntity{}
	this.Instance = instance
	return &this
}

// NewAdverseEventSuspectEntityWithDefaults instantiates a new AdverseEventSuspectEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdverseEventSuspectEntityWithDefaults() *AdverseEventSuspectEntity {
	this := AdverseEventSuspectEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdverseEventSuspectEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEventSuspectEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdverseEventSuspectEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdverseEventSuspectEntity) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AdverseEventSuspectEntity) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEventSuspectEntity) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AdverseEventSuspectEntity) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *AdverseEventSuspectEntity) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *AdverseEventSuspectEntity) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEventSuspectEntity) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *AdverseEventSuspectEntity) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *AdverseEventSuspectEntity) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetInstance returns the Instance field value
func (o *AdverseEventSuspectEntity) GetInstance() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *AdverseEventSuspectEntity) GetInstanceOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *AdverseEventSuspectEntity) SetInstance(v Reference) {
	o.Instance = v
}

// GetCausality returns the Causality field value if set, zero value otherwise.
func (o *AdverseEventSuspectEntity) GetCausality() []AdverseEventCausality {
	if o == nil || IsNil(o.Causality) {
		var ret []AdverseEventCausality
		return ret
	}
	return o.Causality
}

// GetCausalityOk returns a tuple with the Causality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdverseEventSuspectEntity) GetCausalityOk() ([]AdverseEventCausality, bool) {
	if o == nil || IsNil(o.Causality) {
		return nil, false
	}
	return o.Causality, true
}

// HasCausality returns a boolean if a field has been set.
func (o *AdverseEventSuspectEntity) HasCausality() bool {
	if o != nil && !IsNil(o.Causality) {
		return true
	}

	return false
}

// SetCausality gets a reference to the given []AdverseEventCausality and assigns it to the Causality field.
func (o *AdverseEventSuspectEntity) SetCausality(v []AdverseEventCausality) {
	o.Causality = v
}

func (o AdverseEventSuspectEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdverseEventSuspectEntity) ToMap() (map[string]interface{}, error) {
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
	toSerialize["instance"] = o.Instance
	if !IsNil(o.Causality) {
		toSerialize["causality"] = o.Causality
	}
	return toSerialize, nil
}

func (o *AdverseEventSuspectEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instance",
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

	varAdverseEventSuspectEntity := _AdverseEventSuspectEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdverseEventSuspectEntity)

	if err != nil {
		return err
	}

	*o = AdverseEventSuspectEntity(varAdverseEventSuspectEntity)

	return err
}

type NullableAdverseEventSuspectEntity struct {
	value *AdverseEventSuspectEntity
	isSet bool
}

func (v NullableAdverseEventSuspectEntity) Get() *AdverseEventSuspectEntity {
	return v.value
}

func (v *NullableAdverseEventSuspectEntity) Set(val *AdverseEventSuspectEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAdverseEventSuspectEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAdverseEventSuspectEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdverseEventSuspectEntity(val *AdverseEventSuspectEntity) *NullableAdverseEventSuspectEntity {
	return &NullableAdverseEventSuspectEntity{value: val, isSet: true}
}

func (v NullableAdverseEventSuspectEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdverseEventSuspectEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


