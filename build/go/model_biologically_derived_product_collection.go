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

// checks if the BiologicallyDerivedProductCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiologicallyDerivedProductCollection{}

// BiologicallyDerivedProductCollection A material substance originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity.
type BiologicallyDerivedProductCollection struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Healthcare professional who is performing the collection.
	Collector *Reference `json:"collector,omitempty"`
	// The patient or entity, such as a hospital or vendor in the case of a processed/manipulated/manufactured product, providing the product.
	Source *Reference `json:"source,omitempty"`
	// Time of product collection.
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`
	// Time of product collection.
	CollectedPeriod *Period `json:"collectedPeriod,omitempty"`
}

// NewBiologicallyDerivedProductCollection instantiates a new BiologicallyDerivedProductCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiologicallyDerivedProductCollection() *BiologicallyDerivedProductCollection {
	this := BiologicallyDerivedProductCollection{}
	return &this
}

// NewBiologicallyDerivedProductCollectionWithDefaults instantiates a new BiologicallyDerivedProductCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiologicallyDerivedProductCollectionWithDefaults() *BiologicallyDerivedProductCollection {
	this := BiologicallyDerivedProductCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BiologicallyDerivedProductCollection) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *BiologicallyDerivedProductCollection) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *BiologicallyDerivedProductCollection) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCollector returns the Collector field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetCollector() Reference {
	if o == nil || IsNil(o.Collector) {
		var ret Reference
		return ret
	}
	return *o.Collector
}

// GetCollectorOk returns a tuple with the Collector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetCollectorOk() (*Reference, bool) {
	if o == nil || IsNil(o.Collector) {
		return nil, false
	}
	return o.Collector, true
}

// HasCollector returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasCollector() bool {
	if o != nil && !IsNil(o.Collector) {
		return true
	}

	return false
}

// SetCollector gets a reference to the given Reference and assigns it to the Collector field.
func (o *BiologicallyDerivedProductCollection) SetCollector(v Reference) {
	o.Collector = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetSource() Reference {
	if o == nil || IsNil(o.Source) {
		var ret Reference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetSourceOk() (*Reference, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference and assigns it to the Source field.
func (o *BiologicallyDerivedProductCollection) SetSource(v Reference) {
	o.Source = &v
}

// GetCollectedDateTime returns the CollectedDateTime field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetCollectedDateTime() string {
	if o == nil || IsNil(o.CollectedDateTime) {
		var ret string
		return ret
	}
	return *o.CollectedDateTime
}

// GetCollectedDateTimeOk returns a tuple with the CollectedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetCollectedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectedDateTime) {
		return nil, false
	}
	return o.CollectedDateTime, true
}

// HasCollectedDateTime returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasCollectedDateTime() bool {
	if o != nil && !IsNil(o.CollectedDateTime) {
		return true
	}

	return false
}

// SetCollectedDateTime gets a reference to the given string and assigns it to the CollectedDateTime field.
func (o *BiologicallyDerivedProductCollection) SetCollectedDateTime(v string) {
	o.CollectedDateTime = &v
}

// GetCollectedPeriod returns the CollectedPeriod field value if set, zero value otherwise.
func (o *BiologicallyDerivedProductCollection) GetCollectedPeriod() Period {
	if o == nil || IsNil(o.CollectedPeriod) {
		var ret Period
		return ret
	}
	return *o.CollectedPeriod
}

// GetCollectedPeriodOk returns a tuple with the CollectedPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiologicallyDerivedProductCollection) GetCollectedPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.CollectedPeriod) {
		return nil, false
	}
	return o.CollectedPeriod, true
}

// HasCollectedPeriod returns a boolean if a field has been set.
func (o *BiologicallyDerivedProductCollection) HasCollectedPeriod() bool {
	if o != nil && !IsNil(o.CollectedPeriod) {
		return true
	}

	return false
}

// SetCollectedPeriod gets a reference to the given Period and assigns it to the CollectedPeriod field.
func (o *BiologicallyDerivedProductCollection) SetCollectedPeriod(v Period) {
	o.CollectedPeriod = &v
}

func (o BiologicallyDerivedProductCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BiologicallyDerivedProductCollection) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Collector) {
		toSerialize["collector"] = o.Collector
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.CollectedDateTime) {
		toSerialize["collectedDateTime"] = o.CollectedDateTime
	}
	if !IsNil(o.CollectedPeriod) {
		toSerialize["collectedPeriod"] = o.CollectedPeriod
	}
	return toSerialize, nil
}

type NullableBiologicallyDerivedProductCollection struct {
	value *BiologicallyDerivedProductCollection
	isSet bool
}

func (v NullableBiologicallyDerivedProductCollection) Get() *BiologicallyDerivedProductCollection {
	return v.value
}

func (v *NullableBiologicallyDerivedProductCollection) Set(val *BiologicallyDerivedProductCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableBiologicallyDerivedProductCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableBiologicallyDerivedProductCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiologicallyDerivedProductCollection(val *BiologicallyDerivedProductCollection) *NullableBiologicallyDerivedProductCollection {
	return &NullableBiologicallyDerivedProductCollection{value: val, isSet: true}
}

func (v NullableBiologicallyDerivedProductCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiologicallyDerivedProductCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


