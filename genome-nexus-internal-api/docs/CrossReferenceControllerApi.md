# \CrossReferenceControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchGeneXrefsGET1**](CrossReferenceControllerApi.md#FetchGeneXrefsGET1) | **Get** /xrefs/{accession} | Perform lookups of Ensembl identifiers and retrieve their external references in other databases


# **FetchGeneXrefsGET1**
> []GeneXref FetchGeneXrefsGET1(ctx, accession)
Perform lookups of Ensembl identifiers and retrieve their external references in other databases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accession** | **string**| Ensembl gene accession. For example ENSG00000169083 | 

### Return type

[**[]GeneXref**](GeneXref.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

