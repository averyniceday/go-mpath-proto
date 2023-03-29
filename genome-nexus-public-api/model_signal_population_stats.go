/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"encoding/json"
)

// checks if the SignalPopulationStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalPopulationStats{}

// SignalPopulationStats struct for SignalPopulationStats
type SignalPopulationStats struct {
	// African/African American
	Afr *float64 `json:"afr,omitempty"`
	// Ashkenazi Jewish
	Asj *float64 `json:"asj,omitempty"`
	// Asian
	Asn *float64 `json:"asn,omitempty"`
	// European
	Eur *float64 `json:"eur,omitempty"`
	// Impact
	Impact *float64 `json:"impact,omitempty"`
	// Other
	Oth *float64 `json:"oth,omitempty"`
}

// NewSignalPopulationStats instantiates a new SignalPopulationStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalPopulationStats() *SignalPopulationStats {
	this := SignalPopulationStats{}
	return &this
}

// NewSignalPopulationStatsWithDefaults instantiates a new SignalPopulationStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalPopulationStatsWithDefaults() *SignalPopulationStats {
	this := SignalPopulationStats{}
	return &this
}

// GetAfr returns the Afr field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetAfr() float64 {
	if o == nil || isNil(o.Afr) {
		var ret float64
		return ret
	}
	return *o.Afr
}

// GetAfrOk returns a tuple with the Afr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetAfrOk() (*float64, bool) {
	if o == nil || isNil(o.Afr) {
		return nil, false
	}
	return o.Afr, true
}

// HasAfr returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasAfr() bool {
	if o != nil && !isNil(o.Afr) {
		return true
	}

	return false
}

// SetAfr gets a reference to the given float64 and assigns it to the Afr field.
func (o *SignalPopulationStats) SetAfr(v float64) {
	o.Afr = &v
}

// GetAsj returns the Asj field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetAsj() float64 {
	if o == nil || isNil(o.Asj) {
		var ret float64
		return ret
	}
	return *o.Asj
}

// GetAsjOk returns a tuple with the Asj field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetAsjOk() (*float64, bool) {
	if o == nil || isNil(o.Asj) {
		return nil, false
	}
	return o.Asj, true
}

// HasAsj returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasAsj() bool {
	if o != nil && !isNil(o.Asj) {
		return true
	}

	return false
}

// SetAsj gets a reference to the given float64 and assigns it to the Asj field.
func (o *SignalPopulationStats) SetAsj(v float64) {
	o.Asj = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetAsn() float64 {
	if o == nil || isNil(o.Asn) {
		var ret float64
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetAsnOk() (*float64, bool) {
	if o == nil || isNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasAsn() bool {
	if o != nil && !isNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given float64 and assigns it to the Asn field.
func (o *SignalPopulationStats) SetAsn(v float64) {
	o.Asn = &v
}

// GetEur returns the Eur field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetEur() float64 {
	if o == nil || isNil(o.Eur) {
		var ret float64
		return ret
	}
	return *o.Eur
}

// GetEurOk returns a tuple with the Eur field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetEurOk() (*float64, bool) {
	if o == nil || isNil(o.Eur) {
		return nil, false
	}
	return o.Eur, true
}

// HasEur returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasEur() bool {
	if o != nil && !isNil(o.Eur) {
		return true
	}

	return false
}

// SetEur gets a reference to the given float64 and assigns it to the Eur field.
func (o *SignalPopulationStats) SetEur(v float64) {
	o.Eur = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetImpact() float64 {
	if o == nil || isNil(o.Impact) {
		var ret float64
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetImpactOk() (*float64, bool) {
	if o == nil || isNil(o.Impact) {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasImpact() bool {
	if o != nil && !isNil(o.Impact) {
		return true
	}

	return false
}

// SetImpact gets a reference to the given float64 and assigns it to the Impact field.
func (o *SignalPopulationStats) SetImpact(v float64) {
	o.Impact = &v
}

// GetOth returns the Oth field value if set, zero value otherwise.
func (o *SignalPopulationStats) GetOth() float64 {
	if o == nil || isNil(o.Oth) {
		var ret float64
		return ret
	}
	return *o.Oth
}

// GetOthOk returns a tuple with the Oth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalPopulationStats) GetOthOk() (*float64, bool) {
	if o == nil || isNil(o.Oth) {
		return nil, false
	}
	return o.Oth, true
}

// HasOth returns a boolean if a field has been set.
func (o *SignalPopulationStats) HasOth() bool {
	if o != nil && !isNil(o.Oth) {
		return true
	}

	return false
}

// SetOth gets a reference to the given float64 and assigns it to the Oth field.
func (o *SignalPopulationStats) SetOth(v float64) {
	o.Oth = &v
}

func (o SignalPopulationStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalPopulationStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Afr) {
		toSerialize["afr"] = o.Afr
	}
	if !isNil(o.Asj) {
		toSerialize["asj"] = o.Asj
	}
	if !isNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !isNil(o.Eur) {
		toSerialize["eur"] = o.Eur
	}
	if !isNil(o.Impact) {
		toSerialize["impact"] = o.Impact
	}
	if !isNil(o.Oth) {
		toSerialize["oth"] = o.Oth
	}
	return toSerialize, nil
}

type NullableSignalPopulationStats struct {
	value *SignalPopulationStats
	isSet bool
}

func (v NullableSignalPopulationStats) Get() *SignalPopulationStats {
	return v.value
}

func (v *NullableSignalPopulationStats) Set(val *SignalPopulationStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalPopulationStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalPopulationStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalPopulationStats(val *SignalPopulationStats) *NullableSignalPopulationStats {
	return &NullableSignalPopulationStats{value: val, isSet: true}
}

func (v NullableSignalPopulationStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalPopulationStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

