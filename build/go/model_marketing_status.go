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

// checks if the MarketingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingStatus{}

// MarketingStatus The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type MarketingStatus struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The country in which the marketing authorisation has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements.
	Country CodeableConcept `json:"country"`
	// Where a Medicines Regulatory Agency has granted a marketing authorisation for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified.
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty"`
	// This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples.
	Status CodeableConcept `json:"status"`
	// The date when the Medicinal Product is placed on the market by the Marketing Authorisation Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	DateRange Period `json:"dateRange"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	RestoreDate *string `json:"restoreDate,omitempty"`
}

type _MarketingStatus MarketingStatus

// NewMarketingStatus instantiates a new MarketingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingStatus(country CodeableConcept, status CodeableConcept, dateRange Period) *MarketingStatus {
	this := MarketingStatus{}
	this.Country = country
	this.Status = status
	this.DateRange = dateRange
	return &this
}

// NewMarketingStatusWithDefaults instantiates a new MarketingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingStatusWithDefaults() *MarketingStatus {
	this := MarketingStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarketingStatus) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarketingStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MarketingStatus) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MarketingStatus) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MarketingStatus) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MarketingStatus) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MarketingStatus) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MarketingStatus) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MarketingStatus) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetCountry returns the Country field value
func (o *MarketingStatus) GetCountry() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetCountryOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *MarketingStatus) SetCountry(v CodeableConcept) {
	o.Country = v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *MarketingStatus) GetJurisdiction() CodeableConcept {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret CodeableConcept
		return ret
	}
	return *o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetJurisdictionOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *MarketingStatus) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given CodeableConcept and assigns it to the Jurisdiction field.
func (o *MarketingStatus) SetJurisdiction(v CodeableConcept) {
	o.Jurisdiction = &v
}

// GetStatus returns the Status field value
func (o *MarketingStatus) GetStatus() CodeableConcept {
	if o == nil {
		var ret CodeableConcept
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetStatusOk() (*CodeableConcept, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MarketingStatus) SetStatus(v CodeableConcept) {
	o.Status = v
}

// GetDateRange returns the DateRange field value
func (o *MarketingStatus) GetDateRange() Period {
	if o == nil {
		var ret Period
		return ret
	}

	return o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetDateRangeOk() (*Period, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateRange, true
}

// SetDateRange sets field value
func (o *MarketingStatus) SetDateRange(v Period) {
	o.DateRange = v
}

// GetRestoreDate returns the RestoreDate field value if set, zero value otherwise.
func (o *MarketingStatus) GetRestoreDate() string {
	if o == nil || IsNil(o.RestoreDate) {
		var ret string
		return ret
	}
	return *o.RestoreDate
}

// GetRestoreDateOk returns a tuple with the RestoreDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingStatus) GetRestoreDateOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreDate) {
		return nil, false
	}
	return o.RestoreDate, true
}

// HasRestoreDate returns a boolean if a field has been set.
func (o *MarketingStatus) HasRestoreDate() bool {
	if o != nil && !IsNil(o.RestoreDate) {
		return true
	}

	return false
}

// SetRestoreDate gets a reference to the given string and assigns it to the RestoreDate field.
func (o *MarketingStatus) SetRestoreDate(v string) {
	o.RestoreDate = &v
}

func (o MarketingStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingStatus) ToMap() (map[string]interface{}, error) {
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
	toSerialize["country"] = o.Country
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	toSerialize["status"] = o.Status
	toSerialize["dateRange"] = o.DateRange
	if !IsNil(o.RestoreDate) {
		toSerialize["restoreDate"] = o.RestoreDate
	}
	return toSerialize, nil
}

func (o *MarketingStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
		"status",
		"dateRange",
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

	varMarketingStatus := _MarketingStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketingStatus)

	if err != nil {
		return err
	}

	*o = MarketingStatus(varMarketingStatus)

	return err
}

type NullableMarketingStatus struct {
	value *MarketingStatus
	isSet bool
}

func (v NullableMarketingStatus) Get() *MarketingStatus {
	return v.value
}

func (v *NullableMarketingStatus) Set(val *MarketingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingStatus(val *MarketingStatus) *NullableMarketingStatus {
	return &NullableMarketingStatus{value: val, isSet: true}
}

func (v NullableMarketingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


