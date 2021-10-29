package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (m mongodb) NotifyInsertMany(ctx context.Context, notify []*model.Notify) error {
	var many []interface{}
	for i := 0; i < len(notify); i++ {
		id, err := m.crudSequence.Next(ctx, database.MongoCollectionNotify)
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

func (m mongodb) NotifyFindAllByReceiveUserId(ctx context.Context, receiveUserId int64) ([]*model.Notify, error) {
	return m.crudNotify.FindManyFilter(ctx, bson.D{{"receive_user_id", receiveUserId}, {"deleted_at", bson.M{"$eq": nil}}})
}

func (m mongodb) NotifyFindOneReceiveUserIdById(ctx context.Context, notifyId int64) (int64, error) {
	filter, err := m.crudNotify.FindOneByFilter(ctx, bson.D{{"_id", notifyId}})
	if err != nil {
		return 0, err
	}
	return filter.ReceiveUserId, nil
}
