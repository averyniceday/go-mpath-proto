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

// Hg19 struct for Hg19
type Hg19 struct {
	// end
	End *int32 `json:"end,omitempty"`
	// start
	Start *int32 `json:"start,omitempty"`
}

// NewHg19 instantiates a new Hg19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHg19() *Hg19 {
	this := Hg19{}
	return &this
}

// NewHg19WithDefaults instantiates a new Hg19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHg19WithDefaults() *Hg19 {
	this := Hg19{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Hg19) GetEnd() int32 {
	if o == nil || isNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hg19) GetEndOk() (*int32, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Hg19) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *Hg19) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Hg19) GetStart() int32 {
	if o == nil || isNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hg19) GetStartOk() (*int32, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Hg19) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *Hg19) SetStart(v int32) {
	o.Start = &v
}

func (o Hg19) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableHg19 struct {
	value *Hg19
	isSet bool
}

func (v NullableHg19) Get() *Hg19 {
	return v.value
}

func (v *NullableHg19) Set(val *Hg19) {
	v.value = val
	v.isSet = true
}

func (v NullableHg19) IsSet() bool {
	return v.isSet
}

func (v *NullableHg19) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHg19(val *Hg19) *NullableHg19 {
	return &NullableHg19{value: val, isSet: true}
}

func (v NullableHg19) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHg19) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

