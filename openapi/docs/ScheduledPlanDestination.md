# ScheduledPlanDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique Id | [optional] [readonly] 
**ScheduledPlanId** | **int64** | Id of a scheduled plan you own | [optional] 
**Format** | **string** | The data format to send to the given destination. Supported formats vary by destination, but include: \&quot;txt\&quot;, \&quot;csv\&quot;, \&quot;inline_json\&quot;, \&quot;json\&quot;, \&quot;json_detail\&quot;, \&quot;xlsx\&quot;, \&quot;html\&quot;, \&quot;wysiwyg_pdf\&quot;, \&quot;assembled_pdf\&quot;, \&quot;wysiwyg_png\&quot; | [optional] 
**ApplyFormatting** | **bool** | Are values formatted? (containing currency symbols, digit separators, etc. | [optional] 
**ApplyVis** | **bool** | Whether visualization options are applied to the results. | [optional] 
**Address** | **string** | Address for recipient. For email e.g. &#39;user@example.com&#39;. For webhooks e.g. &#39;https://domain/path&#39;. For Amazon S3 e.g. &#39;s3://bucket-name/path/&#39;. For SFTP e.g. &#39;sftp://host-name/path/&#39;.  | [optional] 
**LookerRecipient** | **bool** | Whether the recipient is a Looker user on the current instance (only applicable for email recipients) | [optional] [readonly] 
**Type** | **string** | Type of the address (&#39;email&#39;, &#39;webhook&#39;, &#39;s3&#39;, or &#39;sftp&#39;) | [optional] 
**Parameters** | **string** | JSON object containing parameters for external scheduling. For Amazon S3, this requires keys and values for access_key_id and region. For SFTP, this requires a key and value for username. | [optional] 
**SecretParameters** | **string** | (Write-Only) JSON object containing secret parameters for external scheduling. For Amazon S3, this requires a key and value for secret_access_key. For SFTP, this requires a key and value for password. | [optional] 
**Message** | **string** | Optional message to be included in scheduled emails | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


