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

// checks if the SubstanceSpecificationOfficial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstanceSpecificationOfficial{}

// SubstanceSpecificationOfficial The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecificationOfficial struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Which authority uses this official name.
	Authority *CodeableConcept `json:"authority,omitempty"`
	// The status of the official name.
	Status *CodeableConcept `json:"status,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	Date *string `json:"date,omitempty"`
}

// NewSubstanceSpecificationOfficial instantiates a new SubstanceSpecificationOfficial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstanceSpecificationOfficial() *SubstanceSpecificationOfficial {
	this := SubstanceSpecificationOfficial{}
	return &this
}

// NewSubstanceSpecificationOfficialWithDefaults instantiates a new SubstanceSpecificationOfficial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstanceSpecificationOfficialWithDefaults() *SubstanceSpecificationOfficial {
	this := SubstanceSpecificationOfficial{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubstanceSpecificationOfficial) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SubstanceSpecificationOfficial) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SubstanceSpecificationOfficial) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetAuthority() CodeableConcept {
	if o == nil || IsNil(o.Authority) {
		var ret CodeableConcept
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetAuthorityOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Authority) {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasAuthority() bool {
	if o != nil && !IsNil(o.Authority) {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given CodeableConcept and assigns it to the Authority field.
func (o *SubstanceSpecificationOfficial) SetAuthority(v CodeableConcept) {
	o.Authority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetStatus() CodeableConcept {
	if o == nil || IsNil(o.Status) {
		var ret CodeableConcept
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetStatusOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CodeableConcept and assigns it to the Status field.
func (o *SubstanceSpecificationOfficial) SetStatus(v CodeableConcept) {
	o.Status = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SubstanceSpecificationOfficial) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstanceSpecificationOfficial) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SubstanceSpecificationOfficial) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SubstanceSpecificationOfficial) SetDate(v string) {
	o.Date = &v
}

func (o SubstanceSpecificationOfficial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubstanceSpecificationOfficial) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Authority) {
		toSerialize["authority"] = o.Authority
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableSubstanceSpecificationOfficial struct {
	value *SubstanceSpecificationOfficial
	isSet bool
}

func (v NullableSubstanceSpecificationOfficial) Get() *SubstanceSpecificationOfficial {
	return v.value
}

func (v *NullableSubstanceSpecificationOfficial) Set(val *SubstanceSpecificationOfficial) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstanceSpecificationOfficial) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstanceSpecificationOfficial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstanceSpecificationOfficial(val *SubstanceSpecificationOfficial) *NullableSubstanceSpecificationOfficial {
	return &NullableSubstanceSpecificationOfficial{value: val, isSet: true}
}

func (v NullableSubstanceSpecificationOfficial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstanceSpecificationOfficial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


