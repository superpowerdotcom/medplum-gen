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

// checks if the AuditEventEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEventEntity{}

// AuditEventEntity A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEventEntity struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies a specific instance of the entity. The reference should be version specific.
	What *Reference `json:"what,omitempty"`
	// The type of the object that was involved in this audit event.
	Type *Coding `json:"type,omitempty"`
	// Code representing the role the entity played in the event being audited.
	Role *Coding `json:"role,omitempty"`
	// Identifier for the data life-cycle stage for the entity.
	Lifecycle *Coding `json:"lifecycle,omitempty"`
	// Security labels for the identified entity.
	SecurityLabel []Coding `json:"securityLabel,omitempty"`
	// A sequence of Unicode characters
	Name *string `json:"name,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// A stream of bytes
	Query *string `json:"query,omitempty"`
	// Tagged value pairs for conveying additional information about the entity.
	Detail []AuditEventDetail `json:"detail,omitempty"`
}

// NewAuditEventEntity instantiates a new AuditEventEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEventEntity() *AuditEventEntity {
	this := AuditEventEntity{}
	return &this
}

// NewAuditEventEntityWithDefaults instantiates a new AuditEventEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEventEntityWithDefaults() *AuditEventEntity {
	this := AuditEventEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditEventEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditEventEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditEventEntity) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AuditEventEntity) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AuditEventEntity) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *AuditEventEntity) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *AuditEventEntity) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *AuditEventEntity) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *AuditEventEntity) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetWhat returns the What field value if set, zero value otherwise.
func (o *AuditEventEntity) GetWhat() Reference {
	if o == nil || IsNil(o.What) {
		var ret Reference
		return ret
	}
	return *o.What
}

// GetWhatOk returns a tuple with the What field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetWhatOk() (*Reference, bool) {
	if o == nil || IsNil(o.What) {
		return nil, false
	}
	return o.What, true
}

// HasWhat returns a boolean if a field has been set.
func (o *AuditEventEntity) HasWhat() bool {
	if o != nil && !IsNil(o.What) {
		return true
	}

	return false
}

// SetWhat gets a reference to the given Reference and assigns it to the What field.
func (o *AuditEventEntity) SetWhat(v Reference) {
	o.What = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditEventEntity) GetType() Coding {
	if o == nil || IsNil(o.Type) {
		var ret Coding
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetTypeOk() (*Coding, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditEventEntity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Coding and assigns it to the Type field.
func (o *AuditEventEntity) SetType(v Coding) {
	o.Type = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AuditEventEntity) GetRole() Coding {
	if o == nil || IsNil(o.Role) {
		var ret Coding
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetRoleOk() (*Coding, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AuditEventEntity) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given Coding and assigns it to the Role field.
func (o *AuditEventEntity) SetRole(v Coding) {
	o.Role = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *AuditEventEntity) GetLifecycle() Coding {
	if o == nil || IsNil(o.Lifecycle) {
		var ret Coding
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetLifecycleOk() (*Coding, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *AuditEventEntity) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given Coding and assigns it to the Lifecycle field.
func (o *AuditEventEntity) SetLifecycle(v Coding) {
	o.Lifecycle = &v
}

// GetSecurityLabel returns the SecurityLabel field value if set, zero value otherwise.
func (o *AuditEventEntity) GetSecurityLabel() []Coding {
	if o == nil || IsNil(o.SecurityLabel) {
		var ret []Coding
		return ret
	}
	return o.SecurityLabel
}

// GetSecurityLabelOk returns a tuple with the SecurityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetSecurityLabelOk() ([]Coding, bool) {
	if o == nil || IsNil(o.SecurityLabel) {
		return nil, false
	}
	return o.SecurityLabel, true
}

// HasSecurityLabel returns a boolean if a field has been set.
func (o *AuditEventEntity) HasSecurityLabel() bool {
	if o != nil && !IsNil(o.SecurityLabel) {
		return true
	}

	return false
}

// SetSecurityLabel gets a reference to the given []Coding and assigns it to the SecurityLabel field.
func (o *AuditEventEntity) SetSecurityLabel(v []Coding) {
	o.SecurityLabel = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuditEventEntity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuditEventEntity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuditEventEntity) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuditEventEntity) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuditEventEntity) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuditEventEntity) SetDescription(v string) {
	o.Description = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *AuditEventEntity) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *AuditEventEntity) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *AuditEventEntity) SetQuery(v string) {
	o.Query = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AuditEventEntity) GetDetail() []AuditEventDetail {
	if o == nil || IsNil(o.Detail) {
		var ret []AuditEventDetail
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventEntity) GetDetailOk() ([]AuditEventDetail, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AuditEventEntity) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given []AuditEventDetail and assigns it to the Detail field.
func (o *AuditEventEntity) SetDetail(v []AuditEventDetail) {
	o.Detail = v
}

func (o AuditEventEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEventEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.What) {
		toSerialize["what"] = o.What
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Lifecycle) {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if !IsNil(o.SecurityLabel) {
		toSerialize["securityLabel"] = o.SecurityLabel
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableAuditEventEntity struct {
	value *AuditEventEntity
	isSet bool
}

func (v NullableAuditEventEntity) Get() *AuditEventEntity {
	return v.value
}

func (v *NullableAuditEventEntity) Set(val *AuditEventEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEventEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEventEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEventEntity(val *AuditEventEntity) *NullableAuditEventEntity {
	return &NullableAuditEventEntity{value: val, isSet: true}
}

func (v NullableAuditEventEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEventEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


