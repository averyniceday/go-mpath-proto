# \CuriousCasesControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCuriousCasesGET**](CuriousCasesControllerApi.md#FetchCuriousCasesGET) | **Get** /curious_cases/{genomicLocation} | Retrieves Curious Cases info by a genomic location


# **FetchCuriousCasesGET**
> CuriousCases FetchCuriousCasesGET(ctx, genomicLocation)
Retrieves Curious Cases info by a genomic location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **genomicLocation** | **string**| Genomic location, for example: 7,116411883,116411905,TTCTTTCTCTCTGTTTTAAGATC,- | 

### Return type

[**CuriousCases**](CuriousCases.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

