# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | ID of the integration. | [optional] [readonly] 
**IntegrationHubId** | **int64** | ID of the integration hub. | [optional] [readonly] 
**Label** | **string** | Label for the integration. | [optional] [readonly] 
**Description** | **string** | Description of the integration. | [optional] [readonly] 
**Enabled** | **bool** | Whether the integration is available to users. | [optional] 
**Params** | [**[]IntegrationParam**](IntegrationParam.md) | Array of params for the integration. | [optional] 
**SupportedFormats** | **[]string** | A list of data formats the integration supports. If unspecified, the default is all data formats. Valid values are: \&quot;txt\&quot;, \&quot;csv\&quot;, \&quot;inline_json\&quot;, \&quot;json\&quot;, \&quot;json_label\&quot;, \&quot;json_detail\&quot;, \&quot;json_detail_lite_stream\&quot;, \&quot;xlsx\&quot;, \&quot;html\&quot;, \&quot;wysiwyg_pdf\&quot;, \&quot;assembled_pdf\&quot;, \&quot;wysiwyg_png\&quot;, \&quot;csv_zip\&quot;. | [optional] [readonly] 
**SupportedActionTypes** | **[]string** | A list of action types the integration supports. Valid values are: \&quot;cell\&quot;, \&quot;query\&quot;, \&quot;dashboard\&quot;. | [optional] [readonly] 
**SupportedFormattings** | **[]string** | A list of formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \&quot;formatted\&quot;, \&quot;unformatted\&quot;. | [optional] [readonly] 
**SupportedVisualizationFormattings** | **[]string** | A list of visualization formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \&quot;apply\&quot;, \&quot;noapply\&quot;. | [optional] [readonly] 
**SupportedDownloadSettings** | **[]string** | A list of all the download mechanisms the integration supports. The order of values is not significant: Looker will select the most appropriate supported download mechanism for a given query. The integration must ensure it can handle any of the mechanisms it claims to support. If unspecified, this defaults to all download setting values. Valid values are: \&quot;push\&quot;, \&quot;url\&quot;. | [optional] [readonly] 
**IconUrl** | **string** | URL to an icon for the integration. | [optional] [readonly] 
**UsesOauth** | **bool** | Whether the integration uses oauth. | [optional] [readonly] 
**RequiredFields** | [**[]IntegrationRequiredField**](IntegrationRequiredField.md) | A list of descriptions of required fields that this integration is compatible with. If there are multiple entries in this list, the integration requires more than one field. If unspecified, no fields will be required. | [optional] [readonly] 
**DelegateOauth** | **bool** | Whether the integration uses delegate oauth, which allows federation between an integration installation scope specific entity (like org, group, and team, etc.) and Looker. | [optional] [readonly] 
**InstalledDelegateOauthTargets** | **[]string** | Whether the integration is available to users. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


