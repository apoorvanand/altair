package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/codefluence-x/altair/provider/plugin/oauth/entity"
	"github.com/codefluence-x/altair/provider/plugin/oauth/eobject"
	"github.com/codefluence-x/altair/provider/plugin/oauth/interfaces"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Authorization struct handle all of things related to oauth2 authorization
type Authorization struct {
	oauthApplicationModel interfaces.OauthApplicationModel
	oauthAccessTokenModel interfaces.OauthAccessTokenModel
	oauthAccessGrantModel interfaces.OauthAccessGrantModel
	oauthValidator        interfaces.OauthValidator

	modelFormatter interfaces.ModelFormater
	oauthFormatter interfaces.OauthFormatter

	config entity.PluginConfig
}

// NewAuthorization create new service to handler authorize related flow
func NewAuthorization(
	oauthApplicationModel interfaces.OauthApplicationModel,
	oauthAccessTokenModel interfaces.OauthAccessTokenModel,
	oauthAccessGrantModel interfaces.OauthAccessGrantModel,
	modelFormatter interfaces.ModelFormater,
	oauthValidator interfaces.OauthValidator,
	oauthFormatter interfaces.OauthFormatter,
	config entity.PluginConfig,
) *Authorization {
	return &Authorization{
		oauthApplicationModel: oauthApplicationModel,
		oauthAccessTokenModel: oauthAccessTokenModel,
		oauthAccessGrantModel: oauthAccessGrantModel,
		modelFormatter:        modelFormatter,
		oauthValidator:        oauthValidator,
		oauthFormatter:        oauthFormatter,
		config:                config,
	}
}

// Grantor provide granting logic for authorization request
func (a *Authorization) Grantor(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (interface{}, *entity.Error) {
	if authorizationReq.ResponseType == nil {
		return nil, &entity.Error{
			HttpStatus: http.StatusUnprocessableEntity,
			Errors:     eobject.Wrap(eobject.ValidationError("response_type cannot be empty")),
		}
	}

	switch *authorizationReq.ResponseType {
	case "token":
		return a.GrantToken(ctx, authorizationReq)
	case "code":
		return a.Grant(ctx, authorizationReq)
	}

	err := &entity.Error{
		HttpStatus: http.StatusUnprocessableEntity,
		Errors:     eobject.Wrap(eobject.ValidationError("response_type is invalid. Should be either `token` or `code`.")),
	}

	log.Error().
		Err(err).
		Stack().
		Interface("request_id", ctx.Value("request_id")).
		Interface("request", authorizationReq).
		Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grantor")).
		Msg("invalid response type sent by client")
	return nil, err
}

// Grant authorization an access code
func (a *Authorization) Grant(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessGrantJSON, *entity.Error) {
	var oauthAccessGrantJSON entity.OauthAccessGrantJSON

	oauthApplication, entityErr := a.findAndValidateApplication(ctx, authorizationReq.ClientUID, authorizationReq.ClientSecret)
	if entityErr != nil {
		return oauthAccessGrantJSON, entityErr
	}

	if err := a.oauthValidator.ValidateAuthorizationGrant(ctx, authorizationReq, oauthApplication); err != nil {
		return oauthAccessGrantJSON, err
	}

	id, err := a.oauthAccessGrantModel.Create(ctx, a.modelFormatter.AccessGrantFromAuthorizationRequest(authorizationReq, oauthApplication))
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Interface("request", authorizationReq).
			Int("application_id", oauthApplication.ID).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant")).
			Msg("Error creating access grant")

		return entity.OauthAccessGrantJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	oauthAccessGrant, err := a.oauthAccessGrantModel.One(ctx, id)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Interface("request", authorizationReq).
			Int("application_id", oauthApplication.ID).
			Int("last_inserted_id", id).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant")).
			Msg("Error selecting one access grant after creating the data")
		return entity.OauthAccessGrantJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	return a.oauthFormatter.AccessGrant(oauthAccessGrant), nil
}

// GrantToken will grant an access token
func (a *Authorization) GrantToken(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	oauthApplication, entityErr := a.findAndValidateApplication(ctx, authorizationReq.ClientUID, authorizationReq.ClientSecret)
	if entityErr != nil {
		return entity.OauthAccessTokenJSON{}, entityErr
	}

	if err := a.oauthValidator.ValidateAuthorizationGrant(ctx, authorizationReq, oauthApplication); err != nil {
		return entity.OauthAccessTokenJSON{}, err
	}

	id, err := a.oauthAccessTokenModel.Create(ctx, a.modelFormatter.AccessTokenFromAuthorizationRequest(authorizationReq, oauthApplication))
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Interface("request", authorizationReq).
			Int("application_id", oauthApplication.ID).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant_token")).
			Msg("Error creating access token after creating the data")
		return entity.OauthAccessTokenJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	oauthAccessToken, err := a.oauthAccessTokenModel.One(ctx, id)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Interface("request", authorizationReq).
			Int("application_id", oauthApplication.ID).
			Int("last_inserted_id", id).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant_token")).
			Msg("Error selecting one access token")

		return entity.OauthAccessTokenJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	return a.oauthFormatter.AccessToken(oauthAccessToken, *authorizationReq.RedirectURI), nil
}

func (a *Authorization) findAndValidateApplication(ctx context.Context, clientUID, clientSecret *string) (entity.OauthApplication, *entity.Error) {
	if clientUID == nil {
		return entity.OauthApplication{}, &entity.Error{
			HttpStatus: http.StatusUnprocessableEntity,
			Errors:     eobject.Wrap(eobject.ValidationError("client_uid cannot be empty")),
		}
	}

	if clientSecret == nil {
		return entity.OauthApplication{}, &entity.Error{
			HttpStatus: http.StatusUnprocessableEntity,
			Errors:     eobject.Wrap(eobject.ValidationError("client_secret cannot be empty")),
		}
	}

	oauthApplication, err := a.oauthApplicationModel.OneByUIDandSecret(ctx, *clientUID, *clientSecret)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Str("client_uid", *clientUID).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("find_secret")).
			Msg("application cannot be found because there was an error")
		if err == sql.ErrNoRows {
			return entity.OauthApplication{}, &entity.Error{
				HttpStatus: http.StatusNotFound,
				Errors:     eobject.Wrap(eobject.NotFoundError(ctx, "client_uid & client_secret")),
			}
		}

		return entity.OauthApplication{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	return oauthApplication, nil
}

// Token will grant a token from authorization code
func (a *Authorization) Token(ctx context.Context, accessTokenReq entity.AccessTokenRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	oauthApplication, entityErr := a.findAndValidateApplication(ctx, accessTokenReq.ClientUID, accessTokenReq.ClientSecret)
	if entityErr != nil {
		return entity.OauthAccessTokenJSON{}, entityErr
	}

	if entityErr := a.oauthValidator.ValidateTokenGrant(ctx, accessTokenReq); entityErr != nil {
		return entity.OauthAccessTokenJSON{}, entityErr
	}

	oauthAccessGrant, err := a.oauthAccessGrantModel.OneByCode(ctx, *accessTokenReq.Code)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("one_by_code")).
			Msg("authorization code cannot be found because there was an error")
		if err == sql.ErrNoRows {
			return entity.OauthAccessTokenJSON{}, &entity.Error{
				HttpStatus: http.StatusNotFound,
				Errors:     eobject.Wrap(eobject.NotFoundError(ctx, "authorization_code")),
			}
		}

		return entity.OauthAccessTokenJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	if err := a.oauthValidator.ValidateTokenAuthorizationCode(ctx, accessTokenReq, oauthAccessGrant); err != nil {
		return entity.OauthAccessTokenJSON{}, err
	}

	id, err := a.oauthAccessTokenModel.Create(ctx, a.modelFormatter.AccessTokenFromOauthAccessGrant(oauthAccessGrant, oauthApplication))
	if err != nil {

		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant_token")).
			Msg("Error creating access token after creating the data")

		return entity.OauthAccessTokenJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	oauthAccessToken, err := a.oauthAccessTokenModel.One(ctx, id)
	if err != nil {

		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant_token")).
			Msg("Error selecting one access token")

		return entity.OauthAccessTokenJSON{}, &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	err = a.oauthAccessGrantModel.Revoke(ctx, *accessTokenReq.Code)
	if err != nil {
		// TODO: Error is intended to be suppressed until database transaction is implemented. After database transaction is implemented, then it will be rollbacked if there is error in revoke oauth access grants process
		log.Error().
			Err(err).
			Stack().
			Interface("request_id", ctx.Value("request_id")).
			Array("tags", zerolog.Arr().Str("service").Str("authorization").Str("grant_token")).
			Msg("Error revoke oauth access grant")
	}

	return a.oauthFormatter.AccessToken(oauthAccessToken, oauthAccessGrant.RedirectURI.String), nil
}

// RevokeToken revoke given access token request
func (a *Authorization) RevokeToken(ctx context.Context, revokeAccessTokenReq entity.RevokeAccessTokenRequestJSON) *entity.Error {

	if revokeAccessTokenReq.Token == nil {
		return &entity.Error{
			HttpStatus: http.StatusUnprocessableEntity,
			Errors:     eobject.Wrap(eobject.ValidationError("token is empty")),
		}
	}

	err := a.oauthAccessTokenModel.Revoke(ctx, *revokeAccessTokenReq.Token)
	if err == sql.ErrNoRows {
		return &entity.Error{
			HttpStatus: http.StatusNotFound,
			Errors:     eobject.Wrap(eobject.NotFoundError(ctx, "token")),
		}
	} else if err != nil {
		return &entity.Error{
			HttpStatus: http.StatusInternalServerError,
			Errors:     eobject.Wrap(eobject.InternalServerError(ctx)),
		}
	}

	return nil
}
