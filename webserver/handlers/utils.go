package handlers

import (
	"net/http"
	"strings"
	"time"
)

func setAuthCookies(w http.ResponseWriter, tokenString string) {
	firstDot := strings.Index(tokenString, ".") + 1
	secondDot := strings.Index(tokenString[firstDot:], ".") + firstDot
	cookieHeaderPayload := http.Cookie{
		Name:    "header.payload",
		Value:   tokenString[:secondDot],
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
		Domain:  ".jackal.online",
		Path:    "/",
	}
	cookieSignature := http.Cookie{
		Name:     "signature",
		Value:    tokenString[secondDot+1:],
		Expires:  time.Now().Add(720 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Domain:   ".jackal.online",
		Path:     "/",
	}
	http.SetCookie(w, &cookieHeaderPayload)
	http.SetCookie(w, &cookieSignature)
}

func removeAuthCookies(w http.ResponseWriter, r *http.Request) {
	signature, err := r.Cookie("signature")
	if err == nil {
		signature.Expires = time.Now().AddDate(-10, 0, 0)
		http.SetCookie(w, signature)
	}

	headerPayload, err := r.Cookie("header.payload")
	if err == nil {
		headerPayload.Expires = time.Now().AddDate(-10, 0, 0)
		http.SetCookie(w, headerPayload)
	}
}
