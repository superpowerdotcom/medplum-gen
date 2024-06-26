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

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location Details and position information for a physical place where services are provided and resources and participants may be stored, found, contained, or accommodated.
type Location struct {
	// This is a Location resource
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
	// Unique code or number identifying the location to its users.
	Identifier []Identifier `json:"identifier,omitempty"`
	// The status property covers the general availability of the resource, not the current value which may be covered by the operationStatus, or by a schedule/slots if they are configured for the location.
	Status *string `json:"status,omitempty"`
	// The operational status covers operation values most relevant to beds (but can also apply to rooms/units/chairs/etc. such as an isolation unit/dialysis chair). This typically covers concepts such as contamination, housekeeping, and other activities like maintenance.
	OperationalStatus *Coding `json:"operationalStatus,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A list of alternate names that the location is known as, or was known as, in the past.
	Alias []string `json:"alias,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// Indicates whether a resource instance represents a specific location or a class of locations.
	Mode *string `json:"mode,omitempty"`
	// Indicates the type of function performed at the location.
	Type []CodeableConcept `json:"type,omitempty"`
	// The contact details of communication devices available at the location. This can include phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Physical location.
	Address *Address `json:"address,omitempty"`
	// Physical form of the location, e.g. building, room, vehicle, road.
	PhysicalType *CodeableConcept `json:"physicalType,omitempty"`
	// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
	Position *LocationPosition `json:"position,omitempty"`
	// The organization responsible for the provisioning and upkeep of the location.
	ManagingOrganization *Reference `json:"managingOrganization,omitempty"`
	// Another Location of which this Location is physically a part of.
	PartOf *Reference `json:"partOf,omitempty"`
	// What days/times during a week is this location usually open.
	HoursOfOperation []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`
	// A sequence of Unicode characters
	AvailabilityExceptions *string `json:"availabilityExceptions,omitempty"`
	// Technical endpoints providing access to services operated for the location.
	Endpoint []Reference `json:"endpoint,omitempty"`
}

type _Location Location

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(resourceType string) *Location {
	this := Location{}
	this.ResourceType = resourceType
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Location) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Location) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Location) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Location) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Location) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Location) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Location) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Location) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Location) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Location) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Location) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Location) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Location) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Location) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Location) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Location) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Location) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Location) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Location) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Location) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *Location) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Location) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Location) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Location) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Location) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Location) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Location) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Location) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Location) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *Location) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Location) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Location) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Location) SetStatus(v string) {
	o.Status = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *Location) GetOperationalStatus() Coding {
	if o == nil || IsNil(o.OperationalStatus) {
		var ret Coding
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetOperationalStatusOk() (*Coding, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *Location) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given Coding and assigns it to the OperationalStatus field.
func (o *Location) SetOperationalStatus(v Coding) {
	o.OperationalStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Location) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Location) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Location) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *Location) GetAlias() []string {
	if o == nil || IsNil(o.Alias) {
		var ret []string
		return ret
	}
	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAliasOk() ([]string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *Location) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given []string and assigns it to the Alias field.
func (o *Location) SetAlias(v []string) {
	o.Alias = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Location) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Location) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Location) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Location) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Location) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Location) SetMode(v string) {
	o.Mode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Location) GetType() []CodeableConcept {
	if o == nil || IsNil(o.Type) {
		var ret []CodeableConcept
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetTypeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Location) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []CodeableConcept and assigns it to the Type field.
func (o *Location) SetType(v []CodeableConcept) {
	o.Type = v
}

// GetTelecom returns the Telecom field value if set, zero value otherwise.
func (o *Location) GetTelecom() []ContactPoint {
	if o == nil || IsNil(o.Telecom) {
		var ret []ContactPoint
		return ret
	}
	return o.Telecom
}

// GetTelecomOk returns a tuple with the Telecom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetTelecomOk() ([]ContactPoint, bool) {
	if o == nil || IsNil(o.Telecom) {
		return nil, false
	}
	return o.Telecom, true
}

// HasTelecom returns a boolean if a field has been set.
func (o *Location) HasTelecom() bool {
	if o != nil && !IsNil(o.Telecom) {
		return true
	}

	return false
}

// SetTelecom gets a reference to the given []ContactPoint and assigns it to the Telecom field.
func (o *Location) SetTelecom(v []ContactPoint) {
	o.Telecom = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Location) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Location) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Location) SetAddress(v Address) {
	o.Address = &v
}

// GetPhysicalType returns the PhysicalType field value if set, zero value otherwise.
func (o *Location) GetPhysicalType() CodeableConcept {
	if o == nil || IsNil(o.PhysicalType) {
		var ret CodeableConcept
		return ret
	}
	return *o.PhysicalType
}

// GetPhysicalTypeOk returns a tuple with the PhysicalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPhysicalTypeOk() (*CodeableConcept, bool) {
	if o == nil || IsNil(o.PhysicalType) {
		return nil, false
	}
	return o.PhysicalType, true
}

// HasPhysicalType returns a boolean if a field has been set.
func (o *Location) HasPhysicalType() bool {
	if o != nil && !IsNil(o.PhysicalType) {
		return true
	}

	return false
}

// SetPhysicalType gets a reference to the given CodeableConcept and assigns it to the PhysicalType field.
func (o *Location) SetPhysicalType(v CodeableConcept) {
	o.PhysicalType = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Location) GetPosition() LocationPosition {
	if o == nil || IsNil(o.Position) {
		var ret LocationPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPositionOk() (*LocationPosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Location) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given LocationPosition and assigns it to the Position field.
func (o *Location) SetPosition(v LocationPosition) {
	o.Position = &v
}

// GetManagingOrganization returns the ManagingOrganization field value if set, zero value otherwise.
func (o *Location) GetManagingOrganization() Reference {
	if o == nil || IsNil(o.ManagingOrganization) {
		var ret Reference
		return ret
	}
	return *o.ManagingOrganization
}

// GetManagingOrganizationOk returns a tuple with the ManagingOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetManagingOrganizationOk() (*Reference, bool) {
	if o == nil || IsNil(o.ManagingOrganization) {
		return nil, false
	}
	return o.ManagingOrganization, true
}

// HasManagingOrganization returns a boolean if a field has been set.
func (o *Location) HasManagingOrganization() bool {
	if o != nil && !IsNil(o.ManagingOrganization) {
		return true
	}

	return false
}

// SetManagingOrganization gets a reference to the given Reference and assigns it to the ManagingOrganization field.
func (o *Location) SetManagingOrganization(v Reference) {
	o.ManagingOrganization = &v
}

// GetPartOf returns the PartOf field value if set, zero value otherwise.
func (o *Location) GetPartOf() Reference {
	if o == nil || IsNil(o.PartOf) {
		var ret Reference
		return ret
	}
	return *o.PartOf
}

// GetPartOfOk returns a tuple with the PartOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPartOfOk() (*Reference, bool) {
	if o == nil || IsNil(o.PartOf) {
		return nil, false
	}
	return o.PartOf, true
}

// HasPartOf returns a boolean if a field has been set.
func (o *Location) HasPartOf() bool {
	if o != nil && !IsNil(o.PartOf) {
		return true
	}

	return false
}

// SetPartOf gets a reference to the given Reference and assigns it to the PartOf field.
func (o *Location) SetPartOf(v Reference) {
	o.PartOf = &v
}

// GetHoursOfOperation returns the HoursOfOperation field value if set, zero value otherwise.
func (o *Location) GetHoursOfOperation() []LocationHoursOfOperation {
	if o == nil || IsNil(o.HoursOfOperation) {
		var ret []LocationHoursOfOperation
		return ret
	}
	return o.HoursOfOperation
}

// GetHoursOfOperationOk returns a tuple with the HoursOfOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetHoursOfOperationOk() ([]LocationHoursOfOperation, bool) {
	if o == nil || IsNil(o.HoursOfOperation) {
		return nil, false
	}
	return o.HoursOfOperation, true
}

// HasHoursOfOperation returns a boolean if a field has been set.
func (o *Location) HasHoursOfOperation() bool {
	if o != nil && !IsNil(o.HoursOfOperation) {
		return true
	}

	return false
}

// SetHoursOfOperation gets a reference to the given []LocationHoursOfOperation and assigns it to the HoursOfOperation field.
func (o *Location) SetHoursOfOperation(v []LocationHoursOfOperation) {
	o.HoursOfOperation = v
}

// GetAvailabilityExceptions returns the AvailabilityExceptions field value if set, zero value otherwise.
func (o *Location) GetAvailabilityExceptions() string {
	if o == nil || IsNil(o.AvailabilityExceptions) {
		var ret string
		return ret
	}
	return *o.AvailabilityExceptions
}

// GetAvailabilityExceptionsOk returns a tuple with the AvailabilityExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAvailabilityExceptionsOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityExceptions) {
		return nil, false
	}
	return o.AvailabilityExceptions, true
}

// HasAvailabilityExceptions returns a boolean if a field has been set.
func (o *Location) HasAvailabilityExceptions() bool {
	if o != nil && !IsNil(o.AvailabilityExceptions) {
		return true
	}

	return false
}

// SetAvailabilityExceptions gets a reference to the given string and assigns it to the AvailabilityExceptions field.
func (o *Location) SetAvailabilityExceptions(v string) {
	o.AvailabilityExceptions = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *Location) GetEndpoint() []Reference {
	if o == nil || IsNil(o.Endpoint) {
		var ret []Reference
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetEndpointOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *Location) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given []Reference and assigns it to the Endpoint field.
func (o *Location) SetEndpoint(v []Reference) {
	o.Endpoint = v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OperationalStatus) {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Telecom) {
		toSerialize["telecom"] = o.Telecom
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.PhysicalType) {
		toSerialize["physicalType"] = o.PhysicalType
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ManagingOrganization) {
		toSerialize["managingOrganization"] = o.ManagingOrganization
	}
	if !IsNil(o.PartOf) {
		toSerialize["partOf"] = o.PartOf
	}
	if !IsNil(o.HoursOfOperation) {
		toSerialize["hoursOfOperation"] = o.HoursOfOperation
	}
	if !IsNil(o.AvailabilityExceptions) {
		toSerialize["availabilityExceptions"] = o.AvailabilityExceptions
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

func (o *Location) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
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

	varLocation := _Location{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocation)

	if err != nil {
		return err
	}

	*o = Location(varLocation)

	return err
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


