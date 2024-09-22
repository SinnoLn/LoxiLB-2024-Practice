package api

type Account struct {
	CommonAPI
}

type AccountModGet struct {
	Attr []AccountGet `json:"Attr"`
}

type AccountGet struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
