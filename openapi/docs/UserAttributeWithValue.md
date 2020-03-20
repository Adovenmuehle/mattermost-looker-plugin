# UserAttributeWithValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Name** | **string** | Name of user attribute | [optional] [readonly] 
**Label** | **string** | Human-friendly label for user attribute | [optional] [readonly] 
**Rank** | **int64** | Precedence for setting value on user (lowest wins) | [optional] [readonly] 
**Value** | **string** | Value of attribute for user | [optional] 
**UserId** | **int64** | Id of User | [optional] [readonly] 
**UserCanEdit** | **bool** | Can the user set this value | [optional] [readonly] 
**ValueIsHidden** | **bool** | If true, the \&quot;value\&quot; field will be null, because the attribute settings block access to this value | [optional] [readonly] 
**UserAttributeId** | **int64** | Id of User Attribute | [optional] [readonly] 
**Source** | **string** | How user got this value for this attribute | [optional] [readonly] 
**HiddenValueDomainWhitelist** | **string** | If this user attribute is hidden, whitelist of destinations to which it may be sent. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


