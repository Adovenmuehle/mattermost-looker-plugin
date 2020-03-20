# LookmlModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**AllowedDbConnectionNames** | **[]string** | Array of names of connections this model is allowed to use | [optional] 
**Explores** | [**[]LookmlModelNavExplore**](LookmlModelNavExplore.md) | Array of explores (if has_content) | [optional] [readonly] 
**HasContent** | **bool** | Does this model declaration have have lookml content? | [optional] [readonly] 
**Label** | **string** | UI-friendly name for this model | [optional] [readonly] 
**Name** | **string** | Name of the model. Also used as the unique identifier | [optional] 
**ProjectName** | **string** | Name of project containing the model | [optional] 
**UnlimitedDbConnections** | **bool** | Is this model allowed to use all current and future connections | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


