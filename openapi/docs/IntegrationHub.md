# IntegrationHub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | ID of the hub. | [optional] [readonly] 
**Url** | **string** | URL of the hub. | [optional] 
**Label** | **string** | Label of the hub. | [optional] [readonly] 
**Official** | **bool** | Whether this hub is a first-party integration hub operated by Looker. | [optional] [readonly] 
**FetchErrorMessage** | **string** | An error message, present if the integration hub metadata could not be fetched. If this is present, the integration hub is unusable. | [optional] [readonly] 
**AuthorizationToken** | **string** | (Write-Only) An authorization key that will be sent to the integration hub on every request. | [optional] 
**HasAuthorizationToken** | **bool** | Whether the authorization_token is set for the hub. | [optional] [readonly] 
**LegalAgreementSigned** | **bool** | Whether the legal agreement message has been signed by the user. This only matters if legal_agreement_required is true. | [optional] [readonly] 
**LegalAgreementRequired** | **bool** | Whether the legal terms for the integration hub are required before use. | [optional] [readonly] 
**LegalAgreementText** | **string** | The legal agreement text for this integration hub. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


