package handler

type LoginResponse struct {
	Token string `json:"token"`
}

func ToLoginReponse(tkn string) LoginResponse {
	return LoginResponse{
		Token: tkn,
	}
}
