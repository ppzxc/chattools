package mongodb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m mongodb) SubscriptionExistsByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (int64, error) {
	filter, err := m.crudSubs.FindOneByFilter(ctx, bson.D{{"user_id", userId}, {"topic_id", topicId}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}
	return filter.Id, nil
}

func (m mongodb) SubscriptionFindAllByUserId(ctx context.Context, userId int64) (subscriptions []model.Subscription, err error) {
	return m.crudSubs.FindManyByFilter(ctx, bson.D{{"user_id", userId}})
}

func (m mongodb) SubscriptionFindAllByTopicId(ctx context.Context, topicId int64) ([]model.Subscription, error) {
	return m.crudSubs.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}})
}

func (m mongodb) SubscriptionFindOneByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (model.Subscription, error) {
	return m.crudSubs.FindOneByFilter(ctx, bson.D{{"user_id", userId}, {"topic_id", topicId}})
}

func (m mongodb) SubscriptionInsert(ctx context.Context, subscription model.Subscription) (int64, error) {
	id, err := m.crudSequence.Next(ctx, database.MongoCollectionSubscriptions)
	if err != nil {
		return 0, err
	}
	subscription.Id = id

	return id, m.crudSubs.InsertOne(ctx, subscription)
}

func (m mongodb) SubscriptionDeleteByTopicIdAndUserId(ctx context.Context, topicId int64, userId int64) error {
	return m.crudSubs.DeleteOneByFilter(ctx, bson.D{{"user_id", userId}, {"topic_id", topicId}})
}

func (m mongodb) SubscriptionDeleteAllByUserId(ctx context.Context, userId int64) error {
	err := m.crudSubs.DeleteAllByFilter(ctx, bson.D{{"user_id", userId}})
	if err != nil && err == mongo.ErrNoDocuments {
		return sql.ErrNoRows
	} else if err != nil {
		return err
	} else {
		return nil
	}
}

func (m mongodb) SubscriptionUpdateAck(ctx context.Context, subscription model.Subscription) error {
	var update bson.D
	if subscription.ReceiveSequenceId > 0 {
		update = bson.D{{"$set", bson.D{
			{"receive_sequence_id", subscription.ReceiveSequenceId},
			{"updated_at", subscription.UpdatedAt},
		}}}
	} else if subscription.ReadSequenceId > 0 {
		update = bson.D{{"$set", bson.D{
			{"read_sequence_id", subscription.ReadSequenceId},
			{"updated_at", subscription.UpdatedAt},
		}}}
	} else {
		update = bson.D{{"$set", bson.D{{"updated_at", subscription.UpdatedAt}}}}
	}

	// where sequence_id is not null
	if subscription.ReceiveSequenceId > 0 {
		return m.crudSubs.UpdateOneByFilter(
			ctx,
			bson.M{
				"user_id":     subscription.UserId,
				"topic_id":    subscription.TopicId,
				"sequence_id": bson.D{{"$lt", subscription.ReceiveSequenceId}},
			},
			update)
	} else if subscription.ReadSequenceId > 0 {
		return m.crudSubs.UpdateOneByFilter(
			ctx,
			bson.M{
				"user_id":     subscription.UserId,
				"topic_id":    subscription.TopicId,
				"sequence_id": bson.D{{"$lt", subscription.ReadSequenceId}},
			},
			update)
		//return m.crudSubs.UpdateOneByFilter(
		//	ctx,
		//	bson.D{
		//		{"user_id", subscription.UserId},
		//		{"topic_id", subscription.TopicId},
		//		{"sequence_id", bson.D{{"$lt", subscription.ReadSequenceId}}},
		//	},
		//	update,
		//)
	} else {
		return m.crudSubs.UpdateOneByFilter(
			ctx,
			bson.M{
				"user_id":  subscription.UserId,
				"topic_id": subscription.TopicId,
			},
			update)
		//return m.crudSubs.UpdateOneByFilter(
		//	ctx,
		//	bson.D{
		//		{"user_id", subscription.UserId},
		//		{"topic_id", subscription.TopicId},
		//	},
		//	update,
		//)
	}
}
