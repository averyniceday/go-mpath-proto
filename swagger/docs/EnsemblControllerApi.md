# \EnsemblControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET) | **Get** /ensembl/canonical-gene/entrez/{entrezGeneId} | Retrieves Ensembl canonical gene id by Entrez Gene Id
[**FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST) | **Post** /ensembl/canonical-gene/entrez | Retrieves canonical Ensembl Gene ID by Entrez Gene Ids
[**FetchCanonicalEnsemblGeneIdByHugoSymbolGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByHugoSymbolGET) | **Get** /ensembl/canonical-gene/hgnc/{hugoSymbol} | Retrieves Ensembl canonical gene id by Hugo Symbol
[**FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST) | **Post** /ensembl/canonical-gene/hgnc | Retrieves canonical Ensembl Gene ID by Hugo Symbols
[**FetchCanonicalEnsemblTranscriptByHugoSymbolGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblTranscriptByHugoSymbolGET) | **Get** /ensembl/canonical-transcript/hgnc/{hugoSymbol} | Retrieves Ensembl canonical transcript by Hugo Symbol
[**FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST) | **Post** /ensembl/canonical-transcript/hgnc | Retrieves Ensembl canonical transcripts by Hugo Symbols
[**FetchEnsemblTranscriptByTranscriptIdGET**](EnsemblControllerApi.md#FetchEnsemblTranscriptByTranscriptIdGET) | **Get** /ensembl/transcript/{transcriptId} | Retrieves the transcript by an Ensembl transcript ID
[**FetchEnsemblTranscriptsByEnsemblFilterPOST**](EnsemblControllerApi.md#FetchEnsemblTranscriptsByEnsemblFilterPOST) | **Post** /ensembl/transcript | Retrieves Ensembl Transcripts by Ensembl transcript IDs, hugo Symbols, protein IDs, or gene IDs
[**FetchEnsemblTranscriptsGET**](EnsemblControllerApi.md#FetchEnsemblTranscriptsGET) | **Get** /ensembl/transcript | Retrieves Ensembl Transcripts by protein ID, and gene ID. Retrieves all transcripts in case no query parameter provided
[**FetchGeneXrefsGET**](EnsemblControllerApi.md#FetchGeneXrefsGET) | **Get** /ensembl/xrefs | Perform lookups of Ensembl identifiers and retrieve their external references in other databases


# **FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET**
> EnsemblGene FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET(ctx, entrezGeneId)
Retrieves Ensembl canonical gene id by Entrez Gene Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **entrezGeneId** | **string**| An Entrez Gene Id. For example 23140 | 

### Return type

[**EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST**
> []EnsemblGene FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST(ctx, entrezGeneIds)
Retrieves canonical Ensembl Gene ID by Entrez Gene Ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **entrezGeneIds** | **[]string**| List of Entrez Gene Ids. For example [\&quot;23140\&quot;,\&quot;26009\&quot;,\&quot;100131879\&quot;] | 

### Return type

[**[]EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCanonicalEnsemblGeneIdByHugoSymbolGET**
> EnsemblGene FetchCanonicalEnsemblGeneIdByHugoSymbolGET(ctx, hugoSymbol)
Retrieves Ensembl canonical gene id by Hugo Symbol

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hugoSymbol** | **string**| A Hugo Symbol. For example TP53 | 

### Return type

[**EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST**
> []EnsemblGene FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST(ctx, hugoSymbols)
Retrieves canonical Ensembl Gene ID by Hugo Symbols

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hugoSymbols** | **[]string**| List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;] | 

### Return type

[**[]EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCanonicalEnsemblTranscriptByHugoSymbolGET**
> EnsemblTranscript FetchCanonicalEnsemblTranscriptByHugoSymbolGET(ctx, hugoSymbol, optional)
Retrieves Ensembl canonical transcript by Hugo Symbol

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hugoSymbol** | **string**| A Hugo Symbol. For example TP53 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoSymbol** | **string**| A Hugo Symbol. For example TP53 | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | [default to uniprot]

### Return type

[**EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST**
> []EnsemblTranscript FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST(ctx, hugoSymbols, optional)
Retrieves Ensembl canonical transcripts by Hugo Symbols

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hugoSymbols** | **[]string**| List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;] | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoSymbols** | **[]string**| List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;] | 
 **isoformOverrideSource** | **string**| Isoform override source. For example uniprot | [default to uniprot]

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchEnsemblTranscriptByTranscriptIdGET**
> EnsemblTranscript FetchEnsemblTranscriptByTranscriptIdGET(ctx, transcriptId)
Retrieves the transcript by an Ensembl transcript ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transcriptId** | **string**| An Ensembl transcript ID. For example ENST00000361390 | 

### Return type

[**EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchEnsemblTranscriptsByEnsemblFilterPOST**
> []EnsemblTranscript FetchEnsemblTranscriptsByEnsemblFilterPOST(ctx, ensemblFilter)
Retrieves Ensembl Transcripts by Ensembl transcript IDs, hugo Symbols, protein IDs, or gene IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ensemblFilter** | [**EnsemblFilter**](EnsemblFilter.md)| List of Ensembl transcript IDs. For example [\&quot;ENST00000361390\&quot;, \&quot;ENST00000361453\&quot;, \&quot;ENST00000361624\&quot;]&lt;br&gt;OR&lt;br&gt;List of Hugo Symbols. For example [\&quot;TP53\&quot;, \&quot;PIK3CA\&quot;, \&quot;BRCA1\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl protein IDs. For example [\&quot;ENSP00000439985\&quot;, \&quot;ENSP00000478460\&quot;, \&quot;ENSP00000346196\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl gene IDs. For example [\&quot;ENSG00000136999\&quot;, \&quot;ENSG00000272398\&quot;, \&quot;ENSG00000198695\&quot;] | 

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchEnsemblTranscriptsGET**
> []EnsemblTranscript FetchEnsemblTranscriptsGET(ctx, optional)
Retrieves Ensembl Transcripts by protein ID, and gene ID. Retrieves all transcripts in case no query parameter provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **geneId** | **string**| An Ensembl gene ID. For example ENSG00000136999 | 
 **proteinId** | **string**| An Ensembl protein ID. For example ENSP00000439985 | 
 **hugoSymbol** | **string**| A Hugo Symbol For example ARF5 | 

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchGeneXrefsGET**
> []GeneXref FetchGeneXrefsGET(ctx, accession)
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

