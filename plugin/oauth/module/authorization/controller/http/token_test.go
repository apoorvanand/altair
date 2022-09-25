package http_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kodefluence/altair/module/apierror"
	"github.com/kodefluence/altair/plugin/oauth/entity"
	authorizationHttp "github.com/kodefluence/altair/plugin/oauth/module/authorization/controller/http"
	"github.com/kodefluence/altair/plugin/oauth/module/authorization/controller/http/mock"
	"github.com/kodefluence/altair/plugin/oauth/module/formatter"
	"github.com/kodefluence/altair/testhelper"
	"github.com/kodefluence/altair/util"
	"github.com/kodefluence/aurelia"
	"github.com/kodefluence/monorepo/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Method", func(t *testing.T) {
		authorizationService := mock.NewMockAuthorization(mockCtrl)
		assert.Equal(t, "POST", authorizationHttp.NewToken(authorizationService, apierror.Provide()).Method())
	})

	t.Run("Path", func(t *testing.T) {
		authorizationService := mock.NewMockAuthorization(mockCtrl)
		assert.Equal(t, "/oauth/authorizations/token", authorizationHttp.NewToken(authorizationService, apierror.Provide()).Path())
	})

	t.Run("Control", func(t *testing.T) {
		t.Run("Given request with json body", func(t *testing.T) {
			t.Run("Return oauth application data with status 202", func(t *testing.T) {
				apiEngine := gin.Default()

				accessTokenRequest := entity.AccessTokenRequestJSON{
					GrantType:    util.StringToPointer("authorization_code"),
					ClientUID:    util.StringToPointer(aurelia.Hash("x", "y")),
					ClientSecret: util.StringToPointer(aurelia.Hash("z", "a")),
					RedirectURI:  util.StringToPointer("http://github.com"),
					Code:         util.StringToPointer("authorization_code"),
				}
				encodedBytes, err := json.Marshal(accessTokenRequest)
				assert.Nil(t, err)

				oauthAccessToken := entity.OauthAccessToken{
					ID:                 1,
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              aurelia.Hash("x", "y"),
					Scopes: sql.NullString{
						String: "user",
						Valid:  true,
					},
					ExpiresIn: time.Now().Add(time.Hour * 4),
					CreatedAt: time.Now().Truncate(time.Second),
				}
				oauthAccessTokenJSON := formatter.Provide(time.Hour, time.Hour, time.Hour).AccessToken(oauthAccessToken, *accessTokenRequest.RedirectURI, nil)
				expectedReponseByte, _ := json.Marshal(jsonapi.BuildResponse(jsonapi.WithData(oauthAccessTokenJSON)))

				authorizationService := mock.NewMockAuthorization(mockCtrl)
				authorizationService.EXPECT().GrantToken(gomock.Any(), accessTokenRequest).Return(oauthAccessTokenJSON, nil)

				ctrl := authorizationHttp.NewToken(authorizationService, apierror.Provide())
				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))
				responseByte, err := io.ReadAll(w.Body)
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, string(expectedReponseByte), string(responseByte))
			})

			t.Run("Unexpected error in authorization services", func(t *testing.T) {
				t.Run("Return entity error status", func(t *testing.T) {
					apiEngine := gin.Default()

					accessTokenRequest := entity.AccessTokenRequestJSON{
						GrantType:    util.StringToPointer("authorization_code"),
						ClientUID:    util.StringToPointer(aurelia.Hash("x", "y")),
						ClientSecret: util.StringToPointer(aurelia.Hash("z", "a")),
						RedirectURI:  util.StringToPointer("http://github.com"),
						Code:         util.StringToPointer("authorization_code"),
					}
					encodedBytes, err := json.Marshal(accessTokenRequest)
					assert.Nil(t, err)

					authorizationService := mock.NewMockAuthorization(mockCtrl)
					authorizationService.EXPECT().GrantToken(gomock.Any(), accessTokenRequest).Return(entity.OauthAccessTokenJSON{}, testhelper.ErrInternalServer())

					ctrl := authorizationHttp.NewToken(authorizationService, apierror.Provide())
					apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

					w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))
					responseByte, err := io.ReadAll(w.Body)
					assert.Nil(t, err)
					assert.Equal(t, http.StatusInternalServerError, w.Code)
					assert.Equal(t, "{\"errors\":[{\"title\":\"Internal server error\",\"detail\":\"Something is not right, help us fix this problem. Contribute to https://github.com/kodefluence/altair. Tracing code: '\\u003cnil\\u003e'\",\"code\":\"ERR0500\",\"status\":500}]}", string(responseByte))
				})
			})
		})

		t.Run("Given invalid request body", func(t *testing.T) {
			t.Run("Return bad request", func(t *testing.T) {
				apiEngine := gin.Default()

				authorizationService := mock.NewMockAuthorization(mockCtrl)

				ctrl := authorizationHttp.NewToken(authorizationService, apierror.Provide())
				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), testhelper.MockErrorIoReader{})
				responseByte, err := io.ReadAll(w.Body)
				assert.Nil(t, err)
				assert.Equal(t, http.StatusBadRequest, w.Code)
				assert.Equal(t, "{\"errors\":[{\"title\":\"Bad request error\",\"detail\":\"You've send malformed request in your `request body`\",\"code\":\"ERR0400\",\"status\":400}]}", string(responseByte))
			})
		})

		t.Run("Given request body but not json", func(t *testing.T) {
			t.Run("Return bad request", func(t *testing.T) {
				apiEngine := gin.Default()

				authorizationService := mock.NewMockAuthorization(mockCtrl)

				ctrl := authorizationHttp.NewToken(authorizationService, apierror.Provide())
				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader([]byte(`this is gonna be error`)))
				responseByte, err := io.ReadAll(w.Body)
				assert.Nil(t, err)
				assert.Equal(t, http.StatusBadRequest, w.Code)
				assert.Equal(t, "{\"errors\":[{\"title\":\"Bad request error\",\"detail\":\"You've send malformed request in your `request body`\",\"code\":\"ERR0400\",\"status\":400}]}", string(responseByte))
			})
		})
	})
}
