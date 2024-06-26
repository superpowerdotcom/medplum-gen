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

// checks if the ChargeItemPerformer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeItemPerformer{}

// ChargeItemPerformer The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItemPerformer struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Describes the type of performance or participation(e.g. primary surgeon, anesthesiologiest, etc.).
	Function *CodeableConcept `json:"function,omitempty"`
	// The device, practitioner, etc. who performed or participated in the service.
	Actor Reference `json:"actor"`
}

type _ChargeItemPerformer ChargeItemPerformer

// NewChargeItemPerformer instantiates a new ChargeItemPerformer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeItemPerformer(actor Reference) *ChargeItemPerformer {
	this := ChargeItemPerformer{}
	this.Actor = actor
	return &this
}

// NewChargeItemPerformerWithDefaults instantiates a new ChargeItemPerformer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeItemPerformerWithDefaults() *ChargeItemPerformer {
	this := ChargeItemPerformer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeItemPerformer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItemPerformer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeItemPerformer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeItemPerformer) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ChargeItemPerformer) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItemPerformer) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ChargeItemPerformer) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ChargeItemPerformer) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ChargeItemPerformer) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItemPerformer) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ChargeItemPerformer) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ChargeItemPerformer) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *ChargeItemPerformer) GetFunction() CodeableConcept {
	if o == nil || IsNil(o.Function) {
		var ret CodeableConcept
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItemPerformer) GetFunctionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *ChargeItemPerformer) HasFunction() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given CodeableConcept and assigns it to the Function field.
func (o *ChargeItemPerformer) SetFunction(v CodeableConcept) {
	o.Function = &v
}

// GetActor returns the Actor field value
func (o *ChargeItemPerformer) GetActor() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *ChargeItemPerformer) GetActorOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *ChargeItemPerformer) SetActor(v Reference) {
	o.Actor = v
}

func (o ChargeItemPerformer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeItemPerformer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	toSerialize["actor"] = o.Actor
	return toSerialize, nil
}

func (o *ChargeItemPerformer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actor",
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

	varChargeItemPerformer := _ChargeItemPerformer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargeItemPerformer)

	if err != nil {
		return err
	}

	*o = ChargeItemPerformer(varChargeItemPerformer)

	return err
}

type NullableChargeItemPerformer struct {
	value *ChargeItemPerformer
	isSet bool
}

func (v NullableChargeItemPerformer) Get() *ChargeItemPerformer {
	return v.value
}

func (v *NullableChargeItemPerformer) Set(val *ChargeItemPerformer) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeItemPerformer) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeItemPerformer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeItemPerformer(val *ChargeItemPerformer) *NullableChargeItemPerformer {
	return &NullableChargeItemPerformer{value: val, isSet: true}
}

func (v NullableChargeItemPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeItemPerformer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


