package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2018_2_LSP_USER_GRPC/user_proto"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func extractFields(u user_proto.User, fieldsToReturn []string) map[string]interface{} {
	answer := map[string]interface{}{}
	for _, f := range fieldsToReturn {
		switch f {
		case "id":
			answer["id"] = u.ID
		case "firstname":
			answer["firstname"] = u.FirstName
		case "lastname":
			answer["lastname"] = u.LastName
		case "email":
			answer["email"] = u.Email
		case "username":
			answer["username"] = u.Username
		case "avatar":
			answer["avatar"] = u.Avatar
		}
	}
	return answer
}

func setAuthCookies(w http.ResponseWriter, tokenString string) {
	firstDot := strings.Index(tokenString, ".") + 1
	secondDot := strings.Index(tokenString[firstDot:], ".") + firstDot
	cookieHeaderPayload := http.Cookie{
		Name:    "header.payload",
		Value:   tokenString[:secondDot],
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
		Domain:  ".jackal.online",
	}
	cookieSignature := http.Cookie{
		Name:     "signature",
		Value:    tokenString[secondDot+1:],
		Expires:  time.Now().Add(720 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Domain:   ".jackal.online",
	}
	http.SetCookie(w, &cookieHeaderPayload)
	http.SetCookie(w, &cookieSignature)
}

func removeAuthCookies(w http.ResponseWriter, r *http.Request) {
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
}
