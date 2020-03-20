# LdapConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**AlternateEmailLoginAllowed** | **bool** | Allow alternate email-based login via &#39;/login/email&#39; for admins and for specified users with the &#39;login_special_email&#39; permission. This option is useful as a fallback during ldap setup, if ldap config problems occur later, or if you need to support some users who are not in your ldap directory. Looker email/password logins are always disabled for regular users when ldap is enabled. | [optional] 
**AuthPassword** | **string** | (Write-Only)  Password for the LDAP account used to access the LDAP server | [optional] 
**AuthRequiresRole** | **bool** | Users will not be allowed to login at all unless a role for them is found in LDAP if set to true | [optional] 
**AuthUsername** | **string** | Distinguished name of LDAP account used to access the LDAP server | [optional] 
**ConnectionHost** | **string** | LDAP server hostname | [optional] 
**ConnectionPort** | **string** | LDAP host port | [optional] 
**ConnectionTls** | **bool** | Use Transport Layer Security | [optional] 
**ConnectionTlsNoVerify** | **bool** | Do not verify peer when using TLS | [optional] 
**DefaultNewUserGroupIds** | **[]int64** | (Write-Only)  Array of ids of groups that will be applied to new users the first time they login via LDAP | [optional] 
**DefaultNewUserGroups** | [**[]Group**](Group.md) | (Read-only) Groups that will be applied to new users the first time they login via LDAP | [optional] [readonly] 
**DefaultNewUserRoleIds** | **[]int64** | (Write-Only)  Array of ids of roles that will be applied to new users the first time they login via LDAP | [optional] 
**DefaultNewUserRoles** | [**[]Role**](Role.md) | (Read-only) Roles that will be applied to new users the first time they login via LDAP | [optional] [readonly] 
**Enabled** | **bool** | Enable/Disable LDAP authentication for the server | [optional] 
**ForceNoPage** | **bool** | Don&#39;t attempt to do LDAP search result paging (RFC 2696) even if the LDAP server claims to support it. | [optional] 
**Groups** | [**[]LdapGroupRead**](LDAPGroupRead.md) | (Read-only) Array of mappings between LDAP Groups and Looker Roles | [optional] [readonly] 
**GroupsBaseDn** | **string** | Base dn for finding groups in LDAP searches | [optional] 
**GroupsFinderType** | **string** | Identifier for a strategy for how Looker will search for groups in the LDAP server | [optional] 
**GroupsMemberAttribute** | **string** | LDAP Group attribute that signifies the members of the groups. Most commonly &#39;member&#39; | [optional] 
**GroupsObjectclasses** | **string** | Optional comma-separated list of supported LDAP objectclass for groups when doing groups searches | [optional] 
**GroupsUserAttribute** | **string** | LDAP Group attribute that signifies the user in a group. Most commonly &#39;dn&#39; | [optional] 
**GroupsWithRoleIds** | [**[]LdapGroupWrite**](LDAPGroupWrite.md) | (Read/Write) Array of mappings between LDAP Groups and arrays of Looker Role ids | [optional] 
**HasAuthPassword** | **bool** | (Read-only) Has the password been set for the LDAP account used to access the LDAP server | [optional] [readonly] 
**MergeNewUsersByEmail** | **bool** | Merge first-time ldap login to existing user account by email addresses. When a user logs in for the first time via ldap this option will connect this user into their existing account by finding the account with a matching email address. Otherwise a new user account will be created for the user. | [optional] 
**ModifiedAt** | **string** | When this config was last modified | [optional] [readonly] 
**ModifiedBy** | **string** | User id of user who last modified this config | [optional] [readonly] 
**SetRolesFromGroups** | **bool** | Set user roles in Looker based on groups from LDAP | [optional] 
**TestLdapPassword** | **string** | (Write-Only)  Test LDAP user password. For ldap tests only. | [optional] 
**TestLdapUser** | **string** | (Write-Only)  Test LDAP user login id. For ldap tests only. | [optional] 
**UserAttributeMapEmail** | **string** | Name of user record attributes used to indicate email address field | [optional] 
**UserAttributeMapFirstName** | **string** | Name of user record attributes used to indicate first name | [optional] 
**UserAttributeMapLastName** | **string** | Name of user record attributes used to indicate last name | [optional] 
**UserAttributeMapLdapId** | **string** | Name of user record attributes used to indicate unique record id | [optional] 
**UserAttributes** | [**[]LdapUserAttributeRead**](LDAPUserAttributeRead.md) | (Read-only) Array of mappings between LDAP User Attributes and Looker User Attributes | [optional] [readonly] 
**UserAttributesWithIds** | [**[]LdapUserAttributeWrite**](LDAPUserAttributeWrite.md) | (Read/Write) Array of mappings between LDAP User Attributes and arrays of Looker User Attribute ids | [optional] 
**UserBindBaseDn** | **string** | Distinguished name of LDAP node used as the base for user searches | [optional] 
**UserCustomFilter** | **string** | (Optional) Custom RFC-2254 filter clause for use in finding user during login. Combined via &#39;and&#39; with the other generated filter clauses. | [optional] 
**UserIdAttributeNames** | **string** | Name(s) of user record attributes used for matching user login id (comma separated list) | [optional] 
**UserObjectclass** | **string** | (Optional) Name of user record objectclass used for finding user during login id | [optional] 
**Url** | **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


