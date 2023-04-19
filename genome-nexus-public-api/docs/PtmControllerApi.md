# \PtmControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPostTranslationalModificationsByPtmFilterPOST**](PtmControllerApi.md#FetchPostTranslationalModificationsByPtmFilterPOST) | **Post** /ptm/experimental | Retrieves PTM entries by Ensembl Transcript IDs
[**FetchPostTranslationalModificationsGET**](PtmControllerApi.md#FetchPostTranslationalModificationsGET) | **Get** /ptm/experimental | Retrieves PTM entries by Ensembl Transcript ID


# **FetchPostTranslationalModificationsByPtmFilterPOST**
> []PostTranslationalModification FetchPostTranslationalModificationsByPtmFilterPOST(ctx, ptmFilter)
Retrieves PTM entries by Ensembl Transcript IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ptmFilter** | [**PtmFilter**](PtmFilter.md)| List of Ensembl transcript IDs. For example [\&quot;ENST00000420316\&quot;, \&quot;ENST00000646891\&quot;, \&quot;ENST00000371953\&quot;] | 

### Return type

[**[]PostTranslationalModification**](PostTranslationalModification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchPostTranslationalModificationsGET**
> []PostTranslationalModification FetchPostTranslationalModificationsGET(ctx, optional)
Retrieves PTM entries by Ensembl Transcript ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ensemblTranscriptId** | **string**| Ensembl Transcript ID. For example ENST00000646891 | 

### Return type

[**[]PostTranslationalModification**](PostTranslationalModification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

