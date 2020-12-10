# \DefaultApi

All URIs are relative to *https://lotw.arrl.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query**](DefaultApi.md#Query) | **Get** /lotwuser/lotwreport.adi | Querying LoTW for Acceptance and Confirmation of Submitted QSOs



## Query

> string Query(ctx, login, password, qsoQuery, optional)

Querying LoTW for Acceptance and Confirmation of Submitted QSOs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string**| Note that while the user&#39;s primary call sign is usually the username, this is not always the case and should not be assumed. | 
**password** | **string**| The user&#39;s plaintext LotW password. | 
**qsoQuery** | **int32**| If absent, ADIF file will contain no QSO records | 
 **optional** | ***QueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **qsoQsl** | [**optional.Interface of YesOrNo**](.md)| If \&quot;yes\&quot;, only QSL records are returned | 
 **qsoQslsince** | **optional.String**| Returns QSL records received (matched or updated) on or after the specified date. Will also accept date/time in \&quot;YYYY-MM-DD HH:MM:SS\&quot; format. Ignored unless qso_qsl&#x3D;\&quot;yes\&quot;. | 
 **qsoQsorxsince** | **optional.String**| Returns QSO records received (uploaded) on or after the specified date. Will also accept date/time in \&quot;YYYY-MM-DD HH:MM:SS\&quot; format. Ignored unless qso_qsl&#x3D;\&quot;no\&quot;. | 
 **qsoOwncall** | **optional.String**| Returns only records whose \&quot;own\&quot; call sign matches. | 
 **qsoCallsign** | **optional.String**| Returns only records whose \&quot;worked\&quot; call sign matches. | 
 **qsoMode** | **optional.String**| Returns only records whose mode matches. Mode must be one of the allowed modes. | 
 **qsoBand** | **optional.String**| Returns only records whose band matches. Mode must be one of the allowed bands. | 
 **qsoDxcc** | **optional.Int32**| Returns only records whose DXCC entity matches. (This implies qso_qsl&#x3D;\&quot;yes\&quot; since the DXCC entity of un-QSL&#39;d stations isn&#39;t known to LoTW.) Value must be the ARRL DXCC entity number. | 
 **qsoStartdate** | **optional.String**| Returns only records with a QSO date on or after the specified value. | 
 **qsoStarttime** | **optional.String**| Returns only records with a QSO time at or after the specified value on the starting date. This value is ignored if qso_startdate is not provided. | 
 **qsoEnddate** | **optional.String**| Returns only records with a QSO date on or before the specified value. | 
 **qsoEndtime** | **optional.String**| Returns only records with a QSO time at or before the specified value on the ending date. This value is ignored if qso_enddate is not provided. | 
 **qsoMydetail** | [**optional.Interface of YesOrNo**](.md)| If \&quot;yes\&quot;, returns fields that contain the Logging station&#39;s location data, if any. | 
 **qsoQsldetail** | [**optional.Interface of YesOrNo**](.md)| If \&quot;yes\&quot;, returns fields that contain the QSLing station&#39;s location data, if any. | 
 **qsoWithown** | [**optional.Interface of YesOrNo**](.md)| If \&quot;yes\&quot;, each record contains the STATION_CALLSIGN and APP_LoTW_OWNCALL fields to identify the \&quot;own\&quot; call sign used for the QSO. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-arrl-adif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

