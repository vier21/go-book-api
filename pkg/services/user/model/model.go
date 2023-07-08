package model

type User struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

type UpdateUser struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}

type UpdatedUser struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

