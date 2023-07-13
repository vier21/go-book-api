package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/book/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
	if db.DB != nil {
		return &BookRepository{
			client:     db.DB,
			collection: db.DB.Database(config.GetConfig().BookDBName).Collection("books"),
		}
	}
	cli := db.NewConnection()
	return &BookRepository{
		client:     cli,
		collection: cli.Database(config.GetConfig().BookDBName).Collection("books"),
	}
}

func (b *BookRepository) FindById(ctx context.Context, id string) (model.Book, error) {
	cur := b.collection.FindOne(ctx, bson.M{
		"_id": id,
	})
	if cur.Err() != nil {
		return model.Book{}, cur.Err()
	}

	var result model.Book

	if err := cur.Decode(&result); err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func (b *BookRepository) FindByName(ctx context.Context, name string) (model.Book, error) {
	cur := b.collection.FindOne(ctx, bson.D{{"name", name}})
	if cur.Err() != nil {
		return model.Book{}, cur.Err()
	}

	var result model.Book

	if err := cur.Decode(&result); err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func (b *BookRepository) FindBySlug(ctx context.Context, slug string) (model.Book, error) {
	cur := b.collection.FindOne(ctx, bson.D{{"slug", slug}})
	if cur.Err() != nil {
		return model.Book{}, cur.Err()
	}

	var result model.Book

	if err := cur.Decode(&result); err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func (b *BookRepository) FindAll(ctx context.Context) ([]model.Book, error) {
	var result []model.Book
	cur, err := b.collection.Find(ctx, bson.D{})

	if err != nil {
		return []model.Book{}, err
	}

	if err := cur.All(ctx, &result); err != nil {
		return []model.Book{}, err
	}
	return result, nil
}

func (b *BookRepository) InsertBook(ctx context.Context, book model.Book) (model.Book, error) {
	id := uuid.NewString()
	book.Id = id
	doc, err := b.collection.InsertOne(ctx, book)
	if err != nil {
		return model.Book{}, err
	}

	if doc.InsertedID.(string) != id {
		return model.Book{}, fmt.Errorf("document not contain expected id %s", id)
	}

	return book, nil
}

func (b *BookRepository) UpdateBook(ctx context.Context, id string, book model.Book) (model.Book, error) {
	
	return model.Book{}, nil
}

func (b *BookRepository) DeleteBook(ctx context.Context, id string) error {
	return nil
}
