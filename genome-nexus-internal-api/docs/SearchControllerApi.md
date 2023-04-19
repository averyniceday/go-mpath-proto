# \SearchControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAnnotationByKeywordGETUsingGET**](SearchControllerApi.md#SearchAnnotationByKeywordGETUsingGET) | **Get** /search | Performs index search.


# **SearchAnnotationByKeywordGETUsingGET**
> []IndexSearch SearchAnnotationByKeywordGETUsingGET(ctx, keyword, optional)
Performs index search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyword** | **string**| keyword. For example 13:g.32890665G&gt;A, TP53 p.R273C, BRAF c.1799T&gt;A | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string**| keyword. For example 13:g.32890665G&gt;A, TP53 p.R273C, BRAF c.1799T&gt;A | 
 **limit** | **int32**| Max number of matching results to return | 

### Return type

[**[]IndexSearch**](IndexSearch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

