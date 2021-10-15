package repository

import (
	"context"
	"github.com/ppzxc/chattools/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

type User interface {
	FindOneByFilter(ctx context.Context, filter bson.D) (model.User, error)
	FindOneAndDelete(ctx context.Context, userId int64) (*model.User, error)
	FindManyByFilter(ctx context.Context, filter bson.D) ([]model.User, error)
	InsertMany(ctx context.Context, many []interface{}) error
	InsertOne(ctx context.Context, user model.User) error
	FindOneAndUpdateByFilter(ctx context.Context, filter bson.D, update bson.D) error
	Delete(ctx context.Context, userId int64) error
}

type Topic interface {
	FindOneAndUpdateByFilter(ctx context.Context, filter bson.D, update bson.D) error
	FindManyFilter(ctx context.Context, filter bson.D) ([]model.Topic, error)
	FindOneByFilter(ctx context.Context, filter bson.D) (model.Topic, error)
	InsertOne(ctx context.Context, topic model.Topic) error
	UpdateFilter(ctx context.Context, filter bson.D, update bson.D) error
	Update(ctx context.Context, topic *model.Topic) error
	Delete(ctx context.Context, topicId int64) error
}

type Subscription interface {
	FindOneByFilter(ctx context.Context, filter interface{}) (model.Subscription, error)
	FindManyByFilter(ctx context.Context, filter interface{}) ([]model.Subscription, error)
	InsertOne(ctx context.Context, subs model.Subscription) error
	DeleteAllByFilter(ctx context.Context, filter interface{}) error
	DeleteOneByFilter(ctx context.Context, filter interface{}) error
	UpdateOneByFilter(ctx context.Context, filter interface{}, update interface{}) error
}

type Notify interface {
	FindOneByFilter(ctx context.Context, filter bson.D) (*model.Notify, error)
	FindManyFilter(ctx context.Context, filter bson.D) ([]*model.Notify, error)
	InsertMany(ctx context.Context, many []interface{}) error
	UpdateOne(ctx context.Context, notify *model.Notify) error
}

type Sequence interface {
	TopicMaxSeq(collectionName string, topicId int64) string
	Current(ctx context.Context, collectionName string, topicId int64) (int64, error)
	Next(ctx context.Context, collectionName string) (int64, error)
}

type File interface {
	FindOneByFilter(ctx context.Context, filter bson.D) (model.File, error)
	InsertOne(ctx context.Context, file model.File) error
}

type Message interface {
	InsertOne(ctx context.Context, message model.Message) error
	InsertMany(ctx context.Context, messages []interface{}) error
	FindManyByFilter(ctx context.Context, filter bson.D) ([]model.Message, error)
	Delete(ctx context.Context, filter bson.D) error
}

//type DeviceRepository interface {
//	FindAllByUserId(int64) ([]model.Device, error)
//	Add(model.Device) error
//	DeleteByUserId(int64) error
//}
//
//type AuthenticationRepository interface {
//	FindOneByUserId(int64) (model.Authentication, error)
//	FindOneByEmail(string) (model.Authentication, error)
//	Add(model.Authentication) error
//	TxDeleteByUserId(int64) error
//}
//
//type ProfileRepository interface {
//	FindAll() ([]*model.Profile, error)
//	FindOneById(int64) (*model.Profile, error)
//	FindOneByUserId(int64) (*model.Profile, error)
//
//	Add(*model.Profile) error
//	Update(*model.Profile) error
//	Delete(int64) error
//}

