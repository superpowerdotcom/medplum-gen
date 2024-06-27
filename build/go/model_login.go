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

// checks if the Login type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Login{}

// Login Login event and session details.
type Login struct {
	// This is a Login resource
	ResourceType string `json:"resourceType"`
	// A sequence of Unicode characters
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
	Contained []string `json:"contained,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The client requesting the code.
	Client *Reference `json:"client,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	ProfileType *string `json:"profileType,omitempty"`
	// Optional required project for the login.
	Project *Reference `json:"project,omitempty"`
	// The user requesting the code.
	User Reference `json:"user"`
	// Reference to the project membership which includes FHIR identity (patient, practitioner, etc), access policy, and user configuration.
	Membership *Reference `json:"membership,omitempty"`
	// A sequence of Unicode characters
	Scope *string `json:"scope,omitempty"`
	// The authentication method used to obtain the code (password or google).
	AuthMethod string `json:"authMethod"`
	// An instant in time - known at least to the second
	AuthTime string `json:"authTime"`
	// A sequence of Unicode characters
	Cookie *string `json:"cookie,omitempty"`
	// A sequence of Unicode characters
	Code *string `json:"code,omitempty"`
	// A sequence of Unicode characters
	CodeChallenge *string `json:"codeChallenge,omitempty"`
	// OPTIONAL, defaults to \"plain\" if not present in the request.  Code verifier transformation method is \"S256\" or \"plain\".
	CodeChallengeMethod *string `json:"codeChallengeMethod,omitempty"`
	// A sequence of Unicode characters
	RefreshSecret *string `json:"refreshSecret,omitempty"`
	// A sequence of Unicode characters
	Nonce *string `json:"nonce,omitempty"`
	// Value of \"true\" or \"false\"
	MfaVerified *bool `json:"mfaVerified,omitempty"`
	// Value of \"true\" or \"false\"
	Granted *bool `json:"granted,omitempty"`
	// Value of \"true\" or \"false\"
	Revoked *bool `json:"revoked,omitempty"`
	// Value of \"true\" or \"false\"
	Admin *bool `json:"admin,omitempty"`
	// Value of \"true\" or \"false\"
	SuperAdmin *bool `json:"superAdmin,omitempty"`
	// Optional SMART App Launch context for this login.
	Launch *Reference `json:"launch,omitempty"`
	// A sequence of Unicode characters
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// A sequence of Unicode characters
	UserAgent *string `json:"userAgent,omitempty"`
}

type _Login Login

// NewLogin instantiates a new Login object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogin(resourceType string, user Reference, authMethod string, authTime string) *Login {
	this := Login{}
	this.ResourceType = resourceType
	this.User = user
	this.AuthMethod = authMethod
	this.AuthTime = authTime
	return &this
}

// NewLoginWithDefaults instantiates a new Login object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginWithDefaults() *Login {
	this := Login{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Login) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Login) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Login) SetResourceType(v string) {
	o.ResourceType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Login) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Login) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Login) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Login) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Login) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Login) SetMeta(v Meta) {
	o.Meta = &v
}

// GetImplicitRules returns the ImplicitRules field value if set, zero value otherwise.
func (o *Login) GetImplicitRules() string {
	if o == nil || IsNil(o.ImplicitRules) {
		var ret string
		return ret
	}
	return *o.ImplicitRules
}

// GetImplicitRulesOk returns a tuple with the ImplicitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetImplicitRulesOk() (*string, bool) {
	if o == nil || IsNil(o.ImplicitRules) {
		return nil, false
	}
	return o.ImplicitRules, true
}

// HasImplicitRules returns a boolean if a field has been set.
func (o *Login) HasImplicitRules() bool {
	if o != nil && !IsNil(o.ImplicitRules) {
		return true
	}

	return false
}

// SetImplicitRules gets a reference to the given string and assigns it to the ImplicitRules field.
func (o *Login) SetImplicitRules(v string) {
	o.ImplicitRules = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Login) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Login) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Login) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Login) GetText() Narrative {
	if o == nil || IsNil(o.Text) {
		var ret Narrative
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetTextOk() (*Narrative, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Login) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Narrative and assigns it to the Text field.
func (o *Login) SetText(v Narrative) {
	o.Text = &v
}

// GetContained returns the Contained field value if set, zero value otherwise.
func (o *Login) GetContained() []string {
	if o == nil || IsNil(o.Contained) {
		var ret []string
		return ret
	}
	return o.Contained
}

// GetContainedOk returns a tuple with the Contained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetContainedOk() ([]string, bool) {
	if o == nil || IsNil(o.Contained) {
		return nil, false
	}
	return o.Contained, true
}

// HasContained returns a boolean if a field has been set.
func (o *Login) HasContained() bool {
	if o != nil && !IsNil(o.Contained) {
		return true
	}

	return false
}

// SetContained gets a reference to the given []string and assigns it to the Contained field.
func (o *Login) SetContained(v []string) {
	o.Contained = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Login) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Login) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *Login) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *Login) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *Login) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *Login) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *Login) GetClient() Reference {
	if o == nil || IsNil(o.Client) {
		var ret Reference
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetClientOk() (*Reference, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *Login) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given Reference and assigns it to the Client field.
func (o *Login) SetClient(v Reference) {
	o.Client = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *Login) GetProfileType() string {
	if o == nil || IsNil(o.ProfileType) {
		var ret string
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetProfileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *Login) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given string and assigns it to the ProfileType field.
func (o *Login) SetProfileType(v string) {
	o.ProfileType = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Login) GetProject() Reference {
	if o == nil || IsNil(o.Project) {
		var ret Reference
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetProjectOk() (*Reference, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Login) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Reference and assigns it to the Project field.
func (o *Login) SetProject(v Reference) {
	o.Project = &v
}

// GetUser returns the User field value
func (o *Login) GetUser() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Login) GetUserOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Login) SetUser(v Reference) {
	o.User = v
}

// GetMembership returns the Membership field value if set, zero value otherwise.
func (o *Login) GetMembership() Reference {
	if o == nil || IsNil(o.Membership) {
		var ret Reference
		return ret
	}
	return *o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetMembershipOk() (*Reference, bool) {
	if o == nil || IsNil(o.Membership) {
		return nil, false
	}
	return o.Membership, true
}

// HasMembership returns a boolean if a field has been set.
func (o *Login) HasMembership() bool {
	if o != nil && !IsNil(o.Membership) {
		return true
	}

	return false
}

// SetMembership gets a reference to the given Reference and assigns it to the Membership field.
func (o *Login) SetMembership(v Reference) {
	o.Membership = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Login) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Login) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Login) SetScope(v string) {
	o.Scope = &v
}

// GetAuthMethod returns the AuthMethod field value
func (o *Login) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *Login) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value
func (o *Login) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetAuthTime returns the AuthTime field value
func (o *Login) GetAuthTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value
// and a boolean to check if the value has been set.
func (o *Login) GetAuthTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthTime, true
}

// SetAuthTime sets field value
func (o *Login) SetAuthTime(v string) {
	o.AuthTime = v
}

// GetCookie returns the Cookie field value if set, zero value otherwise.
func (o *Login) GetCookie() string {
	if o == nil || IsNil(o.Cookie) {
		var ret string
		return ret
	}
	return *o.Cookie
}

// GetCookieOk returns a tuple with the Cookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetCookieOk() (*string, bool) {
	if o == nil || IsNil(o.Cookie) {
		return nil, false
	}
	return o.Cookie, true
}

// HasCookie returns a boolean if a field has been set.
func (o *Login) HasCookie() bool {
	if o != nil && !IsNil(o.Cookie) {
		return true
	}

	return false
}

// SetCookie gets a reference to the given string and assigns it to the Cookie field.
func (o *Login) SetCookie(v string) {
	o.Cookie = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Login) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Login) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Login) SetCode(v string) {
	o.Code = &v
}

// GetCodeChallenge returns the CodeChallenge field value if set, zero value otherwise.
func (o *Login) GetCodeChallenge() string {
	if o == nil || IsNil(o.CodeChallenge) {
		var ret string
		return ret
	}
	return *o.CodeChallenge
}

// GetCodeChallengeOk returns a tuple with the CodeChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetCodeChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.CodeChallenge) {
		return nil, false
	}
	return o.CodeChallenge, true
}

// HasCodeChallenge returns a boolean if a field has been set.
func (o *Login) HasCodeChallenge() bool {
	if o != nil && !IsNil(o.CodeChallenge) {
		return true
	}

	return false
}

// SetCodeChallenge gets a reference to the given string and assigns it to the CodeChallenge field.
func (o *Login) SetCodeChallenge(v string) {
	o.CodeChallenge = &v
}

// GetCodeChallengeMethod returns the CodeChallengeMethod field value if set, zero value otherwise.
func (o *Login) GetCodeChallengeMethod() string {
	if o == nil || IsNil(o.CodeChallengeMethod) {
		var ret string
		return ret
	}
	return *o.CodeChallengeMethod
}

// GetCodeChallengeMethodOk returns a tuple with the CodeChallengeMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetCodeChallengeMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CodeChallengeMethod) {
		return nil, false
	}
	return o.CodeChallengeMethod, true
}

// HasCodeChallengeMethod returns a boolean if a field has been set.
func (o *Login) HasCodeChallengeMethod() bool {
	if o != nil && !IsNil(o.CodeChallengeMethod) {
		return true
	}

	return false
}

// SetCodeChallengeMethod gets a reference to the given string and assigns it to the CodeChallengeMethod field.
func (o *Login) SetCodeChallengeMethod(v string) {
	o.CodeChallengeMethod = &v
}

// GetRefreshSecret returns the RefreshSecret field value if set, zero value otherwise.
func (o *Login) GetRefreshSecret() string {
	if o == nil || IsNil(o.RefreshSecret) {
		var ret string
		return ret
	}
	return *o.RefreshSecret
}

// GetRefreshSecretOk returns a tuple with the RefreshSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetRefreshSecretOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshSecret) {
		return nil, false
	}
	return o.RefreshSecret, true
}

// HasRefreshSecret returns a boolean if a field has been set.
func (o *Login) HasRefreshSecret() bool {
	if o != nil && !IsNil(o.RefreshSecret) {
		return true
	}

	return false
}

// SetRefreshSecret gets a reference to the given string and assigns it to the RefreshSecret field.
func (o *Login) SetRefreshSecret(v string) {
	o.RefreshSecret = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Login) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Login) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Login) SetNonce(v string) {
	o.Nonce = &v
}

// GetMfaVerified returns the MfaVerified field value if set, zero value otherwise.
func (o *Login) GetMfaVerified() bool {
	if o == nil || IsNil(o.MfaVerified) {
		var ret bool
		return ret
	}
	return *o.MfaVerified
}

// GetMfaVerifiedOk returns a tuple with the MfaVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetMfaVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaVerified) {
		return nil, false
	}
	return o.MfaVerified, true
}

// HasMfaVerified returns a boolean if a field has been set.
func (o *Login) HasMfaVerified() bool {
	if o != nil && !IsNil(o.MfaVerified) {
		return true
	}

	return false
}

// SetMfaVerified gets a reference to the given bool and assigns it to the MfaVerified field.
func (o *Login) SetMfaVerified(v bool) {
	o.MfaVerified = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise.
func (o *Login) GetGranted() bool {
	if o == nil || IsNil(o.Granted) {
		var ret bool
		return ret
	}
	return *o.Granted
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetGrantedOk() (*bool, bool) {
	if o == nil || IsNil(o.Granted) {
		return nil, false
	}
	return o.Granted, true
}

// HasGranted returns a boolean if a field has been set.
func (o *Login) HasGranted() bool {
	if o != nil && !IsNil(o.Granted) {
		return true
	}

	return false
}

// SetGranted gets a reference to the given bool and assigns it to the Granted field.
func (o *Login) SetGranted(v bool) {
	o.Granted = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *Login) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *Login) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *Login) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *Login) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *Login) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *Login) SetAdmin(v bool) {
	o.Admin = &v
}

// GetSuperAdmin returns the SuperAdmin field value if set, zero value otherwise.
func (o *Login) GetSuperAdmin() bool {
	if o == nil || IsNil(o.SuperAdmin) {
		var ret bool
		return ret
	}
	return *o.SuperAdmin
}

// GetSuperAdminOk returns a tuple with the SuperAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetSuperAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.SuperAdmin) {
		return nil, false
	}
	return o.SuperAdmin, true
}

// HasSuperAdmin returns a boolean if a field has been set.
func (o *Login) HasSuperAdmin() bool {
	if o != nil && !IsNil(o.SuperAdmin) {
		return true
	}

	return false
}

// SetSuperAdmin gets a reference to the given bool and assigns it to the SuperAdmin field.
func (o *Login) SetSuperAdmin(v bool) {
	o.SuperAdmin = &v
}

// GetLaunch returns the Launch field value if set, zero value otherwise.
func (o *Login) GetLaunch() Reference {
	if o == nil || IsNil(o.Launch) {
		var ret Reference
		return ret
	}
	return *o.Launch
}

// GetLaunchOk returns a tuple with the Launch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetLaunchOk() (*Reference, bool) {
	if o == nil || IsNil(o.Launch) {
		return nil, false
	}
	return o.Launch, true
}

// HasLaunch returns a boolean if a field has been set.
func (o *Login) HasLaunch() bool {
	if o != nil && !IsNil(o.Launch) {
		return true
	}

	return false
}

// SetLaunch gets a reference to the given Reference and assigns it to the Launch field.
func (o *Login) SetLaunch(v Reference) {
	o.Launch = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *Login) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *Login) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *Login) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *Login) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Login) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *Login) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *Login) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o Login) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Login) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["user"] = o.User
	if !IsNil(o.Membership) {
		toSerialize["membership"] = o.Membership
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["authMethod"] = o.AuthMethod
	toSerialize["authTime"] = o.AuthTime
	if !IsNil(o.Cookie) {
		toSerialize["cookie"] = o.Cookie
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CodeChallenge) {
		toSerialize["codeChallenge"] = o.CodeChallenge
	}
	if !IsNil(o.CodeChallengeMethod) {
		toSerialize["codeChallengeMethod"] = o.CodeChallengeMethod
	}
	if !IsNil(o.RefreshSecret) {
		toSerialize["refreshSecret"] = o.RefreshSecret
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.MfaVerified) {
		toSerialize["mfaVerified"] = o.MfaVerified
	}
	if !IsNil(o.Granted) {
		toSerialize["granted"] = o.Granted
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.SuperAdmin) {
		toSerialize["superAdmin"] = o.SuperAdmin
	}
	if !IsNil(o.Launch) {
		toSerialize["launch"] = o.Launch
	}
	if !IsNil(o.RemoteAddress) {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	return toSerialize, nil
}

func (o *Login) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"user",
		"authMethod",
		"authTime",
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

	varLogin := _Login{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogin)

	if err != nil {
		return err
	}

	*o = Login(varLogin)

	return err
}

type NullableLogin struct {
	value *Login
	isSet bool
}

func (v NullableLogin) Get() *Login {
	return v.value
}

func (v *NullableLogin) Set(val *Login) {
	v.value = val
	v.isSet = true
}

func (v NullableLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogin(val *Login) *NullableLogin {
	return &NullableLogin{value: val, isSet: true}
}

func (v NullableLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


