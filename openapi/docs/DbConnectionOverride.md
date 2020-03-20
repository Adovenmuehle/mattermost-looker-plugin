# DbConnectionOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Context in which to override (&#x60;pdt&#x60; is the only allowed value) | [optional] 
**Host** | **string** | Host name/address of server | [optional] 
**Port** | **string** | Port number on server | [optional] 
**Username** | **string** | Username for server authentication | [optional] 
**Password** | **string** | (Write-Only) Password for server authentication | [optional] 
**HasPassword** | **bool** | Whether or not the password is overridden in this context | [optional] [readonly] 
**Certificate** | **string** | (Write-Only) Base64 encoded Certificate body for server authentication (when appropriate for dialect). | [optional] 
**FileType** | **string** | (Write-Only) Certificate keyfile type - .json or .p12 | [optional] 
**Database** | **string** | Database name | [optional] 
**Schema** | **string** | Scheme name | [optional] 
**JdbcAdditionalParams** | **string** | Additional params to add to JDBC connection string | [optional] 
**AfterConnectStatements** | **string** | SQL statements (semicolon separated) to issue after connecting to the database. Requires &#x60;custom_after_connect_statements&#x60; license feature | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


