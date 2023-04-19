# \PfamControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPfamDomainsByAccessionGET**](PfamControllerApi.md#FetchPfamDomainsByAccessionGET) | **Get** /pfam/domain/{pfamAccession} | Retrieves a PFAM domain by a PFAM domain ID
[**FetchPfamDomainsByPfamAccessionPOST**](PfamControllerApi.md#FetchPfamDomainsByPfamAccessionPOST) | **Post** /pfam/domain | Retrieves PFAM domains by PFAM domain accession IDs


# **FetchPfamDomainsByAccessionGET**
> PfamDomain FetchPfamDomainsByAccessionGET(ctx, pfamAccession)
Retrieves a PFAM domain by a PFAM domain ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pfamAccession** | **string**| A PFAM domain accession ID. For example PF02827 | 

### Return type

[**PfamDomain**](PfamDomain.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchPfamDomainsByPfamAccessionPOST**
> []PfamDomain FetchPfamDomainsByPfamAccessionPOST(ctx, pfamAccessions)
Retrieves PFAM domains by PFAM domain accession IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pfamAccessions** | **[]string**| List of PFAM domain accession IDs. For example [\&quot;PF02827\&quot;,\&quot;PF00093\&quot;,\&quot;PF15276\&quot;] | 

### Return type

[**[]PfamDomain**](PfamDomain.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

