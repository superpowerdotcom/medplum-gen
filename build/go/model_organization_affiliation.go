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

// checks if the OrganizationAffiliation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAffiliation{}

// OrganizationAffiliation Defines an affiliation/assotiation/relationship between 2 distinct oganizations, that is not a part-of relationship/sub-division relationship.
type OrganizationAffiliation struct {
	// This is a OrganizationAffiliation resource
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
	// Business identifiers that are specific to this role.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Value of \"true\" or \"false\"
	Active *bool `json:"active,omitempty"`
	// The period during which the participatingOrganization is affiliated with the primary organization.
	Period *Period `json:"period,omitempty"`
	// Organization where the role is available (primary organization/has members).
	Organization *Reference `json:"organization,omitempty"`
	// The Participating Organization provides/performs the role(s) defined by the code to the Primary Organization (e.g. providing services or is a member of).
	ParticipatingOrganization *Reference `json:"participatingOrganization,omitempty"`
	// Health insurance provider network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined).
	Network []Reference `json:"network,omitempty"`
	// Definition of the role the participatingOrganization plays in the association.
	Code []CodeableConcept `json:"code,omitempty"`
	// Specific specialty of the participatingOrganization in the context of the role.
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// The location(s) at which the role occurs.
	Location []Reference `json:"location,omitempty"`
	// Healthcare services provided through the role.
	HealthcareService []Reference `json:"healthcareService,omitempty"`
	// Contact details at the participatingOrganization relevant to this Affiliation.
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Technical endpoints providing access to services operated for this role.
	Endpoint []Reference `json:"endpoint,omitempty"`
}

type _OrganizationAffiliation OrganizationAffiliation

// NewOrganizationAffiliation instantiates a new OrganizationAffiliation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAffiliation(resourceType string) *OrganizationAffiliation {
	this := OrganizationAffiliation{}
	this.ResourceType = resourceType
	return &this
}

// NewOrganizationAffiliationWithDefaults instantiates a new OrganizationAffiliation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAffiliationWithDefaults() *OrganizationAffiliation {
	this := OrganizationAffiliation{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *OrganizationAffiliation) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *OrganizationAffiliation) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationAffiliation) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *OrganizationAffiliation) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *OrganizationAffiliation) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *OrganizationAffiliation) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *OrganizationAffiliation) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *OrganizationAffiliation) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *OrganizationAffiliation) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *OrganizationAffiliation) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *OrganizationAffiliation) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *OrganizationAffiliation) SetActive(v bool) {
	o.Active = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetPeriod() Period {
	if o == nil || IsNil(o.Period) {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetPeriodOk() (*Period, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *OrganizationAffiliation) SetPeriod(v Period) {
	o.Period = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetOrganization() Reference {
	if o == nil || IsNil(o.Organization) {
		var ret Reference
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetOrganizationOk() (*Reference, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Reference and assigns it to the Organization field.
func (o *OrganizationAffiliation) SetOrganization(v Reference) {
	o.Organization = &v
}

// GetParticipatingOrganization returns the ParticipatingOrganization field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetParticipatingOrganization() Reference {
	if o == nil || IsNil(o.ParticipatingOrganization) {
		var ret Reference
		return ret
	}
	return *o.ParticipatingOrganization
}

// GetParticipatingOrganizationOk returns a tuple with the ParticipatingOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetParticipatingOrganizationOk() (*Reference, bool) {
	if o == nil || IsNil(o.ParticipatingOrganization) {
		return nil, false
	}
	return o.ParticipatingOrganization, true
}

// HasParticipatingOrganization returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasParticipatingOrganization() bool {
	if o != nil && !IsNil(o.ParticipatingOrganization) {
		return true
	}

	return false
}

// SetParticipatingOrganization gets a reference to the given Reference and assigns it to the ParticipatingOrganization field.
func (o *OrganizationAffiliation) SetParticipatingOrganization(v Reference) {
	o.ParticipatingOrganization = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetNetwork() []Reference {
	if o == nil || IsNil(o.Network) {
		var ret []Reference
		return ret
	}
	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetNetworkOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given []Reference and assigns it to the Network field.
func (o *OrganizationAffiliation) SetNetwork(v []Reference) {
	o.Network = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetCode() []CodeableConcept {
	if o == nil || IsNil(o.Code) {
		var ret []CodeableConcept
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetCodeOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given []CodeableConcept and assigns it to the Code field.
func (o *OrganizationAffiliation) SetCode(v []CodeableConcept) {
	o.Code = v
}

// GetSpecialty returns the Specialty field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetSpecialty() []CodeableConcept {
	if o == nil || IsNil(o.Specialty) {
		var ret []CodeableConcept
		return ret
	}
	return o.Specialty
}

// GetSpecialtyOk returns a tuple with the Specialty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetSpecialtyOk() ([]CodeableConcept, bool) {
	if o == nil || IsNil(o.Specialty) {
		return nil, false
	}
	return o.Specialty, true
}

// HasSpecialty returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasSpecialty() bool {
	if o != nil && !IsNil(o.Specialty) {
		return true
	}

	return false
}

// SetSpecialty gets a reference to the given []CodeableConcept and assigns it to the Specialty field.
func (o *OrganizationAffiliation) SetSpecialty(v []CodeableConcept) {
	o.Specialty = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetLocation() []Reference {
	if o == nil || IsNil(o.Location) {
		var ret []Reference
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetLocationOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given []Reference and assigns it to the Location field.
func (o *OrganizationAffiliation) SetLocation(v []Reference) {
	o.Location = v
}

// GetHealthcareService returns the HealthcareService field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetHealthcareService() []Reference {
	if o == nil || IsNil(o.HealthcareService) {
		var ret []Reference
		return ret
	}
	return o.HealthcareService
}

// GetHealthcareServiceOk returns a tuple with the HealthcareService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetHealthcareServiceOk() ([]Reference, bool) {
	if o == nil || IsNil(o.HealthcareService) {
		return nil, false
	}
	return o.HealthcareService, true
}

// HasHealthcareService returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasHealthcareService() bool {
	if o != nil && !IsNil(o.HealthcareService) {
		return true
	}

	return false
}

// SetHealthcareService gets a reference to the given []Reference and assigns it to the HealthcareService field.
func (o *OrganizationAffiliation) SetHealthcareService(v []Reference) {
	o.HealthcareService = v
}

// GetTelecom returns the Telecom field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetTelecom() []ContactPoint {
	if o == nil || IsNil(o.Telecom) {
		var ret []ContactPoint
		return ret
	}
	return o.Telecom
}

// GetTelecomOk returns a tuple with the Telecom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetTelecomOk() ([]ContactPoint, bool) {
	if o == nil || IsNil(o.Telecom) {
		return nil, false
	}
	return o.Telecom, true
}

// HasTelecom returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasTelecom() bool {
	if o != nil && !IsNil(o.Telecom) {
		return true
	}

	return false
}

// SetTelecom gets a reference to the given []ContactPoint and assigns it to the Telecom field.
func (o *OrganizationAffiliation) SetTelecom(v []ContactPoint) {
	o.Telecom = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *OrganizationAffiliation) GetEndpoint() []Reference {
	if o == nil || IsNil(o.Endpoint) {
		var ret []Reference
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAffiliation) GetEndpointOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *OrganizationAffiliation) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given []Reference and assigns it to the Endpoint field.
func (o *OrganizationAffiliation) SetEndpoint(v []Reference) {
	o.Endpoint = v
}

func (o OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAffiliation) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.ParticipatingOrganization) {
		toSerialize["participatingOrganization"] = o.ParticipatingOrganization
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Specialty) {
		toSerialize["specialty"] = o.Specialty
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.HealthcareService) {
		toSerialize["healthcareService"] = o.HealthcareService
	}
	if !IsNil(o.Telecom) {
		toSerialize["telecom"] = o.Telecom
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

func (o *OrganizationAffiliation) UnmarshalJSON(data []byte) (err error) {
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

	varOrganizationAffiliation := _OrganizationAffiliation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationAffiliation)

	if err != nil {
		return err
	}

	*o = OrganizationAffiliation(varOrganizationAffiliation)

	return err
}

type NullableOrganizationAffiliation struct {
	value *OrganizationAffiliation
	isSet bool
}

func (v NullableOrganizationAffiliation) Get() *OrganizationAffiliation {
	return v.value
}

func (v *NullableOrganizationAffiliation) Set(val *OrganizationAffiliation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAffiliation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAffiliation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAffiliation(val *OrganizationAffiliation) *NullableOrganizationAffiliation {
	return &NullableOrganizationAffiliation{value: val, isSet: true}
}

func (v NullableOrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAffiliation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


