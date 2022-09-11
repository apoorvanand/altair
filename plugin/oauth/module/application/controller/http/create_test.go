package http_test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/kodefluence/altair/provider/plugin/oauth/controller"
// 	"github.com/kodefluence/altair/provider/plugin/oauth/entity"
// 	"github.com/kodefluence/altair/provider/plugin/oauth/eobject"
// 	"github.com/kodefluence/altair/provider/plugin/oauth/mock"
// 	"github.com/kodefluence/altair/testhelper"
// 	"github.com/kodefluence/altair/util"
// 	"github.com/stretchr/testify/assert"
// )

// type ErrorResponse struct {
// 	Errors []entity.ErrorObject `json:"errors"`
// }

// func TestCreate(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	t.Run("Method", func(t *testing.T) {
// 		applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 		assert.Equal(t, "POST", controller.NewApplication().Create(applicationManager).Method())
// 	})

// 	t.Run("Path", func(t *testing.T) {
// 		applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 		assert.Equal(t, "/oauth/applications", controller.NewApplication().Create(applicationManager).Path())
// 	})

// 	t.Run("Control", func(t *testing.T) {
// 		t.Run("Given request with json body", func(t *testing.T) {
// 			t.Run("Return oauth application data with status 202", func(t *testing.T) {
// 				apiEngine := gin.Default()

// 				oauthApplicationJSON := entity.OauthApplicationJSON{
// 					OwnerID:      util.IntToPointer(1),
// 					Description:  util.StringToPointer("Application 1"),
// 					Scopes:       util.StringToPointer("public user"),
// 					ClientUID:    util.StringToPointer("clientuid01"),
// 					ClientSecret: util.StringToPointer("clientsecret01"),
// 				}
// 				encodedBytes, err := json.Marshal(oauthApplicationJSON)
// 				assert.Nil(t, err)

// 				applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 				applicationManager.EXPECT().Create(gomock.Any(), oauthApplicationJSON).Return(oauthApplicationJSON, nil)

// 				ctrl := controller.NewApplication().Create(applicationManager)
// 				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

// 				var response responseOne
// 				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))

// 				err = json.Unmarshal(w.Body.Bytes(), &response)
// 				assert.Nil(t, err)

// 				assert.Equal(t, http.StatusCreated, w.Code)
// 				assert.Equal(t, oauthApplicationJSON, response.Data)
// 			})

// 			t.Run("Unexpected error in application manager", func(t *testing.T) {
// 				apiEngine := gin.Default()

// 				oauthApplicationJSON := entity.OauthApplicationJSON{
// 					OwnerID:      util.IntToPointer(1),
// 					Description:  util.StringToPointer("Application 1"),
// 					Scopes:       util.StringToPointer("public user"),
// 					ClientUID:    util.StringToPointer("clientuid01"),
// 					ClientSecret: util.StringToPointer("clientsecret01"),
// 				}
// 				encodedBytes, err := json.Marshal(oauthApplicationJSON)
// 				assert.Nil(t, err)

// 				expectedError := &entity.Error{
// 					HttpStatus: http.StatusInternalServerError,
// 					Errors:     eobject.Wrap(eobject.InternalServerError(context.Background())),
// 				}

// 				applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 				applicationManager.EXPECT().Create(gomock.Any(), oauthApplicationJSON).Return(entity.OauthApplicationJSON{}, expectedError)

// 				ctrl := controller.NewApplication().Create(applicationManager)
// 				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

// 				var response ErrorResponse
// 				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader(encodedBytes))

// 				err = json.Unmarshal(w.Body.Bytes(), &response)
// 				assert.Nil(t, err)

// 				assert.Equal(t, expectedError.HttpStatus, w.Code)
// 				assert.Equal(t, expectedError.Errors, response.Errors)
// 			})
// 		})

// 		t.Run("Given invalid request body", func(t *testing.T) {
// 			t.Run("Return bad request", func(t *testing.T) {
// 				apiEngine := gin.Default()

// 				applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 				applicationManager.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)

// 				ctrl := controller.NewApplication().Create(applicationManager)
// 				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

// 				expectedError := &entity.Error{
// 					HttpStatus: http.StatusBadRequest,
// 					Errors:     eobject.Wrap(eobject.BadRequestError("request body")),
// 				}

// 				var response ErrorResponse
// 				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), testhelper.MockErrorIoReader{})

// 				err := json.Unmarshal(w.Body.Bytes(), &response)
// 				assert.Nil(t, err)

// 				assert.Equal(t, expectedError.HttpStatus, w.Code)
// 				assert.Equal(t, expectedError.Errors, response.Errors)
// 			})
// 		})

// 		t.Run("Given request body but not json", func(t *testing.T) {
// 			t.Run("Return bad request", func(t *testing.T) {
// 				apiEngine := gin.Default()

// 				applicationManager := mock.NewMockApplicationManager(mockCtrl)
// 				applicationManager.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)

// 				ctrl := controller.NewApplication().Create(applicationManager)
// 				apiEngine.Handle(ctrl.Method(), ctrl.Path(), ctrl.Control)

// 				expectedError := &entity.Error{
// 					HttpStatus: http.StatusBadRequest,
// 					Errors:     eobject.Wrap(eobject.BadRequestError("request body")),
// 				}

// 				var response ErrorResponse
// 				w := testhelper.PerformRequest(apiEngine, ctrl.Method(), ctrl.Path(), bytes.NewReader([]byte(`this is gonna be error`)))

// 				err := json.Unmarshal(w.Body.Bytes(), &response)
// 				assert.Nil(t, err)

// 				assert.Equal(t, expectedError.HttpStatus, w.Code)
// 				assert.Equal(t, expectedError.Errors, response.Errors)
// 			})
// 		})
// 	})
// }
