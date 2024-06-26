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

// checks if the NutritionOrderAdministration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NutritionOrderAdministration{}

// NutritionOrderAdministration A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrderAdministration struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The time period and frequency at which the enteral formula should be delivered to the patient.
	Schedule *Timing `json:"schedule,omitempty"`
	// The volume of formula to provide to the patient per the specified administration schedule.
	Quantity *Quantity `json:"quantity,omitempty"`
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	RateQuantity *Quantity `json:"rateQuantity,omitempty"`
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	RateRatio *Ratio `json:"rateRatio,omitempty"`
}

// NewNutritionOrderAdministration instantiates a new NutritionOrderAdministration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNutritionOrderAdministration() *NutritionOrderAdministration {
	this := NutritionOrderAdministration{}
	return &this
}

// NewNutritionOrderAdministrationWithDefaults instantiates a new NutritionOrderAdministration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNutritionOrderAdministrationWithDefaults() *NutritionOrderAdministration {
	this := NutritionOrderAdministration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NutritionOrderAdministration) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *NutritionOrderAdministration) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *NutritionOrderAdministration) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetSchedule() Timing {
	if o == nil || IsNil(o.Schedule) {
		var ret Timing
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetScheduleOk() (*Timing, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Timing and assigns it to the Schedule field.
func (o *NutritionOrderAdministration) SetSchedule(v Timing) {
	o.Schedule = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetQuantity() Quantity {
	if o == nil || IsNil(o.Quantity) {
		var ret Quantity
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given Quantity and assigns it to the Quantity field.
func (o *NutritionOrderAdministration) SetQuantity(v Quantity) {
	o.Quantity = &v
}

// GetRateQuantity returns the RateQuantity field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetRateQuantity() Quantity {
	if o == nil || IsNil(o.RateQuantity) {
		var ret Quantity
		return ret
	}
	return *o.RateQuantity
}

// GetRateQuantityOk returns a tuple with the RateQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetRateQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.RateQuantity) {
		return nil, false
	}
	return o.RateQuantity, true
}

// HasRateQuantity returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasRateQuantity() bool {
	if o != nil && !IsNil(o.RateQuantity) {
		return true
	}

	return false
}

// SetRateQuantity gets a reference to the given Quantity and assigns it to the RateQuantity field.
func (o *NutritionOrderAdministration) SetRateQuantity(v Quantity) {
	o.RateQuantity = &v
}

// GetRateRatio returns the RateRatio field value if set, zero value otherwise.
func (o *NutritionOrderAdministration) GetRateRatio() Ratio {
	if o == nil || IsNil(o.RateRatio) {
		var ret Ratio
		return ret
	}
	return *o.RateRatio
}

// GetRateRatioOk returns a tuple with the RateRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NutritionOrderAdministration) GetRateRatioOk() (*Ratio, bool) {
	if o == nil || IsNil(o.RateRatio) {
		return nil, false
	}
	return o.RateRatio, true
}

// HasRateRatio returns a boolean if a field has been set.
func (o *NutritionOrderAdministration) HasRateRatio() bool {
	if o != nil && !IsNil(o.RateRatio) {
		return true
	}

	return false
}

// SetRateRatio gets a reference to the given Ratio and assigns it to the RateRatio field.
func (o *NutritionOrderAdministration) SetRateRatio(v Ratio) {
	o.RateRatio = &v
}

func (o NutritionOrderAdministration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NutritionOrderAdministration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.RateQuantity) {
		toSerialize["rateQuantity"] = o.RateQuantity
	}
	if !IsNil(o.RateRatio) {
		toSerialize["rateRatio"] = o.RateRatio
	}
	return toSerialize, nil
}

type NullableNutritionOrderAdministration struct {
	value *NutritionOrderAdministration
	isSet bool
}

func (v NullableNutritionOrderAdministration) Get() *NutritionOrderAdministration {
	return v.value
}

func (v *NullableNutritionOrderAdministration) Set(val *NutritionOrderAdministration) {
	v.value = val
	v.isSet = true
}

func (v NullableNutritionOrderAdministration) IsSet() bool {
	return v.isSet
}

func (v *NullableNutritionOrderAdministration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNutritionOrderAdministration(val *NutritionOrderAdministration) *NullableNutritionOrderAdministration {
	return &NullableNutritionOrderAdministration{value: val, isSet: true}
}

func (v NullableNutritionOrderAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNutritionOrderAdministration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


