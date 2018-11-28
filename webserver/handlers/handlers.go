package handlers

import (
	cnt "context"
	"net/http"

	"github.com/go-park-mail-ru/2018_2_LSP_AUTH_GRPC/auth_proto"
	"github.com/go-park-mail-ru/2018_2_LSP_USER_GRPC/user_proto"
	"github.com/thedevsaddam/govalidator"
	"golang.org/x/crypto/bcrypt"
)

// ClearCookiesHandler removes cookies
func ClearCookiesHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	removeAuthCookies(w, r)
	return nil
}

// AuthUserHandler sets cookies for user provided with valid credentials
func AuthUserHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	payload := authPayload{}
	opts := govalidator.Options{
		Request: r,
		Data:    &payload,
		Rules:   authValidationRules,
	}
	v := govalidator.New(opts)
	if e := v.ValidateJSON(); len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return StatusData{http.StatusBadRequest, err}
	}

	ctx := cnt.Background()
	userManager := user_proto.NewUserCheckerClient(env.GRCPUser)
	u, err := userManager.GetOneByEmail(ctx,
		&user_proto.UserEmail{
			Email: payload.Email,
		})
	if err := handleGetOneUserGrpcError(env, err); err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password))
	if err != nil {
		return StatusData{
			Code: http.StatusUnauthorized,
			Data: map[string]string{
				"error": "Wrong user password",
			},
		}
	}

	authManager := auth_proto.NewAuthCheckerClient(env.GRCPAuth)
	token, err := authManager.Generate(ctx,
		&auth_proto.TokenPayload{
			ID: int64(u.ID),
		})
	if err := handleGeneralGrpcError(env, err); err != nil {
		return err
	}

	setAuthCookies(w, token.Token)

	return nil
}
