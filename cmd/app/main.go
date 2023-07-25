package main

import (
	"fmt"
	"reflect"
)

func main() {

	// bookrep := repository.NewBookRepository()
	// bookSvc := service.NewBookService(bookrep)

	// ins, err := bookSvc.StoreBook(context.Background(), model.Book{
	// 	Title:     "booky",
	// 	Author:    "bambang",
	// 	Slug:      "book-a",
	// 	Body:      "ini adalah buku a",
	// 	Publisher: "gramedia",
	// 	Quantity:  1,
	// 	Price:     200000,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(ins)
	cobaReflect()
	// svc := service.NewUserService(repository.NewRepository())
	// defer db.Disconnect()
	// s := server.NewServer(svc)
	// s.Run()

}

type Coba struct {
	Name  string
	Place string
}

func cobaReflect() {
	c1 := Coba{
		Name:  "Agung",
		Place: "Bandung",
	}

	var c2 Coba

	fmt.Println(reflect.ValueOf(c1).IsZero())
	fmt.Println(reflect.ValueOf(c2).IsZero())

}
