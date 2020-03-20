# EmbedSsoUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | The complete URL of the Looker UI page to display in the embed context. For example, to display the dashboard with id 34, &#x60;target_url&#x60; would look like: &#x60;https://mycompany.looker.com:9999/dashboards/34&#x60;. &#x60;target_uri&#x60; MUST contain a scheme (HTTPS), domain name, and URL path. Port must be included if it is required to reach the Looker server from browser clients. If the Looker instance is behind a load balancer or other proxy, &#x60;target_uri&#x60; must be the public-facing domain name and port required to reach the Looker instance, not the actual internal network machine name of the Looker instance. | 
**SessionLength** | **int64** | Number of seconds the SSO embed session will be valid after the embed session is started. Defaults to 300 seconds. Maximum session length accepted is 2592000 seconds (30 days). | [optional] 
**ExternalUserId** | **string** | A value from an external system that uniquely identifies the embed user. Since the user_ids of Looker embed users may change with every embed session, external_user_id provides a way to assign a known, stable user identifier across multiple embed sessions. | [optional] 
**ForceLogoutLogin** | **bool** | When true, the embed session will purge any residual Looker login state (such as in browser cookies) before creating a new login state with the given embed user info. Defaults to true. | [optional] 
**FirstName** | **string** | First name of the embed user. Defaults to &#39;Embed&#39; if not specified | [optional] 
**LastName** | **string** | Last name of the embed user. Defaults to &#39;User&#39; if not specified | [optional] 
**UserTimezone** | **string** | Sets the user timezone for the embed user session, if the User Specific Timezones setting is enabled in the Looker admin settings. A value of &#x60;null&#x60; forces the embed user to use the Looker Application Default Timezone. You MUST omit this property from the request if the User Specific Timezones setting is disabled. Timezone values are validated against the IANA Timezone standard and can be seen in the Application Time Zone dropdown list on the Looker General Settings admin page. | [optional] 
**Permissions** | **[]string** | List of Looker permission names to grant to the embed user. Requested permissions will be filtered to permissions allowed for embed sessions. | [optional] 
**Models** | **[]string** | List of model names that the embed user may access | [optional] 
**GroupIds** | **[]int64** | List of Looker group ids in which to enroll the embed user | [optional] 
**ExternalGroupId** | **int64** | A unique value identifying an embed-exclusive group. Multiple embed users using the same &#x60;external_group_id&#x60; value will be able to share Looker content with each other. Content and embed users associated with the &#x60;external_group_id&#x60; will not be accessible to normal Looker users or embed users not associated with this &#x60;external_group_id&#x60;. | [optional] 
**UserAttributes** | **map[string]string** | A dictionary of name-value pairs associating a Looker user attribute name with a value. | [optional] 
**SecretId** | **int64** | Id of the embed secret to use to sign this SSO url. If specified, the value must be an id of a valid (active) secret defined in the Looker instance. If not specified, the URL will be signed with the newest active embed secret defined in the Looker instance. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


