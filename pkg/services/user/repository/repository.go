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
)

type UserRepository struct {
	UserDB *db.Database
}

type UpdateUser struct {
	Username string `json:"username, omitempty" bson:"username, omitempty"`
	Password string `json:"password, omitempty" bson:"password, omitempty"`
	Email    string `json:"email, omitempty" bson:"email, omitempty"`
}

func UpsertUser(u model.User) UpdateUser {
	return UpdateUser{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}

func NewRepository(userdb *db.Database) *UserRepository {
	return &UserRepository{
		UserDB: userdb,
	}
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var result model.User

	coll := repo.UserDB.Client.Database("auth").Collection("user")
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

func (repo *UserRepository) FindById(ctx context.Context, id string) (model.User, error) {
	var result model.User

	coll := repo.UserDB.Client.Database("auth").Collection("user")
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
	coll := repo.UserDB.Client.Database("auth").Collection("user")

	payload.Id = uuid.NewString()
	_, err := coll.InsertOne(ctx, payload)

	if err != nil {
		return model.User{}, err
	}

	return payload, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, payload model.User) (model.User, error) {
	coll := repo.UserDB.Client.Database("auth").Collection("user")

	filter := bson.M{
		"_id": payload.Id,
	}

	update := bson.M{
		"$set": UpsertUser(payload),
	}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return model.User{}, err
	}

	if res.MatchedCount == 0 {
		return model.User{}, errors.New("failed Update because no id is matched")
	}

	fmt.Println("data succesfull updated")

	return payload, nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, ids ...string) error {

	coll := repo.UserDB.Client.Database("auth").Collection("user")
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

	filter := bson.M{
		"_id": ids[0],
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
