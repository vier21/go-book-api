package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	client *mongo.Client
}

func NewRepository() *UserRepository {
	if db.DB != nil {
		return &UserRepository{
			client: db.DB,
		}
	}
	client := db.NewConnection()
	return &UserRepository{
		client: client,
	}
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var result model.User

	coll := repo.client.Database("auth").Collection("user")
	filter := bson.M{
		"username": username,
	}

	doc := coll.FindOne(ctx, filter)
	if doc.Err() != nil {
		if doc.Err() == mongo.ErrNoDocuments {
			return model.User{}, doc.Err()
		}
	}

	if err := doc.Decode(&result); err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var result model.User

	coll := repo.client.Database("auth").Collection("user")
	filter := bson.M{
		"email": email,
	}

	doc := coll.FindOne(ctx, filter)
	if doc.Err() != nil {
		if doc.Err() == mongo.ErrNoDocuments {
			return model.User{}, doc.Err()
		}
	}

	if err := doc.Decode(&result); err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (repo *UserRepository) FindById(ctx context.Context, id string) (model.User, error) {
	var result model.User

	coll := repo.client.Database("auth").Collection("user")
	filter := bson.M{
		"_id": id,
	}

	doc := coll.FindOne(ctx, filter)
	if doc.Err() != nil {
		if doc.Err() == mongo.ErrNoDocuments {
			return model.User{}, doc.Err()
		}
	}

	if err := doc.Decode(&result); err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (repo *UserRepository) InsertUser(ctx context.Context, payload model.User) (model.User, error) {
	coll := repo.client.Database("auth").Collection("user")

	payload.Id = uuid.NewString()
	_, err := coll.InsertOne(ctx, payload)
	

	if err != nil {
		return model.User{}, err
	}

	return payload, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, id string, payload model.UpdateUser) (model.User, error) {
	coll := repo.client.Database("auth").Collection("user")
	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": payload,
	}

	option := options.FindOneAndUpdate().SetReturnDocument(1)
	res := coll.FindOneAndUpdate(ctx, filter, update, option)
	
	var updatedDoc model.User
	if err := res.Decode(&updatedDoc); err != nil {
		return model.User{}, err
	}

	fmt.Println("data succesfull updated")
	return updatedDoc, nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id string) error {

	coll := repo.client.Database("auth").Collection("user")
	filter := bson.M{
		"_id": id,
	}

	del, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if del.DeletedCount == 0 {
		return errors.New("Item might be deleted or no item with given ID")
	}

	fmt.Println("Success delete item")

	return nil
}

func (repo *UserRepository) BulkDelete(ctx context.Context, ids ...string) error {
	coll := repo.client.Database("auth").Collection("user")

	if len(ids) > 1 {

		models := []mongo.WriteModel{
			mongo.NewDeleteManyModel().SetFilter(bson.D{{"_id", bson.D{{"$in", ids}}}}),
		}

		del, err := coll.BulkWrite(ctx, models)
		if err != nil {
			return err
		}

		fmt.Println(del.DeletedCount)
	}

	err := repo.DeleteUser(ctx, ids[0])

	if err != nil {
		return err
	}

	return nil
}
