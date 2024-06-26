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

// checks if the TestScriptAssert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestScriptAssert{}

// TestScriptAssert A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScriptAssert struct {
	// A sequence of Unicode characters
	Id *string `json:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension `json:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.  Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A sequence of Unicode characters
	Label *string `json:"label,omitempty"`
	// A sequence of Unicode characters
	Description *string `json:"description,omitempty"`
	// The direction to use for the assertion.
	Direction *string `json:"direction,omitempty"`
	// A sequence of Unicode characters
	CompareToSourceId *string `json:"compareToSourceId,omitempty"`
	// A sequence of Unicode characters
	CompareToSourceExpression *string `json:"compareToSourceExpression,omitempty"`
	// A sequence of Unicode characters
	CompareToSourcePath *string `json:"compareToSourcePath,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	ContentType *string `json:"contentType,omitempty"`
	// A sequence of Unicode characters
	Expression *string `json:"expression,omitempty"`
	// A sequence of Unicode characters
	HeaderField *string `json:"headerField,omitempty"`
	// A sequence of Unicode characters
	MinimumId *string `json:"minimumId,omitempty"`
	// Value of \"true\" or \"false\"
	NavigationLinks *bool `json:"navigationLinks,omitempty"`
	// The operator type defines the conditional behavior of the assert. If not defined, the default is equals.
	Operator *string `json:"operator,omitempty"`
	// A sequence of Unicode characters
	Path *string `json:"path,omitempty"`
	// The request method or HTTP operation code to compare against that used by the client system under test.
	RequestMethod *string `json:"requestMethod,omitempty"`
	// A sequence of Unicode characters
	RequestURL *string `json:"requestURL,omitempty"`
	// A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
	Resource *string `json:"resource,omitempty"`
	// okay | created | noContent | notModified | bad | forbidden | notFound | methodNotAllowed | conflict | gone | preconditionFailed | unprocessable.
	Response *string `json:"response,omitempty"`
	// A sequence of Unicode characters
	ResponseCode *string `json:"responseCode,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	SourceId *string `json:"sourceId,omitempty"`
	// Any combination of letters, numerals, \"-\" and \".\", with a length limit of 64 characters.  (This might be an integer, an unprefixed OID, UUID or any other identifier pattern that meets these constraints.)  Ids are case-insensitive.
	ValidateProfileId *string `json:"validateProfileId,omitempty"`
	// A sequence of Unicode characters
	Value *string `json:"value,omitempty"`
	// Value of \"true\" or \"false\"
	WarningOnly *bool `json:"warningOnly,omitempty"`
}

// NewTestScriptAssert instantiates a new TestScriptAssert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestScriptAssert() *TestScriptAssert {
	this := TestScriptAssert{}
	return &this
}

// NewTestScriptAssertWithDefaults instantiates a new TestScriptAssert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestScriptAssertWithDefaults() *TestScriptAssert {
	this := TestScriptAssert{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestScriptAssert) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestScriptAssert) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestScriptAssert) SetId(v string) {
	o.Id = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *TestScriptAssert) GetExtension() []Extension {
	if o == nil || IsNil(o.Extension) {
		var ret []Extension
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *TestScriptAssert) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given []Extension and assigns it to the Extension field.
func (o *TestScriptAssert) SetExtension(v []Extension) {
	o.Extension = v
}

// GetModifierExtension returns the ModifierExtension field value if set, zero value otherwise.
func (o *TestScriptAssert) GetModifierExtension() []Extension {
	if o == nil || IsNil(o.ModifierExtension) {
		var ret []Extension
		return ret
	}
	return o.ModifierExtension
}

// GetModifierExtensionOk returns a tuple with the ModifierExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetModifierExtensionOk() ([]Extension, bool) {
	if o == nil || IsNil(o.ModifierExtension) {
		return nil, false
	}
	return o.ModifierExtension, true
}

// HasModifierExtension returns a boolean if a field has been set.
func (o *TestScriptAssert) HasModifierExtension() bool {
	if o != nil && !IsNil(o.ModifierExtension) {
		return true
	}

	return false
}

// SetModifierExtension gets a reference to the given []Extension and assigns it to the ModifierExtension field.
func (o *TestScriptAssert) SetModifierExtension(v []Extension) {
	o.ModifierExtension = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TestScriptAssert) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TestScriptAssert) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *TestScriptAssert) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TestScriptAssert) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TestScriptAssert) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TestScriptAssert) SetDescription(v string) {
	o.Description = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *TestScriptAssert) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *TestScriptAssert) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *TestScriptAssert) SetDirection(v string) {
	o.Direction = &v
}

// GetCompareToSourceId returns the CompareToSourceId field value if set, zero value otherwise.
func (o *TestScriptAssert) GetCompareToSourceId() string {
	if o == nil || IsNil(o.CompareToSourceId) {
		var ret string
		return ret
	}
	return *o.CompareToSourceId
}

// GetCompareToSourceIdOk returns a tuple with the CompareToSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetCompareToSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompareToSourceId) {
		return nil, false
	}
	return o.CompareToSourceId, true
}

// HasCompareToSourceId returns a boolean if a field has been set.
func (o *TestScriptAssert) HasCompareToSourceId() bool {
	if o != nil && !IsNil(o.CompareToSourceId) {
		return true
	}

	return false
}

// SetCompareToSourceId gets a reference to the given string and assigns it to the CompareToSourceId field.
func (o *TestScriptAssert) SetCompareToSourceId(v string) {
	o.CompareToSourceId = &v
}

// GetCompareToSourceExpression returns the CompareToSourceExpression field value if set, zero value otherwise.
func (o *TestScriptAssert) GetCompareToSourceExpression() string {
	if o == nil || IsNil(o.CompareToSourceExpression) {
		var ret string
		return ret
	}
	return *o.CompareToSourceExpression
}

// GetCompareToSourceExpressionOk returns a tuple with the CompareToSourceExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetCompareToSourceExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CompareToSourceExpression) {
		return nil, false
	}
	return o.CompareToSourceExpression, true
}

// HasCompareToSourceExpression returns a boolean if a field has been set.
func (o *TestScriptAssert) HasCompareToSourceExpression() bool {
	if o != nil && !IsNil(o.CompareToSourceExpression) {
		return true
	}

	return false
}

// SetCompareToSourceExpression gets a reference to the given string and assigns it to the CompareToSourceExpression field.
func (o *TestScriptAssert) SetCompareToSourceExpression(v string) {
	o.CompareToSourceExpression = &v
}

// GetCompareToSourcePath returns the CompareToSourcePath field value if set, zero value otherwise.
func (o *TestScriptAssert) GetCompareToSourcePath() string {
	if o == nil || IsNil(o.CompareToSourcePath) {
		var ret string
		return ret
	}
	return *o.CompareToSourcePath
}

// GetCompareToSourcePathOk returns a tuple with the CompareToSourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetCompareToSourcePathOk() (*string, bool) {
	if o == nil || IsNil(o.CompareToSourcePath) {
		return nil, false
	}
	return o.CompareToSourcePath, true
}

// HasCompareToSourcePath returns a boolean if a field has been set.
func (o *TestScriptAssert) HasCompareToSourcePath() bool {
	if o != nil && !IsNil(o.CompareToSourcePath) {
		return true
	}

	return false
}

// SetCompareToSourcePath gets a reference to the given string and assigns it to the CompareToSourcePath field.
func (o *TestScriptAssert) SetCompareToSourcePath(v string) {
	o.CompareToSourcePath = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TestScriptAssert) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TestScriptAssert) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TestScriptAssert) SetContentType(v string) {
	o.ContentType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *TestScriptAssert) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *TestScriptAssert) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *TestScriptAssert) SetExpression(v string) {
	o.Expression = &v
}

// GetHeaderField returns the HeaderField field value if set, zero value otherwise.
func (o *TestScriptAssert) GetHeaderField() string {
	if o == nil || IsNil(o.HeaderField) {
		var ret string
		return ret
	}
	return *o.HeaderField
}

// GetHeaderFieldOk returns a tuple with the HeaderField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetHeaderFieldOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderField) {
		return nil, false
	}
	return o.HeaderField, true
}

// HasHeaderField returns a boolean if a field has been set.
func (o *TestScriptAssert) HasHeaderField() bool {
	if o != nil && !IsNil(o.HeaderField) {
		return true
	}

	return false
}

// SetHeaderField gets a reference to the given string and assigns it to the HeaderField field.
func (o *TestScriptAssert) SetHeaderField(v string) {
	o.HeaderField = &v
}

// GetMinimumId returns the MinimumId field value if set, zero value otherwise.
func (o *TestScriptAssert) GetMinimumId() string {
	if o == nil || IsNil(o.MinimumId) {
		var ret string
		return ret
	}
	return *o.MinimumId
}

// GetMinimumIdOk returns a tuple with the MinimumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetMinimumIdOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumId) {
		return nil, false
	}
	return o.MinimumId, true
}

// HasMinimumId returns a boolean if a field has been set.
func (o *TestScriptAssert) HasMinimumId() bool {
	if o != nil && !IsNil(o.MinimumId) {
		return true
	}

	return false
}

// SetMinimumId gets a reference to the given string and assigns it to the MinimumId field.
func (o *TestScriptAssert) SetMinimumId(v string) {
	o.MinimumId = &v
}

// GetNavigationLinks returns the NavigationLinks field value if set, zero value otherwise.
func (o *TestScriptAssert) GetNavigationLinks() bool {
	if o == nil || IsNil(o.NavigationLinks) {
		var ret bool
		return ret
	}
	return *o.NavigationLinks
}

// GetNavigationLinksOk returns a tuple with the NavigationLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetNavigationLinksOk() (*bool, bool) {
	if o == nil || IsNil(o.NavigationLinks) {
		return nil, false
	}
	return o.NavigationLinks, true
}

// HasNavigationLinks returns a boolean if a field has been set.
func (o *TestScriptAssert) HasNavigationLinks() bool {
	if o != nil && !IsNil(o.NavigationLinks) {
		return true
	}

	return false
}

// SetNavigationLinks gets a reference to the given bool and assigns it to the NavigationLinks field.
func (o *TestScriptAssert) SetNavigationLinks(v bool) {
	o.NavigationLinks = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *TestScriptAssert) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *TestScriptAssert) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *TestScriptAssert) SetOperator(v string) {
	o.Operator = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TestScriptAssert) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TestScriptAssert) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TestScriptAssert) SetPath(v string) {
	o.Path = &v
}

// GetRequestMethod returns the RequestMethod field value if set, zero value otherwise.
func (o *TestScriptAssert) GetRequestMethod() string {
	if o == nil || IsNil(o.RequestMethod) {
		var ret string
		return ret
	}
	return *o.RequestMethod
}

// GetRequestMethodOk returns a tuple with the RequestMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetRequestMethodOk() (*string, bool) {
	if o == nil || IsNil(o.RequestMethod) {
		return nil, false
	}
	return o.RequestMethod, true
}

// HasRequestMethod returns a boolean if a field has been set.
func (o *TestScriptAssert) HasRequestMethod() bool {
	if o != nil && !IsNil(o.RequestMethod) {
		return true
	}

	return false
}

// SetRequestMethod gets a reference to the given string and assigns it to the RequestMethod field.
func (o *TestScriptAssert) SetRequestMethod(v string) {
	o.RequestMethod = &v
}

// GetRequestURL returns the RequestURL field value if set, zero value otherwise.
func (o *TestScriptAssert) GetRequestURL() string {
	if o == nil || IsNil(o.RequestURL) {
		var ret string
		return ret
	}
	return *o.RequestURL
}

// GetRequestURLOk returns a tuple with the RequestURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetRequestURLOk() (*string, bool) {
	if o == nil || IsNil(o.RequestURL) {
		return nil, false
	}
	return o.RequestURL, true
}

// HasRequestURL returns a boolean if a field has been set.
func (o *TestScriptAssert) HasRequestURL() bool {
	if o != nil && !IsNil(o.RequestURL) {
		return true
	}

	return false
}

// SetRequestURL gets a reference to the given string and assigns it to the RequestURL field.
func (o *TestScriptAssert) SetRequestURL(v string) {
	o.RequestURL = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *TestScriptAssert) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *TestScriptAssert) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *TestScriptAssert) SetResource(v string) {
	o.Resource = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TestScriptAssert) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TestScriptAssert) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *TestScriptAssert) SetResponse(v string) {
	o.Response = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *TestScriptAssert) GetResponseCode() string {
	if o == nil || IsNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetResponseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *TestScriptAssert) HasResponseCode() bool {
	if o != nil && !IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *TestScriptAssert) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *TestScriptAssert) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *TestScriptAssert) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *TestScriptAssert) SetSourceId(v string) {
	o.SourceId = &v
}

// GetValidateProfileId returns the ValidateProfileId field value if set, zero value otherwise.
func (o *TestScriptAssert) GetValidateProfileId() string {
	if o == nil || IsNil(o.ValidateProfileId) {
		var ret string
		return ret
	}
	return *o.ValidateProfileId
}

// GetValidateProfileIdOk returns a tuple with the ValidateProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetValidateProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ValidateProfileId) {
		return nil, false
	}
	return o.ValidateProfileId, true
}

// HasValidateProfileId returns a boolean if a field has been set.
func (o *TestScriptAssert) HasValidateProfileId() bool {
	if o != nil && !IsNil(o.ValidateProfileId) {
		return true
	}

	return false
}

// SetValidateProfileId gets a reference to the given string and assigns it to the ValidateProfileId field.
func (o *TestScriptAssert) SetValidateProfileId(v string) {
	o.ValidateProfileId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestScriptAssert) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestScriptAssert) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TestScriptAssert) SetValue(v string) {
	o.Value = &v
}

// GetWarningOnly returns the WarningOnly field value if set, zero value otherwise.
func (o *TestScriptAssert) GetWarningOnly() bool {
	if o == nil || IsNil(o.WarningOnly) {
		var ret bool
		return ret
	}
	return *o.WarningOnly
}

// GetWarningOnlyOk returns a tuple with the WarningOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestScriptAssert) GetWarningOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.WarningOnly) {
		return nil, false
	}
	return o.WarningOnly, true
}

// HasWarningOnly returns a boolean if a field has been set.
func (o *TestScriptAssert) HasWarningOnly() bool {
	if o != nil && !IsNil(o.WarningOnly) {
		return true
	}

	return false
}

// SetWarningOnly gets a reference to the given bool and assigns it to the WarningOnly field.
func (o *TestScriptAssert) SetWarningOnly(v bool) {
	o.WarningOnly = &v
}

func (o TestScriptAssert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestScriptAssert) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.CompareToSourceId) {
		toSerialize["compareToSourceId"] = o.CompareToSourceId
	}
	if !IsNil(o.CompareToSourceExpression) {
		toSerialize["compareToSourceExpression"] = o.CompareToSourceExpression
	}
	if !IsNil(o.CompareToSourcePath) {
		toSerialize["compareToSourcePath"] = o.CompareToSourcePath
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.HeaderField) {
		toSerialize["headerField"] = o.HeaderField
	}
	if !IsNil(o.MinimumId) {
		toSerialize["minimumId"] = o.MinimumId
	}
	if !IsNil(o.NavigationLinks) {
		toSerialize["navigationLinks"] = o.NavigationLinks
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RequestMethod) {
		toSerialize["requestMethod"] = o.RequestMethod
	}
	if !IsNil(o.RequestURL) {
		toSerialize["requestURL"] = o.RequestURL
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.ValidateProfileId) {
		toSerialize["validateProfileId"] = o.ValidateProfileId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.WarningOnly) {
		toSerialize["warningOnly"] = o.WarningOnly
	}
	return toSerialize, nil
}

type NullableTestScriptAssert struct {
	value *TestScriptAssert
	isSet bool
}

func (v NullableTestScriptAssert) Get() *TestScriptAssert {
	return v.value
}

func (v *NullableTestScriptAssert) Set(val *TestScriptAssert) {
	v.value = val
	v.isSet = true
}

func (v NullableTestScriptAssert) IsSet() bool {
	return v.isSet
}

func (v *NullableTestScriptAssert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestScriptAssert(val *TestScriptAssert) *NullableTestScriptAssert {
	return &NullableTestScriptAssert{value: val, isSet: true}
}

func (v NullableTestScriptAssert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestScriptAssert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

