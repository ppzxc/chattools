package mongodb

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ppzxc/chattools/storage/database"
	model2 "github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TopicInsert is call to require new room
func (m mongodb) TopicInsert(ctx context.Context, topic model2.Topic) (int64, error) {
	id, err := m.crudSequence.Next(ctx, database.MongoCollectionTopic)
	if err != nil {
		return 0, err
	}
	topic.Id = id

	err = m.crudTopic.InsertOne(ctx, topic)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m mongodb) TopicDelete(ctx context.Context, topicId int64) error {
	err := m.crudTopic.Delete(ctx, topicId)
	if err != nil {
		return err
	}

	return m.crudMessage.Delete(ctx, bson.D{{"topic_id", topicId}})
}

func (m mongodb) TopicDeleteByEmptySubs(ctx context.Context) error {
	all, err := m.crudTopic.FindManyFilter(ctx, bson.D{})
	if err != nil {
		return err
	}

	for _, a := range all {
		fmt.Printf("%+#v\n", a)
	}

	// topic empty
	if all == nil || len(all) <= 0 {
		return nil
	}

	for _, topic := range all {
		subs, err := m.crudSubs.FindManyByFilter(ctx, bson.M{"topic_id": topic.Id})
		if err != nil {
			return err
		}

		if err != nil {
			if err != mongo.ErrNoDocuments {
				err := m.crudTopic.Delete(ctx, topic.Id)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else if subs != nil || len(subs) <= 0 {
			err := m.crudTopic.Delete(ctx, topic.Id)
			if err != nil {
				return err
			}

			_ = m.crudMessage.Delete(ctx, bson.D{{"topic_id", topic.Id}})
		}
	}

	return nil
}

func (m mongodb) TopicFindAll(ctx context.Context, paging model2.Paging) ([]model2.Topic, error) {
	var filter bson.D
	//filter := bson.M{"updated_at": bson.M{"$gte": paging.UpdatedAt, "$lte": paging.CreatedAt}}

	if paging != (model2.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
	} else if paging != (model2.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"created_at", bson.M{"$gt": paging.CreatedAt}}}
	} else {
		filter = bson.D{}
	}

	return m.crudTopic.FindManyFilter(ctx, filter)
}

func (m mongodb) TopicFindAllByUserId(ctx context.Context, userId int64) ([]model2.Topic, error) {
	subs, err := m.crudSubs.FindManyByFilter(ctx, bson.M{"user_id": userId})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	var topicIds []int64
	for _, v := range subs {
		topicIds = append(topicIds, v.TopicId)
	}

	if len(topicIds) <= 0 {
		return nil, sql.ErrNoRows
	}

	return m.crudTopic.FindManyFilter(ctx, bson.D{{"_id", bson.M{"$in": topicIds}}})
}

func (m mongodb) TopicFindOneById(ctx context.Context, topicId int64) (model2.Topic, error) {
	return m.crudTopic.FindOneByFilter(ctx, bson.D{{"_id", topicId}})
}
