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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (b *BookRepository) BulkInsertBook(ctx context.Context, books []model.Book) ([]interface{}, error) {
	docs := make([]interface{}, len(books))

	for el := range docs {
		docs[el] = books[el]
	}

	result, err := b.collection.InsertMany(ctx, docs)
	if err != nil {
		return []any{}, err
	}

	return result.InsertedIDs, nil
}

func (b *BookRepository) UpdateBook(ctx context.Context, id string, book model.Book) (model.Book, error) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", book}}
	opt := options.FindOneAndUpdate().SetReturnDocument(1)

	res := b.collection.FindOneAndUpdate(ctx, filter, update, opt)
	var updated model.Book

	if res.Err() != nil {
		return model.Book{}, res.Err()
	}

	if err := res.Decode(&updated); err != nil {
		return model.Book{}, err
	}

	return updated, nil
}

func (b *BookRepository) DeleteBook(ctx context.Context, id string) error {
	del, err := b.collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	if err != nil {
		return err
	}

	if del.DeletedCount == 0 {
		return fmt.Errorf("cannot delete document with given id document deleted: %v", del.DeletedCount)
	}

	return nil
}

func (b *BookRepository) BulkDeleteBook(ctx context.Context, ids []string) (int, error) {
	deleteCount := 0
	if len(ids) > 1 {
		fmt.Println("masukk sini")

		models := []mongo.WriteModel{
			mongo.NewDeleteManyModel().SetFilter(bson.D{{"_id", bson.D{{"$in", ids}}}}),
		}

		res, err := b.collection.BulkWrite(ctx, models)
		if err != nil {
			fmt.Println(res)
			return deleteCount, err
		}
		deleteCount = int(res.DeletedCount)

	}
	return deleteCount, nil
}
