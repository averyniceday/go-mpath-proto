# \NucleotideContextControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchNucleotideContextAnnotationGET**](NucleotideContextControllerApi.md#FetchNucleotideContextAnnotationGET) | **Get** /nucleotide_context/{variant} | Retrieves nucleotide context information for the provided list of variants
[**PostNucleotideContextAnnotation**](NucleotideContextControllerApi.md#PostNucleotideContextAnnotation) | **Post** /nucleotide_context | Retrieves nucleotide context information for the provided list of variants


# **FetchNucleotideContextAnnotationGET**
> NucleotideContext FetchNucleotideContextAnnotationGET(ctx, variant)
Retrieves nucleotide context information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| A variant. For example 7:g.140453136A&gt;T | 

### Return type

[**NucleotideContext**](NucleotideContext.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostNucleotideContextAnnotation**
> []NucleotideContext PostNucleotideContextAnnotation(ctx, variants)
Retrieves nucleotide context information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variants** | **[]string**| List of variants. For example [\&quot;7:g.140453136A&gt;T\&quot;,\&quot;12:g.25398285C&gt;A\&quot;] | 

### Return type

[**[]NucleotideContext**](NucleotideContext.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

