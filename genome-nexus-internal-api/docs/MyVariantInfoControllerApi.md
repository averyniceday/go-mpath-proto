# \MyVariantInfoControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMyVariantInfoAnnotationGET**](MyVariantInfoControllerApi.md#FetchMyVariantInfoAnnotationGET) | **Get** /my_variant_info/variant/{variant} | Retrieves myvariant information for the provided list of variants
[**PostMyVariantInfoAnnotation**](MyVariantInfoControllerApi.md#PostMyVariantInfoAnnotation) | **Post** /my_variant_info/variant | Retrieves myvariant information for the provided list of variants


# **FetchMyVariantInfoAnnotationGET**
> MyVariantInfo FetchMyVariantInfoAnnotationGET(ctx, variant)
Retrieves myvariant information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| . For example 7:g.140453136A&gt;T | 

### Return type

[**MyVariantInfo**](MyVariantInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMyVariantInfoAnnotation**
> []MyVariantInfo PostMyVariantInfoAnnotation(ctx, variants)
Retrieves myvariant information for the provided list of variants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variants** | **[]string**| List of variants. For example [\&quot;7:g.140453136A&gt;T\&quot;,\&quot;12:g.25398285C&gt;A\&quot;] | 

### Return type

[**[]MyVariantInfo**](MyVariantInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

