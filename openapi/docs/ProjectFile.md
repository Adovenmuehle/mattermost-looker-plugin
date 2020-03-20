# ProjectFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | An opaque token uniquely identifying a file within a project. Avoid parsing or decomposing the text of this token. This token is stable within a Looker release but may change between Looker releases | [optional] [readonly] 
**Path** | **string** | Path, file name, and extension of the file relative to the project root directory | [optional] [readonly] 
**Title** | **string** | Display name | [optional] [readonly] 
**Type** | **string** | File type: model, view, etc | [optional] [readonly] 
**Extension** | **string** | The extension of the file: .view.lkml, .model.lkml, etc | [optional] [readonly] 
**MimeType** | **string** | File mime type | [optional] [readonly] 
**Editable** | **bool** | State of editability for the file. | [optional] [readonly] 
**GitStatus** | [**GitStatus**](GitStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


