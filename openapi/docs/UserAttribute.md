# UserAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Name** | **string** | Name of user attribute | [optional] 
**Label** | **string** | Human-friendly label for user attribute | [optional] 
**Type** | **string** | Type of user attribute (\&quot;string\&quot;, \&quot;number\&quot;, \&quot;datetime\&quot;, \&quot;yesno\&quot;, \&quot;zipcode\&quot;) | [optional] 
**DefaultValue** | **string** | Default value for when no value is set on the user | [optional] 
**IsSystem** | **bool** | Attribute is a system default | [optional] [readonly] 
**IsPermanent** | **bool** | Attribute is permanent and cannot be deleted | [optional] [readonly] 
**ValueIsHidden** | **bool** | If true, users will not be able to view values of this attribute | [optional] 
**UserCanView** | **bool** | Non-admin users can see the values of their attributes and use them in filters | [optional] 
**UserCanEdit** | **bool** | Users can change the value of this attribute for themselves | [optional] 
**HiddenValueDomainWhitelist** | **string** | Destinations to which a hidden attribute may be sent. Once set, cannot be edited. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


