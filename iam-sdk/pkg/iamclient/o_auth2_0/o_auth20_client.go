// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new o auth2 0 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for o auth2 0 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AuthCodeRequestV3(params *AuthCodeRequestV3Params, authInfo runtime.ClientAuthInfoWriter) (*AuthCodeRequestV3Found, error)

	AuthorizeV3(params *AuthorizeV3Params, authInfo runtime.ClientAuthInfoWriter) (*AuthorizeV3Found, error)

	GetJWKSV3(params *GetJWKSV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetJWKSV3OK, error)

	GetRevocationListV3(params *GetRevocationListV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetRevocationListV3OK, *GetRevocationListV3Unauthorized, error)

	PlatformTokenGrantV3(params *PlatformTokenGrantV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformTokenGrantV3OK, *PlatformTokenGrantV3BadRequest, *PlatformTokenGrantV3Unauthorized, error)

	RetrieveUserThirdPartyPlatformTokenV3(params *RetrieveUserThirdPartyPlatformTokenV3Params, authInfo runtime.ClientAuthInfoWriter) (*RetrieveUserThirdPartyPlatformTokenV3OK, *RetrieveUserThirdPartyPlatformTokenV3Unauthorized, *RetrieveUserThirdPartyPlatformTokenV3Forbidden, *RetrieveUserThirdPartyPlatformTokenV3NotFound, error)

	RevokeUserV3(params *RevokeUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*RevokeUserV3OK, *RevokeUserV3BadRequest, *RevokeUserV3Unauthorized, *RevokeUserV3Forbidden, error)

	TokenGrantV3(params *TokenGrantV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenGrantV3OK, *TokenGrantV3BadRequest, *TokenGrantV3Unauthorized, *TokenGrantV3Forbidden, error)

	TokenIntrospectionV3(params *TokenIntrospectionV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenIntrospectionV3OK, *TokenIntrospectionV3BadRequest, *TokenIntrospectionV3Unauthorized, error)

	TokenRevocationV3(params *TokenRevocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenRevocationV3OK, *TokenRevocationV3BadRequest, *TokenRevocationV3Unauthorized, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AuthCodeRequestV3 generates url to request auth code from third party platform

  'Generate url to request auth code from third party platform <br>
               <h2>Supported platforms:</h2><ul>
               <li><strong>steamopenid</strong></li>This endpoint redirects to steam login page, then redirect back to platform
               authenticate endpoint after successfully authenticating user steam.
               <li><strong>xblweb</strong></li>This endpoint redirects to xbox login page, then redirect back to platform
               authenticate endpoint after successfully authenticating xbox user.
               <li><strong>ps4web</strong></li>This endpoint redirects to psn login page, then redirect back to platform
               authenticate endpoint after successfully authenticating psn user.
               <li><strong>epicgames</strong></li>This endpoint redirects to Epicgames OAuth login page. then redirect to platform
               authenticate endpoint after successfully authenticating an Epicgames credential
               <li><strong>twitch</strong></li>This endpoint redirects to twitch login page, then redirect back to platform
               authenticate endpoint after successfully authenticating twitch user.
               </ul> action code : 10702'
*/
func (a *Client) AuthCodeRequestV3(params *AuthCodeRequestV3Params, authInfo runtime.ClientAuthInfoWriter) (*AuthCodeRequestV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthCodeRequestV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AuthCodeRequestV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/oauth/platforms/{platformId}/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthCodeRequestV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AuthCodeRequestV3Found:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AuthorizeV3 os auth2 authorize API

  Initializes OAuth2.0 authorization code flow<br/>
<p>The endpoint stores authorization request and redirects to login page with the authorization request id.
	The user can then do the authentication on the login page.
	The user will be redirected back to the requesting client with authorization code if successfully authenticated.
</p>
<p>Only authorization code flow supported by this endpoint, implicit flow is not supported.</p>
<ul>
	<li><p><strong>Authorize success</strong>:
		redirects to login page with the following information: ?request_id={authorization_request_id}</p>
	</li>
	<li><p><strong>Authorize failure</strong>:
		redirects to the given redirect uri with the following information:
		?error={error_code}&error_description={error description}</p>
	</li>
</ul>
	<p>Following are the error code based on the specification:</p>
<ul>
	<li><p>invalid_request: The request is missing a required parameter,
		includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.</p>
	</li>
	<li><p>server_error:
		The authorization server encountered an unexpected condition that prevented it from fulfilling the request.</p>
	</li>
	<li><p>unauthorized_client: The client is not authorized to request a token using this method.</p></li>
	<li><p>access_denied: The resource owner or authorization server denied the request.</p></li>
	<li><p>invalid_scope: The requested scope is invalid, unknown, or malformed.</p></li>
	<li><p>unsupported_response_type: The authorization server does not support obtaining a token using this method.</p></li>
	<li><p>temporarily_unavailable: The authorization server is currently unable to handle the request
		due to a temporary overloading or maintenance of the server.</p>
	</li>
</ul>
	<p>Please refer to the RFC for more information about authorization code flow: https://tools.ietf.org/html/rfc6749#section-4.1</p><br>
	action code: 10701


*/
func (a *Client) AuthorizeV3(params *AuthorizeV3Params, authInfo runtime.ClientAuthInfoWriter) (*AuthorizeV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthorizeV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AuthorizeV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/oauth/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthorizeV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AuthorizeV3Found:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetJWKSV3 JSONs web key set for verifying j w t

  <p>This endpoint serves public keys for verifying JWT access tokens generated by this service.</p>
			<p>When a client application wants to verify a JWT token, it needs to get the 'kid' value found in the JWT token header and use it
			to look up the corresponding public key from a set returned by this endpoint. The client application can then use that public key to verify the JWT.</p>
			<p>A client application might cache the keys so it doesn't need to do request every time it needs to verify a JWT token. If a client application
			caches the keys and a key with the same 'kid' cannot be found in the cache, it should then try to refresh the keys by making a request to this
			endpoint again.</p>
			<p>Please refer to the RFC for more information about JWK (JSON Web Key): https://tools.ietf.org/html/rfc7517</p>
			<br>action code : 10709
*/
func (a *Client) GetJWKSV3(params *GetJWKSV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetJWKSV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJWKSV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetJWKSV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/oauth/jwks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJWKSV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetJWKSV3OK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetRevocationListV3 os auth2 revocation list API

  <p>This endpoint will return a list of revoked users and revoked tokens. List of revoked tokens in bloom filter format.</p>
					 <p>This endpoint requires authorized requests header with valid access token.</p>
					 <p>The bloom filter uses MurmurHash3 algorithm for hashing the values</p>
					 <p>action code : 10708</p>
*/
func (a *Client) GetRevocationListV3(params *GetRevocationListV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetRevocationListV3OK, *GetRevocationListV3Unauthorized, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRevocationListV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRevocationListV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/oauth/revocationlist",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRevocationListV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetRevocationListV3OK:
		return v, nil, nil
	case *GetRevocationListV3Unauthorized:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PlatformTokenGrantV3 os auth2 access token generation specific to platform

  <p>Platform token grant specifically used for performing token grant using platform, e.g. Steam, Justice, etc. The endpoint automatically create an account if the account associated with the platform is not exists yet.
			This endpoint requires all requests to have Authorization header set with Basic access authentication
			constructed from client id and client secret. For publisher-game namespace schema : Specify only either platform_token or device_id. Device token grant
			should be requested along with device_id parameter against game namespace. Another 3rd party platform token grant should be requested
			along with platform_token parameter against publisher namespace.</p>
			<h2>Supported platforms:</h2>
			<ul>
				<li><strong>steam</strong>: The platform_token’s value is the authentication code returned by Steam.</li>
				<li><strong>steamopenid</strong>: Steam's user authentication method using OpenID 2.0. The platform_token's value is URL generated by Steam on web authentication</li>
				<li><strong>facebook</strong>: The platform_token’s value is the authorization code returned by Facebook OAuth</li>
				<li><strong>google</strong>: The platform_token’s value is the authorization code returned by Google OAuth</li>
				<li><strong>oculus</strong>: The platform_token’s value is a string composed of Oculus's user ID and the nonce separated by a colon (:).</li>
				<li><strong>twitch</strong>: The platform_token’s value is the authorization code returned by Twitch OAuth.</li>
				<li><strong>discord</strong>: The platform_token’s value is the authorization code returned by Discord OAuth</li>
				<li><strong>android</strong>: The device_id is the Android’s device ID</li>
				<li><strong>ios</strong>: The device_id is the iOS’s device ID.</li>
				<li><strong>device</strong>: Every device that does’nt run Android and iOS is categorized as a device. The device_id is the device’s ID.</li>
				<li><strong>justice</strong>: The platform_token’s value is the designated user’s access token.</li>
				<li><strong>epicgames</strong>: The platform_token’s value is an access-token obtained from Epicgames EOS Account Service.</li>
                <li><strong>stadia</strong>: The platform_token's value is a JWT Token, which can be obtained after calling the Stadia SDK's function.</li>
				<li><strong>ps4</strong>: The platform_token’s value is the authorization code returned by Sony OAuth.</li>
				<li><strong>ps5</strong>: The platform_token’s value is the authorization code returned by Sony OAuth.</li>
				<li><strong>nintendo</strong>: The platform_token’s value is the authorization code(id_token) returned by Nintendo OAuth.</li>
			</ul>
			<h2>Account Group</h2>
			<p>Several platforms are grouped under account groups. The accounts on these platforms have the same platform user id.
			Login using one of these platform will returns the same IAM user. </p>
			<p>Following is the current registered account grouping: </p>
			<ul>
			<li> (psn) ps4web </li>
			<li> (psn) ps4 </li>
			<li> (psn) ps5 </li>
			</ul>
			<h2>Access Token Content</h2>
			<p>Following is the access token’s content:</p>
			<ul>
			<li>
				<p><strong>namespace</strong>. It is the namespace the token was generated from.</p>
			</li>
			<li>
				<p><strong>display_name</strong>. The display name of the sub. It is empty if the token is generated from the client credential</p>
			</li>
			<li>
				<p><strong>roles</strong>. The sub’s roles. It is empty if the token is generated from the client credential</p>
			</li>
			<li>
				<p><strong>namespace_roles</strong>. The sub’s roles scoped to namespace. Improvement from roles, which make the role scoped to specific namespace instead of global to publisher namespace</p>
			</li>
			<li>
				<p><strong>permissions</strong>. The sub or aud’ permissions</p>
			</li>
			<li>
				<p><strong>bans</strong>. The sub’s list of bans. It is used by the IAM client for validating the token.</p>
			</li>
			<li>
				<p><strong>jflgs</strong>. It stands for Justice Flags. It is a special flag used for storing additional status information regarding the sub. It is implemented as a bit mask. Following explains what each bit represents:</p>
			<ul>
				<li><p>1: Email Address Verified</p></li>
				<li><p>2: Phone Number Verified</p></li>
				<li><p>4: Anonymous</p></li>
				<li><p>8: Suspicious Login</p></li>
			</ul>
			</li>
			<li>
				<p><strong>aud</strong>. The aud is the client ID.</p>
			</li>
			<li>
				<p><strong>iat</strong>. The time the token issues at. It is in Epoch time format</p>
			</li>
			<li>
				<p><strong>exp</strong>. The time the token expires. It is in Epoch time format</p>
			</li>
			<li>
				<p><strong>sub</strong>. The UserID. The sub is omitted if the token is generated from client credential</p>
			</li>
			<h2>Bans</h2>
			<p>The JWT contains user's active bans with its expiry date. List of ban types can be obtained from /bans.</p>
			<br>action code : 10704
*/
func (a *Client) PlatformTokenGrantV3(params *PlatformTokenGrantV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformTokenGrantV3OK, *PlatformTokenGrantV3BadRequest, *PlatformTokenGrantV3Unauthorized, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlatformTokenGrantV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PlatformTokenGrantV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/oauth/platforms/{platformId}/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PlatformTokenGrantV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PlatformTokenGrantV3OK:
		return v, nil, nil, nil
	case *PlatformTokenGrantV3BadRequest:
		return nil, v, nil, nil
	case *PlatformTokenGrantV3Unauthorized:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveUserThirdPartyPlatformTokenV3 retrieves user third party platform token

  Retrieve User Third Party Platform Token<br/>
<p>
This endpoint used for retrieving third party platform token for user that login using third party.
Passing platform group name or it's member will return same access token that can be used across the platform members.
</p>
<p>The third party platform and platform group covered for this is:</p>
<ul>
	<li>(psn) ps4web</li>
	<li>(psn) ps4</li>
	<li>(psn) ps5</li>
	<li>epicgames</li>
	<li>twitch</li>
</ul>
*/
func (a *Client) RetrieveUserThirdPartyPlatformTokenV3(params *RetrieveUserThirdPartyPlatformTokenV3Params, authInfo runtime.ClientAuthInfoWriter) (*RetrieveUserThirdPartyPlatformTokenV3OK, *RetrieveUserThirdPartyPlatformTokenV3Unauthorized, *RetrieveUserThirdPartyPlatformTokenV3Forbidden, *RetrieveUserThirdPartyPlatformTokenV3NotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveUserThirdPartyPlatformTokenV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveUserThirdPartyPlatformTokenV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveUserThirdPartyPlatformTokenV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveUserThirdPartyPlatformTokenV3OK:
		return v, nil, nil, nil, nil
	case *RetrieveUserThirdPartyPlatformTokenV3Unauthorized:
		return nil, v, nil, nil, nil
	case *RetrieveUserThirdPartyPlatformTokenV3Forbidden:
		return nil, nil, v, nil, nil
	case *RetrieveUserThirdPartyPlatformTokenV3NotFound:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RevokeUserV3 revokes user s tokens

  <p>This endpoint revokes all access tokens and refresh tokens a user has prior the revocation time.</p>
			<p>This endpoint requires authorized requests header with valid access token.</p>
			<p>Required permission 'ADMIN:NAMESPACE:{namespace}:USER:{userId} [UPDATE]'</p>
			<p>It is a convenient feature for the developer (or admin) who wanted to revokes all user's access tokens and refresh tokens generated before some period of time.</p>
			<p>action code : 10707</p>
*/
func (a *Client) RevokeUserV3(params *RevokeUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*RevokeUserV3OK, *RevokeUserV3BadRequest, *RevokeUserV3Unauthorized, *RevokeUserV3Forbidden, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeUserV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RevokeUserV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/oauth/admin/namespaces/{namespace}/users/{userId}/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RevokeUserV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RevokeUserV3OK:
		return v, nil, nil, nil, nil
	case *RevokeUserV3BadRequest:
		return nil, v, nil, nil, nil
	case *RevokeUserV3Unauthorized:
		return nil, nil, v, nil, nil
	case *RevokeUserV3Forbidden:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  TokenGrantV3 os auth2 access token generation endpoint

  <p>This endpoint supports grant type:</p><ol>
	 		<li>Grant Type == <code>authorization_code</code>:<br />
				&nbsp;&nbsp;&nbsp; It generates the user token by given the authorization
				code which generated in "/v3/oauth/auth" API response. It should also pass
				in the redirect_uri, which should be the same as generating the
				authorization code request.
			</li>
			<li>Grant Type == <code>refresh_token</code>:<br />
	 			&nbsp;&nbsp;&nbsp; Used to get a new access token for a valid refresh token.
			</li>
			<li>Grant Type == <code>client_credentials</code>:<br />
	 			&nbsp;&nbsp;&nbsp; It generates a token by checking the client credentials provided through Authorization header.
			</li></ol>
			<h2>Access Token Content</h2>
			<p>Following is the access token’s content:</p>
			<ul>
			<li>
				<p><strong>namespace</strong>. It is the namespace the token was generated from.</p>
			</li>
			<li>
				<p><strong>display_name</strong>. The display name of the sub. It is empty if the token is generated from the client credential</p>
			</li>
			<li>
				<p><strong>roles</strong>. The sub’s roles. It is empty if the token is generated from the client credential</p>
			</li>
			<li>
				<p><strong>namespace_roles</strong>. The sub’s roles scoped to namespace. Improvement from roles, which make the role scoped to specific namespace instead of global to publisher namespace</p>
			</li>
			<li>
				<p><strong>permissions</strong>. The sub or aud’ permissions</p>
			</li>
			<li>
				<p><strong>bans</strong>. The sub’s list of bans. It is used by the IAM client for validating the token.</p>
			</li>
			<li>
				<p><strong>jflgs</strong>. It stands for Justice Flags. It is a special flag used for storing additional status information regarding the sub. It is implemented as a bit mask. Following explains what each bit represents:</p>
			<ul>
				<li><p>1: Email Address Verified</p></li>
				<li><p>2: Phone Number Verified</p></li>
				<li><p>4: Anonymous</p></li>
				<li><p>8: Suspicious Login</p></li>
			</ul>
			</li>
			<li>
				<p><strong>aud</strong>. The aud is the targeted resource server.</p>
			</li>
			<li>
				<p><strong>iat</strong>. The time the token issues at. It is in Epoch time format</p>
			</li>
			<li>
				<p><strong>exp</strong>. The time the token expires. It is in Epoch time format</p>
			</li>
			<li>
				<p><strong>client_id</strong>. The UserID. The sub is omitted if the token is generated from client credential</p>
			</li>
			<li>
				<p><strong>scope</strong>. The scope of the access request, expressed as a list of space-delimited, case-sensitive strings</p>
			</li>
			</ul>
			<h2>Bans</h2>
			<p>The JWT contains user's active bans with its expiry date. List of ban types can be obtained from /bans.</p>
			<h2>Track Login History</h2>
			<p>This endpoint will track login history to detect suspicious login activity, please provide "device_id" (alphanumeric) in request header parameter otherwise we will set to "unknown".</p>
			<p>Align with General Data Protection Regulation in Europe, user login history will be kept within 28 days by default"</p>
			<p>action code: 10703
*/
func (a *Client) TokenGrantV3(params *TokenGrantV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenGrantV3OK, *TokenGrantV3BadRequest, *TokenGrantV3Unauthorized, *TokenGrantV3Forbidden, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenGrantV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TokenGrantV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/oauth/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TokenGrantV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *TokenGrantV3OK:
		return v, nil, nil, nil, nil
	case *TokenGrantV3BadRequest:
		return nil, v, nil, nil, nil
	case *TokenGrantV3Unauthorized:
		return nil, nil, v, nil, nil
	case *TokenGrantV3Forbidden:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  TokenIntrospectionV3 os auth2 token introspection API

  <p>This endpoint returns information about an access token intended to be used by resource servers or other internal servers.</p>
					 <p>This endpoint requires authorized requests header with valid basic or bearer token.</p>
					 <p>action code : 10705</p>
*/
func (a *Client) TokenIntrospectionV3(params *TokenIntrospectionV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenIntrospectionV3OK, *TokenIntrospectionV3BadRequest, *TokenIntrospectionV3Unauthorized, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenIntrospectionV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TokenIntrospectionV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/oauth/introspect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TokenIntrospectionV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *TokenIntrospectionV3OK:
		return v, nil, nil, nil
	case *TokenIntrospectionV3BadRequest:
		return nil, v, nil, nil
	case *TokenIntrospectionV3Unauthorized:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  TokenRevocationV3 os auth2 token revocation API

  <p>This endpoint revokes a token.</p>
					 <p>This endpoint requires authorized requests header with valid access token.</p><br>action code: 10706
*/
func (a *Client) TokenRevocationV3(params *TokenRevocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*TokenRevocationV3OK, *TokenRevocationV3BadRequest, *TokenRevocationV3Unauthorized, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenRevocationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TokenRevocationV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/oauth/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TokenRevocationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *TokenRevocationV3OK:
		return v, nil, nil, nil
	case *TokenRevocationV3BadRequest:
		return nil, v, nil, nil
	case *TokenRevocationV3Unauthorized:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
