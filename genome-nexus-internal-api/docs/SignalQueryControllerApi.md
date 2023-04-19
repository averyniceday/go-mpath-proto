# \SignalQueryControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSignalByKeywordGETUsingGET**](SignalQueryControllerApi.md#SearchSignalByKeywordGETUsingGET) | **Get** /signal/search | Performs search by gene, protein change, variant or region.


# **SearchSignalByKeywordGETUsingGET**
> []SignalQuery SearchSignalByKeywordGETUsingGET(ctx, keyword, optional)
Performs search by gene, protein change, variant or region.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyword** | **string**| keyword. For example BRCA; BRAF V600; 13:32906640-32906640; 13:g.32890665G&gt;A | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string**| keyword. For example BRCA; BRAF V600; 13:32906640-32906640; 13:g.32890665G&gt;A | 
 **limit** | **int32**| Max number of matching results to return | 

### Return type

[**[]SignalQuery**](SignalQuery.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

