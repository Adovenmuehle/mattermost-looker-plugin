# DataActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** | ID of the webhook event that sent this data action. In some error conditions, this may be null. | [optional] [readonly] 
**Success** | **bool** | Whether the data action was successful. | [optional] [readonly] 
**RefreshQuery** | **bool** | When true, indicates that the client should refresh (rerun) the source query because the data may have been changed by the action. | [optional] [readonly] 
**ValidationErrors** | [**ValidationError**](ValidationError.md) |  | [optional] 
**Message** | **string** | Optional message returned by the data action server describing the state of the action that took place. This can be used to implement custom failure messages. If a failure is related to a particular form field, the server should send back a validation error instead. The Looker web UI does not currently display any message if the action indicates &#39;success&#39;, but may do so in the future. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


