package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m mongodb) NotifyMaxIdByUserId(ctx context.Context, userId int64) (maxId int64, err error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)
	messages, err := m.crudNotify.FindManyFilter(ctx, bson.D{{"receive_user_id", userId}}, &findOptions)
	if err != nil {
		return 0, err
	}

	if len(messages) <= 0 {
		return 0, mongo.ErrNoDocuments
	}

	return messages[0].Id, nil
}

func (m mongodb) NotifyInsertOne(ctx context.Context, notify model.Notify) (int64, error) {
	id, err := m.crudSeq.Next(ctx, database.MongoCollectionNotify)
	if err != nil {
		return 0, err
	}
	notify.Id = id
	return m.crudNotify.InsertOne(ctx, notify)
}

func (m mongodb) NotifyInsertMany(ctx context.Context, notify []*model.Notify) error {
	var many []interface{}
	for i := 0; i < len(notify); i++ {
		id, err := m.crudSeq.Next(ctx, database.MongoCollectionNotify)
		if err != nil {
			return err
		}
		notify[i].Id = id
		many = append(many, notify[i])
	}

	return m.crudNotify.InsertMany(ctx, many)
}

func (m mongodb) NotifyUpdate(ctx context.Context, notify *model.Notify) error {
	return m.crudNotify.UpdateOne(ctx, notify)
}

func (m mongodb) NotifyFindAllByReceiveUserId(ctx context.Context, receiveUserId int64, paging model.Paging) ([]*model.Notify, error) {
	var filter bson.D
	if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"receive_user_id", receiveUserId}, {"updated_at", bson.M{"$gt": paging.UpdatedAt}}, {"deleted_at", bson.M{"$eq": nil}}}
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"receive_user_id", receiveUserId}, {"created_at", bson.M{"$gt": paging.CreatedAt}}, {"deleted_at", bson.M{"$eq": nil}}}
	} else if paging != (model.Paging{}) && paging.Id > 0 {
		filter = bson.D{{"receive_user_id", receiveUserId}, {"_id", bson.M{"$lt": paging.Id}}, {"deleted_at", bson.M{"$eq": nil}}}
	}

	return m.crudNotify.FindManyFilter(ctx, filter)
}

func (m mongodb) NotifyFindOneById(ctx context.Context, notifyId int64) (*model.Notify, error) {
	filter, err := m.crudNotify.FindOneByFilter(ctx, bson.D{{"_id", notifyId}})
	if err != nil {
		return &model.Notify{}, err
	}
	return filter, nil
}

func (m mongodb) NotifyFindOneReceiveUserIdById(ctx context.Context, notifyId int64) (int64, error) {
	filter, err := m.crudNotify.FindOneByFilter(ctx, bson.D{{"_id", notifyId}})
	if err != nil {
		return 0, err
	}
	return filter.ReceiveUserId, nil
}
