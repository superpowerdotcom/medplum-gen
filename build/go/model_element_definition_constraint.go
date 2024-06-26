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

// checks if the ElementDefinitionConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ElementDefinitionConstraint{}

// ElementDefinitionConstraint Captures constraints on each element within the resource, profile, or extension.
type ElementDefinitionConstraint struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Key *string `json:"key,omitempty"`
	// A sequence of Unicode characters
	Requirements *string `json:"requirements,omitempty"`
	// Identifies the impact constraint violation has on the conformance of the instance.
	Severity *string `json:"severity,omitempty"`
	// A sequence of Unicode characters
	Human *string `json:"human,omitempty"`
	// A sequence of Unicode characters
	Expression *string `json:"expression,omitempty"`
	// A sequence of Unicode characters
	Xpath *string `json:"xpath,omitempty"`
	// A URI that is a reference to a canonical URL on a FHIR resource
	Source *string `json:"source,omitempty"`
}

// NewElementDefinitionConstraint instantiates a new ElementDefinitionConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElementDefinitionConstraint() *ElementDefinitionConstraint {
	this := ElementDefinitionConstraint{}
	return &this
}

// NewElementDefinitionConstraintWithDefaults instantiates a new ElementDefinitionConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElementDefinitionConstraintWithDefaults() *ElementDefinitionConstraint {
	this := ElementDefinitionConstraint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ElementDefinitionConstraint) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *ElementDefinitionConstraint) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *ElementDefinitionConstraint) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ElementDefinitionConstraint) SetKey(v string) {
	o.Key = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetRequirements() string {
	if o == nil || IsNil(o.Requirements) {
		var ret string
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetRequirementsOk() (*string, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given string and assigns it to the Requirements field.
func (o *ElementDefinitionConstraint) SetRequirements(v string) {
	o.Requirements = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ElementDefinitionConstraint) SetSeverity(v string) {
	o.Severity = &v
}

// GetHuman returns the Human field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetHuman() string {
	if o == nil || IsNil(o.Human) {
		var ret string
		return ret
	}
	return *o.Human
}

// GetHumanOk returns a tuple with the Human field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetHumanOk() (*string, bool) {
	if o == nil || IsNil(o.Human) {
		return nil, false
	}
	return o.Human, true
}

// HasHuman returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasHuman() bool {
	if o != nil && !IsNil(o.Human) {
		return true
	}

	return false
}

// SetHuman gets a reference to the given string and assigns it to the Human field.
func (o *ElementDefinitionConstraint) SetHuman(v string) {
	o.Human = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ElementDefinitionConstraint) SetExpression(v string) {
	o.Expression = &v
}

// GetXpath returns the Xpath field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetXpath() string {
	if o == nil || IsNil(o.Xpath) {
		var ret string
		return ret
	}
	return *o.Xpath
}

// GetXpathOk returns a tuple with the Xpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetXpathOk() (*string, bool) {
	if o == nil || IsNil(o.Xpath) {
		return nil, false
	}
	return o.Xpath, true
}

// HasXpath returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasXpath() bool {
	if o != nil && !IsNil(o.Xpath) {
		return true
	}

	return false
}

// SetXpath gets a reference to the given string and assigns it to the Xpath field.
func (o *ElementDefinitionConstraint) SetXpath(v string) {
	o.Xpath = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ElementDefinitionConstraint) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElementDefinitionConstraint) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ElementDefinitionConstraint) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ElementDefinitionConstraint) SetSource(v string) {
	o.Source = &v
}

func (o ElementDefinitionConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ElementDefinitionConstraint) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Human) {
		toSerialize["human"] = o.Human
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Xpath) {
		toSerialize["xpath"] = o.Xpath
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableElementDefinitionConstraint struct {
	value *ElementDefinitionConstraint
	isSet bool
}

func (v NullableElementDefinitionConstraint) Get() *ElementDefinitionConstraint {
	return v.value
}

func (v *NullableElementDefinitionConstraint) Set(val *ElementDefinitionConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableElementDefinitionConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableElementDefinitionConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElementDefinitionConstraint(val *ElementDefinitionConstraint) *NullableElementDefinitionConstraint {
	return &NullableElementDefinitionConstraint{value: val, isSet: true}
}

func (v NullableElementDefinitionConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElementDefinitionConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


