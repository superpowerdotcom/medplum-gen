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

// checks if the OrganizationContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationContact{}

// OrganizationContact A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.  Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, payer/insurer, etc.
type OrganizationContact struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Indicates a purpose for which the contact can be reached.
	Purpose *CodeableConcept `json:"purpose,omitempty"`
	// A name associated with the contact.
	Name *HumanName `json:"name,omitempty"`
	// A contact detail (e.g. a telephone number or an email address) by which the party may be contacted.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Visiting or postal addresses for the contact.
	Address *Address `json:"address,omitempty"`
}

// NewOrganizationContact instantiates a new OrganizationContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationContact() *OrganizationContact {
	this := OrganizationContact{}
	return &this
}

// NewOrganizationContactWithDefaults instantiates a new OrganizationContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationContactWithDefaults() *OrganizationContact {
	this := OrganizationContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationContact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationContact) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *OrganizationContact) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *OrganizationContact) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *OrganizationContact) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *OrganizationContact) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *OrganizationContact) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *OrganizationContact) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *OrganizationContact) GetPurpose() CodeableConcept {
	if o == nil || IsNil(o.Purpose) {
		var ret CodeableConcept
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetPurposeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *OrganizationContact) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given CodeableConcept and assigns it to the Purpose field.
func (o *OrganizationContact) SetPurpose(v CodeableConcept) {
	o.Purpose = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationContact) GetName() HumanName {
	if o == nil || IsNil(o.Name) {
		var ret HumanName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetNameOk() (*HumanName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationContact) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given HumanName and assigns it to the Name field.
func (o *OrganizationContact) SetName(v HumanName) {
	o.Name = &v
}

// GetTelecom returns the Telecom field value if set, zero value otherwise.
func (o *OrganizationContact) GetTelecom() []ContactPoint {
	if o == nil || IsNil(o.Telecom) {
		var ret []ContactPoint
		return ret
	}
	return o.Telecom
}

// GetTelecomOk returns a tuple with the Telecom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetTelecomOk() ([]ContactPoint, bool) {
	if o == nil || IsNil(o.Telecom) {
		return nil, false
	}
	return o.Telecom, true
}

// HasTelecom returns a boolean if a field has been set.
func (o *OrganizationContact) HasTelecom() bool {
	if o != nil && !IsNil(o.Telecom) {
		return true
	}

	return false
}

// SetTelecom gets a reference to the given []ContactPoint and assigns it to the Telecom field.
func (o *OrganizationContact) SetTelecom(v []ContactPoint) {
	o.Telecom = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationContact) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationContact) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationContact) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *OrganizationContact) SetAddress(v Address) {
	o.Address = &v
}

func (o OrganizationContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationContact) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Telecom) {
		toSerialize["telecom"] = o.Telecom
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableOrganizationContact struct {
	value *OrganizationContact
	isSet bool
}

func (v NullableOrganizationContact) Get() *OrganizationContact {
	return v.value
}

func (v *NullableOrganizationContact) Set(val *OrganizationContact) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationContact) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationContact(val *OrganizationContact) *NullableOrganizationContact {
	return &NullableOrganizationContact{value: val, isSet: true}
}

func (v NullableOrganizationContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


