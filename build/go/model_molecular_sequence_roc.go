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

// checks if the MolecularSequenceRoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularSequenceRoc{}

// MolecularSequenceRoc Raw data describing a biological sequence.
type MolecularSequenceRoc struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Invidual data point representing the GQ (genotype quality) score threshold.
	Score []float32 `json:"score,omitempty"`
	// The number of true positives if the GQ score threshold was set to \"score\" field value.
	NumTP []float32 `json:"numTP,omitempty"`
	// The number of false positives if the GQ score threshold was set to \"score\" field value.
	NumFP []float32 `json:"numFP,omitempty"`
	// The number of false negatives if the GQ score threshold was set to \"score\" field value.
	NumFN []float32 `json:"numFN,omitempty"`
	// Calculated precision if the GQ score threshold was set to \"score\" field value.
	Precision []float32 `json:"precision,omitempty"`
	// Calculated sensitivity if the GQ score threshold was set to \"score\" field value.
	Sensitivity []float32 `json:"sensitivity,omitempty"`
	// Calculated fScore if the GQ score threshold was set to \"score\" field value.
	FMeasure []float32 `json:"fMeasure,omitempty"`
}

// NewMolecularSequenceRoc instantiates a new MolecularSequenceRoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularSequenceRoc() *MolecularSequenceRoc {
	this := MolecularSequenceRoc{}
	return &this
}

// NewMolecularSequenceRocWithDefaults instantiates a new MolecularSequenceRoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularSequenceRocWithDefaults() *MolecularSequenceRoc {
	this := MolecularSequenceRoc{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MolecularSequenceRoc) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MolecularSequenceRoc) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MolecularSequenceRoc) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetScore() []float32 {
	if o == nil || IsNil(o.Score) {
		var ret []float32
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetScoreOk() ([]float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given []float32 and assigns it to the Score field.
func (o *MolecularSequenceRoc) SetScore(v []float32) {
	o.Score = v
}

// GetNumTP returns the NumTP field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetNumTP() []float32 {
	if o == nil || IsNil(o.NumTP) {
		var ret []float32
		return ret
	}
	return o.NumTP
}

// GetNumTPOk returns a tuple with the NumTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetNumTPOk() ([]float32, bool) {
	if o == nil || IsNil(o.NumTP) {
		return nil, false
	}
	return o.NumTP, true
}

// HasNumTP returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasNumTP() bool {
	if o != nil && !IsNil(o.NumTP) {
		return true
	}

	return false
}

// SetNumTP gets a reference to the given []float32 and assigns it to the NumTP field.
func (o *MolecularSequenceRoc) SetNumTP(v []float32) {
	o.NumTP = v
}

// GetNumFP returns the NumFP field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetNumFP() []float32 {
	if o == nil || IsNil(o.NumFP) {
		var ret []float32
		return ret
	}
	return o.NumFP
}

// GetNumFPOk returns a tuple with the NumFP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetNumFPOk() ([]float32, bool) {
	if o == nil || IsNil(o.NumFP) {
		return nil, false
	}
	return o.NumFP, true
}

// HasNumFP returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasNumFP() bool {
	if o != nil && !IsNil(o.NumFP) {
		return true
	}

	return false
}

// SetNumFP gets a reference to the given []float32 and assigns it to the NumFP field.
func (o *MolecularSequenceRoc) SetNumFP(v []float32) {
	o.NumFP = v
}

// GetNumFN returns the NumFN field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetNumFN() []float32 {
	if o == nil || IsNil(o.NumFN) {
		var ret []float32
		return ret
	}
	return o.NumFN
}

// GetNumFNOk returns a tuple with the NumFN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetNumFNOk() ([]float32, bool) {
	if o == nil || IsNil(o.NumFN) {
		return nil, false
	}
	return o.NumFN, true
}

// HasNumFN returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasNumFN() bool {
	if o != nil && !IsNil(o.NumFN) {
		return true
	}

	return false
}

// SetNumFN gets a reference to the given []float32 and assigns it to the NumFN field.
func (o *MolecularSequenceRoc) SetNumFN(v []float32) {
	o.NumFN = v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetPrecision() []float32 {
	if o == nil || IsNil(o.Precision) {
		var ret []float32
		return ret
	}
	return o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetPrecisionOk() ([]float32, bool) {
	if o == nil || IsNil(o.Precision) {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasPrecision() bool {
	if o != nil && !IsNil(o.Precision) {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given []float32 and assigns it to the Precision field.
func (o *MolecularSequenceRoc) SetPrecision(v []float32) {
	o.Precision = v
}

// GetSensitivity returns the Sensitivity field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetSensitivity() []float32 {
	if o == nil || IsNil(o.Sensitivity) {
		var ret []float32
		return ret
	}
	return o.Sensitivity
}

// GetSensitivityOk returns a tuple with the Sensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetSensitivityOk() ([]float32, bool) {
	if o == nil || IsNil(o.Sensitivity) {
		return nil, false
	}
	return o.Sensitivity, true
}

// HasSensitivity returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasSensitivity() bool {
	if o != nil && !IsNil(o.Sensitivity) {
		return true
	}

	return false
}

// SetSensitivity gets a reference to the given []float32 and assigns it to the Sensitivity field.
func (o *MolecularSequenceRoc) SetSensitivity(v []float32) {
	o.Sensitivity = v
}

// GetFMeasure returns the FMeasure field value if set, zero value otherwise.
func (o *MolecularSequenceRoc) GetFMeasure() []float32 {
	if o == nil || IsNil(o.FMeasure) {
		var ret []float32
		return ret
	}
	return o.FMeasure
}

// GetFMeasureOk returns a tuple with the FMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequenceRoc) GetFMeasureOk() ([]float32, bool) {
	if o == nil || IsNil(o.FMeasure) {
		return nil, false
	}
	return o.FMeasure, true
}

// HasFMeasure returns a boolean if a field has been set.
func (o *MolecularSequenceRoc) HasFMeasure() bool {
	if o != nil && !IsNil(o.FMeasure) {
		return true
	}

	return false
}

// SetFMeasure gets a reference to the given []float32 and assigns it to the FMeasure field.
func (o *MolecularSequenceRoc) SetFMeasure(v []float32) {
	o.FMeasure = v
}

func (o MolecularSequenceRoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularSequenceRoc) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.NumTP) {
		toSerialize["numTP"] = o.NumTP
	}
	if !IsNil(o.NumFP) {
		toSerialize["numFP"] = o.NumFP
	}
	if !IsNil(o.NumFN) {
		toSerialize["numFN"] = o.NumFN
	}
	if !IsNil(o.Precision) {
		toSerialize["precision"] = o.Precision
	}
	if !IsNil(o.Sensitivity) {
		toSerialize["sensitivity"] = o.Sensitivity
	}
	if !IsNil(o.FMeasure) {
		toSerialize["fMeasure"] = o.FMeasure
	}
	return toSerialize, nil
}

type NullableMolecularSequenceRoc struct {
	value *MolecularSequenceRoc
	isSet bool
}

func (v NullableMolecularSequenceRoc) Get() *MolecularSequenceRoc {
	return v.value
}

func (v *NullableMolecularSequenceRoc) Set(val *MolecularSequenceRoc) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularSequenceRoc) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularSequenceRoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularSequenceRoc(val *MolecularSequenceRoc) *NullableMolecularSequenceRoc {
	return &NullableMolecularSequenceRoc{value: val, isSet: true}
}

func (v NullableMolecularSequenceRoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularSequenceRoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

