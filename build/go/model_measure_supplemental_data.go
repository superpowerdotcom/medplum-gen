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

// checks if the MeasureSupplementalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasureSupplementalData{}

// MeasureSupplementalData The Measure resource provides the definition of a quality measure.
type MeasureSupplementalData struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Indicates a meaning for the supplemental data. This can be as simple as a unique identifier, or it can establish meaning in a broader context by drawing from a terminology, allowing supplemental data to be correlated across measures.
	Code *CodeableConcept `json:"code,omitempty"`
	// An indicator of the intended usage for the supplemental data element. Supplemental data indicates the data is additional information requested to augment the measure information. Risk adjustment factor indicates the data is additional information used to calculate risk adjustment factors when applying a risk model to the measure calculation.
	Usage []CodeableConcept `json:"usage,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// The criteria for the supplemental data. This is typically the name of a valid expression defined within a referenced library, but it may also be a path to a specific data element. The criteria defines the data to be returned for this element.
	Criteria Expression `json:"criteria"`
}

type _MeasureSupplementalData MeasureSupplementalData

// NewMeasureSupplementalData instantiates a new MeasureSupplementalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasureSupplementalData(criteria Expression) *MeasureSupplementalData {
	this := MeasureSupplementalData{}
	this.Criteria = criteria
	return &this
}

// NewMeasureSupplementalDataWithDefaults instantiates a new MeasureSupplementalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasureSupplementalDataWithDefaults() *MeasureSupplementalData {
	this := MeasureSupplementalData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MeasureSupplementalData) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MeasureSupplementalData) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MeasureSupplementalData) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetCode() CodeableConcept {
	if o == nil || IsNil(o.Code) {
		var ret CodeableConcept
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetCodeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given CodeableConcept and assigns it to the Code field.
func (o *MeasureSupplementalData) SetCode(v CodeableConcept) {
	o.Code = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetUsage() []CodeableConcept {
	if o == nil || IsNil(o.Usage) {
		var ret []CodeableConcept
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetUsageOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []CodeableConcept and assigns it to the Usage field.
func (o *MeasureSupplementalData) SetUsage(v []CodeableConcept) {
	o.Usage = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MeasureSupplementalData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MeasureSupplementalData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MeasureSupplementalData) SetDescription(v string) {
	o.Description = &v
}

// GetCriteria returns the Criteria field value
func (o *MeasureSupplementalData) GetCriteria() Expression {
	if o == nil {
		var ret Expression
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *MeasureSupplementalData) GetCriteriaOk() (*Expression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// SetCriteria sets field value
func (o *MeasureSupplementalData) SetCriteria(v Expression) {
	o.Criteria = v
}

func (o MeasureSupplementalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasureSupplementalData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["criteria"] = o.Criteria
	return toSerialize, nil
}

func (o *MeasureSupplementalData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"criteria",
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

	varMeasureSupplementalData := _MeasureSupplementalData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeasureSupplementalData)

	if err != nil {
		return err
	}

	*o = MeasureSupplementalData(varMeasureSupplementalData)

	return err
}

type NullableMeasureSupplementalData struct {
	value *MeasureSupplementalData
	isSet bool
}

func (v NullableMeasureSupplementalData) Get() *MeasureSupplementalData {
	return v.value
}

func (v *NullableMeasureSupplementalData) Set(val *MeasureSupplementalData) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasureSupplementalData) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasureSupplementalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasureSupplementalData(val *MeasureSupplementalData) *NullableMeasureSupplementalData {
	return &NullableMeasureSupplementalData{value: val, isSet: true}
}

func (v NullableMeasureSupplementalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasureSupplementalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


