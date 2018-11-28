package handlers

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fields   string `json:"fields"`
}
