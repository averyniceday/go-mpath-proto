# \AnnotationControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVariantAnnotationByGenomicLocationGET**](AnnotationControllerApi.md#FetchVariantAnnotationByGenomicLocationGET) | **Get** /annotation/genomic/{genomicLocation} | Retrieves VEP annotation for the provided genomic location
[**FetchVariantAnnotationByGenomicLocationPOST**](AnnotationControllerApi.md#FetchVariantAnnotationByGenomicLocationPOST) | **Post** /annotation/genomic | Retrieves VEP annotation for the provided list of genomic locations
[**FetchVariantAnnotationByIdGET**](AnnotationControllerApi.md#FetchVariantAnnotationByIdGET) | **Get** /annotation/dbsnp/{variantId} | Retrieves VEP annotation for the give dbSNP id
[**FetchVariantAnnotationByIdPOST**](AnnotationControllerApi.md#FetchVariantAnnotationByIdPOST) | **Post** /annotation/dbsnp/ | Retrieves VEP annotation for the provided list of dbSNP ids
[**FetchVariantAnnotationGET**](AnnotationControllerApi.md#FetchVariantAnnotationGET) | **Get** /annotation/{variant} | Retrieves VEP annotation for the provided variant
[**FetchVariantAnnotationPOST**](AnnotationControllerApi.md#FetchVariantAnnotationPOST) | **Post** /annotation | Retrieves VEP annotation for the provided list of variants


# **FetchVariantAnnotationByGenomicLocationGET**
> VariantAnnotation FetchVariantAnnotationByGenomicLocationGET(ctx, genomicLocation, optional)
Retrieves VEP annotation for the provided genomic location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **genomicLocation** | **string**| A genomic location. For example 7,140453136,140453136,A,T | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genomicLocation** | **string**| A genomic location. For example 7,140453136,140453136,A,T | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: hotspots | [default to hotspots]

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationByGenomicLocationPOST**
> []VariantAnnotation FetchVariantAnnotationByGenomicLocationPOST(ctx, genomicLocations, optional)
Retrieves VEP annotation for the provided list of genomic locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **genomicLocations** | [**[]GenomicLocation**](GenomicLocation.md)| List of Genomic Locations | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genomicLocations** | [**[]GenomicLocation**](GenomicLocation.md)| List of Genomic Locations | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: hotspots | [default to hotspots]

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationByIdGET**
> VariantAnnotation FetchVariantAnnotationByIdGET(ctx, variantId, optional)
Retrieves VEP annotation for the give dbSNP id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variantId** | **string**| dbSNP id. For example rs116035550. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variantId** | **string**| dbSNP id. For example rs116035550. | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: annotation_summary | [default to annotation_summary]

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationByIdPOST**
> []VariantAnnotation FetchVariantAnnotationByIdPOST(ctx, variantIds, optional)
Retrieves VEP annotation for the provided list of dbSNP ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variantIds** | **[]string**| List of variant IDs. For example [\&quot;rs116035550\&quot;] | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variantIds** | **[]string**| List of variant IDs. For example [\&quot;rs116035550\&quot;] | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: annotation_summary | [default to annotation_summary]

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationGET**
> VariantAnnotation FetchVariantAnnotationGET(ctx, variant, optional)
Retrieves VEP annotation for the provided variant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| Variant. For example 17:g.41242962_41242963insGA | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | **string**| Variant. For example 17:g.41242962_41242963insGA | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: hotspots | [default to hotspots]

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationPOST**
> []VariantAnnotation FetchVariantAnnotationPOST(ctx, variants, optional)
Retrieves VEP annotation for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variants** | **[]string**| List of variants. For example [\&quot;X:g.66937331T&gt;A\&quot;,\&quot;17:g.41242962_41242963insGA\&quot;] (GRCh37) or [\&quot;1:g.182712A&gt;C\&quot;, \&quot;2:g.265023C&gt;T\&quot;, \&quot;3:g.319781del\&quot;, \&quot;19:g.110753dup\&quot;, \&quot;1:g.1385015_1387562del\&quot;] (GRCh38) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variants** | **[]string**| List of variants. For example [\&quot;X:g.66937331T&gt;A\&quot;,\&quot;17:g.41242962_41242963insGA\&quot;] (GRCh37) or [\&quot;1:g.182712A&gt;C\&quot;, \&quot;2:g.265023C&gt;T\&quot;, \&quot;3:g.319781del\&quot;, \&quot;19:g.110753dup\&quot;, \&quot;1:g.1385015_1387562del\&quot;] (GRCh38) | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | 
 **token** | **string**| Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | [**[]string**](string.md)| Comma separated list of fields to include (case-sensitive!). For example: hotspots | [default to hotspots]

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

