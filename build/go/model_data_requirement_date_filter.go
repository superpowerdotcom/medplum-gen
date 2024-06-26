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

// checks if the DataRequirementDateFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataRequirementDateFilter{}

// DataRequirementDateFilter Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirementDateFilter struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Path *string `json:"path,omitempty"`
	// A sequence of Unicode characters
	SearchParam *string `json:"searchParam,omitempty"`
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDuration *Duration `json:"valueDuration,omitempty"`
}

// NewDataRequirementDateFilter instantiates a new DataRequirementDateFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataRequirementDateFilter() *DataRequirementDateFilter {
	this := DataRequirementDateFilter{}
	return &this
}

// NewDataRequirementDateFilterWithDefaults instantiates a new DataRequirementDateFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataRequirementDateFilterWithDefaults() *DataRequirementDateFilter {
	this := DataRequirementDateFilter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataRequirementDateFilter) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *DataRequirementDateFilter) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *DataRequirementDateFilter) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DataRequirementDateFilter) SetPath(v string) {
	o.Path = &v
}

// GetSearchParam returns the SearchParam field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetSearchParam() string {
	if o == nil || IsNil(o.SearchParam) {
		var ret string
		return ret
	}
	return *o.SearchParam
}

// GetSearchParamOk returns a tuple with the SearchParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetSearchParamOk() (*string, bool) {
	if o == nil || IsNil(o.SearchParam) {
		return nil, false
	}
	return o.SearchParam, true
}

// HasSearchParam returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasSearchParam() bool {
	if o != nil && !IsNil(o.SearchParam) {
		return true
	}

	return false
}

// SetSearchParam gets a reference to the given string and assigns it to the SearchParam field.
func (o *DataRequirementDateFilter) SetSearchParam(v string) {
	o.SearchParam = &v
}

// GetValueDateTime returns the ValueDateTime field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetValueDateTime() string {
	if o == nil || IsNil(o.ValueDateTime) {
		var ret string
		return ret
	}
	return *o.ValueDateTime
}

// GetValueDateTimeOk returns a tuple with the ValueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetValueDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueDateTime) {
		return nil, false
	}
	return o.ValueDateTime, true
}

// HasValueDateTime returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasValueDateTime() bool {
	if o != nil && !IsNil(o.ValueDateTime) {
		return true
	}

	return false
}

// SetValueDateTime gets a reference to the given string and assigns it to the ValueDateTime field.
func (o *DataRequirementDateFilter) SetValueDateTime(v string) {
	o.ValueDateTime = &v
}

// GetValuePeriod returns the ValuePeriod field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetValuePeriod() Period {
	if o == nil || IsNil(o.ValuePeriod) {
		var ret Period
		return ret
	}
	return *o.ValuePeriod
}

// GetValuePeriodOk returns a tuple with the ValuePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetValuePeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.ValuePeriod) {
		return nil, false
	}
	return o.ValuePeriod, true
}

// HasValuePeriod returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasValuePeriod() bool {
	if o != nil && !IsNil(o.ValuePeriod) {
		return true
	}

	return false
}

// SetValuePeriod gets a reference to the given Period and assigns it to the ValuePeriod field.
func (o *DataRequirementDateFilter) SetValuePeriod(v Period) {
	o.ValuePeriod = &v
}

// GetValueDuration returns the ValueDuration field value if set, zero value otherwise.
func (o *DataRequirementDateFilter) GetValueDuration() Duration {
	if o == nil || IsNil(o.ValueDuration) {
		var ret Duration
		return ret
	}
	return *o.ValueDuration
}

// GetValueDurationOk returns a tuple with the ValueDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRequirementDateFilter) GetValueDurationOk() (*Duration, bool) {
	if o == nil || IsNil(o.ValueDuration) {
		return nil, false
	}
	return o.ValueDuration, true
}

// HasValueDuration returns a boolean if a field has been set.
func (o *DataRequirementDateFilter) HasValueDuration() bool {
	if o != nil && !IsNil(o.ValueDuration) {
		return true
	}

	return false
}

// SetValueDuration gets a reference to the given Duration and assigns it to the ValueDuration field.
func (o *DataRequirementDateFilter) SetValueDuration(v Duration) {
	o.ValueDuration = &v
}

func (o DataRequirementDateFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataRequirementDateFilter) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SearchParam) {
		toSerialize["searchParam"] = o.SearchParam
	}
	if !IsNil(o.ValueDateTime) {
		toSerialize["valueDateTime"] = o.ValueDateTime
	}
	if !IsNil(o.ValuePeriod) {
		toSerialize["valuePeriod"] = o.ValuePeriod
	}
	if !IsNil(o.ValueDuration) {
		toSerialize["valueDuration"] = o.ValueDuration
	}
	return toSerialize, nil
}

type NullableDataRequirementDateFilter struct {
	value *DataRequirementDateFilter
	isSet bool
}

func (v NullableDataRequirementDateFilter) Get() *DataRequirementDateFilter {
	return v.value
}

func (v *NullableDataRequirementDateFilter) Set(val *DataRequirementDateFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDataRequirementDateFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDataRequirementDateFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataRequirementDateFilter(val *DataRequirementDateFilter) *NullableDataRequirementDateFilter {
	return &NullableDataRequirementDateFilter{value: val, isSet: true}
}

func (v NullableDataRequirementDateFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataRequirementDateFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

