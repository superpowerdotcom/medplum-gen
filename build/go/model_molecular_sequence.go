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

// checks if the MolecularSequence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MolecularSequence{}

// MolecularSequence Raw data describing a biological sequence.
type MolecularSequence struct {
	// This is a MolecularSequence resource
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
	// A unique identifier for this particular sequence instance. This is a FHIR-defined id.
	Identifier []Identifier `json:"identifier,omitempty"`
	// Amino Acid Sequence/ DNA Sequence / RNA Sequence.
	Type *string `json:"type,omitempty"`
	// A whole number
	CoordinateSystem *float32 `json:"coordinateSystem,omitempty"`
	// The patient whose sequencing results are described by this resource.
	Patient *Reference `json:"patient,omitempty"`
	// Specimen used for sequencing.
	Specimen *Reference `json:"specimen,omitempty"`
	// The method for sequencing, for example, chip information.
	Device *Reference `json:"device,omitempty"`
	// The organization or lab that should be responsible for this result.
	Performer *Reference `json:"performer,omitempty"`
	// The number of copies of the sequence of interest. (RNASeq).
	Quantity *Quantity `json:"quantity,omitempty"`
	// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
	ReferenceSeq *MolecularSequenceReferenceSeq `json:"referenceSeq,omitempty"`
	// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
	Variant []MolecularSequenceVariant `json:"variant,omitempty"`
	// A sequence of Unicode characters
	ObservedSeq *string `json:"observedSeq,omitempty"`
	// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
	Quality []MolecularSequenceQuality `json:"quality,omitempty"`
	// A whole number
	ReadCoverage *float32 `json:"readCoverage,omitempty"`
	// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
	Repository []MolecularSequenceRepository `json:"repository,omitempty"`
	// Pointer to next atomic sequence which at most contains one variant.
	Pointer []Reference `json:"pointer,omitempty"`
	// Information about chromosome structure variation.
	StructureVariant []MolecularSequenceStructureVariant `json:"structureVariant,omitempty"`
}

type _MolecularSequence MolecularSequence

// NewMolecularSequence instantiates a new MolecularSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMolecularSequence(resourceType string) *MolecularSequence {
	this := MolecularSequence{}
	this.ResourceType = resourceType
	return &this
}

// NewMolecularSequenceWithDefaults instantiates a new MolecularSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMolecularSequenceWithDefaults() *MolecularSequence {
	this := MolecularSequence{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *MolecularSequence) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *MolecularSequence) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MolecularSequence) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MolecularSequence) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MolecularSequence) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MolecularSequence) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MolecularSequence) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *MolecularSequence) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *MolecularSequence) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *MolecularSequence) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *MolecularSequence) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MolecularSequence) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MolecularSequence) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MolecularSequence) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MolecularSequence) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MolecularSequence) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *MolecularSequence) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *MolecularSequence) GetContained() []ResourceList {
	if o == nil || IsNil(o.Contained) {
		var ret []ResourceList
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetContainedOk() ([]ResourceList, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *MolecularSequence) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []ResourceList and assigns it to the Contained field.
func (o *MolecularSequence) SetContained(v []ResourceList) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *MolecularSequence) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *MolecularSequence) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *MolecularSequence) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *MolecularSequence) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *MolecularSequence) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *MolecularSequence) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MolecularSequence) GetIdentifier() []Identifier {
	if o == nil || IsNil(o.Identifier) {
		var ret []Identifier
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetIdentifierOk() ([]Identifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MolecularSequence) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given []Identifier and assigns it to the Identifier field.
func (o *MolecularSequence) SetIdentifier(v []Identifier) {
	o.Identifier = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MolecularSequence) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MolecularSequence) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MolecularSequence) SetType(v string) {
	o.Type = &v
}

// GetCoordinateSystem returns the CoordinateSystem field value if set, zero value otherwise.
func (o *MolecularSequence) GetCoordinateSystem() float32 {
	if o == nil || IsNil(o.CoordinateSystem) {
		var ret float32
		return ret
	}
	return *o.CoordinateSystem
}

// GetCoordinateSystemOk returns a tuple with the CoordinateSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetCoordinateSystemOk() (*float32, bool) {
	if o == nil || IsNil(o.CoordinateSystem) {
		return nil, false
	}
	return o.CoordinateSystem, true
}

// HasCoordinateSystem returns a boolean if a field has been set.
func (o *MolecularSequence) HasCoordinateSystem() bool {
	if o != nil && !IsNil(o.CoordinateSystem) {
		return true
	}

	return false
}

// SetCoordinateSystem gets a reference to the given float32 and assigns it to the CoordinateSystem field.
func (o *MolecularSequence) SetCoordinateSystem(v float32) {
	o.CoordinateSystem = &v
}

// GetPatient returns the Patient field value if set, zero value otherwise.
func (o *MolecularSequence) GetPatient() Reference {
	if o == nil || IsNil(o.Patient) {
		var ret Reference
		return ret
	}
	return *o.Patient
}

// GetPatientOk returns a tuple with the Patient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetPatientOk() (*Reference, bool) {
	if o == nil || IsNil(o.Patient) {
		return nil, false
	}
	return o.Patient, true
}

// HasPatient returns a boolean if a field has been set.
func (o *MolecularSequence) HasPatient() bool {
	if o != nil && !IsNil(o.Patient) {
		return true
	}

	return false
}

// SetPatient gets a reference to the given Reference and assigns it to the Patient field.
func (o *MolecularSequence) SetPatient(v Reference) {
	o.Patient = &v
}

// GetSpecimen returns the Specimen field value if set, zero value otherwise.
func (o *MolecularSequence) GetSpecimen() Reference {
	if o == nil || IsNil(o.Specimen) {
		var ret Reference
		return ret
	}
	return *o.Specimen
}

// GetSpecimenOk returns a tuple with the Specimen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetSpecimenOk() (*Reference, bool) {
	if o == nil || IsNil(o.Specimen) {
		return nil, false
	}
	return o.Specimen, true
}

// HasSpecimen returns a boolean if a field has been set.
func (o *MolecularSequence) HasSpecimen() bool {
	if o != nil && !IsNil(o.Specimen) {
		return true
	}

	return false
}

// SetSpecimen gets a reference to the given Reference and assigns it to the Specimen field.
func (o *MolecularSequence) SetSpecimen(v Reference) {
	o.Specimen = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *MolecularSequence) GetDevice() Reference {
	if o == nil || IsNil(o.Device) {
		var ret Reference
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetDeviceOk() (*Reference, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *MolecularSequence) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Reference and assigns it to the Device field.
func (o *MolecularSequence) SetDevice(v Reference) {
	o.Device = &v
}

// GetPerformer returns the Performer field value if set, zero value otherwise.
func (o *MolecularSequence) GetPerformer() Reference {
	if o == nil || IsNil(o.Performer) {
		var ret Reference
		return ret
	}
	return *o.Performer
}

// GetPerformerOk returns a tuple with the Performer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetPerformerOk() (*Reference, bool) {
	if o == nil || IsNil(o.Performer) {
		return nil, false
	}
	return o.Performer, true
}

// HasPerformer returns a boolean if a field has been set.
func (o *MolecularSequence) HasPerformer() bool {
	if o != nil && !IsNil(o.Performer) {
		return true
	}

	return false
}

// SetPerformer gets a reference to the given Reference and assigns it to the Performer field.
func (o *MolecularSequence) SetPerformer(v Reference) {
	o.Performer = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MolecularSequence) GetQuantity() Quantity {
	if o == nil || IsNil(o.Quantity) {
		var ret Quantity
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetQuantityOk() (*Quantity, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MolecularSequence) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given Quantity and assigns it to the Quantity field.
func (o *MolecularSequence) SetQuantity(v Quantity) {
	o.Quantity = &v
}

// GetReferenceSeq returns the ReferenceSeq field value if set, zero value otherwise.
func (o *MolecularSequence) GetReferenceSeq() MolecularSequenceReferenceSeq {
	if o == nil || IsNil(o.ReferenceSeq) {
		var ret MolecularSequenceReferenceSeq
		return ret
	}
	return *o.ReferenceSeq
}

// GetReferenceSeqOk returns a tuple with the ReferenceSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetReferenceSeqOk() (*MolecularSequenceReferenceSeq, bool) {
	if o == nil || IsNil(o.ReferenceSeq) {
		return nil, false
	}
	return o.ReferenceSeq, true
}

// HasReferenceSeq returns a boolean if a field has been set.
func (o *MolecularSequence) HasReferenceSeq() bool {
	if o != nil && !IsNil(o.ReferenceSeq) {
		return true
	}

	return false
}

// SetReferenceSeq gets a reference to the given MolecularSequenceReferenceSeq and assigns it to the ReferenceSeq field.
func (o *MolecularSequence) SetReferenceSeq(v MolecularSequenceReferenceSeq) {
	o.ReferenceSeq = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *MolecularSequence) GetVariant() []MolecularSequenceVariant {
	if o == nil || IsNil(o.Variant) {
		var ret []MolecularSequenceVariant
		return ret
	}
	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetVariantOk() ([]MolecularSequenceVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *MolecularSequence) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given []MolecularSequenceVariant and assigns it to the Variant field.
func (o *MolecularSequence) SetVariant(v []MolecularSequenceVariant) {
	o.Variant = v
}

// GetObservedSeq returns the ObservedSeq field value if set, zero value otherwise.
func (o *MolecularSequence) GetObservedSeq() string {
	if o == nil || IsNil(o.ObservedSeq) {
		var ret string
		return ret
	}
	return *o.ObservedSeq
}

// GetObservedSeqOk returns a tuple with the ObservedSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetObservedSeqOk() (*string, bool) {
	if o == nil || IsNil(o.ObservedSeq) {
		return nil, false
	}
	return o.ObservedSeq, true
}

// HasObservedSeq returns a boolean if a field has been set.
func (o *MolecularSequence) HasObservedSeq() bool {
	if o != nil && !IsNil(o.ObservedSeq) {
		return true
	}

	return false
}

// SetObservedSeq gets a reference to the given string and assigns it to the ObservedSeq field.
func (o *MolecularSequence) SetObservedSeq(v string) {
	o.ObservedSeq = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *MolecularSequence) GetQuality() []MolecularSequenceQuality {
	if o == nil || IsNil(o.Quality) {
		var ret []MolecularSequenceQuality
		return ret
	}
	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetQualityOk() ([]MolecularSequenceQuality, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *MolecularSequence) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given []MolecularSequenceQuality and assigns it to the Quality field.
func (o *MolecularSequence) SetQuality(v []MolecularSequenceQuality) {
	o.Quality = v
}

// GetReadCoverage returns the ReadCoverage field value if set, zero value otherwise.
func (o *MolecularSequence) GetReadCoverage() float32 {
	if o == nil || IsNil(o.ReadCoverage) {
		var ret float32
		return ret
	}
	return *o.ReadCoverage
}

// GetReadCoverageOk returns a tuple with the ReadCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetReadCoverageOk() (*float32, bool) {
	if o == nil || IsNil(o.ReadCoverage) {
		return nil, false
	}
	return o.ReadCoverage, true
}

// HasReadCoverage returns a boolean if a field has been set.
func (o *MolecularSequence) HasReadCoverage() bool {
	if o != nil && !IsNil(o.ReadCoverage) {
		return true
	}

	return false
}

// SetReadCoverage gets a reference to the given float32 and assigns it to the ReadCoverage field.
func (o *MolecularSequence) SetReadCoverage(v float32) {
	o.ReadCoverage = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *MolecularSequence) GetRepository() []MolecularSequenceRepository {
	if o == nil || IsNil(o.Repository) {
		var ret []MolecularSequenceRepository
		return ret
	}
	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetRepositoryOk() ([]MolecularSequenceRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *MolecularSequence) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given []MolecularSequenceRepository and assigns it to the Repository field.
func (o *MolecularSequence) SetRepository(v []MolecularSequenceRepository) {
	o.Repository = v
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *MolecularSequence) GetPointer() []Reference {
	if o == nil || IsNil(o.Pointer) {
		var ret []Reference
		return ret
	}
	return o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetPointerOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Pointer) {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *MolecularSequence) HasPointer() bool {
	if o != nil && !IsNil(o.Pointer) {
		return true
	}

	return false
}

// SetPointer gets a reference to the given []Reference and assigns it to the Pointer field.
func (o *MolecularSequence) SetPointer(v []Reference) {
	o.Pointer = v
}

// GetStructureVariant returns the StructureVariant field value if set, zero value otherwise.
func (o *MolecularSequence) GetStructureVariant() []MolecularSequenceStructureVariant {
	if o == nil || IsNil(o.StructureVariant) {
		var ret []MolecularSequenceStructureVariant
		return ret
	}
	return o.StructureVariant
}

// GetStructureVariantOk returns a tuple with the StructureVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MolecularSequence) GetStructureVariantOk() ([]MolecularSequenceStructureVariant, bool) {
	if o == nil || IsNil(o.StructureVariant) {
		return nil, false
	}
	return o.StructureVariant, true
}

// HasStructureVariant returns a boolean if a field has been set.
func (o *MolecularSequence) HasStructureVariant() bool {
	if o != nil && !IsNil(o.StructureVariant) {
		return true
	}

	return false
}

// SetStructureVariant gets a reference to the given []MolecularSequenceStructureVariant and assigns it to the StructureVariant field.
func (o *MolecularSequence) SetStructureVariant(v []MolecularSequenceStructureVariant) {
	o.StructureVariant = v
}

func (o MolecularSequence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MolecularSequence) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CoordinateSystem) {
		toSerialize["coordinateSystem"] = o.CoordinateSystem
	}
	if !IsNil(o.Patient) {
		toSerialize["patient"] = o.Patient
	}
	if !IsNil(o.Specimen) {
		toSerialize["specimen"] = o.Specimen
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Performer) {
		toSerialize["performer"] = o.Performer
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ReferenceSeq) {
		toSerialize["referenceSeq"] = o.ReferenceSeq
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.ObservedSeq) {
		toSerialize["observedSeq"] = o.ObservedSeq
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !IsNil(o.ReadCoverage) {
		toSerialize["readCoverage"] = o.ReadCoverage
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Pointer) {
		toSerialize["pointer"] = o.Pointer
	}
	if !IsNil(o.StructureVariant) {
		toSerialize["structureVariant"] = o.StructureVariant
	}
	return toSerialize, nil
}

func (o *MolecularSequence) UnmarshalJSON(data []byte) (err error) {
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

	varMolecularSequence := _MolecularSequence{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMolecularSequence)

	if err != nil {
		return err
	}

	*o = MolecularSequence(varMolecularSequence)

	return err
}

type NullableMolecularSequence struct {
	value *MolecularSequence
	isSet bool
}

func (v NullableMolecularSequence) Get() *MolecularSequence {
	return v.value
}

func (v *NullableMolecularSequence) Set(val *MolecularSequence) {
	v.value = val
	v.isSet = true
}

func (v NullableMolecularSequence) IsSet() bool {
	return v.isSet
}

func (v *NullableMolecularSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMolecularSequence(val *MolecularSequence) *NullableMolecularSequence {
	return &NullableMolecularSequence{value: val, isSet: true}
}

func (v NullableMolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMolecularSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


