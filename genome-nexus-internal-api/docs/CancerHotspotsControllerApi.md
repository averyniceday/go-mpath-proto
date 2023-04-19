# \CancerHotspotsControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchHotspotAnnotationByGenomicLocationGET**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByGenomicLocationGET) | **Get** /cancer_hotspots/genomic/{genomicLocation} | Retrieves hotspot annotations for a specific genomic location
[**FetchHotspotAnnotationByGenomicLocationPOST**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByGenomicLocationPOST) | **Post** /cancer_hotspots/genomic | Retrieves hotspot annotations for the provided list of genomic locations
[**FetchHotspotAnnotationByHgvsGET**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByHgvsGET) | **Get** /cancer_hotspots/hgvs/{variant} | Retrieves hotspot annotations for a specific variant
[**FetchHotspotAnnotationByHgvsPOST**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByHgvsPOST) | **Post** /cancer_hotspots/hgvs | Retrieves hotspot annotations for the provided list of variants
[**FetchHotspotAnnotationByProteinLocationsPOST**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByProteinLocationsPOST) | **Post** /cancer_hotspots/proteinLocations | Retrieves hotspot annotations for the provided list of transcript id, protein location and mutation type
[**FetchHotspotAnnotationByTranscriptIdGET**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByTranscriptIdGET) | **Get** /cancer_hotspots/transcript/{transcriptId} | Retrieves hotspot annotations for the provided transcript ID
[**FetchHotspotAnnotationByTranscriptIdPOST**](CancerHotspotsControllerApi.md#FetchHotspotAnnotationByTranscriptIdPOST) | **Post** /cancer_hotspots/transcript | Retrieves hotspot annotations for the provided list of transcript ID


# **FetchHotspotAnnotationByGenomicLocationGET**
> []Hotspot FetchHotspotAnnotationByGenomicLocationGET(ctx, genomicLocation)
Retrieves hotspot annotations for a specific genomic location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **genomicLocation** | **string**| A genomic location. For example 7,140453136,140453136,A,T | 

### Return type

[**[]Hotspot**](Hotspot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByGenomicLocationPOST**
> []AggregatedHotspots FetchHotspotAnnotationByGenomicLocationPOST(ctx, genomicLocations)
Retrieves hotspot annotations for the provided list of genomic locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **genomicLocations** | [**[]GenomicLocation**](GenomicLocation.md)| List of genomic locations. | 

### Return type

[**[]AggregatedHotspots**](AggregatedHotspots.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByHgvsGET**
> []Hotspot FetchHotspotAnnotationByHgvsGET(ctx, variant)
Retrieves hotspot annotations for a specific variant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| A variant. For example 7:g.140453136A&gt;T | 

### Return type

[**[]Hotspot**](Hotspot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByHgvsPOST**
> []AggregatedHotspots FetchHotspotAnnotationByHgvsPOST(ctx, variants)
Retrieves hotspot annotations for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variants** | **[]string**| List of variants. For example [\&quot;7:g.140453136A&gt;T\&quot;,\&quot;12:g.25398285C&gt;A\&quot;] | 

### Return type

[**[]AggregatedHotspots**](AggregatedHotspots.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByProteinLocationsPOST**
> []AggregatedHotspots FetchHotspotAnnotationByProteinLocationsPOST(ctx, proteinLocations)
Retrieves hotspot annotations for the provided list of transcript id, protein location and mutation type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **proteinLocations** | [**[]ProteinLocation**](ProteinLocation.md)| List of transcript id, protein start location, protein end location, mutation type. The mutation types are limited to &#39;Missense_Mutation&#39;, &#39;In_Frame_Ins&#39;, &#39;In_Frame_Del&#39;, &#39;Splice_Site&#39;, and &#39;Splice_Region&#39; | 

### Return type

[**[]AggregatedHotspots**](AggregatedHotspots.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByTranscriptIdGET**
> []Hotspot FetchHotspotAnnotationByTranscriptIdGET(ctx, transcriptId)
Retrieves hotspot annotations for the provided transcript ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transcriptId** | **string**| A Transcript Id. For example ENST00000288602 | 

### Return type

[**[]Hotspot**](Hotspot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchHotspotAnnotationByTranscriptIdPOST**
> []AggregatedHotspots FetchHotspotAnnotationByTranscriptIdPOST(ctx, transcriptIds)
Retrieves hotspot annotations for the provided list of transcript ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transcriptIds** | **[]string**| List of transcript Id. For example [\&quot;ENST00000288602\&quot;,\&quot;ENST00000256078\&quot;] | 

### Return type

[**[]AggregatedHotspots**](AggregatedHotspots.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

