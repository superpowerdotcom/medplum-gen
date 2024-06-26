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

// checks if the SupplyDeliverySuppliedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyDeliverySuppliedItem{}

// SupplyDeliverySuppliedItem Record of delivery of what is supplied.
type SupplyDeliverySuppliedItem struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The amount of supply that has been dispensed. Includes unit of measure.
	Quantity *Quantity `json:"quantity,omitempty"`
	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemReference *Reference `json:"itemReference,omitempty"`
}

// NewSupplyDeliverySuppliedItem instantiates a new SupplyDeliverySuppliedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyDeliverySuppliedItem() *SupplyDeliverySuppliedItem {
	this := SupplyDeliverySuppliedItem{}
	return &this
}

// NewSupplyDeliverySuppliedItemWithDefaults instantiates a new SupplyDeliverySuppliedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyDeliverySuppliedItemWithDefaults() *SupplyDeliverySuppliedItem {
	this := SupplyDeliverySuppliedItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupplyDeliverySuppliedItem) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SupplyDeliverySuppliedItem) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SupplyDeliverySuppliedItem) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetQuantity() Quantity {
	if o == nil || IsNil(o.Quantity) {
		var ret Quantity
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given Quantity and assigns it to the Quantity field.
func (o *SupplyDeliverySuppliedItem) SetQuantity(v Quantity) {
	o.Quantity = &v
}

// GetItemCodeableConcept returns the ItemCodeableConcept field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetItemCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ItemCodeableConcept
}

// GetItemCodeableConceptOk returns a tuple with the ItemCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetItemCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		return nil, false
	}
	return o.ItemCodeableConcept, true
}

// HasItemCodeableConcept returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasItemCodeableConcept() bool {
	if o != nil && !IsNil(o.ItemCodeableConcept) {
		return true
	}

	return false
}

// SetItemCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ItemCodeableConcept field.
func (o *SupplyDeliverySuppliedItem) SetItemCodeableConcept(v CodeableConcept) {
	o.ItemCodeableConcept = &v
}

// GetItemReference returns the ItemReference field value if set, zero value otherwise.
func (o *SupplyDeliverySuppliedItem) GetItemReference() Reference {
	if o == nil || IsNil(o.ItemReference) {
		var ret Reference
		return ret
	}
	return *o.ItemReference
}

// GetItemReferenceOk returns a tuple with the ItemReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyDeliverySuppliedItem) GetItemReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ItemReference) {
		return nil, false
	}
	return o.ItemReference, true
}

// HasItemReference returns a boolean if a field has been set.
func (o *SupplyDeliverySuppliedItem) HasItemReference() bool {
	if o != nil && !IsNil(o.ItemReference) {
		return true
	}

	return false
}

// SetItemReference gets a reference to the given Reference and assigns it to the ItemReference field.
func (o *SupplyDeliverySuppliedItem) SetItemReference(v Reference) {
	o.ItemReference = &v
}

func (o SupplyDeliverySuppliedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyDeliverySuppliedItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ItemCodeableConcept) {
		toSerialize["itemCodeableConcept"] = o.ItemCodeableConcept
	}
	if !IsNil(o.ItemReference) {
		toSerialize["itemReference"] = o.ItemReference
	}
	return toSerialize, nil
}

type NullableSupplyDeliverySuppliedItem struct {
	value *SupplyDeliverySuppliedItem
	isSet bool
}

func (v NullableSupplyDeliverySuppliedItem) Get() *SupplyDeliverySuppliedItem {
	return v.value
}

func (v *NullableSupplyDeliverySuppliedItem) Set(val *SupplyDeliverySuppliedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyDeliverySuppliedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyDeliverySuppliedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyDeliverySuppliedItem(val *SupplyDeliverySuppliedItem) *NullableSupplyDeliverySuppliedItem {
	return &NullableSupplyDeliverySuppliedItem{value: val, isSet: true}
}

func (v NullableSupplyDeliverySuppliedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyDeliverySuppliedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


