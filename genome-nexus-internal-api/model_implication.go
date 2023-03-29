/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_internal_api

import (
	"encoding/json"
)

// checks if the Implication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Implication{}

// Implication struct for Implication
type Implication struct {
	Alterations []string `json:"alterations,omitempty"`
	Description *string `json:"description,omitempty"`
	LevelOfEvidence *string `json:"levelOfEvidence,omitempty"`
	TumorType *TumorType `json:"tumorType,omitempty"`
}

// NewImplication instantiates a new Implication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImplication() *Implication {
	this := Implication{}
	return &this
}

// NewImplicationWithDefaults instantiates a new Implication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImplicationWithDefaults() *Implication {
	this := Implication{}
	return &this
}

// GetAlterations returns the Alterations field value if set, zero value otherwise.
func (o *Implication) GetAlterations() []string {
	if o == nil || isNil(o.Alterations) {
		var ret []string
		return ret
	}
	return o.Alterations
}

// GetAlterationsOk returns a tuple with the Alterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Implication) GetAlterationsOk() ([]string, bool) {
	if o == nil || isNil(o.Alterations) {
		return nil, false
	}
	return o.Alterations, true
}

// HasAlterations returns a boolean if a field has been set.
func (o *Implication) HasAlterations() bool {
	if o != nil && !isNil(o.Alterations) {
		return true
	}

	return false
}

// SetAlterations gets a reference to the given []string and assigns it to the Alterations field.
func (o *Implication) SetAlterations(v []string) {
	o.Alterations = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Implication) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Implication) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Implication) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Implication) SetDescription(v string) {
	o.Description = &v
}

// GetLevelOfEvidence returns the LevelOfEvidence field value if set, zero value otherwise.
func (o *Implication) GetLevelOfEvidence() string {
	if o == nil || isNil(o.LevelOfEvidence) {
		var ret string
		return ret
	}
	return *o.LevelOfEvidence
}

// GetLevelOfEvidenceOk returns a tuple with the LevelOfEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Implication) GetLevelOfEvidenceOk() (*string, bool) {
	if o == nil || isNil(o.LevelOfEvidence) {
		return nil, false
	}
	return o.LevelOfEvidence, true
}

// HasLevelOfEvidence returns a boolean if a field has been set.
func (o *Implication) HasLevelOfEvidence() bool {
	if o != nil && !isNil(o.LevelOfEvidence) {
		return true
	}

	return false
}

// SetLevelOfEvidence gets a reference to the given string and assigns it to the LevelOfEvidence field.
func (o *Implication) SetLevelOfEvidence(v string) {
	o.LevelOfEvidence = &v
}

// GetTumorType returns the TumorType field value if set, zero value otherwise.
func (o *Implication) GetTumorType() TumorType {
	if o == nil || isNil(o.TumorType) {
		var ret TumorType
		return ret
	}
	return *o.TumorType
}

// GetTumorTypeOk returns a tuple with the TumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Implication) GetTumorTypeOk() (*TumorType, bool) {
	if o == nil || isNil(o.TumorType) {
		return nil, false
	}
	return o.TumorType, true
}

// HasTumorType returns a boolean if a field has been set.
func (o *Implication) HasTumorType() bool {
	if o != nil && !isNil(o.TumorType) {
		return true
	}

	return false
}

// SetTumorType gets a reference to the given TumorType and assigns it to the TumorType field.
func (o *Implication) SetTumorType(v TumorType) {
	o.TumorType = &v
}

func (o Implication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Implication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alterations) {
		toSerialize["alterations"] = o.Alterations
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.LevelOfEvidence) {
		toSerialize["levelOfEvidence"] = o.LevelOfEvidence
	}
	if !isNil(o.TumorType) {
		toSerialize["tumorType"] = o.TumorType
	}
	return toSerialize, nil
}

type NullableImplication struct {
	value *Implication
	isSet bool
}

func (v NullableImplication) Get() *Implication {
	return v.value
}

func (v *NullableImplication) Set(val *Implication) {
	v.value = val
	v.isSet = true
}

func (v NullableImplication) IsSet() bool {
	return v.isSet
}

func (v *NullableImplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImplication(val *Implication) *NullableImplication {
	return &NullableImplication{value: val, isSet: true}
}

func (v NullableImplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

