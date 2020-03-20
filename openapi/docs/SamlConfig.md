# SamlConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Enabled** | **bool** | Enable/Disable Saml authentication for the server | [optional] 
**IdpCert** | **string** | Identity Provider Certificate (provided by IdP) | [optional] 
**IdpUrl** | **string** | Identity Provider Url (provided by IdP) | [optional] 
**IdpIssuer** | **string** | Identity Provider Issuer (provided by IdP) | [optional] 
**IdpAudience** | **string** | Identity Provider Audience (set in IdP config). Optional in Looker. Set this only if you want Looker to validate the audience value returned by the IdP. | [optional] 
**AllowedClockDrift** | **int64** | Count of seconds of clock drift to allow when validating timestamps of assertions. | [optional] 
**UserAttributeMapEmail** | **string** | Name of user record attributes used to indicate email address field | [optional] 
**UserAttributeMapFirstName** | **string** | Name of user record attributes used to indicate first name | [optional] 
**UserAttributeMapLastName** | **string** | Name of user record attributes used to indicate last name | [optional] 
**NewUserMigrationTypes** | **string** | Merge first-time saml login to existing user account by email addresses. When a user logs in for the first time via saml this option will connect this user into their existing account by finding the account with a matching email address by testing the given types of credentials for existing users. Otherwise a new user account will be created for the user. This list (if provided) must be a comma separated list of string like &#39;email,ldap,google&#39; | [optional] 
**AlternateEmailLoginAllowed** | **bool** | Allow alternate email-based login via &#39;/login/email&#39; for admins and for specified users with the &#39;login_special_email&#39; permission. This option is useful as a fallback during ldap setup, if ldap config problems occur later, or if you need to support some users who are not in your ldap directory. Looker email/password logins are always disabled for regular users when ldap is enabled. | [optional] 
**TestSlug** | **string** | Slug to identify configurations that are created in order to run a Saml config test | [optional] [readonly] 
**ModifiedAt** | **string** | When this config was last modified | [optional] [readonly] 
**ModifiedBy** | **string** | User id of user who last modified this config | [optional] [readonly] 
**DefaultNewUserRoles** | [**[]Role**](Role.md) | (Read-only) Roles that will be applied to new users the first time they login via Saml | [optional] [readonly] 
**DefaultNewUserGroups** | [**[]Group**](Group.md) | (Read-only) Groups that will be applied to new users the first time they login via Saml | [optional] [readonly] 
**DefaultNewUserRoleIds** | **[]int64** | (Write-Only) Array of ids of roles that will be applied to new users the first time they login via Saml | [optional] 
**DefaultNewUserGroupIds** | **[]int64** | (Write-Only) Array of ids of groups that will be applied to new users the first time they login via Saml | [optional] 
**SetRolesFromGroups** | **bool** | Set user roles in Looker based on groups from Saml | [optional] 
**GroupsAttribute** | **string** | Name of user record attributes used to indicate groups. Used when &#39;groups_finder_type&#39; is set to &#39;grouped_attribute_values&#39; | [optional] 
**Groups** | [**[]SamlGroupRead**](SamlGroupRead.md) | (Read-only) Array of mappings between Saml Groups and Looker Roles | [optional] [readonly] 
**GroupsWithRoleIds** | [**[]SamlGroupWrite**](SamlGroupWrite.md) | (Read/Write) Array of mappings between Saml Groups and arrays of Looker Role ids | [optional] 
**AuthRequiresRole** | **bool** | Users will not be allowed to login at all unless a role for them is found in Saml if set to true | [optional] 
**UserAttributes** | [**[]SamlUserAttributeRead**](SamlUserAttributeRead.md) | (Read-only) Array of mappings between Saml User Attributes and Looker User Attributes | [optional] [readonly] 
**UserAttributesWithIds** | [**[]SamlUserAttributeWrite**](SamlUserAttributeWrite.md) | (Read/Write) Array of mappings between Saml User Attributes and arrays of Looker User Attribute ids | [optional] 
**GroupsFinderType** | **string** | Identifier for a strategy for how Looker will find groups in the SAML response. One of [&#39;grouped_attribute_values&#39;, &#39;individual_attributes&#39;] | [optional] 
**GroupsMemberValue** | **string** | Value for group attribute used to indicate membership. Used when &#39;groups_finder_type&#39; is set to &#39;individual_attributes&#39; | [optional] 
**BypassLoginPage** | **bool** | Bypass the login page when user authentication is required. Redirect to IdP immediately instead. | [optional] 
**Url** | **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


