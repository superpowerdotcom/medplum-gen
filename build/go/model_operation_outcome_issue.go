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

// checks if the OperationOutcomeIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationOutcomeIssue{}

// OperationOutcomeIssue A collection of error, warning, or information messages that result from a system action.
type OperationOutcomeIssue struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Indicates whether the issue indicates a variation from successful processing.
	Severity *string `json:"severity,omitempty"`
	// Describes the type of the issue. The system that creates an OperationOutcome SHALL choose the most applicable code from the IssueType value set, and may additional provide its own code for the error in the details element.
	Code *string `json:"code,omitempty"`
	// Additional details about the error. This may be a text description of the error or a system code that identifies the error.
	Details *CodeableConcept `json:"details,omitempty"`
	// A sequence of Unicode characters
	Diagnostics *string `json:"diagnostics,omitempty"`
	// This element is deprecated because it is XML specific. It is replaced by issue.expression, which is format independent, and simpler to parse.   For resource issues, this will be a simple XPath limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.  For HTTP errors, will be \"http.\" + the parameter name.
	Location []string `json:"location,omitempty"`
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	Expression []string `json:"expression,omitempty"`
}

// NewOperationOutcomeIssue instantiates a new OperationOutcomeIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationOutcomeIssue() *OperationOutcomeIssue {
	this := OperationOutcomeIssue{}
	return &this
}

// NewOperationOutcomeIssueWithDefaults instantiates a new OperationOutcomeIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationOutcomeIssueWithDefaults() *OperationOutcomeIssue {
	this := OperationOutcomeIssue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OperationOutcomeIssue) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *OperationOutcomeIssue) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *OperationOutcomeIssue) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *OperationOutcomeIssue) SetSeverity(v string) {
	o.Severity = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OperationOutcomeIssue) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetDetails() CodeableConcept {
	if o == nil || IsNil(o.Details) {
		var ret CodeableConcept
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetDetailsOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given CodeableConcept and assigns it to the Details field.
func (o *OperationOutcomeIssue) SetDetails(v CodeableConcept) {
	o.Details = &v
}

// GetDiagnostics returns the Diagnostics field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetDiagnostics() string {
	if o == nil || IsNil(o.Diagnostics) {
		var ret string
		return ret
	}
	return *o.Diagnostics
}

// GetDiagnosticsOk returns a tuple with the Diagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetDiagnosticsOk() (*string, bool) {
	if o == nil || IsNil(o.Diagnostics) {
		return nil, false
	}
	return o.Diagnostics, true
}

// HasDiagnostics returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasDiagnostics() bool {
	if o != nil && !IsNil(o.Diagnostics) {
		return true
	}

	return false
}

// SetDiagnostics gets a reference to the given string and assigns it to the Diagnostics field.
func (o *OperationOutcomeIssue) SetDiagnostics(v string) {
	o.Diagnostics = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetLocation() []string {
	if o == nil || IsNil(o.Location) {
		var ret []string
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetLocationOk() ([]string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given []string and assigns it to the Location field.
func (o *OperationOutcomeIssue) SetLocation(v []string) {
	o.Location = v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *OperationOutcomeIssue) GetExpression() []string {
	if o == nil || IsNil(o.Expression) {
		var ret []string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationOutcomeIssue) GetExpressionOk() ([]string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *OperationOutcomeIssue) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []string and assigns it to the Expression field.
func (o *OperationOutcomeIssue) SetExpression(v []string) {
	o.Expression = v
}

func (o OperationOutcomeIssue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationOutcomeIssue) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Diagnostics) {
		toSerialize["diagnostics"] = o.Diagnostics
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

type NullableOperationOutcomeIssue struct {
	value *OperationOutcomeIssue
	isSet bool
}

func (v NullableOperationOutcomeIssue) Get() *OperationOutcomeIssue {
	return v.value
}

func (v *NullableOperationOutcomeIssue) Set(val *OperationOutcomeIssue) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationOutcomeIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationOutcomeIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationOutcomeIssue(val *OperationOutcomeIssue) *NullableOperationOutcomeIssue {
	return &NullableOperationOutcomeIssue{value: val, isSet: true}
}

func (v NullableOperationOutcomeIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationOutcomeIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


