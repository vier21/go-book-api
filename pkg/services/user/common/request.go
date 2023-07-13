package common



type RegisterPayload struct {
	Id       string `json:"id"`
	Username string `json:username`
	Email    string `json:email`
}


type UpdatePayload struct {
	Id       string `json:"id"`
	Username string `json:username`
	Email    string `json:email`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Id       string `json:"id"`
	Username string `json:username`
	Email    string `json:email`
}


