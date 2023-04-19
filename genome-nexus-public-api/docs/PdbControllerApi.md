# \PdbControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPdbHeaderGET**](PdbControllerApi.md#FetchPdbHeaderGET) | **Get** /pdb/header/{pdbId} | Retrieves PDB header info by a PDB id
[**FetchPdbHeaderPOST**](PdbControllerApi.md#FetchPdbHeaderPOST) | **Post** /pdb/header | Retrieves PDB header info by a PDB id


# **FetchPdbHeaderGET**
> PdbHeader FetchPdbHeaderGET(ctx, pdbId)
Retrieves PDB header info by a PDB id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pdbId** | **string**| PDB id, for example 1a37 | 

### Return type

[**PdbHeader**](PdbHeader.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchPdbHeaderPOST**
> []PdbHeader FetchPdbHeaderPOST(ctx, pdbIds)
Retrieves PDB header info by a PDB id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pdbIds** | **[]string**| List of pdb ids, for example [\&quot;1a37\&quot;,\&quot;1a4o\&quot;] | 

### Return type

[**[]PdbHeader**](PdbHeader.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

