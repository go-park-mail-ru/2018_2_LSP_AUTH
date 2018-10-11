package user

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

// User Structure that stores user information retrieved from database or
// entered by user during registration
type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	ID        int    `json:"id"`
	Token     string `json:"token"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Avatar    string `json:"avatar"`
}

// Auth Function that authenticates user
func (u *User) Auth(db *sql.DB, email string, password string) error {
	rows, err := db.Query("SELECT id, username, email, firstname, lastname, rating, avatar, password FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return err
	}

	defer rows.Close()
	if !rows.Next() {
		return errors.New("User not found")
	}

	var firstname sql.NullString
	var lastname sql.NullString
	var avatar sql.NullString
	err = rows.Scan(&u.ID, &u.Username, &u.Email, &firstname, &lastname, &u.Rating, &avatar)
	if err != nil {
		return err
	}
	if temp, err := firstname.Value(); temp != nil && err == nil {
		u.FirstName = temp.(string)
	}
	if temp, err := lastname.Value(); temp != nil && err == nil {
		u.LastName = temp.(string)
	}
	if temp, err := avatar.Value(); temp != nil && err == nil {
		u.Avatar = temp.(string)
	}

	if !validatePassword(u.Password, password) {
		return errors.New("Wrong password for user")
	}

	if err := u.generateToken(); err != nil {
		return err
	}

	return nil
}

func validatePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

func (u *User) generateToken() error {
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        u.ID,
		"generated": time.Now(),
	})
	u.Token, err = token.SignedString([]byte("HeAdfasdf3ref&^%$Dfrtgauyhia"))
	return err
}
