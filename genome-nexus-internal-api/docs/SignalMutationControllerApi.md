# \SignalMutationControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSignalMutationsByHgvsgGETUsingGET**](SignalMutationControllerApi.md#FetchSignalMutationsByHgvsgGETUsingGET) | **Get** /signal/mutation/hgvs/{variant} | Retrieves SignalDB mutations by hgvgs variant
[**FetchSignalMutationsByHugoSymbolGETUsingGET**](SignalMutationControllerApi.md#FetchSignalMutationsByHugoSymbolGETUsingGET) | **Get** /signal/mutation | Retrieves SignalDB mutations by Hugo Gene Symbol
[**FetchSignalMutationsByMutationFilterPOSTUsingPOST**](SignalMutationControllerApi.md#FetchSignalMutationsByMutationFilterPOSTUsingPOST) | **Post** /signal/mutation | Retrieves SignalDB mutations by Mutation Filter


# **FetchSignalMutationsByHgvsgGETUsingGET**
> []SignalMutation FetchSignalMutationsByHgvsgGETUsingGET(ctx, variant)
Retrieves SignalDB mutations by hgvgs variant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **variant** | **string**| A variant. For example 13:g.32890665G&gt;A | 

### Return type

[**[]SignalMutation**](SignalMutation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchSignalMutationsByHugoSymbolGETUsingGET**
> []SignalMutation FetchSignalMutationsByHugoSymbolGETUsingGET(ctx, optional)
Retrieves SignalDB mutations by Hugo Gene Symbol

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoGeneSymbol** | **string**| Hugo Symbol. For example BRCA1 | 

### Return type

[**[]SignalMutation**](SignalMutation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchSignalMutationsByMutationFilterPOSTUsingPOST**
> []SignalMutation FetchSignalMutationsByMutationFilterPOSTUsingPOST(ctx, mutationFilter)
Retrieves SignalDB mutations by Mutation Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mutationFilter** | [**SignalMutationFilter**](SignalMutationFilter.md)| List of Hugo Gene Symbols. For example [\&quot;TP53\&quot;, \&quot;PIK3CA\&quot;, \&quot;BRCA1\&quot;] | 

### Return type

[**[]SignalMutation**](SignalMutation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

