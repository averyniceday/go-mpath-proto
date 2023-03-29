# Clinvar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateAllele** | Pointer to **string** |  | [optional] 
**Chromosome** | Pointer to **string** |  | [optional] 
**ClinicalSignificance** | Pointer to **string** |  | [optional] 
**ClinvarId** | Pointer to **int32** |  | [optional] 
**ConflictingClinicalSignificance** | Pointer to **string** |  | [optional] 
**EndPosition** | Pointer to **int32** |  | [optional] 
**ReferenceAllele** | Pointer to **string** |  | [optional] 
**StartPosition** | Pointer to **int32** |  | [optional] 

## Methods

### NewClinvar

`func NewClinvar() *Clinvar`

NewClinvar instantiates a new Clinvar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClinvarWithDefaults

`func NewClinvarWithDefaults() *Clinvar`

NewClinvarWithDefaults instantiates a new Clinvar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateAllele

`func (o *Clinvar) GetAlternateAllele() string`

GetAlternateAllele returns the AlternateAllele field if non-nil, zero value otherwise.

### GetAlternateAlleleOk

`func (o *Clinvar) GetAlternateAlleleOk() (*string, bool)`

GetAlternateAlleleOk returns a tuple with the AlternateAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAllele

`func (o *Clinvar) SetAlternateAllele(v string)`

SetAlternateAllele sets AlternateAllele field to given value.

### HasAlternateAllele

`func (o *Clinvar) HasAlternateAllele() bool`

HasAlternateAllele returns a boolean if a field has been set.

### GetChromosome

`func (o *Clinvar) GetChromosome() string`

GetChromosome returns the Chromosome field if non-nil, zero value otherwise.

### GetChromosomeOk

`func (o *Clinvar) GetChromosomeOk() (*string, bool)`

GetChromosomeOk returns a tuple with the Chromosome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosome

`func (o *Clinvar) SetChromosome(v string)`

SetChromosome sets Chromosome field to given value.

### HasChromosome

`func (o *Clinvar) HasChromosome() bool`

HasChromosome returns a boolean if a field has been set.

### GetClinicalSignificance

`func (o *Clinvar) GetClinicalSignificance() string`

GetClinicalSignificance returns the ClinicalSignificance field if non-nil, zero value otherwise.

### GetClinicalSignificanceOk

`func (o *Clinvar) GetClinicalSignificanceOk() (*string, bool)`

GetClinicalSignificanceOk returns a tuple with the ClinicalSignificance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalSignificance

`func (o *Clinvar) SetClinicalSignificance(v string)`

SetClinicalSignificance sets ClinicalSignificance field to given value.

### HasClinicalSignificance

`func (o *Clinvar) HasClinicalSignificance() bool`

HasClinicalSignificance returns a boolean if a field has been set.

### GetClinvarId

`func (o *Clinvar) GetClinvarId() int32`

GetClinvarId returns the ClinvarId field if non-nil, zero value otherwise.

### GetClinvarIdOk

`func (o *Clinvar) GetClinvarIdOk() (*int32, bool)`

GetClinvarIdOk returns a tuple with the ClinvarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinvarId

`func (o *Clinvar) SetClinvarId(v int32)`

SetClinvarId sets ClinvarId field to given value.

### HasClinvarId

`func (o *Clinvar) HasClinvarId() bool`

HasClinvarId returns a boolean if a field has been set.

### GetConflictingClinicalSignificance

`func (o *Clinvar) GetConflictingClinicalSignificance() string`

GetConflictingClinicalSignificance returns the ConflictingClinicalSignificance field if non-nil, zero value otherwise.

### GetConflictingClinicalSignificanceOk

`func (o *Clinvar) GetConflictingClinicalSignificanceOk() (*string, bool)`

GetConflictingClinicalSignificanceOk returns a tuple with the ConflictingClinicalSignificance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingClinicalSignificance

`func (o *Clinvar) SetConflictingClinicalSignificance(v string)`

SetConflictingClinicalSignificance sets ConflictingClinicalSignificance field to given value.

### HasConflictingClinicalSignificance

`func (o *Clinvar) HasConflictingClinicalSignificance() bool`

HasConflictingClinicalSignificance returns a boolean if a field has been set.

### GetEndPosition

`func (o *Clinvar) GetEndPosition() int32`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *Clinvar) GetEndPositionOk() (*int32, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *Clinvar) SetEndPosition(v int32)`

SetEndPosition sets EndPosition field to given value.

### HasEndPosition

`func (o *Clinvar) HasEndPosition() bool`

HasEndPosition returns a boolean if a field has been set.

### GetReferenceAllele

`func (o *Clinvar) GetReferenceAllele() string`

GetReferenceAllele returns the ReferenceAllele field if non-nil, zero value otherwise.

### GetReferenceAlleleOk

`func (o *Clinvar) GetReferenceAlleleOk() (*string, bool)`

GetReferenceAlleleOk returns a tuple with the ReferenceAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAllele

`func (o *Clinvar) SetReferenceAllele(v string)`

SetReferenceAllele sets ReferenceAllele field to given value.

### HasReferenceAllele

`func (o *Clinvar) HasReferenceAllele() bool`

HasReferenceAllele returns a boolean if a field has been set.

### GetStartPosition

`func (o *Clinvar) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *Clinvar) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *Clinvar) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.

### HasStartPosition

`func (o *Clinvar) HasStartPosition() bool`

HasStartPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


