# Homepage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentMetadataId** | **int64** | Id of associated content_metadata record | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Date of homepage creatopm | [optional] [readonly] 
**DeletedAt** | [**time.Time**](time.Time.md) | Date of homepage deletion | [optional] 
**Description** | **string** | Description of homepage | [optional] 
**HomepageSections** | [**[]HomepageSection**](HomepageSection.md) | Sections of the homepage | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**SectionOrder** | **[]int64** | ids of the homepage sections in the order they should be displayed | [optional] 
**Title** | **string** | Title of homepage | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Date of last homepage update | [optional] [readonly] 
**UserId** | **int64** | User id of homepage creator | [optional] [readonly] 
**PrimaryHomepage** | **bool** | Whether the homepage is the primary homepage or not | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


