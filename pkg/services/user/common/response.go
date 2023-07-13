package common

type UpdateResponse struct {
	Status  string          `json:"status"`
	Payload UpdatePayload `json:"payload"`
}

type RegisterResponse struct {
	Status  string          `json:"status"`
	Payload RegisterPayload `json:"payload"`
}

type LoginResponse struct {
	Status  string       `json:"status"`
	Token   string       `json:"token`
	Payload LoginPayload `json:"data"`
}

type DeleteResponse struct {
	Status string
	Message any
}