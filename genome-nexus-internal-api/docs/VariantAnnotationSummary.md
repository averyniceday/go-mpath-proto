# VariantAnnotationSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyName** | **string** | Assembly name | [optional] [default to null]
**CanonicalTranscriptId** | **string** | Canonical transcript id | [optional] [default to null]
**GenomicLocation** | [***GenomicLocation**](GenomicLocation.md) | Genomic location | [optional] [default to null]
**StrandSign** | **string** | Strand (- or +) | [optional] [default to null]
**TranscriptConsequenceSummaries** | [**[]TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) | All transcript consequence summaries | [default to null]
**TranscriptConsequenceSummary** | [***TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) | Most impactful transcript consequence of canonical transcript or if non-existent any transcript | [default to null]
**TranscriptConsequences** | [**[]TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) | (Deprecated) Transcript consequence summaries (list of one when using annotation/, multiple when using annotation/summary/ | [default to null]
**Variant** | **string** | Variant key | [default to null]
**VariantType** | **string** | Variant type | [optional] [default to null]
**Vues** | [***Vues**](Vues.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


