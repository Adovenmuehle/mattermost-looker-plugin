# Dialect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the dialect | [optional] [readonly] 
**Label** | **string** | The human-readable label of the connection | [optional] [readonly] 
**SupportsCostEstimate** | **bool** | Whether the dialect supports query cost estimates | [optional] [readonly] 
**PersistentTableIndexes** | **string** | PDT index columns | [optional] [readonly] 
**PersistentTableSortkeys** | **string** | PDT sortkey columns | [optional] [readonly] 
**PersistentTableDistkey** | **string** | PDT distkey column | [optional] [readonly] 
**SupportsStreaming** | **bool** | Suports streaming results | [optional] [readonly] 
**AutomaticallyRunSqlRunnerSnippets** | **bool** | Should SQL Runner snippets automatically be run | [optional] [readonly] 
**ConnectionTests** | **[]string** | Array of names of the tests that can be run on a connection using this dialect | [optional] [readonly] 
**SupportsInducer** | **bool** | Is supported with the inducer (i.e. generate from sql) | [optional] [readonly] 
**SupportsMultipleDatabases** | **bool** | Can multiple databases be accessed from a connection using this dialect | [optional] [readonly] 
**SupportsPersistentDerivedTables** | **bool** | Whether the dialect supports allowing Looker to build persistent derived tables | [optional] [readonly] 
**HasSslSupport** | **bool** | Does the database have client SSL support settable through the JDBC string explicitly? | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


