# DbConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Name** | **string** | Name of the connection. Also used as the unique identifier | [optional] 
**Dialect** | [**Dialect**](Dialect.md) |  | [optional] 
**Snippets** | [**[]Snippet**](Snippet.md) | SQL Runner snippets for this connection | [optional] [readonly] 
**Host** | **string** | Host name/address of server | [optional] 
**Port** | **string** | Port number on server | [optional] 
**Username** | **string** | Username for server authentication | [optional] 
**Password** | **string** | (Write-Only) Password for server authentication | [optional] 
**UsesOauth** | **bool** | Whether the connection uses OAuth for authentication. | [optional] [readonly] 
**Certificate** | **string** | (Write-Only) Base64 encoded Certificate body for server authentication (when appropriate for dialect). | [optional] 
**FileType** | **string** | (Write-Only) Certificate keyfile type - .json or .p12 | [optional] 
**Database** | **string** | Database name | [optional] 
**DbTimezone** | **string** | Time zone of database | [optional] 
**QueryTimezone** | **string** | Timezone to use in queries | [optional] 
**Schema** | **string** | Scheme name | [optional] 
**MaxConnections** | **int64** | Maximum number of concurrent connection to use | [optional] 
**MaxBillingGigabytes** | **string** | Maximum size of query in GBs (BigQuery only, can be a user_attribute name) | [optional] 
**Ssl** | **bool** | Use SSL/TLS when connecting to server | [optional] 
**VerifySsl** | **bool** | Verify the SSL | [optional] 
**TmpDbName** | **string** | Name of temporary database (if used) | [optional] 
**JdbcAdditionalParams** | **string** | Additional params to add to JDBC connection string | [optional] 
**PoolTimeout** | **int64** | Connection Pool Timeout, in seconds | [optional] 
**DialectName** | **string** | (Read/Write) SQL Dialect name | [optional] 
**CreatedAt** | **string** | Creation date for this connection | [optional] [readonly] 
**UserId** | **string** | Id of user who last modified this connection configuration | [optional] [readonly] 
**Example** | **bool** | Is this an example connection? | [optional] [readonly] 
**UserDbCredentials** | **bool** | (Limited access feature) Are per user db credentials enabled. Enabling will remove previously set username and password | [optional] 
**UserAttributeFields** | **[]string** | Fields whose values map to user attribute names | [optional] 
**MaintenanceCron** | **string** | Cron string specifying when maintenance such as PDT trigger checks and drops should be performed | [optional] 
**LastRegenAt** | **string** | Unix timestamp at start of last completed PDT trigger check process | [optional] [readonly] 
**LastReapAt** | **string** | Unix timestamp at start of last completed PDT reap process | [optional] [readonly] 
**SqlRunnerPrecacheTables** | **bool** | Precache tables in the SQL Runner | [optional] 
**AfterConnectStatements** | **string** | SQL statements (semicolon separated) to issue after connecting to the database. Requires &#x60;custom_after_connect_statements&#x60; license feature | [optional] 
**PdtContextOverride** | [**DbConnectionOverride**](DBConnectionOverride.md) |  | [optional] 
**Managed** | **bool** | Is this connection created and managed by Looker | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


