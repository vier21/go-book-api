package model

type Book struct {
	Id        string `json:"id, omitempty"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Slug      string `json:"slug"`
	Body      string `json:"string"`
	Publisher string `json:"publisher"`
}
