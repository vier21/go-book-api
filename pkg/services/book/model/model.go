package model

type Book struct {
	Id        string `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	Slug      string `json:"slug" bson:"slug"`
	Body      string `json:"string" bson:"body"`
	Publisher string `json:"publisher" bson:"publisher"`
}
