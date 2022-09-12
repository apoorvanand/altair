package usecase

import "github.com/kodefluence/altair/plugin/oauth/entity"

type Formatter interface {
	AccessTokenFromAuthorizationRequestInsertable(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessTokenInsertable
	AccessTokenFromOauthAccessGrantInsertable(oauthAccessGrant entity.OauthAccessGrant, application entity.OauthApplication) entity.OauthAccessTokenInsertable
	AccessGrantFromAuthorizationRequestInsertable(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessGrantInsertable
	OauthApplicationInsertable(r entity.OauthApplicationJSON) entity.OauthApplicationInsertable
	AccessTokenFromOauthRefreshTokenInsertable(application entity.OauthApplication, accessToken entity.OauthAccessToken) entity.OauthAccessTokenInsertable
	RefreshTokenInsertable(application entity.OauthApplication, accessToken entity.OauthAccessToken) entity.OauthRefreshTokenInsertable
	AccessGrant(e entity.OauthAccessGrant) entity.OauthAccessGrantJSON
	AccessToken(e entity.OauthAccessToken, redirectURI string, refreshTokenJSON *entity.OauthRefreshTokenJSON) entity.OauthAccessTokenJSON
	RefreshToken(e entity.OauthRefreshToken) entity.OauthRefreshTokenJSON
}
