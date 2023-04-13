# VariantAnnotation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlleleString** | **string** | Allele string (e.g: A/T) | [optional] [default to null]
**AnnotationJSON** | **string** | Annotation data as JSON string | [optional] [default to null]
**AnnotationSummary** | [***VariantAnnotationSummary**](VariantAnnotationSummary.md) | Variant Annotation Summary | [optional] [default to null]
**AssemblyName** | **string** | NCBI build number | [optional] [default to null]
**Clinvar** | [***ClinvarAnnotation**](ClinvarAnnotation.md) | ClinVar | [optional] [default to null]
**ColocatedVariants** | [**[]ColocatedVariant**](ColocatedVariant.md) |  | [optional] [default to null]
**End** | **int32** | End position | [optional] [default to null]
**Hgvsg** | **string** |  | [optional] [default to null]
**Hotspots** | [***HotspotAnnotation**](HotspotAnnotation.md) | Hotspot Annotation | [optional] [default to null]
**Id** | **string** | Variant id | [default to null]
**IntergenicConsequences** | [**[]IntergenicConsequences**](IntergenicConsequences.md) | intergenicConsequences | [default to null]
**MostSevereConsequence** | **string** | Most severe consequence | [optional] [default to null]
**MutationAssessor** | [***MutationAssessorAnnotation**](MutationAssessorAnnotation.md) | Mutation Assessor Annotation | [optional] [default to null]
**MyVariantInfo** | [***MyVariantInfoAnnotation**](MyVariantInfoAnnotation.md) | My Variant Info Annotation | [optional] [default to null]
**NucleotideContext** | [***NucleotideContextAnnotation**](NucleotideContextAnnotation.md) | Nucleotide Context Annotation | [optional] [default to null]
**Oncokb** | [***OncokbAnnotation**](OncokbAnnotation.md) | Oncokb | [optional] [default to null]
**OriginalVariantQuery** | **string** | Original variant query | [default to null]
**Ptms** | [***PtmAnnotation**](PtmAnnotation.md) | Post Translational Modifications | [optional] [default to null]
**SeqRegionName** | **string** | Chromosome | [optional] [default to null]
**SignalAnnotation** | [***SignalAnnotation**](SignalAnnotation.md) |  | [optional] [default to null]
**Start** | **int32** | Start position | [optional] [default to null]
**Strand** | **int32** | Strand (negative or positive) | [optional] [default to null]
**SuccessfullyAnnotated** | **bool** | Status flag indicating whether variant was succesfully annotated | [optional] [default to null]
**TranscriptConsequences** | [**[]TranscriptConsequence**](TranscriptConsequence.md) | List of transcripts | [optional] [default to null]
**Variant** | **string** | Variant key | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


