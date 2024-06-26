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

// checks if the Age type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Age{}

// Age A duration of time during which an organism (or a process) has existed.
type Age struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// A rational number with implicit precision
	Value *float32 `json:"value,omitempty"`
	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is \"<\" , then the real value is < stated value.
	Comparator *string `json:"comparator,omitempty"`
	// A sequence of Unicode characters
	Unit *string `json:"unit,omitempty"`
	// String of characters used to identify a name or a resource
	System *string `json:"system,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Code *string `json:"code,omitempty"`
}

// NewAge instantiates a new Age object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAge() *Age {
	this := Age{}
	return &this
}

// NewAgeWithDefaults instantiates a new Age object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeWithDefaults() *Age {
	this := Age{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Age) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Age) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Age) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Age) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Age) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Age) SetExtension(v []Extension) {
	o.Extension = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Age) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Age) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Age) SetValue(v float32) {
	o.Value = &v
}

// GetComparator returns the Comparator field value if set, zero value otherwise.
func (o *Age) GetComparator() string {
	if o == nil || IsNil(o.Comparator) {
		var ret string
		return ret
	}
	return *o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetComparatorOk() (*string, bool) {
	if o == nil || IsNil(o.Comparator) {
		return nil, false
	}
	return o.Comparator, true
}

// HasComparator returns a boolean if a field has been set.
func (o *Age) HasComparator() bool {
	if o != nil && !IsNil(o.Comparator) {
		return true
	}

	return false
}

// SetComparator gets a reference to the given string and assigns it to the Comparator field.
func (o *Age) SetComparator(v string) {
	o.Comparator = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Age) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Age) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Age) SetUnit(v string) {
	o.Unit = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *Age) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *Age) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *Age) SetSystem(v string) {
	o.System = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Age) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Age) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Age) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Age) SetCode(v string) {
	o.Code = &v
}

func (o Age) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Age) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Comparator) {
		toSerialize["comparator"] = o.Comparator
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableAge struct {
	value *Age
	isSet bool
}

func (v NullableAge) Get() *Age {
	return v.value
}

func (v *NullableAge) Set(val *Age) {
	v.value = val
	v.isSet = true
}

func (v NullableAge) IsSet() bool {
	return v.isSet
}

func (v *NullableAge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAge(val *Age) *NullableAge {
	return &NullableAge{value: val, isSet: true}
}

func (v NullableAge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


