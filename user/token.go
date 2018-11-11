package user

// Token - token fields
type Token struct {
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Domain    string   `json:"domain"`
	Roles     []string `json:"roles"`
}
