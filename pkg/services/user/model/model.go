package model

type User struct {
	Id       string `json:"id, omitempty" bson:"_id, omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}


