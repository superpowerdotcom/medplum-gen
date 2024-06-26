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

// checks if the SupplyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyRequest{}

// SupplyRequest A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest struct {
	// This is a SupplyRequest resource
	ResourceType string `json:"resourceType"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	Id *string `json:"id,omitempty"`
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta `json:"meta,omitempty"`
	// String of characters used to identify a name or a resource
	ImplicitRules *string `json:"implicitRules,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Language *string `json:"language,omitempty"`
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it \"clinically safe\" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative `json:"text,omitempty"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []ResourceList `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Business identifiers assigned to this SupplyRequest by the author and/or other systems. These identifiers remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Status of the supply request.
	Status *string `json:"status,omitempty"`
	// Category of supply, e.g.  central, non-stock, etc. This is used to support work flows associated with the supply process.
	Category *CodeableConcept `json:"category,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Priority *string `json:"priority,omitempty"`
	// The item that is requested to be supplied. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	// The item that is requested to be supplied. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemReference *Reference `json:"itemReference,omitempty"`
	// The amount that is being ordered of the indicated item.
	Quantity Quantity `json:"quantity"`
	// Specific parameters for the ordered item.  For example, the size of the indicated item.
	Parameter []SupplyRequestParameter `json:"parameter,omitempty"`
	// When the request should be fulfilled.
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	// When the request should be fulfilled.
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When the request should be fulfilled.
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// A date, date-time or partial date (e.g. just year or year + month).  If hours and minutes are specified, a time zone SHALL be populated. The format is a union of the schema types gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored.                 Dates SHALL be valid dates.
	AuthoredOn *string `json:"authoredOn,omitempty"`
	// The device, practitioner, etc. who initiated the request.
	Requester *Reference `json:"requester,omitempty"`
	// Who is intended to fulfill the request.
	Supplier []Reference `json:"supplier,omitempty"`
	// The reason why the supply item was requested.
	ReasonCode []CodeableConcept `json:"reasonCode,omitempty"`
	// The reason why the supply item was requested.
	ReasonReference []Reference `json:"reasonReference,omitempty"`
	// Where the supply is expected to come from.
	DeliverFrom *Reference `json:"deliverFrom,omitempty"`
	// Where the supply is destined to go.
	DeliverTo *Reference `json:"deliverTo,omitempty"`
}

type _SupplyRequest SupplyRequest

// NewSupplyRequest instantiates a new SupplyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyRequest(resourceType string, quantity Quantity) *SupplyRequest {
	this := SupplyRequest{}
	this.ResourceType = resourceType
	this.Quantity = quantity
	return &this
}

// NewSupplyRequestWithDefaults instantiates a new SupplyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyRequestWithDefaults() *SupplyRequest {
	this := SupplyRequest{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *SupplyRequest) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *SupplyRequest) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupplyRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupplyRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupplyRequest) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SupplyRequest) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SupplyRequest) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *SupplyRequest) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *SupplyRequest) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *SupplyRequest) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *SupplyRequest) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SupplyRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SupplyRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *SupplyRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SupplyRequest) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SupplyRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *SupplyRequest) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *SupplyRequest) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *SupplyRequest) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *SupplyRequest) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *SupplyRequest) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *SupplyRequest) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *SupplyRequest) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *SupplyRequest) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *SupplyRequest) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *SupplyRequest) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *SupplyRequest) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *SupplyRequest) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *SupplyRequest) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SupplyRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SupplyRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SupplyRequest) SetStatus(v string) {
	o.Status = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SupplyRequest) GetCategory() CodeableConcept {
	if o == nil || IsNil(o.Category) {
		var ret CodeableConcept
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetCategoryOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SupplyRequest) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CodeableConcept and assigns it to the Category field.
func (o *SupplyRequest) SetCategory(v CodeableConcept) {
	o.Category = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SupplyRequest) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SupplyRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SupplyRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetItemCodeableConcept returns the ItemCodeableConcept field value if set, zero value otherwise.
func (o *SupplyRequest) GetItemCodeableConcept() CodeableConcept {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		var ret CodeableConcept
		return ret
	}
	return *o.ItemCodeableConcept
}

// GetItemCodeableConceptOk returns a tuple with the ItemCodeableConcept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetItemCodeableConceptOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.ItemCodeableConcept) {
		return nil, false
	}
	return o.ItemCodeableConcept, true
}

// HasItemCodeableConcept returns a boolean if a field has been set.
func (o *SupplyRequest) HasItemCodeableConcept() bool {
	if o != nil && !IsNil(o.ItemCodeableConcept) {
		return true
	}

	return false
}

// SetItemCodeableConcept gets a reference to the given CodeableConcept and assigns it to the ItemCodeableConcept field.
func (o *SupplyRequest) SetItemCodeableConcept(v CodeableConcept) {
	o.ItemCodeableConcept = &v
}

// GetItemReference returns the ItemReference field value if set, zero value otherwise.
func (o *SupplyRequest) GetItemReference() Reference {
	if o == nil || IsNil(o.ItemReference) {
		var ret Reference
		return ret
	}
	return *o.ItemReference
}

// GetItemReferenceOk returns a tuple with the ItemReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetItemReferenceOk() (*Reference, bool) {
	if o == nil || IsNil(o.ItemReference) {
		return nil, false
	}
	return o.ItemReference, true
}

// HasItemReference returns a boolean if a field has been set.
func (o *SupplyRequest) HasItemReference() bool {
	if o != nil && !IsNil(o.ItemReference) {
		return true
	}

	return false
}

// SetItemReference gets a reference to the given Reference and assigns it to the ItemReference field.
func (o *SupplyRequest) SetItemReference(v Reference) {
	o.ItemReference = &v
}

// GetQuantity returns the Quantity field value
func (o *SupplyRequest) GetQuantity() Quantity {
	if o == nil {
		var ret Quantity
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetQuantityOk() (*Quantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SupplyRequest) SetQuantity(v Quantity) {
	o.Quantity = v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *SupplyRequest) GetParameter() []SupplyRequestParameter {
	if o == nil || IsNil(o.Parameter) {
		var ret []SupplyRequestParameter
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetParameterOk() ([]SupplyRequestParameter, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *SupplyRequest) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []SupplyRequestParameter and assigns it to the Parameter field.
func (o *SupplyRequest) SetParameter(v []SupplyRequestParameter) {
	o.Parameter = v
}

// GetOccurrenceDateTime returns the OccurrenceDateTime field value if set, zero value otherwise.
func (o *SupplyRequest) GetOccurrenceDateTime() string {
	if o == nil || IsNil(o.OccurrenceDateTime) {
		var ret string
		return ret
	}
	return *o.OccurrenceDateTime
}

// GetOccurrenceDateTimeOk returns a tuple with the OccurrenceDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetOccurrenceDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OccurrenceDateTime) {
		return nil, false
	}
	return o.OccurrenceDateTime, true
}

// HasOccurrenceDateTime returns a boolean if a field has been set.
func (o *SupplyRequest) HasOccurrenceDateTime() bool {
	if o != nil && !IsNil(o.OccurrenceDateTime) {
		return true
	}

	return false
}

// SetOccurrenceDateTime gets a reference to the given string and assigns it to the OccurrenceDateTime field.
func (o *SupplyRequest) SetOccurrenceDateTime(v string) {
	o.OccurrenceDateTime = &v
}

// GetOccurrencePeriod returns the OccurrencePeriod field value if set, zero value otherwise.
func (o *SupplyRequest) GetOccurrencePeriod() Period {
	if o == nil || IsNil(o.OccurrencePeriod) {
		var ret Period
		return ret
	}
	return *o.OccurrencePeriod
}

// GetOccurrencePeriodOk returns a tuple with the OccurrencePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetOccurrencePeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.OccurrencePeriod) {
		return nil, false
	}
	return o.OccurrencePeriod, true
}

// HasOccurrencePeriod returns a boolean if a field has been set.
func (o *SupplyRequest) HasOccurrencePeriod() bool {
	if o != nil && !IsNil(o.OccurrencePeriod) {
		return true
	}

	return false
}

// SetOccurrencePeriod gets a reference to the given Period and assigns it to the OccurrencePeriod field.
func (o *SupplyRequest) SetOccurrencePeriod(v Period) {
	o.OccurrencePeriod = &v
}

// GetOccurrenceTiming returns the OccurrenceTiming field value if set, zero value otherwise.
func (o *SupplyRequest) GetOccurrenceTiming() Timing {
	if o == nil || IsNil(o.OccurrenceTiming) {
		var ret Timing
		return ret
	}
	return *o.OccurrenceTiming
}

// GetOccurrenceTimingOk returns a tuple with the OccurrenceTiming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetOccurrenceTimingOk() (*Timing, bool) {
	if o == nil || IsNil(o.OccurrenceTiming) {
		return nil, false
	}
	return o.OccurrenceTiming, true
}

// HasOccurrenceTiming returns a boolean if a field has been set.
func (o *SupplyRequest) HasOccurrenceTiming() bool {
	if o != nil && !IsNil(o.OccurrenceTiming) {
		return true
	}

	return false
}

// SetOccurrenceTiming gets a reference to the given Timing and assigns it to the OccurrenceTiming field.
func (o *SupplyRequest) SetOccurrenceTiming(v Timing) {
	o.OccurrenceTiming = &v
}

// GetAuthoredOn returns the AuthoredOn field value if set, zero value otherwise.
func (o *SupplyRequest) GetAuthoredOn() string {
	if o == nil || IsNil(o.AuthoredOn) {
		var ret string
		return ret
	}
	return *o.AuthoredOn
}

// GetAuthoredOnOk returns a tuple with the AuthoredOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetAuthoredOnOk() (*string, bool) {
	if o == nil || IsNil(o.AuthoredOn) {
		return nil, false
	}
	return o.AuthoredOn, true
}

// HasAuthoredOn returns a boolean if a field has been set.
func (o *SupplyRequest) HasAuthoredOn() bool {
	if o != nil && !IsNil(o.AuthoredOn) {
		return true
	}

	return false
}

// SetAuthoredOn gets a reference to the given string and assigns it to the AuthoredOn field.
func (o *SupplyRequest) SetAuthoredOn(v string) {
	o.AuthoredOn = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *SupplyRequest) GetRequester() Reference {
	if o == nil || IsNil(o.Requester) {
		var ret Reference
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetRequesterOk() (*Reference, bool) {
	if o == nil || IsNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *SupplyRequest) HasRequester() bool {
	if o != nil && !IsNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given Reference and assigns it to the Requester field.
func (o *SupplyRequest) SetRequester(v Reference) {
	o.Requester = &v
}

// GetSupplier returns the Supplier field value if set, zero value otherwise.
func (o *SupplyRequest) GetSupplier() []Reference {
	if o == nil || IsNil(o.Supplier) {
		var ret []Reference
		return ret
	}
	return o.Supplier
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetSupplierOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Supplier) {
		return nil, false
	}
	return o.Supplier, true
}

// HasSupplier returns a boolean if a field has been set.
func (o *SupplyRequest) HasSupplier() bool {
	if o != nil && !IsNil(o.Supplier) {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given []Reference and assigns it to the Supplier field.
func (o *SupplyRequest) SetSupplier(v []Reference) {
	o.Supplier = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *SupplyRequest) GetReasonCode() []CodeableConcept {
	if o == nil || IsNil(o.ReasonCode) {
		var ret []CodeableConcept
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetReasonCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *SupplyRequest) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given []CodeableConcept and assigns it to the ReasonCode field.
func (o *SupplyRequest) SetReasonCode(v []CodeableConcept) {
	o.ReasonCode = v
}

// GetReasonReference returns the ReasonReference field value if set, zero value otherwise.
func (o *SupplyRequest) GetReasonReference() []Reference {
	if o == nil || IsNil(o.ReasonReference) {
		var ret []Reference
		return ret
	}
	return o.ReasonReference
}

// GetReasonReferenceOk returns a tuple with the ReasonReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetReasonReferenceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.ReasonReference) {
		return nil, false
	}
	return o.ReasonReference, true
}

// HasReasonReference returns a boolean if a field has been set.
func (o *SupplyRequest) HasReasonReference() bool {
	if o != nil && !IsNil(o.ReasonReference) {
		return true
	}

	return false
}

// SetReasonReference gets a reference to the given []Reference and assigns it to the ReasonReference field.
func (o *SupplyRequest) SetReasonReference(v []Reference) {
	o.ReasonReference = v
}

// GetDeliverFrom returns the DeliverFrom field value if set, zero value otherwise.
func (o *SupplyRequest) GetDeliverFrom() Reference {
	if o == nil || IsNil(o.DeliverFrom) {
		var ret Reference
		return ret
	}
	return *o.DeliverFrom
}

// GetDeliverFromOk returns a tuple with the DeliverFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetDeliverFromOk() (*Reference, bool) {
	if o == nil || IsNil(o.DeliverFrom) {
		return nil, false
	}
	return o.DeliverFrom, true
}

// HasDeliverFrom returns a boolean if a field has been set.
func (o *SupplyRequest) HasDeliverFrom() bool {
	if o != nil && !IsNil(o.DeliverFrom) {
		return true
	}

	return false
}

// SetDeliverFrom gets a reference to the given Reference and assigns it to the DeliverFrom field.
func (o *SupplyRequest) SetDeliverFrom(v Reference) {
	o.DeliverFrom = &v
}

// GetDeliverTo returns the DeliverTo field value if set, zero value otherwise.
func (o *SupplyRequest) GetDeliverTo() Reference {
	if o == nil || IsNil(o.DeliverTo) {
		var ret Reference
		return ret
	}
	return *o.DeliverTo
}

// GetDeliverToOk returns a tuple with the DeliverTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyRequest) GetDeliverToOk() (*Reference, bool) {
	if o == nil || IsNil(o.DeliverTo) {
		return nil, false
	}
	return o.DeliverTo, true
}

// HasDeliverTo returns a boolean if a field has been set.
func (o *SupplyRequest) HasDeliverTo() bool {
	if o != nil && !IsNil(o.DeliverTo) {
		return true
	}

	return false
}

// SetDeliverTo gets a reference to the given Reference and assigns it to the DeliverTo field.
func (o *SupplyRequest) SetDeliverTo(v Reference) {
	o.DeliverTo = &v
}

func (o SupplyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ImplicitRules) {
		toSerialize["implicitRules"] = o.ImplicitRules
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Contained) {
		toSerialize["contained"] = o.Contained
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.ModifierExtension) {
		toSerialize["modifierExtension"] = o.ModifierExtension
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ItemCodeableConcept) {
		toSerialize["itemCodeableConcept"] = o.ItemCodeableConcept
	}
	if !IsNil(o.ItemReference) {
		toSerialize["itemReference"] = o.ItemReference
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.OccurrenceDateTime) {
		toSerialize["occurrenceDateTime"] = o.OccurrenceDateTime
	}
	if !IsNil(o.OccurrencePeriod) {
		toSerialize["occurrencePeriod"] = o.OccurrencePeriod
	}
	if !IsNil(o.OccurrenceTiming) {
		toSerialize["occurrenceTiming"] = o.OccurrenceTiming
	}
	if !IsNil(o.AuthoredOn) {
		toSerialize["authoredOn"] = o.AuthoredOn
	}
	if !IsNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !IsNil(o.Supplier) {
		toSerialize["supplier"] = o.Supplier
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if !IsNil(o.ReasonReference) {
		toSerialize["reasonReference"] = o.ReasonReference
	}
	if !IsNil(o.DeliverFrom) {
		toSerialize["deliverFrom"] = o.DeliverFrom
	}
	if !IsNil(o.DeliverTo) {
		toSerialize["deliverTo"] = o.DeliverTo
	}
	return toSerialize, nil
}

func (o *SupplyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"quantity",
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

	varSupplyRequest := _SupplyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyRequest)

	if err != nil {
		return err
	}

	*o = SupplyRequest(varSupplyRequest)

	return err
}

type NullableSupplyRequest struct {
	value *SupplyRequest
	isSet bool
}

func (v NullableSupplyRequest) Get() *SupplyRequest {
	return v.value
}

func (v *NullableSupplyRequest) Set(val *SupplyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyRequest(val *SupplyRequest) *NullableSupplyRequest {
	return &NullableSupplyRequest{value: val, isSet: true}
}

func (v NullableSupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


