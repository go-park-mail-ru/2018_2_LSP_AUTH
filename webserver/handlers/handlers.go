package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2018_2_LSP_AUTH/user"
	"github.com/thedevsaddam/govalidator"
)

func init() {
	govalidator.AddCustomRule("fields", func(field string, rule string, message string, value interface{}) error {
		fields := strings.Split(value.(string), ",")
		if len(fields) == 0 {
			return errors.New("Field keyword should be field list divided by comma. Available fields: " + strings.TrimPrefix(rule, "fields:"))
		}
		fieldListStr := strings.TrimPrefix(rule, "fields:")
		fieldListSlice := strings.Split(fieldListStr, ",")
		for _, f := range fields {
			if !contains(fieldListSlice, f) {
				return errors.New("Field keyword should be field list divided by comma. Available fields: " + strings.TrimPrefix(rule, "fields:"))
			}
		}
		return nil
	})
}

func DeleteHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	signature, err := r.Cookie("signature")
	if err == nil {
		signature.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, signature)
	}

	headerPayload, err := r.Cookie("header.payload")
	if err == nil {
		headerPayload.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, headerPayload)
	}

	return nil
}

func PostHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Fields   string `json:"fields"`
	}{"", "", ""}

	rules := govalidator.MapData{
		"email":    []string{"required", "between:4,25", "email"},
		"password": []string{"required", "alpha_space"},
		"fields":   []string{"fields:username,email,firstname,lastname,rating,id,avatar"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &payload,
		Rules:   rules,
	}
	v := govalidator.New(opts)
	if e := v.ValidateJSON(); len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return StatusData{http.StatusBadRequest, err}
	}

	var u user.User
	if err := u.Auth(env.DB, payload.Email, payload.Password); err != nil {
		return StatusData{http.StatusBadRequest, map[string]string{"error": err.Error()}}
	}

	setAuthCookies(w, u.Token)

	fieldsToReturn := strings.Split(payload.Fields, ",")
	answer := extractFields(u, fieldsToReturn)

	return StatusData{http.StatusOK, answer}
}

func extractFields(u user.User, fieldsToReturn []string) map[string]string {
	answer := map[string]string{}
	for _, f := range fieldsToReturn {
		switch f {
		case "firstname":
			answer["firstname"] = u.FirstName
		case "lastname":
			answer["lastname"] = u.LastName
		case "email":
			answer["email"] = u.Email
		case "username":
			answer["username"] = u.Username
		}
	}
	return answer
}
