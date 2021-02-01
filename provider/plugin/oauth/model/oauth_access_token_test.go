package model_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/codefluence-x/altair/provider/plugin/oauth/entity"
	"github.com/codefluence-x/altair/provider/plugin/oauth/model"

	"github.com/stretchr/testify/assert"
)

var oauthAccessTokenModelRows = []string{
	"id",
	"oauth_application_id",
	"resource_owner_id",
	"token",
	"scopes",
	"expires_in",
	"created_at",
	"revoked_at",
}

func TestOauthAccessToken(t *testing.T) {

	t.Run("Name", func(t *testing.T) {
		t.Run("Return a model's name", func(t *testing.T) {
			db, _, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			assert.Equal(t, "oauth-access-token-model", model.NewOauthAccessToken(db).Name())
		})
	})

	t.Run("One", func(t *testing.T) {
		t.Run("Given oauth access token id", func(t *testing.T) {
			t.Run("Return entity oauth access token", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				data := entity.OauthAccessToken{
					ID: 1,
				}

				rows := sqlmock.NewRows(oauthAccessTokenModelRows).
					AddRow(
						data.ID,
						data.OauthApplicationID,
						data.ResourceOwnerID,
						data.Token,
						data.Scopes,
						data.ExpiresIn,
						data.CreatedAt,
						data.RevokedAT,
					)

				mockdb.ExpectQuery(`select id, oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at from oauth_access_tokens where id = \? and revoked_at is null limit 1`).
					WithArgs(1).
					WillReturnRows(rows)

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				dataFromDB, err := oauthAccessTokenModel.One(context.Background(), 1)

				assert.Nil(t, err)
				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Equal(t, data, dataFromDB)
			})

			t.Run("Row not found in database", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				mockdb.ExpectQuery(`select id, oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at from oauth_access_tokens where id = \? and revoked_at is null limit 1`).
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				dataFromDB, err := oauthAccessTokenModel.One(context.Background(), 1)

				assert.NotNil(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Equal(t, entity.OauthAccessToken{}, dataFromDB)
			})
		})
	})

	t.Run("OneByToken", func(t *testing.T) {
		t.Run("Given oauth access token", func(t *testing.T) {
			t.Run("Return entity oauth access token", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				data := entity.OauthAccessToken{
					ID: 1,
				}

				rows := sqlmock.NewRows(oauthAccessTokenModelRows).
					AddRow(
						data.ID,
						data.OauthApplicationID,
						data.ResourceOwnerID,
						data.Token,
						data.Scopes,
						data.ExpiresIn,
						data.CreatedAt,
						data.RevokedAT,
					)

				mockdb.ExpectQuery(`select id, oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at from oauth_access_tokens where token = \? and revoked_at is null limit 1`).
					WithArgs("token").
					WillReturnRows(rows)

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				dataFromDB, err := oauthAccessTokenModel.OneByToken(context.Background(), "token")

				assert.Nil(t, err)
				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Equal(t, data, dataFromDB)
			})

			t.Run("Row not found in database", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				mockdb.ExpectQuery(`select id, oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at from oauth_access_tokens where token = \? and revoked_at is null limit 1`).
					WithArgs("token").
					WillReturnError(sql.ErrNoRows)

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				dataFromDB, err := oauthAccessTokenModel.OneByToken(context.Background(), "token")

				assert.NotNil(t, err)
				assert.Equal(t, sql.ErrNoRows, err)
				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Equal(t, entity.OauthAccessToken{}, dataFromDB)
			})
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("Given context and access token insertable", func(t *testing.T) {
			t.Run("Return last inserted id", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				mockdb.ExpectExec(`insert into oauth_access_tokens \(oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at\) values\(\?, \?, \?, \?, \?, now\(\), null\)`).
					WithArgs(insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn).
					WillReturnResult(sqlmock.NewResult(1, 1))

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				lastInsertedID, err := oauthAccessTokenModel.Create(context.Background(), insertable)

				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Nil(t, err)
				assert.Equal(t, 1, lastInsertedID)
			})

			t.Run("Unexpected error", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				mockdb.ExpectExec(`insert into oauth_access_tokens \(oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at\) values\(\?, \?, \?, \?, \?, now\(\), null\)`).
					WithArgs(insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn).
					WillReturnError(errors.New("unexpected error"))

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				lastInsertedID, err := oauthAccessTokenModel.Create(context.Background(), insertable)

				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.NotNil(t, err)
				assert.Equal(t, 0, lastInsertedID)
			})

			t.Run("Get last inserted id error", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				mockdb.ExpectExec(`insert into oauth_access_tokens \(oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at\) values\(\?, \?, \?, \?, \?, now\(\), null\)`).
					WithArgs(insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("unexpected error")))

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				lastInsertedID, err := oauthAccessTokenModel.Create(context.Background(), insertable)

				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.NotNil(t, err)
				assert.Equal(t, 0, lastInsertedID)
			})
		})

		t.Run("Given context, access token insertable and database transaction", func(t *testing.T) {
			t.Run("Return last inserted id", func(t *testing.T) {
				db, mockdb, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				mockdb.ExpectBegin()
				mockdb.ExpectExec(`insert into oauth_access_tokens \(oauth_application_id, resource_owner_id, token, scopes, expires_in, created_at, revoked_at\) values\(\?, \?, \?, \?, \?, now\(\), null\)`).
					WithArgs(insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn).
					WillReturnResult(sqlmock.NewResult(1, 1))

				tx, _ := db.Begin()

				oauthAccessTokenModel := model.NewOauthAccessToken(db)
				lastInsertedID, err := oauthAccessTokenModel.Create(context.Background(), insertable, tx)

				assert.Nil(t, mockdb.ExpectationsWereMet())
				assert.Nil(t, err)
				assert.Equal(t, 1, lastInsertedID)
			})
		})
	})

	t.Run("Revoke", func(t *testing.T) {
		t.Run("Given context and token", func(t *testing.T) {
			t.Run("Run gracefully", func(t *testing.T) {
				t.Run("Return nil", func(t *testing.T) {
					db, mockdb, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mockdb.ExpectExec(`update oauth_access_tokens set revoked_at = now\(\) where token = \?`).
						WithArgs("token").WillReturnResult(sqlmock.NewResult(1, 1))

					oauthAccessTokenModel := model.NewOauthAccessToken(db)
					err = oauthAccessTokenModel.Revoke(context.Background(), "token")
					assert.Nil(t, err)
				})
			})

			t.Run("Execution error", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					db, mockdb, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mockdb.ExpectExec(`update oauth_access_tokens set revoked_at = now\(\) where token = \?`).
						WithArgs("token").WillReturnError(errors.New("unexpected error"))

					oauthAccessTokenModel := model.NewOauthAccessToken(db)
					err = oauthAccessTokenModel.Revoke(context.Background(), "token")
					assert.NotNil(t, err)
				})
			})

			t.Run("Get rows affected error", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					db, mockdb, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mockdb.ExpectExec(`update oauth_access_tokens set revoked_at = now\(\) where token = \?`).
						WithArgs("token").WillReturnResult(sqlmock.NewErrorResult(errors.New("unexpected error")))

					oauthAccessTokenModel := model.NewOauthAccessToken(db)
					err = oauthAccessTokenModel.Revoke(context.Background(), "token")
					assert.NotNil(t, err)
				})
			})

			t.Run("No rows affected", func(t *testing.T) {
				t.Run("Return sql.ErrNoRows error", func(t *testing.T) {
					db, mockdb, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mockdb.ExpectExec(`update oauth_access_tokens set revoked_at = now\(\) where token = \?`).
						WithArgs("token").WillReturnResult(sqlmock.NewResult(1, 0))

					oauthAccessTokenModel := model.NewOauthAccessToken(db)
					err = oauthAccessTokenModel.Revoke(context.Background(), "token")
					assert.NotNil(t, err)
					assert.Equal(t, sql.ErrNoRows, err)
				})
			})
		})
	})
}
