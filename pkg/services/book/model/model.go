package model

type Book struct {
	Id        string `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	Slug      string `json:"slug" bson:"slug"`
	Body      string `json:"string" bson:"body"`
	Publisher string `json:"publisher" bson:"publisher"`
}

type UpdateBook struct {
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
	Author    string `json:"author,omitempty" bson:"author,omitempty"`
	Slug      string `json:"slug,omitempty" bson:"slug,omitempty"`
	Body      string `json:"string,omitempty" bson:"body,omitempty"`
	Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
}
