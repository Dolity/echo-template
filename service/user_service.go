package service

import (
	"context"
	_interface "myapp/interface"
	"myapp/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) _interface.UserRepository {
	return &UserService{
		collection: db.Collection("users"),
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	_, err := s.collection.InsertOne(ctx, user)
	return err
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


