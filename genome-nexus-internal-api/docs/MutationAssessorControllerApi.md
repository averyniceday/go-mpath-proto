# \MutationAssessorControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMutationAssessorAnnotationGET**](MutationAssessorControllerApi.md#FetchMutationAssessorAnnotationGET) | **Get** /mutation_assessor/{variant} | Retrieves mutation assessor information for the provided list of variants
[**PostMutationAssessorAnnotation**](MutationAssessorControllerApi.md#PostMutationAssessorAnnotation) | **Post** /mutation_assessor | Retrieves mutation assessor information for the provided list of variants


# **FetchMutationAssessorAnnotationGET**
> MutationAssessor FetchMutationAssessorAnnotationGET(ctx, variant)
Retrieves mutation assessor information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| A variant. For example 7:g.140453136A&gt;T | 

### Return type

[**MutationAssessor**](MutationAssessor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMutationAssessorAnnotation**
> []MutationAssessor PostMutationAssessorAnnotation(ctx, variants)
Retrieves mutation assessor information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variants** | **[]string**| List of variants. For example [\&quot;7:g.140453136A&gt;T\&quot;,\&quot;12:g.25398285C&gt;A\&quot;] | 

### Return type

[**[]MutationAssessor**](MutationAssessor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

