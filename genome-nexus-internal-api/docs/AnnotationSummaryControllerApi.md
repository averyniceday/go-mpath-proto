# \AnnotationSummaryControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVariantAnnotationSummaryGET**](AnnotationSummaryControllerApi.md#FetchVariantAnnotationSummaryGET) | **Get** /annotation/summary/{variant} | Retrieves VEP annotation summary for the provided variant
[**FetchVariantAnnotationSummaryPOST**](AnnotationSummaryControllerApi.md#FetchVariantAnnotationSummaryPOST) | **Post** /annotation/summary | Retrieves VEP annotation summary for the provided list of variants


# **FetchVariantAnnotationSummaryGET**
> VariantAnnotationSummary FetchVariantAnnotationSummaryGET(ctx, variant, optional)
Retrieves VEP annotation summary for the provided variant

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
 **projection** | **string**| Indicates whether to return summary for all transcripts or only for canonical transcript | [default to ALL]

### Return type

[**VariantAnnotationSummary**](VariantAnnotationSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVariantAnnotationSummaryPOST**
> []VariantAnnotationSummary FetchVariantAnnotationSummaryPOST(ctx, variants, optional)
Retrieves VEP annotation summary for the provided list of variants

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
 **projection** | **string**| Indicates whether to return summary for all transcripts or only for canonical transcript | [default to ALL]

### Return type

[**[]VariantAnnotationSummary**](VariantAnnotationSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

