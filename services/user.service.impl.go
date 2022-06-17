package services

import (
	"CrudGolang/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	//TODO implement me
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	//TODO implement me
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	//TODO implement me
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found is empty")
	}

	return users, nil
	//panic("implement me")
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	//TODO implement me
	query := bson.D{bson.E{Key: "name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "age", Value: user.Age}, bson.E{Key: "address", Value: user.Adrress}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, query, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update ")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	//TODO implement me
	query := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, query)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete ")
	}
	return nil
}
