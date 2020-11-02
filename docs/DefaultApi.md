# \DefaultApi

All URIs are relative to *https://lotw.arrl.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LotwuserLotwreportAdiGet**](DefaultApi.md#LotwuserLotwreportAdiGet) | **Get** /lotwuser/lotwreport.adi | The do-everything endpoint



## LotwuserLotwreportAdiGet

> string LotwuserLotwreportAdiGet(ctx, login, password, qsoQuery, optional)

The do-everything endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string**|  | 
**password** | **string**|  | 
**qsoQuery** | **string**|  | 
 **optional** | ***LotwuserLotwreportAdiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LotwuserLotwreportAdiGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **qsoQsl** | **optional.String**| If \&quot;yes\&quot;, only QSL records are returned | 
 **qsoQslsince** | **optional.String**| -| Returns QSL records received (matched or updated) on or after the specified date. Will also accept date/time in \&quot;YYYY-MM-DD HH:MM:SS\&quot; format. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-arrl-adif; charset=iso-8859-1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

