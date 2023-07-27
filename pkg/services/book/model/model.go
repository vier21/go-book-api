package model

type Book struct {
	Id        string  `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string  `json:"title" bson:"title"`
	Author    string  `json:"author" bson:"author"`
	Slug      string  `json:"slug" bson:"slug"`
	Body      string  `json:"body" bson:"body"`
	Publisher string  `json:"publisher" bson:"publisher"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"Price"`
}

type UpdateBook struct {
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
	Author    string `json:"author,omitempty" bson:"author,omitempty"`
	Slug      string `json:"slug,omitempty" bson:"slug,omitempty"`
	Body      string `json:"string,omitempty" bson:"body,omitempty"`
	Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Quantity  int     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty" bson:"Price,omitempty"`
}

func (b Book) Update() UpdateBook {
	return UpdateBook{
		Title:     b.Title,
		Author:    b.Author,
		Slug:      b.Slug,
		Body:      b.Body,
		Publisher: b.Publisher,
	}
}
